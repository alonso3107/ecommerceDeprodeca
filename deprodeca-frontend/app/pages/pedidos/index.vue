<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })
const config = useRuntimeConfig()
const pedidos = ref<any[]>([])
const loading = ref(true)

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/pedidos`, { headers: { Authorization: `Bearer ${token}` } })
    if (res.success) pedidos.value = res.data
  } catch (_) {} finally { loading.value = false }
})

function formatearPrecio(p: number) { return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(p) }
function formatearFecha(f: string) { return new Date(f).toLocaleDateString("es-PE", { day: "numeric", month: "short", year: "numeric" }) }
const badge: Record<string, string> = { pendiente: "bg-advertencia/10 text-advertencia", confirmado: "bg-texto/10 text-texto", enviado: "bg-texto/5 text-texto", entregado: "bg-exito/10 text-exito", cancelado: "bg-error/10 text-error" }
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <h1 class="font-display text-display-lg text-texto uppercase mb-10">Mis Pedidos</h1>

    <div v-if="loading" class="flex justify-center py-20"><Spinner size="lg" /></div>

    <div v-else-if="pedidos.length === 0" class="text-center py-20">
      <p class="font-display text-display-md text-texto-muted mb-4">Sin pedidos</p>
      <Button variant="primary" @click="navigateTo('/catalogo')">Ir al Catálogo</Button>
    </div>

    <div v-else class="space-y-4">
      <NuxtLink v-for="p in pedidos" :key="p.id" :to="`/pedidos/${p.id}`" class="block p-5 rounded-2xl bg-white shadow-xs hover-lift">
        <div class="flex items-center justify-between">
          <div>
            <p class="font-display text-heading text-texto">Pedido #{{ p.id }}</p>
            <p class="font-body text-small text-texto-muted mt-1">{{ formatearFecha(p.created_at) }}</p>
          </div>
          <div class="text-right">
            <p class="font-display text-heading text-texto">{{ formatearPrecio(p.total) }}</p>
            <span :class="['inline-block px-3 py-1 rounded-full font-body text-caption font-semibold mt-1 capitalize', badge[p.estado] || 'bg-fondo text-texto-muted']">{{ p.estado }}</span>
          </div>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
