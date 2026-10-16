<script setup lang="ts">
import { computed, watch } from 'vue'
import EmptyState from '@/components/EmptyState.vue'
import ProfileModeHint from '@/components/ProfileModeHint.vue'
import ProfileSensorRow from '@/components/ProfileSensorRow.vue'
import ScheduleTabStrip from '@/components/ScheduleTabStrip.vue'
import { useControls } from '@/composables/useControls'
import { useDayTypes } from '@/composables/useDayTypes'
import { usePersistentId } from '@/composables/usePersistentRef'
import type { ProfileEditMode } from '@/composables/useProfileEditMode'
import { useScheduleReload } from '@/composables/useScheduleReload'

defineProps<{
  mode: ProfileEditMode
}>()

const { orderedControls } = useControls()
const { dayTypes } = useDayTypes()
const { reloadDaytypeView } = useScheduleReload()

const DAYTYPE_KEY = 'profile_daytype'
const { selectedId, select } = usePersistentId(DAYTYPE_KEY)

const selectedDaytype = computed(
  () =>
    dayTypes.value.find((dayType) => dayType.id === selectedId.value) ?? dayTypes.value[0] ?? null,
)

const tabs = computed(() =>
  dayTypes.value.map((dayType) => ({
    id: dayType.id,
    label: dayType.name,
  })),
)

function ensureSelection() {
  if (dayTypes.value.length === 0) return
  if (!dayTypes.value.some((dayType) => dayType.id === selectedId.value)) {
    const first = dayTypes.value[0]
    if (first) select(first.id)
  }
}

function selectAndReload(id: number) {
  // Clicking the active tab is not a change; skip the refetch.
  if (id === selectedId.value) return
  select(id)
  void reloadDaytypeView(id)
}

watch(dayTypes, ensureSelection, { immediate: true })
</script>

<template>
  <ScheduleTabStrip
    class="selector"
    :items="tabs"
    :selected-id="selectedDaytype?.id ?? null"
    group-label="Day type"
    @select="selectAndReload"
  />

  <ProfileModeHint :mode="mode" />

  <EmptyState v-if="orderedControls.length === 0" message="No heating controls found." />
  <div v-else-if="selectedDaytype" class="row-list">
    <ProfileSensorRow
      v-for="control in orderedControls"
      :key="control.id"
      :control="control"
      :daytype="selectedDaytype"
      :mode="mode"
    />
  </div>
</template>

<style scoped>
.selector {
  margin-bottom: 0.25rem;
}

.row-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
</style>
