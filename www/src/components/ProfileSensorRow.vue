<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import type { Control, DayType, Profile } from '@/api'
import ProfileTempChart from '@/components/ProfileTempChart.vue'
import ProfileZonePopover from '@/components/ProfileZonePopover.vue'
import { useControls } from '@/composables/useControls'
import { useControlTemps } from '@/composables/useControlTemps'
import { useProfiles } from '@/composables/useProfiles'
import { useScheduleReload } from '@/composables/useScheduleReload'
import type { ProfileEditMode } from '@/composables/useProfileEditMode'
import {
  DEFAULT_NEW_TEMP,
  SECONDS_PER_DAY,
  SLOT_STEP_SECONDS,
  SLOTS_PER_DAY,
  formatClock,
  parseStart,
  snapToStep,
  validateStart,
} from '@/utils/profileSchedule'

const props = withDefaults(
  defineProps<{
    control: Control
    daytype: DayType
    mode: ProfileEditMode
    title?: string
    showMeta?: boolean
  }>(),
  {
    showMeta: true,
  },
)

const { entriesFor, addEntry, setEntryStart, setEntryTemp, saveTick } = useProfiles()
const { isReloading } = useScheduleReload()
const { updateControl } = useControls()
const { currentTemp, targetNow: effectiveTarget } = useControlTemps(computed(() => props.control))

const activeId = ref<number | null>(null)
const rowError = ref<string | null>(null)
const areaRef = ref<HTMLElement | null>(null)

// entriesFor already returns sorted entries.
const entries = computed<Profile[]>(() =>
  entriesFor(props.control.resourceUri, props.daytype.resourceUri),
)

const activeEntry = computed<Profile | null>(
  () => entries.value.find((entry) => entry.id === activeId.value) ?? null,
)

const anchorPercent = computed(() => {
  if (!activeEntry.value) return 50
  const ratio = (parseStart(activeEntry.value.start) / SECONDS_PER_DAY) * 100
  return Math.min(88, Math.max(12, ratio))
})

// Refresh the device's effective target after any successful schedule change.
// Only the row that actually displays the target does this; in the by-sensor
// orientation the parent refreshes once for the shared control.
watch(saveTick, () => {
  if (!props.showMeta) return
  void updateControl(props.control.sensorId)
})

// Close the popover when the active entry disappears (deleted or reloaded).
watch(activeEntry, (entry) => {
  if (!entry) activeId.value = null
})

function openPopover(id: number) {
  activeId.value = id
  rowError.value = null
}

function closePopover() {
  activeId.value = null
}

function onChangeStart(id: number, seconds: number) {
  const conflict = validateStart(entries.value, seconds, id)
  if (conflict) {
    rowError.value = conflict
    return
  }
  rowError.value = null
  setEntryStart(id, seconds)
}

function onChangeTemp(id: number, temp: number) {
  setEntryTemp(id, temp)
}

function nearestFreeSlot(seconds: number): number | null {
  if (validateStart(entries.value, seconds) == null) return seconds
  for (let step = 1; step <= SLOTS_PER_DAY; step++) {
    for (const direction of [1, -1] as const) {
      const candidate =
        (seconds + direction * step * SLOT_STEP_SECONDS + SECONDS_PER_DAY) % SECONDS_PER_DAY
      if (validateStart(entries.value, candidate) == null) return candidate
    }
  }
  return null
}

async function addAt(seconds: number) {
  const slot = nearestFreeSlot(snapToStep(seconds))
  if (slot == null) {
    rowError.value = 'No free time slot available.'
    return
  }
  rowError.value = null
  const created = await addEntry(
    props.control.resourceUri,
    props.daytype.resourceUri,
    slot,
    DEFAULT_NEW_TEMP,
  )
  if (created) activeId.value = created.id
}

function onAddClick() {
  const list = entries.value
  const last = list[list.length - 1]
  const candidate = last ? (parseStart(last.start) + 2 * 3600) % SECONDS_PER_DAY : 6 * 3600
  void addAt(candidate)
}

function onDocumentPointerDown(event: PointerEvent) {
  if (activeId.value == null) return
  const target = event.target as Node | null
  if (areaRef.value && target && areaRef.value.contains(target)) return
  closePopover()
}

function onKeydown(event: KeyboardEvent) {
  if (event.key === 'Escape') closePopover()
}

onMounted(() => {
  document.addEventListener('pointerdown', onDocumentPointerDown)
  document.addEventListener('keydown', onKeydown)
})

onUnmounted(() => {
  document.removeEventListener('pointerdown', onDocumentPointerDown)
  document.removeEventListener('keydown', onKeydown)
})
</script>

<template>
  <div class="sensor-row">
    <div class="row-main" :class="`mode-${mode}`">
      <div class="row-label">
        <span class="label-heading">
          <span class="sensor-name">{{ title ?? control.name }}</span>
          <span class="zone-count">· {{ entries.length }}</span>
        </span>
        <span v-if="showMeta" class="sensor-meta">
          <span class="meta-value">{{ currentTemp }}</span>
          <span class="meta-arrow" aria-hidden="true">→</span>
          <span class="meta-value">{{ effectiveTarget }}</span>
        </span>
      </div>

      <div ref="areaRef" class="timeline-area">
        <ProfileTempChart
          :entries="entries"
          :mode="mode"
          :selected-id="activeId"
          @select="openPopover"
          @change-start="onChangeStart"
          @change-temp="onChangeTemp"
          @add-at="addAt"
        />
        <p v-if="isReloading && entries.length === 0" class="row-loading">Loading…</p>

        <div v-if="activeEntry" class="zone-popover-anchor" :style="{ left: `${anchorPercent}%` }">
          <ProfileZonePopover
            :entry="activeEntry"
            :entries="entries"
            :mode="mode"
            @close="closePopover"
          />
        </div>
      </div>

      <button
        v-if="mode === 'structure'"
        type="button"
        class="add-button"
        aria-label="Add zone"
        title="Add zone"
        @click="onAddClick"
      >
        +
      </button>
    </div>

    <p v-if="rowError" class="row-error">{{ rowError }}</p>
  </div>
</template>

<style scoped>
.sensor-row {
  border: 1px solid var(--color-border-card);
  border-radius: 0.75rem;
  background: linear-gradient(145deg, var(--color-card-a) 0%, var(--color-card-b) 100%);
  box-shadow: var(--shadow-card);
  padding: 0.5rem 0.75rem;
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
}

.row-main {
  display: grid;
  /* Fixed label track + remainder for the bar, so bars align across rows.
     The action track only exists in structure mode, letting the bar run to
     the row's right edge in temperature mode. */
  grid-template-columns: 11rem minmax(0, 1fr) 1.7rem;
  align-items: center;
  gap: 0.75rem;
}

.row-main.mode-temperature {
  grid-template-columns: 11rem minmax(0, 1fr);
}

@media (max-width: 700px) {
  .row-main {
    grid-template-columns: 7rem minmax(0, 1fr) 1.7rem;
    gap: 0.5rem;
  }

  .row-main.mode-temperature {
    grid-template-columns: 7rem minmax(0, 1fr);
  }
}

.row-label {
  grid-column: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 0.1rem;
}

.label-heading {
  display: flex;
  align-items: baseline;
  gap: 0.35rem;
  min-width: 0;
}

.sensor-name {
  font-size: 0.85rem;
  font-weight: 700;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.zone-count {
  font-size: 0.72rem;
  color: var(--color-text-muted);
  white-space: nowrap;
}

.timeline-area {
  position: relative;
  grid-column: 2;
  min-width: 0;
}

.row-loading {
  position: absolute;
  top: 50%;
  left: 0.5rem;
  transform: translateY(-50%);
  margin: 0;
  font-size: 0.75rem;
  color: var(--color-text-muted);
  pointer-events: none;
}

.sensor-meta {
  min-width: 0;
  overflow: hidden;
  display: flex;
  align-items: baseline;
  gap: 0.3rem;
  font-size: 0.72rem;
  white-space: nowrap;
}

.meta-arrow {
  color: var(--color-text-muted);
}

.meta-value {
  font-weight: 700;
  font-family: 'Courier New', monospace;
  color: var(--color-text-values);
}

.zone-popover-anchor {
  position: absolute;
  top: 100%;
  transform: translateX(-50%);
  margin-top: 0.4rem;
  z-index: 20;
}

.row-error {
  margin: 0;
  color: var(--color-error);
  font-size: 0.78rem;
}

.add-button {
  grid-column: 3;
  background: var(--color-accent-bg);
  border: 1px solid var(--color-accent-border);
  color: var(--color-text);
  border-radius: 0.45rem;
  width: 1.7rem;
  height: 1.7rem;
  line-height: 1;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  padding: 0;
}

.add-button:hover {
  background: var(--color-accent-border);
}
</style>
