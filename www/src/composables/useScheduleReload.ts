import { computed, ref } from 'vue'
import { useControls } from '@/composables/useControls'
import { useDayTypes } from '@/composables/useDayTypes'
import { useProfiles } from '@/composables/useProfiles'
import { readStoredString, writeStoredString } from '@/composables/usePersistentRef'

const pendingReloads = ref(0)

/** True while a schedule reload is in flight. */
const isReloading = computed(() => pendingReloads.value > 0)

// Storage keys mirror the persisted selection in useProfileAxis,
// ProfileBySensor and ProfileByDayType.
const AXIS_KEY = 'profile_group_axis'
const SENSOR_KEY = 'profile_sensor'
const DAYTYPE_KEY = 'profile_daytype'

function readStoredId(key: string): number | null {
  const raw = readStoredString(key)
  if (!raw) return null
  const parsed = Number(raw)
  return Number.isFinite(parsed) ? parsed : null
}

function readStoredAxis(): 'sensor' | 'daytype' {
  return readStoredString(AXIS_KEY) === 'daytype' ? 'daytype' : 'sensor'
}

/**
 * Schedule data is always refetched, never served from a long-lived cache.
 * Base lists (controls + day types) load once when the Schedules view is
 * activated; every sensor/day change then fetches only its filtered
 * `?control=` / `?daytype=` profile slice. All reloads flush pending edits
 * first.
 */
export function useScheduleReload() {
  async function withReload(work: () => Promise<void>): Promise<void> {
    pendingReloads.value++
    try {
      await work()
    } finally {
      pendingReloads.value--
    }
  }

  /** Resolve the stored sensor against loaded controls, falling back to first. */
  function resolveSensorId(): number | null {
    const { orderedControls } = useControls()
    if (orderedControls.value.length === 0) return null
    const stored = readStoredId(SENSOR_KEY)
    if (stored != null && orderedControls.value.some((control) => control.id === stored)) {
      return stored
    }
    const first = orderedControls.value[0]
    if (!first) return null
    writeStoredString(SENSOR_KEY, String(first.id))
    return first.id
  }

  /** Resolve the stored day type against loaded day types, falling back to first. */
  function resolveDaytypeId(): number | null {
    const { dayTypes } = useDayTypes()
    if (dayTypes.value.length === 0) return null
    const stored = readStoredId(DAYTYPE_KEY)
    if (stored != null && dayTypes.value.some((dayType) => dayType.id === stored)) {
      return stored
    }
    const first = dayTypes.value[0]
    if (!first) return null
    writeStoredString(DAYTYPE_KEY, String(first.id))
    return first.id
  }

  /**
   * Initial load for the Schedules view: base lists in parallel, then the
   * filtered profile slice for the persisted axis/selection.
   */
  async function reloadInitialView(): Promise<void> {
    await withReload(async () => {
      const { flushAll, clearEntries, loadEntriesFor } = useProfiles()
      const { loadControls } = useControls()
      const { loadDayTypes } = useDayTypes()

      // Flush debounced temperature edits first: clearing drops local state,
      // so reloading without flushing would drop unsaved input.
      await flushAll()
      clearEntries()
      await Promise.all([loadDayTypes(), loadControls()])

      if (readStoredAxis() === 'daytype') {
        const id = resolveDaytypeId()
        if (id != null) await loadEntriesFor({ daytype: id })
      } else {
        const id = resolveSensorId()
        if (id != null) await loadEntriesFor({ control: id })
      }
    })
  }

  /** Refetch only the profile slice for one sensor (base lists stay as-is). */
  async function reloadSensorView(controlId: number): Promise<void> {
    await withReload(async () => {
      const { flushAll, clearEntries, loadEntriesFor } = useProfiles()
      await flushAll()
      clearEntries()
      await loadEntriesFor({ control: controlId })
    })
  }

  /** Refetch only the profile slice for one day type (base lists stay as-is). */
  async function reloadDaytypeView(daytypeId: number): Promise<void> {
    await withReload(async () => {
      const { flushAll, clearEntries, loadEntriesFor } = useProfiles()
      await flushAll()
      clearEntries()
      await loadEntriesFor({ daytype: daytypeId })
    })
  }

  return {
    isReloading,
    reloadInitialView,
    reloadSensorView,
    reloadDaytypeView,
    resolveSensorId,
    resolveDaytypeId,
  }
}
