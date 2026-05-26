<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })
const route = useRoute()
const config = useRuntimeConfig()
const pedidoId = computed(() => route.params.id as string)
const pedido = ref<any>(null)
const loading = ref(true)

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/pedidos/${pedidoId.value}`, { headers: { Authorization: `Bearer ${token}` } })
    if (res.success) pedido.value = res.data
  } catch (_) {} finally { loading.value = false }
})

function formatearPrecio(p: number) { return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(p) }
function formatearFecha(f: string) { return new Date(f).toLocaleDateString("es-PE", { day: "numeric", month: "long", year: "numeric", hour: "2-digit", minute: "2-digit" }) }
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <div v-if="loading" class="flex justify-center py-20"><Spinner size="lg" /></div>

    <div v-else-if="!pedido" class="text-center py-20">
      <h2 class="font-display text-display-md text-texto mb-4">Pedido no encontrado</h2>
      <Button variant="outline" @click="navigateTo('/pedidos')">Volver</Button>
    </div>

    <template v-else>
      <div class="flex items-center justify-between mb-10">
        <div>
          <h1 class="font-display text-display-lg text-texto uppercase">Pedido #{{ pedido.id }}</h1>
          <p class="font-body text-small text-texto-muted mt-1">{{ formatearFecha(pedido.created_at) }}</p>
        </div>
        <span class="px-4 py-2 rounded-full font-body text-small font-semibold capitalize" :class="{ pendiente: 'bg-advertencia/10 text-advertencia', confirmado: 'bg-texto/10 text-texto', enviado: 'bg-texto/5 text-texto', entregado: 'bg-exito/10 text-exito', cancelado: 'bg-error/10 text-error' }[pedido.estado] || 'bg-fondo text-texto-muted'">{{ pedido.estado }}</span>
      </div>

      <div class="bg-white rounded-2xl shadow-xs p-8 mb-10">
        <h3 class="font-body text-subheading font-bold text-texto mb-6">Productos</h3>
        <div class="space-y-4">
          <div v-for="det in pedido.detalles" :key="det.id" class="flex items-center gap-4 py-3 border-b border-borde last:border-0">
            <div class="flex-1">
              <p class="font-body text-body font-semibold text-texto">{{ det.producto_nombre }}</p>
              <p class="font-body text-caption text-texto-muted">{{ det.cantidad }} x {{ formatearPrecio(det.precio_unitario) }}</p>
            </div>
            <p class="font-display text-heading text-texto">{{ formatearPrecio(det.subtotal) }}</p>
          </div>
        </div>
        <div class="mt-6 pt-6 border-t-2 border-borde text-right">
          <p class="font-body text-small text-texto-muted">Total</p>
          <p class="font-display text-display-md text-texto">{{ formatearPrecio(pedido.total) }}</p>
        </div>
      </div>

      <Button variant="outline" @click="navigateTo('/pedidos')">&larr; Volver a Mis Pedidos</Button>
    </template>
  </div>
</template>
