<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { useTheme } from '@/composables/useTheme'
import { useAuth } from '@/composables/useAuth'
import { useOverviewViewMode } from '@/composables/useOverviewViewMode'
import { useWebSocketSync } from '@/composables/useWebSocketSync'
import { PWA_UPDATE_EVENT, applyPwaUpdate } from '@/pwa'
import AppHeader from '@/components/AppHeader.vue'

// Initialises theme and keeps html[data-theme] in sync
const { pref, setTheme } = useTheme()
const { session, logout: authLogout } = useAuth()
const { viewMode, setViewMode } = useOverviewViewMode()
const route = useRoute()
const router = useRouter()
// Manage websocket lifecycle based on auth state
useWebSocketSync()

const showUpdateBanner = ref(false)

function onPwaUpdateAvailable() {
  showUpdateBanner.value = true
}

function dismissUpdateBanner() {
  showUpdateBanner.value = false
}

function installUpdate() {
  applyPwaUpdate()
}

const currentPage = computed<'overview' | 'heating'>(() =>
  route.name === 'heating' ? 'heating' : 'overview',
)

const showShellHeader = computed(() => route.meta.requiresAuth === true)

const headerTitle = computed(() =>
  currentPage.value === 'heating' ? 'Heating Control' : 'Thermo Center',
)

const headerSubtitle = computed(() =>
  currentPage.value === 'heating' ? 'Heating overview' : 'Sensor overview',
)

const headerIcon = computed(() => (currentPage.value === 'heating' ? '🔥' : '🌡️'))

const showViewToggle = computed(() => currentPage.value === 'overview')

async function logout() {
  await authLogout()
  router.push('/login')
}

onMounted(() => {
  window.addEventListener(PWA_UPDATE_EVENT, onPwaUpdateAvailable)
})

onUnmounted(() => {
  window.removeEventListener(PWA_UPDATE_EVENT, onPwaUpdateAvailable)
})
</script>

<template>
  <div class="app-shell">
    <AppHeader
      v-if="showShellHeader"
      :title="headerTitle"
      :subtitle="headerSubtitle"
      :icon="headerIcon"
      :current-page="currentPage"
      :show-view-toggle="showViewToggle"
      :username="session?.username"
      :current-theme="pref"
      :view-mode="viewMode"
      @logout="logout"
      @theme-change="setTheme"
      @view-mode-change="setViewMode"
    />

    <RouterView />
  </div>

  <div v-if="showUpdateBanner" class="pwa-update-banner" role="status" aria-live="polite">
    <span>A new version is available.</span>
    <button type="button" class="pwa-update-action" @click="installUpdate">Update</button>
    <button type="button" class="pwa-update-dismiss" @click="dismissUpdateBanner">Dismiss</button>
  </div>
</template>

<style>
*,
*::before,
*::after {
  box-sizing: border-box;
}

html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
  font-family:
    'Inter',
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    sans-serif;
  -webkit-font-smoothing: antialiased;
}

#app {
  height: 100%;
}

.app-shell {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  color: var(--color-text);
}

/* ── Dark theme (default) ── */
:root,
html[data-theme='dark'] {
  --color-bg: #0d1117;
  --color-background: #161b22;
  --color-surface: rgba(22, 27, 34, 0.88);
  --color-surface-solid: #161b22;
  --color-card-a: #1e293b;
  --color-card-b: #0f172a;
  --color-card-back-a: #0f172a;
  --color-card-back-b: #1e293b;
  --color-text: #e2e8f0;
  --color-text-secondary: #94a3b8;
  --color-text-muted: #64748b;
  --color-text-faint: #475569;
  --color-text-values: #f1f5f9;
  --color-text-fields: #cbd5e1;
  --color-border: rgba(255, 255, 255, 0.08);
  --color-border-strong: rgba(255, 255, 255, 0.12);
  --color-border-card: rgba(255, 255, 255, 0.1);
  --color-border-card-back: rgba(99, 102, 241, 0.3);
  --color-border-dialog: rgba(255, 255, 255, 0.1);
  --color-accent: #818cf8;
  --color-accent-bg: rgba(99, 102, 241, 0.2);
  --color-accent-border: rgba(99, 102, 241, 0.5);
  --color-error: #fca5a5;
  --color-error-bg: rgba(239, 68, 68, 0.15);
  --color-error-border: rgba(239, 68, 68, 0.4);
  --color-stale: #f87171;
  --color-inactive-badge: rgba(239, 68, 68, 0.2);
  --color-inactive-badge-border: rgba(239, 68, 68, 0.4);
  --color-inactive-badge-text: #fca5a5;
  --shadow-card: 0 8px 32px rgba(0, 0, 0, 0.4), inset 0 1px 0 rgba(255, 255, 255, 0.08);
  --shadow-card-back: 0 8px 32px rgba(0, 0, 0, 0.5), inset 0 1px 0 rgba(99, 102, 241, 0.1);
  --color-reading-divider: rgba(255, 255, 255, 0.08);
  --color-footer-border: rgba(255, 255, 255, 0.06);
  --color-back-footer-border: rgba(99, 102, 241, 0.15);
  --color-input-bg: rgba(255, 255, 255, 0.07);
  --color-border-input: rgba(255, 255, 255, 0.15);
  --color-input-text: #e2e8f0;
  --color-input-placeholder: #475569;
  --color-task-bg: rgba(15, 23, 42, 0.8);
  --color-border-task: rgba(99, 102, 241, 0.2);
  --color-btn-bg: linear-gradient(135deg, #3b82f6, #2563eb);
  --color-btn-shadow: rgba(59, 130, 246, 0.35);
  --color-page-bg: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%);
  --color-login-card: rgba(255, 255, 255, 0.05);
  --color-login-card-border: rgba(255, 255, 255, 0.12);
  --color-login-card-title: #e2e8f0;
  --color-login-subtitle: #94a3b8;
  --color-drag-over-bg: rgba(99, 102, 241, 0.12);
  --color-table-header-bg: rgba(30, 41, 59, 0.9);
  --color-table-bg: rgba(15, 23, 42, 0.7);
  --color-table-row-hover: rgba(255, 255, 255, 0.04);
  --color-handle: #475569;
  --color-primary: #3b82f6;
}

/* ── Light theme ── */
html[data-theme='light'] {
  --color-bg: #f1f5f9;
  --color-background: #ffffff;
  --color-surface: rgba(255, 255, 255, 0.94);
  --color-surface-solid: #ffffff;
  --color-card-a: #ffffff;
  --color-card-b: #f8fafc;
  --color-card-back-a: #f8fafc;
  --color-card-back-b: #f1f5f9;
  --color-text: #0f172a;
  --color-text-secondary: #475569;
  --color-text-muted: #64748b;
  --color-text-faint: #94a3b8;
  --color-text-values: #0f172a;
  --color-text-fields: #334155;
  --color-border: rgba(0, 0, 0, 0.08);
  --color-border-strong: rgba(0, 0, 0, 0.12);
  --color-border-card: rgba(0, 0, 0, 0.1);
  --color-border-card-back: rgba(99, 102, 241, 0.25);
  --color-border-dialog: rgba(0, 0, 0, 0.1);
  --color-accent: #6366f1;
  --color-accent-bg: rgba(99, 102, 241, 0.1);
  --color-accent-border: rgba(99, 102, 241, 0.4);
  --color-error: #dc2626;
  --color-error-bg: rgba(239, 68, 68, 0.08);
  --color-error-border: rgba(239, 68, 68, 0.3);
  --color-stale: #dc2626;
  --color-inactive-badge: rgba(239, 68, 68, 0.1);
  --color-inactive-badge-border: rgba(239, 68, 68, 0.3);
  --color-inactive-badge-text: #dc2626;
  --shadow-card: 0 4px 16px rgba(0, 0, 0, 0.1), inset 0 1px 0 rgba(255, 255, 255, 0.8);
  --shadow-card-back: 0 4px 16px rgba(0, 0, 0, 0.12), inset 0 1px 0 rgba(99, 102, 241, 0.05);
  --color-reading-divider: rgba(0, 0, 0, 0.08);
  --color-footer-border: rgba(0, 0, 0, 0.06);
  --color-back-footer-border: rgba(99, 102, 241, 0.12);
  --color-input-bg: rgba(0, 0, 0, 0.04);
  --color-border-input: rgba(0, 0, 0, 0.18);
  --color-input-text: #0f172a;
  --color-input-placeholder: #94a3b8;
  --color-task-bg: rgba(241, 245, 249, 0.9);
  --color-border-task: rgba(99, 102, 241, 0.2);
  --color-btn-bg: linear-gradient(135deg, #3b82f6, #2563eb);
  --color-btn-shadow: rgba(59, 130, 246, 0.25);
  --color-page-bg: linear-gradient(135deg, #dbeafe 0%, #e0e7ff 50%, #ddd6fe 100%);
  --color-login-card: rgba(255, 255, 255, 0.85);
  --color-login-card-border: rgba(0, 0, 0, 0.1);
  --color-login-card-title: #0f172a;
  --color-login-subtitle: #475569;
  --color-drag-over-bg: rgba(99, 102, 241, 0.08);
  --color-table-header-bg: rgba(241, 245, 249, 0.95);
  --color-table-bg: rgba(255, 255, 255, 0.8);
  --color-table-row-hover: rgba(0, 0, 0, 0.02);
  --color-handle: #94a3b8;
  --color-primary: #3b82f6;
}

/* ── Resync button (shared) ── */
.resync-button {
  background: rgba(239, 68, 68, 0.2);
  border: 1px solid rgba(239, 68, 68, 0.4);
  color: #fca5a5;
  border-radius: 0.3rem;
  font-size: 0.65rem;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.2s;
}

.resync-button:hover:not(:disabled) {
  background: rgba(239, 68, 68, 0.35);
  border-color: rgba(239, 68, 68, 0.6);
}

.resync-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.pwa-update-banner {
  position: fixed;
  right: 1rem;
  bottom: 1rem;
  left: 1rem;
  z-index: 1000;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border: 1px solid var(--color-accent-border);
  border-radius: 0.75rem;
  background: var(--color-surface-solid);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.22);
}

.pwa-update-action,
.pwa-update-dismiss {
  border: 1px solid var(--color-border-strong);
  border-radius: 0.5rem;
  background: transparent;
  cursor: pointer;
  font-size: 0.85rem;
  font-weight: 600;
  padding: 0.35rem 0.7rem;
}

.pwa-update-action {
  margin-left: auto;
  border-color: var(--color-accent-border);
  background: var(--color-accent-bg);
}

@media (min-width: 640px) {
  .pwa-update-banner {
    left: auto;
    min-width: 26rem;
    max-width: 36rem;
  }
}
</style>
