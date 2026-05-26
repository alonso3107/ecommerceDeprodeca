<!--
  ═══════════════════════════════════════════════════════════════
  HomeCtaFinal.vue — Sección de cierre con llamado a la acción.
  Brutalismo en negativo: fondo negro, tipografía blanca masiva,
  botón dorado como único acento de color.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
function irARegistro() {
  navigateTo("/auth/registro")
}
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    CTA FINAL — Fondo negro puro, texto blanco masivo.
    Sin gradientes, sin sombras, sin adornos. Solo el mensaje.
    ═══════════════════════════════════════════════════════════
  -->
  <section class="py-36 md:py-48 bg-texto relative overflow-hidden border-t-4 border-[#D4A017]">

    <!-- Detalle técnico esquina -->
    <div class="absolute top-6 left-6 md:top-8 md:left-8 font-mono text-[10px] text-white/20 uppercase tracking-[0.3em] select-none" aria-hidden="true">
      ─── 04 · Acción
    </div>

    <div class="relative max-w-[1280px] mx-auto px-6 md:px-8 text-center">
      <!-- Tipografía masiva -->
      <h2 class="font-display text-[clamp(3rem,8vw,6rem)] text-white uppercase leading-[0.9] tracking-[-0.03em] mb-6">
        Empezá a<br />comprar<span class="text-[#D4A017]">.</span>
      </h2>

      <p class="font-body text-body text-white/40 max-w-md mx-auto mb-14 leading-relaxed">
        Registrate gratis. Sin costo de inscripción.
        Accedé a precios exclusivos para tu bodega.
      </p>

      <!-- Botón — dorado sólido, sin sombras, brutalista -->
      <button
        class="inline-block bg-[#D4A017] text-black font-display text-heading uppercase
               tracking-[0.05em] px-12 py-5
               hover:bg-white transition-colors duration-300
               active:scale-[0.98]"
        @click="irARegistro"
      >
        Crear Cuenta Gratis
      </button>

      <!-- Detalle inferior: sin tarjeta, sin compromiso -->
      <p class="mt-8 font-mono text-[10px] text-white/15 uppercase tracking-[0.2em]">
        Sin tarjeta · Sin compromiso · Cancelá cuando quieras
      </p>
    </div>

  </section>
</template>
