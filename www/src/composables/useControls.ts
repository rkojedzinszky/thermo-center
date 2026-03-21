import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { Control } from '@/api'

export function useControls() {
  const controlsMap = ref<Map<number, Control>>(new Map())

  const orderedControls = computed<Control[]>(() => {
    return Array.from(controlsMap.value.values())
  })

  async function loadControls() {
    const result = await api.listControl()
    const items: Control[] = result.objects ?? []
    const newMap = new Map<number, Control>()
    for (const c of items) {
      newMap.set(c.id, c)
    }
    controlsMap.value = newMap
  }

  function updateControlDirect(updated: Control) {
    controlsMap.value.set(updated.id, updated)
  }

  async function updateControl(controlId: number) {
    // Only fetch if this control is already known (was fetched via list call)
    if (!controlsMap.value.has(controlId)) {
      return
    }
    try {
      const control = await api.getControl({ id: controlId })
      updateControlDirect(control)
    } catch {
      // ignore errors
    }
  }

  return {
    orderedControls,
    loadControls,
    updateControlDirect,
    updateControl,
  }
}
