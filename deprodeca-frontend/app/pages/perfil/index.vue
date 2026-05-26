     1|<script setup lang="ts">
     2|definePageMeta({
     3|  layout: "default",
     4|  middleware: "auth",
     5|})
     6|
     7|const config = useRuntimeConfig()
     8|const perfil = ref<any>(null)
     9|const editando = ref(false)
    10|const loading = ref(true)
    11|const saving = ref(false)
    12|const toastMsg = ref("")
    13|const toastType = ref<"success" | "error">("success")
    14|const showToast = ref(false)
    15|
    16|// Form editable
    17|const form = ref({ nombre: "", empresa: "", telefono: "" })
    18|
    19|onMounted(async () => {
    20|  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
    21|  if (!token) return
    22|
    23|  try {
    24|    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/usuarios/perfil`, {
    25|      headers: { Authorization: `Bearer ${token}` },
    26|    })
    27|    if (res.success) {
    28|      perfil.value = res.data
    29|      form.value = { nombre: res.data.nombre, empresa: res.data.empresa, telefono: res.data.telefono }
    30|    }
    31|  } catch (_) {} finally {
    32|    loading.value = false
    33|  }
    34|})
    35|
    36|async function guardarPerfil() {
    37|  saving.value = true
    38|  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
    39|
    40|  try {
    41|    const res = await $fetch<{ success: boolean; message: string }>(`${config.public.apiBase}/usuarios/perfil`, {
    42|      method: "PUT",
    43|      headers: { Authorization: `Bearer ${token}`, "Content-Type": "application/json" },
    44|      body: form.value,
    45|    })
    46|    if (res.success) {
    47|      perfil.value = { ...perfil.value, ...form.value }
    48|      toastMsg.value = "Perfil actualizado"
    49|      toastType.value = "success"
    50|    }
    51|  } catch (e: any) {
    52|    toastMsg.value = e?.data?.message || "Error al actualizar"
    53|    toastType.value = "error"
    54|  } finally {
    55|    saving.value = false
    56|    editando.value = false
    57|    showToast.value = true
    58|    setTimeout(() => showToast.value = false, 3000)
    59|  }
    60|}
    61|</script>
    62|
    63|<template>
    64|  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    65|    <h1 class="font-display text-display-lg text-text-primary uppercase mb-8">Mi Perfil</h1>
    66|
    67|    <div v-if="loading" class="flex justify-center py-20">
    68|      <Spinner size="lg" label="Cargando perfil..." />
    69|    </div>
    70|
    71|    <template v-else-if="perfil">
    72|      <div class="grid grid-cols-1 md:grid-cols-[300px_1fr] gap-8">
    73|        <!-- Sidebar -->
    74|        <div class="space-y-3">
    75|          <NuxtLink
    76|            to="/perfil"
    77|            class="flex items-center gap-2 px-4 py-3 rounded-lg bg-brand-primary/20 text-text-primary font-body text-small font-semibold"
    78|          >
    79|            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21v-2a4 4 0 00-4-4H9a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
    80|            Datos Personales
    81|          </NuxtLink>
    82|          <NuxtLink
    83|            to="/perfil/rangos"
    84|            class="flex items-center gap-2 px-4 py-3 rounded-lg text-text-muted font-body text-small hover:bg-surface hover:text-text-primary transition-colors"
    85|          >
    86|            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M12 15l-4.24-4.24a6 6 0 0112.48 0L20 15"/></svg>
    87|            Mi Rango
    88|          </NuxtLink>
    89|          <NuxtLink
    90|            to="/pedidos"
    91|            class="flex items-center gap-2 px-4 py-3 rounded-lg text-text-muted font-body text-small hover:bg-surface hover:text-text-primary transition-colors"
    92|          >
    93|            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M6 2 3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4Z"/><path d="M3 6h18"/><path d="M16 10a4 4 0 01-8 0"/></svg>
    94|            Mis Pedidos
    95|          </NuxtLink>
    96|        </div>
    97|
    98|        <!-- Contenido -->
    99|        <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-8">
   100|          <div class="flex items-center justify-between mb-6">
   101|            <h2 class="font-body text-heading font-bold text-text-primary">Datos Personales</h2>
   102|            <Button
   103|              v-if="!editando"
   104|              variant="outline"
   105|              size="sm"
   106|              @click="editando = true"
   107|            >
   108|              Editar
   109|            </Button>
   110|          </div>
   111|
   112|          <div v-if="!editando" class="space-y-4">
   113|            <div>
   114|              <p class="font-body text-caption text-text-muted">Nombre</p>
   115|              <p class="font-body text-body font-semibold text-text-primary">{{ perfil.nombre }}</p>
   116|            </div>
   117|            <div>
   118|              <p class="font-body text-caption text-text-muted">Empresa</p>
   119|              <p class="font-body text-body font-semibold text-text-primary">{{ perfil.empresa }}</p>
   120|            </div>
   121|            <div>
   122|              <p class="font-body text-caption text-text-muted">RUC</p>
   123|              <p class="font-body text-body font-semibold text-text-primary">{{ perfil.ruc }}</p>
   124|            </div>
   125|            <div>
   126|              <p class="font-body text-caption text-text-muted">Email</p>
   127|              <p class="font-body text-body font-semibold text-text-primary">{{ perfil.email }}</p>
   128|            </div>
   129|            <div>
   130|              <p class="font-body text-caption text-text-muted">Teléfono</p>
   131|              <p class="font-body text-body font-semibold text-text-primary">{{ perfil.telefono || "No registrado" }}</p>
   132|            </div>
   133|          </div>
   134|
   135|          <form v-else @submit.prevent="guardarPerfil" class="space-y-4">
   136|            <InputField v-model="form.nombre" label="Nombre" :required="true" />
   137|            <InputField v-model="form.empresa" label="Empresa" :required="true" />
   138|            <InputField v-model="form.telefono" label="Teléfono" type="tel" />
   139|            <div class="flex gap-3 pt-2">
   140|              <Button type="submit" variant="primary" :loading="saving">Guardar</Button>
   141|              <Button variant="ghost" @click="editando = false">Cancelar</Button>
   142|            </div>
   143|          </form>
   144|        </div>
   145|      </div>
   146|    </template>
   147|
   148|    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
   149|  </div>
   150|</template>
   151|