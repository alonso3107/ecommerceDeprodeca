<script setup lang="ts">
import { ref } from "vue"

definePageMeta({ layout: "auth" })

const form = ref({ nombre: "", empresa: "", ruc: "", telefono: "", email: "", password: "", confirmPassword: "" })
const loading = ref(false)
const errorMsg = ref("")
const successMsg = ref("")

function validar(): string | null {
  const f = form.value
  if (!f.nombre || !f.empresa || !f.ruc || !f.email || !f.password) return "Completá todos los campos obligatorios (*)"
  if (f.password.length < 6) return "La contraseña debe tener al menos 6 caracteres"
  if (f.password !== f.confirmPassword) return "Las contraseñas no coinciden"
  if (f.ruc.length !== 11) return "El RUC debe tener 11 dígitos"
  return null
}

async function handleRegistro() {
  const err = validar()
  if (err) { errorMsg.value = err; return }
  loading.value = true
  try {
    const config = useRuntimeConfig()
    const res = await $fetch<{ success: boolean; data: any; message: string }>(`${config.public.apiBase}/auth/registro`, {
      method: "POST",
      body: { email: form.value.email, password: form.value.password, nombre: form.value.nombre, empresa: form.value.empresa, ruc: form.value.ruc, telefono: form.value.telefono },
    })
    if (res.success) {
      successMsg.value = "¡Cuenta creada! Redirigiendo..."
      if (import.meta.client) {
        localStorage.setItem("deprodeca_token", res.data.token)
        localStorage.setItem("deprodeca_user", JSON.stringify({ id: res.data.usuario_id, email: res.data.email, nombre: res.data.nombre, rol: res.data.rol }))
      }
      setTimeout(() => navigateTo("/"), 1500)
    }
  } catch (e: any) {
    errorMsg.value = e?.data?.message || "Error al crear la cuenta"
  } finally { loading.value = false }
}
</script>

<template>
  <div class="w-full max-w-[440px] mx-auto">
    <div class="text-center mb-10">
      <NuxtLink to="/" class="inline-block">
        <span class="font-display text-display-md text-texto leading-none">depro<span class="text-texto-muted font-normal">deca</span></span>
      </NuxtLink>
      <p class="mt-2 font-body text-small text-texto-muted">Creá tu cuenta de proveedor</p>
    </div>

    <div class="bg-white rounded-2xl shadow-xs p-8">
      <form @submit.prevent="handleRegistro" class="space-y-4">
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Nombre Completo *</label>
          <input v-model="form.nombre" placeholder="Tu nombre" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Bodega / Empresa *</label>
          <input v-model="form.empresa" placeholder="Bodega El Sol" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>
        <div class="grid grid-cols-2 gap-3">
          <div>
            <label class="block font-body text-small font-semibold text-texto mb-2">RUC *</label>
            <input v-model="form.ruc" placeholder="12345678901" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
          </div>
          <div>
            <label class="block font-body text-small font-semibold text-texto mb-2">Teléfono</label>
            <input v-model="form.telefono" type="tel" placeholder="999888777" class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
          </div>
        </div>
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Email *</label>
          <input v-model="form.email" type="email" placeholder="tu@email.com" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Contraseña *</label>
          <input v-model="form.password" type="password" placeholder="Mínimo 6 caracteres" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Confirmar Contraseña *</label>
          <input v-model="form.confirmPassword" type="password" placeholder="Repetí tu contraseña" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>

        <p v-if="errorMsg" class="text-small text-error font-medium text-center" role="alert">{{ errorMsg }}</p>
        <p v-if="successMsg" class="text-small text-exito font-medium text-center" role="status">{{ successMsg }}</p>

        <Button type="submit" variant="primary" size="lg" :full-width="true" :loading="loading" :disabled="!!successMsg">
          Crear Cuenta Gratis
        </Button>
      </form>

      <p class="mt-6 text-center font-body text-small text-texto-muted">
        ¿Ya tenés cuenta?
        <NuxtLink to="/auth/login" class="text-texto hover:underline font-semibold ml-1">Iniciá Sesión</NuxtLink>
      </p>
    </div>
  </div>
</template>
