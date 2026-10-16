import { ref } from 'vue'
import { usePersistentString } from '@/composables/usePersistentRef'

export type ControlViewMode = 'cards' | 'simple-table' | 'full-table'

const VIEW_MODE_KEY = 'control_view_mode'

function isViewMode(value: string): value is ControlViewMode {
  return value === 'cards' || value === 'simple-table' || value === 'full-table'
}

// Default to cards for new users
const { state: viewMode, set: setViewMode } = usePersistentString<ControlViewMode>(
  VIEW_MODE_KEY,
  'cards',
  isViewMode,
)
const reorderMode = ref(false)

export function useControlViewMode() {
  function toggleReorderMode() {
    reorderMode.value = !reorderMode.value
  }

  return {
    viewMode,
    setViewMode,
    reorderMode,
    toggleReorderMode,
  }
}
