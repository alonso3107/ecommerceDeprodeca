<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })
const config = useRuntimeConfig()
const perfil = ref<any>(null)
const editando = ref(false)
const loading = ref(true)
const saving = ref(false)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)
const form = ref({ nombre: "", empresa: "", telefono: "" })

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/usuarios/perfil`, { headers: { Authorization: `Bearer ${token}` } })
    if (res.success) { perfil.value = res.data; form.value = { nombre: res.data.nombre, empresa: res.data.empresa, telefono: res.data.telefono } }
  } catch (_) {} finally { loading.value = false }
})

async function guardarPerfil() {
  saving.value = true
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  try {
    await $fetch(`${config.public.apiBase}/usuarios/perfil`, { method: "PUT", headers: { Authorization: `Bearer ${token}`, "Content-Type": "application/json" }, body: form.value })
    perfil.value = { ...perfil.value, ...form.value }
    toastMsg.value = "Perfil actualizado"; toastType.value = "success"
  } catch (e: any) { toastMsg.value = e?.data?.message || "Error"; toastType.value = "error" }
  finally { saving.value = false; editando.value = false; showToast.value = true; setTimeout(() => showToast.value = false, 3000) }
}
</script>

<template>
  <div class="page-enter max-w-[960px] mx-auto px-6 md:px-8 py-12">
    <h1 class="font-display text-display-lg text-texto uppercase mb-10">Mi Perfil</h1>

    <div v-if="loading" class="flex justify-center py-20"><Spinner size="lg" /></div>

    <template v-else-if="perfil">
      <div class="grid grid-cols-1 md:grid-cols-[220px_1fr] gap-10">
        <div class="space-y-1">
          <NuxtLink to="/perfil" class="flex items-center gap-2 px-4 py-3 rounded-xl bg-texto text-white font-body text-small font-semibold">Datos Personales</NuxtLink>
          <NuxtLink to="/perfil/rangos" class="flex items-center gap-2 px-4 py-3 rounded-xl text-texto-muted font-body text-small hover:bg-fondo hover:text-texto transition-colors">Mi Rango</NuxtLink>
          <NuxtLink to="/pedidos" class="flex items-center gap-2 px-4 py-3 rounded-xl text-texto-muted font-body text-small hover:bg-fondo hover:text-texto transition-colors">Mis Pedidos</NuxtLink>
        </div>

        <div class="bg-white rounded-2xl shadow-xs p-8">
          <div class="flex items-center justify-between mb-8">
            <h2 class="font-body text-heading font-bold text-texto">Datos Personales</h2>
            <Button v-if="!editando" variant="outline" size="sm" @click="editando = true">Editar</Button>
          </div>

          <div v-if="!editando" class="space-y-5">
            <div><p class="font-body text-caption text-texto-muted uppercase tracking-wide">Nombre</p><p class="font-body text-body font-semibold text-texto mt-1">{{ perfil.nombre }}</p></div>
            <div><p class="font-body text-caption text-texto-muted uppercase tracking-wide">Empresa</p><p class="font-body text-body font-semibold text-texto mt-1">{{ perfil.empresa }}</p></div>
            <div><p class="font-body text-caption text-texto-muted uppercase tracking-wide">RUC</p><p class="font-body text-body font-semibold text-texto mt-1">{{ perfil.ruc }}</p></div>
            <div><p class="font-body text-caption text-texto-muted uppercase tracking-wide">Email</p><p class="font-body text-body font-semibold text-texto mt-1">{{ perfil.email }}</p></div>
            <div><p class="font-body text-caption text-texto-muted uppercase tracking-wide">Teléfono</p><p class="font-body text-body font-semibold text-texto mt-1">{{ perfil.telefono || "—" }}</p></div>
          </div>

          <form v-else @submit.prevent="guardarPerfil" class="space-y-5">
            <div><label class="block font-body text-small font-semibold text-texto mb-2">Nombre</label><input v-model="form.nombre" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all min-h-[48px]" /></div>
            <div><label class="block font-body text-small font-semibold text-texto mb-2">Empresa</label><input v-model="form.empresa" required class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all min-h-[48px]" /></div>
            <div><label class="block font-body text-small font-semibold text-texto mb-2">Teléfono</label><input v-model="form.telefono" type="tel" class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all min-h-[48px]" /></div>
            <div class="flex gap-3 pt-2"><Button type="submit" variant="primary" :loading="saving">Guardar</Button><Button variant="ghost" @click="editando = false">Cancelar</Button></div>
          </form>
        </div>
      </div>
    </template>

    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
  </div>
</template>
