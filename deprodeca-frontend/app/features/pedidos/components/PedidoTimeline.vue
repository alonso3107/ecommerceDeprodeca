<script setup lang="ts">
const props = defineProps<{
  estado: string
  fechaConfirmado?: string
  fechaEnviado?: string
  fechaEntregado?: string
}>()

const pasos = ["pendiente", "confirmado", "enviado", "entregado"] as const

const estadoNormalizado = computed(() => {
  if (props.estado === "en_proceso") return "confirmado"
  return props.estado
})

const indiceActual = computed(() => pasos.indexOf(estadoNormalizado.value as (typeof pasos)[number]))

const config = {
  pendiente: { label: "Pendiente", color: "#F59E0B" },
  confirmado: { label: "Confirmado", color: "#171717" },
  enviado: { label: "Enviado", color: "#737373" },
  entregado: { label: "Entregado", color: "#16A34A" },
} as const

function fechaPaso(paso: string) {
  if (paso === "confirmado") return props.fechaConfirmado
  if (paso === "enviado") return props.fechaEnviado
  if (paso === "entregado") return props.fechaEntregado
  return ""
}

function formatearFecha(fecha?: string) {
  if (!fecha) return ""
  return new Date(fecha).toLocaleDateString("es-PE", {
    day: "numeric",
    month: "short",
    hour: "2-digit",
    minute: "2-digit",
  })
}

function esCompletado(i: number) {
  return indiceActual.value >= 0 && i < indiceActual.value
}

function esActual(i: number) {
  return indiceActual.value === i
}

function esFuturo(i: number) {
  return indiceActual.value < i || indiceActual.value === -1
}

function colorConector(i: number) {
  const paso = pasos[i] || "pendiente"
  return i < indiceActual.value ? config[paso].color : "#E5E5E5"
}
</script>

<template>
  <section class="border border-borde bg-white p-6">
    <div v-if="estado === 'cancelado'" class="flex items-center gap-4">
      <div class="w-12 h-12 rounded-full border-2 border-[#DC2626] bg-[#DC2626] flex items-center justify-center">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path d="M6 6L18 18M18 6L6 18" stroke="white" stroke-width="2" />
        </svg>
      </div>
      <div>
        <p class="font-mono text-[11px] uppercase tracking-[0.18em] text-[#DC2626]">Pedido Cancelado</p>
        <p class="font-mono text-[9px] text-texto-muted mt-1">{{ formatearFecha(fechaConfirmado) }}</p>
      </div>
    </div>

    <div v-else>
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 md:gap-0">
        <div v-for="(paso, i) in pasos" :key="paso" class="relative flex md:block items-start md:items-center gap-4">
          <div
            v-if="i < pasos.length - 1"
            class="hidden md:block absolute left-[calc(50%+24px)] top-[23px] h-[2px] w-[calc(100%-48px)]"
            :style="{ backgroundColor: colorConector(i) }"
          />
          <div
            v-if="i < pasos.length - 1"
            class="md:hidden absolute left-[23px] top-[48px] w-[2px] h-[54px]"
            :style="{ backgroundColor: colorConector(i) }"
          />

          <div
            class="w-12 h-12 rounded-full border-2 flex items-center justify-center shrink-0"
            :style="{
              borderColor: esFuturo(i) ? '#E5E5E5' : config[paso].color,
              backgroundColor: esFuturo(i) ? '#FFFFFF' : config[paso].color,
            }"
          >
            <svg v-if="esCompletado(i)" width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path d="M5 12.5L10 17.5L19 7.5" stroke="white" stroke-width="2" />
            </svg>
            <template v-else>
              <svg v-if="paso === 'pendiente'" width="24" height="24" viewBox="0 0 24 24" fill="none">
                <circle cx="12" cy="12" r="8" :stroke="esActual(i) ? 'white' : '#F59E0B'" stroke-width="2" />
                <path d="M12 8V12L15 14" :stroke="esActual(i) ? 'white' : '#F59E0B'" stroke-width="2" />
              </svg>
              <svg v-else-if="paso === 'confirmado'" width="24" height="24" viewBox="0 0 24 24" fill="none">
                <rect x="5" y="5" width="14" height="14" :stroke="esActual(i) ? 'white' : '#171717'" stroke-width="2" />
                <path d="M8 12L11 15L16 9" :stroke="esActual(i) ? 'white' : '#171717'" stroke-width="2" />
              </svg>
              <svg v-else-if="paso === 'enviado'" width="24" height="24" viewBox="0 0 24 24" fill="none">
                <rect x="3" y="7" width="12" height="10" :stroke="esActual(i) ? 'white' : '#737373'" stroke-width="2" />
                <path d="M15 10H19L21 12V17H15V10Z" :stroke="esActual(i) ? 'white' : '#737373'" stroke-width="2" />
                <circle cx="8" cy="18" r="1.5" :fill="esActual(i) ? 'white' : '#737373'" />
                <circle cx="18" cy="18" r="1.5" :fill="esActual(i) ? 'white' : '#737373'" />
              </svg>
              <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none">
                <rect x="4" y="6" width="16" height="12" :stroke="esActual(i) ? 'white' : '#16A34A'" stroke-width="2" />
                <path d="M4 10L12 14L20 10" :stroke="esActual(i) ? 'white' : '#16A34A'" stroke-width="2" />
                <path d="M9 12L11 14L15 10" :stroke="esActual(i) ? 'white' : '#16A34A'" stroke-width="2" />
              </svg>
            </template>
          </div>

          <div class="md:mt-3 md:text-center">
            <p class="font-mono text-[11px] uppercase tracking-[0.14em]" :style="{ color: esFuturo(i) ? '#A3A3A3' : '#171717' }">
              {{ config[paso].label }}
            </p>
            <p class="font-mono text-[9px] text-texto-muted mt-1">
              {{ formatearFecha(fechaPaso(paso)) }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
