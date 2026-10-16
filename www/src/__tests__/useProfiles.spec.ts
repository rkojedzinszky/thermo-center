import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import api from '../utils/api'
import { ResponseError, type Profile } from '../api'
import { useProfiles } from '../composables/useProfiles'
import { makeProfile as makeBaseProfile } from './factories'

vi.mock('../utils/api', () => ({
  default: {
    listProfile: vi.fn(),
    createProfile: vi.fn(),
    patchProfile: vi.fn(),
    deleteProfile: vi.fn(),
    getProfile: vi.fn(),
  },
}))

const mockedApi = api as unknown as {
  listProfile: ReturnType<typeof vi.fn>
  createProfile: ReturnType<typeof vi.fn>
  patchProfile: ReturnType<typeof vi.fn>
  deleteProfile: ReturnType<typeof vi.fn>
  getProfile: ReturnType<typeof vi.fn>
}

const CONTROL = '/api/v1/control/1/'
const DAYTYPE = '/api/v1/daytype/1/'
// Filters take plain primary keys, not resource URIs.
const CONTROL_ID = 1
const DAYTYPE_ID = 1

function makeProfile(overrides: Partial<Profile> = {}): Profile {
  return makeBaseProfile(
    overrides.id ?? 1,
    overrides.start ?? '06:00:00',
    overrides.targetTemp ?? 20,
    {
      control: CONTROL,
      daytype: DAYTYPE,
      ...overrides,
    },
  )
}

describe('useProfiles', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    mockedApi.listProfile.mockResolvedValue({ objects: [], meta: {} })
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('loads and groups profiles by control and daytype', async () => {
    mockedApi.listProfile.mockResolvedValue({
      objects: [
        makeProfile({ id: 2, start: '22:00:00' }),
        makeProfile({ id: 1, start: '06:00:00' }),
        makeProfile({ id: 3, control: '/api/v1/control/9/', start: '07:00:00' }),
      ],
      meta: {},
    })

    const { loadProfiles, entriesFor } = useProfiles()
    await loadProfiles()

    const entries = entriesFor(CONTROL, DAYTYPE)
    expect(entries.map((entry) => entry.id)).toEqual([1, 2])
    expect(entriesFor('/api/v1/control/9/', DAYTYPE)).toHaveLength(1)
  })

  it('optimistically adds an entry and replaces it with the created one', async () => {
    const { loadProfiles, entriesFor, addEntry } = useProfiles()
    await loadProfiles()

    let resolveCreate: (value: Profile) => void = () => {}
    mockedApi.createProfile.mockReturnValue(
      new Promise<Profile>((resolve) => {
        resolveCreate = resolve
      }),
    )

    const promise = addEntry(CONTROL, DAYTYPE, 6 * 3600, 20)

    // Optimistic entry is visible before the request resolves.
    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(1)
    expect(mockedApi.createProfile).toHaveBeenCalledWith({
      profileW: { control: CONTROL, daytype: DAYTYPE, start: '06:00:00', targetTemp: 20 },
    })

    resolveCreate(makeProfile({ id: 7 }))
    await promise

    const entries = entriesFor(CONTROL, DAYTYPE)
    expect(entries).toHaveLength(1)
    expect(entries[0]?.id).toBe(7)
  })

  it('rolls back a failed add', async () => {
    mockedApi.createProfile.mockRejectedValue(new Error('boom'))

    const { loadProfiles, entriesFor, addEntry } = useProfiles()
    await loadProfiles()
    await addEntry(CONTROL, DAYTYPE, 6 * 3600, 20)

    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(0)
  })

  it('applies a temperature change optimistically and persists it after debounce', async () => {
    vi.useFakeTimers()
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.patchProfile.mockResolvedValue(makeProfile({ id: 1, targetTemp: 22 }))

    const { loadProfiles, entriesFor, setEntryTemp } = useProfiles()
    await loadProfiles()

    setEntryTemp(1, 22)
    expect(entriesFor(CONTROL, DAYTYPE)[0]?.targetTemp).toBe(22)
    expect(mockedApi.patchProfile).not.toHaveBeenCalled()

    await vi.advanceTimersByTimeAsync(500)
    expect(mockedApi.patchProfile).toHaveBeenCalledWith({
      id: 1,
      patchProfileRequest: { targetTemp: 22 },
    })
  })

  it('resyncs and flags an error when a change fails', async () => {
    vi.useFakeTimers()
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.patchProfile.mockRejectedValue(new Error('boom'))
    mockedApi.getProfile.mockResolvedValue(makeProfile({ id: 1, targetTemp: 20 }))

    const { loadProfiles, entriesFor, statusFor, setEntryTemp } = useProfiles()
    await loadProfiles()

    setEntryTemp(1, 25)
    await vi.advanceTimersByTimeAsync(500)
    await vi.runOnlyPendingTimersAsync()

    expect(mockedApi.getProfile).toHaveBeenCalledWith({ id: 1 })
    expect(entriesFor(CONTROL, DAYTYPE)[0]?.targetTemp).toBe(20)
    expect(statusFor(1)).toBe('error')
  })

  it('removes an entry optimistically and calls the API', async () => {
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.deleteProfile.mockResolvedValue(undefined)

    const { loadProfiles, entriesFor, removeEntry } = useProfiles()
    await loadProfiles()
    await removeEntry(1)

    expect(mockedApi.deleteProfile).toHaveBeenCalledWith({ id: 1 })
    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(0)
  })

  it('restores an entry when delete fails', async () => {
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.deleteProfile.mockRejectedValue(new Error('boom'))

    const { loadProfiles, entriesFor, removeEntry } = useProfiles()
    await loadProfiles()
    await removeEntry(1)

    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(1)
  })

  it('cleans up a created entry that was removed while the create was in flight', async () => {
    mockedApi.deleteProfile.mockResolvedValue(undefined)

    const { loadProfiles, entriesFor, addEntry, removeEntry } = useProfiles()
    await loadProfiles()

    let resolveCreate: (value: Profile) => void = () => {}
    mockedApi.createProfile.mockReturnValue(
      new Promise<Profile>((resolve) => {
        resolveCreate = resolve
      }),
    )

    const promise = addEntry(CONTROL, DAYTYPE, 6 * 3600, 20)
    const optimisticId = entriesFor(CONTROL, DAYTYPE)[0]!.id
    await removeEntry(optimisticId)

    resolveCreate(makeProfile({ id: 42 }))
    await promise

    expect(mockedApi.deleteProfile).toHaveBeenCalledWith({ id: 42 })
    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(0)
  })

  it('keeps the optimistic entry when resync fails with a network error', async () => {
    vi.useFakeTimers()
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.patchProfile.mockRejectedValue(new Error('boom'))
    mockedApi.getProfile.mockRejectedValue(new Error('network'))

    const { loadProfiles, entriesFor, statusFor, setEntryTemp } = useProfiles()
    await loadProfiles()

    setEntryTemp(1, 25)
    await vi.advanceTimersByTimeAsync(500)
    await vi.runOnlyPendingTimersAsync()

    // Network blip must not look like a delete.
    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(1)
    expect(entriesFor(CONTROL, DAYTYPE)[0]?.targetTemp).toBe(25)
    expect(statusFor(1)).toBe('error')
  })

  it('drops the entry when resync confirms a 404', async () => {
    vi.useFakeTimers()
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.patchProfile.mockRejectedValue(new Error('boom'))
    mockedApi.getProfile.mockRejectedValue(
      new ResponseError(new Response(null, { status: 404 }), 'Not found'),
    )

    const { loadProfiles, entriesFor, setEntryTemp } = useProfiles()
    await loadProfiles()

    setEntryTemp(1, 25)
    await vi.advanceTimersByTimeAsync(500)
    await vi.runOnlyPendingTimersAsync()

    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(0)
  })

  it('flushes pending patches on demand', async () => {
    vi.useFakeTimers()
    mockedApi.listProfile.mockResolvedValue({ objects: [makeProfile({ id: 1 })], meta: {} })
    mockedApi.patchProfile.mockResolvedValue(makeProfile({ id: 1, targetTemp: 23 }))

    const { loadProfiles, setEntryTemp, flushAll } = useProfiles()
    await loadProfiles()

    setEntryTemp(1, 23)
    expect(mockedApi.patchProfile).not.toHaveBeenCalled()
    await flushAll()
    expect(mockedApi.patchProfile).toHaveBeenCalledWith({
      id: 1,
      patchProfileRequest: { targetTemp: 23 },
    })
  })

  it('loads a filtered slice with the control filter', async () => {
    mockedApi.listProfile.mockResolvedValue({
      objects: [makeProfile({ id: 5, start: '08:00:00' })],
      meta: {},
    })

    const { loadEntriesFor, entriesFor } = useProfiles()
    await loadEntriesFor({ control: CONTROL_ID })

    expect(mockedApi.listProfile).toHaveBeenCalledWith({ control: CONTROL_ID })
    expect(entriesFor(CONTROL, DAYTYPE)).toHaveLength(1)
    expect(entriesFor(CONTROL, DAYTYPE)[0]?.id).toBe(5)
  })

  it('loads a filtered slice with the daytype filter', async () => {
    const { loadEntriesFor } = useProfiles()
    await loadEntriesFor({ daytype: DAYTYPE_ID })

    expect(mockedApi.listProfile).toHaveBeenCalledWith({ daytype: DAYTYPE_ID })
  })

  it('ignores a stale slice response when a newer load wins', async () => {
    let resolveFirst!: (value: { objects: ReturnType<typeof makeProfile>[]; meta: object }) => void
    mockedApi.listProfile
      .mockReturnValueOnce(
        new Promise((resolve) => {
          resolveFirst = resolve
        }),
      )
      .mockResolvedValue({
        objects: [makeProfile({ id: 9, control: '/api/v1/control/2/', start: '09:00:00' })],
        meta: {},
      })

    const { loadEntriesFor, entriesFor } = useProfiles()
    const first = loadEntriesFor({ control: CONTROL_ID })
    await loadEntriesFor({ control: 2 })
    resolveFirst({ objects: [makeProfile({ id: 1, start: '06:00:00' })], meta: {} })
    await first

    // The superseded first response must not clobber the newer slice.
    expect(entriesFor('/api/v1/control/2/', DAYTYPE)[0]?.id).toBe(9)
  })
})
