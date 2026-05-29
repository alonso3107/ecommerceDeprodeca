<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion } from "@vueuse/core"

function irARegistro() {
  navigateTo("/auth/registro")
}

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
const pulse = ref(false)
const reducedMotion = usePreferredReducedMotion()

useIntersectionObserver(
  sectionRef,
  ([entry]) => {
    if (entry?.isIntersecting && !visible.value) {
      visible.value = true
      if (reducedMotion.value !== "reduce") {
        setTimeout(() => {
          pulse.value = true
          setTimeout(() => {
            pulse.value = false
          }, 2000)
        }, 2000)
      }
    }
  },
  { threshold: 0.3 },
)
</script>

<template>
  <section ref="sectionRef" class="py-24 md:py-32 bg-[#A16207] border-t border-[#C5BFB5]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="max-w-[780px] mx-auto text-center">
        <p class="font-mono text-[11px] text-[#1C1917]/60 uppercase tracking-[0.2em] mb-7" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'" style="transition: opacity 500ms ease-out; transition-delay: 0ms">Activa tu canal mayorista</p>

        <h2 class="font-display text-[clamp(3rem,8vw,6rem)] font-bold text-[#1C1917] leading-[0.9] mb-6" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'" style="transition: opacity 600ms ease-out, transform 600ms ease-out; transition-delay: 150ms">
          Empeza a comprar
        </h2>

        <div class="h-[2px] bg-[#1C1917] mx-auto mb-8" :style="{ width: (visible || reducedMotion === 'reduce') ? '6rem' : '0rem', transition: 'width 500ms ease-out', transitionDelay: '300ms' }" />

        <p class="font-body text-[16px] text-[#1C1917]/60 max-w-lg mx-auto mb-12 leading-relaxed" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'" style="transition: opacity 500ms ease-out; transition-delay: 400ms">
          Registrate gratis en minutos y accede a precios exclusivos para tu bodega.
        </p>

        <button
          class="inline-flex items-center bg-[#1C1917] text-[#FAFAF9] font-bold px-12 py-4 transition-colors duration-300 hover:bg-[#FAFAF9] hover:text-[#1C1917]"
          :class="[(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3', pulse ? 'flat-pulse' : '']"
          style="transition: opacity 500ms ease-out, transform 500ms ease-out; transition-delay: 550ms"
          @click="irARegistro"
        >
          Crear Cuenta Gratis
        </button>

        <div class="mt-10 flex flex-wrap items-center justify-center gap-x-6 gap-y-2" :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'" style="transition: opacity 500ms ease-out; transition-delay: 700ms">
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">Sin tarjeta</span>
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">Sin compromiso</span>
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">2 min</span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
@keyframes flat-pulse {
  0%,
  100% { box-shadow: 0 0 0 0 rgba(28, 25, 23, 0); }
  50% { box-shadow: 0 0 0 6px rgba(28, 25, 23, 0.08); }
}

.flat-pulse {
  animation: flat-pulse 2s ease-out 1;
}
</style>
