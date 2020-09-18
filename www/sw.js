// Service-worker for Thermo-center
self.addEventListener('fetch', (event) => {
  return fetch(event.request);
});