import { describe, it, expect, vi, beforeEach } from 'vitest'
import api from '../utils/api'
import { ResponseError, type Calendar } from '../api'
import { useCalendar } from '../composables/useCalendar'
import { makeCalendar } from './factories'

vi.mock('../utils/api', () => ({
  default: {
    listCalendar: vi.fn(),
    patchCalendar: vi.fn(),
  },
}))

const mockedApi = api as unknown as {
  listCalendar: ReturnType<typeof vi.fn>
  patchCalendar: ReturnType<typeof vi.fn>
}

const DAY = '2026-10-26'
const ENTRY_ID = 42
const RANGE_START = new Date(Date.UTC(2026, 9, 26))
const RANGE_END = new Date(Date.UTC(2026, 10, 1))
const OLD_DAYTYPE = '/api/v1/daytype/1/'
const NEW_DAYTYPE = '/api/v1/daytype/2/'

describe('useCalendar', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    mockedApi.listCalendar.mockResolvedValue({ objects: [], meta: {} })
  })

  it('loads a range and maps entries by their ISO day', async () => {
    mockedApi.listCalendar.mockResolvedValue({
      objects: [makeCalendar(ENTRY_ID, DAY, NEW_DAYTYPE)],
      meta: {},
    })

    const { loadRange, dayTypeUri } = useCalendar()
    await loadRange(RANGE_START, RANGE_END)

    expect(mockedApi.listCalendar).toHaveBeenCalledWith({
      dayGte: RANGE_START,
      dayLte: RANGE_END,
    })
    expect(dayTypeUri(DAY)).toBe(NEW_DAYTYPE)
    expect(dayTypeUri('2026-10-27')).toBeNull()
  })

  it('applies a day type optimistically and patches the entry by id', async () => {
    mockedApi.listCalendar.mockResolvedValue({
      objects: [makeCalendar(ENTRY_ID, DAY, OLD_DAYTYPE)],
      meta: {},
    })
    let resolvePatch: (value: Calendar) => void = () => {}
    mockedApi.patchCalendar.mockReturnValue(
      new Promise<Calendar>((resolve) => {
        resolvePatch = resolve
      }),
    )

    const { loadRange, dayTypeUri, setDayType, statusFor } = useCalendar()
    await loadRange(RANGE_START, RANGE_END)

    const promise = setDayType(DAY, NEW_DAYTYPE)

    expect(dayTypeUri(DAY)).toBe(NEW_DAYTYPE)
    expect(statusFor(DAY)).toBe('saving')
    expect(mockedApi.patchCalendar).toHaveBeenCalledWith({
      id: ENTRY_ID,
      patchCalendarRequest: { daytype: NEW_DAYTYPE },
    })

    resolvePatch(makeCalendar(ENTRY_ID, DAY, NEW_DAYTYPE))
    await promise

    expect(dayTypeUri(DAY)).toBe(NEW_DAYTYPE)
    expect(statusFor(DAY)).toBeNull()
  })

  it('rolls back and flags an error when the patch fails', async () => {
    mockedApi.listCalendar.mockResolvedValue({
      objects: [makeCalendar(ENTRY_ID, DAY, OLD_DAYTYPE)],
      meta: {},
    })
    mockedApi.patchCalendar.mockRejectedValue(
      new ResponseError(new Response(null, { status: 404 }), 'Not Found'),
    )

    const { loadRange, dayTypeUri, setDayType, statusFor, lastError } = useCalendar()
    await loadRange(RANGE_START, RANGE_END)

    await setDayType(DAY, NEW_DAYTYPE)

    expect(dayTypeUri(DAY)).toBe(OLD_DAYTYPE)
    expect(statusFor(DAY)).toBe('error')
    expect(lastError.value).toBe('Failed to save the day type.')
  })

  it('refuses to patch a day without a loaded entry', async () => {
    const { loadRange, setDayType, statusFor, lastError } = useCalendar()
    await loadRange(RANGE_START, RANGE_END)

    await setDayType(DAY, NEW_DAYTYPE)

    expect(mockedApi.patchCalendar).not.toHaveBeenCalled()
    expect(statusFor(DAY)).toBe('error')
    expect(lastError.value).toBe('This day has no calendar entry to update.')
  })

  it('ignores a stale range response', async () => {
    let resolveFirst: (value: unknown) => void = () => {}
    mockedApi.listCalendar.mockReturnValueOnce(
      new Promise((resolve) => {
        resolveFirst = resolve
      }),
    )
    mockedApi.listCalendar.mockResolvedValueOnce({
      objects: [makeCalendar(ENTRY_ID, DAY, NEW_DAYTYPE)],
      meta: {},
    })

    const { loadRange, dayTypeUri } = useCalendar()
    const first = loadRange(new Date(Date.UTC(2026, 7, 1)), new Date(Date.UTC(2026, 8, 12)))
    const second = loadRange(RANGE_START, RANGE_END)
    await second

    resolveFirst({ objects: [makeCalendar(1, '2026-08-15', OLD_DAYTYPE)], meta: {} })
    await first

    expect(dayTypeUri(DAY)).toBe(NEW_DAYTYPE)
    expect(dayTypeUri('2026-08-15')).toBeNull()
  })

  it('clears the error banner', async () => {
    mockedApi.listCalendar.mockResolvedValue({
      objects: [makeCalendar(ENTRY_ID, DAY, OLD_DAYTYPE)],
      meta: {},
    })
    mockedApi.patchCalendar.mockRejectedValue(new Error('boom'))

    const { loadRange, setDayType, lastError, clearError } = useCalendar()
    await loadRange(RANGE_START, RANGE_END)
    await setDayType(DAY, NEW_DAYTYPE)

    expect(lastError.value).not.toBeNull()
    clearError()
    expect(lastError.value).toBeNull()
  })
})
