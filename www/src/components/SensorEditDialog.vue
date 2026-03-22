<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import type { THSensor, PatchSensorRequest } from '@/api'
import api from '@/utils/api'
import { useConfigureSensorTask } from '@/composables/useConfigureSensorTask'

const props = defineProps<{
  sensor: THSensor | null
  open: boolean
}>()

const emit = defineEmits<{
  close: []
  updated: [sensor: THSensor]
  deleted: []
}>()

const sensorName = ref('')
const isSaving = ref(false)
const isDeleting = ref(false)
const deleteConfirmation = ref(false)
const apiError = ref<string | null>(null)

const {
  task,
  isLoading: isConfiguringLoading,
  isFinished,
  error: configError,
  createAndPollTask,
  reset: resetTask,
} = useConfigureSensorTask()

const isConfiguring = computed(() => task.value !== null)
const showTaskStatus = computed(() => task.value !== null)

watch(
  () => props.sensor,
  (newSensor) => {
    if (newSensor) {
      sensorName.value = newSensor.name
      apiError.value = null
    }
  },
  { immediate: true },
)

watch(
  () => props.open,
  (isOpen) => {
    if (!isOpen) {
      resetTask()
      deleteConfirmation.value = false
    }
  },
)

// Auto-close dialog 5 seconds after configuration finishes (unless there's an error)
watch(
  () => isFinished.value,
  (finished) => {
    if (finished && !task.value?.error) {
      const timer = setTimeout(() => {
        emit('close')
      }, 5000)
      return () => clearTimeout(timer)
    }
  },
)

async function handleSave(): Promise<void> {
  if (!props.sensor) return

  isSaving.value = true
  apiError.value = null

  try {
    const updated = await api.patchSensor({
      id: props.sensor.id,
      patchSensorRequest: { name: sensorName.value } as PatchSensorRequest,
    })
    emit('updated', updated)
  } catch (e: unknown) {
    apiError.value = e instanceof Error ? e.message : 'Failed to save sensor'
  } finally {
    isSaving.value = false
  }
}

async function handleConfigure(): Promise<void> {
  if (!props.sensor) return
  await createAndPollTask(props.sensor.id)
}

async function handleRetry(): Promise<void> {
  if (!props.sensor) return
  await createAndPollTask(props.sensor.id)
}

async function handleDelete(): Promise<void> {
  if (!props.sensor) return

  isDeleting.value = true
  apiError.value = null

  try {
    await api.deleteSensor({ id: props.sensor.id })
    emit('deleted')
    emit('close')
  } catch (e: unknown) {
    apiError.value = e instanceof Error ? e.message : 'Failed to delete sensor'
  } finally {
    isDeleting.value = false
    deleteConfirmation.value = false
  }
}

function handleCancel(): void {
  emit('close')
}

function handleKeyDown(e: KeyboardEvent): void {
  if (e.key === 'Escape') {
    e.preventDefault()
    emit('close')
  }
}

onMounted(() => {
  if (props.open) {
    document.addEventListener('keydown', handleKeyDown)
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeyDown)
})

watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      document.addEventListener('keydown', handleKeyDown)
    } else {
      document.removeEventListener('keydown', handleKeyDown)
    }
  },
)
</script>

<template>
  <div v-if="open" class="dialog-overlay" @click="handleCancel">
    <div class="dialog-container" @click.stop>
      <!-- Main Edit Panel -->
      <div v-if="!showTaskStatus" class="dialog-content">
        <h2 class="dialog-title">Edit Sensor</h2>

        <div v-if="apiError" class="error-message">{{ apiError }}</div>

        <div class="form-group">
          <label for="sensor-name" class="form-label">Name</label>
          <input
            id="sensor-name"
            v-model="sensorName"
            type="text"
            class="form-input"
            placeholder="Sensor name"
          />
        </div>

        <div class="dialog-buttons">
          <button
            v-if="!deleteConfirmation"
            class="btn btn-configure"
            :disabled="isConfiguringLoading || isSaving || isDeleting"
            @click="handleConfigure"
          >
            Configure
          </button>
          <button
            v-if="!deleteConfirmation"
            class="btn btn-save"
            :disabled="isConfiguringLoading || isSaving || isDeleting"
            @click="handleSave"
          >
            {{ isSaving ? 'Saving...' : 'Save' }}
          </button>

          <div v-if="!deleteConfirmation" style="flex: 1" />

          <button
            v-if="!deleteConfirmation"
            class="btn btn-delete"
            :disabled="isConfiguringLoading || isSaving || isDeleting"
            @click="deleteConfirmation = true"
          >
            Delete
          </button>

          <div v-if="deleteConfirmation" class="delete-confirmation">
            <button class="btn btn-delete-confirm" :disabled="isDeleting" @click="handleDelete">
              {{ isDeleting ? 'Deleting...' : 'Confirm delete' }}
            </button>
            <button
              class="btn btn-delete-cancel"
              :disabled="isDeleting"
              @click="deleteConfirmation = false"
            >
              Cancel
            </button>
          </div>

          <button v-if="!deleteConfirmation" class="btn btn-cancel" @click="handleCancel">Cancel</button>
        </div>
      </div>

      <!-- Configure Task Status Panel -->
      <div v-else class="dialog-content">
        <h2 class="dialog-title">Configure Sensor</h2>

        <div v-if="configError" class="error-message">{{ configError }}</div>

        <div class="task-status">
          <div class="status-group">
            <div class="status-item">
              <span class="status-label">ID:</span>
              <span class="status-value">{{ props.sensor?.id }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">Name:</span>
              <span class="status-value">{{ props.sensor?.name }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">Created:</span>
              <span class="status-value">{{
                task?.created ? new Date(task.created).toLocaleString() : '—'
              }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">Started:</span>
              <span class="status-value">{{
                task?.started ? new Date(task.started).toLocaleString() : '—'
              }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">First Discovery:</span>
              <span class="status-value">{{
                task?.firstDiscovery ? new Date(task.firstDiscovery).toLocaleString() : '—'
              }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">Last Discovery:</span>
              <span class="status-value">{{
                task?.lastDiscovery ? new Date(task.lastDiscovery).toLocaleString() : '—'
              }}</span>
            </div>

            <div class="status-item">
              <span class="status-label">Finished:</span>
              <span class="status-value">{{
                task?.finished ? new Date(task.finished).toLocaleString() : '—'
              }}</span>
            </div>

            <div class="status-item" :class="{ 'error-item': task?.error }">
              <span class="status-label">Error:</span>
              <span class="status-value">{{ task?.error || '—' }}</span>
            </div>

            <div class="status-item" :class="{ 'success-item': isFinished && !task?.error }">
              <span class="status-label">Status:</span>
              <span class="status-value">
                <template v-if="!isFinished">⏳ Running...</template>
                <template v-else-if="task?.error">❌ Error</template>
                <template v-else>✓ Configuration complete</template>
              </span>
            </div>
          </div>
        </div>

        <div class="dialog-buttons">
          <button v-if="task?.error && isFinished" class="btn btn-configure" @click="handleRetry">
            Retry
          </button>
          <div style="flex: 1" />
          <button
            class="btn"
            :class="isFinished && !task?.error ? 'btn-close' : 'btn-cancel'"
            @click="handleCancel"
          >
            {{ isFinished && !task?.error ? 'Close' : 'Cancel' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dialog-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(2px);
}

.dialog-container {
  background: var(--color-background);
  border: 1px solid var(--color-border-dialog);
  border-radius: 0.8rem;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.3);
  max-width: 500px;
  width: 90%;
  max-height: 90vh;
  overflow: auto;
}

.dialog-content {
  padding: 1.5rem;
}

.dialog-title {
  margin: 0 0 1rem 0;
  font-size: 1.3rem;
  font-weight: 700;
}

.error-message {
  background: rgba(239, 68, 68, 0.15);
  border: 1px solid rgba(239, 68, 68, 0.3);
  color: #fca5a5;
  padding: 0.75rem;
  border-radius: 0.4rem;
  margin-bottom: 1rem;
  font-size: 0.9rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
}

.form-label {
  font-size: 0.9rem;
  font-weight: 600;
}

.form-input {
  padding: 0.6rem;
  background: var(--color-input-bg);
  border: 1px solid var(--color-border-input);
  border-radius: 0.4rem;
  font-family: inherit;
  font-size: 0.95rem;
}

.form-input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.1);
}

.task-status {
  background: var(--color-task-bg);
  border: 1px solid var(--color-border-task);
  border-radius: 0.5rem;
  padding: 1rem;
  margin-bottom: 1.5rem;
}

.status-group {
  display: flex;
  flex-direction: column;
  gap: 0.7rem;
}

.status-item {
  display: flex;
  align-items: flex-start;
  gap: 0.8rem;
  font-size: 0.9rem;
}

.status-item.error-item {
  color: #fca5a5;
  background: rgba(239, 68, 68, 0.1);
  padding: 0.5rem;
  border-radius: 0.3rem;
  margin: -0.5rem;
}

.status-item.success-item {
  color: #86efac;
  background: rgba(34, 197, 94, 0.1);
  padding: 0.5rem;
  border-radius: 0.3rem;
  margin: -0.5rem;
}

.status-label {
  font-weight: 600;
  color: var(--color-text-muted);
  min-width: 140px;
  flex-shrink: 0;
}

.status-value {
  flex: 1;
  word-break: break-word;
}

.dialog-buttons {
  display: flex;
  gap: 0.6rem;
  align-items: center;
  flex-wrap: wrap;
}

.btn {
  padding: 0.6rem 1rem;
  border: none;
  border-radius: 0.4rem;
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-configure {
  background: rgba(59, 130, 246, 0.2);
  border: 1px solid rgba(59, 130, 246, 0.6);
  color: #3b82f6;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
  min-width: 80px;
}

.btn-configure:hover:not(:disabled) {
  background: rgba(59, 130, 246, 0.35);
  border-color: rgba(59, 130, 246, 0.8);
}

.btn-save {
  background: rgba(34, 197, 94, 0.2);
  border: 1px solid rgba(34, 197, 94, 0.6);
  color: #16a34a;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
}

.btn-save:hover:not(:disabled) {
  background: rgba(34, 197, 94, 0.35);
  border-color: rgba(34, 197, 94, 0.8);
}

.btn-delete {
  background: rgba(239, 68, 68, 0.2);
  border: 1px solid rgba(239, 68, 68, 0.6);
  color: #dc2626;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
}

.btn-delete:hover:not(:disabled) {
  background: rgba(239, 68, 68, 0.35);
  border-color: rgba(239, 68, 68, 0.8);
}

.btn-delete-confirm {
  background: rgba(239, 68, 68, 0.3);
  border: 1px solid rgba(239, 68, 68, 0.6);
  color: #dc2626;
  padding: 0.5rem 0.8rem;
  font-size: 0.85rem;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
}

.btn-delete-confirm:hover:not(:disabled) {
  background: rgba(239, 68, 68, 0.45);
}

.btn-delete-cancel {
  background: rgba(107, 114, 128, 0.2);
  border: 1px solid rgba(107, 114, 128, 0.5);
  color: #4b5563;
  padding: 0.5rem 0.8rem;
  font-size: 0.85rem;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.3);
}

.btn-delete-cancel:hover:not(:disabled) {
  background: rgba(107, 114, 128, 0.3);
}

.btn-cancel {
  background: rgba(107, 114, 128, 0.2);
  border: 1px solid rgba(107, 114, 128, 0.5);
  color: #4b5563;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.3);
  min-width: 80px;
}

.btn-cancel:hover:not(:disabled) {
  background: rgba(107, 114, 128, 0.3);
  border-color: rgba(107, 114, 128, 0.7);
}

.btn-close {
  background: rgba(34, 197, 94, 0.2);
  border: 1px solid rgba(34, 197, 94, 0.6);
  color: #16a34a;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
  min-width: 80px;
}

.btn-close:hover:not(:disabled) {
  background: rgba(34, 197, 94, 0.35);
  border-color: rgba(34, 197, 94, 0.8);
}

.btn-back {
  background: rgba(59, 130, 246, 0.2);
  border: 1px solid rgba(59, 130, 246, 0.6);
  color: #3b82f6;
  width: 100%;
  text-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
}

.btn-back:hover:not(:disabled) {
  background: rgba(59, 130, 246, 0.35);
  border-color: rgba(59, 130, 246, 0.8);
}

.error-message-text {
  background: rgba(239, 68, 68, 0.15);
  color: #fca5a5;
  border: 1px solid rgba(239, 68, 68, 0.3);
}

.delete-confirmation {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  width: 100%;
  justify-content: space-between;
}
</style>
