<!--
  ═══════════════════════════════════════════════════════════════
  admin/dashboard.vue — Panel de Control · DEPRODECA
  Brutalismo funcional: métricas masivas, tabla de gestión
  con acciones claras. Flujo: Verificar → Enviar → Entregar.
  Gamificación se dispara al entregar.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "admim", middleware: "admin" })

const config = useRuntimeConfig()
const pedidos = ref<any[]>([])
const pagos = ref<any[]>([])
const loading = ref(true)
const accionLoading = ref<number | null>(null)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

function getToken() {
  return import.meta.client ? localStorage.getItem("deprodeca_token") : null
}

async function cargarDatos() {
  loading.value = true
  const token = getToken()
  if (!token) return
  try {
    const [resPedidos, resPagos] = await Promise.all([
      $fetch<{ success: boolean; data: any[] }>(
        `${config.public.apiBase}/admin/pedidos?limite=50`,
        { headers: { Authorization: `Bearer ${token}` } },
      ),
      $fetch<{ success: boolean; data: any[] }>(
        `${config.public.apiBase}/admin/pagos?limite=100`,
        { headers: { Authorization: `Bearer ${token}` } },
      ),
    ])
    if (resPedidos.success) pedidos.value = resPedidos.data
    if (resPagos.success) pagos.value = resPagos.data
  } catch (_) { /* fallback */ }
  finally { loading.value = false }
}

function pagoDePedido(pedidoId: number): any {
  return pagos.value.find((p: any) => p.pedido_id === pedidoId)
}

// ─── Acciones ──────────────────────────────────────────
async function verificarPago(pagoId: number) {
  accionLoading.value = pagoId
  try {
    await $fetch(`${config.public.apiBase}/pagos/${pagoId}/verificar`, {
      method: "PUT",
      headers: { Authorization: `Bearer ${getToken()}` },
    })
    mostrarToast("Pago verificado. Pedido confirmado.", "success")
    await cargarDatos()
  } catch (e: any) {
    mostrarToast(e?.data?.message || "Error al verificar pago", "error")
  } finally { accionLoading.value = null }
}

async function enviarPedido(pedidoId: number) {
  accionLoading.value = pedidoId
  try {
    await $fetch(`${config.public.apiBase}/admin/pedidos/${pedidoId}/enviar`, {
      method: "PUT",
      headers: { Authorization: `Bearer ${getToken()}` },
    })
    mostrarToast("Pedido marcado como enviado", "success")
    await cargarDatos()
  } catch (e: any) {
    mostrarToast(e?.data?.message || "Error al enviar pedido", "error")
  } finally { accionLoading.value = null }
}

async function entregarPedido(pedidoId: number) {
  accionLoading.value = pedidoId
  try {
    await $fetch(`${config.public.apiBase}/admin/pedidos/${pedidoId}/entregar`, {
      method: "PUT",
      headers: { Authorization: `Bearer ${getToken()}` },
    })
    mostrarToast("Pedido entregado. Puntos de gamificacion acreditados.", "success")
    await cargarDatos()
  } catch (e: any) {
    mostrarToast(e?.data?.message || "Error al entregar pedido", "error")
  } finally { accionLoading.value = null }
}

function mostrarToast(msg: string, type: "success" | "error") {
  toastMsg.value = msg
  toastType.value = type
  showToast.value = true
  setTimeout(() => (showToast.value = false), 4000)
}

// ─── Utilidades ────────────────────────────────────────
function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency", currency: "PEN", minimumFractionDigits: 2,
  }).format(p)
}

const estadoConfig: Record<string, { label: string; color: string; bg: string }> = {
  pendiente:  { label: "Pendiente",  color: "#D97706", bg: "#FEF3C7" },
  confirmado: { label: "Confirmado", color: "#171717", bg: "#E5E5E5" },
  enviado:    { label: "Enviado",    color: "#525252", bg: "#F5F5F5" },
  entregado:  { label: "Entregado",  color: "#15803D", bg: "#DCFCE7" },
  cancelado:  { label: "Cancelado",  color: "#DC2626", bg: "#FEE2E2" },
}

// ─── Métricas ──────────────────────────────────────────
const totales = computed(() => ({
  pendientes: pedidos.value.filter(p => p.estado === "pendiente").length,
  confirmados: pedidos.value.filter(p => p.estado === "confirmado").length,
  enviados: pedidos.value.filter(p => p.estado === "enviado").length,
  entregados: pedidos.value.filter(p => p.estado === "entregado").length,
}))

const metricas = computed(() => [
  { label: "Pendientes", valor: totales.value.pendientes, color: "#D97706", icon: "pending" },
  { label: "Confirmados", valor: totales.value.confirmados, color: "#171717", icon: "confirmed" },
  { label: "Enviados", valor: totales.value.enviados, color: "#525252", icon: "shipped" },
  { label: "Entregados", valor: totales.value.entregados, color: "#15803D", icon: "delivered" },
])

onMounted(cargarDatos)
</script>

<template>
  <div class="page-enter">

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-xs text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Panel de Control
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Dashboard<span class="text-[#D4A017]">.</span>
      </h1>
      <p class="mt-2 font-body text-sm text-texto-muted">
        Gestion de pedidos, verificacion de pagos y seguimiento de envios
      </p>
    </div>

    <!-- ═══ METRICAS ═══ -->
    <div class="grid grid-cols-2 lg:grid-cols-4 gap-0 border border-borde mb-10">
      <div
        v-for="(m, i) in metricas"
        :key="m.label"
        class="bg-white p-6 flex flex-col gap-4 border-b lg:border-b-0 border-borde"
        :class="i < 3 ? 'border-r' : ''"
      >
        <!-- Indicador geometrico -->
        <div class="flex items-center gap-2">
          <svg v-if="m.icon === 'pending'" width="14" height="14" viewBox="0 0 14 14">
            <circle cx="7" cy="7" r="5" :stroke="m.color" stroke-width="2" fill="none"/>
            <circle cx="7" cy="7" r="2" :fill="m.color"/>
          </svg>
          <svg v-if="m.icon === 'confirmed'" width="14" height="14" viewBox="0 0 14 14">
            <rect width="14" height="14" :stroke="m.color" stroke-width="2" fill="none"/>
            <rect x="3" y="3" width="8" height="8" :fill="m.color"/>
          </svg>
          <svg v-if="m.icon === 'shipped'" width="16" height="12" viewBox="0 0 16 12">
            <path d="M1 6H13M13 6L9 1.5M13 6L9 10.5" :stroke="m.color" stroke-width="2" stroke-linecap="square" stroke-linejoin="miter"/>
          </svg>
          <svg v-if="m.icon === 'delivered'" width="14" height="14" viewBox="0 0 14 14">
            <path d="M2 7.5L5.5 11L12 4" :stroke="m.color" stroke-width="2.5" stroke-linecap="square" stroke-linejoin="miter"/>
          </svg>
          <span class="font-mono text-xs uppercase tracking-[0.15em] font-bold" :style="{ color: m.color }">
            {{ m.label }}
          </span>
        </div>

        <!-- Numero masivo -->
        <p class="font-display text-[clamp(2.5rem,4vw,3.5rem)] leading-none" :style="{ color: m.color }">
          {{ m.valor }}
        </p>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-24 gap-4">
      <div class="w-10 h-10 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-sm text-texto-muted uppercase tracking-[0.2em]">Cargando pedidos...</span>
    </div>

    <!-- Vacio -->
    <div v-else-if="pedidos.length === 0" class="text-center py-24 border border-borde">
      <svg class="mx-auto mb-6 text-texto/10" width="64" height="64" viewBox="0 0 64 64" fill="none">
        <rect x="10" y="8" width="44" height="52" stroke="currentColor" stroke-width="2"/>
        <line x1="22" y1="20" x2="42" y2="20" stroke="currentColor" stroke-width="1.5"/>
        <line x1="22" y1="28" x2="42" y2="28" stroke="currentColor" stroke-width="1.5"/>
        <line x1="22" y1="36" x2="34" y2="36" stroke="currentColor" stroke-width="1.5"/>
      </svg>
      <p class="font-display text-display-md text-texto/10 uppercase">Sin pedidos</p>
      <p class="font-body text-sm text-texto-muted mt-3">No hay pedidos registrados en el sistema.</p>
    </div>

    <!-- ═══ TABLA DE PEDIDOS ═══ -->
    <div v-else class="border border-borde">
      <!-- Header -->
      <div class="px-6 py-5 bg-texto text-white flex items-center justify-between">
        <div class="flex items-center gap-3">
          <svg width="18" height="18" viewBox="0 0 18 18" fill="none">
            <rect x="3" y="2" width="12" height="15" stroke="white" stroke-width="1.5"/>
            <line x1="6" y1="6" x2="12" y2="6" stroke="white" stroke-width="1.2"/>
            <line x1="6" y1="9" x2="12" y2="9" stroke="white" stroke-width="1.2"/>
            <line x1="6" y1="12" x2="9" y2="12" stroke="white" stroke-width="1.2"/>
          </svg>
          <span class="font-mono text-sm uppercase tracking-[0.2em] font-bold">
            Gestion de Pedidos
          </span>
        </div>
        <span class="font-mono text-xs text-white/40 uppercase tracking-[0.15em]">
          {{ pedidos.length }} pedidos
        </span>
      </div>

      <!-- Columnas (desktop) -->
      <div class="hidden md:grid grid-cols-[80px_1fr_130px_130px_1fr] gap-4 px-6 py-4
                  bg-fondo border-b border-borde font-mono text-xs uppercase tracking-[0.15em] text-texto-muted">
        <span>ID</span>
        <span>Cliente</span>
        <span>Total</span>
        <span>Estado</span>
        <span>Accion</span>
      </div>

      <!-- Filas -->
      <div
        v-for="(p, i) in pedidos"
        :key="p.id"
        class="grid grid-cols-1 md:grid-cols-[80px_1fr_130px_130px_1fr] gap-3 md:gap-4 px-6 py-5
               items-center border-t border-borde transition-colors duration-200"
        :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
      >
        <!-- ID -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">ID</span>
          <p class="font-mono text-sm text-texto font-bold tracking-wider">
            #{{ p.id }}
          </p>
        </div>

        <!-- Cliente -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Cliente</span>
          <p class="font-body text-sm text-texto font-medium truncate">
            {{ p.usuario_nombre || p.usuario_email || 'Usuario #' + p.usuario_id }}
          </p>
        </div>

        <!-- Total -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Total</span>
          <p class="font-body text-sm font-bold text-texto">
            {{ formatearPrecio(p.total) }}
          </p>
        </div>

        <!-- Estado -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Estado</span>
          <span
            class="inline-flex items-center gap-1.5 px-3 py-1.5 font-mono text-xs uppercase tracking-[0.1em] font-bold"
            :style="{
              color: estadoConfig[p.estado]?.color,
              backgroundColor: estadoConfig[p.estado]?.bg,
              border: '1px solid ' + estadoConfig[p.estado]?.color,
            }"
          >
            {{ estadoConfig[p.estado]?.label || p.estado }}
          </span>
        </div>

        <!-- Acciones -->
        <div class="flex flex-wrap gap-1.5">
          <!-- Pendiente con pago: Verificar -->
          <button
            v-if="p.estado === 'pendiente' && pagoDePedido(p.id)"
            class="font-mono text-xs uppercase tracking-[0.1em] px-4 py-2 border cursor-pointer
                   transition-all duration-200 disabled:opacity-40 font-bold"
            :style="{
              color: '#D97706',
              borderColor: '#D97706',
              backgroundColor: 'transparent',
            }"
            :disabled="accionLoading !== null"
            @click="verificarPago(pagoDePedido(p.id).id)"
            @mouseenter="(e: MouseEvent) => {
              (e.target as HTMLElement).style.backgroundColor = '#D97706';
              (e.target as HTMLElement).style.color = '#000';
            }"
            @mouseleave="(e: MouseEvent) => {
              (e.target as HTMLElement).style.backgroundColor = 'transparent';
              (e.target as HTMLElement).style.color = '#D97706';
            }"
          >
            <span v-if="accionLoading === pagoDePedido(p.id)?.id" class="inline-block w-3.5 h-3.5 border-2 border-[#D97706] border-t-transparent animate-spin mr-1.5" />
            Verificar Pago
          </button>
          <span
            v-else-if="p.estado === 'pendiente'"
            class="font-mono text-xs text-texto-muted/50 uppercase tracking-[0.1em] px-3 py-2 italic"
          >
            Sin comprobante
          </span>

          <!-- Confirmado: Enviar -->
          <button
            v-if="p.estado === 'confirmado'"
            class="font-mono text-xs uppercase tracking-[0.1em] px-4 py-2 border border-texto
                   text-texto hover:bg-texto hover:text-white cursor-pointer
                   transition-all duration-200 disabled:opacity-40 font-bold"
            :disabled="accionLoading !== null"
            @click="enviarPedido(p.id)"
          >
            <span v-if="accionLoading === p.id" class="inline-block w-3.5 h-3.5 border-2 border-texto border-t-transparent animate-spin mr-1.5" />
            Marcar Enviado
          </button>

          <!-- Enviado: Entregar -->
          <button
            v-if="p.estado === 'enviado'"
            class="font-mono text-xs uppercase tracking-[0.1em] px-4 py-2 border cursor-pointer
                   transition-all duration-200 disabled:opacity-40 font-bold"
            :style="{
              color: '#15803D',
              borderColor: '#15803D',
              backgroundColor: 'transparent',
            }"
            :disabled="accionLoading !== null"
            @click="entregarPedido(p.id)"
            @mouseenter="(e: MouseEvent) => {
              (e.target as HTMLElement).style.backgroundColor = '#15803D';
              (e.target as HTMLElement).style.color = '#fff';
            }"
            @mouseleave="(e: MouseEvent) => {
              (e.target as HTMLElement).style.backgroundColor = 'transparent';
              (e.target as HTMLElement).style.color = '#15803D';
            }"
          >
            <span v-if="accionLoading === p.id" class="inline-block w-3.5 h-3.5 border-2 border-[#15803D] border-t-transparent animate-spin mr-1.5" />
            Confirmar Entrega
          </button>

          <!-- Entregado / Cancelado -->
          <span
            v-if="p.estado === 'entregado' || p.estado === 'cancelado'"
            class="font-mono text-xs uppercase tracking-[0.1em] px-3 py-2 font-bold"
            :style="{ color: estadoConfig[p.estado]?.color }"
          >
            {{ p.estado === 'entregado' ? 'Completado' : 'Cancelado' }}
          </span>
        </div>
      </div>

      <!-- Footer con totales -->
      <div class="border-t-2 border-texto px-6 py-4 bg-fondo flex flex-wrap items-center justify-between gap-4">
        <div class="flex flex-wrap gap-6">
          <div v-for="m in metricas" :key="m.label" class="flex items-center gap-2">
            <span class="font-mono text-xs uppercase tracking-[0.1em]" :style="{ color: m.color }">
              {{ m.label }}
            </span>
            <span class="font-display text-lg" :style="{ color: m.color }">
              {{ m.valor }}
            </span>
          </div>
        </div>
        <span class="font-mono text-xs text-texto-muted uppercase tracking-[0.15em]">
          Total: {{ pedidos.length }} pedidos
        </span>
      </div>
    </div>

    <!-- Toast -->
    <Toast
      :message="toastMsg"
      :type="toastType"
      :show="showToast"
      @close="showToast = false"
    />
  </div>
</template>
