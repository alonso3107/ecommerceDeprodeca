<!--
  ═══════════════════════════════════════════════════════════════
  HomeCtaFinal.vue — CTA Premium · DEPRODECA
  Flat Design V2: bg-negro, SVG geométricos decorativos,
  tipografía masiva, staggered reveals. CERO sombras/degradados.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion } from "@vueuse/core"

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
const reducedMotion = usePreferredReducedMotion()

useIntersectionObserver(
  sectionRef,
  ([entry]) => {
    if (entry?.isIntersecting && !visible.value) {
      visible.value = true
    }
  },
  { threshold: 0.25 },
)

function irARegistro() {
  navigateTo("/auth/registro")
}
</script>

<template>
  <section
    ref="sectionRef"
    class="relative py-24 md:py-32 bg-negro border-t border-stone overflow-hidden"
  >
    <!-- ═══ SVG Geométrico Decorativo (fondo) ═══ -->
    <div class="absolute inset-0 pointer-events-none" aria-hidden="true">
      <!-- Círculo masivo dorado -->
      <svg
        class="absolute -top-32 -right-32 w-[600px] h-[600px] text-dorado opacity-[0.04]"
        :class="visible ? 'scale-100' : 'scale-75'"
        :style="{ transition: 'transform 1200ms ease-out' }"
        viewBox="0 0 600 600" fill="none"
      >
        <circle cx="300" cy="300" r="280" stroke="currentColor" stroke-width="1" />
        <circle cx="300" cy="300" r="200" stroke="currentColor" stroke-width="1" />
        <circle cx="300" cy="300" r="120" stroke="currentColor" stroke-width="1" />
        <circle cx="300" cy="300" r="50" stroke="currentColor" stroke-width="1.5" />
      </svg>

      <!-- Rombo grande abajo-izquierda -->
      <svg
        class="absolute -bottom-24 -left-24 w-[400px] h-[400px] text-stone-oscuro opacity-[0.06]"
        :class="visible ? 'scale-100 rotate-0' : 'scale-75 rotate-12'"
        :style="{ transition: 'transform 1400ms ease-out' }"
        viewBox="0 0 400 400" fill="none"
      >
        <rect x="200" y="20" width="254" height="254" transform="rotate(45 200 20)" stroke="currentColor" stroke-width="1" />
        <rect x="200" y="60" width="198" height="198" transform="rotate(45 200 60)" stroke="currentColor" stroke-width="1" />
      </svg>

      <!-- Líneas diagonales decorativas -->
      <svg class="absolute top-0 right-0 w-full h-full opacity-[0.03]" viewBox="0 0 1200 600" fill="none">
        <line x1="900" y1="0" x2="1200" y2="600" stroke="#A16207" stroke-width="1" />
        <line x1="1000" y1="0" x2="1200" y2="400" stroke="#A16207" stroke-width="1" />
        <line x1="1100" y1="0" x2="1200" y2="200" stroke="#A16207" stroke-width="1" />
      </svg>
    </div>

    <!-- ═══ Contenido ═══ -->
    <div class="relative z-10 max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="grid lg:grid-cols-[1.2fr_0.8fr] gap-12 lg:gap-16 items-center">
        <!-- ═══ COLUMNA IZQUIERDA: Texto + CTA ═══ -->
        <div>
          <!-- Label -->
          <p
            class="font-mono text-[11px] text-stone-oscuro uppercase tracking-[0.2em] mb-6"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'"
            style="transition: opacity 500ms ease-out; transition-delay: 0ms"
          >
            Activa tu canal mayorista
          </p>

          <!-- Headline masivo -->
          <h2
            class="font-display text-[clamp(2.8rem,7vw,5.6rem)] font-black text-blanco leading-[0.88] tracking-[-0.02em]"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
            style="transition: opacity 600ms ease-out, transform 600ms ease-out; transition-delay: 100ms"
          >
            Empieza a<br />comprar
            <span class="text-dorado">como negocio.</span>
          </h2>

          <!-- Descripción -->
          <p
            class="font-body text-[16px] md:text-[17px] text-stone-oscuro max-w-[48ch] mt-6 leading-relaxed"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'"
            style="transition: opacity 500ms ease-out; transition-delay: 250ms"
          >
            Registrate gratis en minutos, valida tu cuenta y desbloquea precios mayoristas reales para surtir tu bodega con mejor margen.
          </p>

          <!-- Botones -->
          <div
            class="mt-8 flex flex-wrap items-center gap-4"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3'"
            style="transition: opacity 500ms ease-out, transform 500ms ease-out; transition-delay: 380ms"
          >
            <!-- CTA Principal -->
            <button
              class="flex items-center gap-3 bg-dorado text-negro font-display text-[15px] font-bold uppercase tracking-[0.04em] px-8 py-4 border border-dorado hover:bg-dorado-claro hover:border-dorado-claro transition-colors duration-300 min-h-[56px]"
              @click="irARegistro"
            >
              Crear cuenta gratis
              <svg width="16" height="12" viewBox="0 0 16 12" fill="none">
                <path d="M0 6H14M14 6L9 1.5M14 6L9 10.5" stroke="currentColor" stroke-width="2" stroke-linecap="square" stroke-linejoin="miter"/>
              </svg>
            </button>

            <!-- Link secundario -->
            <NuxtLink
              to="/catalogo"
              class="flex items-center gap-2 font-mono text-[11px] uppercase tracking-[0.16em] text-stone-oscuro hover:text-dorado transition-colors duration-300"
            >
              Ver catálogo
              <svg width="12" height="10" viewBox="0 0 12 10" fill="none" aria-hidden="true">
                <path d="M1 5H11M11 5L7.5 1.5M11 5L7.5 8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
              </svg>
            </NuxtLink>
          </div>

          <!-- Pills informativas -->
          <div
            class="mt-8 flex flex-wrap items-center gap-2 md:gap-3"
            :class="(visible || reducedMotion === 'reduce') ? 'opacity-100' : 'opacity-0'"
            style="transition: opacity 500ms ease-out; transition-delay: 520ms"
          >
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.14em] border border-stone/30 px-2.5 py-1">
              Sin tarjeta
            </span>
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.14em] border border-stone/30 px-2.5 py-1">
              Sin contrato
            </span>
            <span class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.14em] border border-stone/30 px-2.5 py-1">
              Aprobación en 2 min
            </span>
          </div>
        </div>

        <!-- ═══ COLUMNA DERECHA: Steps verticales ═══ -->
        <div
          class="relative"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'"
          style="transition: opacity 600ms ease-out, transform 600ms ease-out; transition-delay: 250ms"
        >
          <!-- Línea conectora vertical -->
          <div class="absolute left-[13px] top-6 bottom-6 w-px bg-stone/20" aria-hidden="true" />

          <div class="flex flex-col gap-5 relative">
            <!-- Paso 01 -->
            <div class="flex items-start gap-4">
              <div class="relative z-10 w-[26px] h-[26px] flex items-center justify-center border border-dorado bg-negro shrink-0">
                <span class="font-mono text-[10px] text-dorado font-bold">01</span>
              </div>
              <div>
                <p class="font-body text-[15px] font-bold text-blanco leading-tight">Crea tu cuenta</p>
                <p class="font-body text-[13px] text-stone-oscuro leading-snug mt-1">Con RUC y datos básicos de tu bodega.</p>
              </div>
            </div>

            <!-- Paso 02 -->
            <div class="flex items-start gap-4">
              <div class="relative z-10 w-[26px] h-[26px] flex items-center justify-center border border-dorado bg-negro shrink-0">
                <span class="font-mono text-[10px] text-dorado font-bold">02</span>
              </div>
              <div>
                <p class="font-body text-[15px] font-bold text-blanco leading-tight">Elige y pide</p>
                <p class="font-body text-[13px] text-stone-oscuro leading-snug mt-1">Categorías, cantidades y fecha de entrega.</p>
              </div>
            </div>

            <!-- Paso 03 -->
            <div class="flex items-start gap-4">
              <div class="relative z-10 w-[26px] h-[26px] flex items-center justify-center border border-dorado bg-negro shrink-0">
                <span class="font-mono text-[10px] text-dorado font-bold">03</span>
              </div>
              <div>
                <p class="font-body text-[15px] font-bold text-blanco leading-tight">Recibe y gana</p>
                <p class="font-body text-[13px] text-stone-oscuro leading-snug mt-1">Tu pedido y recompensas por recurrencia.</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
/* 
 * Solo lo que Tailwind no puede expresar.
 * Flat Design V2: CERO border-radius, CERO sombras, CERO degradados.
 */
</style>
