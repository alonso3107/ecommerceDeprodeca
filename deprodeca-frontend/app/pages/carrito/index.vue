     1|<script setup lang="ts">
     2|definePageMeta({
     3|  layout: "default",
     4|})
     5|
     6|const carrito = ref<any[]>([])
     7|const loading = ref(false)
     8|const toastMsg = ref("")
     9|const toastType = ref<"success" | "error">("success")
    10|const showToast = ref(false)
    11|
    12|onMounted(() => {
    13|  if (import.meta.client) {
    14|    carrito.value = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    15|  }
    16|})
    17|
    18|const total = computed(() =>
    19|  carrito.value.reduce((sum, item) => sum + item.precio * item.cantidad, 0),
    20|)
    21|
    22|function formatearPrecio(precio: number) {
    23|  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(precio)
    24|}
    25|
    26|function guardarCarrito() {
    27|  if (import.meta.client) {
    28|    localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito.value))
    29|    window.dispatchEvent(new Event("carrito-actualizado"))
    30|  }
    31|}
    32|
    33|function actualizarCantidad(item: any, delta: number) {
    34|  item.cantidad += delta
    35|  if (item.cantidad <= 0) {
    36|    eliminarItem(item)
    37|    return
    38|  }
    39|  guardarCarrito()
    40|}
    41|
    42|function eliminarItem(item: any) {
    43|  carrito.value = carrito.value.filter((i: any) => i.id !== item.id)
    44|  guardarCarrito()
    45|}
    46|
    47|async function crearPedido() {
    48|  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
    49|  if (!token) {
    50|    toastMsg.value = "Debes iniciar sesión para crear un pedido"
    51|    toastType.value = "error"
    52|    showToast.value = true
    53|    setTimeout(() => showToast.value = false, 3000)
    54|    return
    55|  }
    56|
    57|  loading.value = true
    58|  try {
    59|    const config = useRuntimeConfig()
    60|    const items = carrito.value.map((item: any) => ({
    61|      producto_id: item.id,
    62|      cantidad: item.cantidad,
    63|    }))
    64|
    65|    const res = await $fetch<{ success: boolean; data: any; message: string }>(
    66|      `${config.public.apiBase}/pedidos`,
    67|      {
    68|        method: "POST",
    69|        headers: { Authorization: `Bearer ${token}` },
    70|        body: { items },
    71|      },
    72|    )
    73|
    74|    if (res.success) {
    75|      carrito.value = []
    76|      if (import.meta.client) localStorage.removeItem("deprodeca_carrito")
    77|      toastMsg.value = "¡Pedido creado exitosamente!"
    78|      toastType.value = "success"
    79|      showToast.value = true
    80|      setTimeout(() => {
    81|        showToast.value = false
    82|        navigateTo(`/pedidos/${res.data.id}`)
    83|      }, 1500)
    84|    }
    85|  } catch (e: any) {
    86|    toastMsg.value = e?.data?.message || "Error al crear el pedido"
    87|    toastType.value = "error"
    88|    showToast.value = true
    89|    setTimeout(() => showToast.value = false, 3000)
    90|  } finally {
    91|    loading.value = false
    92|  }
    93|}
    94|</script>
    95|
    96|<template>
    97|  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    98|    <h1 class="font-display text-display-lg text-text-primary uppercase mb-8">Carrito</h1>
    99|
   100|    <!-- Vacío -->
   101|    <div
   102|      v-if="carrito.length === 0"
   103|      class="text-center py-20"
   104|    >
   105|      <div class="w-20 h-20 mx-auto mb-4 rounded-full bg-surface flex items-center justify-center">
   106|        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-text-muted"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 002 1.58h9.78a2 2 0 001.95-1.57l1.65-7.43H5.12"/></svg>
   107|      </div>
   108|      <h3 class="font-body text-subheading font-bold text-text-primary mb-2">Carrito vacío</h3>
   109|      <p class="font-body text-small text-text-muted mb-6">Agrega productos desde el catálogo.</p>
   110|      <Button variant="primary" size="md" @click="navigateTo('/catalogo')">
   111|        Explorar Catálogo
   112|      </Button>
   113|    </div>
   114|
   115|    <!-- Items -->
   116|    <template v-else>
   117|      <div class="grid grid-cols-1 lg:grid-cols-[1fr_380px] gap-8">
   118|        <!-- Lista de items -->
   119|        <div class="space-y-4">
   120|          <div
   121|            v-for="item in carrito"
   122|            :key="item.id"
   123|            class="flex items-center gap-4 p-4 rounded-xl border-2 border-border-base hover:border-brand-primary transition-all duration-300"
   124|          >
   125|            <img
   126|              :src="item.imagen_url || 'https://images.unsplash.com/photo-1543168256-418811576931?w=100&q=60'"
   127|              :alt="item.nombre"
   128|              class="w-20 h-20 rounded-lg object-cover flex-shrink-0"
   129|            />
   130|
   131|            <div class="flex-1 min-w-0">
   132|              <h3 class="font-body text-body font-semibold text-text-primary truncate">{{ item.nombre }}</h3>
   133|              <p class="font-body text-small text-text-muted">{{ formatearPrecio(item.precio) }} / {{ item.unidad }}</p>
   134|            </div>
   135|
   136|            <div class="flex items-center border-2 border-border-base rounded-lg">
   137|              <button
   138|                class="px-2 py-1 font-body text-small hover:bg-surface transition-colors min-h-[36px] min-w-[36px]"
   139|                @click="actualizarCantidad(item, -1)"
   140|              >−</button>
   141|              <span class="px-3 font-body text-small font-semibold min-w-[36px] text-center">{{ item.cantidad }}</span>
   142|              <button
   143|                class="px-2 py-1 font-body text-small hover:bg-surface transition-colors min-h-[36px] min-w-[36px]"
   144|                @click="actualizarCantidad(item, 1)"
   145|              >+</button>
   146|            </div>
   147|
   148|            <p class="font-display text-heading text-text-primary w-28 text-right">
   149|              {{ formatearPrecio(item.precio * item.cantidad) }}
   150|            </p>
   151|
   152|            <button
   153|              class="p-2 rounded-lg hover:bg-surface transition-colors min-h-[44px] min-w-[44px] flex items-center justify-center text-text-muted hover:text-status-error"
   154|              @click="eliminarItem(item)"
   155|              aria-label="Eliminar producto"
   156|            >
   157|              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
   158|            </button>
   159|          </div>
   160|        </div>
   161|
   162|        <!-- Resumen -->
   163|        <div class="bg-surface rounded-2xl border-2 border-border-base p-6 h-fit lg:sticky lg:top-[88px]">
   164|          <h3 class="font-body text-subheading font-bold text-text-primary mb-4">Resumen del Pedido</h3>
   165|
   166|          <div class="space-y-3 mb-6">
   167|            <div class="flex justify-between font-body text-small text-text-muted">
   168|              <span>Productos ({{ carrito.reduce((s: number, i: any) => s + i.cantidad, 0) }})</span>
   169|              <span>{{ formatearPrecio(total) }}</span>
   170|            </div>
   171|            <div class="flex justify-between font-body text-small text-text-muted">
   172|              <span>Envío</span>
   173|              <span class="text-status-success font-semibold">Gratis</span>
   174|            </div>
   175|          </div>
   176|
   177|          <div class="border-t border-border-base pt-4 mb-6">
   178|            <div class="flex justify-between items-baseline">
   179|              <span class="font-body text-body font-bold text-text-primary">Total</span>
   180|              <span class="font-display text-display-md text-text-primary">{{ formatearPrecio(total) }}</span>
   181|            </div>
   182|          </div>
   183|
   184|          <Button
   185|            variant="primary"
   186|            size="lg"
   187|            :full-width="true"
   188|            :loading="loading"
   189|            @click="crearPedido"
   190|          >
   191|            Crear Pedido
   192|          </Button>
   193|
   194|          <p class="mt-3 text-center font-body text-caption text-text-muted">
   195|            Al crear el pedido aceptas nuestros términos.
   196|          </p>
   197|        </div>
   198|      </div>
   199|    </template>
   200|
   201|    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
   202|  </div>
   203|</template>
   204|