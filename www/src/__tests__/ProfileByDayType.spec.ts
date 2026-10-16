import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { ref } from 'vue'
import { mount, enableAutoUnmount } from '@vue/test-utils'
import ProfileByDayType from '../components/ProfileByDayType.vue'
import type { Control, DayType } from '../api'
import { makeControl, makeDayType } from './factories'

enableAutoUnmount(afterEach)

const mockReloadDaytypeView = vi.fn()
const mockControls = ref<Control[]>([])
const mockDayTypes = ref<DayType[]>([])

vi.mock('../composables/useControls', () => ({
  useControls: () => ({
    orderedControls: mockControls,
    loadControls: vi.fn(),
    updateControlDirect: vi.fn(),
    updateControl: vi.fn(),
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
    reloadSensorView: vi.fn(),
    reloadDaytypeView: mockReloadDaytypeView,
    resolveSensorId: vi.fn(),
    resolveDaytypeId: vi.fn(),
    isReloading: ref(false),
  }),
}))

describe('ProfileByDayType', () => {
  beforeEach(() => {
    localStorage.clear()
    vi.clearAllMocks()
    mockControls.value = [makeControl(1, 'Living Room')]
    mockDayTypes.value = [makeDayType(1, 'workday'), makeDayType(2, 'Sunday')]
  })

  it('renders a tab per day type and a row per control', () => {
    const wrapper = mount(ProfileByDayType, { props: { mode: 'temperature' } })

    expect(wrapper.findAll('.tab')).toHaveLength(2)
    expect(wrapper.find('.tab-count').exists()).toBe(false)
    expect(wrapper.findAll('.sensor-row')).toHaveLength(1)
  })

  it('refetches the filtered slice when switching day types', async () => {
    const wrapper = mount(ProfileByDayType, { props: { mode: 'temperature' } })

    await wrapper.findAll('.tab')[1]!.trigger('click')

    expect(mockReloadDaytypeView).toHaveBeenCalledTimes(1)
    expect(mockReloadDaytypeView).toHaveBeenCalledWith(2)
  })

  it('does not refetch when clicking the active day type', async () => {
    const wrapper = mount(ProfileByDayType, { props: { mode: 'temperature' } })
    mockReloadDaytypeView.mockClear()

    await wrapper.findAll('.tab')[0]!.trigger('click')

    expect(mockReloadDaytypeView).not.toHaveBeenCalled()
  })
})
