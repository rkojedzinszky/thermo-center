<script setup lang="ts">
import Navigation from './Navigation.vue'

interface Props {
  title: string
  subtitle: string
  icon: string
  currentPage: 'overview' | 'heating'
  showViewToggle?: boolean
  username?: string
  currentTheme?: 'light' | 'system' | 'dark'
  viewMode?: 'table' | 'cards'
}

interface Emits {
  logout: []
  themeChange: [theme: 'light' | 'system' | 'dark']
  viewModeChange: [mode: 'table' | 'cards']
}

withDefaults(defineProps<Props>(), {
  showViewToggle: false,
  username: '',
  currentTheme: 'system',
  viewMode: 'cards',
})

const emit = defineEmits<Emits>()

function handleThemeChange(theme: 'light' | 'system' | 'dark') {
  emit('themeChange', theme)
}

function handleViewModeChange(mode: 'table' | 'cards') {
  emit('viewModeChange', mode)
}
</script>

<template>
  <header class="app-header">
    <div class="header-left">
      <span class="header-icon">{{ icon }}</span>
      <div>
        <h1 class="header-title">{{ title }}</h1>
        <p class="header-sub">{{ subtitle }}</p>
      </div>
    </div>

    <div class="header-right">
      <!-- View toggle (only on overview page) -->
      <div v-if="showViewToggle" class="button-group" role="group" aria-label="View mode">
        <button
          class="btn"
          :class="{ active: viewMode === 'table' }"
          title="Table view"
          @click="handleViewModeChange('table')"
        >
          <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
            <rect x="2" y="4" width="20" height="3" rx="1" />
            <rect x="2" y="10" width="20" height="3" rx="1" />
            <rect x="2" y="16" width="20" height="3" rx="1" />
          </svg>
        </button>
        <button
          class="btn"
          :class="{ active: viewMode === 'cards' }"
          title="Card view"
          @click="handleViewModeChange('cards')"
        >
          <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
            <rect x="2" y="2" width="9" height="9" rx="2" />
            <rect x="13" y="2" width="9" height="9" rx="2" />
            <rect x="2" y="13" width="9" height="9" rx="2" />
            <rect x="13" y="13" width="9" height="9" rx="2" />
          </svg>
        </button>
      </div>

      <!-- Theme toggle -->
      <div class="button-group" role="group" aria-label="Theme">
        <button
          class="btn"
          :class="{ active: currentTheme === 'light' }"
          title="Light theme"
          @click="handleThemeChange('light')"
        >
          ☀️
        </button>
        <button
          class="btn"
          :class="{ active: currentTheme === 'system' }"
          title="System default"
          @click="handleThemeChange('system')"
        >
          🖥️
        </button>
        <button
          class="btn"
          :class="{ active: currentTheme === 'dark' }"
          title="Dark theme"
          @click="handleThemeChange('dark')"
        >
          🌙
        </button>
      </div>

      <!-- User menu -->
      <div class="user-menu">
        <span class="username">{{ username }}</span>
        <button class="btn logout-btn" title="Logout" @click="$emit('logout')">↪</button>
      </div>

      <!-- Navigation menu -->
      <Navigation :current="currentPage" />
    </div>
  </header>
</template>

<style scoped>
.app-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.8rex;
  border-bottom: 1px solid var(--color-border);
  background: linear-gradient(
    to right,
    rgba(var(--color-card-a-rgb), 0.3),
    rgba(var(--color-card-b-rgb), 0.2)
  );
  backdrop-filter: blur(4px);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.header-icon {
  font-size: 2.5rem;
  line-height: 1;
}

.header-title {
  margin: 0;
  font-size: 1.4rem;
  font-weight: 700;
  letter-spacing: -0.02em;
}

.header-sub {
  margin: 0.2rem 0 0;
  font-size: 0.85rem;
  color: var(--color-text-muted);
}

.header-right {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.button-group {
  display: flex;
  gap: 0.3rem;
  background: rgba(0, 0, 0, 0.06);
  border: 1px solid var(--color-border);
  border-radius: 0.5rem;
  padding: 0.3rem;
}

.btn {
  background: transparent;
  border: none;
  color: var(--color-text-muted);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0.2rem;
  border-radius: 0.4rem;
  font-size: 1rem;
  transition: all 0.2s;
  font-weight: 500;
}

.btn:hover {
  background: rgba(0, 0, 0, 0.04);
  color: var(--color-text);
}

.btn.active {
  background: var(--color-accent-bg, rgba(99, 102, 241, 0.1));
  color: var(--color-text);
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.username {
  font-size: 0.9rem;
  color: var(--color-text-muted);
  font-weight: 500;
}

.logout-btn {
  background: transparent;
  border: 1px solid var(--color-border);
  padding: 0.4rem 0.8rem;
  border-radius: 0.5rem;
  font-size: 0.95rem;
  margin-left: 0.5rem;
}

.logout-btn:hover {
  background: rgba(0, 0, 0, 0.04);
  border-color: var(--color-text);
}

@media (max-width: 768px) {
  .header-title {
    font-size: 1.2rem;
  }

  .header-right {
    gap: 0.5rem;
  }

  .username {
    display: none;
  }

  .button-group {
    gap: 0.2rem;
    padding: 0.2rem;
  }

  .btn {
    padding: 0.35rem 0.5rem;
    font-size: 0.85rem;
  }
}
</style>
