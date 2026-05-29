<script setup lang="ts">
const props = defineProps<{
  gamificacion: {
    rango?: string
    volumen_total?: number
    puntos_totales?: number
  } | null
  leaderboard: Array<{
    posicion: number
    usuario_empresa: string
    rango: string
    puntos_totales: number
  }>
  wsConnected: boolean
}>()

const monto = new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN", maximumFractionDigits: 0 })
const rango = computed(() => (props.gamificacion?.rango || "bronce").toLowerCase())
const volumen = computed(() => props.gamificacion?.volumen_total || 0)

const umbrales: Record<string, number> = { bronce: 5000, plata: 20000, oro: 50000, platino: 50000 }
const siguiente: Record<string, string> = { bronce: "plata", plata: "oro", oro: "platino", platino: "platino" }

const tope = computed(() => umbrales[rango.value] || 5000)
const progreso = computed(() => Math.min((volumen.value / tope.value) * 100, 100))

function colorRango(r: string) {
  return {
    bronce: { hex: "#CD7F32", bg: "#CD7F32" },
    plata: { hex: "#A8A8A8", bg: "#A8A8A8" },
    oro: { hex: "#D4A017", bg: "#D4A017" },
    platino: { hex: "#7B2FBE", bg: "#7B2FBE" },
  }[r] || { hex: "#CD7F32", bg: "#CD7F32" }
}

const beneficios = computed(() => {
  const out = ["Precios B2B"]
  if (["plata", "oro", "platino"].includes(rango.value)) out.push("Envio gratis")
  if (["oro", "platino"].includes(rango.value)) out.push("5% desc")
  if (rango.value === "platino") out.push("Acceso VIP")
  return out
})
</script>

<template>
  <section>
    <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-texto-muted mb-3">--- Mi Rango</p>
    <div class="grid grid-cols-1 lg:grid-cols-3 border border-borde bg-white">
      <div class="lg:col-span-2 border-b lg:border-b-0 lg:border-r border-borde p-6 md:p-8">
        <div class="flex items-center gap-4 mb-6">
          <div class="w-16 h-16 border border-borde flex items-center justify-center">
            <svg v-if="rango === 'bronce'" width="64" height="64" viewBox="0 0 64 64" fill="none"><circle cx="32" cy="32" r="20" :stroke="colorRango(rango).hex" stroke-width="2"/></svg>
            <svg v-else-if="rango === 'plata'" width="64" height="64" viewBox="0 0 64 64" fill="none"><path d="M32 8L56 32L32 56L8 32L32 8Z" :stroke="colorRango(rango).hex" stroke-width="2"/></svg>
            <svg v-else-if="rango === 'oro'" width="64" height="64" viewBox="0 0 64 64" fill="none"><path d="M32 8L50 18V46L32 56L14 46V18L32 8Z" :stroke="colorRango(rango).hex" stroke-width="2"/></svg>
            <svg v-else width="64" height="64" viewBox="0 0 64 64" fill="none"><path d="M32 8L38 24L54 32L38 40L32 56L26 40L10 32L26 24L32 8Z" :stroke="colorRango(rango).hex" stroke-width="2"/></svg>
          </div>
          <div>
            <h3 class="font-display text-display-md uppercase text-texto">{{ rango }}</h3>
            <p class="font-mono text-[11px] text-texto-muted mt-1">{{ (gamificacion?.puntos_totales || 0).toLocaleString("es-PE") }} puntos</p>
          </div>
        </div>

        <div v-if="rango !== 'platino'" class="mb-6">
          <div class="flex justify-between mb-2 font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em]">
            <span>{{ monto.format(volumen) }}</span><span>{{ Math.round(progreso) }}%</span><span>{{ monto.format(tope) }}</span>
          </div>
          <div class="h-4 border border-borde bg-fondo">
            <div class="h-full transition-all duration-700" :style="{ width: `${progreso}%`, backgroundColor: colorRango(rango).bg }" />
          </div>
          <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mt-2">Siguiente: {{ siguiente[rango] }}</p>
        </div>
        <div v-else class="mb-6 flex items-center gap-2">
          <svg width="14" height="14" viewBox="0 0 14 14" fill="none"><path d="M7 1L8.4 4.6L12 7L8.4 9.4L7 13L5.6 9.4L2 7L5.6 4.6L7 1Z" fill="#7B2FBE"/></svg>
          <p class="font-mono text-[10px] uppercase tracking-[0.2em] text-texto-muted">Rango Maximo</p>
        </div>

        <div class="border border-borde">
          <div v-for="(b, i) in beneficios" :key="b" class="flex items-center gap-2 p-3" :class="i < beneficios.length - 1 ? 'border-b border-borde' : ''">
            <svg width="10" height="10" viewBox="0 0 10 10" fill="none"><rect x="1" y="1" width="8" height="8" stroke="#171717"/><path d="M3 5.2L4.3 6.5L7 3.8" stroke="#171717" stroke-width="1"/></svg>
            <span class="font-mono text-[10px] uppercase tracking-[0.15em] text-texto">{{ b }}</span>
          </div>
        </div>
      </div>

      <div class="p-6 md:p-8">
        <div class="flex items-center justify-between mb-4">
          <h4 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto">Leaderboard</h4>
          <span class="w-2.5 h-2.5 border border-borde" :style="{ backgroundColor: wsConnected ? '#16A34A' : '#737373' }" />
        </div>
        <div class="border border-borde max-h-[500px] overflow-auto">
          <div v-for="item in leaderboard.slice(0, 10)" :key="`${item.posicion}-${item.usuario_empresa}`" class="grid grid-cols-[40px_1fr_auto_auto] gap-2 items-center px-3 py-2 border-b border-borde last:border-b-0">
            <div class="font-mono text-[10px] text-texto-muted">
              <svg v-if="item.posicion === 1" width="12" height="12" viewBox="0 0 12 12" fill="none"><path d="M1.5 9.5H10.5L9.6 4.2L7.6 6L6 3.4L4.4 6L2.4 4.2L1.5 9.5Z" stroke="#171717"/></svg>
              <svg v-else-if="item.posicion === 2" width="12" height="12" viewBox="0 0 12 12" fill="none"><path d="M6 1.5L10.5 6L6 10.5L1.5 6L6 1.5Z" stroke="#171717"/></svg>
              <svg v-else-if="item.posicion === 3" width="12" height="12" viewBox="0 0 12 12" fill="none"><path d="M6 1.5L9.8 3.7V8.3L6 10.5L2.2 8.3V3.7L6 1.5Z" stroke="#171717"/></svg>
              <span v-else>{{ item.posicion }}</span>
            </div>
            <div class="font-mono text-[10px] uppercase truncate">{{ item.usuario_empresa }}</div>
            <span class="w-2 h-2 border border-borde" :style="{ backgroundColor: colorRango((item.rango || 'bronce').toLowerCase()).bg }" />
            <div class="font-mono text-[10px] text-texto-muted">{{ (item.puntos_totales || 0).toLocaleString("es-PE") }}</div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
