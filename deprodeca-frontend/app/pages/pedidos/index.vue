<script setup lang="ts">
definePageMeta({
  layout: "default",
  middleware: "auth",
})

const config = useRuntimeConfig()
const pedidos = ref<any[]>([])
const loading = ref(true)

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return

  try {
    const res = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/pedidos`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    if (res.success) pedidos.value = res.data
  } catch (_) {
    pedidos.value = []
  } finally {
    loading.value = false
  }
})

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(precio)
}

function formatearFecha(fecha: string) {
  return new Date(fecha).toLocaleDateString("es-PE", {
    day: "numeric", month: "short", year: "numeric",
  })
}

const estadoBadge: Record<string, string> = {
  pendiente: "bg-status-warning/20 text-status-warning",
  confirmado: "bg-brand-primary/20 text-text-primary",
  enviado: "bg-brand-secondary/20 text-brand-secondary",
  entregado: "bg-status-success/20 text-status-success",
  cancelado: "bg-status-error/20 text-status-error",
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    <h1 class="font-display text-display-lg text-text-primary uppercase mb-8">Mis Pedidos</h1>

    <div v-if="loading" class="flex justify-center py-20">
      <Spinner size="lg" label="Cargando pedidos..." />
    </div>

    <div
      v-else-if="pedidos.length === 0"
      class="text-center py-20"
    >
      <div class="w-20 h-20 mx-auto mb-4 rounded-full bg-surface flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-text-muted"><path d="M6 2 3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4Z"/><path d="M3 6h18"/><path d="M16 10a4 4 0 01-8 0"/></svg>
      </div>
      <h3 class="font-body text-subheading font-bold text-text-primary mb-2">Sin pedidos aún</h3>
      <p class="font-body text-small text-text-muted mb-6">Crea tu primer pedido desde el carrito.</p>
      <Button variant="primary" @click="navigateTo('/catalogo')">Ir al Catálogo</Button>
    </div>

    <div v-else class="space-y-4">
      <NuxtLink
        v-for="p in pedidos"
        :key="p.id"
        :to="`/pedidos/${p.id}`"
        class="block p-5 rounded-xl border-2 border-border-base hover:border-brand-primary hover:shadow-brand transition-all duration-300"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="font-display text-heading text-text-primary">Pedido #{{ p.id }}</p>
            <p class="font-body text-small text-text-muted mt-0.5">{{ formatearFecha(p.created_at) }}</p>
          </div>
          <div class="text-right">
            <p class="font-display text-heading text-text-primary">{{ formatearPrecio(p.total) }}</p>
            <span
              :class="['inline-block px-2.5 py-0.5 rounded-full font-body text-caption font-semibold mt-1', estadoBadge[p.estado] || 'bg-surface text-text-muted']"
            >
              {{ p.estado }}
            </span>
          </div>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>
