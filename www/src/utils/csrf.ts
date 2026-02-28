// /home/krichy/prog/thermo-center/vue-www/src/utils/csrf.ts
export function getCookie(name: string): string | null {
  let cookieValue = null
  if (document.cookie && document.cookie !== '') {
    const cookies = document.cookie.split(';')
    for (const cookie of cookies) {
      const trimmed = cookie.trim()
      // Does this cookie string begin with the name we want?
      if (trimmed.substring(0, name.length + 1) === name + '=') {
        cookieValue = decodeURIComponent(trimmed.substring(name.length + 1))
        break
      }
    }
  }
  return cookieValue
}
