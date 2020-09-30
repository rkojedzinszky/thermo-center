// Service-worker for Thermo-center

const cacheVersion = "default";

const filesToCache = [
    "/",
    "/manifest.json",
    "/dist/bundles/admin-overview.js",
    "/dist/bundles/edit-heatcontrol.js",
    "/dist/bundles/thermo-center/index.css",
    "/dist/bundles/thermo-center/index.js",
    "/dist/bundles/thermo-center/pages/admin/admin.js",
    "/dist/bundles/thermo-center/pages/edit/edit.js",
    "/dist/bundles/thermo-center/pages/heatcontrol/heatcontrol.js",
    "/dist/bundles/thermo-center/pages/login/login.css",
    "/dist/bundles/thermo-center/pages/login/login.js",
    "/dist/bundles/thermo-center/pages/logout/logout.js",
    "/dist/bundles/thermo-center/pages/overview/overview.js",
    "/dist/node_modules/@fortawesome/fontawesome-free/webfonts/fa-solid-900.eot",
    "/dist/node_modules/@fortawesome/fontawesome-free/webfonts/fa-solid-900.svg",
    "/dist/node_modules/@fortawesome/fontawesome-free/webfonts/fa-solid-900.ttf",
    "/dist/node_modules/@fortawesome/fontawesome-free/webfonts/fa-solid-900.woff",
    "/dist/node_modules/@fortawesome/fontawesome-free/webfonts/fa-solid-900.woff2",
    "/dist/steal.production.js",
    "/icons/thermometer-half-solid-288x288.png",
    "/icons/thermometer-half-solid.png",
    "/icons/thermometer-half-solid.svg",
];

self.addEventListener('install', event => {
    console.log("Installing; V=", cacheVersion);

    // cache all files
    event.waitUntil(
        caches.open(cacheVersion).then(cache => cache.addAll(filesToCache))
    );
});

self.addEventListener('activate', event => {
    // Keep only CACHE
    event.waitUntil(
        caches.keys().then(keys => Promise.all(
            keys.map(key => {
                if (key != cacheVersion) {
                    return caches.delete(key);
                }
            })
        )).then(() => {
            console.log("ServiceWorker", cacheVersion, " is ready to handle fetches!");
        })
    );
});

self.addEventListener('fetch', (event) => {
    // Dont handle /api/ and /admin/ paths
    const url = new URL(event.request.url);
    if (url.pathname.startsWith("/api/") || url.pathname.startsWith("/admin/")) {
        return;
    }

    // Let browser handle non-GET requests
    if (event.request.method != "GET") {
        return;
    }

    event.respondWith(
        caches.open(cacheVersion).then(cache => cache.match(event.request).then(response => {
                if (response) {
                    return response;
                }

                return fetch(event.request).then(response => response);
            })
        )
    );
});
