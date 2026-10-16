import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { PatchProfileRequest, Profile } from '@/api'
import { ResponseError } from '@/api'
import { formatStartValue, sortEntries } from '@/utils/profileSchedule'

export type SaveStatus = 'saving' | 'error'

const PATCH_DEBOUNCE_MS = 500

const profilesMap = ref<Map<number, Profile>>(new Map())
const statusMap = ref<Map<number, SaveStatus>>(new Map())
const lastError = ref<string | null>(null)
// Increments after every successful create/patch/delete so views can refresh
// derived state (e.g. the device's effective target).
const saveTick = ref(0)

// Per-entry optimistic version so late responses cannot clobber newer edits.
const versions = new Map<number, number>()
// Pending (debounced) patch payloads, merged per entry.
const pendingChanges = new Map<number, PatchProfileRequest>()
const flushTimers = new Map<number, ReturnType<typeof setTimeout>>()

let tempIdCounter = -1

// Increments on every list fetch; stale responses never overwrite fresher state
// (e.g. rapid tab switches where an earlier request finishes last).
let loadGeneration = 0

function nextTempId(): number {
  return tempIdCounter--
}

function setMapEntry<K, V>(target: { value: Map<K, V> }, key: K, value: V) {
  target.value = new Map(target.value).set(key, value)
}

function deleteMapEntry<K, V>(target: { value: Map<K, V> }, key: K) {
  if (!target.value.has(key)) return
  const next = new Map(target.value)
  next.delete(key)
  target.value = next
}

function upsert(profile: Profile) {
  setMapEntry(profilesMap, profile.id, profile)
}

function removeLocal(id: number) {
  deleteMapEntry(profilesMap, id)
}

function setStatus(id: number, status: SaveStatus | null) {
  if (status) {
    setMapEntry(statusMap, id, status)
  } else {
    deleteMapEntry(statusMap, id)
  }
}

function isNotFound(error: unknown): boolean {
  return error instanceof ResponseError && error.response?.status === 404
}

function bumpVersion(id: number): number {
  const version = (versions.get(id) ?? 0) + 1
  versions.set(id, version)
  return version
}

function applyChanges(profile: Profile, changes: PatchProfileRequest): Profile {
  const next: Profile = { ...profile }
  if (changes.start !== undefined) {
    next.start = changes.start
  }
  if ('targetTemp' in changes) {
    next.targetTemp = changes.targetTemp ?? null
  }
  return next
}

const allProfiles = computed<Profile[]>(() => Array.from(profilesMap.value.values()))

export interface ProfileSliceFilter {
  control?: number
  daytype?: number
}

export function useProfiles() {
  async function loadProfiles() {
    const generation = ++loadGeneration
    try {
      const result = await api.listProfile()
      if (generation !== loadGeneration) return
      const items: Profile[] = result.objects ?? []
      const newMap = new Map<number, Profile>()
      for (const profile of items) {
        newMap.set(profile.id, profile)
      }
      profilesMap.value = newMap
      statusMap.value = new Map()
      pendingChanges.clear()
      versions.clear()
      for (const timer of flushTimers.values()) {
        clearTimeout(timer)
      }
      flushTimers.clear()
      lastError.value = null
    } catch (error) {
      if (generation !== loadGeneration) return
      console.error('Failed to load profiles:', error)
      lastError.value = 'Failed to load schedules.'
    }
  }

  /**
   * Load one filtered slice (e.g. a single sensor or day type) via the
   * `?control=` / `?daytype=` API filters, replacing local state.
   * Nothing else is cached, so callers refetch on every selection change.
   */
  async function loadEntriesFor(filter: ProfileSliceFilter) {
    const generation = ++loadGeneration
    try {
      const result = await api.listProfile({ ...filter })
      if (generation !== loadGeneration) return
      const items: Profile[] = result.objects ?? []
      const newMap = new Map<number, Profile>()
      for (const profile of items) {
        newMap.set(profile.id, profile)
      }
      profilesMap.value = newMap
      statusMap.value = new Map()
      // Drop optimistic versions for entries evicted by the slice replace so
      // a returning id never inherits a stale version guard.
      for (const id of Array.from(versions.keys())) {
        if (!newMap.has(id)) versions.delete(id)
      }
      lastError.value = null
    } catch (error) {
      if (generation !== loadGeneration) return
      console.error('Failed to load profiles:', error)
      lastError.value = 'Failed to load schedules.'
    }
  }

  function entriesFor(controlUri: string, daytypeUri: string): Profile[] {
    return sortEntries(
      allProfiles.value.filter(
        (profile) => profile.control === controlUri && profile.daytype === daytypeUri,
      ),
    )
  }

  function statusFor(id: number): SaveStatus | null {
    return statusMap.value.get(id) ?? null
  }

  async function resyncEntry(id: number) {
    bumpVersion(id)
    try {
      const fresh = await api.getProfile({ id })
      upsert(fresh)
    } catch (error) {
      // Only drop the row when the server confirms it is gone. A network
      // blip must not look like a delete.
      if (isNotFound(error)) {
        removeLocal(id)
        setStatus(id, null)
      }
    }
  }

  async function flushPatch(id: number) {
    const timer = flushTimers.get(id)
    if (timer) {
      clearTimeout(timer)
      flushTimers.delete(id)
    }

    const changes = pendingChanges.get(id)
    pendingChanges.delete(id)
    if (!changes || Object.keys(changes).length === 0) return

    const version = versions.get(id) ?? 0
    setStatus(id, 'saving')

    try {
      const updated = await api.patchProfile({ id, patchProfileRequest: changes })
      // Only apply if no newer edit landed while the request was in flight.
      if ((versions.get(id) ?? 0) === version) {
        const existing = profilesMap.value.get(id)
        upsert({ ...updated, resourceUri: updated.resourceUri || existing?.resourceUri || '' })
        setStatus(id, null)
        saveTick.value++
      }
    } catch (error) {
      console.error('Failed to save profile change:', error)
      lastError.value = 'Failed to save schedule change.'
      setStatus(id, 'error')
      await resyncEntry(id)
    }
  }

  function schedulePatch(id: number, changes: PatchProfileRequest, immediate: boolean) {
    const current = profilesMap.value.get(id)
    if (!current) return

    upsert(applyChanges(current, changes))
    bumpVersion(id)
    pendingChanges.set(id, { ...pendingChanges.get(id), ...changes })

    const existing = flushTimers.get(id)
    if (existing) {
      clearTimeout(existing)
      flushTimers.delete(id)
    }

    if (immediate) {
      void flushPatch(id)
    } else {
      flushTimers.set(
        id,
        setTimeout(() => void flushPatch(id), PATCH_DEBOUNCE_MS),
      )
    }
  }

  async function addEntry(
    controlUri: string,
    daytypeUri: string,
    startSeconds: number,
    targetTemp: number | null,
  ): Promise<Profile | null> {
    const tempId = nextTempId()
    const optimistic: Profile = {
      resourceUri: '',
      id: tempId,
      control: controlUri,
      daytype: daytypeUri,
      start: formatStartValue(startSeconds),
      targetTemp,
    }
    upsert(optimistic)
    setStatus(tempId, 'saving')

    try {
      const created = await api.createProfile({
        profileW: {
          control: controlUri,
          daytype: daytypeUri,
          start: formatStartValue(startSeconds),
          targetTemp,
        },
      })
      setStatus(tempId, null)

      if (profilesMap.value.has(tempId)) {
        removeLocal(tempId)
        upsert(created)
        saveTick.value++
      } else {
        // The optimistic row was removed while the create was in flight.
        try {
          await api.deleteProfile({ id: created.id })
        } catch {
          /* ignore cleanup failure */
        }
      }
      return created
    } catch (error) {
      console.error('Failed to add profile entry:', error)
      lastError.value = 'Failed to add time point.'
      removeLocal(tempId)
      setStatus(tempId, null)
      return null
    }
  }

  function setEntryStart(id: number, startSeconds: number) {
    schedulePatch(id, { start: formatStartValue(startSeconds) }, true)
  }

  function setEntryTemp(id: number, targetTemp: number) {
    schedulePatch(id, { targetTemp }, false)
  }

  function setEntryDefault(id: number) {
    schedulePatch(id, { targetTemp: null }, true)
  }

  async function removeEntry(id: number) {
    const current = profilesMap.value.get(id)
    if (!current) return

    const timer = flushTimers.get(id)
    if (timer) {
      clearTimeout(timer)
      flushTimers.delete(id)
    }
    pendingChanges.delete(id)
    removeLocal(id)

    if (id < 0) {
      // Never persisted (optimistic create still in flight); nothing to delete.
      setStatus(id, null)
      return
    }

    setStatus(id, 'saving')
    try {
      await api.deleteProfile({ id })
      setStatus(id, null)
      saveTick.value++
    } catch (error) {
      console.error('Failed to remove profile entry:', error)
      lastError.value = 'Failed to remove time point.'
      upsert(current)
      setStatus(id, 'error')
    }
  }

  function clearError() {
    lastError.value = null
  }

  /**
   * Drop all local entries (e.g. before a refetch) so stale slices are never
   * shown. Call only after `flushAll()` — pending edits live in the map.
   */
  function clearEntries() {
    profilesMap.value = new Map()
    statusMap.value = new Map()
  }

  /** Immediately flush all debounced patches (e.g. before unload/unmount). */
  async function flushAll() {
    const ids = Array.from(pendingChanges.keys())
    await Promise.all(ids.map((id) => flushPatch(id)))
  }

  return {
    allProfiles,
    lastError,
    saveTick,
    clearError,
    clearEntries,
    loadProfiles,
    loadEntriesFor,
    entriesFor,
    statusFor,
    addEntry,
    setEntryStart,
    setEntryTemp,
    setEntryDefault,
    removeEntry,
    flushAll,
    flushPatch,
  }
}
