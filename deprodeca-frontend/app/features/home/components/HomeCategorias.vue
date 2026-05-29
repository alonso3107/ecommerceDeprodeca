<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion } from "@vueuse/core"
import { imagenesProductos } from "~/shared/utils/imagenes-productos"

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

const imagenCategoria = (slug: string): string => {
  const mapa: Record<string, string> = {
    bebidas: imagenesProductos.bebidas[0],
    snacks: imagenesProductos.snacks[0],
    lacteos: imagenesProductos.lacteos[0],
    abarrotes: imagenesProductos.abarrotes[0],
    limpieza: imagenesProductos.limpieza[0],
    "cuidado-personal": imagenesProductos["cuidado-personal"][0],
  }
  return mapa[slug] || ""
}
</script>

<template>
  <section ref="sectionRef" class="py-24 md:py-32 bg-crema border-t border-stone">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-14 md:mb-16">
        <div>
          <p class="font-mono text-[11px] uppercase text-stone-oscuro tracking-[0.2em] mb-4">Catalogo por categoria</p>
          <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-negro leading-[0.92]">
            Explora<span class="text-dorado">.</span>
          </h2>
        </div>
        <NuxtLink
          to="/catalogo"
          class="font-mono text-[11px] uppercase tracking-[0.18em] text-stone-oscuro hover:text-dorado transition-colors duration-300 flex items-center gap-2"
        >
          Ver todas
          <svg width="12" height="10" viewBox="0 0 12 10" class="text-dorado" fill="none" aria-hidden="true">
            <path d="M1 5H11M11 5L7.5 1.5M11 5L7.5 8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
          </svg>
        </NuxtLink>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-7 md:gap-8">
        <NuxtLink
          v-for="(cat, index) in categorias"
          :key="cat.slug"
          :to="`/catalogo?categoria=${cat.slug}`"
          class="explora-card group relative overflow-hidden border border-stone hover:border-dorado transition-colors duration-500 min-h-[300px] md:min-h-[340px] lg:min-h-[360px]"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-5'"
          :style="{ transition: 'opacity 520ms cubic-bezier(0.22, 1, 0.36, 1), transform 520ms cubic-bezier(0.22, 1, 0.36, 1)', transitionDelay: `${index * 90}ms` }"
        >
          <img
            :src="imagenCategoria(cat.slug)"
            :alt="cat.nombre"
            class="absolute inset-0 w-full h-full object-cover scale-[1.01] group-hover:scale-[1.08] transition-transform duration-700 ease-out"
            loading="lazy"
          >

          <div class="explora-overlay-base absolute inset-0" />
          <div class="explora-overlay-radar absolute inset-0" />
          <div class="explora-overlay-grid absolute inset-0" />
          <div class="explora-overlay-sweep absolute inset-0" />

          <div class="relative z-[2] h-full flex items-end p-3 md:p-4">
            <div class="w-full explora-content-panel translate-y-2 group-hover:translate-y-0 transition-transform duration-500 ease-out">
              <div class="flex items-center justify-between gap-4">
                <div class="flex items-center gap-3 min-w-0">
                <svg width="20" height="20" viewBox="0 0 20 20" fill="none" class="text-dorado shrink-0">
                  <circle v-if="cat.tipo === 'circulo'" cx="10" cy="10" r="6.5" stroke="currentColor" stroke-width="1.5" />
                  <rect v-else-if="cat.tipo === 'rombo'" x="10" y="3" width="9.9" height="9.9" transform="rotate(45 10 3)" stroke="currentColor" stroke-width="1.5" />
                  <path v-else-if="cat.tipo === 'hexagono'" d="M10 3L16 6.5V13.5L10 17L4 13.5V6.5L10 3Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
                  <template v-else-if="cat.tipo === 'cuadrado-cruz'">
                    <rect x="4" y="4" width="12" height="12" stroke="currentColor" stroke-width="1.5" />
                    <path d="M10 6V14M6 10H14" stroke="currentColor" stroke-width="1.5" />
                  </template>
                  <path v-else-if="cat.tipo === 'triangulo'" d="M10 4L16.5 16H3.5L10 4Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
                  <template v-else>
                    <circle cx="10" cy="10" r="6.5" stroke="currentColor" stroke-width="1.5" />
                    <circle cx="10" cy="10" r="4" stroke="currentColor" stroke-width="1.5" />
                    <circle cx="10" cy="10" r="1.5" stroke="currentColor" stroke-width="1.5" />
                  </template>
                </svg>
                <span class="font-display text-[clamp(1rem,2vw,1.5rem)] font-bold text-blanco uppercase leading-none tracking-[-0.02em] truncate">
                  {{ cat.nombre }}
                </span>
                </div>

                <div class="flex items-center gap-3 shrink-0">
                  <span class="font-mono text-[10px] text-dorado uppercase tracking-[0.15em] px-2 py-1 border border-dorado/40 bg-negro/35">{{ cat.count }} prod</span>
                  <svg width="14" height="10" viewBox="0 0 14 10" class="text-blanco group-hover:text-dorado group-hover:translate-x-0.5 transition-all duration-300" fill="none" aria-hidden="true">
                    <path d="M0 5H12M12 5L8 1.5M12 5L8 8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
                  </svg>
                </div>
              </div>

              <p class="mt-2 text-blanco/92 text-[12px] md:text-[13px] leading-relaxed line-clamp-2 tracking-[0.01em]">
                {{ cat.desc }}
              </p>
              <div class="mt-2 h-px bg-gradient-to-r from-dorado/60 via-blanco/40 to-transparent" />
              <div class="mt-2 flex items-center justify-between text-[10px] uppercase tracking-[0.14em] font-mono text-blanco/80">
                <span>Coleccion destacada</span>
                <span class="text-dorado">Explorar</span>
              </div>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </section>
</template>

<style scoped>
.explora-card {
  isolation: isolate;
  box-shadow: 0 18px 36px rgb(20 14 10 / 0.16);
  border-radius: 26px;
  clip-path: polygon(0 14px, 14px 0, calc(100% - 40px) 0, 100% 30px, 100% 100%, 0 100%);
  transform-origin: center;
}

.explora-card::before {
  content: "";
  position: absolute;
  inset: 0;
  border-radius: inherit;
  border: 1px solid rgb(255 255 255 / 0.1);
  pointer-events: none;
  z-index: 3;
}

.explora-card::after {
  content: "";
  position: absolute;
  top: 0;
  right: 0;
  width: 78px;
  height: 66px;
  background: linear-gradient(135deg, rgb(201 161 92 / 0.28), rgb(201 161 92 / 0));
  clip-path: polygon(42% 0, 100% 0, 100% 58%);
  pointer-events: none;
  z-index: 3;
}

.explora-overlay-base,
.explora-overlay-radar,
.explora-overlay-grid,
.explora-overlay-sweep {
  pointer-events: none;
  z-index: 1;
}

.explora-overlay-base {
  background:
    linear-gradient(to top, rgb(14 11 9 / 0.9) 8%, rgb(14 11 9 / 0.58) 48%, rgb(14 11 9 / 0.26) 100%),
    radial-gradient(circle at 12% 84%, rgb(201 161 92 / 0.24) 0%, rgb(201 161 92 / 0) 56%);
  transition: opacity 500ms ease;
}

.explora-overlay-radar {
  opacity: 0;
  mix-blend-mode: screen;
  background:
    radial-gradient(circle at 12% 84%, rgb(201 161 92 / 0.4) 0%, rgb(201 161 92 / 0.08) 18%, rgb(201 161 92 / 0) 48%),
    repeating-radial-gradient(circle at 12% 84%, rgb(201 161 92 / 0.24) 0 1px, transparent 1px 18px);
  transform: scale(1.12);
  transition: opacity 450ms ease;
}

.explora-overlay-grid {
  opacity: 0;
  background-image:
    linear-gradient(to right, rgb(255 255 255 / 0.13) 1px, transparent 1px),
    linear-gradient(to bottom, rgb(255 255 255 / 0.13) 1px, transparent 1px);
  background-size: 28px 28px;
  mask-image: linear-gradient(to top, black 5%, black 56%, transparent 100%);
  transition: opacity 350ms ease;
}

.explora-overlay-sweep {
  opacity: 0;
  background: linear-gradient(112deg, transparent 0%, rgb(255 255 255 / 0.42) 50%, transparent 100%);
  transform: translateX(-145%);
}

.explora-card:hover .explora-overlay-base {
  opacity: 0.96;
}

.explora-card:hover .explora-overlay-radar {
  opacity: 0.9;
  animation: radar-pulse 1.6s ease-out;
}

.explora-card:hover .explora-overlay-grid {
  opacity: 0.42;
}

.explora-card:hover .explora-overlay-sweep {
  opacity: 0.75;
  animation: light-sweep 820ms cubic-bezier(0.22, 1, 0.36, 1);
}

@keyframes light-sweep {
  from { transform: translateX(-145%); }
  to { transform: translateX(145%); }
}

@keyframes radar-pulse {
  0% { transform: scale(1.18); opacity: 0.2; }
  40% { opacity: 0.9; }
  100% { transform: scale(1); opacity: 0.4; }
}

.explora-content-panel {
  border: 1px solid rgb(255 255 255 / 0.14);
  background:
    linear-gradient(145deg, rgb(15 12 10 / 0.7) 0%, rgb(15 12 10 / 0.5) 100%);
  backdrop-filter: blur(3px);
  padding: 0.85rem;
  border-radius: 16px;
}

.explora-card:hover .explora-content-panel {
  border-color: rgb(201 161 92 / 0.48);
}

@media (prefers-reduced-motion: reduce) {
  .explora-overlay-sweep,
  .explora-overlay-radar,
  .explora-card img,
  .explora-card div {
    animation: none !important;
    transition-duration: 0.01ms !important;
  }
}

@media (max-width: 768px) {
  .explora-card {
    border-radius: 22px;
    clip-path: polygon(0 12px, 12px 0, calc(100% - 28px) 0, 100% 22px, 100% 100%, 0 100%);
  }
}
</style>
