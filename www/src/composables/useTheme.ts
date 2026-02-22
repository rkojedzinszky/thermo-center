import { ref, computed, watchEffect } from 'vue'

export type ThemePref = 'light' | 'dark' | 'system'

const THEME_KEY = 'theme_preference'

const storedTheme =
  typeof localStorage !== 'undefined' ? (localStorage.getItem(THEME_KEY) as ThemePref | null) : null

const pref = ref<ThemePref>(storedTheme ?? 'system')

function getSystemTheme(): 'light' | 'dark' {
  if (typeof window === 'undefined' || !window.matchMedia) return 'dark'
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

function applyTheme(p: ThemePref) {
  if (typeof document === 'undefined') return
  const effective = p === 'system' ? getSystemTheme() : p
  document.documentElement.setAttribute('data-theme', effective)
}

// Apply immediately and whenever pref changes
watchEffect(() => applyTheme(pref.value))

// React to OS-level theme changes when in "system" mode
if (typeof window !== 'undefined' && window.matchMedia) {
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
    if (pref.value === 'system') applyTheme('system')
  })
}

export function useTheme() {
  const effectiveTheme = computed<'light' | 'dark'>(() =>
    pref.value === 'system' ? getSystemTheme() : pref.value,
  )

  function setTheme(p: ThemePref) {
    pref.value = p
    if (typeof localStorage !== 'undefined') localStorage.setItem(THEME_KEY, p)
  }

  return { pref, effectiveTheme, setTheme }
}
