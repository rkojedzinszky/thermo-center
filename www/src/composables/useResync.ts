import { ref, watch, type Ref } from 'vue'
import type { THSensor } from '@/api'
import api from '@/utils/api'

/**
 * Creates resync functionality for a single sensor.
 * @param sensor - The sensor to manage resync for
 * @returns Object with disabled state and handleResync function
 */
export function useResync(sensor: Ref<THSensor>) {
  const resyncDisabled = ref(false)

  watch(
    () => sensor.value.sensorResync,
    () => {
      // Re-enable the button when the sensor is updated
      resyncDisabled.value = false
    },
  )

  async function handleResync(e: Event) {
    e.stopPropagation()
    if (resyncDisabled.value || sensor.value.valid !== false) return

    try {
      resyncDisabled.value = true
      await api.createSensorResync({ sensorResyncW: { sensor: sensor.value.resourceUri } })
    } catch (error) {
      console.error('Failed to create sensor resync:', error)
      resyncDisabled.value = false
    }
  }

  return { resyncDisabled, handleResync }
}

/**
 * Creates resync functionality for managing multiple sensors.
 * @param sensors - Ref to array of sensors
 * @returns Object with disabled map and single-sensor handleResync function
 */
export function useResyncMap(sensors: Ref<THSensor[]>) {
  const resyncDisabledMap = ref<Map<number, boolean>>(new Map())

  watch(
    () => sensors.value.map((s) => s.sensorResync),
    () => {
      // Reset disabled state for all sensors when they update
      resyncDisabledMap.value = new Map()
    },
  )

  async function handleResync(sensor: THSensor, e: Event) {
    e.stopPropagation()
    if (resyncDisabledMap.value.get(sensor.id) || sensor.valid !== false) return

    try {
      resyncDisabledMap.value.set(sensor.id, true)
      await api.createSensorResync({ sensorResyncW: { sensor: sensor.resourceUri } })
    } catch (error) {
      console.error('Failed to create sensor resync:', error)
      resyncDisabledMap.value.set(sensor.id, false)
    }
  }

  return { resyncDisabledMap, handleResync }
}
