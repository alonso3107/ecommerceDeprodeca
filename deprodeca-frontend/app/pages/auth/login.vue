<!--
  ═══════════════════════════════════════════════════════════════
  auth/login.vue — Iniciar Sesión · DEPRODECA
  Validación estricta, feedback en tiempo real.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "auth" })

const email = ref("")
const password = ref("")
const loading = ref(false)
const errorMsg = ref("")
const mostrarPassword = ref(false)

// ─── Validación en tiempo real ────────────────────────────
const campoValido = (campo: string): boolean | null => {
  const val = campo === "email" ? email.value : password.value
  if (!val) return null
  if (campo === "email") return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(val)
  return val.length >= 2 // solo que no esté vacío
}

async function handleLogin() {
  errorMsg.value = ""
  if (!email.value || !password.value) {
    errorMsg.value = "Todos los campos son obligatorios"
    return
  }
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    errorMsg.value = "Email inválido"
    return
  }
  loading.value = true
  try {
    const config = useRuntimeConfig()
    const res = await $fetch<{ success: boolean; data: any; message: string }>(
      `${config.public.apiBase}/auth/login`,
      { method: "POST", body: { email: email.value, password: password.value } },
    )
    if (res.success) {
      if (import.meta.client) {
        localStorage.setItem("deprodeca_token", res.data.token)
        localStorage.setItem("deprodeca_user", JSON.stringify({
          id: res.data.usuario_id, email: res.data.email,
          nombre: res.data.nombre, rol: res.data.rol,
        }))
        window.dispatchEvent(new Event("auth-cambiado"))
      }
      await navigateTo("/")
    }
  } catch (e: any) {
    errorMsg.value = e?.data?.message || "Error de conexión"
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>

    <!-- Encabezado -->
    <div class="mb-8">
      <p class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.3em] mb-3">─── Acceso</p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Iniciar<br />Sesión<span class="text-[#D4A017]">.</span>
      </h1>
      <p class="mt-3 font-mono text-[10px] text-texto-muted uppercase tracking-[0.15em]">
        Ingresá a tu cuenta de proveedor
      </p>
    </div>

    <div class="border border-borde bg-white p-8">

      <form @submit.prevent="handleLogin" class="space-y-5" autocomplete="on" novalidate>

        <!-- Email -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="login-email">
            <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em]">Email</span>
            <span v-if="campoValido('email') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('email') === false" class="font-mono text-[10px] text-error">Email inválido</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="14" viewBox="0 0 16 14" fill="none">
              <rect x="0.5" y="1.5" width="15" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M0.5 1.5L8 7.5L15.5 1.5" stroke="currentColor" stroke-width="1.2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="login-email" v-model="email" type="email" autocomplete="email"
                   placeholder="tu@email.com" required
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-texto bg-white placeholder:text-texto-muted focus:outline-none transition-colors min-h-[48px]',
                            campoValido('email') === false ? 'border-error focus:border-error' : 'border-borde focus:border-[#D4A017]']" />
          </div>
        </div>

        <!-- Contraseña -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="login-password">
            <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em]">Contraseña</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <circle cx="8" cy="11" r="1" fill="currentColor"/>
            </svg>
            <input id="login-password" v-model="password"
                   :type="mostrarPassword ? 'text' : 'password'"
                   autocomplete="current-password" placeholder="••••••••" required
                   class="w-full border border-borde pl-11 pr-12 py-3 font-body text-body text-texto bg-white placeholder:text-texto-muted focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
            <button type="button"
                    class="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-texto-muted hover:text-texto transition-colors"
                    @click="mostrarPassword = !mostrarPassword"
                    :aria-label="mostrarPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'">
              <svg v-if="!mostrarPassword" width="18" height="14" viewBox="0 0 18 14" fill="none">
                <rect x="1" y="2" width="16" height="10" rx="5" stroke="currentColor" stroke-width="1.5"/>
                <circle cx="9" cy="7" r="2" stroke="currentColor" stroke-width="1.5"/>
              </svg>
              <svg v-else width="18" height="14" viewBox="0 0 18 14" fill="none">
                <rect x="1" y="2" width="16" height="10" rx="5" stroke="currentColor" stroke-width="1.5"/>
                <line x1="3" y1="3" x2="15" y2="11" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
              </svg>
            </button>
          </div>
        </div>

        <!-- Error -->
        <p v-if="errorMsg" class="font-mono text-[11px] text-error uppercase tracking-[0.1em] text-center font-bold border border-error px-3 py-2"
           role="alert">{{ errorMsg }}</p>

        <!-- Botón -->
        <button type="submit" :disabled="loading"
                class="w-full bg-texto text-white font-display text-heading uppercase tracking-[0.05em]
                       py-4 hover:bg-[#D4A017] hover:text-black transition-colors duration-200
                       disabled:opacity-50 disabled:cursor-not-allowed
                       min-h-[56px] flex items-center justify-center gap-3">
          <span v-if="loading" class="w-5 h-5 border-2 border-white border-t-transparent animate-spin" />
          <span v-else>Iniciar Sesión</span>
        </button>
      </form>

      <div class="relative my-8">
        <div class="absolute inset-0 flex items-center"><div class="w-full border-t-2 border-borde" /></div>
        <div class="relative flex justify-center">
          <span class="bg-white px-4 font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em]">¿Nuevo?</span>
        </div>
      </div>

      <NuxtLink to="/auth/registro"
                class="block text-center w-full py-4 border border-borde
                       font-mono text-[11px] text-texto-muted uppercase tracking-[0.15em]
                       hover:text-texto hover:border-texto/30 hover:bg-fondo transition-all duration-200">
        Crear una Cuenta Gratis
      </NuxtLink>

    </div>
  </div>
</template>
