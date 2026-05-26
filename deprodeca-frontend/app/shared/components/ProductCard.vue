<script setup lang="ts">
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
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN", minimumFractionDigits: 2 }).format(precio)
}
</script>

<template>
  <NuxtLink
    :to="`/catalogo/${slug}`"
    class="group block bg-white rounded-2xl shadow-xs hover-lift overflow-hidden"
  >
    <div class="relative aspect-[4/3] overflow-hidden bg-fondo">
      <img
        :src="imagenUrl"
        :alt="`Producto: ${nombre}`"
        class="w-full h-full object-cover transition-transform duration-700 ease-out group-hover:scale-[1.02]"
        loading="lazy"
      />
      <span class="absolute top-3 left-3 px-2.5 py-1 rounded-lg bg-white/90 text-texto text-caption font-body font-semibold shadow-xs backdrop-blur-sm">
        {{ categoria }}
      </span>
      <span
        v-if="stock === 0"
        class="absolute inset-0 bg-black/40 flex items-center justify-center"
      >
        <span class="text-white font-display text-display-md opacity-80">AGOTADO</span>
      </span>
    </div>

    <div class="p-5">
      <h3 class="font-body text-body font-semibold text-texto line-clamp-2 mb-1 leading-snug">
        {{ nombre }}
      </h3>
      <div class="flex items-baseline gap-1 mt-2">
        <span class="font-display text-heading text-texto">{{ formatearPrecio(precio) }}</span>
        <span class="text-small text-texto-muted font-body">/ {{ unidad }}</span>
      </div>
    </div>
  </NuxtLink>
</template>
