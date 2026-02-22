const RECONNECT_DELAY = 1000

let ws: WebSocket | null = null
let stopped = false
let isConnected = false

type WebSocketListener = {
  onConnected?: () => void
  onUpdate?: (sensorId: number) => void
}

let listener: WebSocketListener | null = null

function getWsUrl(): string {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
  return `${protocol}//${window.location.host}/ws/`
}

function notifyConnected() {
  listener?.onConnected?.()
}

function notifyUpdate(sensorId: number) {
  listener?.onUpdate?.(sensorId)
}

function connect() {
  if (stopped) return

  ws = new WebSocket(getWsUrl())

  ws.onopen = () => {
    isConnected = true
    // Notify subscriber of connection
    notifyConnected()
  }

  ws.onmessage = (event: MessageEvent) => {
    const raw = String(event.data).trim()
    const id = Number(raw)
    if (!isNaN(id) && id > 0) {
      notifyUpdate(id)
    }
  }

  ws.onclose = () => {
    isConnected = false
    if (!stopped) {
      setTimeout(connect, RECONNECT_DELAY)
    }
  }

  ws.onerror = () => {
    ws?.close()
  }
}

export function subscribeToWebSocket(newListener: WebSocketListener): () => void {
  listener = newListener
  // If already connected, notify immediately
  if (isConnected) {
    notifyConnected()
  }
  // Return unsubscribe function
  return () => {
    listener = null
  }
}

export function startWebSocket() {
  stopped = false
  connect()
}

export function stopWebSocket() {
  stopped = true
  isConnected = false
  ws?.close()
  ws = null
}
