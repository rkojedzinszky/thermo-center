import { ref, computed } from 'vue'
import api from '@/utils/api'
import type { Session } from '@/api'

const session = ref<Session | null>(null)
const loading = ref(false)
const error = ref<string | null>(null)

export function useAuth() {
  const isLoggedIn = computed(() => session.value !== null)

  async function checkSession(): Promise<boolean> {
    loading.value = true
    error.value = null
    try {
      session.value = await api.getSession({ id: 1 })
      return true
    } catch {
      session.value = null
      return false
    } finally {
      loading.value = false
    }
  }

  async function login(username: string, password: string): Promise<boolean> {
    loading.value = true
    error.value = null
    try {
      session.value = await api.createSession({ sessionW: { username, password } })
      return true
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'Login failed'
      return false
    } finally {
      loading.value = false
    }
  }

  async function logout(): Promise<void> {
    loading.value = true
    try {
      await api.deleteSession({ id: 1 })
    } catch {
      // Session may already be gone on the server; proceed with local cleanup
    } finally {
      session.value = null
      loading.value = false
    }
  }

  return { session, isLoggedIn, loading, error, checkSession, login, logout }
}
