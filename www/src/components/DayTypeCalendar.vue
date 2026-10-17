<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import type { DayType } from '@/api'
import EmptyState from '@/components/EmptyState.vue'
import { useCalendar, type CalendarSaveStatus } from '@/composables/useCalendar'
import { useDayTypes } from '@/composables/useDayTypes'
import { monthGrid, monthLabel, shiftMonth, todayIsoDay, toIsoDay } from '@/utils/calendarDays'
import { dayTypeColor } from '@/utils/dayTypeColors'

defineOptions({
  name: 'DayTypeCalendar',
})

const WEEKDAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']

interface DayCell {
  iso: string
  date: Date
  inMonth: boolean
  isToday: boolean
  isPast: boolean
  hasEntry: boolean
  dayType: DayType | null
  status: CalendarSaveStatus | null
}

const { dayTypes, loadDayTypes } = useDayTypes()
const { isLoading, lastError, dayTypeUri, statusFor, loadRange, setDayType, clearError } =
  useCalendar()

const today = todayIsoDay()
const initial = new Date()
const viewYear = ref(initial.getFullYear())
const viewMonth = ref(initial.getMonth())
const openDay = ref<string | null>(null)
const dayTypesLoaded = ref(false)

const grid = computed(() => monthGrid(viewYear.value, viewMonth.value))
const header = computed(() => monthLabel(viewYear.value, viewMonth.value))

const dayTypeByUri = computed(() => {
  return new Map(dayTypes.value.map((dayType) => [dayType.resourceUri, dayType]))
})
const dayTypeIndexOf = computed(() => {
  return new Map(dayTypes.value.map((dayType, index) => [dayType.resourceUri, index]))
})

const cells = computed<DayCell[]>(() => {
  return grid.value.days.map((date) => {
    const iso = toIsoDay(date)
    const uri = dayTypeUri(iso)
    return {
      iso,
      date,
      inMonth: date.getUTCMonth() === viewMonth.value && date.getUTCFullYear() === viewYear.value,
      isToday: iso === today,
      isPast: iso < today,
      hasEntry: uri !== null,
      dayType: uri ? (dayTypeByUri.value.get(uri) ?? null) : null,
      status: statusFor(iso),
    }
  })
})

function colorFor(dayType: DayType | null) {
  if (!dayType) return dayTypeColor(-1)
  return dayTypeColor(dayTypeIndexOf.value.get(dayType.resourceUri) ?? 0)
}

function cellAriaLabel(cell: DayCell): string {
  const base = cell.dayType ? `${cell.iso}, ${cell.dayType.name}` : `${cell.iso}, no day type`
  return cell.isPast ? `${base} (past)` : base
}

function openPicker(iso: string) {
  openDay.value = openDay.value === iso ? null : iso
}

function selectDayType(iso: string, dayType: DayType) {
  openDay.value = null
  void setDayType(iso, dayType.resourceUri)
}

async function showMonth(year: number, month: number): Promise<void> {
  viewYear.value = year
  viewMonth.value = month
  openDay.value = null
  const { start, end } = monthGrid(year, month)
  await loadRange(start, end)
}

function shiftView(delta: number) {
  const next = shiftMonth(viewYear.value, viewMonth.value, delta)
  void showMonth(next.year, next.month)
}

function goToday() {
  const current = new Date()
  void showMonth(current.getFullYear(), current.getMonth())
}

onMounted(async () => {
  await loadDayTypes()
  dayTypesLoaded.value = true
  const { start, end } = grid.value
  await loadRange(start, end)
})
</script>

<template>
  <EmptyState
    v-if="dayTypesLoaded && dayTypes.length === 0"
    message="No day types configured."
    min-height="300px"
  />

  <section v-else class="daytype-calendar">
    <header class="calendar-toolbar">
      <div class="month-nav">
        <button type="button" class="nav-button" aria-label="Previous month" @click="shiftView(-1)">
          ‹
        </button>
        <span class="month-label" aria-live="polite">{{ header }}</span>
        <button type="button" class="nav-button" aria-label="Next month" @click="shiftView(1)">
          ›
        </button>
        <button type="button" class="today-button" @click="goToday">Today</button>
      </div>

      <div class="legend" aria-label="Day types">
        <span v-for="(dayType, index) in dayTypes" :key="dayType.id" class="legend-item">
          <span class="swatch" :style="{ background: dayTypeColor(index).accent }"></span>
          {{ dayType.name }}
        </span>
      </div>
    </header>

    <p v-if="lastError" class="calendar-error" role="alert">
      <span>{{ lastError }}</span>
      <button type="button" class="error-dismiss" @click="clearError">Dismiss</button>
    </p>

    <div class="weekday-row">
      <span v-for="label in WEEKDAYS" :key="label" class="weekday">{{ label }}</span>
    </div>

    <div class="calendar-grid" :class="{ 'is-loading': isLoading }">
      <div v-for="cell in cells" :key="cell.iso" class="cell-wrap">
        <button
          type="button"
          class="day-cell"
          :class="{
            outside: !cell.inMonth,
            today: cell.isToday,
            past: cell.isPast,
            'no-entry': !cell.hasEntry,
            saving: cell.status === 'saving',
            error: cell.status === 'error',
            open: openDay === cell.iso,
          }"
          :disabled="cell.isPast || !cell.hasEntry"
          :aria-label="cellAriaLabel(cell)"
          @click="openPicker(cell.iso)"
        >
          <span class="day-number">{{ cell.date.getUTCDate() }}</span>
          <span
            v-if="cell.dayType"
            class="daytype-chip"
            :style="{
              color: colorFor(cell.dayType).accent,
              background: colorFor(cell.dayType).soft,
            }"
          >
            {{ cell.dayType.name }}
          </span>
          <span v-else class="daytype-none">—</span>
        </button>

        <div
          v-if="openDay === cell.iso"
          class="picker"
          role="dialog"
          :aria-label="`Day type for ${cell.iso}`"
        >
          <button
            v-for="dayType in dayTypes"
            :key="dayType.id"
            type="button"
            class="picker-option"
            :class="{ active: cell.dayType?.resourceUri === dayType.resourceUri }"
            @click="selectDayType(cell.iso, dayType)"
          >
            <span class="swatch" :style="{ background: colorFor(dayType).accent }"></span>
            <span class="picker-name">{{ dayType.name }}</span>
            <span v-if="cell.dayType?.resourceUri === dayType.resourceUri" class="picker-check"
              >✓</span
            >
          </button>
        </div>
      </div>
    </div>

    <div v-if="openDay" class="picker-backdrop" @click="openDay = null"></div>
  </section>
</template>

<style scoped>
.daytype-calendar {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  max-width: 900px;
  margin: 0 auto;
}

.calendar-toolbar {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  gap: 0.75rem;
}

.month-nav {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.nav-button {
  width: 2rem;
  height: 2rem;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--color-border);
  border-radius: 0.5rem;
  background: var(--color-card-a);
  color: var(--color-text);
  font-size: 1.1rem;
  line-height: 1;
  cursor: pointer;
  transition: border-color 0.2s;
}

.nav-button:hover {
  border-color: var(--color-accent-border);
}

.month-label {
  min-width: 10rem;
  text-align: center;
  font-size: 1rem;
  font-weight: 700;
  color: var(--color-text);
}

.today-button {
  border: 1px solid var(--color-border-card);
  border-radius: 0.5rem;
  background: var(--color-card-a);
  color: var(--color-text);
  font-size: 0.85rem;
  font-weight: 600;
  padding: 0.35rem 0.7rem;
  cursor: pointer;
  transition: border-color 0.2s;
}

.today-button:hover {
  border-color: var(--color-accent-border);
}

.legend {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.35rem;
  font-size: 0.8rem;
  color: var(--color-text-muted);
}

.swatch {
  width: 0.7rem;
  height: 0.7rem;
  border-radius: 50%;
  flex-shrink: 0;
}

.calendar-error {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  margin: 0;
  padding: 0.5rem 0.7rem;
  border: 1px solid var(--color-error-border);
  border-radius: 0.5rem;
  background: var(--color-error-bg);
  color: var(--color-error);
  font-size: 0.85rem;
}

.error-dismiss {
  margin-left: auto;
  border: 1px solid var(--color-error-border);
  border-radius: 0.35rem;
  background: transparent;
  color: var(--color-error);
  font-size: 0.75rem;
  font-weight: 600;
  padding: 0.15rem 0.5rem;
  cursor: pointer;
}

.weekday-row,
.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 0.35rem;
}

.weekday {
  text-align: center;
  font-size: 0.72rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--color-text-muted);
}

.calendar-grid {
  position: relative;
}

.calendar-grid.is-loading {
  opacity: 0.6;
}

.cell-wrap {
  position: relative;
  min-width: 0;
}

.day-cell {
  width: 100%;
  min-height: 3.8rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  gap: 0.25rem;
  padding: 0.35rem 0.25rem;
  border: 1px solid var(--color-border-card);
  border-radius: 0.55rem;
  background: var(--color-card-a);
  color: var(--color-text);
  cursor: pointer;
  transition:
    border-color 0.15s,
    background-color 0.15s;
}

.day-cell:hover:not(:disabled) {
  border-color: var(--color-accent-border);
}

.day-cell.outside {
  opacity: 0.55;
}

.day-cell.past,
.day-cell.no-entry {
  opacity: 0.45;
  cursor: default;
}

.day-cell.today {
  border-color: var(--color-accent);
  box-shadow: 0 0 0 1px var(--color-accent);
}

.day-cell.open {
  border-color: var(--color-accent-border);
  background: var(--color-accent-bg);
}

.day-cell.saving {
  border-style: dashed;
}

.day-cell.error {
  border-color: var(--color-error-border);
}

.day-number {
  font-size: 0.85rem;
  font-weight: 700;
}

.daytype-chip {
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 0.68rem;
  font-weight: 600;
  padding: 0.08rem 0.4rem;
  border-radius: 0.75rem;
}

.daytype-none {
  color: var(--color-text-faint);
  font-size: 0.75rem;
}

.picker {
  position: absolute;
  top: calc(100% + 0.25rem);
  left: 50%;
  transform: translateX(-50%);
  z-index: 50;
  min-width: 9rem;
  display: flex;
  flex-direction: column;
  gap: 0.15rem;
  padding: 0.3rem;
  border: 1px solid var(--color-border-dialog);
  border-radius: 0.6rem;
  background: var(--color-surface-solid);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.28);
}

.picker-option {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.4rem 0.5rem;
  border: none;
  border-radius: 0.4rem;
  background: transparent;
  color: var(--color-text);
  font-size: 0.85rem;
  text-align: left;
  cursor: pointer;
}

.picker-option:hover {
  background: var(--color-accent-bg);
}

.picker-option.active {
  background: var(--color-accent-bg);
  font-weight: 600;
}

.picker-name {
  flex: 1;
}

.picker-backdrop {
  position: fixed;
  inset: 0;
  z-index: 40;
}

@media (max-width: 768px) {
  .day-cell {
    min-height: 3.2rem;
    padding: 0.25rem 0.15rem;
  }

  .daytype-chip {
    font-size: 0.6rem;
    padding: 0.05rem 0.25rem;
  }

  .month-label {
    min-width: 7.5rem;
    font-size: 0.9rem;
  }
}
</style>
