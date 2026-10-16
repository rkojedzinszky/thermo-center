<script setup lang="ts">
import { onMounted, onUnmounted, watch } from 'vue'
import EmptyState from '@/components/EmptyState.vue'
import ProfileByDayType from '@/components/ProfileByDayType.vue'
import ProfileBySensor from '@/components/ProfileBySensor.vue'
import { useDayTypes } from '@/composables/useDayTypes'
import { useProfileAxis } from '@/composables/useProfileAxis'
import { useProfileEditMode } from '@/composables/useProfileEditMode'
import { useProfiles } from '@/composables/useProfiles'
import { useScheduleReload } from '@/composables/useScheduleReload'

const { dayTypes } = useDayTypes()
const { lastError, clearError, flushAll } = useProfiles()
const { axis, setAxis } = useProfileAxis()
const { mode, setMode } = useProfileEditMode()
const {
  isReloading,
  reloadInitialView,
  reloadSensorView,
  reloadDaytypeView,
  resolveSensorId,
  resolveDaytypeId,
} = useScheduleReload()

onMounted(() => {
  void reloadInitialView()
})

// Grouping toggle only needs the other axis's filtered slice: base lists
// were already loaded on activation.
watch(axis, (next) => {
  if (next === 'daytype') {
    const id = resolveDaytypeId()
    if (id != null) void reloadDaytypeView(id)
  } else {
    const id = resolveSensorId()
    if (id != null) void reloadSensorView(id)
  }
})

onUnmounted(() => {
  // Don't lose debounced temperature edits on navigation.
  void flushAll()
})
</script>

<template>
  <section class="schedule-editor" :aria-busy="isReloading">
    <p v-if="lastError" class="error-banner" role="alert">
      <span>{{ lastError }}</span>
      <button type="button" @click="clearError">Dismiss</button>
    </p>

    <div class="toolbar">
      <div class="switch" role="group" aria-label="Group by">
        <span class="switch-label">Group by</span>
        <button
          type="button"
          class="switch-button"
          :class="{ active: axis === 'sensor' }"
          :aria-pressed="axis === 'sensor'"
          @click="setAxis('sensor')"
        >
          Sensor
        </button>
        <button
          type="button"
          class="switch-button"
          :class="{ active: axis === 'daytype' }"
          :aria-pressed="axis === 'daytype'"
          @click="setAxis('daytype')"
        >
          Day type
        </button>
      </div>

      <div class="switch" role="group" aria-label="Edit mode">
        <button
          type="button"
          class="switch-button"
          :class="{ active: mode === 'temperature' }"
          :aria-pressed="mode === 'temperature'"
          @click="setMode('temperature')"
        >
          Temperatures
        </button>
        <button
          type="button"
          class="switch-button"
          :class="{ active: mode === 'structure' }"
          :aria-pressed="mode === 'structure'"
          @click="setMode('structure')"
        >
          Edit zones
        </button>
      </div>
    </div>

    <EmptyState v-if="isReloading && dayTypes.length === 0" message="Loading schedules…" />
    <EmptyState v-else-if="dayTypes.length === 0" message="No day types are configured." />
    <template v-else>
      <ProfileByDayType v-if="axis === 'daytype'" :mode="mode" />
      <ProfileBySensor v-else :mode="mode" />
    </template>
  </section>
</template>

<style scoped>
.schedule-editor {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-width: 900px;
  margin: 0 auto;
}

.error-banner {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin: 0;
  padding: 0.5rem 0.75rem;
  border: 1px solid var(--color-error-border);
  background: var(--color-error-bg);
  color: var(--color-error);
  border-radius: 0.5rem;
  font-size: 0.85rem;
}

.error-banner button {
  margin-left: auto;
  background: transparent;
  border: 1px solid var(--color-error-border);
  color: inherit;
  border-radius: 0.35rem;
  padding: 0.2rem 0.55rem;
  cursor: pointer;
}

.toolbar {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
}

.switch {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  background: rgba(127, 127, 127, 0.12);
  border: 1px solid var(--color-border);
  border-radius: 0.5rem;
  padding: 0.2rem;
}

.switch-label {
  font-size: 0.7rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--color-text-muted);
  padding: 0 0.4rem 0 0.5rem;
}

.switch-button {
  border: none;
  background: transparent;
  color: var(--color-text-muted);
  border-radius: 0.35rem;
  padding: 0.35rem 0.7rem;
  font-size: 0.8rem;
  font-weight: 600;
  cursor: pointer;
  transition:
    background-color 0.2s,
    color 0.2s;
}

.switch-button:hover {
  color: var(--color-text);
}

.switch-button.active {
  background: var(--color-accent-bg);
  color: var(--color-text);
}
</style>
