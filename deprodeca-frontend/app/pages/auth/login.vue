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
      <p class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.3em] mb-3">─── Acceso</p>
      <h1 class="font-display text-display-lg text-negro uppercase leading-[0.95]">
        Iniciar<br />Sesión<span class="text-dorado">.</span>
      </h1>
      <p class="mt-3 font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.15em]">
        Ingresá a tu cuenta de proveedor
      </p>
    </div>

    <div class="border border-stone bg-blanco p-8">

      <form @submit.prevent="handleLogin" class="space-y-5" autocomplete="on" novalidate>

        <!-- Email -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="login-email">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Email
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('email') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('email') === false" class="font-mono text-[10px] text-error">Email inválido</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="14" viewBox="0 0 16 14" fill="none">
              <rect x="0.5" y="1.5" width="15" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M0.5 1.5L8 7.5L15.5 1.5" stroke="currentColor" stroke-width="1.2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="login-email" v-model="email" type="email" autocomplete="email"
                   placeholder="tu@email.com" required
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('email') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
          </div>
        </div>

        <!-- Contraseña -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="login-password">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Contraseña
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <circle cx="8" cy="11" r="1" fill="currentColor"/>
            </svg>
            <input id="login-password" v-model="password"
                   :type="mostrarPassword ? 'text' : 'password'"
                   autocomplete="current-password" placeholder="••••••••" required
                   class="w-full border border-stone pl-11 pr-12 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:border-dorado focus:outline-none transition-colors duration-200 min-h-[48px]" />
            <button type="button"
                    class="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-stone-oscuro hover:text-negro transition-colors duration-200"
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
        <p v-if="errorMsg" class="font-mono text-[11px] text-error uppercase tracking-[0.1em] text-center font-bold border border-error px-3 py-2 flex items-center justify-center gap-2"
           role="alert">
          <svg width="14" height="14" viewBox="0 0 14 14" class="text-error shrink-0" fill="none" aria-hidden="true">
            <polygon points="7,1 13,12 1,12" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/>
            <line x1="7" y1="5" x2="7" y2="8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
            <circle cx="7" cy="10.5" r="0.8" fill="currentColor"/>
          </svg>
          <span>{{ errorMsg }}</span>
        </p>

        <!-- Botón -->
        <button type="submit" :disabled="loading"
                class="w-full bg-negro text-blanco font-display text-heading uppercase tracking-[0.05em]
                       py-4 hover:bg-dorado hover:text-negro transition-colors duration-200
                       disabled:opacity-50 disabled:cursor-not-allowed
                       min-h-[56px] flex items-center justify-center gap-3">
          <span v-if="loading" class="w-5 h-5 border-2 border-blanco border-t-transparent animate-spin" />
          <template v-else>
            <span>Iniciar Sesión</span>
            <svg width="16" height="12" viewBox="0 0 16 12" fill="none" aria-hidden="true">
              <path d="M0 6H14M14 6L9.5 1.5M14 6L9.5 10.5" stroke="currentColor" stroke-width="2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
          </template>
        </button>
      </form>

      <div class="relative my-8">
        <div class="absolute inset-0 flex items-center"><div class="w-full border-t border-stone" /></div>
        <div class="relative flex justify-center">
          <div class="bg-blanco px-3 flex items-center gap-2">
            <svg width="8" height="8" viewBox="0 0 8 8" class="text-dorado" aria-hidden="true">
              <rect x="4" y="0" width="5.6" height="5.6" transform="rotate(45 4 0)" fill="currentColor"/>
            </svg>
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em]">¿Nuevo?</span>
            <svg width="8" height="8" viewBox="0 0 8 8" class="text-dorado" aria-hidden="true">
              <rect x="4" y="0" width="5.6" height="5.6" transform="rotate(45 4 0)" fill="currentColor"/>
            </svg>
          </div>
        </div>
      </div>

      <NuxtLink to="/auth/registro"
                class="block text-center w-full py-4 border border-stone
                       font-mono text-[11px] text-stone-oscuro uppercase tracking-[0.15em]
                       hover:text-negro hover:border-dorado transition-colors duration-200">
        <span class="inline-flex items-center gap-2">
          Crear una Cuenta Gratis
          <svg width="8" height="8" viewBox="0 0 8 8" class="text-dorado" aria-hidden="true">
            <rect x="4" y="0" width="5.6" height="5.6" transform="rotate(45 4 0)" fill="currentColor"/>
          </svg>
        </span>
      </NuxtLink>

    </div>
  </div>
</template>
