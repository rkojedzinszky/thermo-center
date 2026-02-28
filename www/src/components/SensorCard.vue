<script setup lang="ts">
import { ref, computed } from 'vue'
import type { THSensor } from '@/api'
import { useTimerSync, fmt, formatAge, checkInactive } from '@/composables/useSensorFormatting'
import { useResync } from '@/composables/useResync'

const props = defineProps<{
  sensor: THSensor
  index: number
  total: number
}>()

const emit = defineEmits<{
  dragStart: [index: number]
  dragOver: [index: number]
  dragEnd: []
  updated: [sensor: THSensor]
  deleted: [sensorId: number]
}>()

const flipped = ref(false)
const { now } = useTimerSync()
const sensorRef = computed(() => props.sensor)
const { resyncDisabled, handleResync } = useResync(sensorRef)

const isInactive = computed(() => checkInactive(props.sensor, now.value))
const isInvalid = computed(() => props.sensor.valid === false)

const ageLabel = computed(() => formatAge(props.sensor.lastTsf))

const temperature = computed(() =>
  props.sensor.temperature != null ? `${fmt(props.sensor.temperature)} °C` : '—',
)
const humidity = computed(() =>
  props.sensor.humidity != null ? `${fmt(props.sensor.humidity)} %` : '—',
)

// Back-side fields: all excluding valid and sensorResync
const backFields = computed(() => {
  const s = props.sensor
  return [
    { label: 'ID', value: String(s.id) },
    { label: 'Name', value: s.name },
    { label: 'Temperature', value: temperature.value },
    { label: 'Humidity', value: humidity.value },
    { label: 'Last Data', value: ageLabel.value },
    { label: 'VCC', value: s.vcc != null ? `${fmt(s.vcc)} V` : '—' },
    { label: 'RSSI', value: s.rssi != null ? `${fmt(s.rssi)} dBm` : '—' },
    { label: 'LQI', value: s.lqi != null ? fmt(s.lqi) : '—' },
    { label: 'Interval', value: s.interval != null ? `${fmt(s.interval)} s` : '—' },
    { label: 'Seq', value: s.lastSeq != null ? String(s.lastSeq) : '—' },
  ]
})

// ── Flip ────────────────────────────────────────────
let touchMoved = false

function toggleFlip(e: Event) {
  e.stopPropagation()
  flipped.value = !flipped.value
}

// ── Mouse drag ──────────────────────────────────────
function onDragStart(e: DragEvent) {
  // Custom ghost avoids the browser rendering the 3D-rotated back face
  const ghost = document.createElement('div')
  ghost.textContent = props.sensor.name
  ghost.setAttribute(
    'style',
    'position:fixed;top:-9999px;left:-9999px;padding:0.4rem 0.9rem;' +
      'background:#1e293b;color:#e2e8f0;border:1px solid rgba(255,255,255,0.15);' +
      'border-radius:0.5rem;font-size:0.875rem;font-family:inherit;' +
      'pointer-events:none;white-space:nowrap;',
  )
  document.body.appendChild(ghost)
  if (e.dataTransfer) {
    e.dataTransfer?.setDragImage?.(ghost, ghost.offsetWidth / 2, 16)
    e.dataTransfer.effectAllowed = 'move'
    e.dataTransfer.setData('text/plain', String(props.index))
  }
  setTimeout(() => document.body.removeChild(ghost), 0)
  emit('dragStart', props.index)
}

function onDragOver(e: DragEvent) {
  e.preventDefault()
  emit('dragOver', props.index)
}

// ── Touch drag (mobile) ─────────────────────────────
function onTouchStart(e: TouchEvent) {
  touchMoved = false
  emit('dragStart', props.index)
}

function onTouchMove(e: TouchEvent) {
  touchMoved = true
  e.preventDefault() // prevent page scroll while dragging
  const touch = e.touches[0]
  if (!touch) return
  const el = document.elementFromPoint(touch.clientX, touch.clientY)
  const card = el?.closest('[data-card-index]') as HTMLElement | null
  if (card) {
    const idx = Number(card.dataset.cardIndex)
    if (!isNaN(idx) && idx !== props.index) emit('dragOver', idx)
  }
}

function onTouchEnd(e: TouchEvent) {
  if (touchMoved) {
    // Prevent the synthetic click that would flip the card after a drag
    e.preventDefault()
  }
  touchMoved = false
  emit('dragEnd')
}

function onGripPointerDown(e: Event) {
  // Drag should start only from the grip and should not flip the card.
  e.stopPropagation()
}
</script>

<template>
  <div
    class="card-wrapper"
    :class="{ inactive: isInactive }"
    :aria-label="`Sensor ${sensor.name}`"
    :data-card-index="index"
    @dragover="onDragOver"
    @click="toggleFlip"
  >
    <div class="card" :class="{ flipped }">
      <!-- FRONT -->
      <div class="card-face card-front">
        <div class="card-header">
          <span class="sensor-name">{{ sensor.name }}</span>
          <div class="header-right">
            <span
              class="card-grip"
              title="Drag to reorder"
              draggable="true"
              aria-label="Drag card"
              @dragstart="onDragStart"
              @dragend="$emit('dragEnd')"
              @touchstart.passive="onTouchStart"
              @touchmove="onTouchMove"
              @touchend="onTouchEnd"
              @mousedown="onGripPointerDown"
              @click="onGripPointerDown"
            >
              ⠿
            </span>
          </div>
        </div>
        <div class="card-readings">
          <div class="reading">
            <span class="reading-value">{{ temperature }}</span>
            <span class="reading-label">Temp</span>
          </div>
          <div class="reading-divider" />
          <div class="reading">
            <span class="reading-value">{{ humidity }}</span>
            <span class="reading-label">Humidity</span>
          </div>
        </div>
        <div class="card-footer">
          <button
            v-if="isInvalid"
            class="resync-button"
            :disabled="resyncDisabled"
            title="Request sensor resynchronization"
            @click="handleResync"
          >
            ⚠ Resync
          </button>
          <span v-else class="age-label">{{ ageLabel }}</span>
          <span class="flip-hint">Click to flip</span>
        </div>
      </div>

      <!-- BACK -->
      <div class="card-face card-back">
        <ul class="back-fields">
          <li v-for="field in backFields" :key="field.label" class="back-field">
            <span class="field-label">{{ field.label }}</span>
            <span class="field-value">{{ field.value }}</span>
          </li>
        </ul>
        <div class="card-footer back-footer">
          <span class="flip-hint">Click to flip back</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card-wrapper {
  width: 168px;
  height: 196px;
  perspective: 1000px;
  cursor: pointer;
  user-select: none;
  flex-shrink: 0;
  transition: opacity 0.4s;
  touch-action: auto;
}

.card-wrapper.inactive {
  opacity: 0.45;
}

.card-wrapper:hover {
  z-index: 10;
}

.card-wrapper.inactive:hover {
  opacity: 0.65;
}

.card {
  width: 100%;
  height: 100%;
  position: relative;
  transform-style: preserve-3d;
  transition: transform 0.55s cubic-bezier(0.4, 0.2, 0.2, 1);
}

.card.flipped {
  transform: rotateY(180deg);
}

.card-face {
  position: absolute;
  inset: 0;
  backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
  border-radius: 1.1rem;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* ── FRONT ── */
.card-front {
  background: linear-gradient(145deg, var(--color-card-a) 0%, var(--color-card-b) 100%);
  border: 1px solid var(--color-border-card);
  box-shadow: var(--shadow-card);
}

.card-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 0.55rem 0.6rem 0.2rem;
  gap: 0.3rem;
}

.header-right {
  display: flex;
  align-items: flex-start;
  gap: 0.35rem;
}

.card-grip {
  color: var(--color-handle);
  cursor: grab;
  font-size: 0.88rem;
  line-height: 1;
  user-select: none;
  flex-shrink: 0;
  touch-action: none;
}

.card-edit-btn {
  background: transparent;
  border: none;
  color: var(--color-text-muted);
  font-size: 0.85rem;
  cursor: pointer;
  padding: 0.2rem 0.3rem;
  border-radius: 0.3rem;
  transition: all 0.2s;
  line-height: 1;
  user-select: none;
  flex-shrink: 0;
}

.card-edit-btn:hover {
  color: var(--color-text);
  background: rgba(99, 102, 241, 0.15);
}

.sensor-name {
  color: var(--color-text);
  font-size: 0.9rem;
  font-weight: 700;
  letter-spacing: -0.01em;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  min-width: 0;
}

.card-readings {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  flex: 1;
  padding: 0.12rem 0.36rem;
  justify-content: center;
  gap: 0.08rem;
}

.reading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 0 0 auto;
  gap: 0.08rem;
  min-width: 0;
  padding: 0.08rem;
}

.reading-value {
  color: var(--color-text-values);
  font-size: 1.22rem;
  font-weight: 700;
  letter-spacing: -0.02em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
}

.reading-label {
  color: var(--color-text-muted);
  font-size: 0.62rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.06em;
}

.reading-divider {
  width: 72%;
  height: 1px;
  margin: 0 auto;
  background: var(--color-reading-divider);
  flex-shrink: 0;
}

.card-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.26rem 0.56rem;
  border-top: 1px solid var(--color-footer-border);
  gap: 0.5rem;
}

.age-label {
  color: var(--color-text-secondary);
  font-size: 0.67rem;
  font-weight: 500;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
}

.resync-button {
  padding: 0.15rem 0.35rem;
  flex: 1;
}

.flip-hint {
  color: var(--color-text-faint);
  font-size: 0.58rem;
  font-style: italic;
  white-space: nowrap;
  flex-shrink: 0;
}

/* ── BACK ── */
.card-back {
  background: linear-gradient(145deg, var(--color-card-back-a) 0%, var(--color-card-back-b) 100%);
  border: 1px solid var(--color-border-card-back);
  box-shadow: var(--shadow-card-back);
  transform: rotateY(180deg);
}

.back-fields {
  list-style: none;
  margin: 0;
  padding: 0.3rem 0.56rem 0.14rem;
  flex: 1;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  gap: 0.04rem;
}

.back-field {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 0.24rem;
  padding: 0.02rem 0;
  font-weight: 500;
}

.field-label {
  color: var(--color-text-muted);
  font-size: 0.66rem;
  letter-spacing: 0.02em;
  white-space: nowrap;
  flex-shrink: 0;
}

.field-value {
  color: var(--color-text-fields);
  font-size: 0.66rem;
  text-align: right;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  min-width: 0;
}

.back-footer {
  border-top: 1px solid var(--color-back-footer-border);
  padding: 0.24rem 0.56rem;
}
</style>
