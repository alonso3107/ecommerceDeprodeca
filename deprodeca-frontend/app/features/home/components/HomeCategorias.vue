<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion } from "@vueuse/core"

const categorias = [
  { nombre: "Bebidas", slug: "bebidas", count: "120+", desc: "Gaseosas, aguas, jugos y energizantes", tipo: "circulo" },
  { nombre: "Snacks", slug: "snacks", count: "85+", desc: "Galletas, chocolates y chips", tipo: "rombo" },
  { nombre: "Lacteos", slug: "lacteos", count: "60+", desc: "Leches, yogures y quesos", tipo: "hexagono" },
  { nombre: "Abarrotes", slug: "abarrotes", count: "150+", desc: "Arroz, azucar, aceite y menestras", tipo: "cuadrado-cruz" },
  { nombre: "Limpieza", slug: "limpieza", count: "45+", desc: "Detergentes y desinfectantes", tipo: "triangulo" },
  { nombre: "Cuidado Personal", slug: "cuidado-personal", count: "40+", desc: "Shampoo, jabon y desodorante", tipo: "concentrico" },
] as const

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
const reducedMotion = usePreferredReducedMotion()

useIntersectionObserver(
  sectionRef,
  ([entry]) => {
    if (entry?.isIntersecting) visible.value = true
  },
  { threshold: 0.15 },
)
</script>

<template>
  <section ref="sectionRef" class="py-24 md:py-32 bg-white border-t border-[#D6D3D1]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-14 md:mb-16">
        <div>
          <p class="font-mono text-[11px] uppercase text-[#78716C] tracking-[0.2em] mb-4">02 · Categorias</p>
          <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-[#1C1917] leading-[0.92]">Explora</h2>
        </div>
        <NuxtLink to="/catalogo" class="font-mono text-[11px] uppercase tracking-[0.18em] text-[#78716C] hover:text-[#A16207] transition-colors duration-300">
          Ver todas ->
        </NuxtLink>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <NuxtLink
          v-for="(cat, index) in categorias"
          :key="cat.slug"
          :to="`/catalogo?categoria=${cat.slug}`"
          class="bg-[#FAFAF9] border border-[#D6D3D1] p-8 hover:border-[#A16207] transition-colors duration-300"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-5'"
          :style="{ transition: 'opacity 450ms ease-out, transform 450ms ease-out', transitionDelay: `${index * 80}ms` }"
        >
          <svg width="40" height="40" viewBox="0 0 40 40" fill="none" class="mb-5 text-[#A16207] opacity-30" :style="{ transition: 'transform 600ms ease-out', transform: (visible || reducedMotion === 'reduce') ? 'scale(1)' : 'scale(0.9)' }">
            <circle v-if="cat.tipo === 'circulo'" cx="20" cy="20" r="13" stroke="currentColor" stroke-width="1.5" />
            <path v-else-if="cat.tipo === 'rombo'" d="M20 6L34 20L20 34L6 20L20 6Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
            <path v-else-if="cat.tipo === 'hexagono'" d="M20 6L32 13V27L20 34L8 27V13L20 6Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
            <template v-else-if="cat.tipo === 'cuadrado-cruz'">
              <rect x="8" y="8" width="24" height="24" stroke="currentColor" stroke-width="1.5" />
              <path d="M20 12V28M12 20H28" stroke="currentColor" stroke-width="1.5" />
            </template>
            <path v-else-if="cat.tipo === 'triangulo'" d="M20 8L33 32H7L20 8Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
            <template v-else>
              <circle cx="20" cy="20" r="13" stroke="currentColor" stroke-width="1.5" />
              <circle cx="20" cy="20" r="8" stroke="currentColor" stroke-width="1.5" />
              <circle cx="20" cy="20" r="3" stroke="currentColor" stroke-width="1.5" />
            </template>
          </svg>

          <p class="font-body text-[18px] font-bold text-[#1C1917]">{{ cat.nombre }}</p>
          <p class="font-body text-[14px] text-[#78716C] mt-1">{{ cat.desc }}</p>
          <p class="font-mono text-[11px] text-[#A16207] mt-4 uppercase tracking-[0.15em]">{{ cat.count }} productos</p>
        </NuxtLink>
      </div>
    </div>
  </section>
</template>
