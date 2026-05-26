// Plugin API: $fetch wrapper con base URL y token JWT automático.
export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()

  // Interceptor global de $fetch para agregar Authorization header
  const originalFetch = globalThis.$fetch

  // No needed — Nuxt $fetch can be configured per-call.
  // This plugin provides API base URL utilities.
  return {
    provide: {
      apiBase: config.public.apiBase,
      apiFetch: async (path: string, opts: any = {}) => {
        const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
        const headers: Record<string, string> = {
          "Content-Type": "application/json",
          ...opts.headers,
        }
        if (token) headers["Authorization"] = `Bearer ${token}`

        return $fetch(`${config.public.apiBase}${path}`, {
          ...opts,
          headers,
        })
      },
    },
  }
})
