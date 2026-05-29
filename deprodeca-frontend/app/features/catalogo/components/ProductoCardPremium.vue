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

const emit = defineEmits<{ (e: 'agregar', id: string): void }>()

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
    class="group bg-white border border-[#C5BFB5] overflow-hidden transition-all duration-500 ease-out"
    :class="visible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
    :style="{ transitionDelay: `${indice * 80}ms` }"
  >
    <!-- Imagen -->
    <NuxtLink :to="`/catalogo/${slug}`" class="block relative aspect-[4/5] bg-[#F5F0E8] overflow-hidden">
      <img
        :src="imagenUrl"
        :alt="nombre"
        class="w-full h-full object-cover transition-transform duration-700 ease-out group-hover:scale-105"
        loading="lazy"
        @error="manejarErrorImagen"
      />

      <!-- Badge categoría -->
      <span class="absolute top-3 left-3 bg-[#1C1917] text-[#FAFAF9] font-mono text-[9px] uppercase tracking-[0.15em] px-3 py-1.5">
        {{ categoria }}
      </span>

      <!-- Overlay agotado -->
      <div v-if="stock === 0" class="absolute inset-0 bg-[#1C1917]/60 flex items-center justify-center">
        <span class="font-display text-3xl font-black text-[#FAFAF9] uppercase tracking-wider">Agotado</span>
      </div>
    </NuxtLink>

    <!-- Info -->
    <div class="p-5 flex flex-col gap-3">
      <NuxtLink :to="`/catalogo/${slug}`">
        <h3 class="font-body text-[15px] font-bold text-[#1C1917] leading-snug line-clamp-2 group-hover:text-[#A16207] transition-colors duration-300">
          {{ nombre }}
        </h3>
      </NuxtLink>

      <p class="font-mono text-[13px] text-[#5C554D]">
        {{ formatearPrecio(precio) }} <span class="uppercase text-[11px]">/ {{ unidad }}</span>
      </p>

      <div class="flex items-center justify-between mt-1">
        <span class="font-mono text-[10px] uppercase tracking-[0.1em]" :class="claseStock(stock)">
          ● {{ textoStock(stock) }}
        </span>

        <button
          v-if="stock > 0"
          class="bg-[#1C1917] text-[#FAFAF9] font-mono text-[10px] uppercase tracking-[0.15em] px-4 py-2 hover:bg-[#A16207] active:scale-95 transition-all duration-300"
          @click="emit('agregar', slug)"
        >
          Añadir
        </button>
        <span v-else class="font-mono text-[10px] uppercase text-[#C5BFB5] tracking-[0.15em]">Sin stock</span>
      </div>
    </div>
  </article>
</template>
