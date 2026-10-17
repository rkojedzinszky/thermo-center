import { ref } from 'vue'
import api from '@/utils/api'
import type { Calendar } from '@/api'
import { toIsoDay } from '@/utils/calendarDays'

export type CalendarSaveStatus = 'saving' | 'error'

// Calendar entries keyed by their ISO day (YYYY-MM-DD).
const entriesMap = ref<Map<string, Calendar>>(new Map())
const statusMap = ref<Map<string, CalendarSaveStatus>>(new Map())
const isLoading = ref(false)
const lastError = ref<string | null>(null)

// Increments on every range fetch; stale responses never overwrite fresher state
// (e.g. rapid month switches where an earlier request finishes last).
let loadGeneration = 0
// Per-day optimistic version so late responses cannot clobber newer edits.
const versions = new Map<string, number>()

function setEntry(day: string, entry: Calendar) {
  entriesMap.value = new Map(entriesMap.value).set(day, entry)
}

function setStatus(day: string, status: CalendarSaveStatus | null) {
  const next = new Map(statusMap.value)
  if (status) {
    next.set(day, status)
  } else {
    next.delete(day)
  }
  statusMap.value = next
}

export function useCalendar() {
  /** Fetch all calendar entries whose day falls into the inclusive range. */
  async function loadRange(from: Date, to: Date): Promise<void> {
    const generation = ++loadGeneration
    isLoading.value = true
    try {
      const result = await api.listCalendar({ dayGte: from, dayLte: to })
      if (generation !== loadGeneration) return
      const items: Calendar[] = result.objects ?? []
      const newMap = new Map<string, Calendar>()
      for (const entry of items) {
        newMap.set(toIsoDay(entry.day), entry)
      }
      entriesMap.value = newMap
      statusMap.value = new Map()
      versions.clear()
      lastError.value = null
    } catch (error) {
      if (generation !== loadGeneration) return
      console.error('Failed to load calendar:', error)
      lastError.value = 'Failed to load the calendar.'
    } finally {
      if (generation === loadGeneration) isLoading.value = false
    }
  }

  /** Day type resource URI for a day, or null when the day has no entry. */
  function dayTypeUri(day: string): string | null {
    return entriesMap.value.get(day)?.daytype ?? null
  }

  function statusFor(day: string): CalendarSaveStatus | null {
    return statusMap.value.get(day) ?? null
  }

  /**
   * Assign a day type to a day's entry. Only entries present in the loaded
   * range can be patched; days without an entry (e.g. beyond the generated
   * calendar horizon) cannot be created from the UI.
   */
  async function setDayType(day: string, daytypeUri: string): Promise<void> {
    const previous = entriesMap.value.get(day) ?? null
    if (!previous) {
      lastError.value = 'This day has no calendar entry to update.'
      setStatus(day, 'error')
      return
    }

    const version = (versions.get(day) ?? 0) + 1
    versions.set(day, version)

    // Apply optimistically, keeping the entry's id.
    setEntry(day, { ...previous, daytype: daytypeUri })
    setStatus(day, 'saving')
    lastError.value = null

    try {
      const updated = await api.patchCalendar({
        id: previous.id,
        patchCalendarRequest: { daytype: daytypeUri },
      })

      // Only apply if no newer edit landed while the request was in flight.
      if ((versions.get(day) ?? 0) !== version) return
      setEntry(day, updated)
      setStatus(day, null)
    } catch (error) {
      if ((versions.get(day) ?? 0) !== version) return
      console.error('Failed to save day type:', error)
      lastError.value = 'Failed to save the day type.'
      setEntry(day, previous)
      setStatus(day, 'error')
    }
  }

  function clearError(): void {
    lastError.value = null
  }

  return {
    isLoading,
    lastError,
    dayTypeUri,
    statusFor,
    loadRange,
    setDayType,
    clearError,
  }
}
