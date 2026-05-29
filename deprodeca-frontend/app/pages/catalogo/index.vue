<script setup lang="ts">
import FiltroCategoriaPremium from "~/features/catalogo/components/FiltroCategoriaPremium.vue"
import ProductoCardPremium from "~/features/catalogo/components/ProductoCardPremium.vue"
import { resolverImagenProducto } from "~/shared/utils/imagenes-productos"

definePageMeta({ layout: "default" })

const route = useRoute()
const config = useRuntimeConfig()

const productos = ref<any[]>([])
const categorias = ref<any[]>([])
const loading = ref(true)
const pagina = ref(1)

const categoriaSlug = computed(() => (route.query.categoria as string) || "")
const queryBusqueda = computed(() => (route.query.q as string) || "")

const productosConImagen = computed(() => {
  const imagenesUsadas = new Set<string>()
  return productos.value.map((producto, indice) => ({
    ...producto,
    imagen_resuelta: resolverImagenProducto(producto, indice, imagenesUsadas),
  }))
})

onMounted(async () => {
  try {
    const catRes = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/categorias`)
    if (catRes.success) categorias.value = catRes.data
  } catch {
    categorias.value = []
  }
  await cargarProductos()
})

watch(
  () => route.query,
  () => {
    pagina.value = 1
    cargarProductos()
  },
)

async function cargarProductos() {
  loading.value = true
  try {
    const params = new URLSearchParams()
    params.set("pagina", String(pagina.value))
    params.set("limite", "12")
    if (categoriaSlug.value) params.set("categoria_id", categoriaSlug.value)
    if (queryBusqueda.value) params.set("q", queryBusqueda.value)

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
  if (slug) navigateTo(`/catalogo?categoria=${slug}`)
  else navigateTo("/catalogo")
}

function buscar(query: string) {
  if (query) navigateTo(`/catalogo?q=${encodeURIComponent(query)}`)
  else navigateTo("/catalogo")
}
</script>

<template>
  <main class="max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <header class="bg-blanco py-20 md:py-28 mb-10">
      <p class="font-mono text-[11px] uppercase tracking-[0.4em] text-dorado">Catalogo Premium</p>
      <h1 class="mt-5 font-display text-[clamp(4rem,10vw,8rem)] font-black text-negro uppercase leading-[0.85] tracking-[-0.03em]">Productos</h1>
      <p class="font-body text-lg text-stone-oscuro max-w-[50ch] mt-6 leading-relaxed">
        {{ queryBusqueda ? `Resultados para "${queryBusqueda}"` : "Descubre un catalogo de productos con seleccion premium para bodegas modernas." }}
      </p>
      <div v-if="queryBusqueda || categoriaSlug" class="mt-6 flex flex-wrap gap-2">
        <span v-if="queryBusqueda" class="bg-crema px-4 py-2 font-mono text-[11px] uppercase tracking-[0.18em] text-stone-oscuro">Busqueda: {{ queryBusqueda }}</span>
        <span v-if="categoriaSlug" class="bg-crema px-4 py-2 font-mono text-[11px] uppercase tracking-[0.18em] text-stone-oscuro">Categoria: {{ categoriaSlug }}</span>
      </div>
    </header>

    <FiltroCategoriaPremium
      :categorias="categorias"
      :categoria-activa="categoriaSlug"
      :buscando="!!queryBusqueda"
      :query-texto="queryBusqueda"
      :resultados="productos.length"
      @filtrar="filtrar"
      @buscar="buscar"
    />

    <div v-if="loading" class="flex flex-col items-center justify-center py-32 gap-4">
      <div class="w-8 h-8 border-2 border-negro border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-stone-oscuro uppercase tracking-[0.2em]">Cargando productos</span>
    </div>

    <div v-else-if="productosConImagen.length === 0" class="border border-stone bg-crema px-6 py-24 text-center">
      <p class="font-display text-[clamp(2.5rem,7vw,5rem)] font-black uppercase leading-none text-negro">Sin resultados</p>
      <p class="mx-auto mt-5 max-w-[52ch] font-body text-[16px] leading-relaxed text-stone-oscuro">No encontramos productos con esos filtros.</p>
      <button class="mt-8 border border-negro bg-negro px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-blanco transition-colors duration-300 hover:bg-dorado hover:text-negro" @click="navigateTo('/catalogo')">
        Ver todo
      </button>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
      <ProductoCardPremium
        v-for="p in productosConImagen"
        :key="p.id"
        :slug="p.slug"
        :nombre="p.nombre"
        :precio="p.precio"
        :unidad="p.unidad"
        :imagen-url="p.imagen_resuelta"
        :categoria="p.categoria_nombre || 'General'"
        :stock="p.stock"
      />
    </div>

    <div v-if="productos.length >= 12" class="mt-14 flex items-center justify-center">
      <div class="flex items-stretch">
        <button
          class="border border-stone px-8 py-4 font-mono text-[11px] uppercase tracking-[0.3em] text-stone-oscuro transition-colors duration-300 hover:border-negro hover:text-negro disabled:opacity-30 disabled:cursor-not-allowed"
          :disabled="pagina <= 1"
          @click="pagina--; cargarProductos()"
        >
          Anterior
        </button>
        <span class="bg-negro text-blanco px-8 py-4 font-mono text-[11px] uppercase tracking-[0.2em]">Pagina {{ pagina }}</span>
        <button class="border border-stone px-8 py-4 font-mono text-[11px] uppercase tracking-[0.3em] text-stone-oscuro transition-colors duration-300 hover:border-negro hover:text-negro" @click="pagina++; cargarProductos()">
          Siguiente
        </button>
      </div>
    </div>
  </main>
</template>
