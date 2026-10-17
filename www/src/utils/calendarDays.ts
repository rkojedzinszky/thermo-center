/**
 * Date helpers for the day-type calendar.
 *
 * Calendar days are plain dates (no time zone on the server), so all math is
 * done on UTC-midnight dates and formatted as `YYYY-MM-DD`.
 */

const DAY_MS = 24 * 60 * 60 * 1000

/** Number of weeks rendered per month view. */
export const GRID_WEEKS = 6

/** Format a date as `YYYY-MM-DD` using its UTC parts. */
export function toIsoDay(date: Date): string {
  return date.toISOString().slice(0, 10)
}

/** Today as `YYYY-MM-DD` in the browser's local time zone. */
export function todayIsoDay(now: Date = new Date()): string {
  const month = String(now.getMonth() + 1).padStart(2, '0')
  const day = String(now.getDate()).padStart(2, '0')
  return `${now.getFullYear()}-${month}-${day}`
}

export interface MonthGrid {
  /** First cell of the grid (Monday of the week containing the 1st). */
  start: Date
  /** Last cell of the grid. */
  end: Date
  days: Date[]
}

/** Six-week grid (Monday first) covering the given 0-based month. */
export function monthGrid(year: number, month: number): MonthGrid {
  const first = new Date(Date.UTC(year, month, 1))
  const mondayOffset = (first.getUTCDay() + 6) % 7
  const start = new Date(Date.UTC(year, month, 1 - mondayOffset))
  const days = Array.from({ length: GRID_WEEKS * 7 }, (_, index) => {
    return new Date(start.getTime() + index * DAY_MS)
  })
  return { start, end: days[days.length - 1]!, days }
}

const MONTH_FORMAT = new Intl.DateTimeFormat('en-US', { month: 'long', year: 'numeric' })

/** e.g. "October 2026" for a 0-based month. */
export function monthLabel(year: number, month: number): string {
  return MONTH_FORMAT.format(new Date(Date.UTC(year, month, 1)))
}

/** Shift a year/month pair by whole months. */
export function shiftMonth(
  year: number,
  month: number,
  delta: number,
): { year: number; month: number } {
  const shifted = new Date(Date.UTC(year, month + delta, 1))
  return { year: shifted.getUTCFullYear(), month: shifted.getUTCMonth() }
}
