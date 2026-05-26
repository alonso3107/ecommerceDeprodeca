<!--
  ═══════════════════════════════════════════════════════════════
  admin/dashboard.vue — Dashboard Admin · DEPRODECA
  Gestión de pedidos: Verificar Pago → Enviar → Entregar.
  Gamificación se dispara al ENTREGAR.
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

async function cargarPedidos() {
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

// Encontrar el pago asociado a un pedido
function pagoDePedido(pedidoId: number): any {
  return pagos.value.find((p: any) => p.pedido_id === pedidoId)
}

// ─── Acciones del admin ──────────────────────────────────

async function verificarPago(pagoId: number) {
  accionLoading.value = pagoId
  try {
    await $fetch(`${config.public.apiBase}/pagos/${pagoId}/verificar`, {
      method: "PUT",
      headers: { Authorization: `Bearer ${getToken()}` },
    })
    mostrarToast("Pago verificado. Pedido confirmado.", "success")
    await cargarPedidos()
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
    await cargarPedidos()
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
    mostrarToast("¡Pedido entregado! Puntos de gamificación acreditados.", "success")
    await cargarPedidos()
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

// ─── Utilidades ───────────────────────────────────────────
function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency", currency: "PEN", minimumFractionDigits: 2,
  }).format(p)
}

function formatearFecha(f: string) {
  return new Date(f).toLocaleDateString("es-PE", {
    day: "numeric", month: "short", hour: "2-digit", minute: "2-digit",
  })
}

const estadoConfig: Record<string, { label: string; color: string }> = {
  pendiente:  { label: "Pendiente",  color: "#F59E0B" },
  confirmado: { label: "Confirmado", color: "#171717" },
  enviado:    { label: "Enviado",    color: "#737373" },
  entregado:  { label: "Entregado",  color: "#16A34A" },
  cancelado:  { label: "Cancelado",  color: "#DC2626" },
}

// ─── Métricas calculadas ──────────────────────────────────
const totales = computed(() => ({
  pendientes: pedidos.value.filter(p => p.estado === "pendiente").length,
  confirmados: pedidos.value.filter(p => p.estado === "confirmado").length,
  enviados: pedidos.value.filter(p => p.estado === "enviado").length,
  entregados: pedidos.value.filter(p => p.estado === "entregado").length,
}))

onMounted(cargarPedidos)
</script>

<template>
  <div class="page-enter">

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Admin
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Dashboard<span class="text-[#D4A017]">.</span>
      </h1>
    </div>

    <!-- ═══ MÉTRICAS · Estados ═══ -->
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-0 border border-borde mb-10">
      <div class="p-5 bg-white border-r border-b sm:border-b-0 border-borde flex flex-col gap-2">
        <p class="font-display text-[clamp(2rem,3vw,2.5rem)] text-[#F59E0B] leading-none">{{ totales.pendientes }}</p>
        <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">Pendientes</p>
      </div>
      <div class="p-5 bg-white border-r border-b sm:border-b-0 border-borde flex flex-col gap-2">
        <p class="font-display text-[clamp(2rem,3vw,2.5rem)] text-texto leading-none">{{ totales.confirmados }}</p>
        <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">Confirmados</p>
      </div>
      <div class="p-5 bg-white border-r border-borde flex flex-col gap-2">
        <p class="font-display text-[clamp(2rem,3vw,2.5rem)] text-[#737373] leading-none">{{ totales.enviados }}</p>
        <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">Enviados</p>
      </div>
      <div class="p-5 bg-white flex flex-col gap-2">
        <p class="font-display text-[clamp(2rem,3vw,2.5rem)] text-[#16A34A] leading-none">{{ totales.entregados }}</p>
        <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">Entregados</p>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando pedidos...</span>
    </div>

    <!-- ═══ TABLA DE PEDIDOS ═══ -->
    <div v-else class="border border-borde">
      <div class="px-6 py-4 bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em]">
        Gestión de Pedidos
      </div>

      <!-- Header -->
      <div class="hidden md:grid grid-cols-[80px_1fr_120px_100px_1fr] gap-4 px-6 py-3
                  bg-fondo border-b border-borde font-mono text-[9px] uppercase tracking-[0.15em] text-texto-muted">
        <span>ID</span>
        <span>Cliente</span>
        <span>Total</span>
        <span>Estado</span>
        <span>Acción</span>
      </div>

      <div v-if="pedidos.length === 0" class="px-6 py-16 text-center">
        <p class="font-display text-display-md text-texto/10 uppercase">Sin pedidos</p>
      </div>

      <!-- Filas -->
      <div
        v-for="(p, i) in pedidos"
        :key="p.id"
        class="grid grid-cols-1 md:grid-cols-[80px_1fr_120px_100px_1fr] gap-3 md:gap-4 px-6 py-4
               items-center border-t border-borde hover:bg-[#D4A017]/[0.02] transition-colors"
        :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
      >
        <!-- ID -->
        <div class="flex md:block items-center gap-2">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">ID</span>
          <span class="font-mono text-[12px] text-texto font-bold tracking-wider">#{{ p.id }}</span>
        </div>

        <!-- Cliente -->
        <div class="flex md:block items-center gap-2">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Cliente</span>
          <span class="font-body text-small text-texto truncate">{{ p.usuario_nombre || p.usuario_email || `Usuario #${p.usuario_id}` }}</span>
        </div>

        <!-- Total -->
        <div class="flex md:block items-center gap-2">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Total</span>
          <span class="font-body text-small font-bold text-texto">{{ formatearPrecio(p.total) }}</span>
        </div>

        <!-- Estado -->
        <div class="flex md:block items-center gap-2">
          <span class="font-mono text-[10px] text-texto-muted md:hidden uppercase tracking-wider">Estado</span>
          <span
            class="inline-flex items-center gap-1.5 px-2.5 py-1 border font-mono text-[9px]
                   uppercase tracking-[0.1em] font-bold"
            :style="{ color: estadoConfig[p.estado]?.color, borderColor: estadoConfig[p.estado]?.color }"
          >
            {{ estadoConfig[p.estado]?.label || p.estado }}
          </span>
        </div>

        <!-- Acciones -->
        <div class="flex flex-wrap gap-1">
          <!-- Pendiente: Verificar Pago -->
          <button
            v-if="p.estado === 'pendiente' && pagoDePedido(p.id)"
            class="font-mono text-[9px] uppercase tracking-[0.1em] px-3 py-1.5 border border-[#F59E0B] text-[#F59E0B]
                   hover:bg-[#F59E0B] hover:text-black transition-colors duration-200 disabled:opacity-40"
            :disabled="accionLoading !== null"
            @click="verificarPago(pagoDePedido(p.id).id)"
          >
            <span v-if="accionLoading === pagoDePedido(p.id)?.id" class="inline-block w-3 h-3 border border-[#F59E0B] border-t-transparent animate-spin mr-1" />
            Verificar
          </button>
          <span
            v-else-if="p.estado === 'pendiente'"
            class="font-mono text-[9px] text-texto-muted/40 uppercase tracking-[0.1em] px-2"
          >
            — Sin pago —
          </span>

          <!-- Confirmado: Enviar -->
          <button
            v-if="p.estado === 'confirmado'"
            class="font-mono text-[9px] uppercase tracking-[0.1em] px-3 py-1.5 border border-texto text-texto
                   hover:bg-texto hover:text-white transition-colors duration-200 disabled:opacity-40"
            :disabled="accionLoading !== null"
            @click="enviarPedido(p.id)"
          >
            <span v-if="accionLoading === p.id" class="inline-block w-3 h-3 border border-texto border-t-transparent animate-spin mr-1" />
            Enviar
          </button>

          <!-- Enviado: Entregar (con gamificación) -->
          <button
            v-if="p.estado === 'enviado'"
            class="font-mono text-[9px] uppercase tracking-[0.1em] px-3 py-1.5 border border-[#16A34A] text-[#16A34A]
                   hover:bg-[#16A34A] hover:text-white transition-colors duration-200 disabled:opacity-40"
            :disabled="accionLoading !== null"
            @click="entregarPedido(p.id)"
          >
            <span v-if="accionLoading === p.id" class="inline-block w-3 h-3 border border-[#16A34A] border-t-transparent animate-spin mr-1" />
            Entregar 🏆
          </button>

          <!-- Entregado / Cancelado: sin acción -->
          <span
            v-if="p.estado === 'entregado' || p.estado === 'cancelado'"
            class="font-mono text-[9px] text-texto-muted/40 uppercase tracking-[0.1em] px-2"
          >
            {{ p.estado === 'entregado' ? '— Completado —' : '— Cancelado —' }}
          </span>
        </div>
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
