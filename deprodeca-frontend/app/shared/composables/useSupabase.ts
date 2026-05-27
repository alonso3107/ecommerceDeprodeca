// Composable de Supabase — cliente singleton para DEPRODECA
import { createClient } from "@supabase/supabase-js"

let cliente: ReturnType<typeof createClient> | null = null

export function usarSupabase() {
  if (!cliente) {
    const config = useRuntimeConfig()
    cliente = createClient(
      config.public.supabaseUrl,
      config.public.supabaseAnonKey,
    )
  }
  return cliente
}
