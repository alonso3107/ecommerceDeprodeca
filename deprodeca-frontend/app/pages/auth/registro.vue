     1|<script setup lang="ts">
     2|import { ref } from "vue"
     3|
     4|definePageMeta({
     5|  layout: "auth",
     6|  middleware: "guest",
     7|})
     8|
     9|const nombre = ref("")
    10|const empresa = ref("")
    11|const ruc = ref("")
    12|const telefono = ref("")
    13|const email = ref("")
    14|const password = ref("")
    15|const confirmPassword = ref("")
    16|const loading = ref(false)
    17|const errorMsg = ref("")
    18|const successMsg = ref("")
    19|
    20|function validar() {
    21|  errorMsg.value = ""
    22|  if (!nombre.value || !empresa.value || !ruc.value || !email.value || !password.value) {
    23|    return "Todos los campos obligatorios (*) deben completarse"
    24|  }
    25|  if (password.value.length < 6) {
    26|    return "La contraseña debe tener al menos 6 caracteres"
    27|  }
    28|  if (password.value !== confirmPassword.value) {
    29|    return "Las contraseñas no coinciden"
    30|  }
    31|  if (ruc.value.length !== 11) {
    32|    return "El RUC debe tener 11 dígitos"
    33|  }
    34|  return null
    35|}
    36|
    37|async function handleRegistro() {
    38|  const err = validar()
    39|  if (err) {
    40|    errorMsg.value = err
    41|    return
    42|  }
    43|
    44|  loading.value = true
    45|  try {
    46|    const config = useRuntimeConfig()
    47|    const res = await $fetch<{
    48|      success: boolean
    49|      data: { token: string; usuario_id: number; email: string; nombre: string; rol: string }
    50|      message: string
    51|    }>(`${config.public.apiBase}/auth/registro`, {
    52|      method: "POST",
    53|      body: {
    54|        email: email.value,
    55|        password: password.value,
    56|        nombre: nombre.value,
    57|        empresa: empresa.value,
    58|        ruc: ruc.value,
    59|        telefono: telefono.value,
    60|      },
    61|    })
    62|
    63|    if (res.success) {
    64|      successMsg.value = "¡Cuenta creada! Redirigiendo..."
    65|      if (import.meta.client) {
    66|        localStorage.setItem("deprodeca_token", res.data.token)
    67|        localStorage.setItem("deprodeca_user", JSON.stringify({
    68|          id: res.data.usuario_id,
    69|          email: res.data.email,
    70|          nombre: res.data.nombre,
    71|          rol: res.data.rol,
    72|        }))
    73|      }
    74|      setTimeout(() => navigateTo("/"), 1500)
    75|    }
    76|  } catch (e: any) {
    77|    errorMsg.value = e?.data?.message || "Error al crear la cuenta"
    78|  } finally {
    79|    loading.value = false
    80|  }
    81|}
    82|</script>
    83|
    84|<template>
    85|  <div class="w-full">
    86|    <!-- Logo -->
    87|    <div class="text-center mb-8">
    88|      <NuxtLink to="/" class="inline-block">
    89|        <span class="font-display text-display-md text-text-primary leading-none">
    90|          DEPRO<span class="text-brand-primary">DECA</span>
    91|        </span>
    92|      </NuxtLink>
    93|      <p class="mt-2 font-body text-small text-text-muted">
    94|        Crea tu cuenta de proveedor gratis
    95|      </p>
    96|    </div>
    97|
    98|    <!-- Card del formulario -->
    99|    <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-8">
   100|      <form @submit.prevent="handleRegistro" class="space-y-4">
   101|        <!-- Datos personales -->
   102|        <InputField
   103|          v-model="nombre"
   104|          label="Nombre Completo"
   105|          placeholder="Tu nombre"
   106|          :required="true"
   107|          icon="user"
   108|        />
   109|
   110|        <InputField
   111|          v-model="empresa"
   112|          label="Nombre de la Bodega / Empresa"
   113|          placeholder="Bodega El Sol"
   114|          :required="true"
   115|        />
   116|
   117|        <div class="grid grid-cols-2 gap-3">
   118|          <InputField
   119|            v-model="ruc"
   120|            label="RUC"
   121|            placeholder="12345678901"
   122|            :required="true"
   123|          />
   124|          <InputField
   125|            v-model="telefono"
   126|            label="Teléfono"
   127|            type="tel"
   128|            placeholder="999888777"
   129|          />
   130|        </div>
   131|
   132|        <InputField
   133|          v-model="email"
   134|          label="Correo Electrónico"
   135|          type="email"
   136|          placeholder="tu@email.com"
   137|          :required="true"
   138|        />
   139|
   140|        <InputField
   141|          v-model="password"
   142|          label="Contraseña"
   143|          type="password"
   144|          placeholder="Mínimo 6 caracteres"
   145|          :required="true"
   146|          icon="lock"
   147|        />
   148|
   149|        <InputField
   150|          v-model="confirmPassword"
   151|          label="Confirmar Contraseña"
   152|          type="password"
   153|          placeholder="Repite tu contraseña"
   154|          :required="true"
   155|          icon="lock"
   156|        />
   157|
   158|        <!-- Error -->
   159|        <p
   160|          v-if="errorMsg"
   161|          class="text-small text-status-error font-medium text-center"
   162|          role="alert"
   163|        >
   164|          {{ errorMsg }}
   165|        </p>
   166|
   167|        <!-- Success -->
   168|        <p
   169|          v-if="successMsg"
   170|          class="text-small text-status-success font-medium text-center"
   171|          role="status"
   172|        >
   173|          {{ successMsg }}
   174|        </p>
   175|
   176|        <!-- Submit -->
   177|        <Button
   178|          type="submit"
   179|          variant="primary"
   180|          size="lg"
   181|          :full-width="true"
   182|          :loading="loading"
   183|          :disabled="!!successMsg"
   184|        >
   185|          Crear Cuenta Gratis
   186|        </Button>
   187|      </form>
   188|
   189|      <!-- Link a login -->
   190|      <p class="mt-6 text-center font-body text-small text-text-muted">
   191|        ¿Ya tienes cuenta?
   192|        <NuxtLink to="/auth/login" class="text-brand-secondary hover:underline font-semibold ml-1">
   193|          Inicia Sesión
   194|        </NuxtLink>
   195|      </p>
   196|    </div>
   197|  </div>
   198|</template>
   199|