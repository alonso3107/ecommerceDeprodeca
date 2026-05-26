<!--
  ═══════════════════════════════════════════════════════════════
  HomeGamificacion.vue — Programa de Lealtad DEPRODECA
  Brutalismo de datos: rangos como tabla cruda, progreso visual
  sin adornos.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Rangos del programa de lealtad ───────────────────────
const rangos = [
  { nombre: "Bronce",  minimo: "S/ 0",       icono: "🥉", beneficio: "Precios exclusivos B2B",           nivel: 1 },
  { nombre: "Plata",   minimo: "S/ 5,000",   icono: "🥈", beneficio: "Envío gratis + prioridad",         nivel: 2 },
  { nombre: "Oro",     minimo: "S/ 20,000",  icono: "🥇", beneficio: "5% descuento en volumen",          nivel: 3 },
  { nombre: "Platino", minimo: "S/ 50,000",  icono: "💎", beneficio: "Acceso anticipado + soporte VIP",  nivel: 4 },
] as const
</script>

<template>
  <!--
    ═══════════════════════════════════════════════════════════
    GAMIFICACIÓN — Tabla de rangos con estética brutalista.
    Sin cards redondeadas: filas con bordes duros, datos crudos.
    ═══════════════════════════════════════════════════════════
  -->
  <section class="py-24 md:py-32 bg-white border-t border-borde">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">

      <!-- Encabezado técnico -->
      <div class="mb-16">
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-4">
          ─── 03 · Lealtad
        </p>
        <h2 class="font-display text-display-lg text-texto uppercase leading-[0.95] mb-4">
          Crece con<br />Nosotros<span class="text-[#D4A017]">.</span>
        </h2>
        <p class="font-body text-body text-texto-muted max-w-lg">
          Cada compra suma puntos. Subí de rango y desbloqueá beneficios exclusivos.
        </p>
      </div>

      <!-- Tabla de rangos — brutalista: filas, no cards -->
      <div class="border border-borde overflow-hidden">
        <!-- Header row -->
        <div class="hidden md:grid grid-cols-[1fr_1.5fr_1fr_2fr] bg-texto text-white font-mono text-[11px] uppercase tracking-[0.2em] px-6 py-4">
          <span>Rango</span>
          <span>Desde</span>
          <span>Nivel</span>
          <span>Beneficio</span>
        </div>

        <!-- Data rows -->
        <div
          v-for="(rango, i) in rangos"
          :key="rango.nombre"
          class="grid grid-cols-1 md:grid-cols-[1fr_1.5fr_1fr_2fr] border-t border-borde px-6 py-5
                 hover:bg-[#D4A017]/[0.03] transition-colors duration-200
                 items-center gap-2 md:gap-0"
        >
          <!-- Rango nombre + ícono -->
          <div class="flex items-center gap-3">
            <span class="text-xl" aria-hidden="true">{{ rango.icono }}</span>
            <span class="font-display text-heading text-texto uppercase">{{ rango.nombre }}</span>
          </div>

          <!-- Mínimo -->
          <span class="font-body text-body text-texto-muted md:text-center">{{ rango.minimo }}</span>

          <!-- Nivel (barra visual) -->
          <div class="flex items-center gap-2">
            <div class="flex gap-0.5">
              <div
                v-for="n in 4"
                :key="n"
                class="w-3 h-3 border"
                :class="n <= rango.nivel
                  ? 'bg-[#D4A017] border-[#D4A017]'
                  : 'bg-transparent border-borde'"
              />
            </div>
            <span class="font-mono text-[10px] text-texto-muted">{{ rango.nivel }}/4</span>
          </div>

          <!-- Beneficio -->
          <span class="font-body text-small text-[#D4A017] font-semibold">{{ rango.beneficio }}</span>
        </div>
      </div>

      <!-- CTA inferior -->
      <div class="mt-10">
        <NuxtLink
          to="/perfil/rangos"
          class="inline-block font-mono text-[11px] text-texto-muted hover:text-[#D4A017] uppercase tracking-widest transition-colors duration-200 border-b border-transparent hover:border-[#D4A017] pb-0.5"
        >
          Ver todos los beneficios →
        </NuxtLink>
      </div>

    </div>
  </section>
</template>
