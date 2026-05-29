<!--
  ═══════════════════════════════════════════════════════════════
  HomeHero.vue — DEPRODECA
  Estética: Swiss-Modernist × Digital Brutalism
  Alto contraste, cero sombras, alineación quirúrgica, tipografía masiva.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Navegación ───────────────────────────────────────────
// Rutas principales del hero para navegación directa
const enlaces = [
  { texto: "PRODUCTOS",   ruta: "/catalogo" },
  { texto: "SOLUCIONES",  ruta: "/soluciones" },
  { texto: "TIENDA",      ruta: "/catalogo" },
  { texto: "CONTACTO",    ruta: "/contacto" },
] as const

function navegar(ruta: string) {
  navigateTo(ruta)
}
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    HERO — Fondo sólido dorado (#D4A017), sin gradientes,
    sin sombras. Brutalismo puro: la tipografía es el diseño.
    ═══════════════════════════════════════════════════════════
  -->
  <section
    class="relative bg-[#A16207] min-h-screen flex flex-col overflow-hidden"
    aria-label="Hero principal DEPRODECA"
  >

    <!--
      ─── DETALLES TÉCNICOS · ESQUINAS ─────────────────────
      Monospace pequeño, fine-line. Información corporativa
      que refuerza la estética brutalista-industrial.
    -->
    <!-- TOP-LEFT: Sede central -->
    <!-- Aumentado contraste: black/50 → black/60 para mejor accesibilidad sobre #D4A017 -->
    <div class="absolute top-6 left-6 md:top-8 md:left-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wider uppercase select-none" aria-hidden="true">
      <p>Ica, Perú</p>
      <p>Sede Central</p>
    </div>

    <!-- TOP-RIGHT: Centro de distribución -->
    <div class="absolute top-6 right-6 md:top-8 md:right-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wider uppercase text-right select-none" aria-hidden="true">
      <p>Centro de</p>
      <p>Distribución</p>
    </div>

    <!-- BOTTOM-LEFT: Misión -->
    <div class="absolute bottom-6 left-6 md:bottom-8 md:left-8 font-mono text-[10px] md:text-[11px] text-[#1C1917]/50 leading-tight tracking-wide uppercase max-w-[200px] select-none" aria-hidden="true">
      <p>Distribución Premium</p>
      <p>Calidad Nestlé</p>
    </div>

    <!-- BOTTOM-RIGHT: CTA "Comprar Ahora" -->
    <div class="absolute bottom-6 right-6 md:bottom-8 md:right-8">
      <button
        class="font-mono text-[11px] md:text-xs text-[#1C1917]/60 hover:text-[#1C1917] border border-[#1C1917]/40 hover:border-[#1C1917] px-4 py-2 transition-colors duration-300 uppercase tracking-widest"
        @click="navegar('/catalogo')"
      >
        Comprar Ahora
      </button>
    </div>

    <!--
      ─── NAVBAR · CÁPSULA FLOTANTE ─────────────────────────
      Glassmorphism: frosted glass, centrada, flotante.
      La marca DEPRODECA como ancla visual superior.
    -->
    <nav
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-[42vh] z-10
             bg-[#1C1917] text-[#FAFAF9] font-bold px-8 py-3 border-b-2 border-[#FAFAF9]/30"
      aria-label="Navegación principal"
    >
      <span class="font-display text-xl md:text-2xl tracking-[0.15em] uppercase select-none">
        DEPRODECA
      </span>
    </nav>

    <!--
      ─── TIPOGRAFÍA MASIVA · STACKED ───────────────────────
      Las palabras apiladas verticalmente, condensadas,
      negras sobre dorado. Este ES el diseño.
    -->
    <div class="flex-1 flex items-center justify-center px-4">
      <div class="flex flex-col items-start gap-0 leading-none">
        <NuxtLink
          v-for="enlace in enlaces"
          :key="enlace.texto"
          :to="enlace.ruta"
          class="font-display text-[clamp(4rem,12vw,9rem)] text-black
                 leading-[0.82] tracking-[-0.04em] uppercase
                 hover:text-[#1C1917] transition-colors duration-300
                 cursor-pointer select-none"
        >
          {{ enlace.texto }}
        </NuxtLink>
      </div>
    </div>

    <!--
      ─── LÍNEA INFERIOR · DETALLE TÉCNICO ──────────────────
      Una línea sutil que recorre el borde inferior, con
      metadata corporativa en monospace.
    -->
    <div class="absolute bottom-0 left-0 right-0 border-t border-black/15 flex justify-between px-6 md:px-8 py-3 font-mono text-[9px] md:text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em] select-none" aria-hidden="true">
      <span>Filial Oficial Nestlé Perú</span>
      <span>B2B · Mayorista</span>
    </div>

  </section>
</template>

<style scoped>
/* Sin estilos innecesarios — el brutalismo vive en las clases utilitarias.
   Solo animaciones mínimas y esenciales. */
</style>
