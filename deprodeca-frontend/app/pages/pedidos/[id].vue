<!--
  ═══════════════════════════════════════════════════════════════
  pedidos/[id].vue — Detalle de Pedido · DEPRODECA
  Brutalismo: tabla de productos con bordes duros, timeline
  de estado con indicadores geométricos, total masivo.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })

const route = useRoute()
const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const pedidoId = computed(() => route.params.id as string)
const pedido = ref<any>(null)
const loading = ref(true)

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any }>(
      `${config.public.apiBase}/pedidos/${pedidoId.value}`,
      { headers: { Authorization: `Bearer ${token}` } },
    )
    if (res.success) pedido.value = res.data
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
    day: "numeric", month: "long", year: "numeric",
    hour: "2-digit", minute: "2-digit",
  })
}

// ─── Timeline de estados ──────────────────────────────────
const estados = ["pendiente", "confirmado", "enviado", "entregado"] as const
const estadoIdx = computed(() => estados.indexOf(pedido.value?.estado))

const estadoConfig: Record<string, { label: string; color: string }> = {
  pendiente:  { label: "Pendiente",  color: "#F59E0B" },
  confirmado: { label: "Confirmado", color: "#171717" },
  enviado:    { label: "Enviado",    color: "#737373" },
  entregado:  { label: "Entregado",  color: "#16A34A" },
  cancelado:  { label: "Cancelado",  color: "#DC2626" },
}
</script>

<template>
  <div class="page-enter max-w-[960px] mx-auto px-6 md:px-8 py-12">

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <!-- 404 -->
    <div v-else-if="!pedido" class="text-center py-32">
      <p class="font-display text-[clamp(3rem,8vw,6rem)] text-texto/10 uppercase leading-none mb-6">
        Pedido No<br />Encontrado
      </p>
      <button
        class="font-mono text-[11px] uppercase tracking-[0.2em] border border-borde px-6 py-3
               text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017] transition-colors"
        @click="navigateTo('/pedidos')"
      >
        Volver a Mis Pedidos
      </button>
    </div>

    <!-- Detalle del pedido -->
    <template v-else>

      <!-- Breadcrumb + Encabezado -->
      <nav class="mb-8 font-mono text-[11px] uppercase tracking-[0.15em] flex items-center gap-2 flex-wrap">
        <NuxtLink to="/pedidos" class="text-texto-muted hover:text-texto transition-colors">
          Mis Pedidos
        </NuxtLink>
        <span class="text-borde">/</span>
        <span class="text-texto font-bold">#{{ pedido.id }}</span>
      </nav>

      <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 mb-10">
        <div>
          <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
            Pedido #{{ pedido.id }}
          </h1>
          <p class="font-mono text-[11px] text-texto-muted mt-2 tracking-wider">
            {{ formatearFecha(pedido.created_at) }}
          </p>
        </div>

        <!-- Badge estado -->
        <span
          class="inline-flex items-center gap-2 px-4 py-2 border font-mono text-[11px]
                 uppercase tracking-[0.12em] font-bold self-start"
          :style="{
            color: estadoConfig[pedido.estado]?.color || '#737373',
            borderColor: estadoConfig[pedido.estado]?.color || '#737373',
          }"
        >
          <!-- Indicador geométrico -->
          <svg v-if="pedido.estado === 'pendiente'" width="8" height="8"><circle cx="4" cy="4" r="3" fill="currentColor"/></svg>
          <svg v-if="pedido.estado === 'confirmado'" width="8" height="8"><rect width="8" height="8" fill="currentColor"/></svg>
          <svg v-if="pedido.estado === 'enviado'" width="10" height="8" viewBox="0 0 10 8"><path d="M0 4H8M8 4L5 1M8 4L5 7" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/></svg>
          <svg v-if="pedido.estado === 'entregado'" width="8" height="8" viewBox="0 0 8 8"><path d="M1.5 4.5L3.5 6.5L7 2.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/></svg>
          <svg v-if="pedido.estado === 'cancelado'" width="8" height="8" viewBox="0 0 8 8"><path d="M1 1L7 7M7 1L1 7" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/></svg>
          {{ estadoConfig[pedido.estado]?.label || pedido.estado }}
        </span>
      </div>

      <!-- Timeline de estados (no cancelado) -->
      <div v-if="pedido.estado !== 'cancelado'" class="mb-10 border border-borde p-6">
        <p class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-5">
          Progreso del Pedido
        </p>
        <div class="flex items-center gap-0">
          <template v-for="(est, i) in estados" :key="est">
            <!-- Nodo -->
            <div class="flex flex-col items-center gap-2 flex-1 relative">
              <!-- Conector izquierdo -->
              <div
                v-if="i > 0"
                class="absolute top-4 right-1/2 w-full h-0.5 -z-0"
                :style="{ backgroundColor: i <= estadoIdx ? estadoConfig[estados[i-1]].color : '#E5E5E5' }"
              />
              <!-- Círculo indicador -->
              <div
                class="w-8 h-8 flex items-center justify-center flex-shrink-0 relative z-10
                       transition-colors duration-500"
                :style="{
                  border: `2px solid ${i <= estadoIdx ? estadoConfig[est].color : '#E5E5E5'}`,
                  backgroundColor: i <= estadoIdx ? estadoConfig[est].color : 'white',
                }"
              >
                <!-- Check si completado -->
                <svg v-if="i < estadoIdx" width="14" height="10" viewBox="0 0 14 10" fill="none">
                  <path d="M1 5L5 9L13 1" stroke="white" stroke-width="2"
                        stroke-linecap="square" stroke-linejoin="miter"/>
                </svg>
                <!-- Punto si actual -->
                <div v-else-if="i === estadoIdx" class="w-2 h-2 bg-white" />
              </div>
              <!-- Label -->
              <span
                class="font-mono text-[9px] uppercase tracking-[0.15em] text-center leading-tight"
                :style="{ color: i <= estadoIdx ? estadoConfig[est].color : '#A3A3A3' }"
              >
                {{ estadoConfig[est].label }}
              </span>
            </div>
          </template>
        </div>
      </div>

      <!-- Productos del pedido -->
      <div class="border border-borde mb-10">
        <div class="px-6 py-4 bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em]">
          Productos
        </div>

        <div
          v-for="(det, i) in pedido.detalles"
          :key="det.id"
          class="flex items-center gap-4 px-6 py-4"
          :class="i > 0 ? 'border-t border-borde' : ''"
        >
          <!-- Imagen pequeña -->
          <div class="w-12 h-12 border border-borde bg-fondo flex-shrink-0">
            <img
              :src="det.imagen_url || 'https://images.unsplash.com/photo-1583258292688-d0213dc154b3?w=60&q=60'"
              :alt="det.producto_nombre"
              class="w-full h-full object-cover"
            />
          </div>

          <!-- Info -->
          <div class="flex-1 min-w-0">
            <p class="font-body text-small font-bold text-texto truncate">
              {{ det.producto_nombre }}
            </p>
            <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider mt-0.5">
              {{ det.cantidad }} × {{ formatearPrecio(det.precio_unitario) }}
            </p>
          </div>

          <!-- Subtotal -->
          <p class="font-display text-heading text-texto flex-shrink-0">
            {{ formatearPrecio(det.subtotal) }}
          </p>
        </div>

        <!-- Total -->
        <div class="border-t-2 border-texto px-6 py-5 flex justify-between items-baseline">
          <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
            Total
          </span>
          <span class="font-display text-[clamp(1.75rem,3vw,2.5rem)] text-texto leading-none">
            {{ formatearPrecio(pedido.total) }}
          </span>
        </div>
      </div>

      <!-- Volver -->
      <NuxtLink
        to="/pedidos"
        class="inline-block font-mono text-[11px] uppercase tracking-[0.2em] border border-borde
               px-6 py-3 text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017]
               transition-colors duration-200"
      >
        ← Volver a Mis Pedidos
      </NuxtLink>

    </template>
  </div>
</template>
