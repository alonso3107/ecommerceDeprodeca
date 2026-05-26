     1|<script setup lang="ts">
     2|import { ref } from "vue"
     3|
     4|definePageMeta({
     5|  layout: "auth",
     6|  middleware: "guest",
     7|})
     8|
     9|const email = ref("")
    10|const password = ref("")
    11|const loading = ref(false)
    12|const errorMsg = ref("")
    13|
    14|async function handleLogin() {
    15|  errorMsg.value = ""
    16|  if (!email.value || !password.value) {
    17|    errorMsg.value = "Todos los campos son obligatorios"
    18|    return
    19|  }
    20|
    21|  loading.value = true
    22|  try {
    23|    const config = useRuntimeConfig()
    24|    const res = await $fetch<{
    25|      success: boolean
    26|      data: { token: string; usuario_id: number; email: string; nombre: string; rol: string }
    27|      message: string
    28|    }>(`${config.public.apiBase}/auth/login`, {
    29|      method: "POST",
    30|      body: { email: email.value, password: password.value },
    31|    })
    32|
    33|    if (res.success) {
    34|      // Guardar token en localStorage
    35|      if (import.meta.client) {
    36|        localStorage.setItem("deprodeca_token", res.data.token)
    37|        localStorage.setItem("deprodeca_user", JSON.stringify({
    38|          id: res.data.usuario_id,
    39|          email: res.data.email,
    40|          nombre: res.data.nombre,
    41|          rol: res.data.rol,
    42|        }))
    43|      }
    44|      await navigateTo("/")
    45|    } else {
    46|      errorMsg.value = res.message || "Error al iniciar sesión"
    47|    }
    48|  } catch (e: any) {
    49|    errorMsg.value = e?.data?.message || "Error de conexión con el servidor"
    50|  } finally {
    51|    loading.value = false
    52|  }
    53|}
    54|</script>
    55|
    56|<template>
    57|  <div class="w-full">
    58|    <!-- Logo -->
    59|    <div class="text-center mb-8">
    60|      <NuxtLink to="/" class="inline-block">
    61|        <span class="font-display text-display-md text-text-primary leading-none">
    62|          DEPRO<span class="text-brand-primary">DECA</span>
    63|        </span>
    64|      </NuxtLink>
    65|      <p class="mt-2 font-body text-small text-text-muted">
    66|        Ingresa a tu cuenta de proveedor
    67|      </p>
    68|    </div>
    69|
    70|    <!-- Card del formulario -->
    71|    <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-8">
    72|      <form @submit.prevent="handleLogin" class="space-y-5">
    73|        <!-- Email -->
    74|        <InputField
    75|          v-model="email"
    76|          label="Correo Electrónico"
    77|          type="email"
    78|          placeholder="tu@email.com"
    79|          :required="true"
    80|          icon="user"
    81|        />
    82|
    83|        <!-- Password -->
    84|        <InputField
    85|          v-model="password"
    86|          label="Contraseña"
    87|          type="password"
    88|          placeholder="••••••••"
    89|          :required="true"
    90|          icon="lock"
    91|        />
    92|
    93|        <!-- Error -->
    94|        <p
    95|          v-if="errorMsg"
    96|          class="text-small text-status-error font-medium text-center"
    97|          role="alert"
    98|        >
    99|          {{ errorMsg }}
   100|        </p>
   101|
   102|        <!-- Submit -->
   103|        <Button
   104|          type="submit"
   105|          variant="primary"
   106|          size="lg"
   107|          :full-width="true"
   108|          :loading="loading"
   109|        >
   110|          Iniciar Sesión
   111|        </Button>
   112|      </form>
   113|
   114|      <!-- Divider -->
   115|      <div class="relative my-6">
   116|        <div class="absolute inset-0 flex items-center">
   117|          <div class="w-full border-t border-border-base" />
   118|        </div>
   119|        <div class="relative flex justify-center">
   120|          <span class="bg-white px-4 font-body text-caption text-text-muted">
   121|            ¿Nuevo en DEPRODECA?
   122|          </span>
   123|        </div>
   124|      </div>
   125|
   126|      <!-- Link a registro -->
   127|      <NuxtLink
   128|        to="/auth/registro"
   129|        class="block text-center w-full py-3 rounded-lg border-2 border-border-base font-body text-body font-semibold text-text-primary hover:border-brand-primary hover:text-brand-primary transition-all duration-300"
   130|      >
   131|        Crear una Cuenta Gratis
   132|      </NuxtLink>
   133|    </div>
   134|  </div>
   135|</template>
   136|