import type { Profile } from '@/api'

/** Number of seconds in a full day. */
export const SECONDS_PER_DAY = 24 * 60 * 60

/** Temperature bounds and step used by the editor. */
export const MIN_TEMP = 5
export const MAX_TEMP = 30
export const TEMP_STEP = 0.5
/** Display range of the temperature chart (fixed, independent of edit bounds). */
export const CHART_MIN_TEMP = 10
export const CHART_MAX_TEMP = 30
/** Default temperature applied to freshly created time points. */
export const DEFAULT_NEW_TEMP = 20

/** Scheduling grid: zones snap to 15-minute boundaries (900s). */
export const SLOT_STEP_SECONDS = 15 * 60
/** Number of distinct slots per day at the default grid resolution. */
export const SLOTS_PER_DAY = SECONDS_PER_DAY / SLOT_STEP_SECONDS

/** Parse a profile start string ("HH:MM" or "HH:MM:SS") into seconds since midnight. */
export function parseStart(value: string): number {
  const trimmed = value.trim()
  if (!trimmed) return 0
  const [hoursRaw = '0', minutesRaw = '0', secondsRaw = '0'] = trimmed.split(':')
  // Reject empty parts ("::", "12:") and non-numeric input.
  if ([hoursRaw, minutesRaw, secondsRaw].some((part) => part.trim() === '')) return 0
  const hours = Number(hoursRaw)
  const minutes = Number(minutesRaw)
  const seconds = Number(secondsRaw)
  if (![hours, minutes, seconds].every((part) => Number.isFinite(part))) {
    return 0
  }
  if (hours < 0 || hours > 23 || minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59) {
    return 0
  }
  return (hours * 60 + minutes) * 60 + seconds
}

/** Format seconds since midnight as "HH:MM", wrapping into a single day. */
export function formatClock(seconds: number): string {
  const wrapped = ((Math.floor(seconds) % SECONDS_PER_DAY) + SECONDS_PER_DAY) % SECONDS_PER_DAY
  const hours = Math.floor(wrapped / 3600)
  const minutes = Math.floor((wrapped % 3600) / 60)
  return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}`
}

/** Format seconds since midnight as a value accepted by the API ("HH:MM:00"). */
export function formatStartValue(seconds: number): string {
  return `${formatClock(seconds)}:00`
}

/** Snap seconds to the nearest `stepMinutes` boundary, clamped to the day. */
export function snapToStep(seconds: number, stepMinutes = 15): number {
  const step = stepMinutes * 60
  const snapped = Math.round(seconds / step) * step
  return Math.min(SECONDS_PER_DAY, Math.max(0, snapped))
}

/** Return a copy of the entries sorted by start time. */
export function sortEntries(entries: Profile[]): Profile[] {
  return [...entries].sort((a, b) => parseStart(a.start) - parseStart(b.start))
}

/** True when another entry (optionally ignoring `ignoreId`) already starts at `seconds`. */
export function hasStartConflict(entries: Profile[], seconds: number, ignoreId?: number): boolean {
  return entries.some((entry) => entry.id !== ignoreId && parseStart(entry.start) === seconds)
}

/**
 * Validate a candidate zone start. Returns an error message when the slot is
 * taken, otherwise null. Single place for the "already starts at" message.
 */
export function validateStart(
  entries: Profile[],
  seconds: number,
  ignoreId?: number,
): string | null {
  if (hasStartConflict(entries, seconds, ignoreId)) {
    return `A zone already starts at ${formatClock(seconds)}.`
  }
  return null
}

export interface ScheduleSegment {
  /** Segment start in seconds since midnight. */
  start: number
  /** Segment end in seconds since midnight (exclusive). */
  end: number
  /** Target temperature, or null for the system default. */
  temp: number | null
  /** Whether this segment represents the system default (no explicit target). */
  isDefault: boolean
  /** Source entry, or null for the implicit leading default segment. */
  entry: Profile | null
}

/**
 * Turn profile entries into contiguous 24h segments.
 *
 * Gaps before the first entry are filled with an implicit default segment. Each
 * entry extends until the next entry's start, and the final entry extends to the
 * end of the day.
 */
export function buildSegments(entries: Profile[]): ScheduleSegment[] {
  const sorted = sortEntries(entries)

  if (sorted.length === 0) {
    return [{ start: 0, end: SECONDS_PER_DAY, temp: null, isDefault: true, entry: null }]
  }

  const segments: ScheduleSegment[] = []
  const first = sorted[0]!
  const firstStart = parseStart(first.start)
  if (firstStart > 0) {
    segments.push({ start: 0, end: firstStart, temp: null, isDefault: true, entry: null })
  }

  sorted.forEach((entry, index) => {
    const start = parseStart(entry.start)
    const nextEntry = sorted[index + 1]
    const end = nextEntry ? parseStart(nextEntry.start) : SECONDS_PER_DAY
    // Skip zero-width segments caused by duplicate starts; the later entry wins.
    if (end <= start) return
    const temp = entry.targetTemp ?? null
    segments.push({
      start,
      end: Math.max(start, end),
      temp,
      isDefault: temp == null,
      entry,
    })
  })

  return segments
}

/** Map a target temperature to a cool→warm color. Returns null for the default case. */
function tempHue(temp: number): number {
  const ratio = Math.min(1, Math.max(0, (temp - MIN_TEMP) / (MAX_TEMP - MIN_TEMP)))
  return Math.round(220 - ratio * 220)
}

/** Map a target temperature to a cool→warm color. Returns null for the default case. */
export function tempColor(temp: number | null): string | null {
  if (temp == null) return null
  return `hsl(${tempHue(temp)}, 65%, 45%)`
}

/** Translucent fill for the zone bar, matching `tempColor`'s hue. */
export function tempFill(temp: number | null, alpha = 0.22): string | null {
  if (temp == null) return null
  return `hsl(${tempHue(temp)}, 65%, 45%, ${alpha})`
}

/** Clamp a temperature to the editor bounds and round to the configured step. */
export function clampTemp(temp: number): number {
  const clamped = Math.min(MAX_TEMP, Math.max(MIN_TEMP, temp))
  return Math.round(clamped / TEMP_STEP) * TEMP_STEP
}

/**
 * Map a temperature to a 0..1 ratio inside the chart display range.
 * Values outside the range clamp to the plot edges.
 */
export function tempToRatio(temp: number): number {
  const ratio = (temp - CHART_MIN_TEMP) / (CHART_MAX_TEMP - CHART_MIN_TEMP)
  return Math.min(1, Math.max(0, ratio))
}

/** Snap a temperature to the editor step and clamp it to the chart range. */
export function snapChartTemp(temp: number): number {
  const snapped = Math.round(temp / TEMP_STEP) * TEMP_STEP
  return Math.min(CHART_MAX_TEMP, Math.max(CHART_MIN_TEMP, snapped))
}
