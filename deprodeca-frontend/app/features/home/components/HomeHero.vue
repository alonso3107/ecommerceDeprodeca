<script setup lang="ts">
import { usePreferredReducedMotion } from "@vueuse/core"

const enlaces = [
  { texto: "PRODUCTOS", ruta: "/catalogo" },
  { texto: "SOLUCIONES", ruta: "/soluciones" },
  { texto: "TIENDA", ruta: "/catalogo" },
  { texto: "CONTACTO", ruta: "/contacto" },
] as const

function navegar(ruta: string) {
  navigateTo(ruta)
}

const reducedMotion = usePreferredReducedMotion()
const capsuleVisible = ref(false)
const metaVisible = ref(false)
const enlacesVisibles = ref([false, false, false, false])
const enlaceHover = "hover:text-[#FAFAF9] hover:tracking-[-0.025em]"

onMounted(() => {
  if (reducedMotion.value === "reduce") {
    capsuleVisible.value = true
    metaVisible.value = true
    enlacesVisibles.value = [true, true, true, true]
    return
  }

  setTimeout(() => {
    capsuleVisible.value = true
  }, 100)

  enlaces.forEach((_, index) => {
    setTimeout(() => {
      enlacesVisibles.value[index] = true
    }, index * 100)
  })

  setTimeout(() => {
    metaVisible.value = true
  }, 500)
})
</script>

<template>
  <section class="relative bg-[#A16207] min-h-screen flex flex-col overflow-hidden" aria-label="Hero principal DEPRODECA">
    <div class="absolute top-6 left-6 md:top-8 md:left-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wider uppercase select-none transition-opacity duration-[800ms]" :class="metaVisible ? 'opacity-100' : 'opacity-0'" aria-hidden="true">
      <p>Ica, Perú</p>
      <p>Sede Central</p>
    </div>

    <div class="absolute top-6 right-6 md:top-8 md:right-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wider uppercase text-right select-none transition-opacity duration-[800ms]" :class="metaVisible ? 'opacity-100' : 'opacity-0'" aria-hidden="true">
      <p>Centro de</p>
      <p>Distribución</p>
    </div>

    <div class="absolute bottom-6 left-6 md:bottom-8 md:left-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wide uppercase max-w-[200px] select-none transition-opacity duration-[800ms]" :class="metaVisible ? 'opacity-100' : 'opacity-0'" aria-hidden="true">
      <p>Distribución Premium</p>
      <p>Calidad Nestlé</p>
    </div>

    <div class="absolute bottom-6 right-6 md:bottom-8 md:right-8">
      <button
        class="font-mono text-[11px] md:text-xs text-[#1C1917]/70 hover:text-[#1C1917] border border-[#1C1917]/40 hover:border-[#1C1917] px-4 py-2 transition-colors duration-300 uppercase tracking-widest focus:outline-none focus-visible:border-[#FAFAF9]"
        @click="navegar('/catalogo')"
      >
        Comprar Ahora
      </button>
    </div>

    <nav
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-[42vh] z-10
             bg-[#1C1917] text-[#FAFAF9] font-bold px-8 py-3 border-b-2 border-[#FAFAF9]/30"
      :class="capsuleVisible ? 'opacity-100 -translate-y-[42vh]' : 'opacity-0 -translate-y-[calc(42vh+12px)]'"
      style="transition: opacity 500ms ease-out, transform 500ms ease-out"
      aria-label="Navegación principal"
    >
      <span class="font-display text-xl md:text-2xl tracking-[0.15em] uppercase select-none">
        DEPRODECA
      </span>
    </nav>

    <div class="flex-1 flex items-center justify-center px-4">
      <div class="flex flex-col items-start gap-0 leading-none">
        <NuxtLink
          v-for="(enlace, index) in enlaces"
          :key="enlace.texto"
          :to="enlace.ruta"
          class="font-display text-[clamp(4rem,12vw,9rem)] text-black
                  leading-[0.82] tracking-[-0.04em] uppercase
                 transition-[color,letter-spacing] duration-300
                  cursor-pointer select-none"
          :class="enlaceHover"
          :style="{ transition: 'opacity 600ms ease-out, transform 600ms ease-out, color 300ms ease', opacity: enlacesVisibles[index] ? '1' : '0', transform: enlacesVisibles[index] ? 'translateY(0)' : 'translateY(32px)' }"
        >
          {{ enlace.texto }}
        </NuxtLink>
      </div>
    </div>

    <div class="absolute bottom-0 left-0 right-0 border-t border-black/15 flex justify-between px-6 md:px-8 py-3 font-mono text-[9px] md:text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em] select-none transition-opacity duration-[800ms]" :class="metaVisible ? 'opacity-100' : 'opacity-0'" aria-hidden="true">
      <span>Filial Oficial Nestlé Perú</span>
      <span>B2B · Mayorista</span>
    </div>
  </section>
</template>
