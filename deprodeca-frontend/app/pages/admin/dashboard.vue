<!--
  ═══════════════════════════════════════════════════════════════
  admin/dashboard.vue — Dashboard Admin · DEPRODECA
  Brutalismo de datos: métricas como bloques con indicadores
  geométricos y tendencias. Sin adornos, solo números.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "admim", middleware: "admin" })

// ─── Métricas (datos mock para el diseño) ────────────────
const metricas = [
  {
    label: "Pedidos Hoy",
    valor: "12",
    cambio: "+3",
    tendencia: "up",
    icon: "orders",
    color: "#D4A017",
  },
  {
    label: "Productos",
    valor: "500+",
    cambio: "+12",
    tendencia: "up",
    icon: "products",
    color: "#171717",
  },
  {
    label: "Bodegas",
    valor: "200+",
    cambio: "+8",
    tendencia: "up",
    icon: "stores",
    color: "#16A34A",
  },
  {
    label: "Ventas del Mes",
    valor: "S/ 45K",
    cambio: "+18%",
    tendencia: "up",
    icon: "sales",
    color: "#7B2FBE",
  },
] as const

// ─── Pedidos recientes (mock) ─────────────────────────────
const pedidosRecientes = [
  { id: 1087, empresa: "Bodega El Sol", total: "S/ 1,250.00", estado: "pendiente" },
  { id: 1086, empresa: "Market La Esquina", total: "S/ 890.50", estado: "confirmado" },
  { id: 1085, empresa: "Abarrotes Don Pepe", total: "S/ 2,100.00", estado: "enviado" },
  { id: 1084, empresa: "Minimarket Central", total: "S/ 450.00", estado: "entregado" },
] as const

const estadoColor: Record<string, string> = {
  pendiente: "#F59E0B", confirmado: "#171717", enviado: "#737373", entregado: "#16A34A",
}
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

    <!-- ═══ MÉTRICAS · Grid brutalista ═══ -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-0 border border-borde mb-10">
      <div
        v-for="(m, i) in metricas"
        :key="m.label"
        class="p-6 bg-white flex flex-col justify-between gap-4"
        :class="i < 3 ? 'sm:border-r border-b sm:border-b-0 border-borde' : 'border-b sm:border-b-0 border-borde'"
      >
        <!-- Ícono + tendencia -->
        <div class="flex items-start justify-between">
          <!-- Dashboard: gráfico mini -->
          <svg v-if="m.icon === 'orders'" width="22" height="22" viewBox="0 0 22 22" fill="none" :color="m.color">
            <rect x="2" y="2" width="7" height="7" stroke="currentColor" stroke-width="1.5"/>
            <rect x="13" y="2" width="7" height="7" stroke="currentColor" stroke-width="1.5"/>
            <rect x="2" y="13" width="7" height="7" stroke="currentColor" stroke-width="1.5"/>
            <rect x="13" y="13" width="7" height="7" stroke="currentColor" stroke-width="1.5"/>
            <rect x="3" y="3" width="5" height="5" fill="currentColor" opacity="0.2"/>
          </svg>
          <!-- Productos: caja -->
          <svg v-if="m.icon === 'products'" width="22" height="22" viewBox="0 0 22 22" fill="none" :color="m.color">
            <path d="M11 2L19 6.5V15L11 19.5L3 15V6.5L11 2Z" stroke="currentColor" stroke-width="1.5"/>
            <path d="M3 6.5L11 11L19 6.5" stroke="currentColor" stroke-width="1.5"/>
          </svg>
          <!-- Bodegas: edificio -->
          <svg v-if="m.icon === 'stores'" width="22" height="22" viewBox="0 0 22 22" fill="none" :color="m.color">
            <rect x="6" y="6" width="10" height="14" stroke="currentColor" stroke-width="1.5"/>
            <path d="M3 6L11 2L19 6" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
            <rect x="10" y="12" width="2" height="8" fill="currentColor" opacity="0.3"/>
          </svg>
          <!-- Ventas: diamante -->
          <svg v-if="m.icon === 'sales'" width="22" height="22" viewBox="0 0 22 22" fill="none" :color="m.color">
            <path d="M11 1L21 11L11 21L1 11L11 1Z" stroke="currentColor" stroke-width="1.5"/>
            <circle cx="11" cy="11" r="4" fill="currentColor" opacity="0.2"/>
            <circle cx="11" cy="11" r="2" fill="currentColor"/>
          </svg>

          <!-- Indicador de tendencia -->
          <span class="font-mono text-[10px] text-exito uppercase tracking-wider font-bold">
            {{ m.cambio }}
          </span>
        </div>

        <!-- Valor masivo -->
        <div>
          <p class="font-display text-[clamp(1.75rem,3vw,2.5rem)] text-texto leading-none">
            {{ m.valor }}
          </p>
          <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mt-2">
            {{ m.label }}
          </p>
        </div>
      </div>
    </div>

    <!-- ═══ PEDIDOS RECIENTES ═══ -->
    <div class="border border-borde">
      <div class="px-6 py-4 bg-texto text-white font-mono text-[10px] uppercase tracking-[0.2em] flex items-center justify-between">
        <span>Pedidos Recientes</span>
        <NuxtLink to="/admin/pedidos" class="text-white/50 hover:text-[#D4A017] transition-colors">
          Ver Todos →
        </NuxtLink>
      </div>

      <div
        v-for="(p, i) in pedidosRecientes"
        :key="p.id"
        class="grid grid-cols-[80px_1fr_120px_100px] gap-4 px-6 py-4 items-center"
        :class="i > 0 ? 'border-t border-borde' : ''"
      >
        <span class="font-mono text-[11px] text-texto font-bold tracking-wider">#{{ p.id }}</span>
        <span class="font-body text-small text-texto truncate">{{ p.empresa }}</span>
        <span class="font-body text-small font-bold text-texto">{{ p.total }}</span>
        <span
          class="font-mono text-[9px] uppercase tracking-[0.15em] font-bold text-right"
          :style="{ color: estadoColor[p.estado] }"
        >
          {{ p.estado }}
        </span>
      </div>
    </div>

  </div>
</template>
