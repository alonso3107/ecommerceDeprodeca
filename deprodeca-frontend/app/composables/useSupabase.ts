// Composable de Supabase — cliente singleton para DEPRODECA
import { createClient } from "@supabase/supabase-js"

let cliente: ReturnType<typeof createClient> | null = null

function normalizarValor(valor: string | undefined) {
  return (valor || "").trim()
}

export function useSupabase() {
  if (!cliente) {
    const config = useRuntimeConfig()
    const supabaseUrl = normalizarValor(config.public.supabaseUrl)
    const supabaseAnonKey = normalizarValor(config.public.supabaseAnonKey)

    if (!supabaseUrl || !supabaseAnonKey) {
      throw new Error(
        "Supabase no está configurado: verifica SUPABASE_URL / SUPABASE_ANON_KEY o NUXT_PUBLIC_SUPABASE_URL / NUXT_PUBLIC_SUPABASE_ANON_KEY.",
      )
    }

    cliente = createClient(
      supabaseUrl,
      supabaseAnonKey,
    )
  }
  return cliente
}
