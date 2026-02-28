<script setup lang="ts">
import { useTheme } from '@/composables/useTheme'
import { useWebSocketSync } from '@/composables/useWebSocketSync'
// Initialises theme and keeps html[data-theme] in sync
useTheme()
// Manage websocket lifecycle based on auth state
useWebSocketSync()
</script>

<template>
  <RouterView />
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
  font-family:
    'Inter',
    system-ui,
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    sans-serif;
  -webkit-font-smoothing: antialiased;
}

/* ── Dark theme (default) ── */
:root,
html[data-theme='dark'] {
  --color-bg: #0d1117;
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
  --color-input-border: rgba(255, 255, 255, 0.15);
  --color-input-text: #e2e8f0;
  --color-input-placeholder: #475569;
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
}

/* ── Light theme ── */
html[data-theme='light'] {
  --color-bg: #f1f5f9;
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
  --color-input-border: rgba(0, 0, 0, 0.18);
  --color-input-text: #0f172a;
  --color-input-placeholder: #94a3b8;
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
}

/* ── Resync button (shared) ── */
.resync-button {
  background: rgba(239, 68, 68, 0.2);
  border: 1px solid rgba(239, 68, 68, 0.4);
  color: #000000;
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
</style>
