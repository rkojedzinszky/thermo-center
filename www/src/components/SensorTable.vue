<script setup lang="ts">
import { ref } from 'vue'
import type { THSensor } from '@/api'
import { useTimerSync, fmt, formatAge, checkInactive } from '@/composables/useSensorFormatting'

const props = defineProps<{
  sensors: THSensor[]
}>()

const emit = defineEmits<{
  reorder: [fromIndex: number, toIndex: number]
}>()

const { now } = useTimerSync()

// ── Mouse drag ──────────────────────────────────────
const dragIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

function onDragStart(index: number, e: DragEvent) {
  dragIndex.value = index
  if (e.dataTransfer) {
    e.dataTransfer.effectAllowed = 'move'
    e.dataTransfer.setData('text/plain', String(index))
    // Custom ghost prevents the browser from rendering a garbled / white row
    const sensor = props.sensors[index]
    const ghost = document.createElement('div')
    ghost.textContent = sensor?.name ?? `Row ${index + 1}`
    ghost.setAttribute(
      'style',
      'position:fixed;top:-9999px;left:-9999px;padding:0.4rem 0.9rem;' +
        'background:#1e293b;color:#cbd5e1;border:1px solid rgba(255,255,255,0.1);' +
        'border-radius:0.4rem;font-size:0.875rem;font-family:inherit;' +
        'pointer-events:none;white-space:nowrap;',
    )
    document.body.appendChild(ghost)
    e.dataTransfer?.setDragImage?.(ghost, ghost.offsetWidth / 2, 16)
    setTimeout(() => document.body.removeChild(ghost), 0)
  }
}

function onDragOver(index: number, e: DragEvent) {
  e.preventDefault()
  dragOverIndex.value = index
}

function onDrop(toIndex: number, e: DragEvent) {
  e.preventDefault()
  if (dragIndex.value !== null && dragIndex.value !== toIndex) {
    emit('reorder', dragIndex.value, toIndex)
  }
  dragIndex.value = null
  dragOverIndex.value = null
}

function onDragEnd() {
  dragIndex.value = null
  dragOverIndex.value = null
}

// ── Touch drag (mobile) ─────────────────────────────
function onRowTouchStart(index: number, _e: TouchEvent) {
  dragIndex.value = index
}

function onTableTouchMove(e: TouchEvent) {
  if (dragIndex.value === null) return
  e.preventDefault() // prevent scroll while reordering
  const touch = e.touches[0]
  if (!touch) return
  const el = document.elementFromPoint(touch.clientX, touch.clientY)
  const row = el?.closest('[data-row-index]') as HTMLElement | null
  if (row) {
    const idx = Number(row.dataset.rowIndex)
    if (!isNaN(idx)) dragOverIndex.value = idx
  }
}

function onTableTouchEnd() {
  if (
    dragIndex.value !== null &&
    dragOverIndex.value !== null &&
    dragIndex.value !== dragOverIndex.value
  ) {
    emit('reorder', dragIndex.value, dragOverIndex.value)
  }
  dragIndex.value = null
  dragOverIndex.value = null
}

const columns = [
  { key: 'name', label: 'Name' },
  { key: 'temperature', label: 'Temp (°C)' },
  { key: 'humidity', label: 'Humidity (%)' },
  { key: 'vcc', label: 'VCC (V)' },
  { key: 'rssi', label: 'RSSI (dBm)' },
  { key: 'lqi', label: 'LQI' },
  { key: 'interval', label: 'Interval (s)' },
  { key: 'lastSeq', label: 'Seq' },
  { key: 'lastTsf', label: 'Last Data' },
]

function cellValue(sensor: THSensor, key: string): string {
  switch (key) {
    case 'name':
      return sensor.name
    case 'temperature':
      return sensor.temperature != null ? fmt(sensor.temperature) : '—'
    case 'humidity':
      return sensor.humidity != null ? fmt(sensor.humidity) : '—'
    case 'vcc':
      return sensor.vcc != null ? fmt(sensor.vcc) : '—'
    case 'rssi':
      return sensor.rssi != null ? fmt(sensor.rssi) : '—'
    case 'lqi':
      return sensor.lqi != null ? fmt(sensor.lqi) : '—'
    case 'interval':
      return sensor.interval != null ? fmt(sensor.interval) : '—'
    case 'lastSeq':
      return sensor.lastSeq != null ? String(sensor.lastSeq) : '—'
    case 'lastTsf':
      return formatAge(sensor.lastTsf)
    default:
      return '—'
  }
}
</script>

<template>
  <div class="table-wrapper" @touchmove="onTableTouchMove" @touchend="onTableTouchEnd">
    <table class="sensor-table">
      <thead>
        <tr>
          <th class="th-drag" />
          <th v-for="col in columns" :key="col.key" class="th">{{ col.label }}</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(sensor, idx) in sensors"
          :key="sensor.id"
          class="table-row"
          :class="{
            inactive: checkInactive(sensor, now),
            'drag-over': dragOverIndex === idx,
            dragging: dragIndex === idx,
          }"
          draggable="true"
          :data-row-index="idx"
          @dragstart="onDragStart(idx, $event)"
          @dragover="onDragOver(idx, $event)"
          @drop="onDrop(idx, $event)"
          @dragend="onDragEnd"
          @touchstart.passive="onRowTouchStart(idx, $event)"
        >
          <td class="td-drag">
            <span class="drag-handle" title="Drag to reorder">⠿</span>
          </td>
          <td v-for="col in columns" :key="col.key" class="td" :data-label="col.label">
            <span
              v-if="col.key === 'lastTsf'"
              :class="{ 'stale-text': checkInactive(sensor, now) }"
            >
              {{ cellValue(sensor, col.key) }}
            </span>
            <span v-else>{{ cellValue(sensor, col.key) }}</span>
          </td>
        </tr>
        <tr v-if="sensors.length === 0">
          <td :colspan="columns.length + 1" class="empty-row">No sensors found</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table-wrapper {
  width: 100%;
  overflow-x: auto;
  border-radius: 1rem;
  border: 1px solid var(--color-border);
  background: var(--color-table-bg);
  backdrop-filter: blur(8px);
}

.sensor-table {
  border-collapse: collapse;
  width: 100%;
  min-width: 700px;
  font-size: 0.875rem;
}

.th {
  background: var(--color-table-header-bg);
  border-bottom: 1px solid var(--color-border);
  color: var(--color-text-muted);
  font-size: 0.72rem;
  font-weight: 700;
  letter-spacing: 0.06em;
  padding: 0.85rem 1rem;
  text-align: left;
  text-transform: uppercase;
  white-space: nowrap;
}

.th-drag {
  background: var(--color-table-header-bg);
  border-bottom: 1px solid var(--color-border);
  width: 36px;
  min-width: 36px;
}

.table-row {
  border-bottom: 1px solid var(--color-border);
  cursor: grab;
  transition:
    background 0.15s,
    opacity 0.3s;
}

.table-row:last-child {
  border-bottom: none;
}

.table-row:hover {
  background: var(--color-table-row-hover);
}

.table-row.inactive {
  opacity: 0.4;
}

.table-row.inactive:hover {
  opacity: 0.6;
}

.table-row.drag-over {
  background: var(--color-drag-over-bg);
  border-top: 2px solid var(--color-accent-border);
}

.table-row.dragging {
  opacity: 0.25;
}

.td {
  color: var(--color-text-fields, var(--color-text));
  padding: 0.7rem 1rem;
  white-space: nowrap;
}

.td-drag {
  padding: 0.7rem 0.5rem;
  text-align: center;
  color: var(--color-handle);
}

.drag-handle {
  font-size: 1rem;
  cursor: grab;
  user-select: none;
}

.stale-text {
  color: var(--color-stale);
}

.empty-row {
  color: var(--color-text-muted);
  font-style: italic;
  padding: 2rem;
  text-align: center;
}
</style>
