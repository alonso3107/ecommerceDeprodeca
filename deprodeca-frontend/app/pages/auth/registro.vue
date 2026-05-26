<!--
  ═══════════════════════════════════════════════════════════════
  auth/registro.vue — Crear Cuenta · DEPRODECA
  Brutalismo + autocomplete seguro + validación client-side.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "auth" })

const form = ref({
  nombre: "", empresa: "", ruc: "", telefono: "",
  email: "", password: "", confirmPassword: "",
})
const loading = ref(false)
const errorMsg = ref("")
const successMsg = ref("")
const mostrarPassword = ref(false)

function validar(): string | null {
  const f = form.value
  if (!f.nombre || !f.empresa || !f.ruc || !f.email || !f.password)
    return "Completá todos los campos obligatorios (*)"
  if (f.password.length < 6) return "La contraseña debe tener al menos 6 caracteres"
  if (f.password !== f.confirmPassword) return "Las contraseñas no coinciden"
  if (f.ruc.length !== 11) return "El RUC debe tener 11 dígitos"
  return null
}

async function handleRegistro() {
  const err = validar()
  if (err) { errorMsg.value = err; return }
  loading.value = true; errorMsg.value = ""
  try {
    const config = useRuntimeConfig()
    const res = await $fetch<{ success: boolean; data: any; message: string }>(
      `${config.public.apiBase}/auth/registro`, {
        method: "POST",
        body: {
          email: form.value.email, password: form.value.password,
          nombre: form.value.nombre, empresa: form.value.empresa,
          ruc: form.value.ruc, telefono: form.value.telefono,
        },
      })
    if (res.success) {
      successMsg.value = "¡Cuenta creada! Redirigiendo..."
      if (import.meta.client) {
        localStorage.setItem("deprodeca_token", res.data.token)
        localStorage.setItem("deprodeca_user", JSON.stringify({
          id: res.data.usuario_id, email: res.data.email,
          nombre: res.data.nombre, rol: res.data.rol,
        }))
      }
      setTimeout(() => navigateTo("/"), 1500)
    }
  } catch (e: any) {
    errorMsg.value = e?.data?.message || "Error al crear la cuenta"
  } finally { loading.value = false }
}
</script>

<template>
  <div>

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Registro
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Crear<br />Cuenta<span class="text-[#D4A017]">.</span>
      </h1>
      <p class="mt-3 font-mono text-[10px] text-texto-muted uppercase tracking-[0.15em]">
        Registrate como proveedor
      </p>
    </div>

    <div class="border border-borde bg-white p-8">

      <!-- Éxito -->
      <div v-if="successMsg" class="text-center py-8">
        <svg class="mx-auto mb-4 text-exito" width="48" height="48" viewBox="0 0 48 48" fill="none">
          <path d="M24 4L44 24L24 44L4 24L24 4Z" stroke="currentColor" stroke-width="2"/>
          <path d="M18 25L22 30L32 20" stroke="currentColor" stroke-width="3"
                stroke-linecap="square" stroke-linejoin="miter"/>
        </svg>
        <p class="font-display text-display-md text-texto uppercase mb-2">
          ¡Cuenta Creada<span class="text-exito">!</span>
        </p>
        <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider">Redirigiendo...</p>
      </div>

      <!-- Formulario -->
      <form v-else @submit.prevent="handleRegistro" class="space-y-4" autocomplete="on">

        <!-- Nombre -->
        <div>
          <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-nombre">
            Nombre Completo *
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <circle cx="8" cy="5" r="2.5" stroke="currentColor" stroke-width="1.2"/>
              <path d="M3 14C3 11.2 5.2 9 8 9C10.8 9 13 11.2 13 14"
                    stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
            </svg>
            <input id="reg-nombre" v-model="form.nombre" autocomplete="name"
                   placeholder="Tu nombre completo" required
                   class="w-full border border-borde pl-11 pr-4 py-3 font-body text-body text-texto
                          bg-white placeholder:text-texto-muted
                          focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
          </div>
        </div>

        <!-- Empresa -->
        <div>
          <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-empresa">
            Bodega / Empresa *
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="4" y="4" width="8" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M2.5 4L8 1L13.5 4" stroke="currentColor" stroke-width="1.2"
                    stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="reg-empresa" v-model="form.empresa" autocomplete="organization"
                   placeholder="Bodega El Sol" required
                   class="w-full border border-borde pl-11 pr-4 py-3 font-body text-body text-texto
                          bg-white placeholder:text-texto-muted
                          focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
          </div>
        </div>

        <!-- RUC + Teléfono -->
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-ruc">RUC *</label>
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                   width="14" height="14" viewBox="0 0 14 14" fill="none">
                <rect x="2" y="1" width="10" height="12" stroke="currentColor" stroke-width="1.2"/>
                <line x1="5" y1="4" x2="9" y2="4" stroke="currentColor" stroke-width="1"/>
              </svg>
              <input id="reg-ruc" v-model="form.ruc" placeholder="12345678901" required
                     class="w-full border border-borde pl-9 pr-3 py-3 font-body text-body text-texto
                            bg-white placeholder:text-texto-muted
                            focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
            </div>
          </div>
          <div>
            <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-tel">Teléfono</label>
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                   width="14" height="14" viewBox="0 0 14 14" fill="none">
                <rect x="4" y="1" width="6" height="12" rx="1" stroke="currentColor" stroke-width="1.2"/>
              </svg>
              <input id="reg-tel" v-model="form.telefono" type="tel" autocomplete="tel"
                     placeholder="999888777"
                     class="w-full border border-borde pl-9 pr-3 py-3 font-body text-body text-texto
                            bg-white placeholder:text-texto-muted
                            focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
            </div>
          </div>
        </div>

        <!-- Email -->
        <div>
          <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-email">Email *</label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="14" viewBox="0 0 16 14" fill="none">
              <rect x="0.5" y="1.5" width="15" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M0.5 1.5L8 7.5L15.5 1.5" stroke="currentColor" stroke-width="1.2"
                    stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="reg-email" v-model="form.email" type="email" autocomplete="email"
                   placeholder="tu@email.com" required
                   class="w-full border border-borde pl-11 pr-4 py-3 font-body text-body text-texto
                          bg-white placeholder:text-texto-muted
                          focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
          </div>
        </div>

        <!-- Contraseña -->
        <div>
          <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-pass">Contraseña *</label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7"
                    stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <circle cx="8" cy="11" r="1" fill="currentColor"/>
            </svg>
            <input id="reg-pass" v-model="form.password"
                   :type="mostrarPassword ? 'text' : 'password'"
                   autocomplete="new-password"
                   placeholder="Mínimo 6 caracteres" required
                   class="w-full border border-borde pl-11 pr-12 py-3 font-body text-body text-texto
                          bg-white placeholder:text-texto-muted
                          focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
            <button type="button"
                    class="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-texto-muted hover:text-texto transition-colors"
                    @click="mostrarPassword = !mostrarPassword"
                    :aria-label="mostrarPassword ? 'Ocultar' : 'Mostrar'">
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

        <!-- Confirmar -->
        <div>
          <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2" for="reg-pass2">Confirmar *</label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7"
                    stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <rect x="7" y="10" width="2" height="2" fill="currentColor"/>
            </svg>
            <input id="reg-pass2" v-model="form.confirmPassword" type="password"
                   autocomplete="new-password"
                   placeholder="Repetí tu contraseña" required
                   class="w-full border border-borde pl-11 pr-4 py-3 font-body text-body text-texto
                          bg-white placeholder:text-texto-muted
                          focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]" />
          </div>
        </div>

        <!-- Mensajes -->
        <p v-if="errorMsg" class="font-mono text-[11px] text-error uppercase tracking-[0.1em] text-center font-bold"
           role="alert">{{ errorMsg }}</p>

        <!-- Botón -->
        <button type="submit" :disabled="loading"
                class="w-full bg-texto text-white font-display text-heading uppercase tracking-[0.05em]
                       py-4 hover:bg-[#D4A017] hover:text-black
                       transition-colors duration-200
                       disabled:opacity-50 disabled:cursor-not-allowed
                       min-h-[56px] flex items-center justify-center gap-3">
          <span v-if="loading" class="w-5 h-5 border-2 border-white border-t-transparent animate-spin" />
          <span v-else>Crear Cuenta Gratis</span>
        </button>
      </form>

      <!-- Link login -->
      <p class="mt-6 text-center font-mono text-[10px] text-texto-muted uppercase tracking-[0.1em]">
        ¿Ya tenés cuenta?
        <NuxtLink to="/auth/login" class="text-texto hover:text-[#D4A017] font-bold ml-1 transition-colors">
          Iniciá Sesión
        </NuxtLink>
      </p>

    </div>
  </div>
</template>
