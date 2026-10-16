import { describe, it, expect } from 'vitest'
import type { Profile } from '../api'
import {
  SECONDS_PER_DAY,
  buildSegments,
  clampTemp,
  formatClock,
  formatStartValue,
  hasStartConflict,
  parseStart,
  snapChartTemp,
  snapToStep,
  tempColor,
  tempFill,
  tempToRatio,
  validateStart,
} from '../utils/profileSchedule'
import { makeProfile } from './factories'

function profile(overrides: Partial<Profile> & { id: number; start: string }): Profile {
  return makeProfile(overrides.id, overrides.start, overrides.targetTemp ?? 20, overrides)
}

describe('profileSchedule time helpers', () => {
  it('parses "HH:MM" and "HH:MM:SS" to seconds', () => {
    expect(parseStart('06:30')).toBe(6 * 3600 + 30 * 60)
    expect(parseStart('06:30:15')).toBe(6 * 3600 + 30 * 60 + 15)
    expect(parseStart('00:00:00')).toBe(0)
  })

  it('formats seconds as HH:MM and as an API start value', () => {
    expect(formatClock(6 * 3600 + 30 * 60)).toBe('06:30')
    expect(formatStartValue(6 * 3600 + 30 * 60)).toBe('06:30:00')
  })

  it('snaps seconds to the nearest step and clamps to the day', () => {
    expect(snapToStep(800, 15)).toBe(900)
    expect(snapToStep(100, 15)).toBe(0)
    expect(snapToStep(-100, 15)).toBe(0)
    expect(snapToStep(SECONDS_PER_DAY + 100, 15)).toBe(SECONDS_PER_DAY)
  })

  it('detects start conflicts while ignoring the edited entry', () => {
    const entries = [profile({ id: 1, start: '06:00:00' }), profile({ id: 2, start: '08:00:00' })]
    expect(hasStartConflict(entries, 8 * 3600)).toBe(true)
    expect(hasStartConflict(entries, 8 * 3600, 2)).toBe(false)
    expect(hasStartConflict(entries, 9 * 3600)).toBe(false)
  })

  it('clamps and steps temperatures', () => {
    expect(clampTemp(35)).toBe(30)
    expect(clampTemp(3)).toBe(5)
    expect(clampTemp(20.3)).toBe(20.5)
  })

  it('returns null color for the default case and a hue for temperatures', () => {
    expect(tempColor(null)).toBeNull()
    expect(tempColor(5)).toContain('hsl(220')
    expect(tempColor(30)).toContain('hsl(0')
  })

  it('returns a translucent fill matching the temperature hue', () => {
    expect(tempFill(null)).toBeNull()
    expect(tempFill(20)).toContain('hsl(')
    expect(tempFill(20)).toContain('0.22')
    expect(tempFill(20, 0.4)).toContain('0.4')
  })

  it('maps temperatures onto the 10–30 °C chart range with clamping', () => {
    expect(tempToRatio(10)).toBe(0)
    expect(tempToRatio(30)).toBe(1)
    expect(tempToRatio(20)).toBe(0.5)
    expect(tempToRatio(5)).toBe(0)
    expect(tempToRatio(35)).toBe(1)
  })

  it('snaps chart temperatures to the step within the display range', () => {
    expect(snapChartTemp(22.3)).toBe(22.5)
    expect(snapChartTemp(5)).toBe(10)
    expect(snapChartTemp(35)).toBe(30)
  })
})

describe('buildSegments', () => {
  it('returns a single default segment when there are no entries', () => {
    const segments = buildSegments([])
    expect(segments).toHaveLength(1)
    expect(segments[0]).toMatchObject({
      start: 0,
      end: SECONDS_PER_DAY,
      isDefault: true,
      entry: null,
    })
  })

  it('adds an implicit default head before the first entry', () => {
    const segments = buildSegments([profile({ id: 1, start: '06:00:00', targetTemp: 21 })])
    expect(segments).toHaveLength(2)
    expect(segments[0]).toMatchObject({ start: 0, end: 6 * 3600, isDefault: true })
    expect(segments[1]).toMatchObject({
      start: 6 * 3600,
      end: SECONDS_PER_DAY,
      temp: 21,
      isDefault: false,
    })
  })

  it('extends each entry until the next and sorts unordered input', () => {
    const segments = buildSegments([
      profile({ id: 2, start: '22:00:00', targetTemp: 18 }),
      profile({ id: 1, start: '06:00:00', targetTemp: 21 }),
    ])
    expect(segments).toHaveLength(3)
    expect(segments[1]).toMatchObject({ start: 6 * 3600, end: 22 * 3600, temp: 21 })
    expect(segments[2]).toMatchObject({ start: 22 * 3600, end: SECONDS_PER_DAY, temp: 18 })
  })

  it('does not add a head when the first entry starts at midnight', () => {
    const segments = buildSegments([profile({ id: 1, start: '00:00:00', targetTemp: 19 })])
    expect(segments).toHaveLength(1)
    expect(segments[0]).toMatchObject({ start: 0, end: SECONDS_PER_DAY, temp: 19 })
  })

  it('treats a null target as a default segment', () => {
    const segments = buildSegments([profile({ id: 1, start: '00:00:00', targetTemp: null })])
    expect(segments[0]?.isDefault).toBe(true)
  })

  it('skips zero-width segments from duplicate starts', () => {
    const segments = buildSegments([
      profile({ id: 1, start: '06:00:00', targetTemp: 21 }),
      profile({ id: 2, start: '06:00:00', targetTemp: 18 }),
    ])
    // Later entry wins; no zero-width segment is rendered.
    expect(segments.every((segment) => segment.end > segment.start)).toBe(true)
    expect(segments[segments.length - 1]).toMatchObject({ temp: 18 })
  })

  it('rejects out-of-range and malformed start strings', () => {
    expect(parseStart('')).toBe(0)
    expect(parseStart('99:99')).toBe(0)
    expect(parseStart('24:00')).toBe(0)
    expect(parseStart('12:60')).toBe(0)
    expect(parseStart('::')).toBe(0)
  })

  it('validates start conflicts with a shared message', () => {
    const entries = [profile({ id: 1, start: '06:00:00' })]
    expect(validateStart(entries, 6 * 3600)).toContain('06:00')
    expect(validateStart(entries, 6 * 3600, 1)).toBeNull()
    expect(validateStart(entries, 7 * 3600)).toBeNull()
  })
})
