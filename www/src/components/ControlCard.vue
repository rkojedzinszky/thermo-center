<script setup lang="ts">
import { computed } from 'vue'
import type { Control } from '@/api'
import { formatDecimal2 } from '@/composables/useSensorFormatting'

const props = defineProps<{
  control: Control
}>()

const current = computed(() =>
  props.control.temperature != null ? formatDecimal2(props.control.temperature) : '—',
)

const target = computed(() =>
  props.control.targetTemp != null ? formatDecimal2(props.control.targetTemp) : '—',
)

const difference = computed(() => {
  if (props.control.temperature == null || props.control.targetTemp == null) return null
  const diff = props.control.targetTemp - props.control.temperature
  return diff >= 0 ? `+${formatDecimal2(diff)}` : formatDecimal2(diff)
})
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
        <span class="temp-value">{{ target }} °C</span>
      </div>
    </div>

    <div v-if="difference !== null" class="card-diff">
      <span class="diff-label">Difference</span>
      <span
        class="diff-value"
        :class="{ positive: difference.startsWith('+'), negative: difference.startsWith('-') }"
      >
        {{ difference }} °C
      </span>
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

.control-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-card-hover, var(--shadow-card));
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.7rem 0.6rem;
  border-bottom: 1px solid var(--color-footer-border);
}

.control-name {
  color: var(--color-text);
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
  padding: 0.5rem 0.6rem;
  border-top: 1px solid var(--color-footer-border);
  background: rgba(0, 0, 0, 0.05);
}

.diff-label {
  color: var(--color-text-muted);
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.diff-value {
  color: var(--color-text);
  font-size: 0.95rem;
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
