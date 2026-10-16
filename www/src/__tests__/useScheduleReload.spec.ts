import { describe, it, expect, vi, beforeEach } from 'vitest'
import api from '../utils/api'
import { useScheduleReload } from '../composables/useScheduleReload'
import { useProfiles } from '../composables/useProfiles'
import { CONTROL_URI, DAYTYPE_URI, makeControl, makeDayType, makeProfile } from './factories'

vi.mock('../utils/api', () => ({
  default: {
    listProfile: vi.fn(),
    createProfile: vi.fn(),
    patchProfile: vi.fn(),
    deleteProfile: vi.fn(),
    getProfile: vi.fn(),
    listControl: vi.fn(),
    listDayType: vi.fn(),
  },
}))

const mockedApi = api as unknown as {
  listProfile: ReturnType<typeof vi.fn>
  patchProfile: ReturnType<typeof vi.fn>
  listControl: ReturnType<typeof vi.fn>
  listDayType: ReturnType<typeof vi.fn>
}

describe('useScheduleReload', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    localStorage.clear()
    mockedApi.listProfile.mockResolvedValue({ objects: [], meta: {} })
    mockedApi.listControl.mockResolvedValue({ objects: [], meta: {} })
    mockedApi.listDayType.mockResolvedValue({ objects: [], meta: {} })
    mockedApi.patchProfile.mockResolvedValue(makeProfile(1, '06:00:00', 23))
  })

  it('loads base lists then the filtered slice for the persisted sensor', async () => {
    mockedApi.listControl.mockResolvedValue({
      objects: [makeControl(1, 'Living Room'), makeControl(2, 'Bedroom')],
      meta: {},
    })
    mockedApi.listDayType.mockResolvedValue({
      objects: [makeDayType(1, 'workday')],
      meta: {},
    })
    localStorage.setItem('profile_sensor', '2')

    const { reloadInitialView } = useScheduleReload()
    await reloadInitialView()

    expect(mockedApi.listControl).toHaveBeenCalledTimes(1)
    expect(mockedApi.listDayType).toHaveBeenCalledTimes(1)
    // No unfiltered profile fetch: only the selected sensor's slice.
    expect(mockedApi.listProfile).toHaveBeenCalledTimes(1)
    expect(mockedApi.listProfile).toHaveBeenCalledWith({ control: 2 })
  })

  it('loads the daytype slice when the persisted axis is daytype', async () => {
    mockedApi.listControl.mockResolvedValue({
      objects: [makeControl(1, 'Living Room')],
      meta: {},
    })
    mockedApi.listDayType.mockResolvedValue({
      objects: [makeDayType(1, 'workday'), makeDayType(2, 'Sunday')],
      meta: {},
    })
    localStorage.setItem('profile_group_axis', 'daytype')
    localStorage.setItem('profile_daytype', '2')

    const { reloadInitialView } = useScheduleReload()
    await reloadInitialView()

    expect(mockedApi.listProfile).toHaveBeenCalledTimes(1)
    expect(mockedApi.listProfile).toHaveBeenCalledWith({ daytype: 2 })
  })

  it('falls back to the first sensor when the stored selection is stale', async () => {
    mockedApi.listControl.mockResolvedValue({
      objects: [makeControl(1, 'Living Room')],
      meta: {},
    })
    mockedApi.listDayType.mockResolvedValue({
      objects: [makeDayType(1, 'workday')],
      meta: {},
    })
    localStorage.setItem('profile_sensor', '99')

    const { reloadInitialView } = useScheduleReload()
    await reloadInitialView()

    expect(mockedApi.listProfile).toHaveBeenCalledWith({ control: 1 })
    expect(localStorage.getItem('profile_sensor')).toBe('1')
  })

  it('skips the profile fetch when there is nothing to select', async () => {
    const { reloadInitialView } = useScheduleReload()
    await reloadInitialView()

    expect(mockedApi.listControl).toHaveBeenCalledTimes(1)
    expect(mockedApi.listDayType).toHaveBeenCalledTimes(1)
    expect(mockedApi.listProfile).not.toHaveBeenCalled()
  })

  it('flushes pending edits before reloading a sensor slice', async () => {
    mockedApi.listProfile.mockResolvedValue({
      objects: [makeProfile(1, '06:00:00', 20)],
      meta: {},
    })
    const { loadProfiles, entriesFor, setEntryTemp } = useProfiles()
    await loadProfiles()
    expect(mockedApi.listProfile).toHaveBeenCalledTimes(1)

    setEntryTemp(1, 23)
    expect(entriesFor('/api/v1/control/1/', '/api/v1/daytype/1/')[0]?.targetTemp).toBe(23)

    const { reloadSensorView } = useScheduleReload()
    await reloadSensorView(1)

    // Pending temperature edit is persisted, not wiped by the reload.
    expect(mockedApi.patchProfile).toHaveBeenCalledWith({
      id: 1,
      patchProfileRequest: { targetTemp: 23 },
    })
    expect(mockedApi.listProfile).toHaveBeenLastCalledWith({ control: 1 })
  })

  it('reports reloading state while in flight', async () => {
    let resolveList!: (value: { objects: never[]; meta: object }) => void
    mockedApi.listProfile.mockReturnValue(
      new Promise((resolve) => {
        resolveList = resolve
      }),
    )

    const { reloadSensorView, isReloading } = useScheduleReload()
    expect(isReloading.value).toBe(false)

    const promise = reloadSensorView(1)
    expect(isReloading.value).toBe(true)

    resolveList({ objects: [], meta: {} })
    await promise
    expect(isReloading.value).toBe(false)
  })

  it('fetches only the filtered slice when reloading a sensor view', async () => {
    const { reloadSensorView } = useScheduleReload()
    await reloadSensorView(2)

    expect(mockedApi.listProfile).toHaveBeenCalledWith({ control: 2 })
    expect(mockedApi.listControl).not.toHaveBeenCalled()
    expect(mockedApi.listDayType).not.toHaveBeenCalled()
  })

  it('fetches only the filtered slice when reloading a day type view', async () => {
    const { reloadDaytypeView } = useScheduleReload()
    await reloadDaytypeView(2)

    expect(mockedApi.listProfile).toHaveBeenCalledWith({ daytype: 2 })
    expect(mockedApi.listControl).not.toHaveBeenCalled()
    expect(mockedApi.listDayType).not.toHaveBeenCalled()
  })

  it('clears stale entries while the fresh slice is in flight', async () => {
    mockedApi.listProfile.mockResolvedValueOnce({
      objects: [makeProfile(1, '06:00:00', 20)],
      meta: {},
    })
    const { loadProfiles, entriesFor } = useProfiles()
    await loadProfiles()
    expect(entriesFor(CONTROL_URI, DAYTYPE_URI)).toHaveLength(1)

    let resolveSlice!: (value: { objects: ReturnType<typeof makeProfile>[]; meta: object }) => void
    mockedApi.listProfile.mockReturnValueOnce(
      new Promise((resolve) => {
        resolveSlice = resolve
      }),
    )

    const { reloadSensorView } = useScheduleReload()
    const pending = reloadSensorView(2)

    // No stale rows during the flight: the old slice is dropped immediately.
    await vi.waitFor(() => expect(entriesFor(CONTROL_URI, DAYTYPE_URI)).toHaveLength(0))

    resolveSlice({ objects: [makeProfile(9, '09:00:00', 21)], meta: {} })
    await pending
    expect(entriesFor(CONTROL_URI, DAYTYPE_URI)).toHaveLength(1)
    expect(entriesFor(CONTROL_URI, DAYTYPE_URI)[0]?.id).toBe(9)
  })

  it('clears a previous visit’s entries on initial load', async () => {
    mockedApi.listProfile.mockResolvedValueOnce({
      objects: [makeProfile(1, '06:00:00', 20)],
      meta: {},
    })
    const { loadProfiles, entriesFor } = useProfiles()
    await loadProfiles()
    expect(entriesFor(CONTROL_URI, DAYTYPE_URI)).toHaveLength(1)

    mockedApi.listControl.mockResolvedValue({
      objects: [makeControl(1, 'Living Room')],
      meta: {},
    })
    mockedApi.listDayType.mockResolvedValue({
      objects: [makeDayType(1, 'workday')],
      meta: {},
    })
    let resolveSlice!: (value: { objects: never[]; meta: object }) => void
    mockedApi.listProfile.mockReturnValueOnce(
      new Promise((resolve) => {
        resolveSlice = resolve
      }),
    )

    const { reloadInitialView } = useScheduleReload()
    const pending = reloadInitialView()

    await vi.waitFor(() => expect(entriesFor(CONTROL_URI, DAYTYPE_URI)).toHaveLength(0))

    resolveSlice({ objects: [], meta: {} })
    await pending
    expect(mockedApi.listProfile).toHaveBeenLastCalledWith({ control: 1 })
  })
})
