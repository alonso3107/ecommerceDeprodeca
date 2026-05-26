<!--
  ═══════════════════════════════════════════════════════════════
  pedidos/index.vue — Mis Pedidos · DEPRODECA
  Brutalismo: tabla de pedidos con bordes duros, badges de estado
  como bloques rectangulares con indicadores geométricos.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })

const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const pedidos = ref<any[]>([])
const loading = ref(true)

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any[] }>(
      `${config.public.apiBase}/pedidos`,
      { headers: { Authorization: `Bearer ${token}` } },
    )
    if (res.success) pedidos.value = res.data
  } catch (_) {
    /* fallback */
  } finally {
    loading.value = false
  }
})

// ─── Utilidades ───────────────────────────────────────────
function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency", currency: "PEN", minimumFractionDigits: 2,
  }).format(p)
}

function formatearFecha(f: string) {
  return new Date(f).toLocaleDateString("es-PE", {
    day: "numeric", month: "short", year: "numeric",
  })
}

// ─── Configuración de estado (íconos geométricos) ─────────
const estadoConfig: Record<string, { label: string; color: string; icon: string }> = {
  pendiente:  { label: "Pendiente",  color: "#F59E0B", icon: "circle" },
  confirmado: { label: "Confirmado", color: "#171717", icon: "square" },
  enviado:    { label: "Enviado",    color: "#737373", icon: "arrow" },
  entregado:  { label: "Entregado",  color: "#16A34A", icon: "check" },
  cancelado:  { label: "Cancelado",  color: "#DC2626", icon: "x" },
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Pedidos
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Mis Pedidos<span class="text-[#D4A017]">.</span>
      </h1>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <!-- Vacío -->
    <div v-else-if="pedidos.length === 0" class="text-center py-32">
      <!-- Ícono: Caja/documento geométrico -->
      <svg class="mx-auto mb-8 text-texto/10" width="72" height="72" viewBox="0 0 72 72" fill="none">
        <rect x="12" y="10" width="48" height="56" stroke="currentColor" stroke-width="2"/>
        <line x1="24" y1="22" x2="48" y2="22" stroke="currentColor" stroke-width="1.5"/>
        <line x1="24" y1="30" x2="48" y2="30" stroke="currentColor" stroke-width="1.5"/>
        <line x1="24" y1="38" x2="40" y2="38" stroke="currentColor" stroke-width="1.5"/>
      </svg>
      <p class="font-display text-[clamp(3rem,8vw,6rem)] text-texto/10 uppercase leading-none mb-6">
        Sin<br />Pedidos
      </p>
      <p class="font-body text-body text-texto-muted mb-8">
        Aún no has realizado ningún pedido.
      </p>
      <button
        class="font-mono text-[11px] uppercase tracking-[0.2em] border border-borde px-6 py-3
               text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017]
               transition-colors duration-200"
        @click="navigateTo('/catalogo')"
      >
        Ir al Catálogo →
      </button>
    </div>

    <!-- Tabla de pedidos -->
    <div v-else class="border border-borde">

      <!-- Header row -->
      <div class="hidden md:grid grid-cols-[100px_1fr_140px_160px] gap-4 px-6 py-4
                  bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em]">
        <span>Pedido</span>
        <span>Fecha</span>
        <span>Total</span>
        <span>Estado</span>
      </div>

      <!-- Filas de pedidos -->
      <NuxtLink
        v-for="(p, i) in pedidos"
        :key="p.id"
        :to="`/pedidos/${p.id}`"
        class="grid grid-cols-1 md:grid-cols-[100px_1fr_140px_160px] gap-4 px-6 py-5
               items-center border-t border-borde
               hover:bg-[#D4A017]/[0.02] transition-colors duration-200"
        :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
      >
        <!-- ID -->
        <div>
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Pedido</span>
          <p class="font-mono text-[13px] text-texto font-bold tracking-wider">
            #{{ p.id }}
          </p>
        </div>

        <!-- Fecha -->
        <div>
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Fecha</span>
          <p class="font-body text-small text-texto-muted">{{ formatearFecha(p.created_at) }}</p>
        </div>

        <!-- Total -->
        <div class="md:text-right">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Total</span>
          <p class="font-display text-heading text-texto leading-none">
            {{ formatearPrecio(p.total) }}
          </p>
        </div>

        <!-- Estado · Badge geométrico -->
        <div class="md:text-right">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Estado</span>
          <span
            class="inline-flex items-center gap-2 px-3 py-1.5 border font-mono text-[10px]
                   uppercase tracking-[0.12em] font-bold"
            :style="{
              color: estadoConfig[p.estado]?.color || '#737373',
              borderColor: estadoConfig[p.estado]?.color || '#737373',
            }"
          >
            <!-- Indicador geométrico por estado -->
            <!-- Pendiente: círculo -->
            <svg v-if="estadoConfig[p.estado]?.icon === 'circle'" width="8" height="8" viewBox="0 0 8 8">
              <circle cx="4" cy="4" r="3" fill="currentColor"/>
            </svg>
            <!-- Confirmado: cuadrado -->
            <svg v-if="estadoConfig[p.estado]?.icon === 'square'" width="8" height="8" viewBox="0 0 8 8">
              <rect width="8" height="8" fill="currentColor"/>
            </svg>
            <!-- Enviado: flecha -->
            <svg v-if="estadoConfig[p.estado]?.icon === 'arrow'" width="10" height="8" viewBox="0 0 10 8">
              <path d="M0 4H8M8 4L5 1M8 4L5 7" stroke="currentColor" stroke-width="1.5"
                    stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <!-- Entregado: check -->
            <svg v-if="estadoConfig[p.estado]?.icon === 'check'" width="8" height="8" viewBox="0 0 8 8">
              <path d="M1.5 4.5L3.5 6.5L7 2.5" stroke="currentColor" stroke-width="1.5"
                    stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
            <!-- Cancelado: X -->
            <svg v-if="estadoConfig[p.estado]?.icon === 'x'" width="8" height="8" viewBox="0 0 8 8">
              <path d="M1 1L7 7M7 1L1 7" stroke="currentColor" stroke-width="1.5"
                    stroke-linecap="square"/>
            </svg>
            {{ estadoConfig[p.estado]?.label || p.estado }}
          </span>
        </div>
      </NuxtLink>
    </div>

  </div>
</template>
