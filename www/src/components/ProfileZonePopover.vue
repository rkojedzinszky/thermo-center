<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { Profile } from '@/api'
import type { ProfileEditMode } from '@/composables/useProfileEditMode'
import { useProfiles } from '@/composables/useProfiles'
import {
  DEFAULT_NEW_TEMP,
  MAX_TEMP,
  MIN_TEMP,
  SLOT_STEP_SECONDS,
  TEMP_STEP,
  clampTemp,
  formatClock,
  parseStart,
  validateStart,
} from '@/utils/profileSchedule'

const props = defineProps<{
  entry: Profile
  entries: Profile[]
  mode: ProfileEditMode
}>()

const emit = defineEmits<{
  close: []
}>()

const { setEntryTemp, setEntryDefault, setEntryStart, removeEntry, statusFor } = useProfiles()

const timeError = ref<string | null>(null)
const entryStartClock = computed(() => formatClock(parseStart(props.entry.start)))
const timeValue = ref(entryStartClock.value)

watch(entryStartClock, (next) => {
  // Keep the input in sync when the entry moves (e.g. after a drag),
  // unless the user is mid-edit with an error shown.
  if (timeError.value == null) timeValue.value = next
})

const isDefault = computed(() => props.entry.targetTemp == null)
const displayTemp = computed(() => props.entry.targetTemp ?? DEFAULT_NEW_TEMP)
const entryStatus = computed(() => statusFor(props.entry.id))
const statusTitle = computed(() =>
  entryStatus.value === 'error' ? 'Save failed' : entryStatus.value === 'saving' ? 'Saving…' : '',
)

function onTempInput(raw: unknown) {
  const value = Number(raw)
  if (Number.isNaN(value)) return
  setEntryTemp(props.entry.id, clampTemp(value))
}

function onSliderInput(event: Event) {
  onTempInput((event.target as HTMLInputElement).value)
}

function onNumberChange(event: Event) {
  onTempInput((event.target as HTMLInputElement).value)
}

function onDefaultToggle(event: Event) {
  const checked = (event.target as HTMLInputElement).checked
  if (checked) {
    setEntryDefault(props.entry.id)
  } else {
    setEntryTemp(props.entry.id, DEFAULT_NEW_TEMP)
  }
}

function onTimeChange(event: Event) {
  // Read from the event for test/compat (v-model updates on input, tests fire change).
  const raw = (event.target as HTMLInputElement).value
  timeValue.value = raw
  const seconds = parseStart(raw)
  const conflict = validateStart(props.entries, seconds, props.entry.id)
  if (conflict) {
    timeError.value = conflict
    // Reset via state (Vue re-renders the input); no direct DOM mutation.
    timeValue.value = entryStartClock.value
    return
  }
  timeError.value = null
  setEntryStart(props.entry.id, seconds)
}

function onDelete() {
  void removeEntry(props.entry.id)
  emit('close')
}
</script>

<template>
  <div
    class="zone-popover"
    role="dialog"
    :aria-label="`Zone at ${formatClock(parseStart(entry.start))}`"
  >
    <div class="popover-header">
      <span class="popover-title">Zone · {{ formatClock(parseStart(entry.start)) }}</span>
      <span class="popover-status" :class="entryStatus ?? ''" :title="statusTitle"></span>
      <button type="button" class="popover-done" @click="emit('close')">Done</button>
    </div>

    <div class="popover-row">
      <input
        class="popover-slider"
        type="range"
        :min="MIN_TEMP"
        :max="MAX_TEMP"
        :step="TEMP_STEP"
        :value="displayTemp"
        :disabled="isDefault"
        aria-label="Target temperature"
        @input="onSliderInput"
      />
      <input
        class="popover-number"
        type="number"
        :min="MIN_TEMP"
        :max="MAX_TEMP"
        :step="TEMP_STEP"
        :value="displayTemp"
        :disabled="isDefault"
        aria-label="Target temperature value"
        @change="onNumberChange"
      />
      <span class="popover-unit">°C</span>
    </div>

    <label class="popover-default">
      <input type="checkbox" :checked="isDefault" @change="onDefaultToggle" />
      Default / Off
    </label>

    <template v-if="mode === 'structure'">
      <div class="popover-row">
        <label class="popover-field">
          <span class="popover-field-label">Start</span>
          <input
            v-model="timeValue"
            class="popover-time"
            type="time"
            :step="SLOT_STEP_SECONDS"
            aria-label="Zone start time"
            @change="onTimeChange"
          />
        </label>
        <button type="button" class="popover-delete" @click="onDelete">Delete</button>
      </div>
      <p v-if="timeError" class="popover-error">{{ timeError }}</p>
    </template>
  </div>
</template>

<style scoped>
.zone-popover {
  width: 16rem;
  display: flex;
  flex-direction: column;
  gap: 0.55rem;
  padding: 0.65rem 0.75rem;
  border-radius: 0.6rem;
  border: 1px solid var(--color-border-dialog, var(--color-border));
  background: var(--color-surface-solid, var(--color-table-bg));
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.28);
}

.popover-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.popover-title {
  font-size: 0.85rem;
  font-weight: 700;
  color: var(--color-text);
}

.popover-status {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
  background: transparent;
}

.popover-status.saving {
  background: var(--color-accent);
}

.popover-status.error {
  background: var(--color-error, #ef4444);
}

.popover-done {
  margin-left: auto;
  background: var(--color-accent-bg);
  border: 1px solid var(--color-accent-border);
  color: var(--color-text);
  border-radius: 0.35rem;
  padding: 0.2rem 0.55rem;
  font-size: 0.75rem;
  font-weight: 600;
  cursor: pointer;
}

.popover-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.popover-slider {
  flex: 1;
  min-width: 0;
  accent-color: var(--color-accent);
}

.popover-number {
  width: 3.6rem;
  text-align: right;
  background: var(--color-input-bg, rgba(127, 127, 127, 0.08));
  border: 1px solid var(--color-border-input, var(--color-border));
  color: var(--color-input-text, var(--color-text));
  border-radius: 0.35rem;
  padding: 0.2rem 0.35rem;
  font-size: 0.85rem;
}

.popover-slider:disabled,
.popover-number:disabled {
  opacity: 0.4;
}

.popover-unit {
  font-size: 0.75rem;
  color: var(--color-text-muted);
}

.popover-default {
  display: flex;
  align-items: center;
  gap: 0.35rem;
  font-size: 0.78rem;
  color: var(--color-text-muted);
  cursor: pointer;
}

.popover-field {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  font-size: 0.78rem;
  color: var(--color-text-muted);
}

.popover-time {
  background: var(--color-input-bg, rgba(127, 127, 127, 0.08));
  border: 1px solid var(--color-border-input, var(--color-border));
  color: var(--color-input-text, var(--color-text));
  border-radius: 0.35rem;
  padding: 0.2rem 0.35rem;
  font-size: 0.82rem;
}

.popover-delete {
  margin-left: auto;
  background: var(--color-error-bg, rgba(239, 68, 68, 0.15));
  border: 1px solid var(--color-error-border, rgba(239, 68, 68, 0.4));
  color: var(--color-error, #ef4444);
  border-radius: 0.35rem;
  padding: 0.25rem 0.6rem;
  font-size: 0.78rem;
  font-weight: 600;
  cursor: pointer;
}

.popover-error {
  margin: 0;
  font-size: 0.75rem;
  color: var(--color-error);
}
</style>
