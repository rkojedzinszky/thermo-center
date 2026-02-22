<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useSensors } from '@/composables/useSensors'
import { useAuth } from '@/composables/useAuth'
import { useTheme } from '@/composables/useTheme'
import { startWebSocket, stopWebSocket } from '@/services/websocket'
import SensorCard from '@/components/SensorCard.vue'
import SensorTable from '@/components/SensorTable.vue'

const router = useRouter()
const { orderedSensors, loadSensors, reorder } = useSensors()
const { session, logout: authLogout } = useAuth()
const { pref, setTheme } = useTheme()

type ViewMode = 'cards' | 'table'
const VIEW_MODE_KEY = 'sensor_view_mode'
const viewMode = ref<ViewMode>((localStorage.getItem(VIEW_MODE_KEY) as ViewMode | null) ?? 'cards')

function setViewMode(mode: ViewMode) {
  viewMode.value = mode
  localStorage.setItem(VIEW_MODE_KEY, mode)
}

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

async function logout() {
  stopWebSocket()
  await authLogout()
  router.push('/login')
}

onMounted(async () => {
  await loadSensors()
  startWebSocket()
})

onUnmounted(() => {
  stopWebSocket()
})
</script>

<template>
  <div class="overview">
    <!-- Header -->
    <header class="overview-header">
      <div class="header-left">
        <span class="header-icon">🌡️</span>
        <div>
          <h1 class="header-title">Thermo Center</h1>
          <p class="header-sub">
            {{ orderedSensors.length }} sensor{{ orderedSensors.length !== 1 ? 's' : '' }}
          </p>
        </div>
      </div>
      <div class="header-right">
        <!-- Theme toggle -->
        <div class="view-toggle" role="group" aria-label="Theme">
          <button
            class="toggle-btn"
            :class="{ active: pref === 'light' }"
            title="Light theme"
            aria-label="Light theme"
            @click="setTheme('light')"
          >
            ☀️
          </button>
          <button
            class="toggle-btn"
            :class="{ active: pref === 'system' }"
            title="System default"
            aria-label="System default theme"
            @click="setTheme('system')"
          >
            🖥️
          </button>
          <button
            class="toggle-btn"
            :class="{ active: pref === 'dark' }"
            title="Dark theme"
            aria-label="Dark theme"
            @click="setTheme('dark')"
          >
            🌙
          </button>
        </div>

        <!-- View toggle -->
        <div class="view-toggle" role="group" aria-label="View mode">
          <button
            class="toggle-btn"
            :class="{ active: viewMode === 'cards' }"
            title="Card view"
            aria-label="Card view"
            @click="setViewMode('cards')"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <rect x="2" y="2" width="9" height="9" rx="2" />
              <rect x="13" y="2" width="9" height="9" rx="2" />
              <rect x="2" y="13" width="9" height="9" rx="2" />
              <rect x="13" y="13" width="9" height="9" rx="2" />
            </svg>
          </button>
          <button
            class="toggle-btn"
            :class="{ active: viewMode === 'table' }"
            title="Table view"
            aria-label="Table view"
            @click="setViewMode('table')"
          >
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <rect x="2" y="4" width="20" height="3" rx="1" />
              <rect x="2" y="10" width="20" height="3" rx="1" />
              <rect x="2" y="16" width="20" height="3" rx="1" />
            </svg>
          </button>
        </div>

        <div class="user-info">
          <span class="username">{{ session?.username }}</span>
          <button class="logout-btn" @click="logout">Sign out</button>
        </div>
      </div>
    </header>

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
        />
        <div v-if="orderedSensors.length === 0" class="empty-state">
          <span class="empty-icon">📡</span>
          <p>No sensors found</p>
        </div>
      </div>

      <!-- Table view -->
      <div v-else class="table-container">
        <SensorTable :sensors="orderedSensors" @reorder="reorder" />
      </div>
    </main>
  </div>
</template>

<style scoped>
.overview {
  background: var(--color-bg);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  color: var(--color-text);
}

/* Header */
.overview-header {
  align-items: center;
  background: var(--color-surface);
  backdrop-filter: blur(16px);
  border-bottom: 1px solid var(--color-border);
  display: flex;
  justify-content: space-between;
  padding: 1rem 1.5rem;
  position: sticky;
  top: 0;
  z-index: 100;
  gap: 1rem;
  flex-wrap: wrap;
}

.header-left {
  align-items: center;
  display: flex;
  gap: 0.75rem;
}

.header-icon {
  font-size: 2rem;
}

.header-title {
  color: var(--color-text);
  font-size: 1.4rem;
  font-weight: 700;
  margin: 0;
  letter-spacing: -0.01em;
}

.header-sub {
  color: var(--color-text-muted);
  font-size: 0.8rem;
  margin: 0;
}

.header-right {
  align-items: center;
  display: flex;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.view-toggle {
  display: flex;
  background: var(--color-border);
  border: 1px solid var(--color-border-strong);
  border-radius: 0.6rem;
  overflow: hidden;
}

.toggle-btn {
  background: transparent;
  border: none;
  color: var(--color-text-muted);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.45rem 0.65rem;
  font-size: 0.9rem;
  transition:
    background 0.15s,
    color 0.15s;
}

.toggle-btn:hover {
  background: var(--color-border-strong);
  color: var(--color-text-secondary);
}

.toggle-btn.active {
  background: var(--color-accent-bg);
  color: var(--color-accent);
}

.user-info {
  align-items: center;
  display: flex;
  gap: 0.75rem;
}

.username {
  color: var(--color-text-secondary);
  font-size: 0.9rem;
  font-weight: 500;
}

.logout-btn {
  background: var(--color-error-bg);
  border: 1px solid var(--color-error-border);
  border-radius: 0.5rem;
  color: var(--color-error);
  cursor: pointer;
  font-size: 0.82rem;
  padding: 0.4rem 0.85rem;
  transition:
    background 0.15s,
    border-color 0.15s;
}

.logout-btn:hover {
  filter: brightness(1.15);
}

/* Content */
.overview-content {
  flex: 1;
  padding: 1.5rem;
}

.cards-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 1.25rem;
  align-items: flex-start;
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
  max-width: 100%;
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
