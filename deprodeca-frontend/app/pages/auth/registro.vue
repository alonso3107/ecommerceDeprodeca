<!--
  ═══════════════════════════════════════════════════════════════
  auth/registro.vue — Crear Cuenta · DEPRODECA
  Validación estricta: RUC exacto 11 dígitos, contraseña ≥6,
  feedback en tiempo real por campo. El usuario no juega.
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

// ─── Validaciones por campo (tiempo real) ─────────────────
const errores = computed(() => {
  const e: Record<string, string | null> = {
    nombre: null, empresa: null, ruc: null, email: null, password: null, confirmPassword: null,
  }
  const f = form.value

  if (f.nombre && f.nombre.trim().length < 3)
    e.nombre = "Mínimo 3 caracteres"

  if (f.empresa && f.empresa.trim().length < 2)
    e.empresa = "Mínimo 2 caracteres"

  if (f.ruc && !/^\d+$/.test(f.ruc))
    e.ruc = "Solo dígitos"
  else if (f.ruc && f.ruc.length !== 11)
    e.ruc = `Exacto 11 dígitos (${f.ruc.length}/11)`

  if (f.email && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(f.email))
    e.email = "Email inválido"

  if (f.password && f.password.length < 6)
    e.password = "Mínimo 6 caracteres"

  if (f.confirmPassword && f.password !== f.confirmPassword)
    e.confirmPassword = "No coincide"

  return e
})

const campoValido = (campo: string) => {
  const f = form.value
  const val = (f as any)[campo]
  if (!val) return null // sin tocar
  return errores.value[campo] === null // true = ✅, false = ❌
}

// ─── Validación final ─────────────────────────────────────
function validar(): string | null {
  const f = form.value
  if (!f.nombre || !f.empresa || !f.ruc || !f.email || !f.password)
    return "Completá todos los campos obligatorios (*)"
  if (f.nombre.trim().length < 3) return "Nombre: mínimo 3 caracteres"
  if (f.empresa.trim().length < 2) return "Empresa: mínimo 2 caracteres"
  if (!/^\d{11}$/.test(f.ruc)) return "El RUC debe tener exactamente 11 dígitos"
  if (f.password.length < 6) return "La contraseña debe tener al menos 6 caracteres"
  if (f.password !== f.confirmPassword) return "Las contraseñas no coinciden"
  return null
}

// ─── Solo dígitos para RUC y teléfono ─────────────────────
function soloDigitos(event: KeyboardEvent) {
  if (!/[\d]/.test(event.key) && event.key !== "Backspace" && event.key !== "Tab" && event.key !== "ArrowLeft" && event.key !== "ArrowRight") {
    event.preventDefault()
  }
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
        body: { email: form.value.email, password: form.value.password,
                nombre: form.value.nombre, empresa: form.value.empresa,
                ruc: form.value.ruc, telefono: form.value.telefono },
      })
    if (res.success) {
      successMsg.value = "¡Cuenta creada! Redirigiendo..."
      if (import.meta.client) {
        localStorage.setItem("deprodeca_token", res.data.token)
        localStorage.setItem("deprodeca_user", JSON.stringify({
          id: res.data.usuario_id, email: res.data.email,
          nombre: res.data.nombre, rol: res.data.rol,
        }))
        window.dispatchEvent(new Event("auth-cambiado"))
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
    <div class="mb-8">
      <p class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.3em] mb-3">─── Registro</p>
      <h1 class="font-display text-display-lg text-negro uppercase leading-[0.95]">
        Crear<br />Cuenta<span class="text-dorado">.</span>
      </h1>
      <p class="mt-3 font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.15em]">
        Registrate como proveedor
      </p>
    </div>

    <div class="border border-stone bg-blanco p-8">

      <!-- Éxito -->
      <div v-if="successMsg" class="text-center py-8 page-enter">
        <svg class="mx-auto mb-4 text-exito" width="48" height="48" viewBox="0 0 48 48" fill="none">
          <path d="M24 4L44 24L24 44L4 24L24 4Z" stroke="currentColor" stroke-width="2"/>
          <path d="M18 25L22 30L32 20" stroke="currentColor" stroke-width="3" stroke-linecap="square" stroke-linejoin="miter"/>
        </svg>
        <p class="font-display text-display-md text-negro uppercase mb-2">¡Cuenta Creada<span class="text-exito">!</span></p>
        <p class="font-mono text-[10px] text-stone-oscuro uppercase tracking-wider">Redirigiendo...</p>
      </div>

      <!-- Formulario -->
      <form v-else @submit.prevent="handleRegistro" class="space-y-4" autocomplete="on" novalidate>

        <!-- ═══ NOMBRE ═══ -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="reg-nombre">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Nombre Completo
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('nombre') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('nombre') === false" class="font-mono text-[10px] text-error">{{ errores.nombre }}</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <circle cx="8" cy="5" r="2.5" stroke="currentColor" stroke-width="1.2"/>
              <path d="M3 14C3 11.2 5.2 9 8 9C10.8 9 13 11.2 13 14" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
            </svg>
            <input id="reg-nombre" v-model="form.nombre" autocomplete="name"
                   placeholder="Tu nombre completo" required minlength="3"
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('nombre') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
          </div>
        </div>

        <!-- ═══ EMPRESA ═══ -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="reg-empresa">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Bodega / Empresa
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('empresa') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('empresa') === false" class="font-mono text-[10px] text-error">{{ errores.empresa }}</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="4" y="4" width="8" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M2.5 4L8 1L13.5 4" stroke="currentColor" stroke-width="1.2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="reg-empresa" v-model="form.empresa" autocomplete="organization"
                   placeholder="Bodega El Sol" required minlength="2"
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('empresa') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
          </div>
        </div>

        <!-- ═══ RUC + TELÉFONO ═══ -->
        <div class="grid grid-cols-2 gap-3">
          <!-- RUC · BLOQUEADO a 11 dígitos exactos -->
          <div>
            <label class="flex justify-between items-baseline mb-2" for="reg-ruc">
              <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">RUC
                <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                  <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
                </svg>
              </span>
              <span v-if="campoValido('ruc') === true" class="font-mono text-[10px] text-exito">✓</span>
              <span v-else-if="campoValido('ruc') === false" class="font-mono text-[10px] text-error">{{ errores.ruc }}</span>
              <span v-else class="font-mono text-[9px] text-stone-oscuro/50 tracking-wider">11 dígitos</span>
            </label>
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                   width="14" height="14" viewBox="0 0 14 14" fill="none">
                <rect x="2" y="1" width="10" height="12" stroke="currentColor" stroke-width="1.2"/>
                <line x1="5" y1="4" x2="9" y2="4" stroke="currentColor" stroke-width="1"/>
              </svg>
              <input id="reg-ruc" v-model="form.ruc"
                     placeholder="12345678901" required
                     inputmode="numeric" maxlength="11" minlength="11"
                     pattern="\d{11}"
                     @keydown="soloDigitos"
                     :class="['w-full border pl-9 pr-3 py-3 font-mono text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px] tracking-[0.1em]',
                               campoValido('ruc') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
            </div>
            <!-- Contador de dígitos -->
            <p class="mt-1 font-mono text-[9px] text-stone-oscuro/50 text-right tracking-wider">
              {{ form.ruc.length }}/11
            </p>
          </div>

          <!-- Teléfono -->
          <div>
            <label class="flex justify-between items-baseline mb-2" for="reg-tel">
              <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em]">Teléfono</span>
            </label>
            <div class="relative">
              <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                   width="14" height="14" viewBox="0 0 14 14" fill="none">
                <rect x="4" y="1" width="6" height="12" rx="1" stroke="currentColor" stroke-width="1.2"/>
              </svg>
              <input id="reg-tel" v-model="form.telefono" type="tel" autocomplete="tel"
                     placeholder="999888777" inputmode="numeric" maxlength="9"
                     @keydown="soloDigitos"
                     class="w-full border border-stone pl-9 pr-3 py-3 font-mono text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:border-dorado focus:outline-none transition-colors duration-200 min-h-[48px] tracking-[0.1em]" />
            </div>
          </div>
        </div>

        <!-- ═══ EMAIL ═══ -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="reg-email">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Email
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('email') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('email') === false" class="font-mono text-[10px] text-error">{{ errores.email }}</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="14" viewBox="0 0 16 14" fill="none">
              <rect x="0.5" y="1.5" width="15" height="11" stroke="currentColor" stroke-width="1.2"/>
              <path d="M0.5 1.5L8 7.5L15.5 1.5" stroke="currentColor" stroke-width="1.2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <input id="reg-email" v-model="form.email" type="email" autocomplete="email"
                   placeholder="tu@email.com" required
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('email') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
          </div>
        </div>

        <!-- ═══ CONTRASEÑA ═══ -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="reg-pass">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Contraseña
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('password') === true" class="font-mono text-[10px] text-exito">✓</span>
            <span v-else-if="campoValido('password') === false" class="font-mono text-[10px] text-error">{{ errores.password }}</span>
            <span v-else class="font-mono text-[9px] text-stone-oscuro/50 tracking-wider">mín 6</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <circle cx="8" cy="11" r="1" fill="currentColor"/>
            </svg>
            <input id="reg-pass" v-model="form.password"
                   :type="mostrarPassword ? 'text' : 'password'"
                   autocomplete="new-password" placeholder="Mínimo 6 caracteres" required minlength="6"
                   :class="['w-full border pl-11 pr-12 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('password') === false ? 'border-error focus:border-error' : 'border-stone focus:border-dorado']" />
            <button type="button"
                    class="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-stone-oscuro hover:text-negro transition-colors duration-200"
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

        <!-- ═══ CONFIRMAR CONTRASEÑA ═══ -->
        <div>
          <label class="flex justify-between items-baseline mb-2" for="reg-pass2">
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.2em] inline-flex items-center gap-1">Confirmar
              <svg width="6" height="6" viewBox="0 0 6 6" class="text-dorado inline-block" aria-hidden="true">
                <rect x="3" y="0" width="4.2" height="4.2" transform="rotate(45 3 0)" fill="currentColor"/>
              </svg>
            </span>
            <span v-if="campoValido('confirmPassword') === true" class="font-mono text-[10px] text-exito">✓ Iguales</span>
            <span v-else-if="campoValido('confirmPassword') === false" class="font-mono text-[10px] text-error">✗ No coincide</span>
          </label>
          <div class="relative">
            <svg class="absolute left-4 top-1/2 -translate-y-1/2 text-stone-oscuro pointer-events-none"
                 width="16" height="16" viewBox="0 0 16 16" fill="none">
              <rect x="2.5" y="7" width="11" height="8" stroke="currentColor" stroke-width="1.2"/>
              <path d="M4.5 7V5C4.5 3 6 1.5 8 1.5C10 1.5 11.5 3 11.5 5V7" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
              <rect x="7" y="10" width="2" height="2" fill="currentColor"/>
            </svg>
            <input id="reg-pass2" v-model="form.confirmPassword" type="password"
                   autocomplete="new-password" placeholder="Repetí tu contraseña" required
                   :class="['w-full border pl-11 pr-4 py-3 font-body text-body text-negro bg-blanco placeholder:text-stone-oscuro hover:bg-crema/30 focus:bg-blanco focus:outline-none transition-colors duration-200 min-h-[48px]',
                             campoValido('confirmPassword') === false ? 'border-error focus:border-error' : campoValido('confirmPassword') === true ? 'border-exito focus:border-exito' : 'border-stone focus:border-dorado']" />
          </div>
        </div>

        <!-- Mensajes -->
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
                       py-4 hover:bg-dorado hover:text-negro
                       transition-colors duration-200
                       disabled:opacity-50 disabled:cursor-not-allowed
                       min-h-[56px] flex items-center justify-center gap-3">
          <span v-if="loading" class="w-5 h-5 border-2 border-blanco border-t-transparent animate-spin" />
          <template v-else>
            <span>Crear Cuenta Gratis</span>
            <svg width="16" height="12" viewBox="0 0 16 12" fill="none" aria-hidden="true">
              <path d="M0 6H14M14 6L9.5 1.5M14 6L9.5 10.5" stroke="currentColor" stroke-width="2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
          </template>
        </button>
      </form>

      <p class="mt-6 text-center font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.1em]">
        ¿Ya tenés cuenta?
        <NuxtLink to="/auth/login" class="text-stone-oscuro hover:text-dorado font-bold ml-1 transition-colors duration-200 inline-flex items-center gap-1">
          Iniciá Sesión
          <svg width="8" height="8" viewBox="0 0 8 8" class="text-dorado" aria-hidden="true">
            <rect x="4" y="0" width="5.6" height="5.6" transform="rotate(45 4 0)" fill="currentColor"/>
          </svg>
        </NuxtLink>
      </p>

    </div>
  </div>
</template>
