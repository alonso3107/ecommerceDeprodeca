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
  <div class="bg-[#FBF7F1]">
    <div class="max-w-[1360px] mx-auto px-6 md:px-8 py-10 md:py-14 pb-20">
      <!-- Header -->
      <header class="mb-10 rounded-[32px] border border-[#E7E0D5] bg-white p-6 md:p-10 shadow-[0_28px_90px_-60px_rgba(28,25,23,0.45)] lg:grid lg:grid-cols-[1.5fr_.85fr] lg:gap-8 lg:items-end">
        <div>
          <p class="mb-4 font-mono text-[11px] uppercase tracking-[0.3em] text-[#A16207]">
            Catálogo / Productos
          </p>
          <h1 class="max-w-[12ch] font-display text-[clamp(2.8rem,7vw,5.8rem)] font-black uppercase leading-[0.92] text-[#1C1917]">
            {{ queryBusqueda ? `"${queryBusqueda}"` : 'Productos' }}
          </h1>
          <p class="mt-5 max-w-[52ch] text-[16px] leading-relaxed text-[#5C554D]">
            {{ queryBusqueda ? 'Resultados de búsqueda ajustados al filtro activo.' : 'Selección curada con imágenes locales y composición más limpia para navegar rápido.' }}
          </p>
        </div>

        <div class="mt-6 grid grid-cols-2 gap-3 lg:mt-0">
          <div class="rounded-3xl bg-[#1C1917] p-5 text-[#FAFAF9]">
            <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#D6D3D1]">Resultados</p>
            <p class="mt-3 font-display text-4xl font-black leading-none">{{ productos.length }}</p>
          </div>
          <div class="rounded-3xl border border-[#E7E0D5] bg-[#FBF7F1] p-5 text-[#1C1917]">
            <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#5C554D]">Estado</p>
            <p class="mt-3 text-[15px] font-semibold leading-tight">
              {{ loading ? 'Cargando...' : (queryBusqueda ? 'Búsqueda activa' : 'Vista general') }}
            </p>
          </div>
        </div>
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
      <div v-else-if="productosConImagen.length === 0" class="rounded-[28px] border border-dashed border-[#D6D3D1] bg-white py-24 text-center">
        <p class="mb-6 font-display text-[clamp(2.5rem,7vw,5rem)] uppercase leading-none text-[#C5BFB5]">
          Sin resultados
        </p>
        <p class="mb-8 text-[16px] text-[#5C554D]">
          No encontramos productos con esos filtros
        </p>
        <button
          class="bg-[#1C1917] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-[#FAFAF9] transition-colors duration-300 hover:bg-[#A16207]"
          @click="navigateTo('/catalogo')"
        >
          Ver todo
        </button>
      </div>

      <!-- Grid de productos -->
      <div v-else class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
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
            class="border border-[#C5BFB5] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#5C554D] transition-colors duration-300 hover:border-[#1C1917] hover:text-[#1C1917] disabled:cursor-not-allowed disabled:opacity-30"
            :disabled="pagina <= 1"
            @click="pagina--; cargarProductos()"
          >
            Anterior
          </button>
          <span class="bg-[#1C1917] px-8 py-3 font-mono text-[11px] font-bold text-[#FAFAF9]">
            {{ pagina }}
          </span>
          <button
            class="border border-[#C5BFB5] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#5C554D] transition-colors duration-300 hover:border-[#1C1917] hover:text-[#1C1917]"
            @click="pagina++; cargarProductos()"
          >
            Siguiente
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
