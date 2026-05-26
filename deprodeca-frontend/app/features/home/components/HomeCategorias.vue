<!--
  ═══════════════════════════════════════════════════════════════
  HomeCategorias.vue — Grid de categorías de producto
  Brutalismo contenido: etiquetas sharp, hover mínimo, enfoque
  utilitario.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Datos estáticos ──────────────────────────────────────
const categorias = [
  { nombre: "Bebidas",          slug: "bebidas",           count: "120+" },
  { nombre: "Snacks",           slug: "snacks",            count: "85+" },
  { nombre: "Lácteos",          slug: "lacteos",           count: "60+" },
  { nombre: "Abarrotes",        slug: "abarrotes",         count: "150+" },
  { nombre: "Limpieza",         slug: "limpieza",          count: "45+" },
  { nombre: "Cuidado Personal", slug: "cuidado-personal",  count: "40+" },
] as const
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    CATEGORÍAS — Grid utilitario. Brutalismo: sin íconos
    decorativos, tipografía como elemento estructural,
    contadores de producto como dato crudo.
    ═══════════════════════════════════════════════════════════
  -->
  <section class="py-24 md:py-32 bg-fondo border-t border-borde">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">

      <!-- Encabezado técnico -->
      <div class="flex flex-col md:flex-row md:items-end justify-between mb-14 gap-4">
        <div>
          <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-4">
            ─── 02 · Categorías
          </p>
          <h2 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
            Explorá
          </h2>
        </div>
        <NuxtLink
          to="/catalogo"
          class="font-mono text-[11px] text-texto-muted hover:text-[#D4A017] uppercase tracking-widest transition-colors duration-200 border-b border-transparent hover:border-[#D4A017] pb-0.5"
        >
          Ver todas →
        </NuxtLink>
      </div>

      <!-- Grid brutalista: bordes finos, sin sombras, pura estructura -->
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-px bg-borde">
        <NuxtLink
          v-for="cat in categorias"
          :key="cat.slug"
          :to="`/catalogo?categoria=${cat.slug}`"
          class="group bg-white flex flex-col justify-between p-6 md:p-8 min-h-[160px]
                 hover:bg-[#D4A017]/5 border border-transparent hover:border-[#D4A017]/30 transition-all duration-200"
        >
          <span class="font-body text-body font-semibold text-texto uppercase tracking-wide
                       group-hover:text-[#D4A017] transition-colors duration-200">
            {{ cat.nombre }}
          </span>
          <span class="font-mono text-[10px] text-texto-muted tracking-widest mt-4">
            {{ cat.count }} items
          </span>
        </NuxtLink>
      </div>

    </div>
  </section>
</template>
