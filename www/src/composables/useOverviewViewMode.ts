import { ref } from 'vue'

export type OverviewViewMode = 'table' | 'cards'

const VIEW_MODE_KEY = 'sensor_view_mode'

function loadInitialViewMode(): OverviewViewMode {
  const raw = localStorage.getItem(VIEW_MODE_KEY)
  return raw === 'cards' || raw === 'table' ? raw : 'table'
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
