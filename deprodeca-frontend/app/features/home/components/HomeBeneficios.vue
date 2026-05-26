<!--
  ═══════════════════════════════════════════════════════════════
  HomeBeneficios.vue — Ventajas competitivas DEPRODECA
  Brutalismo suave: bordes duros, alto contraste, datos crudos.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Datos estáticos · sin reactividad innecesaria ────────
const beneficios = [
  {
    titulo: "Precios B2B",
    desc: "Sin intermediarios. Precios directos de Nestlé para tu bodega.",
    cifra: "−30%",
    etiqueta: "vs retail",
  },
  {
    titulo: "Pedidos al Mayor",
    desc: "Comprá por caja o unidad. Stock siempre disponible.",
    cifra: "500+",
    etiqueta: "productos",
  },
  {
    titulo: "Entrega 24-48h",
    desc: "Despacho rápido a todo Ica. Seguimiento en vivo.",
    cifra: "48h",
    etiqueta: "máximo",
  },
  {
    titulo: "Gamificación",
    desc: "Subí de Bronce a Platino. Cada compra te da puntos y beneficios.",
    cifra: "4",
    etiqueta: "rangos",
  },
] as const
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    BENEFICIOS — Grid 4-columnas con datos crudos.
    Estética brutalista: bordes finos, números grandes, sin
    sombras pesadas, enfoque en la información.
    ═══════════════════════════════════════════════════════════
  -->
  <section class="py-24 md:py-32 bg-white border-t border-[#D4A017]/20">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">

      <!-- Encabezado técnico -->
      <div class="mb-16">
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-4">
          ─── 01 · Ventajas
        </p>
        <h2 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
          ¿Por qué<br />DEPRODECA<span class="text-[#D4A017]">?</span>
        </h2>
      </div>

      <!-- Grid de beneficios — brutalista: bordes finos, datos first -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-px bg-borde">
        <div
          v-for="b in beneficios"
          :key="b.titulo"
          class="bg-white p-8 md:p-10 flex flex-col justify-between gap-6"
        >
          <!-- Cifra grande — brutalist data display -->
          <div>
            <p class="font-display text-[clamp(2.5rem,5vw,3.5rem)] text-texto leading-none">
              {{ b.cifra }}
            </p>
            <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider mt-1">
              {{ b.etiqueta }}
            </p>
          </div>

          <!-- Título + descripción -->
          <div>
            <h3 class="font-body text-body font-bold text-texto mb-2">
              {{ b.titulo }}
            </h3>
            <p class="font-body text-small text-texto-muted leading-relaxed">
              {{ b.desc }}
            </p>
          </div>
        </div>
      </div>

    </div>
  </section>
</template>
