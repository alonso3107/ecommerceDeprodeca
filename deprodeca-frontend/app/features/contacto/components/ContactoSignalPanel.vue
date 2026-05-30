<template>
  <aside class="signal-panel panel-appear relative overflow-hidden rounded-[32px] border border-[#1E1A17] bg-[#181614] text-[#FFF9F1] shadow-[0_28px_110px_-70px_rgba(24,22,20,0.75)]" :class="isReady ? 'is-visible delay-1' : ''">
    <div class="absolute inset-0 bg-[radial-gradient(circle_at_50%_25%,_rgba(212,160,23,0.26),_transparent_42%)]" aria-hidden="true" />
    <div class="relative flex h-full flex-col justify-between p-6 sm:p-7 md:p-9">
      <div>
        <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#D7B25A]">Ruta de respuesta</p>
        <p class="mt-4 max-w-[34ch] text-[15px] leading-relaxed text-[#FAF4EA]/76">
          Este panel funciona como una brújula viva: un centro de atención, tres canales y un pulso visual que marca si tu mensaje ya está en camino.
        </p>

        <div class="signal-orbit relative mx-auto mt-6 aspect-square w-full max-w-[240px] sm:mt-8 sm:max-w-[300px] md:max-w-[360px]">
          <div class="orbit orbit-1" aria-hidden="true" />
          <div class="orbit orbit-2" aria-hidden="true" />
          <div class="orbit orbit-3" aria-hidden="true" />
          <div class="orbit-core">
            <span class="font-mono text-[10px] uppercase tracking-[0.26em] text-[#D7B25A]">Pulse</span>
            <strong class="mt-2 block font-display text-3xl font-black uppercase leading-none text-[#FFF9F1]">24h</strong>
            <span class="mt-2 block text-[12px] text-[#FAF4EA]/70">ventana de respuesta</span>
          </div>
          <span class="orbit-node node-a" aria-hidden="true" />
          <span class="orbit-node node-b" aria-hidden="true" />
          <span class="orbit-node node-c" aria-hidden="true" />
        </div>
      </div>

      <div class="hidden gap-3 md:grid sm:grid-cols-3">
        <div v-for="stat in responseStats" :key="stat.label" class="rounded-[18px] border border-white/10 bg-white/5 p-4 backdrop-blur-sm">
          <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#D7B25A]">{{ stat.label }}</p>
          <p class="mt-2 font-body text-[15px] leading-tight text-[#FFF9F1]">{{ stat.value }}</p>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
defineProps<{ responseStats: Array<{ label: string; value: string }>; isReady: boolean }>()
</script>
