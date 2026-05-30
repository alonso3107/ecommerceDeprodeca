<script setup lang="ts">
definePageMeta({ layout: "default" });
// Sin imports manuales — Nuxt auto-importa desde features/home/components/
// Los componentes Lazy* son manejados por Nuxt automáticamente
</script>

<template>
    <div>
        <HomeHero />

        <LazyHomeBeneficios />

        <LazyHomeCategorias />

        <LazyHomeGamificacion />

        <LazyHomeCtaFinal />
    </div>
</template>
