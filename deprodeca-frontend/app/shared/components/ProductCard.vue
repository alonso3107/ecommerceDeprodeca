<script setup lang="ts">
// ProductCard: Tarjeta de producto para catálogo / featured
// Usa tokens de Tailwind, sin colores hex hardcodeados.

defineProps<{
  slug: string
  nombre: string
  precio: number
  unidad: string
  imagenUrl: string
  categoria: string
  stock: number
}>()

function formatearPrecio(precio: number): string {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}
</script>

<template>
  <NuxtLink
    :to="`/catalogo/${slug}`"
    class="group block bg-white rounded-xl border border-border-base overflow-hidden hover-lift"
  >
    <!-- Imagen del producto -->
    <div class="relative aspect-[4/3] overflow-hidden bg-surface">
      <img
        :src="imagenUrl"
        :alt="`Producto: ${nombre}`"
        class="w-full h-full object-cover transition-transform duration-700 ease-out group-hover:scale-[1.02]"
        loading="lazy"
      />

      <!-- Badge de categoría -->
      <span
        class="absolute top-3 left-3 px-2.5 py-1 rounded-full bg-brand-primary/90 text-text-primary text-caption font-body font-semibold"
      >
        {{ categoria }}
      </span>

      <!-- Stock bajo -->
      <span
        v-if="stock <= 10 && stock > 0"
        class="absolute top-3 right-3 px-2.5 py-1 rounded-full bg-status-warning/90 text-white text-caption font-body font-semibold"
      >
        ¡Últimas {{ stock }}!
      </span>

      <!-- Sin stock -->
      <div
        v-if="stock === 0"
        class="absolute inset-0 bg-black/50 flex items-center justify-center"
      >
        <span class="text-white font-display text-display-md opacity-80">AGOTADO</span>
      </div>
    </div>

    <!-- Info del producto -->
    <div class="p-4">
      <h3 class="font-body text-body font-semibold text-text-primary line-clamp-2 mb-1">
        {{ nombre }}
      </h3>

      <div class="flex items-baseline gap-1 mt-2">
        <span class="font-display text-heading text-text-primary">
          {{ formatearPrecio(precio) }}
        </span>
        <span class="text-small text-text-muted font-body">/ {{ unidad }}</span>
      </div>
    </div>
  </NuxtLink>
</template>
