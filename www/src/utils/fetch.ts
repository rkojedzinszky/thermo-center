import { getCookie } from './csrf'

const CSRF_METHODS = ['POST', 'PUT', 'PATCH', 'DELETE']

export default async function (input: RequestInfo | URL, init?: RequestInit): Promise<Response> {
  const options = { ...init }

  const method = options.method?.toUpperCase() || 'GET'

  if (CSRF_METHODS.includes(method)) {
    const csrfToken = getCookie('csrftoken')
    if (csrfToken) {
      const headers = new Headers(options?.headers)
      headers.set('X-CSRFToken', csrfToken)
      options.headers = headers
    }
  }

  return fetch(input, options)
}
