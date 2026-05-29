<script setup lang="ts">
function irARegistro() {
  navigateTo("/auth/registro")
}
</script>

<template>
  <section class="py-24 md:py-32 bg-[#A16207] border-t border-[#D6D3D1]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="max-w-[780px] mx-auto text-center">
        <p class="font-mono text-[11px] text-[#1C1917]/60 uppercase tracking-[0.2em] mb-7">Sin costo de inscripcion</p>

        <h2 class="font-display text-[clamp(3rem,8vw,6rem)] font-bold text-[#1C1917] leading-[0.9] mb-6">
          Empeza a comprar
        </h2>

        <div class="h-[2px] w-24 bg-[#1C1917] mx-auto mb-8" />

        <p class="font-body text-[16px] text-[#1C1917]/60 max-w-lg mx-auto mb-12 leading-relaxed">
          Registrate gratis en minutos y accede a precios exclusivos para tu bodega.
        </p>

        <button
          class="inline-flex items-center bg-[#1C1917] text-[#FAFAF9] font-bold px-12 py-4 transition-colors duration-300 hover:bg-[#FAFAF9] hover:text-[#1C1917]"
          @click="irARegistro"
        >
          Crear Cuenta Gratis
        </button>

        <div class="mt-10 flex flex-wrap items-center justify-center gap-x-6 gap-y-2">
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">Sin tarjeta</span>
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">Sin compromiso</span>
          <span class="font-mono text-[10px] text-[#1C1917]/40 uppercase tracking-[0.2em]">2 min</span>
        </div>
      </div>
    </div>
  </section>
</template>
