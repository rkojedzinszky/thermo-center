<script setup lang="ts">
import { ref, nextTick, watch } from 'vue'

defineOptions({
  name: 'NavigationMenu',
})

defineProps<{
  current?: 'overview' | 'heating'
  currentTheme?: 'light' | 'system' | 'dark'
  reorderMode?: boolean
}>()

const emit = defineEmits<{
  logout: []
  themeChange: [theme: 'light' | 'system' | 'dark']
  toggleReorder: []
}>()

const isOpen = ref(false)
const isThemeOpen = ref(false)
const menuRef = ref<HTMLElement | null>(null)
const dropdownRef = ref<HTMLElement | null>(null)
const dropdownStyle = ref<{ top: string; right: string }>({ top: '0', right: '0' })

function closeMenu() {
  isOpen.value = false
  isThemeOpen.value = false
}

function toggleThemeSubmenu() {
  isThemeOpen.value = !isThemeOpen.value
}

function selectTheme(theme: 'light' | 'system' | 'dark') {
  emit('themeChange', theme)
  isThemeOpen.value = false
}

function logout() {
  emit('logout')
  closeMenu()
}

function onToggleReorder() {
  emit('toggleReorder')
  closeMenu()
}

async function updateDropdownPosition() {
  if (!menuRef.value || !isOpen.value) return

  await nextTick()

  const button = menuRef.value.querySelector('.menu-button') as HTMLElement
  if (!button) return

  const rect = button.getBoundingClientRect()
  dropdownStyle.value = {
    top: `${rect.bottom + 8}px`,
    right: `${window.innerWidth - rect.right}px`,
  }
}

watch(isOpen, async () => {
  if (isOpen.value) {
    await updateDropdownPosition()
  }
})

// Close menu when clicking outside
if (typeof window !== 'undefined') {
  const handleClickOutside = (e: MouseEvent) => {
    const target = e.target as Node

    const isInMenu = menuRef.value?.contains(target)
    const isInDropdown = dropdownRef.value?.contains(target)

    if (!isInMenu && !isInDropdown) {
      isOpen.value = false
      isThemeOpen.value = false
    }
  }

  watch(isOpen, (newVal) => {
    if (newVal) {
      document.addEventListener('click', handleClickOutside)
    } else {
      document.removeEventListener('click', handleClickOutside)
    }
  })
}
</script>

<template>
  <div class="nav-menu" ref="menuRef">
    <button
      class="menu-button"
      :class="{ active: isOpen }"
      title="Navigation menu"
      @click="isOpen = !isOpen"
    >
      ☰
    </button>

    <Teleport to="body">
      <div v-if="isOpen" class="menu-dropdown" ref="dropdownRef" :style="dropdownStyle">
        <router-link
          to="/overview"
          class="menu-link"
          :class="{ active: current === 'overview' }"
          @click="closeMenu"
        >
          <span class="menu-icon">🌡️</span>
          <span>Sensors Overview</span>
        </router-link>
        <router-link
          to="/heating"
          class="menu-link"
          :class="{ active: current === 'heating' }"
          @click="closeMenu"
        >
          <span class="menu-icon">🔥</span>
          <span>Heating Control</span>
        </router-link>

        <div class="menu-divider"></div>

        <button
          v-if="current === 'overview'"
          class="menu-link"
          :class="{ active: reorderMode }"
          type="button"
          @click="onToggleReorder"
        >
          <span class="menu-icon">⇅</span>
          <span>Reorder</span>
          <span v-if="reorderMode" class="menu-check">✓</span>
        </button>

        <button class="menu-link submenu-trigger" type="button" @click.stop="toggleThemeSubmenu">
          <span class="menu-icon">🎨</span>
          <span>Theme</span>
          <span class="submenu-arrow">{{ isThemeOpen ? '▲' : '▼' }}</span>
        </button>

        <div v-if="isThemeOpen" class="submenu">
          <button
            class="submenu-item"
            :class="{ active: currentTheme === 'light' }"
            @click.stop="selectTheme('light')"
            type="button"
          >
            Light
          </button>
          <button
            class="submenu-item"
            :class="{ active: currentTheme === 'system' }"
            @click.stop="selectTheme('system')"
            type="button"
          >
            System
          </button>
          <button
            class="submenu-item"
            :class="{ active: currentTheme === 'dark' }"
            @click.stop="selectTheme('dark')"
            type="button"
          >
            Dark
          </button>
        </div>

        <button class="menu-link" type="button" @click="logout">
          <span class="menu-icon">↪</span>
          <span>Logout</span>
        </button>
      </div>
    </Teleport>
  </div>
</template>

<style scoped>
.nav-menu {
  position: relative;
  display: flex;
  align-items: center;
}

.menu-button {
  background: transparent;
  border: 1px solid var(--color-border);
  color: var(--color-text-muted);
  width: 2.2rem;
  height: 2.2rem;
  border-radius: 0.5rem;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.menu-button:hover {
  background: rgba(0, 0, 0, 0.04);
}

.menu-button.active {
  background: rgba(99, 102, 241, 0.1);
  border-color: rgba(99, 102, 241, 0.3);
}

.menu-dropdown {
  position: fixed;
  background: var(--color-surface-solid, var(--color-table-bg));
  border: 1px solid var(--color-border);
  border-radius: 0.6rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  overflow: hidden;
  min-width: 180px;
  z-index: 9999;
}

.menu-link {
  display: flex;
  align-items: center;
  gap: 0.8rem;
  padding: 0.75rem 1rem;
  text-decoration: none;
  color: var(--color-text-muted);
  transition: all 0.2s;
  border: none;
  background: transparent;
  cursor: pointer;
  width: 100%;
  text-align: left;
  font-size: 0.95rem;
  font-weight: 500;
}

.menu-link:hover {
  background: rgba(0, 0, 0, 0.05);
}

.menu-link.active {
  background: rgba(99, 102, 241, 0.1);
}

.menu-icon {
  font-size: 1.1rem;
  line-height: 1;
}

.menu-divider {
  border-top: 1px solid var(--color-border);
  margin: 0.4rem 0;
}

.submenu-trigger {
  justify-content: space-between;
  width: 100%;
}

.submenu-arrow {
  margin-left: auto;
  font-size: 0.8rem;
  color: var(--color-text-muted);
}

.menu-check {
  margin-left: auto;
  font-size: 0.85rem;
  color: var(--color-accent, #6366f1);
}

.submenu {
  display: flex;
  flex-direction: column;
  border-left: 2px solid var(--color-border);
  margin: 0 0.5rem 0.5rem;
  padding-left: 0.5rem;
}

.submenu-item {
  text-align: left;
  padding: 0.5rem;
  border: none;
  background: transparent;
  color: var(--color-text-muted);
  cursor: pointer;
  border-radius: 0.35rem;
  margin-top: 0.15rem;
  font-size: 0.9rem;
}

.submenu-item:hover {
  background: rgba(0, 0, 0, 0.05);
}

.submenu-item.active {
  background: rgba(99, 102, 241, 0.12);
  color: var(--color-text);
}
</style>
