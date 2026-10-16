import { describe, it, expect, beforeEach } from 'vitest'
import { usePersistentId, usePersistentString } from '../composables/usePersistentRef'

type Mode = 'a' | 'b'
function isMode(value: string): value is Mode {
  return value === 'a' || value === 'b'
}

describe('usePersistentRef', () => {
  beforeEach(() => {
    localStorage.clear()
  })

  it('falls back when nothing is stored', () => {
    const { state } = usePersistentString<Mode>('k1', 'a', isMode)
    expect(state.value).toBe('a')
  })

  it('loads a valid stored value and persists updates', () => {
    localStorage.setItem('k2', 'b')
    const { state, set } = usePersistentString<Mode>('k2', 'a', isMode)
    expect(state.value).toBe('b')
    set('a')
    expect(localStorage.getItem('k2')).toBe('a')
  })

  it('ignores invalid stored values', () => {
    localStorage.setItem('k3', 'nope')
    const { state } = usePersistentString<Mode>('k3', 'a', isMode)
    expect(state.value).toBe('a')
  })

  it('stores numeric ids and rejects non-numeric input', () => {
    const { selectedId, select } = usePersistentId('k4')
    expect(selectedId.value).toBeNull()
    select(7)
    expect(localStorage.getItem('k4')).toBe('7')

    localStorage.setItem('k5', 'oops')
    expect(usePersistentId('k5').selectedId.value).toBeNull()
  })
})
