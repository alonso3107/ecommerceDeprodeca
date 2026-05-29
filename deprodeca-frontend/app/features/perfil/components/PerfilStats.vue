<script setup lang="ts">
const props = defineProps<{
  stats: {
    puntos: number
    pedidos: number
    volumen: number
    descuento: number
  }
}>()

const moneyFormatter = new Intl.NumberFormat("es-PE", {
  style: "currency",
  currency: "PEN",
  maximumFractionDigits: 0,
})
</script>

<template>
  <section class="bg-white border border-borde border-t-0">
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-px bg-borde">
      <article class="bg-white p-6 hover:bg-fondo transition-colors duration-300">
        <div class="flex items-center gap-2 mb-2">
          <svg width="12" height="12" viewBox="0 0 12 12" fill="none"><circle cx="6" cy="6" r="4" stroke="#D4A017" /></svg>
          <p class="font-display text-display-md text-texto">{{ stats.puntos.toLocaleString("es-PE") }}</p>
        </div>
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Puntos</p>
      </article>
      <article class="bg-white p-6 hover:bg-fondo transition-colors duration-300">
        <div class="flex items-center gap-2 mb-2">
          <svg width="12" height="12" viewBox="0 0 12 12" fill="none"><rect x="2" y="2" width="8" height="8" stroke="#171717" /></svg>
          <p class="font-display text-display-md text-texto">{{ stats.pedidos.toLocaleString("es-PE") }}</p>
        </div>
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Pedidos</p>
      </article>
      <article class="bg-white p-6 hover:bg-fondo transition-colors duration-300">
        <p class="font-display text-display-md text-texto mb-2">{{ moneyFormatter.format(stats.volumen || 0) }}</p>
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Volumen</p>
      </article>
      <article class="bg-white p-6 hover:bg-fondo transition-colors duration-300">
        <p class="font-display text-display-md text-texto mb-2">{{ stats.descuento }}%</p>
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Descuento</p>
      </article>
    </div>
  </section>
</template>
