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

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}
</script>

<template>
  <NuxtLink :to="`/catalogo/${slug}`" class="group block bg-blanco border border-stone hover:border-dorado transition-colors duration-500">
    <div class="relative overflow-hidden aspect-[5/6] bg-crema">
      <img :src="imagenUrl" :alt="nombre" class="w-full h-full object-cover transition-transform duration-700 ease-out group-hover:scale-[1.03]" loading="lazy" />
      <div class="absolute inset-x-0 bottom-0 h-1/2 bg-gradient-to-t from-black/40 to-black/0" />
      <p class="absolute inset-x-0 bottom-0 p-6 font-display text-2xl font-black text-blanco uppercase tracking-tight leading-none">
        {{ nombre }}
      </p>
      <div v-if="stock === 0" class="absolute inset-0 flex items-center justify-center bg-black/45">
        <span class="font-display text-5xl font-black text-blanco/80 uppercase tracking-tight">Agotado</span>
      </div>
    </div>

    <div class="p-6">
      <p class="font-mono text-[10px] uppercase text-stone-oscuro tracking-[0.2em]">{{ categoria }}</p>
      <div class="mt-3 flex items-end gap-2">
        <p class="font-display text-4xl font-black text-negro leading-none">{{ formatearPrecio(precio) }}</p>
        <p class="font-mono text-[11px] text-stone-oscuro pb-1">/{{ unidad }}</p>
      </div>
      <div class="mt-4 flex items-center gap-2">
        <span class="w-3 h-3" :class="stock > 0 ? 'bg-exito' : 'bg-error'" />
        <span class="font-mono text-[10px] uppercase tracking-[0.15em]" :class="stock > 0 ? 'text-exito' : 'text-error'">
          {{ stock > 0 ? `${stock} disponibles` : "Agotado" }}
        </span>
      </div>
    </div>
  </NuxtLink>
</template>
