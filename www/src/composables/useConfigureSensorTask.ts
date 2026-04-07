import { ref, computed } from 'vue'
import type { ConfigureSensorTask } from '@/api'
import api from '@/utils/api'

export function useConfigureSensorTask() {
  const task = ref<ConfigureSensorTask | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const pollingInterval = ref<ReturnType<typeof setInterval> | null>(null)

  const isFinished = computed(() => task.value?.finished != null)

  async function createAndPollTask(sensorId: number): Promise<void> {
    isLoading.value = true
    error.value = null

    try {
      // Create the task
      task.value = await api.createConfigureSensorTask({
        configureSensorTaskW: { sensorId },
      })

      // Start polling every second
      startPolling()
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'Failed to create configure task'
      task.value = null
    } finally {
      isLoading.value = false
    }
  }

  function startPolling(): void {
    if (pollingInterval.value) clearInterval(pollingInterval.value)

    pollingInterval.value = setInterval(async () => {
      if (!task.value) return

      try {
        const updated = await api.getConfigureSensorTask({ id: task.value.id })
        task.value = updated

        // Stop polling if finished
        if (updated.finished != null) {
          stopPolling()
        }
      } catch (e: unknown) {
        error.value = e instanceof Error ? e.message : 'Failed to fetch task status'
      }
    }, 1000)
  }

  function stopPolling(): void {
    if (pollingInterval.value) {
      clearInterval(pollingInterval.value)
      pollingInterval.value = null
    }
  }

  function reset(): void {
    stopPolling()
    task.value = null
    error.value = null
    isLoading.value = false
  }

  return {
    task,
    isLoading,
    error,
    isFinished,
    createAndPollTask,
    stopPolling,
    reset,
  }
}
