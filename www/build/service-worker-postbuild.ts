import fs from 'fs'
import path from 'path'
import type { Plugin } from 'vite'

/**
 * Vite plugin to inject prefetch file list into service worker
 * Reads all files from dist/ directory and injects them into sw.js
 * Excludes: source maps, and other non-essential files
 */
export function serviceWorkerPostbuild(): Plugin {
  return {
    name: 'service-worker-postbuild',
    apply: 'build',
    enforce: 'post',
    async generateBundle() {
      // This hook runs after all files are written to disk
    },
    writeBundle() {
      const generateCacheVersion = () => {
        const timestamp = Date.now()
        const randomPart = Math.random().toString(36).slice(2, 10)
        return `v${timestamp}-${randomPart}`
      }
      const cacheVersion = generateCacheVersion()
      // This hook is called after all files are written
      const distDir = path.resolve(__dirname, '../dist')
      const swPath = path.join(distDir, 'sw.js')

      if (!fs.existsSync(distDir)) {
        console.warn('[inject-sw-precache] dist/ directory not found')
        return
      }

      // Recursively get all files in dist/
      function getAllFiles(dir: string, basePath = ''): string[] {
        const files: string[] = []

        const entries = fs.readdirSync(dir, { withFileTypes: true })

        for (const entry of entries) {
          const fullPath = path.join(dir, entry.name)
          const relativePath = path.join(basePath, entry.name)

          if (entry.isDirectory()) {
            files.push(...getAllFiles(fullPath, relativePath))
          } else if (entry.isFile()) {
            // Convert to Unix-style path for consistency
            const unixPath = relativePath.replace(/\\/g, '/')
            files.push(unixPath)
          }
        }

        return files
      }

      // Get all files and filter
      const allFiles = getAllFiles(distDir)

      // Files to exclude
      const excludePatterns = [
        /\.map$/, // Source maps
        /sw\.js$/, // Service worker itself
      ]

      const filesToPrefetch = allFiles
        .filter((file) => !excludePatterns.some((pattern) => pattern.test(file)))
        .sort()

      // Read current service worker
      let swContent = fs.readFileSync(swPath, 'utf-8')

      // Replace CACHE_VERSION in service worker with a unique value for each build.
      const cacheVersionDeclaration = `const CACHE_VERSION = '${cacheVersion}'`
      swContent = swContent.replace(/const CACHE_VERSION = ['"][^'"]*['"]/, cacheVersionDeclaration)

      // Find the OFFLINE_URLS array and update it with all files
      // Convert file paths to have ./ prefix for service worker
      const prefetchUrls = filesToPrefetch.map((file) => `./${file}`)

      // Create the new OFFLINE_URLS array
      const offlineUrlsDeclaration = `const OFFLINE_URLS = [${prefetchUrls.map((url) => `'${url}'`).join(', ')}]`

      // Replace the OFFLINE_URLS declaration
      swContent = swContent.replace(/const OFFLINE_URLS = \[.*?\]/s, offlineUrlsDeclaration)

      // Write the updated service worker
      fs.writeFileSync(swPath, swContent)

      console.log(
        `[service-worker-postbuild] Set CACHE_VERSION=${cacheVersion} and injected ${filesToPrefetch.length} files into service worker`,
      )
    },
  }
}
