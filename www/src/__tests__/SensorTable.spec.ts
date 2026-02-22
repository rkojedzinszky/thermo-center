import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { mount } from '@vue/test-utils'
import SensorTable from '../components/SensorTable.vue'
import type { THSensor } from '../api'

const NOW_UNIX = 1_700_000_000

const sensors: THSensor[] = [
  {
    id: 1,
    name: 'Kitchen',
    temperature: 23.2,
    humidity: 48.0,
    lastTsf: NOW_UNIX - 45,
    vcc: 3.0,
    rssi: -70,
    lqi: 180,
    interval: 120,
    lastSeq: 9999,
    valid: true,
    sensorResync: null,
  },
  {
    id: 2,
    name: 'Attic',
    temperature: 18.0,
    humidity: 62.5,
    lastTsf: NOW_UNIX - 700, // inactive
    vcc: 2.9,
    rssi: -80,
    lqi: 120,
    interval: 120,
    lastSeq: 8000,
    valid: false,
    sensorResync: 'abc',
  },
]

beforeEach(() => {
  vi.useFakeTimers()
  vi.setSystemTime(NOW_UNIX * 1000)
})

afterEach(() => {
  vi.useRealTimers()
})

describe('SensorTable', () => {
  describe('layout & structure', () => {
    it('renders a table element', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.find('table').exists()).toBe(true)
    })

    it('renders thead with column headers', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const ths = wrapper.findAll('th.th')
      expect(ths.length).toBeGreaterThan(0)
    })

    it('renders the correct column headers', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const headerText = wrapper.findAll('th.th').map((th) => th.text())
      expect(headerText).toContain('Name')
      expect(headerText).toContain('Temp (°C)')
      expect(headerText).toContain('Humidity (%)')
      expect(headerText).toContain('Last Data')
    })

    it('renders a row for each sensor', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.findAll('tbody tr.table-row')).toHaveLength(sensors.length)
    })

    it('shows empty state when no sensors provided', () => {
      const wrapper = mount(SensorTable, { props: { sensors: [] } })
      const emptyRow = wrapper.find('.empty-row')
      expect(emptyRow.exists()).toBe(true)
      expect(emptyRow.text()).toContain('No sensors found')
    })

    it('renders drag handle cell in each row', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const handles = wrapper.findAll('.drag-handle')
      expect(handles).toHaveLength(sensors.length)
    })
  })

  describe('readability – cell data', () => {
    it('displays sensor name in a row', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.text()).toContain('Kitchen')
    })

    it('displays formatted temperature', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.text()).toContain('23.2')
    })

    it('displays formatted humidity', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.text()).toContain('48')
    })

    it('displays age for last_tsf', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      expect(wrapper.text()).toContain('45 seconds ago')
    })

    it('shows "—" for null fields on the second sensor', () => {
      const sensorWithNulls: THSensor[] = [
        {
          id: 3,
          name: 'Shed',
          temperature: null,
          humidity: null,
          lastTsf: null,
          valid: null,
          sensorResync: null,
        },
      ]
      const wrapper = mount(SensorTable, { props: { sensors: sensorWithNulls } })
      expect(wrapper.text()).toContain('—')
    })

    it('does NOT render "valid" column', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const headerText = wrapper.findAll('th').map((th) => th.text().toLowerCase())
      expect(headerText).not.toContain('valid')
    })

    it('does NOT render "sensor_resync" column', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const headerText = wrapper.findAll('th').map((th) => th.text().toLowerCase())
      expect(headerText.join(' ')).not.toContain('resync')
    })
  })

  describe('inactive state', () => {
    it('adds "inactive" class to the row of an inactive sensor', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      // Second sensor is inactive
      expect(rows[1]!.classes()).toContain('inactive')
    })

    it('does NOT add "inactive" class to the row of an active sensor', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      expect(rows[0]!.classes()).not.toContain('inactive')
    })

    it('applies stale-text class to the age cell of inactive sensor', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const staleElements = wrapper.findAll('.stale-text')
      expect(staleElements.length).toBeGreaterThan(0)
    })

    it('does not apply stale-text to active sensor age', () => {
      const activeSensorOnly: THSensor[] = [sensors[0]!]
      const wrapper = mount(SensorTable, { props: { sensors: activeSensorOnly } })
      const staleElements = wrapper.findAll('.stale-text')
      expect(staleElements).toHaveLength(0)
    })
  })

  describe('drag-and-drop reorder', () => {
    it('all rows are draggable', () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      rows.forEach((row) => {
        expect(row.attributes('draggable')).toBe('true')
      })
    })

    it('emits "reorder" event when dragging from one row to another', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!
      const row1 = rows[1]!

      await row0.trigger('dragstart', {
        dataTransfer: { setData: vi.fn(), effectAllowed: 'move' },
      })
      await row1.trigger('dragover')
      await row1.trigger('drop')

      expect(wrapper.emitted('reorder')).toBeTruthy()
      expect(wrapper.emitted('reorder')?.[0]).toEqual([0, 1])
    })

    it('does NOT emit reorder when dropped on same index', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!

      await row0.trigger('dragstart', {
        dataTransfer: { setData: vi.fn(), effectAllowed: 'move' },
      })
      await row0.trigger('dragover')
      await row0.trigger('drop')

      expect(wrapper.emitted('reorder')).toBeFalsy()
    })

    it('adds "drag-over" class when dragging over a row', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!
      const row1 = rows[1]!

      await row0.trigger('dragstart', {
        dataTransfer: { setData: vi.fn(), effectAllowed: 'move' },
      })
      await row1.trigger('dragover')

      expect(row1.classes()).toContain('drag-over')
    })

    it('clears drag-over class after drag ends', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!
      const row1 = rows[1]!

      await row0.trigger('dragstart', {
        dataTransfer: { setData: vi.fn(), effectAllowed: 'move' },
      })
      await row1.trigger('dragover')
      await row0.trigger('dragend')

      expect(row1.classes()).not.toContain('drag-over')
    })
  })

  describe('touch drag reorder', () => {
    it('emits "reorder" event when touch-dragging from one row to another', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!
      const row1 = rows[1]!

      // Mock elementFromPoint to return the second row's DOM element
      document.elementFromPoint = vi.fn().mockReturnValue(row1.element)

      await row0.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await wrapper.find('.table-wrapper').trigger('touchmove', {
        touches: [{ clientX: 50, clientY: 100 }],
      })
      await wrapper.find('.table-wrapper').trigger('touchend', {
        changedTouches: [{ clientX: 50, clientY: 100 }],
      })

      expect(wrapper.emitted('reorder')).toBeTruthy()
      expect(wrapper.emitted('reorder')?.[0]).toEqual([0, 1])

      delete (document as any).elementFromPoint
    })

    it('does NOT emit reorder when touch ends on same row', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!

      document.elementFromPoint = vi.fn().mockReturnValue(row0.element)

      await row0.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await wrapper.find('.table-wrapper').trigger('touchmove', {
        touches: [{ clientX: 50, clientY: 55 }],
      })
      await wrapper.find('.table-wrapper').trigger('touchend', {
        changedTouches: [{ clientX: 50, clientY: 55 }],
      })

      expect(wrapper.emitted('reorder')).toBeFalsy()

      delete (document as any).elementFromPoint
    })

    it('adds "drag-over" class to target row during touch drag', async () => {
      const wrapper = mount(SensorTable, { props: { sensors } })
      const rows = wrapper.findAll('tbody tr.table-row')
      const row0 = rows[0]!
      const row1 = rows[1]!

      document.elementFromPoint = vi.fn().mockReturnValue(row1.element)

      await row0.trigger('touchstart', { touches: [{ clientX: 50, clientY: 50 }] })
      await wrapper.find('.table-wrapper').trigger('touchmove', {
        touches: [{ clientX: 50, clientY: 100 }],
      })

      expect(row1.classes()).toContain('drag-over')

      delete (document as any).elementFromPoint
    })
  })

  describe('timestamp formatting', () => {
    it('shows "No data" for null lastTsf', () => {
      const s: THSensor[] = [
        { id: 10, name: 'Test', lastTsf: null, valid: null, sensorResync: null },
      ]
      const wrapper = mount(SensorTable, { props: { sensors: s } })
      expect(wrapper.text()).toContain('No data')
    })

    it('shows minute-format for 75 seconds ago', () => {
      const s: THSensor[] = [
        { id: 11, name: 'Test', lastTsf: NOW_UNIX - 75, valid: null, sensorResync: null },
      ]
      const wrapper = mount(SensorTable, { props: { sensors: s } })
      expect(wrapper.text()).toContain('1 minute ago')
    })
  })
})
