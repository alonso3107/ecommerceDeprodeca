<script setup lang="ts">
import { ref } from "vue"

definePageMeta({ layout: "auth" })

const email = ref("")
const password = ref("")
const loading = ref(false)
const errorMsg = ref("")

async function handleLogin() {
  errorMsg.value = ""
  if (!email.value || !password.value) {
    errorMsg.value = "Todos los campos son obligatorios"
    return
  }
  loading.value = true
  try {
    const config = useRuntimeConfig()
    const res = await $fetch<{ success: boolean; data: any; message: string }>(`${config.public.apiBase}/auth/login`, {
      method: "POST",
      body: { email: email.value, password: password.value },
    })
    if (res.success) {
      if (import.meta.client) {
        localStorage.setItem("deprodeca_token", res.data.token)
        localStorage.setItem("deprodeca_user", JSON.stringify({ id: res.data.usuario_id, email: res.data.email, nombre: res.data.nombre, rol: res.data.rol }))
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
  <div class="w-full max-w-[400px] mx-auto">
    <div class="text-center mb-10">
      <NuxtLink to="/" class="inline-block">
        <span class="font-display text-display-md text-texto leading-none">depro<span class="text-texto-muted font-normal">deca</span></span>
      </NuxtLink>
      <p class="mt-2 font-body text-small text-texto-muted">Ingresá a tu cuenta</p>
    </div>

    <div class="bg-white rounded-2xl shadow-xs p-8">
      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Email</label>
          <input v-model="email" type="email" placeholder="tu@email.com" required
            class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>

        <div>
          <label class="block font-body text-small font-semibold text-texto mb-2">Contraseña</label>
          <input v-model="password" type="password" placeholder="••••••••" required
            class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[48px]" />
        </div>

        <p v-if="errorMsg" class="text-small text-error font-medium text-center" role="alert">{{ errorMsg }}</p>

        <Button type="submit" variant="primary" size="lg" :full-width="true" :loading="loading">
          Iniciar Sesión
        </Button>
      </form>

      <div class="relative my-6">
        <div class="absolute inset-0 flex items-center"><div class="w-full border-t border-borde" /></div>
        <div class="relative flex justify-center"><span class="bg-white px-4 font-body text-caption text-texto-muted">¿Nuevo en DEPRODECA?</span></div>
      </div>

      <NuxtLink to="/auth/registro"
        class="block text-center w-full py-3 rounded-xl border border-borde font-body text-body font-semibold text-texto hover:bg-fondo transition-all duration-300">
        Crear una Cuenta Gratis
      </NuxtLink>
    </div>
  </div>
</template>
