<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useControls } from '@/composables/useControls'
import { useInstantProfiles } from '@/composables/useInstantProfiles'
import { subscribeToWebSocket } from '@/services/websocket'
import ControlCard from '@/components/ControlCard.vue'
import InstantProfileButton from '@/components/InstantProfileButton.vue'

const { orderedControls, loadControls, updateControl } = useControls()
const { instantProfiles, loadInstantProfiles, toggleProfile } = useInstantProfiles()

let unsubscribeFromWebSocket: (() => void) | null = null

onMounted(async () => {
  // Load instant profiles
  await loadInstantProfiles()

  // Subscribe to websocket events
  unsubscribeFromWebSocket = subscribeToWebSocket({
    onConnected: () => {
      loadControls()
    },
    onUpdate: (sensorId: number) => {
      updateControl(sensorId)
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
      <!-- Instant Profiles Section -->
      <section v-if="instantProfiles.length > 0" class="profiles-section">
        <div class="profiles-container">
          <InstantProfileButton
            v-for="profile in instantProfiles"
            :key="profile.id"
            :profile="profile"
            @toggle="toggleProfile(profile.id)"
          />
        </div>
      </section>

      <!-- Controls Section -->
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

/* Profiles Section */
.profiles-section {
  margin-bottom: 2rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.section-title {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--color-text);
  margin: 0;
  padding: 0 0.5rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.profiles-container {
  display: flex;
  flex-wrap: wrap;
  gap: 0.8rem;
  align-items: center;
  justify-content: center;
}

@media (max-width: 768px) {
  .profiles-section {
    margin-bottom: 1rem;
    gap: 0.6rem;
  }

  .profiles-container {
    gap: 0.6rem;
  }

  .cards-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
    padding: 1rem;
  }
}
</style>
