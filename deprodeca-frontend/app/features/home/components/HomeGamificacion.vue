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
  <section ref="sectionRef" class="py-24 md:py-32 bg-[#FAFAF9] border-t border-[#D6D3D1]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="mb-14 md:mb-16">
        <p class="font-mono text-[11px] uppercase text-[#78716C] tracking-[0.2em] mb-4">03 · Programa de Lealtad</p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-[#1C1917] leading-[0.92]">Crece con Nosotros</h2>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-14">
        <article v-for="(paso, index) in pasos" :key="paso.numero" class="bg-white border border-[#D6D3D1] p-6" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'" :style="{ transition: 'opacity 500ms ease-out, transform 500ms ease-out', transitionDelay: `${index * 120}ms` }">
          <p class="font-display text-[2.5rem] font-bold text-[#A16207]/20 leading-none">{{ paso.numero }}</p>
          <h3 class="font-body text-[15px] font-bold text-[#1C1917] mt-3">{{ paso.titulo }}</h3>
          <p class="font-body text-[13px] text-[#78716C] mt-2 leading-relaxed">{{ paso.desc }}</p>
        </article>
      </div>

      <div class="border border-[#D6D3D1] bg-white overflow-hidden">
        <div class="hidden md:grid grid-cols-[2fr_1fr_2fr_1fr] bg-[#1C1917] text-[#FAFAF9] font-mono text-[10px] uppercase tracking-[0.2em] px-6 py-4">
          <span>Rango</span>
          <span>Desde</span>
          <span>Progreso</span>
          <span>Nivel</span>
        </div>

        <div v-for="(rango, rowIndex) in rangos" :key="rango.nombre" class="grid grid-cols-1 md:grid-cols-[2fr_1fr_2fr_1fr] px-6 py-5 border-b border-[#D6D3D1] last:border-b-0 hover:bg-[#F5F0E8] transition-colors duration-300 gap-3 md:gap-0 items-center" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3'" :style="{ transition: 'opacity 400ms ease-out, transform 400ms ease-out', transitionDelay: `${rowIndex * 100}ms` }">
          <div class="flex items-center gap-3">
            <svg width="18" height="18" viewBox="0 0 18 18" fill="none"><circle v-if="rango.nivel===1" cx="9" cy="9" r="6" :stroke="rango.color" stroke-width="1.5"/><path v-else-if="rango.nivel===2" d="M9 2L16 9L9 16L2 9L9 2Z" :stroke="rango.color" stroke-width="1.5"/><path v-else-if="rango.nivel===3" d="M9 2L15 5.5V12.5L9 16L3 12.5V5.5L9 2Z" :stroke="rango.color" stroke-width="1.5"/><path v-else d="M9 1.5L11 6.8L16.5 9L11 11.2L9 16.5L7 11.2L1.5 9L7 6.8L9 1.5Z" :stroke="rango.color" stroke-width="1.5"/></svg>
            <span class="font-body text-[15px] font-bold text-[#1C1917]">{{ rango.nombre }}</span>
          </div>
          <span class="font-mono text-[14px] text-[#1C1917]">{{ rango.minimo }}</span>
          <div class="flex items-center gap-1.5">
            <div v-for="n in 4" :key="n" class="w-6 h-2 border border-[#D6D3D1] overflow-hidden">
              <div class="h-full bg-[#A16207]" :style="{ width: n <= rango.nivel ? ((visible || reducedMotion === 'reduce') ? '100%' : '0%') : '0%', transition: 'width 800ms ease-out', transitionDelay: `${300 + rowIndex * 100}ms` }" />
            </div>
          </div>
          <span class="font-mono text-[12px] text-[#78716C]">{{ rango.nivel }}/4</span>
        </div>
      </div>

      <div class="mt-10" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'" style="transition: opacity 500ms ease-out; transition-delay: 600ms">
        <NuxtLink to="/perfil" class="inline-flex items-center border border-[#D6D3D1] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#1C1917] hover:border-[#A16207] hover:text-[#A16207] transition-colors duration-300">
          Ver mis beneficios
        </NuxtLink>
      </div>
    </div>
  </section>
</template>
