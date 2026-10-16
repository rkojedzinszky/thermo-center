import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { DayType } from '@/api'

const dayTypesMap = ref<Map<number, DayType>>(new Map())

const dayTypes = computed<DayType[]>(() => Array.from(dayTypesMap.value.values()))

export function useDayTypes() {
  async function loadDayTypes() {
    try {
      const result = await api.listDayType()
      const items: DayType[] = result.objects ?? []
      const newMap = new Map<number, DayType>()
      for (const dayType of items) {
        newMap.set(dayType.id, dayType)
      }
      dayTypesMap.value = newMap
    } catch (error) {
      console.error('Failed to load day types:', error)
    }
  }

  return { dayTypes, loadDayTypes }
}
