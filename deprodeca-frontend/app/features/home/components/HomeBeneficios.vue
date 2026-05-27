<!--
  ═══════════════════════════════════════════════════════════════
  HomeBeneficios.vue — Ventajas competitivas DEPRODECA
  Swiss Modernism 2.0 × Brutalista: grid matemático, tipografía
  estructural, datos crudos con jerarquía limpia.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
const beneficios = [
  {
    titulo: "Precios B2B",
    desc: "Sin intermediarios. Precios directos de Nestlé para tu bodega. Ahorrá hasta un 30% frente al retail tradicional.",
    cifra: "−30%",
    etiqueta: "vs retail",
    detalle: "Margen superior garantizado en cada pedido",
  },
  {
    titulo: "Pedidos al Mayor",
    desc: "Comprá por caja o unidad. Stock siempre disponible con reposición automática.",
    cifra: "500+",
    etiqueta: "productos",
    detalle: "Catálogo actualizado semanalmente",
  },
  {
    titulo: "Entrega 24-48h",
    desc: "Despacho rápido a todo Ica. Seguimiento en vivo desde que sale de nuestro almacén.",
    cifra: "48h",
    etiqueta: "máximo",
    detalle: "Cobertura: Ica, Chincha, Pisco, Nazca",
  },
  {
    titulo: "Gamificación",
    desc: "Subí de Bronce a Platino. Cada compra te da puntos y beneficios exclusivos.",
    cifra: "4",
    etiqueta: "rangos",
    detalle: "Beneficios acumulables sin caducidad",
  },
] as const
</script>

<template>
  <section class="relative py-28 md:py-36 bg-[#FAFAFA] border-t border-[#E5E5E5] overflow-hidden">
    
    <!-- SVG geométrico de fondo — Swiss asimétrico -->
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute inset-0 h-full w-full opacity-[0.07]"
      viewBox="0 0 1440 600"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M-80 120H1520" stroke="#D4A017" stroke-width="1.5" />
      <path d="M-80 300H1520" stroke="#D4A017" stroke-width="1.5" />
      <path d="M-80 480H1520" stroke="#D4A017" stroke-width="1.5" />
      <path d="M320 0V600" stroke="#D4A017" stroke-width="1.5" />
      <path d="M720 0V600" stroke="#D4A017" stroke-width="1.5" />
      <path d="M1120 0V600" stroke="#D4A017" stroke-width="1.5" />
      <circle cx="720" cy="300" r="180" stroke="#D4A017" stroke-width="2" />
      <circle cx="720" cy="300" r="120" stroke="#D4A017" stroke-width="1" />
      <path d="M540 300H900" stroke="#D4A017" stroke-width="2" />
      <path d="M720 120V480" stroke="#D4A017" stroke-width="2" />
    </svg>

    <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
      
      <!-- Header Swiss — label técnico + título con punto -->
      <div class="mb-16 md:mb-20">
        <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-5">
          01 · Ventajas
        </p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,5rem)] text-[#171717] uppercase leading-[0.92]">
          ¿Por qué<br />DEPRODECA<span class="text-[#D4A017]">?</span>
        </h2>
        <div class="mt-6 h-[2px] w-24 bg-[#D4A017]" />
        <p class="mt-5 font-body text-[16px] text-[#666666] max-w-[55ch] leading-relaxed">
          Ventajas reales para tu bodega. Sin letra chica, sin sorpresas.
        </p>
      </div>

      <!-- Grid 4 columnas — Swiss matemático con gap generoso -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-0 bg-[#E5E5E5]">
        
        <div
          v-for="(b, i) in beneficios"
          :key="b.titulo"
          class="group bg-[#FAFAFA] p-8 md:p-10 flex flex-col justify-between gap-8 border-[1px] border-transparent hover:border-[#D4A017]/30 transition-all duration-500"
          :class="i === 0 ? 'sm:col-span-2 lg:col-span-1' : ''"
        >
          
          <!-- Cifra masiva + línea decorativa -->
          <div>
            <div class="flex items-baseline gap-2">
              <p class="font-display text-[clamp(2.8rem,5vw,4rem)] text-[#171717] leading-none tabular-nums">
                {{ b.cifra }}
              </p>
            </div>
            <div class="mt-3 h-[2px] w-12 bg-[#D4A017]/40 group-hover:w-20 group-hover:bg-[#D4A017] transition-all duration-500" />
            <p class="mt-3 font-mono text-[10px] text-[#999999] uppercase tracking-[0.2em]">
              {{ b.etiqueta }}
            </p>
          </div>

          <!-- Info — título + descripción + detalle Swiss -->
          <div class="space-y-3">
            <h3 class="font-display text-[18px] font-semibold text-[#171717] uppercase tracking-wide">
              {{ b.titulo }}
            </h3>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed">
              {{ b.desc }}
            </p>
            <p class="font-mono text-[10px] text-[#D4A017]/70 uppercase tracking-[0.12em] pt-1 border-t border-[#E5E5E5]">
              {{ b.detalle }}
            </p>
          </div>

        </div>
        
      </div>

    </div>
  </section>
</template>
