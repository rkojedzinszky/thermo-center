<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useControls } from '@/composables/useControls'
import { useInstantProfiles } from '@/composables/useInstantProfiles'
import { usePersistentString } from '@/composables/usePersistentRef'
import { subscribeToWebSocket } from '@/services/websocket'
import ControlCard from '@/components/ControlCard.vue'
import EmptyState from '@/components/EmptyState.vue'
import InstantProfileButton from '@/components/InstantProfileButton.vue'
import ProfileScheduleEditor from '@/components/ProfileScheduleEditor.vue'

type HeatingTab = 'cards' | 'schedules'

const TAB_KEY = 'heating_tab'

function isHeatingTab(value: string): value is HeatingTab {
  return value === 'cards' || value === 'schedules'
}

const { state: activeTab, set: selectTab } = usePersistentString<HeatingTab>(
  TAB_KEY,
  'cards',
  isHeatingTab,
)

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
    <div class="tabbar" role="tablist" aria-label="Heating control view">
      <button
        type="button"
        role="tab"
        class="tab"
        :class="{ active: activeTab === 'cards' }"
        :aria-selected="activeTab === 'cards'"
        @click="selectTab('cards')"
      >
        Cards
      </button>
      <button
        type="button"
        role="tab"
        class="tab"
        :class="{ active: activeTab === 'schedules' }"
        :aria-selected="activeTab === 'schedules'"
        @click="selectTab('schedules')"
      >
        Schedules
      </button>
    </div>

    <!-- Content -->
    <main class="content">
      <template v-if="activeTab === 'cards'">
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
        <EmptyState v-if="orderedControls.length === 0" message="No heating controls found." min-height="300px" />
        <div v-else class="cards-grid">
          <ControlCard v-for="control in orderedControls" :key="control.id" :control="control" />
        </div>
      </template>

      <ProfileScheduleEditor v-else />
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

.tabbar {
  display: flex;
  gap: 0.3rem;
  padding: 0.5rem 2rem 0;
  border-bottom: 1px solid var(--color-border);
}

.tab {
  background: transparent;
  border: none;
  border-bottom: 2px solid transparent;
  color: var(--color-text-muted);
  font-size: 0.9rem;
  font-weight: 600;
  padding: 0.5rem 0.9rem;
  cursor: pointer;
  transition:
    color 0.2s,
    border-color 0.2s;
}

.tab:hover {
  color: var(--color-text);
}

.tab.active {
  color: var(--color-text);
  border-bottom-color: var(--color-accent);
}

.cards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(var(--card-width, 168px), 1fr));
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
