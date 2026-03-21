<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import type { THSensor } from '@/api'
import { useSensors } from '@/composables/useSensors'
import { useOverviewViewMode } from '@/composables/useOverviewViewMode'
import { subscribeToWebSocket } from '@/services/websocket'
import SensorCard from '@/components/SensorCard.vue'
import SensorTable from '@/components/SensorTable.vue'

const { orderedSensors, loadSensors, updateSensor, updateSensorDirect, removeSensor, reorder } =
  useSensors()
const { viewMode } = useOverviewViewMode()

// Card drag-and-drop state
const dragFromIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

function onCardDragStart(index: number) {
  dragFromIndex.value = index
}

function onCardDragOver(index: number) {
  dragOverIndex.value = index
}

function onCardDragEnd() {
  if (dragFromIndex.value !== null && dragOverIndex.value !== null) {
    reorder(dragFromIndex.value, dragOverIndex.value)
  }
  dragFromIndex.value = null
  dragOverIndex.value = null
}

// Handle sensor updates (instantly update local state)
function onSensorUpdated(updatedSensor: THSensor) {
  updateSensorDirect(updatedSensor)
}

// Handle sensor deletions
function onSensorDeleted(deletedId: number) {
  removeSensor(deletedId)
}

let unsubscribeFromWebSocket: (() => void) | null = null

onMounted(async () => {
  // Subscribe to websocket events
  unsubscribeFromWebSocket = subscribeToWebSocket({
    onConnected: () => {
      loadSensors()
    },
    onUpdate: (sensorId: number) => {
      updateSensor(sensorId)
    },
  })
})

onUnmounted(() => {
  unsubscribeFromWebSocket?.()
})
</script>

<template>
  <div class="overview">
    <!-- Main content -->
    <main class="overview-content">
      <!-- Cards view -->
      <div v-if="viewMode === 'cards'" class="cards-grid">
        <SensorCard
          v-for="(sensor, idx) in orderedSensors"
          :key="sensor.id"
          :sensor="sensor"
          :index="idx"
          :total="orderedSensors.length"
          :class="{ 'drag-target': dragOverIndex === idx && dragFromIndex !== idx }"
          @drag-start="onCardDragStart"
          @drag-over="onCardDragOver"
          @drag-end="onCardDragEnd"
          @updated="onSensorUpdated"
          @deleted="() => onSensorDeleted(sensor.id)"
        />
        <div v-if="orderedSensors.length === 0" class="empty-state">
          <span class="empty-icon">📡</span>
          <p>No sensors found</p>
        </div>
      </div>

      <!-- Table view -->
      <div v-else class="table-container">
        <SensorTable
          :sensors="orderedSensors"
          @reorder="reorder"
          @updated="onSensorUpdated"
          @deleted="(sensorId: number) => onSensorDeleted(sensorId)"
        />
      </div>
    </main>
  </div>
</template>

<style scoped>
.overview {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background: var(--color-bg);
  color: var(--color-text);
}

/* Content */
.overview-content {
  flex: 1;
  min-height: 0;
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
}

.cards-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
  align-items: flex-start;
  justify-content: center;
}

.cards-grid > :deep(.card-wrapper) {
  transition:
    transform 0.15s,
    box-shadow 0.15s;
}

.cards-grid > :deep(.card-wrapper.drag-target) {
  transform: scale(1.03);
  box-shadow: 0 0 0 2px var(--color-accent);
  border-radius: 1.1rem;
}

.table-container {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.empty-state {
  align-items: center;
  color: var(--color-text-muted);
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  padding: 4rem;
  width: 100%;
}

.empty-icon {
  font-size: 3rem;
}

.empty-state p {
  font-size: 1.1rem;
  margin: 0;
}
</style>
