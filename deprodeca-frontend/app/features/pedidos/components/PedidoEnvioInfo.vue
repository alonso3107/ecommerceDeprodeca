<script setup lang="ts">
const props = defineProps<{
  envio?: {
    direccion?: string
    fecha_estimada?: string
    metodo?: string
    transportista?: string
    notas?: string
    estado?: string
  } | null
}>()

const tieneEnvio = computed(() => !!props.envio)

const fechaEstimadaFormateada = computed(() => {
  if (!props.envio?.fecha_estimada) return "No disponible"
  return new Date(props.envio.fecha_estimada).toLocaleDateString("es-PE", {
    day: "numeric",
    month: "long",
    year: "numeric",
  })
})

const badgeRuta = computed(() => {
  if (!props.envio?.fecha_estimada) return ""
  const estimada = new Date(props.envio.fecha_estimada).getTime()
  const ahora = Date.now()
  const estado = (props.envio.estado || "").toLowerCase()
  if (estado === "entregado") return ""
  return estimada < ahora ? "En ruta" : "Proximamente"
})
</script>

<template>
  <section class="bg-white border border-borde p-6">
    <div class="flex items-center justify-between mb-4">
      <div class="flex items-center gap-3">
        <svg width="20" height="20" viewBox="0 0 20 20" fill="none">
          <rect x="2" y="6" width="9" height="8" stroke="#171717" />
          <path d="M11 8H15L18 11V14H11V8Z" stroke="#171717" />
          <circle cx="6" cy="15" r="1.3" fill="#171717" />
          <circle cx="14" cy="15" r="1.3" fill="#171717" />
        </svg>
        <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] font-bold text-texto">Datos del Envio</h3>
      </div>
      <span v-if="badgeRuta" class="border border-borde px-2 py-1 font-mono text-[9px] uppercase tracking-[0.14em] text-texto-muted">
        {{ badgeRuta }}
      </span>
    </div>

    <div v-if="!tieneEnvio" class="min-h-[120px] flex items-center justify-center gap-3 border border-borde">
      <svg width="18" height="18" viewBox="0 0 18 18" fill="none"><rect x="2" y="4" width="14" height="11" stroke="#737373"/><path d="M2 7H16" stroke="#737373"/></svg>
      <p class="font-mono text-[11px] uppercase tracking-[0.15em] text-texto-muted">Informacion de envio no disponible</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 border border-borde">
      <div class="p-4 border-b md:border-r border-borde">
        <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-2">Direccion</p>
        <div class="flex items-start gap-2">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none"><path d="M8 14C8 14 12 10.5 12 7.2C12 5 10.2 3.2 8 3.2C5.8 3.2 4 5 4 7.2C4 10.5 8 14 8 14Z" stroke="#171717"/><circle cx="8" cy="7" r="1.5" stroke="#171717"/></svg>
          <p class="font-body text-small text-texto">{{ envio?.direccion || "No disponible" }}</p>
        </div>
      </div>

      <div class="p-4 border-b border-borde">
        <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-2">Fecha estimada</p>
        <div class="flex items-center gap-2">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none"><rect x="2" y="3" width="12" height="11" stroke="#171717"/><path d="M2 6H14" stroke="#171717"/></svg>
          <p class="font-mono text-[12px] text-texto">{{ fechaEstimadaFormateada }}</p>
        </div>
      </div>

      <div class="p-4 md:border-r border-borde">
        <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-2">Metodo</p>
        <div class="flex items-center gap-2">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none"><rect x="2" y="5" width="8" height="6" stroke="#171717"/><path d="M10 6H13L14 8V11H10V6Z" stroke="#171717"/></svg>
          <p class="font-mono text-[11px] uppercase text-texto">{{ envio?.metodo || "No disponible" }}</p>
        </div>
      </div>

      <div class="p-4">
        <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-2">Transportista</p>
        <div class="flex items-center gap-2">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none"><circle cx="8" cy="5" r="2" stroke="#171717"/><path d="M3.5 13C3.5 10.8 5.5 9 8 9C10.5 9 12.5 10.8 12.5 13" stroke="#171717"/></svg>
          <p class="font-body text-small text-texto">{{ envio?.transportista || "No disponible" }}</p>
        </div>
      </div>
    </div>

    <div v-if="envio?.notas" class="mt-4 border border-borde p-4">
      <p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-2">Notas</p>
      <p class="font-body text-small text-texto">{{ envio.notas }}</p>
    </div>
  </section>
</template>
