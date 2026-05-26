<!--
  ═══════════════════════════════════════════════════════════════
  ProductoCardBrutal.vue — Card de producto estilo brutalista
  Sin bordes redondeados, sin sombras, sin transforms.
  Datos crudos, tipografía como elemento estructural.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Props ────────────────────────────────────────────────
defineProps<{
  slug: string
  nombre: string
  precio: number
  unidad: string
  imagenUrl: string
  categoria: string
  stock: number
}>()

// ─── Utilidades ───────────────────────────────────────────
const PLACEHOLDER = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='400' height='300'%3E%3Crect fill='%23FAFAFA' width='400' height='300'/%3E%3Ctext x='200' y='155' text-anchor='middle' font-size='14' font-family='monospace' fill='%23A3A3A3'%3EDEPRODECA%3C/text%3E%3C/svg%3E"

function formatearPrecio(precio: number): string {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    CARD BRUTALISTA
    • Sin border-radius
    • Sin sombras
    • Sin scale en hover
    • Borde duro como único indicador visual
    • Datos first: precio masivo, categoría monospace
    ═══════════════════════════════════════════════════════════
  -->
  <NuxtLink
    :to="`/catalogo/${slug}`"
    class="group block bg-white border border-borde
           hover:border-[#D4A017] hover:bg-[#D4A017]/[0.02]
           transition-all duration-200"
  >

    <!-- Imagen · Rectangular, sin curvas, sin zoom -->
    <div class="relative aspect-[4/3] overflow-hidden bg-fondo border-b border-borde">
      <img
        :src="imagenUrl"
        :alt="`Producto: ${nombre}`"
        class="w-full h-full object-cover"
        loading="lazy"
        @error="(e: Event) => (e.target as HTMLImageElement).src = PLACEHOLDER"
      />

      <!-- Overlay AGOTADO · Brutalista: franja negra con texto blanco -->
      <div
        v-if="stock === 0"
        class="absolute inset-0 bg-black/50 flex items-center justify-center"
      >
        <span class="font-display text-display-md text-white uppercase tracking-widest">
          AGOTADO
        </span>
      </div>
    </div>

    <!-- Datos del producto -->
    <div class="p-5 flex flex-col gap-3">

      <!-- Categoría · Monospace técnico -->
      <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.15em]">
        {{ categoria }}
      </span>

      <!-- Nombre · Display condensado -->
      <h3 class="font-body text-body font-bold text-texto leading-snug line-clamp-2">
        {{ nombre }}
      </h3>

      <!-- Precio · Masivo, brutalista -->
      <div class="flex items-baseline gap-1.5 mt-1">
        <span class="font-display text-[clamp(1.5rem,3vw,2rem)] text-texto leading-none">
          {{ formatearPrecio(precio) }}
        </span>
        <span class="font-mono text-[10px] text-texto-muted tracking-wider">
          /{{ unidad }}
        </span>
      </div>

      <!-- Stock · Indicador técnico -->
      <div class="flex items-center gap-2 mt-1">
        <span
          class="w-2 h-2"
          :class="stock > 0 ? 'bg-exito' : 'bg-error'"
          aria-hidden="true"
        />
        <span class="font-mono text-[10px] uppercase tracking-wider"
              :class="stock > 0 ? 'text-exito' : 'text-error'">
          {{ stock > 0 ? `${stock} en stock` : 'Agotado' }}
        </span>
      </div>

    </div>
  </NuxtLink>
</template>
