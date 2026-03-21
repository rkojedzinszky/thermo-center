<script setup lang="ts">
import { ref, nextTick, watch } from 'vue'

defineOptions({
  name: 'NavigationMenu',
})

defineProps<{
  current?: 'overview' | 'heating'
}>()

const isOpen = ref(false)
const menuRef = ref<HTMLElement | null>(null)
const dropdownStyle = ref<{ top: string; right: string }>({ top: '0', right: '0' })

function closeMenu() {
  isOpen.value = false
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
    if (menuRef.value && !menuRef.value.contains(e.target as Node)) {
      isOpen.value = false
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
      <div v-if="isOpen" class="menu-dropdown" :style="dropdownStyle">
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
  color: var(--color-text);
}

.menu-button.active {
  background: rgba(99, 102, 241, 0.1);
  border-color: rgba(99, 102, 241, 0.3);
  color: var(--color-text);
}

.menu-dropdown {
  position: fixed;
  background: var(--color-popup-bg, var(--color-table-bg));
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
  color: var(--color-text);
}

.menu-link.active {
  background: rgba(99, 102, 241, 0.1);
  color: var(--color-text);
}

.menu-icon {
  font-size: 1.1rem;
  line-height: 1;
}
</style>
