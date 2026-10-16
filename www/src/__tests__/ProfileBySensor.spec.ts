import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { ref, nextTick } from 'vue'
import { mount, enableAutoUnmount } from '@vue/test-utils'
import ProfileBySensor from '../components/ProfileBySensor.vue'
import type { Control, DayType, Profile } from '../api'
import { makeControl, makeDayType } from './factories'

enableAutoUnmount(afterEach)

const mockUpdateControl = vi.fn()
const mockReloadSensorView = vi.fn()
const mockSaveTick = ref(0)
const mockControls = ref<Control[]>([])
const mockDayTypes = ref<DayType[]>([])

vi.mock('../composables/useProfiles', () => ({
  useProfiles: () => ({
    entriesFor: () => [] as Profile[],
    statusFor: () => null,
    saveTick: mockSaveTick,
    addEntry: vi.fn(),
    setEntryStart: vi.fn(),
    setEntryTemp: vi.fn(),
    setEntryDefault: vi.fn(),
    removeEntry: vi.fn(),
    lastError: ref(null),
    clearError: vi.fn(),
    loadProfiles: vi.fn(),
    loadEntriesFor: vi.fn(),
  }),
}))

vi.mock('../composables/useControls', () => ({
  useControls: () => ({
    orderedControls: mockControls,
    loadControls: vi.fn(),
    updateControlDirect: vi.fn(),
    updateControl: mockUpdateControl,
  }),
}))

vi.mock('../composables/useDayTypes', () => ({
  useDayTypes: () => ({
    dayTypes: mockDayTypes,
    loadDayTypes: vi.fn(),
  }),
}))

vi.mock('../composables/useScheduleReload', () => ({
  useScheduleReload: () => ({
    reloadInitialView: vi.fn(),
    reloadSensorView: mockReloadSensorView,
    reloadDaytypeView: vi.fn(),
    resolveSensorId: vi.fn(),
    resolveDaytypeId: vi.fn(),
    isReloading: ref(false),
  }),
}))

describe('ProfileBySensor', () => {
  beforeEach(() => {
    localStorage.clear()
    vi.clearAllMocks()
    mockSaveTick.value = 0
    mockControls.value = [makeControl(1, 'Living Room'), makeControl(2, 'Bedroom')]
    mockDayTypes.value = [makeDayType(1, 'workday'), makeDayType(2, 'Sunday')]
  })

  it('renders a tab per sensor, a summary and a row per day type', () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })

    expect(wrapper.findAll('.tab')).toHaveLength(2)
    expect(wrapper.find('.summary-name').text()).toBe('Living Room')
    expect(wrapper.text()).toContain('21.50 °C')

    const rows = wrapper.findAll('.sensor-row')
    expect(rows).toHaveLength(2)
    expect(rows[0]!.find('.sensor-name').text()).toBe('workday')
    expect(rows[1]!.find('.sensor-name').text()).toBe('Sunday')
  })

  it('hides the per-row meta in the by-sensor orientation', () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })
    for (const row of wrapper.findAll('.sensor-row')) {
      expect(row.find('.sensor-meta').exists()).toBe(false)
    }
  })

  it('renders tabs without count badges', () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })

    expect(wrapper.findAll('.tab')).toHaveLength(2)
    expect(wrapper.find('.tab-count').exists()).toBe(false)
  })

  it('switches the sensor via the tab strip', async () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })

    await wrapper.findAll('.tab')[1]!.trigger('click')

    expect(wrapper.find('.summary-name').text()).toBe('Bedroom')
  })

  it('refetches the filtered slice when switching sensors', async () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })

    await wrapper.findAll('.tab')[1]!.trigger('click')

    expect(mockReloadSensorView).toHaveBeenCalledTimes(1)
    expect(mockReloadSensorView).toHaveBeenCalledWith(2)
  })

  it('does not refetch when clicking the active sensor', async () => {
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })
    mockReloadSensorView.mockClear()

    await wrapper.findAll('.tab')[0]!.trigger('click')

    expect(mockReloadSensorView).not.toHaveBeenCalled()
  })

  it('refreshes the selected control once after a save', async () => {
    mount(ProfileBySensor, { props: { mode: 'temperature' } })

    mockSaveTick.value++
    await nextTick()

    expect(mockUpdateControl).toHaveBeenCalledTimes(1)
    expect(mockUpdateControl).toHaveBeenCalledWith(1)
  })

  it('shows an empty state when there are no controls', () => {
    mockControls.value = []
    const wrapper = mount(ProfileBySensor, { props: { mode: 'temperature' } })
    expect(wrapper.find('.empty-state').exists()).toBe(true)
    expect(wrapper.find('.sensor-summary').exists()).toBe(false)
  })
})
