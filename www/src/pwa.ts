const PWA_UPDATE_EVENT = 'pwa-update-available'

let swRegistration: ServiceWorkerRegistration | null = null
let isRefreshing = false

function notifyUpdateAvailable() {
  window.dispatchEvent(new CustomEvent(PWA_UPDATE_EVENT))
}

export function registerPwaServiceWorker() {
  if (!('serviceWorker' in navigator) || !import.meta.env.PROD) {
    return
  }

  const swUrl = `${import.meta.env.BASE_URL}sw.js?v=${encodeURIComponent(__APP_BUILD_TIME__)}`

  navigator.serviceWorker
    .register(swUrl)
    .then((registration) => {
      swRegistration = registration

      if (registration.waiting) {
        notifyUpdateAvailable()
      }

      registration.addEventListener('updatefound', () => {
        const newWorker = registration.installing

        if (!newWorker) {
          return
        }

        newWorker.addEventListener('statechange', () => {
          if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
            notifyUpdateAvailable()
          }
        })
      })
    })
    .catch((error) => {
      console.error('Service worker registration failed:', error)
    })

  navigator.serviceWorker.addEventListener('controllerchange', () => {
    if (isRefreshing) {
      return
    }

    isRefreshing = true
    window.location.reload()
  })
}

export function applyPwaUpdate() {
  if (swRegistration?.waiting) {
    swRegistration.waiting.postMessage({ type: 'SKIP_WAITING' })
  }
}

export { PWA_UPDATE_EVENT }
