<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { THSensor } from '@/api'
import { useTimerSync, fmt, formatAge, checkInactive } from '@/composables/useSensorFormatting'
import { useResyncMap } from '@/composables/useResync'
import { useAuth } from '@/composables/useAuth'
import SensorEditDialog from './SensorEditDialog.vue'

const props = defineProps<{
  sensors: THSensor[]
}>()

const emit = defineEmits<{
  reorder: [fromIndex: number, toIndex: number]
  updated: [sensor: THSensor]
  deleted: [sensorId: number]
}>()

const { now } = useTimerSync()
const { isAdmin } = useAuth()
const editDialogOpen = ref(false)
const selectedSensor = ref<THSensor | null>(null)
const sensorsRef = ref(props.sensors)
const { resyncDisabledMap, handleResync } = useResyncMap(sensorsRef)

// Keep sensorsRef in sync with prop changes
watch(
  () => props.sensors,
  (newSensors) => {
    sensorsRef.value = newSensors
  },
)

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

const columns = computed(() => [
  { key: 'name', label: 'Name' },
  { key: 'temperature', label: 'Temp (°C)' },
  { key: 'humidity', label: 'Humidity (%)' },
  { key: 'vcc', label: 'VCC (V)' },
  { key: 'rssi', label: 'RSSI (dBm)' },
  { key: 'lqi', label: 'LQI' },
  { key: 'interval', label: 'Interval (s)' },
  { key: 'lastSeq', label: 'Seq' },
  { key: 'lastTsf', label: 'Last Data' },
  ...(isAdmin.value ? [{ key: 'actions', label: 'Actions' }] : []),
])

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
          :data-row-index="idx"
          @dragover="onDragOver(idx, $event)"
          @drop="onDrop(idx, $event)"
        >
          <td class="td-drag">
            <span
              class="drag-handle"
              title="Drag to reorder"
              draggable="true"
              aria-label="Drag row"
              @dragstart="onDragStart(idx, $event)"
              @dragend="onDragEnd"
              @touchstart.passive="onRowTouchStart(idx, $event)"
            >
              ⠿
            </span>
          </td>
          <td v-for="col in columns" :key="col.key" class="td" :data-label="col.label">
            <button
              v-if="col.key === 'lastTsf' && sensor.valid === false"
              class="resync-button"
              :disabled="resyncDisabledMap.get(sensor.id) ?? false"
              title="Request sensor resynchronization"
              @click="handleResync(sensor, $event)"
            >
              ⚠ Resync
            </button>
            <button
              v-else-if="col.key === 'actions'"
              class="action-button"
              title="Edit sensor"
              @click="
                () => {
                  selectedSensor = sensor
                  editDialogOpen = true
                }
              "
            >
              Edit
            </button>
            <span v-else>
              {{ cellValue(sensor, col.key) }}
            </span>
          </td>
        </tr>
        <tr v-if="sensors.length === 0">
          <td :colspan="columns.length + 1" class="empty-row">No sensors found</td>
        </tr>
      </tbody>
    </table>
  </div>

  <SensorEditDialog
    :sensor="selectedSensor"
    :open="editDialogOpen"
    @close="editDialogOpen = false"
    @updated="
      (updatedSensor) => {
        emit('updated', updatedSensor)
        editDialogOpen = false
      }
    "
    @deleted="
      () => {
        if (selectedSensor) {
          emit('deleted', selectedSensor.id)
        }
        editDialogOpen = false
      }
    "
  />
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
  touch-action: none;
}

.resync-button {
  padding: 0.25rem 0.5rem;
}

.action-button {
  background: rgba(59, 130, 246, 0.2);
  border: 1px solid rgba(59, 130, 246, 0.6);
  color: #3b82f6;
  border-radius: 0.3rem;
  font-size: 0.75rem;
  font-weight: 600;
  cursor: pointer;
  padding: 0.35rem 0.6rem;
  transition: all 0.2s;
  white-space: nowrap;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
}

.action-button:hover:not(:disabled) {
  background: rgba(59, 130, 246, 0.35);
  border-color: rgba(59, 130, 246, 0.8);
}

.action-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
.empty-row {
  color: var(--color-text-muted);
  font-style: italic;
  padding: 2rem;
  text-align: center;
}
</style>
