// useFetch — Composable reutilizable para llamadas a la API.
// Encapsula loading, error y tipado genérico.
import { ref } from "vue"

export function useFetch<T>() {
  const data = ref<T | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function execute(url: string, opts: RequestInit = {}) {
    loading.value = true
    error.value = null

    try {
      const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
      const headers: Record<string, string> = {
        "Content-Type": "application/json",
        ...(opts.headers as Record<string, string> || {}),
      }
      if (token) headers["Authorization"] = `Bearer ${token}`

      const res = await fetch(url, { ...opts, headers })
      const json = await res.json()

      if (!res.ok) {
        throw new Error(json.message || `Error ${res.status}`)
      }

      data.value = json.data as T
      return json
    } catch (e: any) {
      error.value = e.message || "Error de conexión"
      throw e
    } finally {
      loading.value = false
    }
  }

  return { data, loading, error, execute }
}
