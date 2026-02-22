import { useSensors } from '@/composables/useSensors'

const RECONNECT_DELAY = 1000

let ws: WebSocket | null = null
let stopped = false

function getWsUrl(): string {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  return `${protocol}//${window.location.host}/ws/`
}

function connect() {
  if (stopped) return
  const { loadSensors, updateSensor } = useSensors()

  ws = new WebSocket(getWsUrl())

  ws.onopen = () => {
    // Full refresh after reconnect
    loadSensors()
  }

  ws.onmessage = (event: MessageEvent) => {
    const raw = String(event.data).trim()
    const id = Number(raw)
    if (!isNaN(id) && id > 0) {
      updateSensor(id)
    }
  }

  ws.onclose = () => {
    if (!stopped) {
      setTimeout(connect, RECONNECT_DELAY)
    }
  }

  ws.onerror = () => {
    ws?.close()
  }
}

export function startWebSocket() {
  stopped = false
  connect()
}

export function stopWebSocket() {
  stopped = true
  ws?.close()
  ws = null
}
