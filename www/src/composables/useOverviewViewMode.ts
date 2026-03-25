import { ref } from 'vue'

export type OverviewViewMode = 'cards' | 'simple-table' | 'full-table'

const VIEW_MODE_KEY = 'sensor_view_mode'

function loadInitialViewMode(): OverviewViewMode {
  const raw = localStorage.getItem(VIEW_MODE_KEY)
  if (raw === 'cards' || raw === 'simple-table' || raw === 'full-table') {
    return raw
  }
  // Default to cards for new users
  return 'cards'
}

const viewMode = ref<OverviewViewMode>(loadInitialViewMode())

export function useOverviewViewMode() {
  function setViewMode(mode: OverviewViewMode) {
    viewMode.value = mode
    localStorage.setItem(VIEW_MODE_KEY, mode)
  }

  return {
    viewMode,
    setViewMode,
  }
}
