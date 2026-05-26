<script setup lang="ts">
definePageMeta({
  layout: "admim",
  middleware: "admin",
})
</script>

<template>
  <div class="page-enter">
    <h1 class="font-display text-display-lg text-text-primary uppercase mb-6">Dashboard Admin</h1>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <div class="p-6 rounded-xl border-2 border-border-base bg-white">
        <p class="font-body text-caption text-text-muted uppercase tracking-wide">Pedidos Hoy</p>
        <p class="font-display text-display-md text-text-primary mt-2">12</p>
      </div>
      <div class="p-6 rounded-xl border-2 border-border-base bg-white">
        <p class="font-body text-caption text-text-muted uppercase tracking-wide">Productos</p>
        <p class="font-display text-display-md text-text-primary mt-2">500+</p>
      </div>
      <div class="p-6 rounded-xl border-2 border-border-base bg-white">
        <p class="font-body text-caption text-text-muted uppercase tracking-wide">Bodegas</p>
        <p class="font-display text-display-md text-text-primary mt-2">200+</p>
      </div>
      <div class="p-6 rounded-xl border-2 border-border-base bg-white">
        <p class="font-body text-caption text-text-muted uppercase tracking-wide">Ventas Mes</p>
        <p class="font-display text-display-md text-text-primary mt-2">S/ 45K</p>
      </div>
    </div>
  </div>
</template>
