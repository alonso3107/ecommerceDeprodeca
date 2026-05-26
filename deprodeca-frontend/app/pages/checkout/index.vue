<!--
  ═══════════════════════════════════════════════════════════════
  checkout/index.vue — Checkout · DEPRODECA
  Brutalismo aplicado al flujo de pago:
  • Métodos de pago como bloques toggle (delegado a MetodoPagoSelector)
  • Panel resumen con bordes duros, sin curvas
  • Éxito con ícono geométrico (diamante + check, no círculo)
  • Botón confirmar: bloque negro sólido masivo
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({
  layout: "default",
  middleware: "auth",
})

const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const carrito = ref<any[]>([])
const metodo = ref("yape")
const loading = ref(false)
const pedidoCreado = ref<any>(null)
const errorMsg = ref("")

// ─── Ciclo de vida ────────────────────────────────────────
onMounted(() => {
  if (import.meta.client) {
    carrito.value = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
  }
  if (carrito.value.length === 0) {
    navigateTo("/carrito")
  }
})

// ─── Computed ─────────────────────────────────────────────
const total = computed(() =>
  carrito.value.reduce((sum: number, item: any) => sum + item.precio * item.cantidad, 0),
)

const totalItems = computed(() =>
  carrito.value.reduce((sum: number, item: any) => sum + item.cantidad, 0),
)

// ─── Utilidades ───────────────────────────────────────────
function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}

// ─── Métodos de pago ──────────────────────────────────────
const metodosPago = [
  {
    id: "yape",
    nombre: "Yape",
    datos: "Yape: 987 654 321",
  },
  {
    id: "plin",
    nombre: "Plin",
    datos: "Plin: 987 654 321",
  },
  {
    id: "transferencia",
    nombre: "Transferencia",
    datos: "BCP: 194-1234567-0-89\nCCI: 0021940123456708999\nTitular: DEPRODECA SAC",
  },
]

// ─── Confirmar pedido ─────────────────────────────────────
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

    <!-- ═══════════════════════════════════════════════════════
         ÉXITO · Confirmación geométrica, sin círculos
         ═══════════════════════════════════════════════════════ -->
    <div v-if="pedidoCreado" class="text-center py-32">
      <!-- Ícono: Diamante geométrico con check -->
      <svg
        class="mx-auto mb-8 text-exito"
        width="80" height="80" viewBox="0 0 80 80" fill="none"
      >
        <!-- Diamante exterior -->
        <path d="M40 4L76 40L40 76L4 40L40 4Z" stroke="currentColor" stroke-width="2"/>
        <!-- Diamante interior -->
        <path d="M40 18L58 40L40 62L22 40L40 18Z" stroke="currentColor" stroke-width="1.5" opacity="0.4"/>
        <!-- Check -->
        <path d="M30 41L37 49L52 33" stroke="currentColor" stroke-width="3"
              stroke-linecap="square" stroke-linejoin="miter"/>
      </svg>

      <p class="font-display text-[clamp(2.5rem,6vw,4rem)] text-texto uppercase leading-[0.9] mb-4">
        ¡Pedido<br />Confirmado<span class="text-exito">!</span>
      </p>

      <div class="flex items-center justify-center gap-3 mb-6">
        <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">
          Pedido #
        </span>
        <span class="font-mono text-[13px] text-texto font-bold tracking-wider">
          {{ pedidoCreado.id }}
        </span>
      </div>

      <p class="font-body text-body text-texto-muted mb-4">
        Redirigiendo a tu pedido...
      </p>

      <!-- Barra de progreso geométrica -->
      <div class="flex justify-center gap-1">
        <span class="w-12 h-1 bg-[#D4A017] animate-pulse" />
        <span class="w-3 h-1 bg-borde" />
        <span class="w-3 h-1 bg-borde" />
      </div>
    </div>

    <!-- ═══════════════════════════════════════════════════════
         CHECKOUT · Layout 2 columnas
         ═══════════════════════════════════════════════════════ -->
    <template v-else>
      <!-- Encabezado -->
      <div class="mb-12">
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-4">
          ─── Checkout · {{ totalItems }} items
        </p>
        <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
          Finalizar<br />Pedido<span class="text-[#D4A017]">.</span>
        </h1>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-[1fr_400px] gap-0 border border-borde">

        <!-- ─── COLUMNA IZQUIERDA · Método de pago ─────────── -->
        <div class="p-8 border-b lg:border-b-0 lg:border-r border-borde">
          <MetodoPagoSelector
            :metodos="metodosPago"
            :metodo-activo="metodo"
            @seleccionar="metodo = $event"
          />
        </div>

        <!-- ─── COLUMNA DERECHA · Resumen ──────────────────── -->
        <div class="bg-white p-8 lg:sticky lg:top-[84px] h-fit">

          <!-- Header con ícono geométrico -->
          <div class="flex items-center gap-3 mb-6">
            <!-- Ícono: Cubo/paquete -->
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" class="text-[#D4A017] flex-shrink-0">
              <rect x="3" y="5" width="14" height="13" stroke="currentColor" stroke-width="1.5"/>
              <path d="M3 5L10 1L17 5" stroke="currentColor" stroke-width="1.5"
                    stroke-linecap="square" stroke-linejoin="miter"/>
              <path d="M10 1V11" stroke="currentColor" stroke-width="1.2" stroke-dasharray="2 2"/>
            </svg>
            <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
              Resumen
            </h3>
          </div>

          <!-- Lista de ítems compacta -->
          <div class="space-y-0 border border-borde mb-6">
            <div
              v-for="(item, i) in carrito"
              :key="item.id"
              class="flex items-center gap-3 px-4 py-3"
              :class="i > 0 ? 'border-t border-borde' : ''"
            >
              <!-- Imagen pequeña, rectangular -->
              <div class="w-10 h-10 border border-borde bg-fondo flex-shrink-0">
                <img
                  :src="item.imagen_url || 'https://images.unsplash.com/photo-1583258292688-d0213dc154b3?w=60&q=60'"
                  :alt="item.nombre"
                  class="w-full h-full object-cover"
                />
              </div>
              <div class="flex-1 min-w-0">
                <p class="font-body text-small font-bold text-texto truncate">
                  {{ item.nombre }}
                </p>
                <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider">
                  {{ item.cantidad }} × {{ formatearPrecio(item.precio) }}
                </p>
              </div>
              <p class="font-body text-small font-bold text-texto flex-shrink-0">
                {{ formatearPrecio(item.precio * item.cantidad) }}
              </p>
            </div>
          </div>

          <!-- Totales -->
          <div class="space-y-3 mb-2">
            <div class="flex justify-between font-body text-small text-texto-muted">
              <span>Subtotal ({{ totalItems }} items)</span>
              <span>{{ formatearPrecio(total) }}</span>
            </div>
            <div class="flex justify-between font-body text-small text-texto-muted">
              <span class="flex items-center gap-1.5">
                Envío
                <svg width="14" height="10" viewBox="0 0 14 10" fill="none" class="text-exito">
                  <rect x="1" y="2" width="9" height="5" stroke="currentColor" stroke-width="1.2"/>
                  <rect x="10" y="1" width="3" height="7" stroke="currentColor" stroke-width="1.2"/>
                  <circle cx="3.5" cy="8" r="1" stroke="currentColor" stroke-width="1"/>
                  <circle cx="8.5" cy="8" r="1" stroke="currentColor" stroke-width="1"/>
                </svg>
              </span>
              <span class="font-mono text-[11px] text-exito uppercase tracking-[0.15em] font-bold">
                Gratis
              </span>
            </div>
          </div>

          <!-- Separador bold -->
          <div class="border-t-2 border-texto my-6" />

          <!-- Total masivo -->
          <div class="flex justify-between items-baseline mb-8">
            <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
              Total
            </span>
            <span class="font-display text-[clamp(1.75rem,3vw,2.25rem)] text-texto leading-none">
              {{ formatearPrecio(total) }}
            </span>
          </div>

          <!-- Error -->
          <p
            v-if="errorMsg"
            class="font-mono text-[11px] text-error uppercase tracking-[0.1em] mb-4 text-center font-bold"
            role="alert"
          >
            {{ errorMsg }}
          </p>

          <!-- Botón CONFIRMAR · Negro masivo -->
          <button
            class="w-full bg-texto text-white font-display text-heading uppercase tracking-[0.05em]
                   py-5 hover:bg-[#D4A017] hover:text-black
                   transition-colors duration-200 min-h-[60px]
                   disabled:opacity-40 disabled:cursor-not-allowed
                   flex items-center justify-center gap-3 group"
            :disabled="loading"
            @click="confirmarPedido"
          >
            <span v-if="loading" class="w-5 h-5 border-2 border-white border-t-transparent animate-spin" />
            <template v-else>
              Confirmar Pedido
              <!-- Ícono: Flecha diagonal bold (checkout) -->
              <svg
                width="18" height="18" viewBox="0 0 18 18" fill="none"
                class="transition-transform duration-300 group-hover:translate-x-0.5 group-hover:-translate-y-0.5"
              >
                <path d="M3 15L15 3M15 3H7M15 3V11" stroke="currentColor" stroke-width="2"
                      stroke-linecap="square" stroke-linejoin="miter"/>
              </svg>
            </template>
          </button>

          <!-- Disclaimer -->
          <p class="mt-5 font-mono text-[10px] text-texto-muted text-center uppercase tracking-[0.1em]">
            Al confirmar aceptas los términos y condiciones
          </p>

        </div>
      </div>
    </template>
  </div>
</template>
