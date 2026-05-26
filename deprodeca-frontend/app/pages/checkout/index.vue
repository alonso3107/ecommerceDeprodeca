<script setup lang="ts">
definePageMeta({
  layout: "default",
  middleware: "auth",
})

const config = useRuntimeConfig()
const carrito = ref<any[]>([])
const metodo = ref("yape")
const loading = ref(false)
const pedidoCreado = ref<any>(null)
const errorMsg = ref("")

onMounted(() => {
  if (import.meta.client) {
    carrito.value = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
  }
  if (carrito.value.length === 0) {
    navigateTo("/carrito")
  }
})

const total = computed(() =>
  carrito.value.reduce((sum: number, item: any) => sum + item.precio * item.cantidad, 0),
)

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(precio)
}

const metodosPago = [
  {
    id: "yape",
    nombre: "Yape",
    icono: "M21 12a9 9 0 11-18 0 9 9 0 0118 0z M9 9.5V12l2 1.5L9 15v2.5 M15 9.5V12l-2 1.5 2 1.5v2.5",
    datos: "Yape: 987 654 321",
  },
  {
    id: "plin",
    nombre: "Plin",
    icono: "M21 12a9 9 0 11-18 0 9 9 0 0118 0z M9 10l2 2 4-4",
    datos: "Plin: 987 654 321",
  },
  {
    id: "transferencia",
    nombre: "Transferencia",
    icono: "M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z",
    datos: "BCP: 194-1234567-0-89\nCCI: 0021940123456708999\nTitular: DEPRODECA SAC",
  },
]

async function confirmarPedido() {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) {
    errorMsg.value = "Debes iniciar sesión"
    return
  }

  loading.value = true
  errorMsg.value = ""

  try {
    // 1. Crear pedido
    const items = carrito.value.map((item: any) => ({
      producto_id: item.id,
      cantidad: item.cantidad,
    }))

    const pedidoRes = await $fetch<{ success: boolean; data: any; message: string }>(
      `${config.public.apiBase}/pedidos`,
      {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
        body: { items },
      },
    )

    if (!pedidoRes.success) {
      errorMsg.value = pedidoRes.message
      return
    }

    // 2. Registrar pago
    await $fetch(`${config.public.apiBase}/pagos`, {
      method: "POST",
      headers: { Authorization: `Bearer ${token}` },
      body: {
        pedido_id: pedidoRes.data.id,
        monto: total.value,
        metodo: metodo.value,
        comprobante_url: "",
      },
    })

    // 3. Limpiar carrito
    if (import.meta.client) {
      localStorage.removeItem("deprodeca_carrito")
      window.dispatchEvent(new Event("carrito-actualizado"))
    }

    pedidoCreado.value = pedidoRes.data
    setTimeout(() => navigateTo(`/pedidos/${pedidoRes.data.id}`), 3000)
  } catch (e: any) {
    errorMsg.value = e?.data?.message || "Error al procesar el pedido"
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <h1 class="font-display text-display-lg text-texto uppercase mb-10">Checkout</h1>

    <!-- Confirmado -->
    <div v-if="pedidoCreado" class="text-center py-20">
      <div class="w-20 h-20 mx-auto mb-6 rounded-full bg-exito/10 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-exito"><path d="M5 13l4 4L19 7"/></svg>
      </div>
      <h2 class="font-display text-display-md text-texto uppercase mb-2">¡Pedido Confirmado!</h2>
      <p class="font-body text-body text-texto-muted mb-2">Pedido #{{ pedidoCreado.id }}</p>
      <p class="font-body text-small text-texto-muted">Redirigiendo a tu pedido...</p>
    </div>

    <template v-else>
      <div class="grid grid-cols-1 lg:grid-cols-[1fr_420px] gap-10">

        <!-- Método de pago -->
        <div>
          <h2 class="font-body text-heading font-bold text-texto mb-6">Método de Pago</h2>

          <div class="space-y-3">
            <button
              v-for="m in metodosPago"
              :key="m.id"
              :class="[
                'w-full flex items-center gap-4 p-5 rounded-2xl border-2 transition-all duration-300 text-left',
                metodo === m.id
                  ? 'border-texto bg-fondo shadow-sm'
                  : 'border-borde hover:border-borde-hover bg-white',
              ]"
              @click="metodo = m.id"
            >
              <div :class="['w-5 h-5 rounded-full border-2 flex items-center justify-center flex-shrink-0', metodo === m.id ? 'border-texto' : 'border-borde']">
                <div v-if="metodo === m.id" class="w-2.5 h-2.5 rounded-full bg-texto" />
              </div>
              <span class="font-body text-body font-semibold text-texto">{{ m.nombre }}</span>
            </button>
          </div>

          <!-- Datos de pago -->
          <div v-if="metodo" class="mt-6 p-6 rounded-2xl bg-fondo border border-borde">
            <h3 class="font-body text-small font-bold text-texto-muted uppercase tracking-wide mb-3">
              Datos para el Pago
            </h3>
            <pre class="font-body text-body text-texto whitespace-pre-line leading-relaxed">{{ metodosPago.find(m => m.id === metodo)?.datos }}</pre>
            <p class="mt-4 font-body text-caption text-texto-muted">
              Realiza el pago por {{ metodosPago.find(m => m.id === metodo)?.nombre }} y luego confirma tu pedido.
            </p>
          </div>
        </div>

        <!-- Resumen -->
        <div class="bg-white rounded-2xl border border-borde shadow-xs p-6 h-fit lg:sticky lg:top-[84px]">
          <h3 class="font-body text-subheading font-bold text-texto mb-5">Resumen del Pedido</h3>

          <div class="space-y-3 mb-6">
            <div
              v-for="item in carrito"
              :key="item.id"
              class="flex items-center gap-3"
            >
              <img
                :src="item.imagen_url || 'https://images.unsplash.com/photo-1583258292688-d0213dc154b3?w=60&q=60'"
                :alt="item.nombre"
                class="w-12 h-12 rounded-lg object-cover flex-shrink-0"
              />
              <div class="flex-1 min-w-0">
                <p class="font-body text-small font-semibold text-texto truncate">{{ item.nombre }}</p>
                <p class="font-body text-caption text-texto-muted">{{ item.cantidad }} x {{ formatearPrecio(item.precio) }}</p>
              </div>
              <p class="font-body text-small font-semibold text-texto">{{ formatearPrecio(item.precio * item.cantidad) }}</p>
            </div>
          </div>

          <div class="border-t border-borde pt-4 space-y-2 mb-6">
            <div class="flex justify-between font-body text-small text-texto-muted">
              <span>Subtotal</span>
              <span>{{ formatearPrecio(total) }}</span>
            </div>
            <div class="flex justify-between font-body text-small text-texto-muted">
              <span>Envío</span>
              <span class="text-exito font-semibold">Gratis</span>
            </div>
          </div>

          <div class="border-t border-borde pt-4 mb-6">
            <div class="flex justify-between items-baseline">
              <span class="font-body text-body font-bold text-texto">Total</span>
              <span class="font-display text-display-md text-texto">{{ formatearPrecio(total) }}</span>
            </div>
          </div>

          <p v-if="errorMsg" class="text-small text-error font-medium mb-4 text-center" role="alert">
            {{ errorMsg }}
          </p>

          <Button
            variant="primary"
            size="lg"
            :full-width="true"
            :loading="loading"
            @click="confirmarPedido"
          >
            Confirmar Pedido — {{ formatearPrecio(total) }}
          </Button>

          <p class="mt-3 text-center font-body text-caption text-texto-muted">
            Al confirmar aceptas nuestros términos y condiciones.
          </p>
        </div>
      </div>
    </template>
  </div>
</template>
