import { ref, type Ref } from 'vue'

/** Read a raw string from localStorage, swallowing storage errors. */
export function readStoredString(key: string): string | null {
  try {
    return localStorage.getItem(key)
  } catch {
    return null
  }
}

/** Write a raw string to localStorage, swallowing storage errors. */
export function writeStoredString(key: string, value: string): void {
  try {
    localStorage.setItem(key, value)
  } catch {
    /* ignore storage errors */
  }
}

export interface PersistentStringApi<T extends string> {
  state: Ref<T>
  set: (next: T) => void
}

/**
 * String-union persisted in localStorage.
 * `isValid` decides whether a stored value is accepted, otherwise `fallback` is used.
 */
export function usePersistentString<T extends string>(
  key: string,
  fallback: T,
  isValid: (value: string) => value is T,
): PersistentStringApi<T> {
  function load(): T {
    const raw = readStoredString(key)
    if (raw != null) {
      try {
        if (isValid(raw)) return raw
      } catch {
        /* fall through to fallback */
      }
    }
    return fallback
  }

  const state = ref<T>(load()) as Ref<T>

  function set(next: T) {
    state.value = next
    writeStoredString(key, next)
  }

  return { state, set }
}

export interface PersistentIdApi {
  selectedId: Ref<number | null>
  select: (id: number) => void
}

/** Persisted numeric id (e.g. selected sensor/daytype), null when absent/invalid. */
export function usePersistentId(key: string, initial?: number | null): PersistentIdApi {
  function load(): number | null {
    if (initial !== undefined) return initial
    const raw = readStoredString(key)
    if (!raw) return null
    const parsed = Number(raw)
    return Number.isFinite(parsed) ? parsed : null
  }

  const selectedId = ref<number | null>(load())

  function select(id: number) {
    selectedId.value = id
    writeStoredString(key, String(id))
  }

  return { selectedId, select }
}
