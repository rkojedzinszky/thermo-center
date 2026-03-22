<script setup lang="ts">
import { computed, ref } from 'vue'
import type { Control, ScheduledOverrideW } from '@/api'
import api from '@/utils/api'
import { formatDecimal2 } from '@/composables/useSensorFormatting'
import { useControls } from '@/composables/useControls'

const props = defineProps<{
  control: Control
}>()

const { updateControl } = useControls()

const current = computed(() =>
  props.control.temperature != null ? formatDecimal2(props.control.temperature) : '—',
)

const isEditingTarget = ref(false)
const editedTarget = ref<number | null>(null)
const isSaving = ref(false)

const displayedTarget = computed(() => {
  if (isEditingTarget.value) return editedTarget.value
  return props.control.targetTemp ?? null
})

const target = computed(() =>
  displayedTarget.value != null ? Math.round(displayedTarget.value).toString() : '—',
)

const difference = computed(() => {
  if (props.control.temperature == null || props.control.targetTemp == null) return '—'
  const diff = props.control.targetTemp - props.control.temperature
  return diff >= 0 ? `+${formatDecimal2(diff)}` : formatDecimal2(diff)
})

function resolveEditBase(): number | null {
  if (isEditingTarget.value && editedTarget.value != null) return editedTarget.value
  if (props.control.targetTemp != null) return props.control.targetTemp
  if (props.control.temperature != null) return props.control.temperature
  return null
}

function stepTarget(delta: number) {
  const base = resolveEditBase()
  if (base == null) return
  editedTarget.value = Math.round(base + delta)
  isEditingTarget.value = true
}

function cancelTargetEdit() {
  isEditingTarget.value = false
  editedTarget.value = null
}

async function confirmTargetEdit() {
  if (editedTarget.value == null || isSaving.value) {
    cancelTargetEdit()
    return
  }

  const start = new Date()
  const end = new Date(start.getTime() + 2 * 60 * 60 * 1000)
  const payload: ScheduledOverrideW = {
    control: `/api/v1/control/${props.control.id}/`,
    start,
    end,
    targetTemp: editedTarget.value,
  }

  isSaving.value = true
  try {
    await api.createScheduledOverride({ scheduledOverrideW: payload })
    console.log('Reloading control after creating override, sensorId:', props.control.sensorId)
    await updateControl(props.control.sensorId)
  } catch (error) {
    console.error('Failed to create scheduled override:', error)
  } finally {
    isSaving.value = false
    cancelTargetEdit()
  }
}
</script>

<template>
  <div class="control-card">
    <div class="card-header">
      <span class="control-name">{{ control.name }}</span>
    </div>

    <div class="card-temps">
      <div class="temp-current">
        <span class="temp-label">Current</span>
        <span class="temp-value">{{ current }} °C</span>
      </div>

      <div class="temp-divider" />

      <div class="temp-target">
        <span class="temp-label">Target</span>
        <div class="target-editor">
          <button
            class="target-step-btn target-step-dec"
            :disabled="isSaving"
            :aria-label="`Decrease target for ${control.name}`"
            @click="stepTarget(-1)"
          >
            -
          </button>
          <span class="temp-value">{{ target }} °C</span>
          <button
            class="target-step-btn target-step-inc"
            :disabled="isSaving"
            :aria-label="`Increase target for ${control.name}`"
            @click="stepTarget(1)"
          >
            +
          </button>
        </div>
      </div>
    </div>

    <div class="card-diff">
      <template v-if="isEditingTarget">
        <button class="edit-action-btn cancel" :disabled="isSaving" @click="cancelTargetEdit">
          Cancel
        </button>
        <button class="edit-action-btn confirm" :disabled="isSaving" @click="confirmTargetEdit">
          {{ isSaving ? 'Saving...' : 'Confirm' }}
        </button>
      </template>
      <template v-else>
        <span class="diff-label">Difference</span>
        <span
          class="diff-value"
          :class="{ positive: difference.startsWith('+'), negative: difference.startsWith('-') }"
        >
          {{ difference }} °C
        </span>
      </template>
    </div>
  </div>
</template>

<style scoped>
.control-card {
  width: 180px;
  border-radius: 0.9rem;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  background: linear-gradient(145deg, var(--color-card-a) 0%, var(--color-card-b) 100%);
  border: 1px solid var(--color-border-card);
  box-shadow: var(--shadow-card);
  transition:
    transform 0.2s,
    box-shadow 0.2s;
  flex-shrink: 0;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.7rem 0.6rem;
  border-bottom: 1px solid var(--color-footer-border);
}

.control-name {
  font-size: 0.95rem;
  font-weight: 700;
  letter-spacing: -0.01em;
  text-align: center;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
  min-width: 0;
}

.card-temps {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0.8rem 0.6rem;
  gap: 0.4rem;
  flex: 1;
}

.temp-current,
.temp-target {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.15rem;
  width: 100%;
}

.target-editor {
  width: 100%;
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  gap: 0.4rem;
}

.temp-label {
  color: var(--color-text-muted);
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.temp-value {
  color: var(--color-text-values);
  font-size: 1.4rem;
  font-weight: 700;
  font-family: 'Courier New', monospace;
  text-align: center;
}

.target-step-btn {
  justify-self: center;
  width: auto;
  height: auto;
  border: none;
  border-radius: 0;
  background: transparent;
  font-size: 1.2rem;
  font-weight: 700;
  cursor: pointer;
  padding: 0.2rem 0.3rem;
  transition: color 0.2s;
}

.target-step-inc {
  color: rgba(255, 0, 0, 0.5);
}

.target-step-dec {
  color: rgba(0, 0, 255, 0.5);
}

.target-step-btn:hover:not(:disabled) {
  color: var(--color-accent-border, var(--color-accent));
  opacity: 0.8;
}

.target-step-btn:disabled {
  opacity: 0.35;
  cursor: not-allowed;
}

.temp-divider {
  width: 70%;
  height: 1px;
  background: var(--color-reading-divider);
  margin: 0.2rem 0;
}

.card-diff {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0rem;
  padding: 0.5rem;
  border-top: 1px solid var(--color-footer-border);
  background: rgba(0, 0, 0, 0.05);
  font-size: 0.7rem;
}

.edit-action-btn {
  flex: 1;
  border: none;
  border-radius: 0;
  background: transparent;
  color: var(--color-text);
  cursor: pointer;
  font-weight: 700;
  text-align: center;
  transition:
    background-color 0.2s,
    opacity 0.2s;
  padding: 0;
}

.edit-action-btn:hover:not(:disabled) {
  background-color: rgba(0, 0, 0, 0.08);
}

.edit-action-btn:active:not(:disabled) {
  background-color: rgba(0, 0, 0, 0.12);
}

.edit-action-btn.cancel:hover:not(:disabled) {
  background-color: rgba(0, 0, 0, 0.08);
}

.edit-action-btn.confirm:hover:not(:disabled) {
  background-color: rgba(255, 255, 255, 0.15);
  color: var(--color-accent);
}

.edit-action-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.diff-label {
  color: var(--color-text-muted);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.diff-value {
  font-weight: 700;
  font-family: 'Courier New', monospace;
}

.diff-value.positive {
  color: #ef4444;
}

.diff-value.negative {
  color: #10b981;
}
</style>
