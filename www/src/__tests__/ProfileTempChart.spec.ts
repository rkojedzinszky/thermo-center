import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import ProfileTempChart from '../components/ProfileTempChart.vue'
import type { Profile } from '../api'
import type { ProfileEditMode } from '../composables/useProfileEditMode'
import { makeProfile } from './factories'

function stubRect(element: HTMLElement, width = 86400, height = 200): void {
  element.getBoundingClientRect = () =>
    ({
      left: 0,
      width,
      top: 0,
      height,
      right: width,
      bottom: height,
      x: 0,
      y: 0,
      toJSON: () => ({}),
    }) as DOMRect
}

function dispatchWindow(type: string, clientX?: number, clientY?: number): void {
  const event = new Event(type)
  if (clientX !== undefined) {
    Object.defineProperty(event, 'clientX', { value: clientX })
  }
  if (clientY !== undefined) {
    Object.defineProperty(event, 'clientY', { value: clientY })
  }
  window.dispatchEvent(event)
}

function mountChart(entries: Profile[], mode: ProfileEditMode, selectedId?: number) {
  return mount(ProfileTempChart, { props: { entries, mode, selectedId } })
}

describe('ProfileTempChart', () => {
  it('renders a grey default band and gridlines for an empty schedule', () => {
    const wrapper = mountChart([], 'temperature')
    expect(wrapper.findAll('.default-band')).toHaveLength(1)
    expect(wrapper.findAll('.zone-bar')).toHaveLength(0)
    expect(wrapper.findAll('.temp-handle')).toHaveLength(0)
    expect(wrapper.findAll('.start-handle')).toHaveLength(0)
    expect(wrapper.findAll('.gridline')).toHaveLength(5)
  })

  it('renders a filled bar rising to the target level on the 10–30 °C scale', () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 20)], 'temperature')
    const bar = wrapper.find('.zone-bar')
    // 20 °C sits halfway between 10 and 30; the bar fills from the bottom up.
    expect(bar.attributes('style')).toContain('top: 50%')
    expect(bar.attributes('style')).toContain('--bar-fill')
    expect(bar.attributes('style')).toContain('--bar-color')
  })

  it('keeps one bar per tempered zone and a grey band per default zone', () => {
    const wrapper = mountChart(
      [makeProfile(1, '06:00:00', 21), makeProfile(2, '22:00:00', null)],
      'temperature',
    )
    expect(wrapper.findAll('.zone-bar')).toHaveLength(1)
    expect(wrapper.findAll('.default-band')).toHaveLength(2)
  })

  it('clamps out-of-range temperatures to the plot edges', () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 5)], 'temperature')
    expect(wrapper.find('.zone-bar').attributes('style')).toContain('top: 100%')
    expect(wrapper.find('.temp-handle').attributes('style')).toContain('top: 100%')
  })

  it('marks the selected entry on its bar and start handle', () => {
    const wrapper = mountChart(
      [makeProfile(1, '06:00:00', 21), makeProfile(2, '22:00:00', 18)],
      'temperature',
      2,
    )
    expect(wrapper.findAll('.zone-bar.selected')).toHaveLength(1)
    const active = wrapper
      .findAll('.start-handle')
      .filter((handle) => handle.classes().includes('active'))
    expect(active).toHaveLength(1)
    expect(wrapper.find('.start-handle').attributes('aria-label')).toContain('06:00')
  })

  it('shows zone info in the bar tooltip', () => {
    const wrapper = mountChart(
      [makeProfile(1, '06:00:00', 21), makeProfile(2, '22:00:00', null)],
      'temperature',
    )
    const title = wrapper.find('.zone-bar').attributes('title')
    expect(title).toContain('06:00')
    expect(title).toContain('22:00')
    expect(title).toContain('21')
    expect(wrapper.find('.default-band').attributes('title')).toContain('default')
  })

  it('selects the segment entry on plot click in temperature mode', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)
    await wrapper.find('.chart-plot').trigger('click', { clientX: 12 * 3600 })
    expect(wrapper.emitted('select')?.[0]).toEqual([1])
    expect(wrapper.emitted('addAt')).toBeFalsy()
  })

  it('adds a zone on plot click (snapped) in structure mode', async () => {
    const wrapper = mountChart([], 'structure')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)
    await wrapper.find('.chart-plot').trigger('click', { clientX: 6 * 3600 + 40 })
    expect(wrapper.emitted('addAt')?.[0]).toEqual([6 * 3600])
  })

  it('moves the zone start on handle drag in temperature mode', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.start-handle').trigger('pointerdown')
    expect(wrapper.emitted('select')?.[0]).toEqual([1])

    dispatchWindow('pointermove', 12 * 3600)
    dispatchWindow('pointerup')

    expect(wrapper.emitted('changeStart')?.[0]).toEqual([1, 12 * 3600])
  })

  it('changes the target temperature on line drag in temperature mode', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    // Plot is 200px tall for 20 °C: 22.5 °C sits 75px from the top.
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.temp-handle').trigger('pointerdown')
    expect(wrapper.emitted('select')?.[0]).toEqual([1])

    dispatchWindow('pointermove', undefined, 75)
    dispatchWindow('pointerup')

    expect(wrapper.emitted('changeTemp')?.[0]).toEqual([1, 22.5])
  })

  it('snaps dragged temperatures to the 0.5 °C step', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    // 77px from top ≈ 22.3 °C, snapping to 22.5 °C.
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.temp-handle').trigger('pointerdown')
    dispatchWindow('pointermove', undefined, 77)
    dispatchWindow('pointerup')

    expect(wrapper.emitted('changeTemp')?.[0]).toEqual([1, 22.5])
  })

  it('does not drag the temperature line in structure mode', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'structure')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.temp-handle').trigger('pointerdown')
    dispatchWindow('pointermove', undefined, 75)
    dispatchWindow('pointerup')

    expect(wrapper.emitted('select')?.[0]).toEqual([1])
    expect(wrapper.emitted('changeTemp')).toBeFalsy()
  })

  it('emits nothing on a plain click without drag', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.temp-handle').trigger('pointerdown')
    dispatchWindow('pointerup')

    expect(wrapper.emitted('select')?.[0]).toEqual([1])
    expect(wrapper.emitted('changeTemp')).toBeFalsy()
    expect(wrapper.emitted('changeStart')).toBeFalsy()
  })

  it('cancels a drag on pointercancel without emitting', async () => {
    const wrapper = mountChart([makeProfile(1, '06:00:00', 21)], 'temperature')
    stubRect(wrapper.find('.chart-plot').element as HTMLElement)

    await wrapper.find('.temp-handle').trigger('pointerdown')
    dispatchWindow('pointermove', undefined, 75)
    dispatchWindow('pointercancel')
    dispatchWindow('pointerup')

    expect(wrapper.emitted('changeTemp')).toBeFalsy()
  })
})
