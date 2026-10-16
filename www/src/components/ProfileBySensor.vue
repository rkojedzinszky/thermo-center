<script setup lang="ts">
import { computed, watch } from 'vue'
import EmptyState from '@/components/EmptyState.vue'
import ProfileModeHint from '@/components/ProfileModeHint.vue'
import ProfileSensorRow from '@/components/ProfileSensorRow.vue'
import ScheduleTabStrip from '@/components/ScheduleTabStrip.vue'
import { useControls } from '@/composables/useControls'
import { useControlTemps } from '@/composables/useControlTemps'
import { useDayTypes } from '@/composables/useDayTypes'
import { usePersistentId } from '@/composables/usePersistentRef'
import type { ProfileEditMode } from '@/composables/useProfileEditMode'
import { useProfiles } from '@/composables/useProfiles'
import { useScheduleReload } from '@/composables/useScheduleReload'

defineProps<{
  mode: ProfileEditMode
}>()

const { orderedControls, updateControl } = useControls()
const { dayTypes } = useDayTypes()
const { saveTick } = useProfiles()
const { reloadSensorView } = useScheduleReload()

const SENSOR_KEY = 'profile_sensor'
const { selectedId, select } = usePersistentId(SENSOR_KEY)

const selectedControl = computed(
  () =>
    orderedControls.value.find((control) => control.id === selectedId.value) ??
    orderedControls.value[0] ??
    null,
)

const { currentTemp, targetNow } = useControlTemps(selectedControl)

const tabs = computed(() =>
  orderedControls.value.map((control) => ({
    id: control.id,
    label: control.name,
  })),
)

function ensureSelection() {
  if (orderedControls.value.length === 0) return
  if (!orderedControls.value.some((control) => control.id === selectedId.value)) {
    const first = orderedControls.value[0]
    if (first) select(first.id)
  }
}

function selectAndReload(id: number) {
  // Clicking the active tab is not a change; skip the refetch.
  if (id === selectedId.value) return
  select(id)
  void reloadSensorView(id)
}

watch(orderedControls, ensureSelection, { immediate: true })

// Single device-target refresh for the selected control after any save.
watch(saveTick, () => {
  if (selectedControl.value) void updateControl(selectedControl.value.sensorId)
})
</script>

<template>
  <ScheduleTabStrip
    class="selector"
    :items="tabs"
    :selected-id="selectedControl?.id ?? null"
    group-label="Sensor"
    @select="selectAndReload"
  />

  <ProfileModeHint :mode="mode" />

  <EmptyState v-if="orderedControls.length === 0" message="No heating controls found." />

  <template v-else-if="selectedControl">
    <div class="sensor-summary">
      <span class="summary-name">{{ selectedControl.name }}</span>
      <span class="summary-meta"
        >Current <strong>{{ currentTemp }}</strong></span
      >
      <span class="summary-meta"
        >Target now <strong>{{ targetNow }}</strong></span
      >
    </div>

    <EmptyState v-if="dayTypes.length === 0" message="No day types are configured." />
    <div v-else class="row-list">
      <ProfileSensorRow
        v-for="dayType in dayTypes"
        :key="dayType.id"
        :control="selectedControl"
        :daytype="dayType"
        :title="dayType.name"
        :show-meta="false"
        :mode="mode"
      />
    </div>
  </template>
</template>

<style scoped>
.selector {
  margin-bottom: 0.25rem;
}

.sensor-summary {
  display: flex;
  align-items: baseline;
  gap: 1rem;
  flex-wrap: wrap;
  padding: 0.55rem 0.8rem;
  border: 1px solid var(--color-border);
  border-radius: 0.6rem;
  background: rgba(127, 127, 127, 0.06);
}

.summary-name {
  font-size: 0.95rem;
  font-weight: 700;
}

.summary-meta {
  font-size: 0.72rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--color-text-muted);
  font-weight: 600;
}

.summary-meta strong {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
  color: var(--color-text-values);
}

.row-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
</style>
