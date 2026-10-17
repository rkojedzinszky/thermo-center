import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { ref } from 'vue'
import { mount, flushPromises, enableAutoUnmount } from '@vue/test-utils'
import DayTypeCalendar from '../components/DayTypeCalendar.vue'
import type { DayType } from '../api'
import { makeDayType } from './factories'
import { monthGrid, monthLabel, shiftMonth, todayIsoDay, toIsoDay } from '../utils/calendarDays'

enableAutoUnmount(afterEach)

const SATURDAY_URI = '/api/v1/daytype/2/'
const SUNDAY_URI = '/api/v1/daytype/3/'

const mockDayTypes = ref<DayType[]>([])
const mockEntries = ref<Map<string, string>>(new Map())
const mockStatuses = ref<Map<string, string>>(new Map())
const mockLastError = ref<string | null>(null)
const mockIsLoading = ref(false)

const mockLoadRange = vi.fn()
const mockSetDayType = vi.fn()
const mockClearError = vi.fn()

vi.mock('../composables/useDayTypes', () => ({
  useDayTypes: () => ({
    dayTypes: mockDayTypes,
    loadDayTypes: vi.fn(),
  }),
}))

vi.mock('../composables/useCalendar', () => ({
  useCalendar: () => ({
    isLoading: mockIsLoading,
    lastError: mockLastError,
    dayTypeUri: (day: string) => mockEntries.value.get(day) ?? null,
    statusFor: (day: string) => mockStatuses.value.get(day) ?? null,
    loadRange: mockLoadRange,
    setDayType: mockSetDayType,
    clearError: mockClearError,
  }),
}))

const now = new Date()
const currentGrid = monthGrid(now.getFullYear(), now.getMonth())
const next = shiftMonth(now.getFullYear(), now.getMonth(), 1)
const nextGrid = monthGrid(next.year, next.month)
// Index 20 always falls inside the month itself (months start within the first 6 cells).
const nextMonthDay = toIsoDay(nextGrid.days[20]!)
// Last cell of the initial grid: always rendered (may belong to the next month).
const lastVisibleDay = toIsoDay(currentGrid.days[currentGrid.days.length - 1]!)

describe('DayTypeCalendar', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    mockDayTypes.value = [
      makeDayType(1, 'workday'),
      makeDayType(2, 'Saturday'),
      makeDayType(3, 'Sunday'),
    ]
    mockEntries.value = new Map()
    mockStatuses.value = new Map()
    mockLastError.value = null
    mockIsLoading.value = false
  })

  it('loads the visible month and renders a week of headers plus 42 cells', async () => {
    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    expect(mockLoadRange).toHaveBeenCalledWith(currentGrid.start, currentGrid.end)
    expect(wrapper.findAll('.weekday')).toHaveLength(7)
    expect(wrapper.findAll('.day-cell')).toHaveLength(42)
    expect(wrapper.find('.month-label').text()).toBe(monthLabel(now.getFullYear(), now.getMonth()))
  })

  it('renders a chip for days with a day type and marks today', async () => {
    mockEntries.value = new Map([[lastVisibleDay, SUNDAY_URI]])

    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    const chipCell = wrapper.find(`.day-cell[aria-label^="${lastVisibleDay}"]`)
    expect(chipCell.find('.daytype-chip').text()).toBe('Sunday')

    const today = todayIsoDay()
    const todayCell = wrapper.find(`.day-cell[aria-label^="${today}"]`)
    expect(todayCell.classes()).toContain('today')

    // Leading/trailing days from other months have no entry.
    expect(wrapper.find('.daytype-none').exists()).toBe(true)
  })

  it('disables past days', async () => {
    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    const pastCells = wrapper.findAll('.day-cell.past')
    expect(pastCells.length).toBeGreaterThan(0)
    expect((pastCells[0]!.element as HTMLButtonElement).disabled).toBe(true)
  })

  it('disables days without a loaded entry', async () => {
    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    const cell = wrapper.find(`.day-cell[aria-label^="${lastVisibleDay}"]`)
    expect(cell.classes()).toContain('no-entry')
    expect((cell.element as HTMLButtonElement).disabled).toBe(true)

    await cell.trigger('click')
    expect(wrapper.find('.picker').exists()).toBe(false)
  })

  it('navigates months and returns with Today', async () => {
    const wrapper = mount(DayTypeCalendar)
    await flushPromises()
    mockLoadRange.mockClear()

    await wrapper.find('[aria-label="Next month"]').trigger('click')

    expect(wrapper.find('.month-label').text()).toBe(monthLabel(next.year, next.month))
    expect(mockLoadRange).toHaveBeenCalledWith(nextGrid.start, nextGrid.end)

    await wrapper.find('.today-button').trigger('click')

    expect(wrapper.find('.month-label').text()).toBe(monthLabel(now.getFullYear(), now.getMonth()))
    expect(mockLoadRange).toHaveBeenLastCalledWith(currentGrid.start, currentGrid.end)
  })

  it('opens a picker on a day and saves the chosen day type', async () => {
    mockEntries.value = new Map([[nextMonthDay, SATURDAY_URI]])

    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    // Move to the next month so the picked day is never in the past.
    await wrapper.find('[aria-label="Next month"]').trigger('click')

    const cell = wrapper.find(`.day-cell[aria-label^="${nextMonthDay}"]`)
    await cell.trigger('click')

    const options = wrapper.findAll('.picker-option')
    expect(options).toHaveLength(3)

    await options[2]!.trigger('click')

    expect(mockSetDayType).toHaveBeenCalledWith(nextMonthDay, SUNDAY_URI)
    expect(wrapper.find('.picker').exists()).toBe(false)
  })

  it('marks the active day type in the picker', async () => {
    mockEntries.value = new Map([[nextMonthDay, SATURDAY_URI]])

    const wrapper = mount(DayTypeCalendar)
    await flushPromises()
    await wrapper.find('[aria-label="Next month"]').trigger('click')
    await wrapper.find(`.day-cell[aria-label^="${nextMonthDay}"]`).trigger('click')

    const active = wrapper.findAll('.picker-option.active')
    expect(active).toHaveLength(1)
    expect(active[0]!.text()).toContain('Saturday')
  })

  it('shows a save status indicator for pending days', async () => {
    mockEntries.value = new Map([[lastVisibleDay, SATURDAY_URI]])
    mockStatuses.value = new Map([[lastVisibleDay, 'saving']])

    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    expect(wrapper.find(`.day-cell[aria-label^="${lastVisibleDay}"]`).classes()).toContain('saving')
  })

  it('shows the error banner and dismisses it', async () => {
    mockLastError.value = 'Failed to save the day type.'

    const wrapper = mount(DayTypeCalendar)
    await flushPromises()

    expect(wrapper.find('.calendar-error').text()).toContain('Failed to save the day type.')

    await wrapper.find('.error-dismiss').trigger('click')
    expect(mockClearError).toHaveBeenCalledTimes(1)
  })
})
