<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import type { Profile } from '@/api'
import type { ProfileEditMode } from '@/composables/useProfileEditMode'
import {
  CHART_MAX_TEMP,
  CHART_MIN_TEMP,
  SECONDS_PER_DAY,
  buildSegments,
  formatClock,
  parseStart,
  snapChartTemp,
  snapToStep,
  tempColor,
  tempFill,
  tempToRatio,
  type ScheduleSegment,
} from '@/utils/profileSchedule'

const props = defineProps<{
  entries: Profile[]
  mode: ProfileEditMode
  selectedId?: number | null
}>()

const emit = defineEmits<{
  select: [id: number]
  changeStart: [id: number, seconds: number]
  changeTemp: [id: number, temp: number]
  addAt: [seconds: number]
}>()

const GRID_TEMPS = [30, 25, 20, 15, 10]

const plotRef = ref<HTMLElement | null>(null)
const dragId = ref<number | null>(null)
const dragKind = ref<'start' | 'temp' | null>(null)
const dragSeconds = ref<number | null>(null)
const dragTemp = ref<number | null>(null)
const dragOriginSeconds = ref<number | null>(null)
const dragOriginTemp = ref<number | null>(null)
const didDrag = ref(false)
const nowSeconds = ref(currentSeconds())

function currentSeconds(): number {
  const now = new Date()
  return now.getHours() * 3600 + now.getMinutes() * 60 + now.getSeconds()
}

const segments = computed<ScheduleSegment[]>(() => buildSegments(props.entries))

function xPercent(seconds: number): string {
  return `${Math.min(100, Math.max(0, (seconds / SECONDS_PER_DAY) * 100))}%`
}

/** CSS top for a temperature inside the 10–30 °C plot (values clamp). */
function yTop(temp: number): string {
  return `${(1 - tempToRatio(temp)) * 100}%`
}

/** Live temperature of a segment while its line is dragged. */
function tempOf(segment: ScheduleSegment): number | null {
  if (segment.temp == null) return null
  if (dragKind.value === 'temp' && segment.entry?.id === dragId.value && dragTemp.value != null) {
    return dragTemp.value
  }
  return segment.temp
}

function bandStyle(segment: ScheduleSegment): Record<string, string> {
  return { left: xPercent(segment.start), width: xPercent(segment.end - segment.start) }
}

/** Filled bar rising from the plot bottom up to the segment's target level. */
function barStyle(segment: ScheduleSegment): Record<string, string> {
  const temp = tempOf(segment)
  const style: Record<string, string> = {
    left: xPercent(segment.start),
    width: xPercent(segment.end - segment.start),
  }
  if (temp != null) {
    style.top = yTop(temp)
    const color = tempColor(temp)
    const fill = tempFill(temp)
    // Custom properties consumed by .zone-bar; Vue applies them via setProperty.
    if (color) style['--bar-color'] = color
    if (fill) style['--bar-fill'] = fill
  }
  return style
}

/** Invisible vertical drag area hugging the bar's top edge. */
function topHandleStyle(segment: ScheduleSegment): Record<string, string> {
  const temp = tempOf(segment)
  const style: Record<string, string> = {
    left: xPercent(segment.start),
    width: xPercent(segment.end - segment.start),
  }
  if (temp != null) style.top = yTop(temp)
  return style
}

function startSeconds(entry: Profile): number {
  if (dragKind.value === 'start' && dragId.value === entry.id && dragSeconds.value != null) {
    return dragSeconds.value
  }
  return parseStart(entry.start)
}

function segmentAt(seconds: number): ScheduleSegment | null {
  for (const segment of segments.value) {
    if (seconds >= segment.start && seconds < segment.end) return segment
  }
  return segments.value[segments.value.length - 1] ?? null
}

function secondsFromClientX(clientX: number): number | null {
  const plot = plotRef.value
  if (!plot) return null
  const rect = plot.getBoundingClientRect()
  if (rect.width === 0) return null
  const ratio = (clientX - rect.left) / rect.width
  return snapToStep(Math.min(SECONDS_PER_DAY, Math.max(0, ratio * SECONDS_PER_DAY)))
}

function tempFromClientY(clientY: number): number | null {
  const plot = plotRef.value
  if (!plot) return null
  const rect = plot.getBoundingClientRect()
  if (rect.height === 0) return null
  const ratio = 1 - (clientY - rect.top) / rect.height
  return snapChartTemp(CHART_MIN_TEMP + ratio * (CHART_MAX_TEMP - CHART_MIN_TEMP))
}

function startDrag(kind: 'start' | 'temp', entry: Profile) {
  dragId.value = entry.id
  dragKind.value = kind
  dragOriginSeconds.value = parseStart(entry.start)
  dragSeconds.value = dragOriginSeconds.value
  dragOriginTemp.value = entry.targetTemp ?? null
  dragTemp.value = dragOriginTemp.value
  didDrag.value = false
  window.addEventListener('pointermove', onPointerMove)
  window.addEventListener('pointerup', onPointerUp)
  window.addEventListener('pointercancel', onPointerCancel)
}

function onStartHandlePointerDown(event: PointerEvent, entry: Profile) {
  event.preventDefault()
  event.stopPropagation()
  emit('select', entry.id)
  startDrag('start', entry)
}

function onTempLinePointerDown(event: PointerEvent, entry: Profile) {
  event.preventDefault()
  event.stopPropagation()
  emit('select', entry.id)
  if (props.mode !== 'temperature' || entry.targetTemp == null) return
  startDrag('temp', entry)
}

function onPointerMove(event: PointerEvent) {
  if (dragId.value == null || dragKind.value == null) return
  if (dragKind.value === 'start') {
    const next = secondsFromClientX(event.clientX)
    if (next == null) return
    if (next !== dragSeconds.value) {
      didDrag.value = true
      dragSeconds.value = next
    }
    return
  }
  const next = tempFromClientY(event.clientY)
  if (next == null) return
  if (next !== dragTemp.value) {
    didDrag.value = true
    dragTemp.value = next
  }
}

function onPointerUp() {
  // Only emit when the handle actually moved; a plain click is a select.
  if (dragId.value != null && didDrag.value) {
    if (
      dragKind.value === 'start' &&
      dragSeconds.value != null &&
      dragSeconds.value !== dragOriginSeconds.value
    ) {
      emit('changeStart', dragId.value, dragSeconds.value)
    } else if (
      dragKind.value === 'temp' &&
      dragTemp.value != null &&
      dragTemp.value !== dragOriginTemp.value
    ) {
      emit('changeTemp', dragId.value, dragTemp.value)
    }
  }
  teardownDrag()
}

function onPointerCancel() {
  teardownDrag()
}

function teardownDrag() {
  dragId.value = null
  dragKind.value = null
  dragSeconds.value = null
  dragTemp.value = null
  dragOriginSeconds.value = null
  dragOriginTemp.value = null
  window.removeEventListener('pointermove', onPointerMove)
  window.removeEventListener('pointerup', onPointerUp)
  window.removeEventListener('pointercancel', onPointerCancel)
}

function onPlotClick(event: MouseEvent) {
  const target = event.target as HTMLElement | null
  if (didDrag.value) {
    didDrag.value = false
    return
  }
  if (target?.closest('.start-handle, .temp-handle')) return

  const plot = plotRef.value
  if (!plot) return
  const rect = plot.getBoundingClientRect()
  if (rect.width === 0) return
  const seconds = Math.min(
    SECONDS_PER_DAY,
    Math.max(0, ((event.clientX - rect.left) / rect.width) * SECONDS_PER_DAY),
  )

  if (props.mode === 'structure') {
    const snapped = secondsFromClientX(event.clientX)
    if (snapped != null) emit('addAt', snapped)
    return
  }

  const segment = segmentAt(seconds)
  if (segment?.entry) emit('select', segment.entry.id)
}

function zoneTitle(segment: ScheduleSegment): string {
  const range = `${formatClock(segment.start)} – ${formatClock(segment.end)}`
  const temp = tempOf(segment)
  return temp == null ? `${range}, default` : `${range}, ${temp} °C`
}

function handleTitle(entry: Profile): string {
  const segment = segments.value.find((candidate) => candidate.entry?.id === entry.id)
  if (!segment) return formatClock(startSeconds(entry))
  return zoneTitle(segment)
}

function tempAriaLabel(entry: Profile): string {
  const time = formatClock(startSeconds(entry))
  if (entry.targetTemp != null) {
    const live =
      dragKind.value === 'temp' && dragId.value === entry.id && dragTemp.value != null
        ? dragTemp.value
        : entry.targetTemp
    return `Zone target ${live} degrees from ${time}`
  }
  return `Zone from ${time}, default`
}

let nowTimer: ReturnType<typeof setInterval> | null = null

onMounted(() => {
  nowTimer = setInterval(() => {
    nowSeconds.value = currentSeconds()
  }, 30000)
})

onUnmounted(() => {
  teardownDrag()
  if (nowTimer) clearInterval(nowTimer)
})
</script>

<template>
  <div class="temp-chart">
    <div
      ref="plotRef"
      class="chart-plot"
      :class="`mode-${mode}`"
      :title="
        mode === 'structure'
          ? 'Click to add a zone, drag a line to move it'
          : 'Drag zone starts or target lines'
      "
      @click="onPlotClick"
    >
      <div v-for="temp in GRID_TEMPS" :key="temp" class="gridline" :style="{ top: yTop(temp) }">
        <span class="grid-label">{{ temp }}°</span>
      </div>

      <div
        v-for="segment in segments.filter((s) => s.temp == null)"
        :key="`band-${segment.start}`"
        class="default-band"
        :style="bandStyle(segment)"
        :title="zoneTitle(segment)"
      />

      <div
        v-for="segment in segments.filter((s) => s.temp != null && s.entry != null)"
        :key="`bar-${segment.entry!.id}`"
        class="zone-bar"
        :class="{ selected: segment.entry!.id === selectedId }"
        :style="barStyle(segment)"
        :title="zoneTitle(segment)"
      />

      <button
        v-for="segment in segments.filter((s) => s.temp != null && s.entry != null)"
        :key="`top-${segment.entry!.id}`"
        type="button"
        class="temp-handle"
        :style="topHandleStyle(segment)"
        :aria-label="tempAriaLabel(segment.entry!)"
        :title="zoneTitle(segment)"
        @pointerdown="onTempLinePointerDown($event, segment.entry!)"
      />

      <button
        v-for="entry in entries"
        :key="`start-${entry.id}`"
        type="button"
        class="start-handle"
        :class="{
          active: entry.id === selectedId,
          dragging: entry.id === dragId && dragKind === 'start',
        }"
        :style="{ left: xPercent(startSeconds(entry)) }"
        :aria-label="`Zone start ${formatClock(startSeconds(entry))}`"
        :title="handleTitle(entry)"
        @pointerdown="onStartHandlePointerDown($event, entry)"
      >
        <span class="handle-stem" />
      </button>

      <div class="now-line" :style="{ left: xPercent(nowSeconds) }" title="Current time" />
    </div>

    <div class="timeline-scale">
      <span v-for="hour in [0, 6, 12, 18, 24]" :key="hour"
        >{{ String(hour).padStart(2, '0') }}:00</span
      >
    </div>
  </div>
</template>

<style scoped>
.temp-chart {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.chart-plot {
  position: relative;
  height: 4.75rem;
  border-radius: 0.4rem;
  background: var(--color-input-bg, rgba(127, 127, 127, 0.08));
  border: 1px solid var(--color-border);
  overflow: hidden;
  touch-action: pan-y;
  cursor: pointer;
}

.chart-plot.mode-structure {
  cursor: copy;
}

.gridline {
  position: absolute;
  left: 0;
  right: 0;
  border-top: 1px dashed rgba(127, 127, 127, 0.25);
  pointer-events: none;
}

.grid-label {
  position: absolute;
  top: 0;
  left: 0.25rem;
  transform: translateY(-100%);
  font-size: 0.6rem;
  color: var(--color-text-muted);
  pointer-events: none;
}

.default-band {
  position: absolute;
  top: 0;
  bottom: 0;
  background: repeating-linear-gradient(
    45deg,
    rgba(127, 127, 127, 0.18),
    rgba(127, 127, 127, 0.18) 6px,
    rgba(127, 127, 127, 0.06) 6px,
    rgba(127, 127, 127, 0.06) 12px
  );
}

.zone-bar {
  position: absolute;
  bottom: 0;
  border-top: 3px solid var(--bar-color, #6366f1);
  background: var(--bar-fill, rgba(99, 102, 241, 0.22));
}

.zone-bar.selected {
  border-top-width: 5px;
  outline: 2px solid var(--color-accent, #6366f1);
  outline-offset: -2px;
}

.temp-handle {
  position: absolute;
  height: 1.1rem;
  transform: translateY(-50%);
  background: transparent;
  border: none;
  padding: 0;
  cursor: ns-resize;
  touch-action: none;
}

.temp-handle::after {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  height: 5px;
  transform: translateY(-50%);
  border-radius: 3px;
  background: rgba(255, 255, 255, 0.4);
  opacity: 0;
  transition: opacity 0.15s;
  pointer-events: none;
}

.temp-handle:hover::after {
  opacity: 1;
}

.start-handle {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 1.25rem;
  transform: translateX(-50%);
  background: transparent;
  border: none;
  padding: 0;
  cursor: ew-resize;
  display: flex;
  justify-content: center;
  touch-action: none;
}

.handle-stem {
  width: 3px;
  height: 100%;
  border-radius: 2px;
  background: rgba(255, 255, 255, 0.75);
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.35);
  pointer-events: none;
  transition:
    width 0.15s,
    background-color 0.15s;
}

.start-handle:hover .handle-stem {
  width: 4px;
  background: rgba(255, 255, 255, 0.95);
  box-shadow:
    0 0 0 1px rgba(0, 0, 0, 0.4),
    0 0 0 4px rgba(127, 127, 127, 0.3);
}

.start-handle.active .handle-stem {
  background: var(--color-accent, #6366f1);
  box-shadow: 0 0 0 1px rgba(255, 255, 255, 0.8);
}

.now-line {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 1px;
  background: var(--color-accent, #6366f1);
  opacity: 0.8;
  pointer-events: none;
  z-index: 2;
}

.timeline-scale {
  display: flex;
  justify-content: space-between;
  font-size: 0.65rem;
  color: var(--color-text-muted);
}
</style>
