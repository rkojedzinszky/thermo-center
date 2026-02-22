import { ref, onMounted, onUnmounted } from 'vue'
import type { THSensor } from '@/api'

const INACTIVE_THRESHOLD = 300 // 5 minutes in seconds

const now = ref(Math.floor(Date.now() / 1000))
let timer: ReturnType<typeof setInterval> | null = null
let refCount = 0

function startTimer() {
  if (refCount === 0) {
    timer = setInterval(() => {
      now.value = Math.floor(Date.now() / 1000)
    }, 1000)
  }
  refCount++
}

function stopTimer() {
  refCount--
  if (refCount === 0 && timer) {
    clearInterval(timer)
    timer = null
  }
}

export function useTimerSync() {
  onMounted(() => startTimer())
  onUnmounted(() => stopTimer())
  return { now }
}

/** Format a number to at most 2 decimal places, stripping trailing zeros. */
export function fmt(n: number): string {
  return parseFloat(n.toFixed(2)).toString()
}

export function formatAge(tsf: number | null | undefined): string {
  if (tsf == null) return 'No data'
  const diff = Math.floor(now.value - tsf)
  if (diff < 2) return 'just now'
  if (diff < 60) return `${diff} second${diff === 1 ? '' : 's'} ago`
  const minutes = Math.floor(diff / 60)
  if (minutes < 60) return `${minutes} minute${minutes === 1 ? '' : 's'} ago`
  const hours = Math.floor(minutes / 60)
  return `${hours} hour${hours === 1 ? '' : 's'} ago`
}

export function checkInactive(sensor: THSensor, current: number): boolean {
  if (sensor.lastTsf == null) return true
  return current - sensor.lastTsf > INACTIVE_THRESHOLD
}

export { INACTIVE_THRESHOLD }
