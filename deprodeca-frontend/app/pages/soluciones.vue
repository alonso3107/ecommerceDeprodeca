<!--
  ═══════════════════════════════════════════════════════════════
  soluciones.vue — Soluciones B2B para bodegas
  Swiss Modernism 2.0 × Brutalista Corporativo: cards de impacto,
  geometría estructural, datos que venden sin adornos.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })

const soluciones = [
  {
    id: "bodega",
    titulo: "Bodega Pequeña",
    subtitulo: "Empezá sin mínimo de compra",
    desc: "Pedí desde una caja. Sin montos mínimos, sin compromisos. Ideal para bodegas que recién empiezan con DEPRODECA.",
    cifra: "S/ 0",
    etiqueta: "mínimo de compra",
    beneficios: ["Catálogo completo", "Precios B2B", "Pedidos por unidad o caja", "Soporte por chat"],
    color: "#CD7F32",
    svgPath: "M28 8 A20 20 0 1 1 28 48 A20 20 0 1 1 28 8 M28 18 A10 10 0 1 0 28 38",
  },
  {
    id: "crecimiento",
    titulo: "Bodega en Crecimiento",
    subtitulo: "Escalá tu volumen, bajá tus costos",
    desc: "Descuentos por volumen desde 5 cajas. Atención prioritaria, ofertas semanales y envíos programados para que nunca te falte stock.",
    cifra: "−15%",
    etiqueta: "en volumen +5 cajas",
    beneficios: ["Descuento por volumen", "Envíos programados", "Atención prioritaria", "Ofertas exclusivas semanales"],
    color: "#A8A8A8",
    svgPath: "M28 4 L44 20 L36 36 L20 36 L12 20 Z M28 20 V36",
  },
  {
    id: "distribuidor",
    titulo: "Distribuidor",
    subtitulo: "Tu centro de distribución, potenciado",
    desc: "Precios mayoristas con margen garantizado. Logística dedicada, gestor de cuenta personal y acceso anticipado a nuevos productos.",
    cifra: "−25%",
    etiqueta: "margen mayorista",
    beneficios: ["Precios mayoristas", "Gestor de cuenta", "Logística dedicada", "Acceso anticipado", "Crédito a 15 días"],
    color: "#D4A017",
    svgPath: "M28 4L36 12V28L28 36L20 28V12Z M28 20 A6 6 0 1 0 28 32",
  },
] as const

const ventajas = [
  { titulo: "Sin intermediarios", desc: "Directo de Nestlé a tu bodega. Sin sobrecostos." },
  { titulo: "Pedidos flexibles", desc: "Por unidad, por caja o por pallet. Vos elegís." },
  { titulo: "Pagos simples", desc: "Yape, Plin, transferencia. Sin complicaciones." },
  { titulo: "Soporte real", desc: "Chat en vivo. Respuesta en menos de 2 horas." },
] as const

const faq = [
  { q: "¿Necesito RUC para comprar?", a: "No. Aceptamos bodegas con DNI o RUC. Sin barreras." },
  { q: "¿Hay pedido mínimo?", a: "No. Podés comprar desde una unidad. Aunque en volumen ahorrás más." },
  { q: "¿Cuánto tarda la entrega?", a: "24-48h en Ica, Chincha, Pisco y Nazca. Seguimiento en vivo." },
  { q: "¿Puedo devolver productos?", a: "Sí. Devoluciones sin costo en 7 días si el producto está cerrado." },
] as const
</script>

<template>
  <div class="page-enter">
    
    <!-- ═══ HERO — Título masivo + SVG geométrico ═══ -->
    <section class="relative py-28 md:py-40 bg-[#FFFFFF] border-b border-[#E5E5E5] overflow-hidden">
      
      <svg
        aria-hidden="true"
        class="pointer-events-none absolute -right-40 top-0 h-full w-[60%] opacity-[0.05]"
        viewBox="0 0 800 600"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <circle cx="400" cy="300" r="280" stroke="#D4A017" stroke-width="1.5" />
        <circle cx="400" cy="300" r="200" stroke="#D4A017" stroke-width="2" />
        <circle cx="400" cy="300" r="120" stroke="#D4A017" stroke-width="2.5" />
        <path d="M120 300H680" stroke="#D4A017" stroke-width="1" />
        <path d="M400 20V580" stroke="#D4A017" stroke-width="1" />
        <path d="M202 98L598 502" stroke="#D4A017" stroke-width="1" />
        <path d="M598 98L202 502" stroke="#D4A017" stroke-width="1" />
      </svg>

      <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
        <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-5">
          Soluciones B2B
        </p>
        <h1 class="font-display text-[clamp(3rem,8vw,7rem)] text-[#171717] uppercase leading-[0.88] tracking-[-0.02em] max-w-[900px]">
          La solución para<br />tu bodega<span class="text-[#D4A017]">.</span>
        </h1>
        <div class="mt-6 h-[2px] w-32 bg-[#D4A017]" />
        <p class="mt-6 font-body text-[18px] text-[#666666] max-w-[580px] leading-relaxed">
          Desde tu primera caja hasta distribuidor. DEPRODECA crece con vos sin letra chica ni sorpresas.
        </p>
      </div>
    </section>

    <!-- ═══ SOLUCIONES — 3 cards de impacto ═══ -->
    <section class="py-24 md:py-32 bg-[#FAFAFA]">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        
        <div class="mb-16">
          <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-4">
            01 · Planes
          </p>
          <h2 class="font-display text-[clamp(2rem,5vw,4rem)] text-[#171717] uppercase leading-[0.92]">
            Elegí tu nivel<span class="text-[#D4A017]">.</span>
          </h2>
        </div>

        <!-- Cards verticales con diseño asimétrico -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-0 bg-[#E5E5E5]">
          
          <div
            v-for="(sol, i) in soluciones"
            :key="sol.id"
            class="group relative bg-[#FAFAFA] p-8 md:p-10 flex flex-col border-[1px] border-transparent hover:border-[#D4A017]/30 transition-all duration-500"
            :class="i === 1 ? 'md:border-x-[1px] md:border-[#E5E5E5]' : ''"
          >
            
            <!-- SVG geométrico de fondo -->
            <svg
              aria-hidden="true"
              class="absolute top-8 right-8 w-24 h-24 opacity-[0.08] group-hover:opacity-[0.15] transition-opacity duration-500"
              viewBox="0 0 56 56"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path :d="sol.svgPath" :stroke="sol.color" stroke-width="2" />
            </svg>

            <!-- Cifra grande -->
            <div class="mb-8">
              <p class="font-display text-[clamp(2.5rem,5vw,4rem)] text-[#171717] leading-none tabular-nums">
                {{ sol.cifra }}
              </p>
              <p class="mt-2 font-mono text-[10px] text-[#999999] uppercase tracking-[0.2em]">
                {{ sol.etiqueta }}
              </p>
            </div>

            <!-- Info -->
            <h3 class="font-display text-[22px] font-bold text-[#171717] uppercase tracking-wide mb-2">
              {{ sol.titulo }}
            </h3>
            <p class="font-mono text-[11px] uppercase tracking-[0.15em] mb-4" :style="{ color: sol.color }">
              {{ sol.subtitulo }}
            </p>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed mb-8 flex-1">
              {{ sol.desc }}
            </p>

            <!-- Beneficios -->
            <ul class="space-y-3 border-t border-[#E5E5E5] pt-6">
              <li
                v-for="ben in sol.beneficios"
                :key="ben"
                class="flex items-start gap-3 font-body text-[13px] text-[#666666]"
              >
                <span class="mt-1.5 w-1.5 h-1.5 shrink-0" :style="{ backgroundColor: sol.color }" />
                {{ ben }}
              </li>
            </ul>

            <!-- CTA -->
            <NuxtLink
              :to="`/auth/registro`"
              class="mt-8 inline-flex items-center gap-2 font-mono text-[11px] uppercase tracking-[0.12em] transition-all duration-300 border px-5 py-3 w-fit"
              :style="{ color: sol.color, borderColor: sol.color }"
              :class="i === 2 ? 'bg-[#D4A017] border-[#D4A017] !text-[#171717] hover:bg-[#171717] hover:!text-[#FAFAFA]' : 'hover:bg-[#171717] hover:text-[#FAFAFA] hover:border-[#171717]'"
            >
              Empezar ahora
              <span class="transition-transform duration-300 group-hover:translate-x-1">→</span>
            </NuxtLink>

          </div>
          
        </div>

      </div>
    </section>

    <!-- ═══ VENTAJAS — Grid 4 columnas ═══ -->
    <section class="py-24 md:py-32 bg-[#FFFFFF] border-t border-[#E5E5E5]">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        
        <div class="mb-16">
          <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-4">
            02 · ¿Por qué DEPRODECA?
          </p>
          <h2 class="font-display text-[clamp(2rem,5vw,4rem)] text-[#171717] uppercase leading-[0.92]">
            Simple, directo,<br />sin vueltas<span class="text-[#D4A017]">.</span>
          </h2>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-px bg-[#E5E5E5]">
          <div
            v-for="v in ventajas"
            :key="v.titulo"
            class="bg-[#FAFAFA] p-8 group hover:bg-[#171717] transition-all duration-500"
          >
            <div class="h-[2px] w-8 bg-[#D4A017] mb-5 group-hover:w-12 transition-all duration-500" />
            <h4 class="font-display text-[16px] font-semibold text-[#171717] uppercase tracking-wide mb-3 group-hover:text-[#FAFAFA] transition-colors duration-500">
              {{ v.titulo }}
            </h4>
            <p class="font-body text-[14px] text-[#666666] leading-relaxed group-hover:text-[#D4A017]/70 transition-colors duration-500">
              {{ v.desc }}
            </p>
          </div>
        </div>

      </div>
    </section>

    <!-- ═══ FAQ — Preguntas frecuentes ═══ -->
    <section class="py-24 md:py-32 bg-[#FAFAFA] border-t border-[#E5E5E5]">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        
        <div class="mb-16">
          <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-4">
            03 · Preguntas frecuentes
          </p>
          <h2 class="font-display text-[clamp(2rem,5vw,3.5rem)] text-[#171717] uppercase leading-[0.92]">
            Sin letra chica<span class="text-[#D4A017]">.</span>
          </h2>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-px bg-[#E5E5E5]">
          <div
            v-for="(item, i) in faq"
            :key="i"
            class="bg-[#FFFFFF] p-8 flex items-start gap-5 group hover:bg-[#171717] transition-all duration-500"
          >
            <span class="font-display text-[2rem] text-[#D4A017] leading-none opacity-30 group-hover:opacity-50 shrink-0 transition-opacity duration-500">
              {{ String(i + 1).padStart(2, '0') }}
            </span>
            <div>
              <h4 class="font-display text-[16px] font-semibold text-[#171717] uppercase tracking-wide mb-2 group-hover:text-[#FAFAFA] transition-colors duration-500">
                {{ item.q }}
              </h4>
              <p class="font-body text-[14px] text-[#666666] leading-relaxed group-hover:text-[#D4A017]/70 transition-colors duration-500">
                {{ item.a }}
              </p>
            </div>
          </div>
        </div>

      </div>
    </section>

    <!-- ═══ CTA FINAL ═══ -->
    <section class="py-28 md:py-40 bg-[#171717] border-t-4 border-[#D4A017] relative overflow-hidden">
      
      <svg
        aria-hidden="true"
        class="pointer-events-none absolute inset-0 h-full w-full opacity-[0.06]"
        viewBox="0 0 1200 500"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <circle cx="600" cy="250" r="220" stroke="#D4A017" stroke-width="1" />
        <circle cx="600" cy="250" r="150" stroke="#D4A017" stroke-width="1.5" />
        <path d="M380 250H820" stroke="#D4A017" stroke-width="1" />
        <path d="M600 30V470" stroke="#D4A017" stroke-width="1" />
      </svg>

      <div class="relative max-w-[1280px] mx-auto px-6 md:px-8 text-center">
        <p class="font-mono text-[11px] text-[#D4A017]/50 uppercase tracking-[0.3em] mb-6">
          Empezá hoy
        </p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,5rem)] text-[#FAFAFA] uppercase leading-[0.9] mb-6">
          Tu bodega<br />merece crecer<span class="text-[#D4A017]">.</span>
        </h2>
        <div class="h-[2px] w-24 bg-[#D4A017] mx-auto mb-8" />
        <NuxtLink
          to="/auth/registro"
          class="inline-flex items-center gap-3 bg-[#D4A017] text-[#171717] font-display text-[18px] uppercase tracking-[0.08em] px-14 py-5 hover:bg-[#FAFAFA] transition-all duration-500 group/cta"
        >
          <span>Crear Cuenta Gratis</span>
          <span class="group-hover/cta:translate-x-1 transition-transform duration-500">→</span>
        </NuxtLink>
      </div>
    </section>

  </div>
</template>
