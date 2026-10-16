<script setup lang="ts">
const props = defineProps<{
  items: { id: number; label: string }[]
  selectedId: number | null
  groupLabel: string
}>()

const emit = defineEmits<{
  select: [id: number]
}>()

function onKeydown(event: KeyboardEvent) {
  if (event.key !== 'ArrowLeft' && event.key !== 'ArrowRight') return
  const currentIndex = props.items.findIndex((item) => item.id === props.selectedId)
  if (currentIndex === -1) return
  event.preventDefault()
  const delta = event.key === 'ArrowRight' ? 1 : -1
  const next = props.items[(currentIndex + delta + props.items.length) % props.items.length]
  if (next) {
    emit('select', next.id)
    // Move focus to the newly selected tab.
    const strip = (event.currentTarget as HTMLElement | null)?.parentElement
    const tabs = strip?.querySelectorAll<HTMLButtonElement>('.tab')
    const idx = (currentIndex + delta + props.items.length) % props.items.length
    tabs?.[idx]?.focus()
  }
}
</script>

<template>
  <div class="tab-strip" role="tablist" :aria-label="groupLabel" @keydown="onKeydown">
    <button
      v-for="item in items"
      :key="item.id"
      type="button"
      role="tab"
      class="tab"
      :class="{ active: item.id === selectedId }"
      :aria-selected="item.id === selectedId"
      :tabindex="item.id === selectedId ? 0 : -1"
      @click="$emit('select', item.id)"
    >
      <span class="tab-name">{{ item.label }}</span>
    </button>
  </div>
</template>

<style scoped>
.tab-strip {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.tab {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  border: 1px solid var(--color-border-card);
  background: var(--color-card-a);
  color: var(--color-text);
  border-radius: 0.5rem;
  padding: 0.45rem 0.8rem;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition:
    background-color 0.2s,
    border-color 0.2s;
}

.tab:hover {
  border-color: var(--color-accent-border);
}

.tab.active {
  background: var(--color-accent-bg);
  border-color: var(--color-accent-border);
}
</style>
