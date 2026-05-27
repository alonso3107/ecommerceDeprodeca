<!--
  ═══════════════════════════════════════════════════════════════
  HomeCtaFinal.vue — Sección de cierre CTA
  Swiss Modernism 2.0 × Brutalista: fondo negro, tipografía
  masiva, geometría asimétrica, botón dorado como ancla visual.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
function irARegistro() {
  navigateTo("/auth/registro")
}
</script>

<template>
  <section class="relative py-36 md:py-48 bg-[#171717] border-t-4 border-[#D4A017] overflow-hidden">
    
    <!-- SVG geométrico masivo de fondo — estructuras asimétricas Swiss -->
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute inset-0 h-full w-full opacity-[0.08]"
      viewBox="0 0 1440 700"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <!-- Círculos concéntricos -->
      <circle cx="720" cy="350" r="320" stroke="#D4A017" stroke-width="1" />
      <circle cx="720" cy="350" r="220" stroke="#D4A017" stroke-width="1.5" />
      <circle cx="720" cy="350" r="120" stroke="#D4A017" stroke-width="2" />
      
      <!-- Líneas radiales -->
      <path d="M720 30V670" stroke="#D4A017" stroke-width="1" />
      <path d="M400 350H1040" stroke="#D4A017" stroke-width="1" />
      <path d="M494 124L946 576" stroke="#D4A017" stroke-width="1" />
      <path d="M946 124L494 576" stroke="#D4A017" stroke-width="1" />
      
      <!-- Cuadrícula sutil -->
      <path d="M180 120H1260M180 240H1260M180 360H1260M180 480H1260M180 600H1260" stroke="#D4A017" stroke-width="0.5" />
      <path d="M240 60V640M480 60V640M720 60V640M960 60V640M1200 60V640" stroke="#D4A017" stroke-width="0.5" />
      
      <!-- Triángulo central -->
      <path d="M720 140L920 500H520Z" stroke="#D4A017" stroke-width="2" />
    </svg>

    <!-- Detalle esquina superior izquierda -->
    <div class="absolute top-8 left-8 font-mono text-[10px] text-[#FAFAFA]/15 uppercase tracking-[0.35em] select-none" aria-hidden="true">
      04 · Comenzar
    </div>

    <!-- Detalle esquina superior derecha -->
    <div class="absolute top-8 right-8 font-mono text-[10px] text-[#FAFAFA]/10 uppercase tracking-[0.25em] select-none hidden md:block" aria-hidden="true">
      deprodeca.com
    </div>

    <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
      
      <div class="max-w-[780px] mx-auto text-center">
        
        <!-- Label superior -->
        <p class="font-mono text-[11px] text-[#D4A017]/60 uppercase tracking-[0.3em] mb-8">
          Sin costo de inscripción
        </p>

        <!-- Título masivo -->
        <h2 class="font-display text-[clamp(3rem,8vw,7rem)] text-[#FAFAFA] uppercase leading-[0.88] tracking-[-0.02em] mb-6">
          Empezá a<br />comprar<span class="text-[#D4A017]">.</span>
        </h2>

        <!-- Línea dorada horizontal -->
        <div class="h-[2px] w-32 bg-[#D4A017] mx-auto mb-8" />

        <!-- Subtítulo -->
        <p class="font-body text-[17px] text-[#FAFAFA]/45 max-w-lg mx-auto mb-14 leading-relaxed">
          Registrate gratis en 2 minutos. Accedé a precios exclusivos para 
          tu bodega y empezá a ahorrar desde tu primer pedido.
        </p>

        <!-- Botón CTA — dorado sólido -->
        <button
          class="group/btn inline-flex items-center gap-3 bg-[#D4A017] text-[#171717] font-display text-[18px] uppercase tracking-[0.08em] px-14 py-5 transition-all duration-500 ease-out hover:bg-[#FAFAFA] hover:gap-5 active:scale-[0.98]"
          @click="irARegistro"
        >
          <span>Crear Cuenta Gratis</span>
          <span class="text-[#171717] group-hover/btn:translate-x-1 transition-transform duration-500">→</span>
        </button>

        <!-- Garantías -->
        <div class="mt-12 flex flex-wrap items-center justify-center gap-x-8 gap-y-3">
          <span class="font-mono text-[10px] text-[#FAFAFA]/20 uppercase tracking-[0.2em]">Sin tarjeta</span>
          <span class="text-[#FAFAFA]/10 hidden sm:inline" aria-hidden="true">|</span>
          <span class="font-mono text-[10px] text-[#FAFAFA]/20 uppercase tracking-[0.2em]">Sin compromiso</span>
          <span class="text-[#FAFAFA]/10 hidden sm:inline" aria-hidden="true">|</span>
          <span class="font-mono text-[10px] text-[#FAFAFA]/20 uppercase tracking-[0.2em]">Cancelá cuando quieras</span>
          <span class="text-[#FAFAFA]/10 hidden sm:inline" aria-hidden="true">|</span>
          <span class="font-mono text-[10px] text-[#FAFAFA]/20 uppercase tracking-[0.2em]">2 min</span>
        </div>

      </div>

    </div>

    <!-- Detalle inferior: línea fina -->
    <div class="absolute bottom-0 left-0 right-0 h-[1px] bg-gradient-to-r from-transparent via-[#D4A017]/20 to-transparent" />

  </section>
</template>
