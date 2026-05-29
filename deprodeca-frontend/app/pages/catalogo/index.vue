<!-- catalogo/index.vue — Fresh Product Gallery Flat Design -->
<script setup lang="ts">
import { resolverImagenProducto } from '~/shared/utils/imagenes-productos'

definePageMeta({ layout: 'default' })

const route = useRoute()
const config = useRuntimeConfig()

const productos = ref<any[]>([])
const categorias = ref<any[]>([])
const loading = ref(true)
const pagina = ref(1)

const categoriaSlug = computed(() => (route.query.categoria as string) || '')
const queryBusqueda = computed(() => (route.query.q as string) || '')

const productosConImagen = computed(() => {
  const imagenesUsadas = new Set<string>()
  return productos.value.map((p, i) => ({
    ...p,
    imagen_resuelta: resolverImagenProducto(p, i, imagenesUsadas),
  }))
})

onMounted(async () => {
  try {
    const catRes = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/categorias`)
    if (catRes.success) categorias.value = catRes.data
  } catch { /* silencioso */ }
  await cargarProductos()
})

watch(() => route.query, () => {
  pagina.value = 1
  cargarProductos()
})

async function cargarProductos() {
  loading.value = true
  try {
    const params = new URLSearchParams()
    params.set('pagina', String(pagina.value))
    params.set('limite', '12')
    if (categoriaSlug.value) params.set('categoria_id', categoriaSlug.value)
    if (queryBusqueda.value) params.set('q', queryBusqueda.value)

    const endpoint = queryBusqueda.value
      ? `${config.public.apiBase}/productos/buscar?${params}`
      : `${config.public.apiBase}/productos?${params}`

    const res = await $fetch<{ success: boolean; data: any[] }>(endpoint)
    productos.value = res.success ? res.data : []
  } catch {
    productos.value = []
  } finally {
    loading.value = false
  }
}

function filtrar(slug: string) {
  navigateTo(slug ? `/catalogo?categoria=${slug}` : '/catalogo')
}

function buscar(query: string) {
  navigateTo(query ? `/catalogo?q=${encodeURIComponent(query)}` : '/catalogo')
}
</script>

<template>
  <div class="max-w-[1280px] mx-auto px-6 md:px-8 py-12 md:py-16">
    <!-- Header -->
    <header class="bg-[#F5F0E8] py-16 md:py-24 text-center mb-10">
      <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#A16207] mb-4">
        Catálogo / Productos
      </p>
      <h1 class="font-display text-[clamp(3rem,8vw,6rem)] font-black text-[#1C1917] uppercase leading-[0.9]">
        {{ queryBusqueda ? `"${queryBusqueda}"` : 'Productos' }}
      </h1>
      <p class="font-body text-[16px] text-[#5C554D] mt-4 max-w-[50ch] mx-auto leading-relaxed">
        {{ queryBusqueda ? `Resultados de búsqueda` : 'Calidad Nestlé para tu bodega' }}
      </p>
    </header>

    <!-- Filtros -->
    <FiltroCategoria
      :categorias="categorias"
      :categoria-activa="categoriaSlug"
      :buscando="!!queryBusqueda"
      :query-texto="queryBusqueda"
      :resultados="productos.length"
      @filtrar="filtrar"
      @buscar="buscar"
    />

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-24 gap-4">
      <div class="w-8 h-8 border-2 border-[#1C1917] border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-[#5C554D] uppercase tracking-[0.2em]">Cargando productos</span>
    </div>

    <!-- Sin resultados -->
    <div v-else-if="productosConImagen.length === 0" class="text-center py-24">
      <p class="font-display text-[clamp(2.5rem,7vw,5rem)] text-[#C5BFB5] uppercase leading-none mb-6">
        Sin resultados
      </p>
      <p class="font-body text-[16px] text-[#5C554D] mb-8">
        No encontramos productos con esos filtros
      </p>
      <button
        class="bg-[#1C1917] text-[#FAFAF9] font-mono text-[11px] uppercase tracking-[0.2em] px-6 py-3 hover:bg-[#A16207] transition-colors duration-300"
        @click="navigateTo('/catalogo')"
      >
        Ver todo
      </button>
    </div>

    <!-- Grid de productos -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <ProductoCardPremium
        v-for="(p, i) in productosConImagen"
        :key="p.id"
        :slug="p.slug"
        :nombre="p.nombre"
        :precio="p.precio"
        :unidad="p.unidad"
        :imagen-url="p.imagen_resuelta"
        :categoria="p.categoria_nombre || 'General'"
        :stock="p.stock"
        :indice="i"
      />
    </div>

    <!-- Paginación -->
    <div v-if="productos.length >= 12" class="mt-14 flex justify-center">
      <div class="flex items-stretch gap-2">
        <button
          class="font-mono text-[11px] uppercase tracking-[0.15em] px-6 py-3 border border-[#C5BFB5] text-[#5C554D] hover:border-[#1C1917] hover:text-[#1C1917] transition-colors duration-300 disabled:opacity-30 disabled:cursor-not-allowed"
          :disabled="pagina <= 1"
          @click="pagina--; cargarProductos()"
        >
          Anterior
        </button>
        <span class="bg-[#1C1917] text-[#FAFAF9] font-mono text-[11px] font-bold px-8 py-3">
          {{ pagina }}
        </span>
        <button
          class="font-mono text-[11px] uppercase tracking-[0.15em] px-6 py-3 border border-[#C5BFB5] text-[#5C554D] hover:border-[#1C1917] hover:text-[#1C1917] transition-colors duration-300"
          @click="pagina++; cargarProductos()"
        >
          Siguiente
        </button>
      </div>
    </div>
  </div>
</template>
