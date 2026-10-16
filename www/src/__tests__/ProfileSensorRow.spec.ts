import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { ref, nextTick } from 'vue'
import { mount, flushPromises, enableAutoUnmount } from '@vue/test-utils'
import ProfileSensorRow from '../components/ProfileSensorRow.vue'
import ProfileTempChart from '../components/ProfileTempChart.vue'
import type { Profile } from '../api'
import type { ProfileEditMode } from '../composables/useProfileEditMode'
import { makeControl, makeDayType, makeProfile } from './factories'

enableAutoUnmount(afterEach)

const mockAddEntry = vi.fn()
const mockSetEntryStart = vi.fn()
const mockSetEntryTemp = vi.fn()
const mockSetEntryDefault = vi.fn()
const mockRemoveEntry = vi.fn()
const mockUpdateControl = vi.fn()
const mockSaveTick = ref(0)
const mockEntries = ref<Profile[]>([])
const mockIsReloading = ref(false)

vi.mock('../composables/useScheduleReload', () => ({
  useScheduleReload: () => ({
    isReloading: mockIsReloading,
    reloadInitialView: vi.fn(),
    reloadSensorView: vi.fn(),
    reloadDaytypeView: vi.fn(),
    resolveSensorId: vi.fn(),
    resolveDaytypeId: vi.fn(),
  }),
}))

vi.mock('../composables/useProfiles', () => ({
  useProfiles: () => ({
    entriesFor: () => mockEntries.value,
    statusFor: () => null,
    saveTick: mockSaveTick,
    addEntry: mockAddEntry,
    setEntryStart: mockSetEntryStart,
    setEntryTemp: mockSetEntryTemp,
    setEntryDefault: mockSetEntryDefault,
    removeEntry: mockRemoveEntry,
    lastError: ref(null),
    clearError: vi.fn(),
    loadProfiles: vi.fn(),
    loadEntriesFor: vi.fn(),
  }),
}))

vi.mock('../composables/useControls', () => ({
  useControls: () => ({
    orderedControls: ref([]),
    loadControls: vi.fn(),
    updateControlDirect: vi.fn(),
    updateControl: mockUpdateControl,
  }),
}))

const control = makeControl(1, 'Living Room')

const daytype = makeDayType(1, 'workday')

function mountRow(mode: ProfileEditMode, extra: Record<string, unknown> = {}) {
  return mount(ProfileSensorRow, { props: { control, daytype, mode, ...extra } })
}

function emitTimeline(wrapper: ReturnType<typeof mountRow>, event: string, ...args: unknown[]) {
  wrapper.findComponent(ProfileTempChart).vm.$emit(event, ...args)
  return nextTick()
}

describe('ProfileSensorRow', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    mockEntries.value = []
    mockSaveTick.value = 0
    mockIsReloading.value = false
  })

  it('renders the sensor name and zone count', () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21), makeProfile(2, '22:00:00', 18)]
    const wrapper = mountRow('temperature')

    expect(wrapper.find('.sensor-name').text()).toBe('Living Room')
    expect(wrapper.find('.zone-count').text()).toContain('· 2')
    expect(wrapper.text()).toContain('21.50 °C')
  })

  it('only shows the add button in structure mode', () => {
    expect(mountRow('temperature').find('.add-button').exists()).toBe(false)
    expect(mountRow('structure').find('.add-button').exists()).toBe(true)
  })

  it('shows a loading hint over the empty bar while refetching', () => {
    mockIsReloading.value = true
    const wrapper = mountRow('temperature')

    expect(wrapper.find('.row-loading').exists()).toBe(true)
    expect(wrapper.findComponent(ProfileTempChart).exists()).toBe(true)
  })

  it('keeps the timeline once entries arrive, even while reloading', () => {
    mockIsReloading.value = true
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')

    expect(wrapper.find('.row-loading').exists()).toBe(false)
    expect(wrapper.findComponent(ProfileTempChart).exists()).toBe(true)
  })

  it('forwards a dragged target temperature to the store', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')
    await emitTimeline(wrapper, 'changeTemp', 1, 22.5)

    expect(mockSetEntryTemp).toHaveBeenCalledWith(1, 22.5)
  })

  it('opens the zone popover when a zone is selected', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')

    expect(wrapper.find('.zone-popover').exists()).toBe(false)
    await emitTimeline(wrapper, 'select', 1)
    expect(wrapper.find('.zone-popover').exists()).toBe(true)
  })

  it('saves a temperature change from the popover', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')
    await emitTimeline(wrapper, 'select', 1)

    const input = wrapper.find('.popover-number')
    ;(input.element as HTMLInputElement).value = '23.5'
    await input.trigger('change')

    expect(mockSetEntryTemp).toHaveBeenCalledWith(1, 23.5)
  })

  it('sets a zone to default from the popover', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')
    await emitTimeline(wrapper, 'select', 1)

    await wrapper.find('.popover-default input').setValue(true)

    expect(mockSetEntryDefault).toHaveBeenCalledWith(1)
  })

  it('deletes a zone from the popover in structure mode', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('structure')
    await emitTimeline(wrapper, 'select', 1)

    expect(wrapper.find('.popover-delete').exists()).toBe(true)
    await wrapper.find('.popover-delete').trigger('click')

    expect(mockRemoveEntry).toHaveBeenCalledWith(1)
    expect(wrapper.find('.zone-popover').exists()).toBe(false)
  })

  it('hides start-time and delete controls in temperature mode', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21)]
    const wrapper = mountRow('temperature')
    await emitTimeline(wrapper, 'select', 1)

    expect(wrapper.find('.popover-time').exists()).toBe(false)
    expect(wrapper.find('.popover-delete').exists()).toBe(false)
  })

  it('adds a zone in structure mode and opens it', async () => {
    mockAddEntry.mockImplementation(async () => {
      const created = makeProfile(9, '06:00:00', 20)
      mockEntries.value = [...mockEntries.value, created]
      return created
    })

    const wrapper = mountRow('structure')
    await wrapper.find('.add-button').trigger('click')
    await flushPromises()

    expect(mockAddEntry).toHaveBeenCalledWith(
      control.resourceUri,
      daytype.resourceUri,
      6 * 3600,
      20,
    )
    expect(wrapper.find('.zone-popover').exists()).toBe(true)
  })

  it('warns and skips a conflicting start time', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21), makeProfile(2, '08:00:00', 18)]
    const wrapper = mountRow('structure')

    await emitTimeline(wrapper, 'changeStart', 1, 8 * 3600)

    expect(mockSetEntryStart).not.toHaveBeenCalled()
    expect(wrapper.find('.row-error').exists()).toBe(true)
  })

  it('applies a non-conflicting start time', async () => {
    mockEntries.value = [makeProfile(1, '06:00:00', 21), makeProfile(2, '08:00:00', 18)]
    const wrapper = mountRow('structure')

    await emitTimeline(wrapper, 'changeStart', 1, 12 * 3600)

    expect(mockSetEntryStart).toHaveBeenCalledWith(1, 12 * 3600)
  })

  it('refreshes the device target after a successful save', async () => {
    const wrapper = mountRow('temperature')
    mockSaveTick.value++
    await nextTick()
    expect(mockUpdateControl).toHaveBeenCalledWith(1)
    expect(wrapper.find('.sensor-name').exists()).toBe(true)
  })

  it('uses a custom title when provided', () => {
    const wrapper = mountRow('temperature', { title: 'workday' })
    expect(wrapper.find('.sensor-name').text()).toBe('workday')
  })

  it('hides meta and skips the device refresh when showMeta is false', async () => {
    const wrapper = mountRow('temperature', { showMeta: false })

    expect(wrapper.find('.sensor-meta').exists()).toBe(false)
    mockSaveTick.value++
    await nextTick()
    expect(mockUpdateControl).not.toHaveBeenCalled()
  })

  it('renders the temperature reading below the name inside the label column', () => {
    const wrapper = mountRow('temperature')

    const label = wrapper.find('.row-label')
    expect(label.find('.sensor-name').exists()).toBe(true)
    expect(label.find('.sensor-meta').exists()).toBe(true)
    expect(label.find('.sensor-meta').text()).toContain('21.50 °C')
  })
})
