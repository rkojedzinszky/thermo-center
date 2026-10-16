<script setup lang="ts">
import { ref, computed } from 'vue'
import type { THSensor, Control } from '@/api'
import { formatDecimal2, formatInteger, formatAge } from '@/composables/useSensorFormatting'
import { useAuth } from '@/composables/useAuth'

type DataItem = THSensor | Control

defineOptions({
  name: 'DataTable',
})

const props = defineProps<{
  items: DataItem[]
  isExpanded?: boolean
  reorderMode?: boolean
  dataType: 'sensor' | 'control'
}>()

const emit = defineEmits<{
  reorder: [fromIndex: number, toIndex: number]
  updated: [item: DataItem]
  deleted: [itemId: number]
}>()

const { isAdmin } = useAuth()

// ── Mouse drag ──────────────────────────────────────
const dragIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

function onDragStart(index: number, e: DragEvent) {
  if (!props.reorderMode) return
  dragIndex.value = index
  if (e.dataTransfer) {
    e.dataTransfer.effectAllowed = 'move'
    e.dataTransfer.setData('text/plain', String(index))
    const item = props.items[index]!
    const name = isSensor(item) ? item.name : (item as Control).name
    const ghost = document.createElement('div')
    ghost.textContent = name ?? `Row ${index + 1}`
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

// ── Touch drag ──────────────────────────────────────
function onRowTouchStart(index: number) {
  if (!props.reorderMode) return
  dragIndex.value = index
}

function onTableTouchMove(e: TouchEvent) {
  if (dragIndex.value === null) return
  e.preventDefault()
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

// Type guards
function isSensor(item: DataItem): item is THSensor {
  return props.dataType === 'sensor'
}

function isControl(item: DataItem): item is Control {
  return props.dataType === 'control'
}

const columns = computed(() => {
  if (props.dataType === 'sensor') {
    const essentialCols = [
      { key: 'name', label: 'Name' },
      { key: 'temperature', label: 'Temp (°C)' },
      { key: 'humidity', label: 'Humidity (%)' },
    ]

    const allCols = [
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
    ]

    return props.isExpanded ? allCols : essentialCols
  } else {
    // Control columns
    const essentialCols = [
      { key: 'name', label: 'Name' },
      { key: 'temperature', label: 'Current (°C)' },
      { key: 'targetTemp', label: 'Target (°C)' },
      { key: 'difference', label: 'Diff (°C)' },
    ]

    const allCols = [
      { key: 'name', label: 'Name' },
      { key: 'temperature', label: 'Current (°C)' },
      { key: 'targetTemp', label: 'Target (°C)' },
      { key: 'difference', label: 'Difference (°C)' },
      { key: 'sensorId', label: 'Sensor ID' },
      ...(isAdmin.value ? [{ key: 'actions', label: 'Actions' }] : []),
    ]

    return props.isExpanded ? allCols : essentialCols
  }
})

function cellValue(item: DataItem, key: string): string {
  if (isSensor(item)) {
    const sensor = item as THSensor
    switch (key) {
      case 'name':
        return sensor.name
      case 'temperature':
        return sensor.temperature != null ? formatDecimal2(sensor.temperature) : '—'
      case 'humidity':
        return sensor.humidity != null ? formatDecimal2(sensor.humidity) : '—'
      case 'vcc':
        return sensor.vcc != null ? formatDecimal2(sensor.vcc) : '—'
      case 'rssi':
        return sensor.rssi != null ? formatInteger(sensor.rssi) : '—'
      case 'lqi':
        return sensor.lqi != null ? formatInteger(sensor.lqi) : '—'
      case 'interval':
        return sensor.interval != null ? formatDecimal2(sensor.interval) : '—'
      case 'lastSeq':
        return sensor.lastSeq != null ? formatInteger(sensor.lastSeq) : '—'
      case 'lastTsf':
        return formatAge(sensor.lastTsf)
      default:
        return '—'
    }
  } else {
    const control = item as Control
    switch (key) {
      case 'name':
        return control.name
      case 'temperature':
        return control.temperature != null ? formatDecimal2(control.temperature) : '—'
      case 'targetTemp':
        return control.targetTemp != null ? formatDecimal2(control.targetTemp) : '—'
      case 'difference': {
        if (control.temperature == null || control.targetTemp == null) return '—'
        const diff = control.targetTemp - control.temperature
        return diff >= 0 ? `+${formatDecimal2(diff)}` : formatDecimal2(diff)
      }
      case 'sensorId':
        return String(control.sensorId)
      default:
        return '—'
    }
  }
}
</script>

<template>
  <div
    class="table-wrapper"
    :class="{ 'reorder-active': reorderMode }"
    @touchmove="onTableTouchMove"
    @touchend="onTableTouchEnd"
  >
    <table class="data-table">
      <thead>
        <tr>
          <th
            v-for="col in columns"
            :key="col.key"
            class="th"
            :class="{
              numeric: [
                'temperature',
                'humidity',
                'vcc',
                'rssi',
                'lqi',
                'interval',
                'lastSeq',
                'lastTsf',
                'targetTemp',
                'difference',
                'sensorId',
              ].includes(col.key),
              actions: col.key === 'actions',
            }"
          >
            {{ col.label }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(item, idx) in items"
          :key="isSensor(item) ? (item as THSensor).id : (item as Control).id"
          class="tr"
          :class="{
            'drag-over': dragOverIndex === idx && dragIndex !== idx,
            'being-dragged': dragIndex === idx,
          }"
          :data-row-index="idx"
          :draggable="reorderMode ? 'true' : undefined"
          @dragstart="onDragStart(idx, $event)"
          @dragover="onDragOver(idx, $event)"
          @drop="onDrop(idx, $event)"
          @dragend="onDragEnd"
          @touchstart.passive="onRowTouchStart(idx)"
        >
          <td
            v-for="col in columns"
            :key="col.key"
            class="td"
            :class="{ numeric: col.key !== 'name' && col.key !== 'actions' }"
          >
            <template v-if="col.key === 'actions'">
              <div class="actions-cell">
                <!-- Actions can be added here if needed -->
              </div>
            </template>
            <template v-else>
              {{ cellValue(item, col.key) }}
            </template>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table-wrapper {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.table-wrapper.reorder-active {
  cursor: grab;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.875rem;
  flex: 1;
  overflow-y: auto;
}

.data-table thead {
  position: sticky;
  top: 0;
  background: var(--color-bg-table-header);
  z-index: 10;
}

.th {
  padding: 0.75rem 1rem;
  text-align: left;
  font-weight: 600;
  border-bottom: 1px solid var(--color-border-table);
  color: var(--color-text-secondary);
  font-size: 0.75rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.th.numeric {
  text-align: right;
}

.tr {
  border-bottom: 1px solid var(--color-border-table);
  transition:
    background-color 0.15s,
    transform 0.15s;
}

.tr:hover {
  background-color: var(--color-bg-table-hover);
}

.tr.drag-over {
  background-color: var(--color-accent-light);
  transform: scale(1.005);
  box-shadow: inset 0 2px 0 0 var(--color-accent);
}

.tr.being-dragged {
  opacity: 0.5;
}

.td {
  padding: 0.75rem 1rem;
  color: var(--color-text);
  font-family: 'Courier New', monospace;
  word-break: break-word;
}

.td.numeric {
  text-align: right;
}

.td:first-child {
  font-weight: 500;
  font-family: inherit;
}

.actions-cell {
  display: flex;
  gap: 0.5rem;
  justify-content: flex-end;
}
</style>
