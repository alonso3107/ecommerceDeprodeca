     1|<script setup lang="ts">
     2|definePageMeta({
     3|  layout: "default",
     4|  middleware: "auth",
     5|})
     6|
     7|const route = useRoute()
     8|const config = useRuntimeConfig()
     9|const pedidoId = computed(() => route.params.id as string)
    10|
    11|const pedido = ref<any>(null)
    12|const loading = ref(true)
    13|
    14|onMounted(async () => {
    15|  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
    16|  if (!token) return
    17|
    18|  try {
    19|    const res = await $fetch<{ success: boolean; data: any }>(
    20|      `${config.public.apiBase}/pedidos/${pedidoId.value}`,
    21|      { headers: { Authorization: `Bearer ${token}` } },
    22|    )
    23|    if (res.success) pedido.value = res.data
    24|  } catch (_) {
    25|    pedido.value = null
    26|  } finally {
    27|    loading.value = false
    28|  }
    29|})
    30|
    31|function formatearPrecio(precio: number) {
    32|  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(precio)
    33|}
    34|
    35|function formatearFecha(fecha: string) {
    36|  return new Date(fecha).toLocaleDateString("es-PE", {
    37|    day: "numeric", month: "long", year: "numeric", hour: "2-digit", minute: "2-digit",
    38|  })
    39|}
    40|</script>
    41|
    42|<template>
    43|  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    44|    <div v-if="loading" class="flex justify-center py-20">
    45|      <Spinner size="lg" label="Cargando pedido..." />
    46|    </div>
    47|
    48|    <div v-else-if="!pedido" class="text-center py-20">
    49|      <h2 class="font-display text-display-md text-text-primary mb-4">Pedido no encontrado</h2>
    50|      <Button variant="outline" @click="navigateTo('/pedidos')">Volver a Pedidos</Button>
    51|    </div>
    52|
    53|    <template v-else>
    54|      <div class="flex items-center justify-between mb-8">
    55|        <div>
    56|          <h1 class="font-display text-display-lg text-text-primary uppercase">
    57|            Pedido #{{ pedido.id }}
    58|          </h1>
    59|          <p class="font-body text-small text-text-muted mt-1">
    60|            {{ formatearFecha(pedido.created_at) }}
    61|          </p>
    62|        </div>
    63|        <span
    64|          class="px-4 py-1.5 rounded-full font-body text-small font-semibold capitalize"
    65|          :class="{
    66|            pendiente: 'bg-status-warning/20 text-status-warning',
    67|            confirmado: 'bg-brand-primary/20 text-text-primary',
    68|            enviado: 'bg-brand-secondary/20 text-brand-secondary',
    69|            entregado: 'bg-status-success/20 text-status-success',
    70|            cancelado: 'bg-status-error/20 text-status-error',
    71|          }[pedido.estado] || 'bg-surface text-text-muted'"
    72|        >
    73|          {{ pedido.estado }}
    74|        </span>
    75|      </div>
    76|
    77|      <!-- Detalles -->
    78|      <div class="bg-surface rounded-2xl border-2 border-border-base p-6 mb-8">
    79|        <h3 class="font-body text-subheading font-bold text-text-primary mb-4">Productos</h3>
    80|        <div class="space-y-3">
    81|          <div
    82|            v-for="det in pedido.detalles"
    83|            :key="det.id"
    84|            class="flex items-center gap-4 py-3 border-b border-border-base last:border-0"
    85|          >
    86|            <img
    87|              v-if="det.producto_imagen"
    88|              :src="det.producto_imagen"
    89|              class="w-12 h-12 rounded-lg object-cover flex-shrink-0"
    90|            />
    91|            <div class="flex-1">
    92|              <p class="font-body text-body font-semibold text-text-primary">{{ det.producto_nombre }}</p>
    93|              <p class="font-body text-caption text-text-muted">{{ det.cantidad }} x {{ formatearPrecio(det.precio_unitario) }}</p>
    94|            </div>
    95|            <p class="font-display text-heading text-text-primary">{{ formatearPrecio(det.subtotal) }}</p>
    96|          </div>
    97|        </div>
    98|
    99|        <div class="mt-4 pt-4 border-t-2 border-border-base text-right">
   100|          <p class="font-body text-small text-text-muted">Total</p>
   101|          <p class="font-display text-display-md text-text-primary">{{ formatearPrecio(pedido.total) }}</p>
   102|        </div>
   103|      </div>
   104|
   105|      <div class="text-center">
   106|        <Button variant="outline" @click="navigateTo('/pedidos')">← Volver a Mis Pedidos</Button>
   107|      </div>
   108|    </template>
   109|  </div>
   110|</template>
   111|