<script setup lang="ts">
import type { InstantProfile } from '@/api'

defineProps<{
  profile: InstantProfile
}>()

defineEmits<{
  toggle: []
}>()
</script>

<template>
  <button
    class="profile-button"
    :class="{ active: profile.active }"
    :aria-label="`Toggle ${profile.name} profile`"
    @click="$emit('toggle')"
  >
    <span class="profile-name">{{ profile.name }}</span>
  </button>
</template>

<style scoped>
.profile-button {
  width: 120px;
  padding: 0.6rem 0.8rem;
  border: 1px solid var(--color-border-card);
  border-radius: 0.6rem;
  background: var(--color-card-a);
  color: var(--color-text);
  font-weight: 600;
  font-size: 0.85rem;
  cursor: pointer;
  transition:
    background-color 0.2s,
    box-shadow 0.2s,
    transform 0.15s;
  box-shadow: var(--shadow-card);
  white-space: nowrap;
  flex-shrink: 0;
  text-align: center;
  overflow: hidden;
  text-overflow: ellipsis;
}

.profile-button:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.25);
}

/* Active state styling */
.profile-button.active {
  background: linear-gradient(145deg, rgba(76, 175, 80, 0.2), rgba(76, 175, 80, 0.1));
  border-color: rgba(76, 175, 80, 0.5);
  color: var(--color-text);
}

.profile-button.active::after {
  content: '';
  display: inline-block;
  width: 0.6rem;
  height: 0.6rem;
  background: #4caf50;
  border-radius: 50%;
  margin-left: 0.5rem;
  vertical-align: middle;
}

/* Inactive state styling */
.profile-button:not(.active) {
  background: linear-gradient(145deg, rgba(200, 200, 200, 0.1), rgba(150, 150, 150, 0.05));
  border-color: rgba(128, 128, 128, 0.2);
  opacity: 0.7;
}
</style>
