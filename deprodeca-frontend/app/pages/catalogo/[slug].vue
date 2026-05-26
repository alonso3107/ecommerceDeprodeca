<script setup lang="ts">
definePageMeta({ layout: "default" })
const route = useRoute()
const config = useRuntimeConfig()
const slug = computed(() => route.params.slug as string)
const producto = ref<any>(null)
const loading = ref(true)
const cantidad = ref(1)
const agregando = ref(false)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

onMounted(async () => {
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/productos/${slug.value}`)
    if (res.success) producto.value = res.data
  } catch (_) { producto.value = null }
  finally { loading.value = false }
})

function formatearPrecio(precio: number) { return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(precio) }

function agregarAlCarrito() {
  if (!producto.value) return
  agregando.value = true
  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    const idx = carrito.findIndex((i: any) => i.id === producto.value.id)
    if (idx >= 0) { carrito[idx].cantidad += cantidad.value }
    else { carrito.push({ id: producto.value.id, nombre: producto.value.nombre, slug: producto.value.slug, precio: producto.value.precio, imagen_url: producto.value.imagen_url, unidad: producto.value.unidad, cantidad: cantidad.value }) }
    localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito))
    window.dispatchEvent(new Event("carrito-actualizado"))
  }
  toastMsg.value = `${cantidad.value} x ${producto.value.nombre} agregado`
  toastType.value = "success"; showToast.value = true
  setTimeout(() => showToast.value = false, 3000)
  agregando.value = false
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <div v-if="loading" class="flex justify-center py-20"><Spinner size="lg" /></div>

    <div v-else-if="!producto" class="text-center py-20">
      <h2 class="font-display text-display-md text-texto mb-4">Producto no encontrado</h2>
      <Button variant="primary" @click="navigateTo('/catalogo')">Volver al Catálogo</Button>
    </div>

    <template v-else>
      <nav class="mb-8 font-body text-small text-texto-muted">
        <NuxtLink to="/catalogo" class="hover:text-texto transition-colors">Catálogo</NuxtLink>
        <span class="mx-2">/</span>
        <span class="text-texto font-semibold">{{ producto.nombre }}</span>
      </nav>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-12">
        <div class="relative rounded-3xl overflow-hidden bg-fondo aspect-square">
          <img :src="producto.imagen_url || 'https://images.unsplash.com/photo-1553530979-3f62c23fb101?w=600&q=80'" :alt="producto.nombre" class="w-full h-full object-cover" />
        </div>

        <div class="flex flex-col justify-center">
          <p class="font-body text-small text-texto-muted uppercase tracking-wide mb-2">{{ producto.categoria_nombre }}</p>
          <h1 class="font-display text-display-lg text-texto uppercase mb-4 leading-[1.05]">{{ producto.nombre }}</h1>
          <p class="font-body text-body text-texto-muted mb-8 leading-relaxed">{{ producto.descripcion || "Producto de calidad para tu bodega." }}</p>

          <div class="flex items-baseline gap-2 mb-6">
            <span class="font-display text-display-xl text-texto leading-none">{{ formatearPrecio(producto.precio) }}</span>
            <span class="font-body text-body text-texto-muted">/ {{ producto.unidad }}</span>
          </div>

          <p class="font-body text-small mb-8">
            <span v-if="producto.stock > 20" class="text-exito font-semibold">{{ producto.stock }} en stock</span>
            <span v-else-if="producto.stock > 0" class="text-advertencia font-semibold">Solo {{ producto.stock }} quedan</span>
            <span v-else class="text-error font-semibold">Agotado</span>
          </p>

          <div v-if="producto.stock > 0" class="flex items-center gap-4">
            <div class="flex items-center border border-borde rounded-xl">
              <button class="px-4 py-3 font-body text-body hover:bg-fondo transition-colors min-h-[48px] min-w-[48px]" :disabled="cantidad <= 1" @click="cantidad--">&minus;</button>
              <span class="px-4 py-3 font-body text-body font-semibold min-w-[48px] text-center">{{ cantidad }}</span>
              <button class="px-4 py-3 font-body text-body hover:bg-fondo transition-colors min-h-[48px] min-w-[48px]" :disabled="cantidad >= producto.stock" @click="cantidad++">+</button>
            </div>
            <Button variant="primary" size="lg" :loading="agregando" @click="agregarAlCarrito">Agregar al Carrito</Button>
          </div>
          <Button v-else variant="outline" size="lg" disabled>Agotado</Button>
        </div>
      </div>
    </template>

    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
  </div>
</template>
