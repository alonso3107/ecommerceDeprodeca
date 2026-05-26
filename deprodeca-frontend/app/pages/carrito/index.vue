<script setup lang="ts">
definePageMeta({ layout: "default" })

const carrito = ref<any[]>([])
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

onMounted(() => { if (import.meta.client) carrito.value = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]") })
const total = computed(() => carrito.value.reduce((s: number, i: any) => s + i.precio * i.cantidad, 0))

function formatearPrecio(p: number) { return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(p) }
function guardarCarrito() { if (import.meta.client) { localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito.value)); window.dispatchEvent(new Event("carrito-actualizado")) } }
function actualizarCantidad(item: any, d: number) { item.cantidad += d; if (item.cantidad <= 0) { carrito.value = carrito.value.filter((i: any) => i.id !== item.id) } guardarCarrito() }
function eliminarItem(item: any) { carrito.value = carrito.value.filter((i: any) => i.id !== item.id); guardarCarrito() }

function irACheckout() {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) { toastMsg.value = "Iniciá sesión para continuar"; toastType.value = "error"; showToast.value = true; setTimeout(() => showToast.value = false, 3000); return }
  navigateTo("/checkout")
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <h1 class="font-display text-display-lg text-texto uppercase mb-10">Carrito</h1>

    <div v-if="carrito.length === 0" class="text-center py-20">
      <p class="font-display text-display-md text-texto-muted mb-4">Carrito vacío</p>
      <Button variant="primary" @click="navigateTo('/catalogo')">Explorar Catálogo</Button>
    </div>

    <template v-else>
      <div class="grid grid-cols-1 lg:grid-cols-[1fr_380px] gap-10">
        <div class="space-y-4">
          <div v-for="item in carrito" :key="item.id" class="flex items-center gap-4 p-5 rounded-2xl bg-white shadow-xs">
            <img :src="item.imagen_url || 'https://images.unsplash.com/photo-1578916171725-4665e9893b41?w=100&q=60'" :alt="item.nombre" class="w-20 h-20 rounded-xl object-cover flex-shrink-0" />
            <div class="flex-1 min-w-0">
              <h3 class="font-body text-body font-semibold text-texto truncate">{{ item.nombre }}</h3>
              <p class="font-body text-small text-texto-muted">{{ formatearPrecio(item.precio) }} / {{ item.unidad }}</p>
            </div>
            <div class="flex items-center border border-borde rounded-xl">
              <button class="px-3 py-2 font-body text-small hover:bg-fondo transition-colors min-h-[40px] min-w-[40px]" @click="actualizarCantidad(item, -1)">&minus;</button>
              <span class="px-3 font-body text-small font-semibold min-w-[40px] text-center">{{ item.cantidad }}</span>
              <button class="px-3 py-2 font-body text-small hover:bg-fondo transition-colors min-h-[40px] min-w-[40px]" @click="actualizarCantidad(item, 1)">+</button>
            </div>
            <p class="font-display text-heading text-texto w-28 text-right">{{ formatearPrecio(item.precio * item.cantidad) }}</p>
            <button class="p-2 rounded-lg hover:bg-fondo transition-colors min-h-[44px] min-w-[44px] text-texto-muted hover:text-error" @click="eliminarItem(item)" aria-label="Eliminar">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
            </button>
          </div>
        </div>

        <div class="bg-white rounded-2xl shadow-xs p-6 h-fit lg:sticky lg:top-[84px]">
          <h3 class="font-body text-subheading font-bold text-texto mb-5">Resumen</h3>
          <div class="space-y-3 mb-6">
            <div class="flex justify-between font-body text-small text-texto-muted"><span>Productos</span><span>{{ formatearPrecio(total) }}</span></div>
            <div class="flex justify-between font-body text-small text-texto-muted"><span>Envío</span><span class="text-exito font-semibold">Gratis</span></div>
          </div>
          <div class="border-t border-borde pt-4 mb-6">
            <div class="flex justify-between items-baseline"><span class="font-body text-body font-bold text-texto">Total</span><span class="font-display text-display-md text-texto">{{ formatearPrecio(total) }}</span></div>
          </div>
          <Button variant="primary" size="lg" :full-width="true" @click="irACheckout">Ir a Checkout</Button>
        </div>
      </div>
    </template>
    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
  </div>
</template>
