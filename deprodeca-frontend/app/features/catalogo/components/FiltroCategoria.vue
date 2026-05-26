<!--
  ═══════════════════════════════════════════════════════════════
  FiltroCategoria.vue — Barra de filtros horizontal brutalista
  Pills toggle con bordes duros. Sin rounded-full, sin sombras.
  Alto contraste: activo = fondo negro + texto blanco.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Props ────────────────────────────────────────────────
const props = defineProps<{
  categorias: { slug: string; nombre: string }[]
  categoriaActiva: string
  buscando: boolean
  queryTexto: string
}>()

// ─── Emits ────────────────────────────────────────────────
const emit = defineEmits<{
  (e: "filtrar", slug: string): void
  (e: "buscar", query: string): void
}>()

// ─── Estado ───────────────────────────────────────────────
const busquedaLocal = ref("")

// ─── Computed ─────────────────────────────────────────────
const placeholderTexto = computed(() =>
  props.buscando ? `Resultados: "${props.queryTexto}"` : "Buscar producto..."
)

// ─── Métodos ──────────────────────────────────────────────
function ejecutarBusqueda() {
  emit("buscar", busquedaLocal.value.trim())
}
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    FILTROS · Barra horizontal monolítica.
    Cada categoría es un bloque rectangular con borde duro.
    Sin transiciones suaves: el toggle es inmediato y crudo.
    ═══════════════════════════════════════════════════════════
  -->
  <div class="space-y-6">

    <!-- Barra de búsqueda · Input brutalista -->
    <form class="flex gap-0" @submit.prevent="ejecutarBusqueda">
      <input
        v-model="busquedaLocal"
        type="search"
        :placeholder="placeholderTexto"
        class="flex-1 border border-borde px-4 py-3 font-body text-body text-texto
               placeholder:text-texto-muted bg-white
               focus:border-[#D4A017] focus:outline-none
               transition-colors duration-200 min-h-[48px]"
      />
      <button
        type="submit"
        class="border border-l-0 border-borde bg-texto text-white px-6 py-3
               font-mono text-[11px] uppercase tracking-[0.15em]
               hover:bg-[#D4A017] hover:text-black hover:border-[#D4A017]
               transition-colors duration-200 min-h-[48px]"
      >
        Buscar
      </button>
    </form>

    <!-- Pills de categorías · Bloques duros -->
    <div class="flex flex-wrap gap-0">
      <!-- "Todas" siempre presente -->
      <button
        class="border border-borde px-5 py-2.5 font-mono text-[11px] uppercase tracking-[0.15em]
               transition-colors duration-150 min-h-[42px]"
        :class="!categoriaActiva
          ? 'bg-texto text-white border-texto font-bold'
          : 'bg-white text-texto-muted hover:text-texto hover:border-texto/30'"
        @click="emit('filtrar', '')"
      >
        Todas
      </button>

      <!-- Categorías dinámicas -->
      <button
        v-for="cat in categorias"
        :key="cat.slug"
        class="border border-l-0 border-borde px-5 py-2.5 font-mono text-[11px] uppercase tracking-[0.15em]
               transition-colors duration-150 min-h-[42px]"
        :class="categoriaActiva === cat.slug
          ? 'bg-texto text-white border-texto font-bold'
          : 'bg-white text-texto-muted hover:text-texto hover:border-texto/30'"
        @click="emit('filtrar', cat.slug)"
      >
        {{ cat.nombre }}
      </button>
    </div>

    <!-- Badge de resultados · Indicador técnico -->
    <div v-if="buscando || categoriaActiva" class="flex items-center gap-2">
      <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em]">
        Filtro activo:
      </span>
      <span class="font-mono text-[10px] text-[#D4A017] uppercase tracking-[0.15em] font-bold">
        {{ buscando ? `"${queryTexto}"` : categoriaActiva }}
      </span>
      <button
        class="font-mono text-[10px] text-texto-muted hover:text-error uppercase tracking-wider ml-2"
        @click="busquedaLocal = ''; emit('filtrar', '')"
      >
        ✕ Limpiar
      </button>
    </div>

  </div>
</template>
