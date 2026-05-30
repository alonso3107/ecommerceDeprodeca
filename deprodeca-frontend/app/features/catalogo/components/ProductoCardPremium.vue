<!-- ProductoCardPremium.vue — Fresh Gallery Flat Design -->
<script setup lang="ts">
import { useIntersectionObserver } from '@vueuse/core'

const props = defineProps<{
  slug: string
  nombre: string
  precio: number
  unidad: string
  imagenUrl: string
  categoria: string
  stock: number
  indice: number
}>()

const PLACEHOLDER = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='400' height='400'%3E%3Crect fill='%23F5F0E8' width='400' height='400'/%3E%3Ctext x='200' y='210' text-anchor='middle' font-size='16' font-family='monospace' fill='%23C5BFB5'%3EDEPRODECA%3C/text%3E%3C/svg%3E"

// Animación al entrar al viewport
const cardRef = ref<HTMLElement | null>(null)
const visible = ref(false)
useIntersectionObserver(cardRef, ([{ isIntersecting }]) => {
  if (isIntersecting) visible.value = true
}, { threshold: 0.1 })

function formatearPrecio(p: number): string {
  return new Intl.NumberFormat('es-PE', { style: 'currency', currency: 'PEN', minimumFractionDigits: 2 }).format(p)
}

function claseStock(s: number): string {
  if (s > 20) return 'text-[#16A34A]'
  if (s > 0) return 'text-[#A16207]'
  return 'text-[#DC2626]'
}

function textoStock(s: number): string {
  if (s > 20) return `${s} en stock`
  if (s > 0) return `Solo ${s}`
  return 'Agotado'
}

function manejarErrorImagen(e: Event) {
  const img = e.target as HTMLImageElement
  if (img && img.src !== PLACEHOLDER) img.src = PLACEHOLDER
}
</script>

<template>
  <article
    ref="cardRef"
    class="group overflow-hidden rounded-[28px] border border-[#E7E0D5] bg-white shadow-[0_24px_80px_-48px_rgba(28,25,23,0.45)] transition-all duration-500 ease-out"
    :class="visible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
    :style="{ transitionDelay: `${indice * 80}ms` }"
  >
    <!-- Imagen -->
    <NuxtLink :to="`/catalogo/${slug}`" class="relative block aspect-[4/3] overflow-hidden bg-gradient-to-br from-[#F6F1E8] via-white to-[#EFE7D8]">
      <img
        :src="imagenUrl"
        :alt="nombre"
        class="h-full w-full object-contain p-5 transition-transform duration-700 ease-out group-hover:scale-[1.03]"
        loading="lazy"
        @error="manejarErrorImagen"
      />

      <div class="pointer-events-none absolute inset-0 bg-gradient-to-t from-[#1C1917]/18 via-transparent to-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100" />

      <!-- Badge categoría -->
      <span class="absolute left-4 top-4 rounded-full bg-[#1C1917] px-3 py-1.5 font-mono text-[9px] uppercase tracking-[0.16em] text-[#FAFAF9]">
        {{ categoria }}
      </span>

      <!-- Overlay agotado -->
      <div v-if="stock === 0" class="absolute inset-0 flex items-center justify-center bg-[#1C1917]/60">
        <span class="font-display text-3xl font-black uppercase tracking-wider text-[#FAFAF9]">Agotado</span>
      </div>
    </NuxtLink>

    <!-- Info -->
    <div class="flex flex-col gap-4 p-5 md:p-6">
      <NuxtLink :to="`/catalogo/${slug}`">
        <h3 class="line-clamp-2 font-body text-[15px] font-semibold leading-snug text-[#1C1917] transition-colors duration-300 group-hover:text-[#A16207]">
          {{ nombre }}
        </h3>
      </NuxtLink>

      <div class="flex items-center justify-between gap-3">
        <span class="font-mono text-[10px] uppercase tracking-[0.14em]" :class="claseStock(stock)">
          ● {{ textoStock(stock) }}
        </span>
        <span class="font-mono text-[13px] text-[#5C554D]">
          {{ formatearPrecio(precio) }} <span class="text-[10px] uppercase tracking-[0.14em]">/ {{ unidad }}</span>
        </span>
      </div>
    </div>
  </article>
</template>
