<!--
  ═══════════════════════════════════════════════════════════════
  catalogo/index.vue — Catálogo de Productos · DEPRODECA
  Brutalismo funcional: grid gap-px, filtros como bloques,
  paginación cruda, sin adornos innecesarios.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })

const route = useRoute()
const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const productos = ref<any[]>([])
const categorias = ref<any[]>([])
const loading = ref(true)
const pagina = ref(1)
const categoriaSlug = computed(() => (route.query.categoria as string) || "")
const queryBusqueda = computed(() => (route.query.q as string) || "")

// ─── Ciclo de vida ────────────────────────────────────────
onMounted(async () => {
  try {
    const catRes = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/categorias`)
    if (catRes.success) categorias.value = catRes.data
  } catch (_) { /* fallback silencioso */ }
  await cargarProductos()
})

// Reaccionar a cambios en query params (filtros, búsqueda)
watch(() => route.query, () => {
  pagina.value = 1
  cargarProductos()
})

// ─── Lógica de carga ──────────────────────────────────────
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
    if (res.success) productos.value = res.data
    else productos.value = []
  } catch (_) {
    productos.value = []
  } finally {
    loading.value = false
  }
}

// ─── Navegación ───────────────────────────────────────────
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
  <!--
    ═══════════════════════════════════════════════════════════
    CATÁLOGO — Layout brutalista.
    Sin sidebar, filtros horizontales como barra de herramientas.
    Grid de productos con gap-px (píxel de borde compartido).
    ═══════════════════════════════════════════════════════════
  -->
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">

    <!-- ─── Encabezado con metadata técnica ──────────────── -->
    <div class="mb-10">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        Catálogo
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        {{ queryBusqueda ? `"${queryBusqueda}"` : "Productos" }}
      </h1>
      <p class="font-body text-body text-texto-muted mt-2">
        {{ queryBusqueda ? 'Resultados de búsqueda' : 'Productos para tu bodega — directo de Nestlé' }}
      </p>
    </div>

    <!-- ─── Filtros · Barra horizontal brutalista ────────── -->
    <FiltroCategoria
      :categorias="categorias"
      :categoria-activa="categoriaSlug"
      :buscando="!!queryBusqueda"
      :query-texto="queryBusqueda"
      @filtrar="filtrar"
      @buscar="buscar"
    />

    <!-- ─── Área de contenido ────────────────────────────── -->
    <div class="mt-10">

      <!-- LOADING · Spinner mínimo, sin cards fantasma -->
      <div v-if="loading" class="flex flex-col items-center justify-center py-32 gap-4">
        <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
        <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
      </div>

      <!-- VACÍO · Tipografía masiva, sin adornos -->
      <div v-else-if="productos.length === 0" class="text-center py-32">
        <p class="font-display text-[clamp(3rem,8vw,6rem)] text-texto/10 uppercase leading-none mb-6">
          Sin Resultados
        </p>
        <p class="font-body text-body text-texto-muted mb-8">
          No encontramos productos con esos filtros.
        </p>
        <button
          class="font-mono text-[11px] text-texto-muted hover:text-[#D4A017] uppercase tracking-[0.2em]
                 border border-borde px-6 py-3 hover:border-[#D4A017] transition-colors duration-200"
          @click="navigateTo('/catalogo')"
        >
          Ver Todo
        </button>
      </div>

      <!-- GRID · gap-px brutalista, sin espacios entre cards -->
      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-px bg-borde">
        <ProductoCardBrutal
          v-for="p in productos"
          :key="p.id"
          :slug="p.slug"
          :nombre="p.nombre"
          :precio="p.precio"
          :unidad="p.unidad"
          :imagen-url="p.imagen_url || 'https://images.unsplash.com/photo-1607622744071-ac94bf8b8c6b?w=400&q=80'"
          :categoria="p.categoria_nombre || 'General'"
          :stock="p.stock"
        />
      </div>

      <!-- PAGINACIÓN · Bloques duros, sin bordes redondeados -->
      <div v-if="productos.length >= 12" class="flex justify-center items-center gap-0 mt-14">
        <button
          class="border border-borde px-5 py-3 font-mono text-[11px] uppercase tracking-[0.15em]
                 text-texto-muted hover:text-texto hover:border-texto/30
                 disabled:opacity-30 disabled:cursor-not-allowed
                 transition-colors duration-150 min-h-[48px]"
          :disabled="pagina <= 1"
          @click="pagina--; cargarProductos()"
        >
          Anterior
        </button>

        <span class="border-t border-b border-borde px-8 py-3 font-mono text-[11px] text-texto
                     uppercase tracking-[0.2em] font-bold min-h-[48px] flex items-center">
          Página {{ pagina }}
        </span>

        <button
          class="border border-borde px-5 py-3 font-mono text-[11px] uppercase tracking-[0.15em]
                 text-texto-muted hover:text-texto hover:border-texto/30
                 transition-colors duration-150 min-h-[48px]"
          @click="pagina++; cargarProductos()"
        >
          Siguiente
        </button>
      </div>

    </div>
  </div>
</template>
