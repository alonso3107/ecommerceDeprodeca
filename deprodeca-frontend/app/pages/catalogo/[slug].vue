<!-- catalogo/[slug].vue — Detalle Producto Flat Design -->
<script setup lang="ts">
import { resolverImagenProducto } from '~/shared/utils/imagenes-productos'

definePageMeta({ layout: 'default' })

const route = useRoute()
const config = useRuntimeConfig()

const slug = computed(() => route.params.slug as string)
const producto = ref<any>(null)
const loading = ref(true)
const cantidad = ref(1)
const agregando = ref(false)
const toastMsg = ref('')
const toastType = ref<'success' | 'error'>('success')
const showToast = ref(false)

// Animaciones de entrada
const imagenVisible = ref(false)
const infoVisible = ref(false)
onMounted(() => {
  setTimeout(() => (imagenVisible.value = true), 100)
  setTimeout(() => (infoVisible.value = true), 300)
})

const imagenProducto = computed(() => {
  if (!producto.value) return ''
  return resolverImagenProducto(producto.value, 0, new Set())
})

const imagenRespaldo = computed(() => {
  const nombre = String(producto.value?.nombre || 'DEPRODECA').toUpperCase()
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='800'%3E%3Crect fill='%23F5F0E8' width='800' height='800'/%3E%3Ctext x='400' y='420' text-anchor='middle' font-size='32' font-family='monospace' fill='%23C5BFB5'%3E${encodeURIComponent(nombre)}%3C/text%3E%3C/svg%3E`
})

watch(slug, () => cargarProducto())

async function cargarProducto() {
  loading.value = true
  cantidad.value = 1
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/productos/${slug.value}`)
    producto.value = res.success ? res.data : null
  } catch {
    producto.value = null
  } finally {
    loading.value = false
  }
}

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat('es-PE', { style: 'currency', currency: 'PEN', minimumFractionDigits: 2 }).format(precio)
}

function manejarErrorImagen(e: Event) {
  const img = e.target as HTMLImageElement
  if (img && img.src !== imagenRespaldo.value) img.src = imagenRespaldo.value
}

function claseStock(s: number) {
  if (s > 20) return 'text-[#16A34A]'
  if (s > 0) return 'text-[#A16207]'
  return 'text-[#DC2626]'
}

function textoStock(s: number) {
  if (s > 20) return `${s} en stock`
  if (s > 0) return `Solo ${s} disponibles`
  return 'Sin inventario'
}

function agregarAlCarrito() {
  if (!producto.value || agregando.value) return
  agregando.value = true

  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem('deprodeca_carrito') || '[]')
    const indice = carrito.findIndex((item: any) => item.id === producto.value.id)
    if (indice >= 0) {
      carrito[indice].cantidad += cantidad.value
    } else {
      carrito.push({
        id: producto.value.id,
        nombre: producto.value.nombre,
        slug: producto.value.slug,
        precio: producto.value.precio,
        imagen_url: imagenProducto.value,
        unidad: producto.value.unidad,
        cantidad: cantidad.value,
      })
    }
    localStorage.setItem('deprodeca_carrito', JSON.stringify(carrito))
    window.dispatchEvent(new Event('carrito-actualizado'))
  }

  toastMsg.value = `${cantidad.value} × ${producto.value.nombre} agregado al carrito`
  toastType.value = 'success'
  showToast.value = true
  setTimeout(() => (showToast.value = false), 3000)
  agregando.value = false
}

cargarProducto()
</script>

<template>
  <div class="max-w-[1280px] mx-auto px-6 md:px-8 py-10 md:py-14">
    <!-- Loading -->
    <div v-if="loading" class="flex min-h-[60vh] items-center justify-center">
      <div class="flex flex-col items-center gap-4">
        <div class="h-8 w-8 animate-spin border-2 border-[#1C1917] border-t-transparent" />
        <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-[#5C554D]">Cargando producto</span>
      </div>
    </div>

    <!-- 404 -->
    <div v-else-if="!producto" class="text-center py-24">
      <p class="font-display text-[clamp(2.5rem,7vw,5rem)] text-[#C5BFB5] uppercase leading-none mb-6">Producto no encontrado</p>
      <button class="bg-[#1C1917] text-[#FAFAF9] font-mono text-[11px] uppercase tracking-[0.2em] px-6 py-3 hover:bg-[#A16207] transition-colors duration-300" @click="navigateTo('/catalogo')">
        Volver al catálogo
      </button>
    </div>

    <!-- Detalle -->
    <template v-else>
      <!-- Breadcrumb -->
      <nav class="mb-8 flex flex-wrap items-center gap-2 font-mono text-[10px] uppercase tracking-[0.15em] text-[#5C554D]">
        <NuxtLink to="/catalogo" class="hover:text-[#A16207] transition-colors">Catálogo</NuxtLink>
        <span class="text-[#C5BFB5]">/</span>
        <span>{{ producto.categoria_nombre || 'General' }}</span>
        <span class="text-[#C5BFB5]">/</span>
        <span class="text-[#1C1917] font-bold">{{ producto.nombre }}</span>
      </nav>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-0">
        <!-- Imagen -->
        <section
          class="relative aspect-square bg-[#F5F0E8] overflow-hidden transition-all duration-600 ease-out"
          :class="imagenVisible ? 'opacity-100 scale-100' : 'opacity-0 scale-95'"
        >
          <img
            :src="imagenProducto"
            :alt="producto.nombre"
            class="block h-full w-full object-cover hover:scale-105 transition-transform duration-700 ease-out"
            loading="eager"
            @error="manejarErrorImagen"
          />

          <!-- Badge categoría -->
          <span class="absolute top-4 left-4 bg-[#1C1917] text-[#FAFAF9] font-mono text-[10px] uppercase tracking-[0.15em] px-4 py-2">
            {{ producto.categoria_nombre || 'General' }}
          </span>

          <!-- Agotado overlay -->
          <div v-if="producto.stock === 0" class="absolute inset-0 bg-[#1C1917]/60 flex items-center justify-center">
            <span class="font-display text-6xl font-black text-[#FAFAF9] uppercase tracking-wider">Agotado</span>
          </div>
        </section>

        <!-- Info -->
        <section
          class="lg:pl-16 lg:py-8 transition-all duration-500 ease-out"
          :class="infoVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3'"
          style="transition-delay: 100ms"
        >
          <!-- Categoría badge inline -->
          <span class="inline-block border border-[#A16207] text-[#A16207] font-mono text-[10px] uppercase tracking-[0.2em] px-4 py-1.5 mb-6">
            {{ producto.categoria_nombre || 'General' }}
          </span>

          <h1 class="font-display text-[clamp(2.5rem,5vw,4rem)] font-black text-[#1C1917] uppercase leading-[0.95]">
            {{ producto.nombre }}
          </h1>

          <p class="font-body text-[16px] text-[#5C554D] leading-relaxed mt-6 max-w-[55ch]">
            {{ producto.descripcion || 'Producto para bodega con presentación profesional y disponibilidad según stock.' }}
          </p>

          <!-- Precio -->
          <div class="border-t border-[#C5BFB5] pt-8 mt-8">
            <div class="flex items-baseline gap-x-3 gap-y-1 flex-wrap">
              <span class="font-display text-[clamp(3.5rem,7vw,5rem)] font-black text-[#A16207] leading-none">
                {{ formatearPrecio(producto.precio) }}
              </span>
              <span class="font-mono text-[12px] text-[#5C554D] uppercase tracking-[0.2em]">
                / {{ producto.unidad }}
              </span>
            </div>
          </div>

          <!-- Stock -->
          <div class="flex items-center gap-2 mt-4">
            <span class="w-2.5 h-2.5" :class="producto.stock > 0 ? 'bg-[#16A34A]' : 'bg-[#DC2626]'" />
            <span class="font-mono text-[11px] uppercase tracking-[0.15em]" :class="claseStock(producto.stock)">
              {{ textoStock(producto.stock) }}
            </span>
          </div>

          <!-- Cantidad + Agregar -->
          <div v-if="producto.stock > 0" class="mt-8 flex flex-col sm:flex-row gap-3">
            <div class="flex w-fit border border-[#C5BFB5] bg-white">
              <button
                class="flex min-h-[56px] min-w-[56px] items-center justify-center font-mono text-[18px] text-[#1C1917] hover:bg-[#F5F0E8] transition-colors duration-200 disabled:opacity-30"
                :disabled="cantidad <= 1"
                @click="cantidad--"
                aria-label="Disminuir cantidad"
              >−</button>
              <span class="flex min-h-[56px] min-w-[64px] items-center justify-center border-x border-[#C5BFB5] font-mono text-[12px] uppercase tracking-[0.2em] text-[#1C1917]">
                {{ cantidad }}
              </span>
              <button
                class="flex min-h-[56px] min-w-[56px] items-center justify-center font-mono text-[18px] text-[#1C1917] hover:bg-[#F5F0E8] transition-colors duration-200 disabled:opacity-30"
                :disabled="cantidad >= producto.stock"
                @click="cantidad++"
                aria-label="Aumentar cantidad"
              >+</button>
            </div>

            <button
              class="min-h-[56px] bg-[#A16207] text-[#1C1917] font-display text-lg font-bold uppercase tracking-[0.05em] px-10 hover:bg-[#1C1917] hover:text-[#FAFAF9] active:scale-95 transition-all duration-300 disabled:opacity-50 flex-1 sm:flex-initial"
              :disabled="agregando"
              @click="agregarAlCarrito"
            >
              <span v-if="agregando" class="inline-flex items-center gap-3">
                <span class="h-4 w-4 animate-spin border-2 border-[#1C1917] border-t-transparent" />
                Agregando
              </span>
              <span v-else>Agregar al carrito</span>
            </button>
          </div>

          <button v-else disabled class="mt-8 min-h-[56px] w-full sm:w-auto border border-[#C5BFB5] bg-white font-mono text-[11px] uppercase tracking-[0.2em] text-[#C5BFB5] px-10 opacity-70">
            Agotado
          </button>
        </section>
      </div>
    </template>

    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
  </div>
</template>
