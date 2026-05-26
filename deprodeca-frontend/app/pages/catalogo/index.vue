<script setup lang="ts">
import { ref, computed, onMounted, watch } from "vue"

definePageMeta({
  layout: "default",
})

const route = useRoute()
const config = useRuntimeConfig()

const productos = ref<any[]>([])
const categorias = ref<any[]>([])
const loading = ref(true)
const pagina = ref(1)

const categoriaSlug = computed(() => (route.query.categoria as string) || "")
const queryBusqueda = computed(() => (route.query.q as string) || "")

onMounted(async () => {
  try {
    const catRes = await $fetch<{ success: boolean; data: any[] }>(
      `${config.public.apiBase}/categorias`,
    )
    if (catRes.success) categorias.value = catRes.data
  } catch (_) {
    // categorías opcionales, continuar sin ellas
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
    if (res.success) productos.value = res.data
  } catch (_) {
    productos.value = []
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="font-display text-display-lg text-text-primary uppercase mb-2">
        {{ queryBusqueda ? `Resultados: "${queryBusqueda}"` : "Catálogo" }}
      </h1>
      <p class="font-body text-body text-text-muted">
        Explora nuestros productos para tu bodega
      </p>
    </div>

    <div class="flex flex-col lg:flex-row gap-8">
      <!-- Sidebar: Categorías -->
      <aside class="lg:w-56 flex-shrink-0">
        <h3 class="font-body text-small font-bold text-text-primary uppercase tracking-wide mb-3">
          Categorías
        </h3>
        <nav class="space-y-1">
          <button
            class="w-full text-left px-3 py-2 rounded-lg font-body text-small transition-colors duration-300"
            :class="!categoriaSlug ? 'bg-brand-primary/20 text-text-primary font-semibold' : 'text-text-muted hover:bg-surface hover:text-text-primary'"
            @click="navigateTo('/catalogo')"
          >
            Todas
          </button>
          <button
            v-for="cat in categorias"
            :key="cat.slug"
            class="w-full text-left px-3 py-2 rounded-lg font-body text-small transition-colors duration-300 min-h-[44px] flex items-center"
            :class="categoriaSlug === cat.slug ? 'bg-brand-primary/20 text-text-primary font-semibold' : 'text-text-muted hover:bg-surface hover:text-text-primary'"
            @click="navigateTo(`/catalogo?categoria=${cat.slug}`)"
          >
            {{ cat.nombre }}
          </button>
        </nav>
      </aside>

      <!-- Grid de Productos -->
      <div class="flex-1">
        <!-- Loading -->
        <div v-if="loading" class="flex justify-center py-20">
          <Spinner size="lg" label="Cargando productos..." />
        </div>

        <!-- Empty state -->
        <div
          v-else-if="productos.length === 0"
          class="text-center py-20"
        >
          <div class="w-20 h-20 mx-auto mb-4 rounded-full bg-surface flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-text-muted"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
          </div>
          <h3 class="font-body text-subheading font-bold text-text-primary mb-1">
            Sin resultados
          </h3>
          <p class="font-body text-small text-text-muted mb-6">
            Intenta con otra búsqueda o categoría.
          </p>
          <Button variant="outline" size="md" @click="navigateTo('/catalogo')">
            Ver Todo el Catálogo
          </Button>
        </div>

        <!-- Grid de productos -->
        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <ProductCard
            v-for="p in productos"
            :key="p.id"
            :slug="p.slug"
            :nombre="p.nombre"
            :precio="p.precio"
            :unidad="p.unidad"
            :imagen-url="p.imagen_url || 'https://images.unsplash.com/photo-1543168256-418811576931?w=400&q=80'"
            :categoria="p.categoria_nombre || 'General'"
            :stock="p.stock"
          />
        </div>

        <!-- Paginación -->
        <div v-if="productos.length >= 12" class="flex justify-center mt-10 gap-2">
          <Button
            variant="outline"
            size="sm"
            :disabled="pagina <= 1"
            @click="pagina--; cargarProductos()"
          >
            Anterior
          </Button>
          <span class="flex items-center px-4 font-body text-small text-text-muted">
            Página {{ pagina }}
          </span>
          <Button
            variant="outline"
            size="sm"
            @click="pagina++; cargarProductos()"
          >
            Siguiente
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>
