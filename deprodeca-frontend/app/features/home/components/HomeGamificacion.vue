<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion } from "@vueuse/core"

const pasos = [
  { numero: "01", titulo: "Compra", desc: "Cada compra suma puntos automaticamente en tu cuenta." },
  { numero: "02", titulo: "Subi", desc: "Al subir de nivel desbloqueas mejores condiciones de compra." },
  { numero: "03", titulo: "Disfruta", desc: "Accede a beneficios exclusivos segun tu rango actual." },
] as const

const rangos = [
  { nombre: "Bronce", minimo: "S/ 0", nivel: 1, color: "#CD7F32" },
  { nombre: "Plata", minimo: "S/ 5,000", nivel: 2, color: "#A8A8A8" },
  { nombre: "Oro", minimo: "S/ 20,000", nivel: 3, color: "#A16207" },
  { nombre: "Platino", minimo: "S/ 50,000", nivel: 4, color: "#8E8E8E" },
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
  <section ref="sectionRef" class="py-24 md:py-32 bg-blanco border-t border-stone">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="mb-14 md:mb-16">
        <p class="font-mono text-[11px] uppercase text-stone-oscuro tracking-[0.2em] mb-4">Beneficios por recompra</p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-negro leading-[0.92]">
          Crece con Nosotros<span class="text-dorado">.</span>
        </h2>
      </div>

      <div class="relative mb-16 md:mb-20">
        <div class="absolute top-8 left-[15%] right-[15%] h-px bg-stone hidden md:block" aria-hidden="true" />

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 md:gap-0 relative">
          <article
            v-for="(paso, index) in pasos"
            :key="paso.numero"
            class="flex flex-col items-center text-center"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
            :style="{ transition: 'opacity 500ms ease-out, transform 500ms ease-out', transitionDelay: `${index * 150}ms` }"
          >
            <div class="relative z-10 w-16 h-16 flex items-center justify-center mb-4">
              <div class="absolute inset-0 border border-stone bg-blanco" />
              <svg v-if="index === 0" width="24" height="24" viewBox="0 0 24 24" fill="none" class="text-dorado relative z-10">
                <rect x="3" y="3" width="18" height="18" stroke="currentColor" stroke-width="1.5" />
                <path d="M8 12H16M12 8V16" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" />
              </svg>
              <svg v-else-if="index === 1" width="24" height="24" viewBox="0 0 24 24" fill="none" class="text-dorado relative z-10">
                <path d="M12 3L20 7V17L12 21L4 17V7L12 3Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
                <path d="M12 12V16M10 10H14" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" />
              </svg>
              <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" class="text-dorado relative z-10">
                <circle cx="12" cy="12" r="8" stroke="currentColor" stroke-width="1.5" />
                <path d="M9 12L11.5 14.5L15.5 10.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter" />
              </svg>
            </div>

            <p class="font-display text-[2rem] font-bold text-dorado/25 leading-none mb-2">{{ paso.numero }}</p>
            <h3 class="font-display text-[18px] font-bold text-negro uppercase mb-1">{{ paso.titulo }}</h3>
            <p class="font-body text-[13px] text-stone-oscuro max-w-[220px] leading-relaxed">{{ paso.desc }}</p>
          </article>
        </div>
      </div>

      <div class="border border-stone bg-blanco overflow-hidden mb-14">
        <div class="px-6 py-4 border-b border-stone bg-crema">
          <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-stone-oscuro">Tus niveles comerciales</p>
        </div>

        <div class="p-6 md:p-8">
          <div class="grid grid-cols-4 gap-2 mb-5">
            <div v-for="(rango, i) in rangos" :key="rango.nombre" class="flex flex-col items-center text-center">
              <svg width="32" height="32" viewBox="0 0 32 32" fill="none" class="mb-2" :class="(visible || reducedMotion === 'reduce') ? 'scale-100' : 'scale-0'" :style="{ transition: 'transform 500ms ease-out', transitionDelay: `${i * 150}ms`, color: rango.color }">
                <circle v-if="rango.nivel === 1" cx="16" cy="16" r="10" stroke="currentColor" stroke-width="1.5" />
                <path v-else-if="rango.nivel === 2" d="M16 4L28 16L16 28L4 16L16 4Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
                <path v-else-if="rango.nivel === 3" d="M16 4L26 10V22L16 28L6 22V10L16 4Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
                <path v-else d="M16 3L19.5 12.5L29 16L19.5 19.5L16 29L12.5 19.5L3 16L12.5 12.5L16 3Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
              </svg>
              <span class="font-display text-[14px] font-bold text-negro uppercase leading-none">{{ rango.nombre }}</span>
              <span class="font-mono text-[10px] text-stone-oscuro mt-0.5">{{ rango.minimo }}</span>
            </div>
          </div>

          <div class="relative h-3 bg-crema overflow-hidden mb-3">
            <div class="absolute inset-y-0 left-0 flex w-full">
              <div v-for="(rango, i) in rangos" :key="i" class="h-full border-r border-blanco last:border-r-0 transition-all duration-1000 ease-out" :style="{ width: '25%', backgroundColor: (visible || reducedMotion === 'reduce') ? rango.color : 'var(--color-stone)', opacity: (visible || reducedMotion === 'reduce') ? 1 : 0, transitionDelay: `${400 + i * 100}ms` }" />
            </div>
          </div>

          <p class="font-mono text-[10px] text-dorado uppercase tracking-[0.15em] text-right">
            Nivel actual: <span class="font-bold">{{ rangos[1].nombre }}</span>
          </p>
        </div>
      </div>

      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-10">
        <div
          v-for="(rango, i) in rangos"
          :key="rango.nombre"
          class="border border-stone bg-blanco p-5 text-center hover:border-dorado transition-colors duration-300"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3'"
          :style="{ transition: 'opacity 400ms ease-out, transform 400ms ease-out', transitionDelay: `${600 + i * 80}ms` }"
        >
          <svg width="40" height="40" viewBox="0 0 40 40" fill="none" class="mx-auto mb-3" :style="{ color: rango.color }">
            <circle v-if="rango.nivel === 1" cx="20" cy="20" r="13" stroke="currentColor" stroke-width="1.5" />
            <path v-else-if="rango.nivel === 2" d="M20 5L35 20L20 35L5 20L20 5Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
            <path v-else-if="rango.nivel === 3" d="M20 5L32 12V28L20 35L8 28V12L20 5Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
            <path v-else d="M20 4L24 16L36 20L24 24L20 36L16 24L4 20L16 16L20 4Z" stroke="currentColor" stroke-width="1.5" stroke-linejoin="miter" />
          </svg>

          <p class="font-display text-[16px] font-bold text-negro uppercase mb-1">{{ rango.nombre }}</p>
          <p class="font-mono text-[11px] text-stone-oscuro mb-3">Desde {{ rango.minimo }}</p>
          <span class="inline-block font-mono text-[10px] text-dorado border border-stone px-2 py-1 uppercase tracking-[0.1em] mb-3">
            Nivel {{ rango.nivel }}
          </span>
          <div class="h-1.5 bg-crema overflow-hidden">
            <div class="h-full transition-all duration-1000 ease-out" :style="{ width: (visible || reducedMotion === 'reduce') ? `${rango.nivel * 25}%` : '0%', backgroundColor: rango.color, transitionDelay: `${800 + i * 100}ms` }" />
          </div>
        </div>
      </div>

      <div class="text-center" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'" style="transition: opacity 500ms ease-out; transition-delay: 900ms">
        <NuxtLink
          to="/perfil"
          class="inline-flex items-center gap-3 border border-stone px-6 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-negro hover:border-dorado hover:text-dorado transition-colors duration-300"
        >
          Ver mis beneficios
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" aria-hidden="true">
            <path d="M1 5H11M11 5L7.5 1.5M11 5L7.5 8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
          </svg>
        </NuxtLink>
      </div>
    </div>
  </section>
</template>
