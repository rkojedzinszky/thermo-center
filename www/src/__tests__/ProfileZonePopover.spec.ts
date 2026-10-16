import { describe, it, expect, vi, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import ProfileZonePopover from '../components/ProfileZonePopover.vue'
import type { Profile } from '../api'
import type { ProfileEditMode } from '../composables/useProfileEditMode'
import { makeProfile } from './factories'

const mockSetEntryTemp = vi.fn()
const mockSetEntryDefault = vi.fn()
const mockSetEntryStart = vi.fn()
const mockRemoveEntry = vi.fn()

vi.mock('../composables/useProfiles', () => ({
  useProfiles: () => ({
    setEntryTemp: mockSetEntryTemp,
    setEntryDefault: mockSetEntryDefault,
    setEntryStart: mockSetEntryStart,
    removeEntry: mockRemoveEntry,
    statusFor: () => null,
  }),
}))

function mountPopover(entry: Profile, entries: Profile[], mode: ProfileEditMode) {
  return mount(ProfileZonePopover, { props: { entry, entries, mode } })
}

describe('ProfileZonePopover', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('changes the start time when the slot is free', async () => {
    const entry = makeProfile(1, '06:00:00', 21)
    const wrapper = mountPopover(entry, [entry], 'structure')

    const input = wrapper.find('.popover-time')
    ;(input.element as HTMLInputElement).value = '08:30'
    await input.trigger('change')

    expect(mockSetEntryStart).toHaveBeenCalledWith(1, 8 * 3600 + 30 * 60)
    expect(wrapper.find('.popover-error').exists()).toBe(false)
  })

  it('rejects a conflicting start time and restores the input', async () => {
    const first = makeProfile(1, '06:00:00', 21)
    const second = makeProfile(2, '08:00:00', 18)
    const wrapper = mountPopover(first, [first, second], 'structure')

    const input = wrapper.find('.popover-time')
    ;(input.element as HTMLInputElement).value = '08:00'
    await input.trigger('change')

    expect(mockSetEntryStart).not.toHaveBeenCalled()
    expect(wrapper.find('.popover-error').exists()).toBe(true)
    expect((input.element as HTMLInputElement).value).toBe('06:00')
  })

  it('exposes temperature controls in both modes', () => {
    const entry = makeProfile(1, '06:00:00', 21)
    expect(mountPopover(entry, [entry], 'temperature').find('.popover-slider').exists()).toBe(true)
    expect(mountPopover(entry, [entry], 'structure').find('.popover-slider').exists()).toBe(true)
  })

  it('hides start-time and delete controls in temperature mode', () => {
    const entry = makeProfile(1, '06:00:00', 21)
    const wrapper = mountPopover(entry, [entry], 'temperature')
    expect(wrapper.find('.popover-time').exists()).toBe(false)
    expect(wrapper.find('.popover-delete').exists()).toBe(false)
  })

  it('sets the temperature from the slider', async () => {
    const entry = makeProfile(1, '06:00:00', 21)
    const wrapper = mountPopover(entry, [entry], 'temperature')

    const slider = wrapper.find('.popover-slider')
    ;(slider.element as HTMLInputElement).value = '24.5'
    await slider.trigger('input')

    expect(mockSetEntryTemp).toHaveBeenCalledWith(1, 24.5)
  })
})
