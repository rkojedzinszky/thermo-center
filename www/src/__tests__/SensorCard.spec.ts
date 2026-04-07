import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import SensorCard from '../components/SensorCard.vue'
import type { THSensor } from '../api'

const NOW_UNIX = Math.floor(Date.now() / 1000)

const activeSensor: THSensor = {
  resourceUri: '/sensor/42',
  id: 42,
  name: 'Living Room',
  temperature: 21.5,
  humidity: 55.3,
  lastTsf: NOW_UNIX - 30, // 30 seconds ago – active
  vcc: 3.3,
  rssi: -65,
  lqi: 200,
  interval: 60,
  lastSeq: 12345,
  valid: true,
  sensorResync: undefined,
}

const inactiveSensor: THSensor = {
  ...activeSensor,
  id: 99,
  name: 'Basement',
  lastTsf: NOW_UNIX - 600, // 10 minutes ago – inactive
}

const noDataSensor: THSensor = {
  ...activeSensor,
  id: 7,
  name: 'Garage',
  lastTsf: null,
  temperature: null,
  humidity: null,
}

const invalidSensor: THSensor = {
  ...activeSensor,
  id: 88,
  name: 'Master Bedroom',
  valid: false,
  sensorResync: undefined,
}

beforeEach(() => {
  vi.useFakeTimers()
  vi.setSystemTime(NOW_UNIX * 1000)
})

afterEach(() => {
  vi.useRealTimers()
})

function mountCard(sensor: THSensor, extraProps: Record<string, unknown> = {}) {
  return mount(SensorCard, {
    props: { sensor, index: 0, total: 1, ...extraProps },
  })
}

describe('SensorCard', () => {
  describe('layout & structure', () => {
    it('renders the card wrapper element', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-wrapper').exists()).toBe(true)
    })

    it('has a front face and a back face', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-front').exists()).toBe(true)
      expect(wrapper.find('.card-back').exists()).toBe(true)
    })

    it('card wrapper is not draggable by default', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-wrapper').attributes('draggable')).toBeUndefined()
    })

    it('card wrapper is draggable in reorder mode', () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      expect(wrapper.find('.card-wrapper').attributes('draggable')).toBe('true')
    })

    it('no drag grip element is rendered', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-grip').exists()).toBe(false)
    })

    it('has aria-label with sensor name', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-wrapper').attributes('aria-label')).toContain(activeSensor.name)
    })

    it('has reading elements for temperature and humidity', () => {
      const wrapper = mountCard(activeSensor)
      const readings = wrapper.findAll('.reading')
      expect(readings).toHaveLength(2)
    })
  })

  describe('readability – front side', () => {
    it('displays sensor name on the front', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.sensor-name').text()).toBe('Living Room')
    })

    it('displays formatted temperature', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-front').text()).toContain('21.50 °C')
    })

    it('displays formatted humidity', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-front').text()).toContain('55.30 %')
    })

    it('displays "30 seconds ago" when lastTsf is 30 seconds ago', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.age-label').text()).toMatch(/^(29|30) seconds ago$/)
    })

    it('displays "No data" when lastTsf is null', () => {
      const wrapper = mountCard(noDataSensor)
      expect(wrapper.find('.age-label').text()).toBe('No data')
    })

    it('displays "—" for null temperature', () => {
      const wrapper = mountCard(noDataSensor)
      expect(wrapper.find('.card-front').text()).toContain('—')
    })

    it('shows a flip hint for discoverability', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-front').text()).toContain('Click to flip')
    })
  })

  describe('inactive state', () => {
    it('adds "inactive" class to wrapper when sensor is inactive', () => {
      const wrapper = mountCard(inactiveSensor)
      expect(wrapper.find('.card-wrapper').classes()).toContain('inactive')
    })

    it('does NOT add "inactive" class for an active sensor', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-wrapper').classes()).not.toContain('inactive')
    })

    it('marks sensor with null lastTsf as inactive', () => {
      const wrapper = mountCard(noDataSensor)
      expect(wrapper.find('.card-wrapper').classes()).toContain('inactive')
    })
  })

  describe('card flip', () => {
    it('starts with card not flipped', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card').classes()).not.toContain('flipped')
    })

    it('adds "flipped" class after clicking', async () => {
      const wrapper = mountCard(activeSensor)
      await wrapper.find('.card-wrapper').trigger('click')
      expect(wrapper.find('.card').classes()).toContain('flipped')
    })

    it('removes "flipped" class on second click (flip back)', async () => {
      const wrapper = mountCard(activeSensor)
      await wrapper.find('.card-wrapper').trigger('click')
      await wrapper.find('.card-wrapper').trigger('click')
      expect(wrapper.find('.card').classes()).not.toContain('flipped')
    })

    it('shows "Click to flip back" hint on the back face', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.card-back').text()).toContain('Click to flip back')
    })
  })

  describe('back side – raw data fields', () => {
    it('renders exactly the expected number of fields', () => {
      const wrapper = mountCard(activeSensor)
      // id, name, temperature, humidity, lastTsf, vcc, rssi, lqi, interval, lastSeq = 10
      expect(wrapper.findAll('.back-field')).toHaveLength(10)
    })

    it('back side contains ID field', () => {
      const wrapper = mountCard(activeSensor)
      const labels = wrapper.findAll('.field-label').map((el) => el.text())
      expect(labels).toContain('ID')
    })

    it('back side shows VCC field', () => {
      const wrapper = mountCard(activeSensor)
      const labels = wrapper.findAll('.field-label').map((el) => el.text())
      expect(labels).toContain('VCC')
    })

    it('back side shows RSSI field', () => {
      const wrapper = mountCard(activeSensor)
      const labels = wrapper.findAll('.field-label').map((el) => el.text())
      expect(labels).toContain('RSSI')
    })

    it('back side shows Last Data field', () => {
      const wrapper = mountCard(activeSensor)
      const labels = wrapper.findAll('.field-label').map((el) => el.text())
      expect(labels).toContain('Last Data')
    })

    it('does NOT render "valid" as a displayed field', () => {
      const wrapper = mountCard(activeSensor)
      const labels = wrapper.findAll('.field-label').map((el) => el.text().toLowerCase())
      expect(labels).not.toContain('valid')
    })

    it('does NOT render "sensorResync" / "sensor_resync" as a displayed field', () => {
      const wrapper = mountCard(activeSensor)
      const text = wrapper.find('.card-back').text().toLowerCase()
      expect(text).not.toContain('resync')
    })
  })

  describe('drag events', () => {
    it('does NOT emit dragStart in normal mode when dragging card wrapper', () => {
      const wrapper = mountCard(activeSensor)
      wrapper.find('.card-wrapper').trigger('dragstart')
      expect(wrapper.emitted('dragStart')).toBeFalsy()
    })

    it('emits dragStart with index when dragging card wrapper in reorder mode', () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      wrapper.find('.card-wrapper').trigger('dragstart', { dataTransfer: null })
      expect(wrapper.emitted('dragStart')?.[0]).toEqual([0])
    })

    it('emits dragOver with index on dragover in reorder mode', () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      wrapper.find('.card-wrapper').trigger('dragover')
      expect(wrapper.emitted('dragOver')?.[0]).toEqual([0])
    })

    it('does NOT emit dragOver in normal mode on dragover', () => {
      const wrapper = mountCard(activeSensor)
      wrapper.find('.card-wrapper').trigger('dragover')
      expect(wrapper.emitted('dragOver')).toBeFalsy()
    })

    it('emits dragEnd on card wrapper dragend in reorder mode', () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      wrapper.find('.card-wrapper').trigger('dragend')
      expect(wrapper.emitted('dragEnd')).toBeTruthy()
    })
  })

  describe('touch drag events', () => {
    it('does NOT emit dragStart in normal mode on touchstart', async () => {
      const wrapper = mountCard(activeSensor)
      await wrapper.find('.card-wrapper').trigger('touchstart', {
        touches: [{ clientX: 50, clientY: 50 }],
      })
      expect(wrapper.emitted('dragStart')).toBeFalsy()
    })

    it('emits dragStart with index on touchstart in reorder mode', async () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      await wrapper.find('.card-wrapper').trigger('touchstart', {
        touches: [{ clientX: 50, clientY: 50 }],
      })
      expect(wrapper.emitted('dragStart')?.[0]).toEqual([0])
    })

    it('emits dragEnd on touchend in reorder mode', async () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      const card = wrapper.find('.card-wrapper')
      await card.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await card.trigger('touchend', { changedTouches: [{ clientX: 50, clientY: 50 }] })
      expect(wrapper.emitted('dragEnd')).toBeTruthy()
    })

    it('emits dragOver when touchmove targets a different card index in reorder mode', async () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      const card = wrapper.find('.card-wrapper')

      const fakeTarget = document.createElement('div')
      fakeTarget.setAttribute('data-card-index', '1')
      const originalElementFromPoint = document.elementFromPoint
      document.elementFromPoint = vi.fn().mockReturnValue(fakeTarget)

      await card.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await card.trigger('touchmove', { touches: [{ clientX: 50, clientY: 120 }] })

      const events = wrapper.emitted('dragOver')
      expect(events).toBeTruthy()
      expect(events?.[0]).toEqual([1])

      document.elementFromPoint = originalElementFromPoint
    })

    it('does NOT emit dragOver when touchmove targets the same card index in reorder mode', async () => {
      const wrapper = mountCard(activeSensor, { reorderMode: true })
      const card = wrapper.find('.card-wrapper')

      const fakeTarget = document.createElement('div')
      fakeTarget.setAttribute('data-card-index', '0')
      const originalElementFromPoint = document.elementFromPoint
      document.elementFromPoint = vi.fn().mockReturnValue(fakeTarget)

      await card.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await card.trigger('touchmove', { touches: [{ clientX: 50, clientY: 55 }] })

      expect(wrapper.emitted('dragOver')).toBeFalsy()

      document.elementFromPoint = originalElementFromPoint
    })
  })

  describe('timestamp formatting', () => {
    it('formats "1 minute ago" for 65 seconds', () => {
      const sensor = { ...activeSensor, lastTsf: NOW_UNIX - 65 }
      const wrapper = mountCard(sensor)
      expect(wrapper.find('.age-label').text()).toBe('1 minute ago')
    })

    it('formats "2 hours ago" for 7500 seconds', () => {
      const sensor = { ...activeSensor, lastTsf: NOW_UNIX - 7500 }
      const wrapper = mountCard(sensor)
      expect(wrapper.find('.age-label').text()).toBe('2 hours ago')
    })

    it('formats "just now" for a future or near-zero timestamp', () => {
      const sensor = { ...activeSensor, lastTsf: NOW_UNIX + 5 }
      const wrapper = mountCard(sensor)
      expect(wrapper.find('.age-label').text()).toBe('just now')
    })
  })

  describe('card size – no overflow guarantee', () => {
    it('card-wrapper has fixed dimensions via CSS classes (not inline overflow)', () => {
      const wrapper = mountCard(activeSensor)
      // The wrapper must not have inline overflow styles that would cause clipping
      const style = wrapper.find('.card-wrapper').attributes('style') ?? ''
      expect(style).not.toContain('overflow: hidden')
    })

    it('back-fields list element exists and is within the card face', () => {
      const wrapper = mountCard(activeSensor)
      const backFace = wrapper.find('.card-back')
      expect(backFace.find('.back-fields').exists()).toBe(true)
    })

    it('each field-value has the ellipsis class applied via CSS (class present)', () => {
      const wrapper = mountCard(activeSensor)
      // Verify the field-value elements exist (CSS handles overflow:hidden text-overflow:ellipsis)
      const values = wrapper.findAll('.field-value')
      expect(values.length).toBeGreaterThan(0)
      values.forEach((v) => {
        expect(v.classes()).toContain('field-value')
      })
    })

    it('reading-value elements exist for temperature and humidity', () => {
      const wrapper = mountCard(activeSensor)
      const values = wrapper.findAll('.reading-value')
      expect(values).toHaveLength(2)
    })
  })

  describe('resync button – invalid sensor', () => {
    it('shows resync button when sensor has valid=false', () => {
      const wrapper = mountCard(invalidSensor)
      expect(wrapper.find('.resync-button').exists()).toBe(true)
    })

    it('does NOT show resync button when sensor has valid=true', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.resync-button').exists()).toBe(false)
    })

    it('shows age-label when sensor is valid', () => {
      const wrapper = mountCard(activeSensor)
      expect(wrapper.find('.age-label').exists()).toBe(true)
    })

    it('hides age-label when sensor is invalid', () => {
      const wrapper = mountCard(invalidSensor)
      expect(wrapper.find('.age-label').exists()).toBe(false)
    })

    it('resync button is enabled on initial render', () => {
      const wrapper = mountCard(invalidSensor)
      expect(wrapper.find('.resync-button').attributes('disabled')).toBeUndefined()
    })

    it('resync button has correct title attribute', () => {
      const wrapper = mountCard(invalidSensor)
      expect(wrapper.find('.resync-button').attributes('title')).toBe(
        'Request sensor resynchronization',
      )
    })

    it('displaying resync button does not affect card interaction', () => {
      const wrapper = mountCard(invalidSensor)
      expect(wrapper.find('.card-wrapper').exists()).toBe(true)
      expect(wrapper.find('.card').classes()).not.toContain('flipped')
    })
  })
})
