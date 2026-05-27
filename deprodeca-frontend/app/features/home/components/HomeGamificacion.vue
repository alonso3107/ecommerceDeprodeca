<!--
  ═══════════════════════════════════════════════════════════════
  HomeGamificacion.vue — Programa de Lealtad DEPRODECA
  Swiss Modernism 2.0 × Brutalista: indicadores geométricos,
  tabla de datos con jerarquía clara, progresión visual.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
const rangos = [
  {
    nombre: "Bronce",
    minimo: "S/ 0",
    beneficio: "Precios exclusivos B2B para tu bodega",
    nivel: 1,
    color: "#CD7F32",
    detalle: "Acceso a todo el catálogo con precios mayoristas",
    extras: ["Catálogo completo", "Precios B2B", "Soporte por chat"],
  },
  {
    nombre: "Plata",
    minimo: "S/ 5,000",
    beneficio: "Envío gratis en todos tus pedidos + atención prioritaria",
    nivel: 2,
    color: "#A8A8A8",
    detalle: "Todo lo de Bronce más envíos sin costo y fila preferencial",
    extras: ["Envío gratis ilimitado", "Atención prioritaria", "Ofertas semanales"],
  },
  {
    nombre: "Oro",
    minimo: "S/ 20,000",
    beneficio: "5% de descuento adicional en compras por volumen",
    nivel: 3,
    color: "#D4A017",
    detalle: "Todo lo de Plata con descuentos acumulables y preventas exclusivas",
    extras: ["5% descuento volumen", "Preventas exclusivas", "Regalos trimestrales"],
  },
  {
    nombre: "Platino",
    minimo: "S/ 50,000",
    beneficio: "Acceso anticipado a nuevos productos + gestor de cuenta dedicado",
    nivel: 4,
    color: "#E5E4E2",
    detalle: "El nivel máximo. Beneficios VIP, soporte personalizado y eventos exclusivos",
    extras: ["Gestor de cuenta VIP", "Acceso anticipado", "Eventos + regalos", "Crédito 15 días"],
  },
] as const
</script>

<template>
  <section class="relative py-28 md:py-36 bg-[#FAFAFA] border-t border-[#E5E5E5] overflow-hidden">
    
    <!-- SVG geométrico de fondo — grid de puntos Swiss -->
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute inset-0 h-full w-full opacity-[0.06]"
      viewBox="0 0 1200 800"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <pattern id="dots" x="0" y="0" width="60" height="60" patternUnits="userSpaceOnUse">
        <circle cx="30" cy="30" r="1.5" fill="#D4A017" />
      </pattern>
      <rect width="1200" height="800" fill="url(#dots)" />
      <path d="M0 200H1200M0 400H1200M0 600H1200" stroke="#D4A017" stroke-width="0.5" />
      <path d="M300 0V800M600 0V800M900 0V800" stroke="#D4A017" stroke-width="0.5" />
    </svg>

    <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
      
      <!-- Header Swiss -->
      <div class="mb-16 md:mb-20">
        <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-5">
          03 · Programa de Lealtad
        </p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,5rem)] text-[#171717] uppercase leading-[0.92]">
          Crece con<br />Nosotros<span class="text-[#D4A017]">.</span>
        </h2>
        <div class="mt-5 h-[2px] w-24 bg-[#D4A017]" />
        <p class="mt-5 font-body text-[16px] text-[#666666] max-w-[60ch] leading-relaxed">
          Cada compra suma puntos automáticamente. Subí de rango y desbloqueá beneficios que 
          aumentan tu margen. Sin letra chica, sin trucos.
        </p>
      </div>

      <!-- Cómo funciona — 3 pasos visuales -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-px bg-[#E5E5E5] mb-16">
        <div class="bg-[#FFFFFF] p-8 flex items-start gap-5">
          <span class="font-display text-[3rem] text-[#D4A017] leading-none opacity-30 shrink-0">01</span>
          <div>
            <h4 class="font-display text-[16px] font-semibold text-[#171717] uppercase tracking-wide mb-2">Comprá</h4>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed">Cada sol que gastás suma puntos. Sin mínimo, sin condiciones.</p>
          </div>
        </div>
        <div class="bg-[#FFFFFF] p-8 flex items-start gap-5">
          <span class="font-display text-[3rem] text-[#D4A017] leading-none opacity-30 shrink-0">02</span>
          <div>
            <h4 class="font-display text-[16px] font-semibold text-[#171717] uppercase tracking-wide mb-2">Subí</h4>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed">Tus puntos te llevan al siguiente rango. El progreso nunca se reinicia.</p>
          </div>
        </div>
        <div class="bg-[#FFFFFF] p-8 flex items-start gap-5">
          <span class="font-display text-[3rem] text-[#D4A017] leading-none opacity-30 shrink-0">03</span>
          <div>
            <h4 class="font-display text-[16px] font-semibold text-[#171717] uppercase tracking-wide mb-2">Disfrutá</h4>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed">Beneficios que aumentan tu margen. Envíos gratis, descuentos, regalos.</p>
          </div>
        </div>
      </div>

      <!-- Tabla de rangos — Swiss: filas con jerarquía clara -->
      <div class="border border-[#E5E5E5] bg-[#FFFFFF] overflow-hidden">
        
        <!-- Header row -->
        <div class="hidden md:grid grid-cols-[2fr_1fr_1fr_2fr_1fr] bg-[#171717] text-[#FAFAFA] font-mono text-[10px] uppercase tracking-[0.25em] px-8 py-5 items-center">
          <span>Rango</span>
          <span>Desde</span>
          <span>Progreso</span>
          <span>Beneficio principal</span>
          <span class="text-right">Extras</span>
        </div>

        <!-- Data rows -->
        <div
          v-for="(rango, i) in rangos"
          :key="rango.nombre"
          class="grid grid-cols-1 md:grid-cols-[2fr_1fr_1fr_2fr_1fr] border-t border-[#E5E5E5] px-6 md:px-8 py-6 hover:bg-[#D4A017]/[0.02] transition-colors duration-300 items-start gap-4 md:gap-0"
        >
          
          <!-- Rango: SVG geométrico + nombre + detalle -->
          <div class="flex items-start gap-4">
            <!-- SVG icon geométrico (reemplaza emojis) -->
            <svg 
              :viewBox="rango.nivel === 1 ? '0 0 40 40' : rango.nivel === 4 ? '0 0 48 48' : '0 0 40 40'"
              class="w-10 h-10 shrink-0 mt-0.5"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <!-- Bronce: círculo -->
              <template v-if="rango.nivel === 1">
                <circle cx="20" cy="20" r="16" :stroke="rango.color" stroke-width="2.5" />
                <circle cx="20" cy="20" r="8" :stroke="rango.color" stroke-width="1.5" fill="none" opacity="0.4" />
              </template>
              <!-- Plata: rombo -->
              <template v-else-if="rango.nivel === 2">
                <path d="M20 4L36 20L20 36L4 20Z" :stroke="rango.color" stroke-width="2.5" />
                <path d="M20 12L28 20L20 28L12 20Z" :stroke="rango.color" stroke-width="1.5" opacity="0.5" />
              </template>
              <!-- Oro: hexágono -->
              <template v-else-if="rango.nivel === 3">
                <path d="M20 4L36 12V28L20 36L4 28V12Z" :stroke="rango.color" stroke-width="2.5" />
                <circle cx="20" cy="20" r="5" :stroke="rango.color" stroke-width="1.5" opacity="0.5" />
              </template>
              <!-- Platino: estrella 6 puntas -->
              <template v-else>
                <path d="M24 4L28 16L40 20L28 24L24 36L20 24L8 20L20 16Z" :stroke="rango.color" stroke-width="2" />
                <circle cx="24" cy="20" r="6" :stroke="rango.color" stroke-width="1" opacity="0.4" />
              </template>
            </svg>

            <div>
              <span class="font-display text-[18px] font-bold text-[#171717] uppercase tracking-wide">{{ rango.nombre }}</span>
              <p class="mt-1 font-body text-[13px] text-[#888888] leading-snug hidden md:block">
                {{ rango.detalle }}
              </p>
            </div>
          </div>

          <!-- Desde (mínimo) -->
          <div class="md:text-center">
            <span class="font-mono text-[14px] text-[#171717] font-semibold tracking-tight">{{ rango.minimo }}</span>
            <span class="block font-mono text-[9px] text-[#999999] uppercase tracking-[0.2em] mt-0.5">acumulado</span>
          </div>

          <!-- Nivel: barras de progreso -->
          <div class="flex items-center gap-3">
            <div class="flex gap-1.5">
              <div
                v-for="n in 4"
                :key="n"
                class="w-5 h-2.5 border transition-colors duration-300"
                :class="n <= rango.nivel
                  ? 'bg-[#D4A017] border-[#D4A017]'
                  : 'bg-transparent border-[#E5E5E5]'"
              />
            </div>
            <span class="font-mono text-[11px] text-[#999999]">{{ rango.nivel }}/4</span>
          </div>

          <!-- Beneficio principal -->
          <div>
            <p class="font-body text-[14px] text-[#D4A017] font-semibold leading-snug">
              {{ rango.beneficio }}
            </p>
          </div>

          <!-- Extras -->
          <div class="md:text-right">
            <div class="flex flex-wrap gap-1.5 md:justify-end">
              <span
                v-for="extra in rango.extras"
                :key="extra"
                class="inline-block font-mono text-[9px] text-[#666666] bg-[#F5F5F5] px-2 py-1 uppercase tracking-[0.08em]"
              >
                {{ extra }}
              </span>
            </div>
          </div>

        </div>
        
      </div>

      <!-- CTA inferior -->
      <div class="mt-12 flex flex-col sm:flex-row items-start sm:items-center gap-6">
        <NuxtLink
          to="/perfil/rangos"
          class="inline-flex items-center gap-2 font-mono text-[11px] text-[#171717] hover:text-[#D4A017] uppercase tracking-[0.15em] transition-colors duration-300 border border-[#E5E5E5] hover:border-[#D4A017] px-6 py-3 group/cta"
        >
          <span>Ver todos los beneficios</span>
          <span class="text-[#D4A017] group-hover/cta:translate-x-1 transition-transform duration-300">→</span>
        </NuxtLink>
        <p class="font-mono text-[10px] text-[#AAAAAA] uppercase tracking-[0.1em]">
          Tus puntos no expiran. Progreso de por vida.
        </p>
      </div>

    </div>
  </section>
</template>
