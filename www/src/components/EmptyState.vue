<script setup lang="ts">
withDefaults(
  defineProps<{
    message?: string
    minHeight?: string
  }>(),
  {
    message: '',
    minHeight: '160px',
  },
)
</script>

<template>
  <div class="empty-state" :style="{ minHeight }">
    <p v-if="message">{{ message }}</p>
    <slot v-else />
  </div>
</template>

<style scoped>
.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 160px;
  color: var(--color-text-muted);
  font-size: 1rem;
}
</style>
