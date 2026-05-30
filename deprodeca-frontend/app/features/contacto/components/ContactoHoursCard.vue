<template>
  <section class="panel-appear is-visible delay-3 rounded-[28px] border border-[#E5D8C2] bg-[#FFF9F1] p-6 shadow-[0_26px_90px_-68px_rgba(24,22,20,0.52)] md:p-8">
    <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#A16207]">Horarios</p>
    <div class="mt-5 space-y-3">
      <div v-for="h in horarios" :key="h.dia" class="flex items-center justify-between rounded-[18px] border border-[#E8DDD0] bg-white px-4 py-3">
        <span class="font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">{{ h.dia }}</span>
        <span class="font-display text-[14px] font-bold uppercase tracking-[0.08em] text-[#181614]">{{ h.hora }}</span>
      </div>
    </div>

    <div class="mt-6 rounded-[22px] bg-[#181614] p-5 text-[#FFF9F1]">
      <p class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#D7B25A]">Detalle útil</p>
      <p class="mt-3 text-[15px] leading-relaxed text-[#FAF4EA]/78">
        Si escribís desde una bodega, agregá tu distrito y el tipo de pedido. Eso acelera mucho la respuesta.
      </p>
    </div>
  </section>
</template>

<script setup lang="ts">
defineProps<{ horarios: Array<{ dia: string; hora: string }> }>()
</script>
