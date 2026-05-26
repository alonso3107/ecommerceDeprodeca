<!--
  ═══════════════════════════════════════════════════════════════
  MetodoPagoSelector.vue — Selector de método de pago brutalista
  Bloques rectangulares duros, sin curvas, sin radio buttons
  circulares. Íconos geométricos únicos para cada método.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// ─── Tipos ────────────────────────────────────────────────
interface MetodoPago {
  id: string
  nombre: string
  datos: string
}

// ─── Props & Emits ────────────────────────────────────────
defineProps<{
  metodos: MetodoPago[]
  metodoActivo: string
}>()
// ─── Emits ────────────────────────────────────────────────
const emit = defineEmits<{
  (e: "seleccionar", id: string): void
}>()

// ─── Helpers ──────────────────────────────────────────────
function claseBoton(id: string, activo: string) {
  return activo === id
    ? "bg-texto text-white border-texto"
    : "bg-white text-texto-muted hover:text-texto hover:border-texto/30 sm:border-r-0 last:border-r"
}
</script>

<template>
  <div>
    <!-- Encabezado técnico -->
    <div class="flex items-center gap-3 mb-6">
      <!-- Ícono: Tarjeta/banco geométrico -->
      <svg width="20" height="16" viewBox="0 0 20 16" fill="none" class="text-[#D4A017] flex-shrink-0">
        <rect x="1" y="2" width="18" height="12" stroke="currentColor" stroke-width="1.5"/>
        <line x1="1" y1="7" x2="19" y2="7" stroke="currentColor" stroke-width="1.5"/>
        <rect x="4" y="10" width="4" height="2" fill="currentColor"/>
        <rect x="10" y="10" width="2" height="2" fill="currentColor"/>
      </svg>
      <h2 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
        Método de Pago
      </h2>
    </div>

    <!-- Opciones · Bloques duros toggle -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-0">
      <button
        v-for="m in metodos"
        :key="m.id"
        class="group flex flex-col items-center gap-3 p-6 border border-borde
               transition-all duration-200 text-center"
        :class="claseBoton(m.id, metodoActivo)"
        @click="emit('seleccionar', m.id)"
      >

        <!-- ÍCONOS GEOMÉTRICOS ÚNICOS -->

        <!-- Yape: Teléfono geométrico con ondas -->
        <svg v-if="m.id === 'yape'" width="28" height="28" viewBox="0 0 28 28" fill="none"
             :class="metodoActivo === m.id ? 'text-white' : 'text-texto'">
          <!-- Cuerpo del teléfono -->
          <rect x="7" y="2" width="14" height="24" rx="2" stroke="currentColor" stroke-width="1.5"/>
          <!-- Pantalla -->
          <rect x="10" y="5" width="8" height="12" stroke="currentColor" stroke-width="1"/>
          <!-- Ondas de sonido (abajo) -->
          <path d="M12 20h4M11 22h6M10 24h8" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
          <!-- Botón home -->
          <rect x="13" y="22" width="2" height="2" fill="currentColor"/>
        </svg>

        <!-- Plin: Escudo/diamante geométrico -->
        <svg v-if="m.id === 'plin'" width="28" height="28" viewBox="0 0 28 28" fill="none"
             :class="metodoActivo === m.id ? 'text-white' : 'text-texto'">
          <!-- Diamante exterior -->
          <path d="M14 2L26 14L14 26L2 14L14 2Z" stroke="currentColor" stroke-width="1.5"/>
          <!-- Check interior -->
          <path d="M9 14L12.5 18L19 11" stroke="currentColor" stroke-width="2" stroke-linecap="square" stroke-linejoin="miter"/>
        </svg>

        <!-- Transferencia: Edificio/banco geométrico -->
        <svg v-if="m.id === 'transferencia'" width="28" height="28" viewBox="0 0 28 28" fill="none"
             :class="metodoActivo === m.id ? 'text-white' : 'text-texto'">
          <!-- Edificio -->
          <rect x="8" y="8" width="12" height="18" stroke="currentColor" stroke-width="1.5"/>
          <!-- Techo -->
          <path d="M6 8L14 3L22 8" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
          <!-- Puerta -->
          <rect x="12" y="18" width="4" height="8" fill="currentColor"/>
          <!-- Ventanas -->
          <rect x="10" y="10" width="2" height="2" fill="currentColor"/>
          <rect x="16" y="10" width="2" height="2" fill="currentColor"/>
          <rect x="10" y="14" width="2" height="2" fill="currentColor"/>
          <rect x="16" y="14" width="2" height="2" fill="currentColor"/>
        </svg>

        <!-- Nombre -->
        <span class="font-mono text-[11px] uppercase tracking-[0.2em] font-bold">
          {{ m.nombre }}
        </span>

        <!-- Indicador de selección: cuadrado, no círculo -->
        <span
          class="w-3 h-3 border"
          :class="metodoActivo === m.id
            ? 'bg-white border-white'
            : 'border-borde group-hover:border-texto/30'"
        />
      </button>
    </div>

    <!-- Datos de pago · Caja monospace técnica -->
    <div
      v-if="metodoActivo"
      class="mt-0 border border-t-0 border-borde p-6 bg-fondo/50"
    >
      <p class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-3">
        Datos para el Pago
      </p>
      <pre class="font-mono text-[13px] text-texto whitespace-pre-line leading-relaxed">{{
        metodos.find(m => m.id === metodoActivo)?.datos
      }}</pre>
      <p class="mt-4 font-body text-caption text-texto-muted">
        Realizá el pago por
        <span class="font-bold text-texto">
          {{ metodos.find(m => m.id === metodoActivo)?.nombre }}
        </span>
        y luego confirmá tu pedido.
      </p>
    </div>
  </div>
</template>
