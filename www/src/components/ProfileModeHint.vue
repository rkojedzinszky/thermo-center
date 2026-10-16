<script setup lang="ts">
import type { ProfileEditMode } from '@/composables/useProfileEditMode'

defineProps<{
  mode: ProfileEditMode
}>()
</script>

<template>
  <p class="hint">
    <template v-if="mode === 'temperature'">
      Click a zone on the bar to set its temperature. Switch to “Edit zones” to add, move or remove
      zones.
    </template>
    <template v-else>
      Click the bar (or “Add zone”) to add a zone, drag a marker to move it, click a marker to edit
      or delete it.
    </template>
  </p>
</template>

<style scoped>
.hint {
  margin: 0;
  font-size: 0.78rem;
  color: var(--color-text-muted);
}
</style>
