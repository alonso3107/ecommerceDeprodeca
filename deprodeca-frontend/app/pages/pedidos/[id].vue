<script setup lang="ts">
import PedidoTimeline from "~/features/pedidos/components/PedidoTimeline.vue"
import PedidoEnvioInfo from "~/features/pedidos/components/PedidoEnvioInfo.vue"
import PedidoAcciones from "~/features/pedidos/components/PedidoAcciones.vue"

definePageMeta({ layout: "default", middleware: "auth" })

const route = useRoute()
const config = useRuntimeConfig()

const pedidoId = computed(() => route.params.id as string)
const pedido = ref<any>(null)
const loading = ref(true)

const estadoConfig: Record<string, { label: string; color: string }> = {
  pendiente: { label: "Pendiente", color: "#F59E0B" },
  confirmado: { label: "Confirmado", color: "#171717" },
  en_proceso: { label: "En Proceso", color: "#171717" },
  enviado: { label: "Enviado", color: "#737373" },
  entregado: { label: "Entregado", color: "#16A34A" },
  cancelado: { label: "Cancelado", color: "#DC2626" },
}

const PLACEHOLDER = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='400' height='300'%3E%3Crect fill='%23FAFAFA' width='400' height='300'/%3E%3Ctext x='200' y='155' text-anchor='middle' font-size='14' font-family='monospace' fill='%23A3A3A3'%3EDEPRODECA%3C/text%3E%3C/svg%3E"

function authHeaders() {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  return token ? { Authorization: `Bearer ${token}` } : null
}

async function cargarPedido() {
  const headers = authHeaders()
  if (!headers) {
    loading.value = false
    return
  }
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/pedidos/${pedidoId.value}`, { headers })
    if (res.success) pedido.value = res.data
    else pedido.value = null
  } catch {
    pedido.value = null
  } finally {
    loading.value = false
  }
}

onMounted(cargarPedido)

function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(p)
}

function formatearFecha(f: string) {
  return new Date(f).toLocaleDateString("es-PE", {
    day: "numeric",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  })
}

const subtotal = computed(() => (pedido.value?.detalles || []).reduce((sum: number, d: any) => sum + Number(d.subtotal || 0), 0))
const envioMonto = computed(() => Number(pedido.value?.envio_monto || 0))
const descuentoMonto = computed(() => Number(pedido.value?.descuento || 0))
const total = computed(() => Number(pedido.value?.total || 0))

async function cancelarPedido() {
  const headers = authHeaders()
  if (!headers) return
  await $fetch(`${config.public.apiBase}/pedidos/${pedidoId.value}/cancelar`, {
    method: "PUT",
    headers,
  })
  await cargarPedido()
}

function repetirPedido(productos: Array<any>) {
  if (!import.meta.client) return
  const carritoActual = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]") as Array<any>

  for (const p of productos || []) {
    const idProducto = p.producto_id || p.id
    if (!idProducto) continue
    const cantidad = Number(p.cantidad || 1)
    const idx = carritoActual.findIndex((item) => item.id === idProducto)
    if (idx >= 0) {
      carritoActual[idx].cantidad += cantidad
    } else {
      carritoActual.push({
        id: idProducto,
        nombre: p.producto_nombre || p.nombre || "Producto",
        slug: p.producto_slug || p.slug || `producto-${idProducto}`,
        precio: Number(p.precio_unitario || p.precio || 0),
        imagen_url: p.imagen_url || "",
        unidad: p.unidad || "unidad",
        cantidad,
      })
    }
  }

  localStorage.setItem("deprodeca_carrito", JSON.stringify(carritoActual))
  window.dispatchEvent(new Event("carrito-actualizado"))
  navigateTo("/carrito")
}
</script>

<template>
  <main class="max-w-[960px] mx-auto px-6 md:px-8 py-12">
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <div v-else-if="!pedido" class="text-center py-32">
      <p class="font-display text-[clamp(3rem,8vw,6rem)] text-texto/10 uppercase leading-none mb-6">Pedido No<br />Encontrado</p>
      <button class="font-mono text-[11px] uppercase tracking-[0.2em] border border-borde px-6 py-3 text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017] transition-colors" @click="navigateTo('/pedidos')">Volver a Mis Pedidos</button>
    </div>

    <template v-else>
      <nav class="mb-8 font-mono text-[11px] uppercase tracking-[0.15em] flex items-center gap-2 flex-wrap">
        <NuxtLink to="/pedidos" class="text-texto-muted hover:text-texto transition-colors">Mis Pedidos</NuxtLink>
        <span class="text-borde">/</span>
        <span class="text-texto font-bold">#{{ pedido.id }}</span>
      </nav>

      <section class="mb-10 flex flex-col sm:flex-row sm:items-end justify-between gap-4">
        <div>
          <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">Pedido #{{ pedido.id }}</h1>
          <p class="font-mono text-[11px] text-texto-muted mt-2 tracking-wider">{{ formatearFecha(pedido.created_at) }}</p>
        </div>
        <div class="text-right">
          <span class="inline-flex items-center gap-2 px-4 py-2 border font-mono text-[14px] uppercase tracking-[0.12em] font-bold" :style="{ color: estadoConfig[pedido.estado]?.color || '#737373', borderColor: estadoConfig[pedido.estado]?.color || '#737373' }">
            <svg v-if="pedido.estado === 'pendiente'" width="12" height="12"><circle cx="6" cy="6" r="4" fill="currentColor"/></svg>
            <svg v-else-if="pedido.estado === 'confirmado' || pedido.estado === 'en_proceso'" width="12" height="12"><rect x="2" y="2" width="8" height="8" fill="currentColor"/></svg>
            <svg v-else-if="pedido.estado === 'enviado'" width="14" height="12" viewBox="0 0 14 12"><path d="M1 6H12M12 6L8 2M12 6L8 10" stroke="currentColor" stroke-width="1.8"/></svg>
            <svg v-else-if="pedido.estado === 'entregado'" width="12" height="12" viewBox="0 0 12 12"><path d="M2 6.5L5 9.5L10 3.5" stroke="currentColor" stroke-width="1.8"/></svg>
            <svg v-else width="12" height="12" viewBox="0 0 12 12"><path d="M2 2L10 10M10 2L2 10" stroke="currentColor" stroke-width="1.8"/></svg>
            {{ estadoConfig[pedido.estado]?.label || pedido.estado }}
          </span>
          <p class="font-display text-display-md text-texto mt-3">{{ formatearPrecio(total) }}</p>
        </div>
      </section>

      <section class="mb-10">
        <PedidoTimeline
          :estado="pedido.estado"
          :fecha-confirmado="pedido.fecha_confirmado"
          :fecha-enviado="pedido.fecha_enviado"
          :fecha-entregado="pedido.fecha_entregado"
        />
      </section>

      <section class="mb-10">
        <PedidoEnvioInfo :envio="pedido.envio" />
      </section>

      <section class="mb-10 border border-borde">
        <div class="px-6 py-4 bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em]">Productos</div>
        <div v-for="(det, i) in pedido.detalles" :key="det.id || i" class="flex items-center gap-4 px-6 py-4" :class="Number(i) > 0 ? 'border-t border-borde' : ''">
          <div class="w-12 h-12 border border-borde bg-fondo flex-shrink-0">
            <img :src="det.imagen_url || 'https://images.unsplash.com/photo-1583258292688-d0213dc154b3?w=60&q=60'" :alt="det.producto_nombre" class="w-full h-full object-cover" @error="(e: Event) => (e.target as HTMLImageElement).src = PLACEHOLDER" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="font-body text-small font-bold text-texto truncate">{{ det.producto_nombre }}</p>
            <p class="font-mono text-[10px] text-texto-muted uppercase tracking-wider mt-0.5">{{ det.cantidad }} x {{ formatearPrecio(det.precio_unitario) }}</p>
          </div>
          <p class="font-display text-heading text-texto flex-shrink-0">{{ formatearPrecio(det.subtotal) }}</p>
        </div>
        <div class="border-t-2 border-texto px-6 py-5 flex justify-between items-baseline">
          <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">Total</span>
          <span class="font-display text-[clamp(1.75rem,3vw,2.5rem)] text-texto leading-none">{{ formatearPrecio(total) }}</span>
        </div>
      </section>

      <section class="mb-10 border border-borde bg-white p-6">
        <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto-muted mb-4">Resumen Financiero</h3>
        <div class="space-y-2">
          <div class="flex items-center justify-between font-mono text-[11px]"><span>Subtotal</span><span>{{ formatearPrecio(subtotal) }}</span></div>
          <div class="flex items-center justify-between font-mono text-[11px]"><span>Envio</span><span>{{ envioMonto > 0 ? formatearPrecio(envioMonto) : 'Gratis' }}</span></div>
          <div v-if="descuentoMonto > 0" class="flex items-center justify-between font-mono text-[11px]"><span>Descuento</span><span>-{{ formatearPrecio(descuentoMonto) }}</span></div>
        </div>
        <div class="mt-4 pt-4 border-t border-borde flex items-end justify-between">
          <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto-muted">Total Final</span>
          <span class="font-display text-[clamp(2rem,4vw,3.2rem)] leading-none text-texto">{{ formatearPrecio(total) }}</span>
        </div>
      </section>

      <section class="mb-10">
        <PedidoAcciones :pedido-id="pedido.id" :estado="pedido.estado" :productos="pedido.detalles || []" @cancelar="cancelarPedido" @repetir="repetirPedido" />
      </section>
    </template>
  </main>
</template>
