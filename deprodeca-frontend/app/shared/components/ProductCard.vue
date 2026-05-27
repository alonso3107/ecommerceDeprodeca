<script setup lang="ts">
const props = defineProps<{
  slug: string
  nombre: string
  precio: number
  unidad: string
  imagenUrl: string
  categoria: string
  stock: number
}>()

const formatoPrecio = new Intl.NumberFormat('es-PE', {
  style: 'currency',
  currency: 'PEN',
  minimumFractionDigits: 2,
})

const fallbackImagen =
  "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='800' height='600' viewBox='0 0 800 600'%3E%3Crect width='800' height='600' fill='%23FAFAFA'/%3E%3Cpath d='M0 460H800' stroke='%23e5e5e5' stroke-width='2'/%3E%3Cpath d='M560 420L720 260L800 340' stroke='%23D4A017' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3Cpath d='M600 520L720 400L780 460' stroke='%23D4A017' stroke-opacity='.18' stroke-width='10' fill='none'/%3E%3Cpath d='M645 140H735V230Z' fill='none' stroke='%23D4A017' stroke-opacity='.14' stroke-width='8'/%3E%3Ctext x='48' y='86' fill='%23171717' font-family='monospace' font-size='26' letter-spacing='3'%3EDEPRODECA%3C/text%3E%3C/svg%3E"

function formatearPrecio(precio: number): string {
  return formatoPrecio.format(precio)
}

function manejarErrorImagen(evento: Event) {
  const img = evento.target as HTMLImageElement | null
  if (img && img.src !== fallbackImagen) {
    img.src = fallbackImagen
  }
}
</script>

<template>
  <NuxtLink
    :to="`/catalogo/${slug}`"
    class="group block border border-[#e5e5e5] bg-[#FAFAFA] transition-transform duration-700 ease-out hover:scale-[1.01]"
  >
    <div class="relative border-b border-[#e5e5e5] bg-[#F2F2F2] aspect-[4/3]">
      <img
        :src="imagenUrl"
        :alt="`Producto: ${nombre}`"
        class="block h-full w-full object-cover"
        loading="lazy"
        @error="manejarErrorImagen"
      />

      <span class="absolute left-4 top-4 border border-[#171717] bg-[#171717] px-3 py-1 font-mono text-[10px] uppercase tracking-[0.2em] text-[#FAFAFA]">
        {{ categoria }}
      </span>

      <span
        v-if="stock === 0"
        class="absolute inset-0 flex items-center justify-center bg-black/50"
      >
        <span class="font-display text-[clamp(1.75rem,4vw,3rem)] uppercase tracking-[0.18em] text-white">
          Agotado
        </span>
      </span>

      <svg
        aria-hidden="true"
        class="pointer-events-none absolute bottom-0 right-0 h-24 w-24 opacity-10"
        viewBox="0 0 96 96"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path d="M8 72H88" stroke="#D4A017" stroke-width="2" />
        <path d="M16 52L36 20L56 52L76 20" stroke="#D4A017" stroke-width="2" />
        <path d="M24 80L48 44L72 80" stroke="#D4A017" stroke-width="2" />
        <path d="M60 28L76 12L88 28" stroke="#D4A017" stroke-width="2" />
      </svg>
    </div>

    <div class="relative p-5 md:p-6 pr-14">
      <h3 class="font-body text-[16px] font-semibold leading-snug text-[#171717] line-clamp-2">
        {{ nombre }}
      </h3>

      <div class="mt-4 flex items-baseline gap-2">
        <span class="font-body text-[18px] font-semibold text-[#171717]">
          {{ formatearPrecio(precio) }}
        </span>
        <span class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#666666]">
          / {{ unidad }}
        </span>
      </div>

      <p class="mt-2 font-mono text-[10px] uppercase tracking-[0.18em] text-[#666666]">
        {{ stock > 0 ? `${stock} en stock` : 'Sin inventario' }}
      </p>
    </div>
  </NuxtLink>
</template>