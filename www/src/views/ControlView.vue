<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useControls } from '@/composables/useControls'
import { subscribeToWebSocket } from '@/services/websocket'
import ControlCard from '@/components/ControlCard.vue'

const { orderedControls, loadControls, updateControl } = useControls()

let unsubscribeFromWebSocket: (() => void) | null = null

onMounted(async () => {
  // Subscribe to websocket events
  unsubscribeFromWebSocket = subscribeToWebSocket({
    onConnected: () => {
      loadControls()
    },
    onUpdate: (controlId: number) => {
      updateControl(controlId)
    },
  })
})

onUnmounted(() => {
  unsubscribeFromWebSocket?.()
})
</script>

<template>
  <div class="heat-control">
    <!-- Content -->
    <main class="content">
      <div v-if="orderedControls.length === 0" class="empty-state">
        <p>No heating controls found.</p>
      </div>
      <div v-else class="cards-grid">
        <ControlCard v-for="control in orderedControls" :key="control.id" :control="control" />
      </div>
    </main>
  </div>
</template>

<style scoped>
.heat-control {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background: var(--color-background);
}

.content {
  flex: 1;
  min-height: 0;
  padding: 2rem;
  overflow-y: auto;
}

.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 300px;
  color: var(--color-text-muted);
  font-size: 1.1rem;
}

.cards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 1.5rem;
  grid-auto-rows: max-content;
}

@media (max-width: 768px) {
  .cards-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
    padding: 1rem;
  }
}
</style>
