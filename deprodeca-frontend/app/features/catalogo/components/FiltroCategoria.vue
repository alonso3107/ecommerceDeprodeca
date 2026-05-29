<!-- FiltroCategoria.vue — Flat Design Fresh Gallery -->
<script setup lang="ts">
const props = defineProps<{
  categorias: { slug: string; nombre: string }[]
  categoriaActiva: string
  buscando: boolean
  queryTexto: string
  resultados: number
}>()

const emit = defineEmits<{
  (e: 'filtrar', slug: string): void
  (e: 'buscar', query: string): void
}>()

const busquedaLocal = ref('')

const placeholderTexto = computed(() =>
  props.buscando ? `Resultados: "${props.queryTexto}"` : 'Buscar productos...'
)

function ejecutarBusqueda() {
  emit('buscar', busquedaLocal.value.trim())
}

function limpiarFiltros() {
  busquedaLocal.value = ''
  emit('filtrar', '')
}

// Animación de entrada de los pills
const pillsVisibles = ref(false)
onMounted(() => { pillsVisibles.value = true })
</script>

<template>
  <div class="space-y-6 mb-10">
    <!-- Pills de categorías -->
    <div class="flex flex-wrap justify-center gap-2">
      <button
        v-for="(cat, i) in [{ slug: '', nombre: 'Todas' }, ...categorias]"
        :key="cat.slug"
        class="font-mono text-[11px] uppercase tracking-[0.1em] px-5 py-2.5 border transition-all duration-300 ease-out"
        :class="[
          pillsVisibles ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2',
          categoriaActiva === cat.slug || (!categoriaActiva && cat.slug === '')
            ? 'bg-[#1C1917] border-[#1C1917] text-[#FAFAF9] font-bold'
            : 'bg-white border-[#C5BFB5] text-[#5C554D] hover:border-[#1C1917] hover:text-[#1C1917]'
        ]"
        :style="{ transitionDelay: `${i * 50}ms` }"
        @click="emit('filtrar', cat.slug)"
      >
        {{ cat.nombre }}
      </button>
    </div>

    <!-- Barra de búsqueda -->
    <form class="flex justify-center" @submit.prevent="ejecutarBusqueda">
      <div class="relative w-full max-w-[400px]">
        <!-- Icono search geométrico -->
        <svg class="absolute left-1 top-1/2 -translate-y-1/2 text-[#C5BFB5]" width="16" height="16" viewBox="0 0 16 16" fill="none">
          <circle cx="6.5" cy="6.5" r="4.5" stroke="currentColor" stroke-width="1.5"/>
          <path d="M10 10L14 14" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
        </svg>
        <input
          v-model="busquedaLocal"
          type="search"
          :placeholder="placeholderTexto"
          class="w-full border-b-2 border-[#C5BFB5] bg-transparent pl-8 pr-4 py-3 font-body text-[14px] text-[#1C1917] placeholder:text-[#C5BFB5] focus:border-[#A16207] focus:outline-none transition-colors duration-300"
        />
      </div>
    </form>

    <!-- Info: filtro activo + contador -->
    <div v-if="buscando || categoriaActiva" class="flex items-center justify-center gap-4">
      <span class="font-mono text-[10px] text-[#5C554D] uppercase tracking-[0.15em]">
        Filtro: <strong class="text-[#A16207]">{{ buscando ? `"${queryTexto}"` : categoriaActiva }}</strong>
      </span>
      <button class="font-mono text-[10px] text-[#5C554D] hover:text-[#DC2626] uppercase transition-colors duration-200" @click="limpiarFiltros">
        ✕ Limpiar
      </button>
    </div>

    <p class="text-center font-mono text-[11px] text-[#5C554D] uppercase tracking-[0.15em]">
      {{ resultados }} productos
    </p>
  </div>
</template>
