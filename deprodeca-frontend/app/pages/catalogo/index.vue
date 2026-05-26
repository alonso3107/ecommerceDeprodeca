<script setup lang="ts">
definePageMeta({ layout: "default" })

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
    const catRes = await $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/categorias`)
    if (catRes.success) categorias.value = catRes.data
  } catch (_) {}
  await cargarProductos()
})

watch(() => route.query, () => { pagina.value = 1; cargarProductos() })

async function cargarProductos() {
  loading.value = true
  try {
    const params = new URLSearchParams()
    params.set("pagina", String(pagina.value))
    params.set("limite", "12")
    if (categoriaSlug.value) params.set("categoria_id", categoriaSlug.value)
    if (queryBusqueda.value) params.set("q", queryBusqueda.value)
    const endpoint = queryBusqueda.value ? `${config.public.apiBase}/productos/buscar?${params}` : `${config.public.apiBase}/productos?${params}`
    const res = await $fetch<{ success: boolean; data: any[] }>(endpoint)
    if (res.success) productos.value = res.data
  } catch (_) { productos.value = [] }
  finally { loading.value = false }
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">
    <div class="mb-12">
      <h1 class="font-display text-display-lg text-texto uppercase mb-2">{{ queryBusqueda ? `"${queryBusqueda}"` : "Catálogo" }}</h1>
      <p class="font-body text-body text-texto-muted">Productos para tu bodega</p>
    </div>

    <div class="flex flex-col lg:flex-row gap-10">
      <aside class="lg:w-48 flex-shrink-0">
        <h3 class="font-body text-caption font-bold text-texto uppercase tracking-wider mb-4">Categorías</h3>
        <nav class="space-y-1">
          <button class="w-full text-left px-3 py-2 rounded-lg font-body text-small transition-colors duration-300 min-h-[40px]" :class="!categoriaSlug ? 'bg-texto text-white font-semibold' : 'text-texto-muted hover:bg-fondo hover:text-texto'" @click="navigateTo('/catalogo')">Todas</button>
          <button v-for="cat in categorias" :key="cat.slug" class="w-full text-left px-3 py-2 rounded-lg font-body text-small transition-colors duration-300 min-h-[40px]" :class="categoriaSlug === cat.slug ? 'bg-texto text-white font-semibold' : 'text-texto-muted hover:bg-fondo hover:text-texto'" @click="navigateTo(`/catalogo?categoria=${cat.slug}`)">{{ cat.nombre }}</button>
        </nav>
      </aside>

      <div class="flex-1">
        <div v-if="loading" class="flex justify-center py-20"><Spinner size="lg" label="Cargando..." /></div>

        <div v-else-if="productos.length === 0" class="text-center py-20">
          <p class="font-display text-display-md text-texto-muted mb-4">Sin resultados</p>
          <Button variant="outline" size="md" @click="navigateTo('/catalogo')">Ver Todo</Button>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <ProductCard v-for="p in productos" :key="p.id" :slug="p.slug" :nombre="p.nombre" :precio="p.precio" :unidad="p.unidad" :imagen-url="p.imagen_url || 'https://images.unsplash.com/photo-1607622744071-ac94bf8b8c6b?w=400&q=80'" :categoria="p.categoria_nombre || 'General'" :stock="p.stock" />
        </div>

        <div v-if="productos.length >= 12" class="flex justify-center mt-12 gap-3">
          <Button variant="outline" size="sm" :disabled="pagina <= 1" @click="pagina--; cargarProductos()">Anterior</Button>
          <span class="flex items-center px-4 font-body text-small text-texto-muted">Página {{ pagina }}</span>
          <Button variant="outline" size="sm" @click="pagina++; cargarProductos()">Siguiente</Button>
        </div>
      </div>
    </div>
  </div>
</template>
