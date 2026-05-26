<!--
  ═══════════════════════════════════════════════════════════════
  index.vue — Página de Inicio · DEPRODECA
  ─────────────────────────────────────────────────────────────
  Orquestador modular. Cada sección es un componente independiente
  en features/home/components/. Nuxt los auto-importa.

  • HomeHero → carga INMEDIATA (above-the-fold, LCP prioritario)
  • LazyHomeBeneficios → lazy-load al hacer scroll
  • LazyHomeCategorias → lazy-load al hacer scroll
  • LazyHomeGamificacion → lazy-load al hacer scroll
  • LazyHomeCtaFinal → lazy-load al hacer scroll

  El prefijo "Lazy" es nativo de Nuxt: cada componente se carga
  bajo demanda sin JS adicional en el bundle inicial.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })
// Sin imports manuales — Nuxt auto-importa desde features/home/components/
// Los componentes Lazy* son manejados por Nuxt automáticamente
</script>

<template>
  <div>
    <!--
      ═══════════════════════════════════════════════════════
      SECCIÓN 01 · HERO — Carga inmediata.
      Above-the-fold, sin lazy-load. Lo primero que ve
      el usuario. Prioridad máxima de renderizado.
      ═══════════════════════════════════════════════════════
    -->
    <HomeHero />

    <!--
      ═══════════════════════════════════════════════════════
      SECCIÓN 02 · BENEFICIOS — Lazy-load nativo Nuxt.
      ═══════════════════════════════════════════════════════
    -->
    <LazyHomeBeneficios />

    <!--
      ═══════════════════════════════════════════════════════
      SECCIÓN 03 · CATEGORÍAS
      ═══════════════════════════════════════════════════════
    -->
    <LazyHomeCategorias />

    <!--
      ═══════════════════════════════════════════════════════
      SECCIÓN 04 · GAMIFICACIÓN
      ═══════════════════════════════════════════════════════
    -->
    <LazyHomeGamificacion />

    <!--
      ═══════════════════════════════════════════════════════
      SECCIÓN 05 · CTA FINAL
      ═══════════════════════════════════════════════════════
    -->
    <LazyHomeCtaFinal />
  </div>
</template>
