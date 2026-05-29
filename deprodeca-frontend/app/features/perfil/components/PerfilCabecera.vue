<script setup lang="ts">
const props = defineProps<{
  perfil: {
    nombre?: string
    empresa?: string
    ruc?: string
    rango?: string
  }
  gamificacion?: {
    rango?: string
    puntos_totales?: number
  } | null
}>()

const rangoActual = computed(() => (props.gamificacion?.rango || props.perfil?.rango || "bronce").toLowerCase())
const puntos = computed(() => props.gamificacion?.puntos_totales || 0)

const iniciales = computed(() => {
  const nombre = (props.perfil?.nombre || "").trim()
  if (!nombre) return "--"
  const parts = nombre.split(/\s+/).slice(0, 2)
  return parts.map((p) => p[0]?.toUpperCase() || "").join("")
})

function colorRango(rango: string) {
  const mapa: Record<string, string> = {
    bronce: "#CD7F32",
    plata: "#A8A8A8",
    oro: "#D4A017",
    platino: "#7B2FBE",
  }
  return mapa[rango] || "#CD7F32"
}
</script>

<template>
  <section class="bg-white border border-borde border-b border-borde p-8 pb-8 mb-0">
    <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
      <div class="flex items-center gap-5">
        <div class="w-[72px] h-[72px] border flex items-center justify-center bg-white border-[#D4A017]">
          <span class="font-display text-2xl text-[#D4A017] uppercase">{{ iniciales }}</span>
        </div>
        <div>
          <h1 class="font-display text-display-lg uppercase text-texto leading-none">{{ perfil.nombre }}</h1>
          <p class="font-mono text-[11px] text-texto-muted mt-2 uppercase tracking-[0.12em]">{{ perfil.empresa }}</p>
          <p class="font-mono text-[9px] text-texto-muted mt-1 uppercase tracking-[0.12em]">RUC {{ perfil.ruc }}</p>
        </div>
      </div>

      <div class="border border-borde px-4 py-3 min-w-[200px] bg-white">
        <div class="flex items-center gap-3">
          <svg v-if="rangoActual === 'bronce'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <circle cx="9" cy="9" r="6" :stroke="colorRango(rangoActual)" stroke-width="1.5" />
          </svg>
          <svg v-else-if="rangoActual === 'plata'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <path d="M9 2L16 9L9 16L2 9L9 2Z" :stroke="colorRango(rangoActual)" stroke-width="1.5" />
          </svg>
          <svg v-else-if="rangoActual === 'oro'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <path d="M9 2L15 5.5V12.5L9 16L3 12.5V5.5L9 2Z" :stroke="colorRango(rangoActual)" stroke-width="1.5" />
          </svg>
          <svg v-else width="18" height="18" viewBox="0 0 18 18" fill="none">
            <path d="M9 1.5L11 6.8L16.5 9L11 11.2L9 16.5L7 11.2L1.5 9L7 6.8L9 1.5Z" :stroke="colorRango(rangoActual)" stroke-width="1.5" />
          </svg>
          <span class="font-mono text-[10px] uppercase tracking-[0.2em] text-texto">{{ rangoActual }}</span>
        </div>
        <p class="font-mono text-[10px] text-texto-muted mt-2">{{ puntos.toLocaleString("es-PE") }} puntos</p>
      </div>
    </div>
  </section>
</template>
