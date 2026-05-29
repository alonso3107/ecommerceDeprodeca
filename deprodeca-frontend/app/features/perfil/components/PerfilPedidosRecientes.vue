<script setup lang="ts">
type Pedido = { id: number | string; created_at: string; total: number; estado: string }
const props = defineProps<{ pedidos: Pedido[] }>()

const money = new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN", maximumFractionDigits: 0 })

function estadoColor(estado: string) {
  const key = (estado || "").toLowerCase()
  const mapa: Record<string, string> = {
    pendiente: "#F59E0B",
    confirmado: "#171717",
    enviado: "#737373",
    entregado: "#16A34A",
    cancelado: "#DC2626",
  }
  return mapa[key] || "#737373"
}
</script>

<template>
  <section>
    <div class="flex items-center justify-between mb-3">
      <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-texto-muted">--- Pedidos Recientes</p>
      <NuxtLink to="/pedidos" class="font-mono text-[10px] uppercase tracking-[0.2em] text-texto-muted hover:text-texto transition-colors duration-500">Ver todos -></NuxtLink>
    </div>

    <div class="border border-borde bg-white overflow-hidden">
      <template v-if="pedidos.length">
        <div class="grid grid-cols-[1fr_1fr_1fr_1fr] bg-texto text-white font-mono text-[10px] uppercase tracking-[0.15em]">
          <div class="px-4 py-3">Pedido</div><div class="px-4 py-3">Fecha</div><div class="px-4 py-3">Total</div><div class="px-4 py-3">Estado</div>
        </div>
        <NuxtLink v-for="pedido in pedidos.slice(0, 5)" :key="pedido.id" :to="`/pedidos/${pedido.id}`" class="grid grid-cols-[1fr_1fr_1fr_1fr] px-0 border-t border-borde hover:bg-fondo transition-colors duration-500">
          <div class="px-4 py-3 font-mono text-[15px] font-bold text-texto">#{{ pedido.id }}</div>
          <div class="px-4 py-3 font-mono text-[12px] text-texto-muted">{{ new Date(pedido.created_at).toLocaleDateString("es-PE") }}</div>
          <div class="px-4 py-3 font-mono text-[13px] font-bold text-texto">{{ money.format(pedido.total || 0) }}</div>
          <div class="px-4 py-3">
            <span class="inline-flex items-center gap-2 border px-2 py-1 font-mono text-[10px] uppercase tracking-[0.1em]" :style="{ borderColor: estadoColor(pedido.estado), color: estadoColor(pedido.estado) }">
              <svg width="8" height="8" viewBox="0 0 8 8" fill="none">
                <circle v-if="(pedido.estado || '').toLowerCase() === 'pendiente'" cx="4" cy="4" r="2.5" :stroke="estadoColor(pedido.estado)" />
                <rect v-else-if="(pedido.estado || '').toLowerCase() === 'confirmado'" x="1.5" y="1.5" width="5" height="5" :stroke="estadoColor(pedido.estado)" />
                <path v-else-if="(pedido.estado || '').toLowerCase() === 'enviado'" d="M1 4H7M4 1L7 4L4 7" :stroke="estadoColor(pedido.estado)" />
                <path v-else-if="(pedido.estado || '').toLowerCase() === 'entregado'" d="M1.5 4L3.3 5.8L6.5 2.6" :stroke="estadoColor(pedido.estado)" />
                <path v-else d="M2 2L6 6M6 2L2 6" :stroke="estadoColor(pedido.estado)" />
              </svg>
              {{ pedido.estado }}
            </span>
          </div>
        </NuxtLink>
      </template>

      <div v-else class="min-h-[180px] flex items-center justify-center gap-3">
        <svg width="18" height="18" viewBox="0 0 18 18" fill="none"><rect x="2" y="4" width="14" height="11" stroke="#737373"/><path d="M2 7H16" stroke="#737373"/></svg>
        <p class="font-mono text-[12px] uppercase tracking-[0.15em] text-texto-muted">Aun no tienes pedidos</p>
      </div>
    </div>
  </section>
</template>
