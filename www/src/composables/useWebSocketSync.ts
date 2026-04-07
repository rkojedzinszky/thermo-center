import { watch } from 'vue'
import { startWebSocket, stopWebSocket } from '@/services/websocket'
import { useAuth } from '@/composables/useAuth'

export function useWebSocketSync() {
  const { isLoggedIn } = useAuth()

  watch(isLoggedIn, (logged) => {
    if (logged) {
      startWebSocket()
    } else {
      stopWebSocket()
    }
  })
}
