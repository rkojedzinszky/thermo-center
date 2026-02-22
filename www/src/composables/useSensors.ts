import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { THSensor } from '@/api'

const ORDER_KEY = 'sensor_order'

const sensorsMap = ref<Map<number, THSensor>>(new Map())
const orderIds = ref<number[]>(loadOrder())

function loadOrder(): number[] {
  try {
    const raw = localStorage.getItem(ORDER_KEY)
    return raw ? (JSON.parse(raw) as number[]) : []
  } catch {
    return []
  }
}

function saveOrder(ids: number[]) {
  try {
    localStorage.setItem(ORDER_KEY, JSON.stringify(ids))
  } catch {
    /* ignore */
  }
}

export function useSensors() {
  const orderedSensors = computed<THSensor[]>(() => {
    const all = Array.from(sensorsMap.value.values())
    const known = orderIds.value.filter((id) => sensorsMap.value.has(id))
    const unknownIds = all.map((s) => s.id).filter((id) => !orderIds.value.includes(id))
    const merged = [...known, ...unknownIds]
    return merged.map((id) => sensorsMap.value.get(id)!).filter(Boolean)
  })

  async function loadSensors() {
    const result = await api.listTHSensor()
    const items: THSensor[] = result.objects ?? []
    const newMap = new Map<number, THSensor>()
    for (const s of items) {
      newMap.set(s.id, s)
    }
    sensorsMap.value = newMap
    // Sync order with current sensors
    const existing = orderIds.value.filter((id) => newMap.has(id))
    const fresh = items.map((s) => s.id).filter((id) => !existing.includes(id))
    orderIds.value = [...existing, ...fresh]
    saveOrder(orderIds.value)
  }

  async function updateSensor(id: number) {
    try {
      const sensor = await api.getTHSensor({ id })
      sensorsMap.value = new Map(sensorsMap.value).set(id, sensor)
      if (!orderIds.value.includes(id)) {
        orderIds.value = [...orderIds.value, id]
        saveOrder(orderIds.value)
      }
    } catch {
      /* silently ignore individual sensor fetch errors */
    }
  }

  function reorder(fromIndex: number, toIndex: number) {
    const ids = [...orderIds.value]
    const moved = ids.splice(fromIndex, 1)[0]
    if (moved === undefined) return
    ids.splice(toIndex, 0, moved)
    orderIds.value = ids
    saveOrder(orderIds.value)
  }

  return { orderedSensors, loadSensors, updateSensor, reorder }
}
