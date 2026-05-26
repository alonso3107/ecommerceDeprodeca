<!--
  ═══════════════════════════════════════════════════════════════
  carrito/index.vue — Carrito de Compras · DEPRODECA
  Brutalismo funcional: ítems como bloques rectangulares,
  íconos geométricos únicos (no los típicos redondeados),
  panel de resumen con bordes duros, tipografía masiva.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })

// ─── Estado ───────────────────────────────────────────────
const carrito = ref<any[]>([])
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

onMounted(() => {
  if (import.meta.client) {
    carrito.value = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
  }
})

// ─── Computed ─────────────────────────────────────────────
const total = computed(() =>
  carrito.value.reduce((s: number, i: any) => s + i.precio * i.cantidad, 0)
)

const totalItems = computed(() =>
  carrito.value.reduce((s: number, i: any) => s + i.cantidad, 0)
)

// ─── Utilidades ───────────────────────────────────────────
function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(p)
}

function guardarCarrito() {
  if (import.meta.client) {
    localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito.value))
    window.dispatchEvent(new Event("carrito-actualizado"))
  }
}

function actualizarCantidad(item: any, d: number) {
  item.cantidad += d
  if (item.cantidad <= 0) {
    carrito.value = carrito.value.filter((i: any) => i.id !== item.id)
  }
  guardarCarrito()
}

function eliminarItem(item: any) {
  carrito.value = carrito.value.filter((i: any) => i.id !== item.id)
  guardarCarrito()
}

function irACheckout() {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) {
    toastMsg.value = "Iniciá sesión para continuar"
    toastType.value = "error"
    showToast.value = true
    setTimeout(() => (showToast.value = false), 3000)
    return
  }
  navigateTo("/checkout")
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">

    <!-- ═══════════════════════════════════════════════════════
         ENCABEZADO · Metadata técnica + título masivo
         ═══════════════════════════════════════════════════════ -->
    <div class="mb-12">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-4">
        ─── Carrito · {{ totalItems }} items
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Tu Pedido<span class="text-[#D4A017]">.</span>
      </h1>
    </div>

    <!-- ═══════════════════════════════════════════════════════
         CARRITO VACÍO · Watermark masivo con ícono geométrico
         ═══════════════════════════════════════════════════════ -->
    <div v-if="carrito.length === 0" class="text-center py-32">
      <!-- Ícono: Caja geométrica hueca (estilo isométrico abstracto) -->
      <svg
        class="mx-auto mb-10 text-texto/10"
        width="80" height="80" viewBox="0 0 80 80" fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <!-- Cara superior -->
        <path d="M40 8L68 24V56L40 72L12 56V24L40 8Z" stroke="currentColor" stroke-width="2"/>
        <!-- Borde frontal -->
        <path d="M12 24L40 40L68 24" stroke="currentColor" stroke-width="2"/>
        <!-- Borde inferior -->
        <path d="M12 56L40 72L68 56" stroke="currentColor" stroke-width="2"/>
        <!-- Línea vertical -->
        <path d="M40 40V72" stroke="currentColor" stroke-width="2"/>
        <!-- Apertura superior (tapa abierta) -->
        <path d="M28 20L40 12L52 20" stroke="currentColor" stroke-width="1.5" stroke-dasharray="4 2"/>
      </svg>

      <p class="font-display text-[clamp(3rem,8vw,7rem)] text-texto/10 uppercase leading-none mb-6">
        Carrito<br />Vacío
      </p>
      <p class="font-body text-body text-texto-muted mb-10 max-w-md mx-auto">
        Tu caja está esperando. Explorá el catálogo y llenala con productos Nestlé para tu bodega.
      </p>
      <button
        class="font-mono text-[11px] uppercase tracking-[0.2em] border border-borde px-8 py-4
               text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017]
               transition-colors duration-200"
        @click="navigateTo('/catalogo')"
      >
        Explorar Catálogo →
      </button>
    </div>

    <!-- ═══════════════════════════════════════════════════════
         CARRITO CON ITEMS · Layout 2 columnas brutalista
         ═══════════════════════════════════════════════════════ -->
    <template v-else>
      <div class="grid grid-cols-1 lg:grid-cols-[1fr_380px] gap-0 border border-borde">

        <!-- ─── COLUMNA IZQUIERDA · Lista de ítems ─────────── -->
        <div class="border-b lg:border-b-0 lg:border-r border-borde">

          <!-- Header row de la tabla -->
          <div class="hidden md:grid grid-cols-[80px_1fr_140px_120px_48px] gap-4 px-6 py-4
                      bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em]">
            <span></span>
            <span>Producto</span>
            <span>Cantidad</span>
            <span>Subtotal</span>
            <span></span>
          </div>

          <!-- Items -->
          <div
            v-for="(item, i) in carrito"
            :key="item.id"
            class="grid grid-cols-1 md:grid-cols-[80px_1fr_140px_120px_48px] gap-4 px-6 py-5
                   items-center border-b border-borde last:border-b-0
                   hover:bg-[#D4A017]/[0.02] transition-colors duration-200"
            :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
          >

            <!-- Imagen · Rectangular, sin curvas -->
            <div class="w-20 h-20 border border-borde bg-fondo flex-shrink-0 mx-auto md:mx-0">
              <img
                :src="item.imagen_url || 'https://images.unsplash.com/photo-1578916171725-4665e9893b41?w=100&q=60'"
                :alt="item.nombre"
                class="w-full h-full object-cover"
              />
            </div>

            <!-- Info del producto -->
            <div class="min-w-0 text-center md:text-left">
              <h3 class="font-body text-body font-bold text-texto truncate">
                {{ item.nombre }}
              </h3>
              <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider mt-1">
                {{ formatearPrecio(item.precio) }} <span class="text-borde">/</span> {{ item.unidad }}
              </p>
            </div>

            <!-- Cantidad · Bloques duros -->
            <div class="flex justify-center md:justify-start">
              <div class="flex border border-borde">
                <button
                  class="px-4 py-2.5 font-mono text-sm hover:bg-fondo transition-colors
                         min-h-[42px] min-w-[42px] flex items-center justify-center"
                  @click="actualizarCantidad(item, -1)"
                >
                  <!-- Ícono: Menos geométrico (línea horizontal en cuadrado) -->
                  <svg width="10" height="2" viewBox="0 0 10 2" fill="none">
                    <rect width="10" height="2" fill="currentColor"/>
                  </svg>
                </button>
                <span class="px-4 py-2.5 font-mono text-sm font-bold text-texto
                             border-x border-borde min-w-[48px] flex items-center justify-center">
                  {{ item.cantidad }}
                </span>
                <button
                  class="px-4 py-2.5 font-mono text-sm hover:bg-fondo transition-colors
                         min-h-[42px] min-w-[42px] flex items-center justify-center"
                  @click="actualizarCantidad(item, 1)"
                >
                  <!-- Ícono: Más geométrico (cruz en cuadrado) -->
                  <svg width="10" height="10" viewBox="0 0 10 10" fill="none">
                    <rect x="4" width="2" height="10" fill="currentColor"/>
                    <rect y="4" width="10" height="2" fill="currentColor"/>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Subtotal · Tipografía display -->
            <p class="font-display text-heading text-texto text-center md:text-right">
              {{ formatearPrecio(item.precio * item.cantidad) }}
            </p>

            <!-- Eliminar · Ícono geométrico único -->
            <div class="flex justify-center md:justify-end">
              <button
                class="group p-2 hover:bg-error/5 transition-colors duration-200
                       min-h-[44px] min-w-[44px] flex items-center justify-center"
                @click="eliminarItem(item)"
                aria-label="Eliminar producto"
              >
                <!-- Ícono: "X" en rombo — forma geométrica no estándar -->
                <svg
                  width="18" height="18" viewBox="0 0 18 18" fill="none"
                  class="text-texto-muted group-hover:text-error transition-colors duration-200"
                >
                  <!-- Rombo exterior -->
                  <path d="M9 1L17 9L9 17L1 9L9 1Z" stroke="currentColor" stroke-width="1.5"/>
                  <!-- X interior -->
                  <path d="M6 6L12 12M12 6L6 12" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
                </svg>
              </button>
            </div>
          </div>
        </div>

        <!-- ─── COLUMNA DERECHA · Panel de resumen brutalista ─── -->
        <div class="bg-white p-8 lg:sticky lg:top-[84px] h-fit">

          <!-- Ícono: Diamante geométrico para "Total" -->
          <div class="flex items-center gap-3 mb-6">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" class="text-[#D4A017] flex-shrink-0">
              <path d="M10 1L19 10L10 19L1 10L10 1Z" stroke="currentColor" stroke-width="1.5"/>
              <circle cx="10" cy="10" r="3" fill="currentColor"/>
            </svg>
            <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
              Resumen
            </h3>
          </div>

          <!-- Líneas de detalle -->
          <div class="space-y-4 mb-6">
            <!-- Productos -->
            <div class="flex justify-between items-baseline">
              <span class="font-body text-small text-texto-muted">Productos ({{ totalItems }})</span>
              <span class="font-mono text-small text-texto tracking-wider">{{ formatearPrecio(total) }}</span>
            </div>

            <!-- Envío · Ícono camión geométrico -->
            <div class="flex justify-between items-baseline">
              <span class="font-body text-small text-texto-muted flex items-center gap-2">
                Envío
                <!-- Ícono: Camión geométrico abstracto -->
                <svg width="16" height="12" viewBox="0 0 16 12" fill="none" class="text-exito">
                  <rect x="1" y="3" width="11" height="6" stroke="currentColor" stroke-width="1.5"/>
                  <rect x="12" y="1" width="3" height="8" stroke="currentColor" stroke-width="1.5"/>
                  <circle cx="4" cy="10" r="1.5" stroke="currentColor" stroke-width="1.2"/>
                  <circle cx="10" cy="10" r="1.5" stroke="currentColor" stroke-width="1.2"/>
                </svg>
              </span>
              <span class="font-mono text-[11px] text-exito uppercase tracking-[0.15em] font-bold">
                Gratis
              </span>
            </div>
          </div>

          <!-- Separador -->
          <div class="border-t-2 border-texto mb-6" />

          <!-- Total · Masivo -->
          <div class="flex justify-between items-baseline mb-8">
            <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
              Total
            </span>
            <span class="font-display text-[clamp(1.75rem,3vw,2.25rem)] text-texto leading-none">
              {{ formatearPrecio(total) }}
            </span>
          </div>

          <!-- Botón Checkout · Negro sólido, brutalista -->
          <button
            class="w-full bg-texto text-white font-display text-heading uppercase tracking-[0.05em]
                   py-5 hover:bg-[#D4A017] hover:text-black
                   transition-colors duration-200 min-h-[60px]
                   flex items-center justify-center gap-3 group"
            @click="irACheckout"
          >
            Ir a Checkout
            <!-- Ícono: Flecha geométrica bold -->
            <svg
              width="20" height="12" viewBox="0 0 20 12" fill="none"
              class="transition-transform duration-300 group-hover:translate-x-1"
            >
              <path d="M0 6H18M18 6L13 1M18 6L13 11" stroke="currentColor" stroke-width="2.5"
                    stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
          </button>

          <!-- Info adicional · Seguridad -->
          <p class="mt-6 font-mono text-[10px] text-texto-muted text-center uppercase tracking-[0.15em]">
            Pago seguro · Sin comisiones ocultas
          </p>
        </div>

      </div>
    </template>

    <!-- Toast -->
    <Toast
      :message="toastMsg"
      :type="toastType"
      :show="showToast"
      @close="showToast = false"
    />
  </div>
</template>
