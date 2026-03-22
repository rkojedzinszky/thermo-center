import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { Control } from '@/api'

const controlsMap = ref<Map<number, Control>>(new Map())

const orderedControls = computed<Control[]>(() => {
  return Array.from(controlsMap.value.values())
})

export function useControls() {
  async function loadControls() {
    const result = await api.listControl()
    const items: Control[] = result.objects ?? []
    const newMap = new Map<number, Control>()
    for (const c of items) {
      newMap.set(c.sensorId, c)
    }
    controlsMap.value = newMap
  }

  function updateControlDirect(updated: Control) {
    controlsMap.value.set(updated.sensorId, updated)
  }

  async function updateControl(sensorId: number) {
    // Only fetch if this control is already known (was fetched via list call)
    if (!controlsMap.value.has(sensorId)) {
      return
    }
    try {
      const control = await api.listControl({ sensorId: sensorId }).then((res) => res.objects?.[0])
      if (control) {
        updateControlDirect(control)
      }
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
