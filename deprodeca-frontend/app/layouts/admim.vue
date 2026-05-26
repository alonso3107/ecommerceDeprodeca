<!--
  ═══════════════════════════════════════════════════════════════
  admim.vue — Layout Admin · DEPRODECA
  Brutalismo industrial: sidebar negro con íconos geométricos,
  contenido con bordes duros. Toolbar técnica superior.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
// Layout admin — protegido por middleware/admin.ts
const route = useRoute()

const links = [
  {
    to: "/admin/dashboard",
    label: "Dashboard",
    icon: "dashboard",
  },
  {
    to: "/admin/productos",
    label: "Productos",
    icon: "products",
  },
  {
    to: "/admin/pedidos",
    label: "Pedidos",
    icon: "orders",
  },
] as const

function isActive(path: string) { return route.path === path }
</script>

<template>
  <div class="min-h-screen flex bg-fondo">

    <!-- ═══ SIDEBAR · Negro industrial ═══ -->
    <aside class="w-64 bg-texto text-white hidden lg:flex flex-col flex-shrink-0">

      <!-- Logo -->
      <div class="px-6 py-8 border-b border-white/10">
        <NuxtLink to="/" class="inline-block">
          <span class="font-display text-display-md leading-none tracking-tight">
            DEPRO<span class="text-[#D4A017]">DECA</span>
          </span>
        </NuxtLink>
        <p class="mt-2 font-mono text-[9px] text-white/30 uppercase tracking-[0.2em]">
          Panel de Administración
        </p>
      </div>

      <!-- Navegación -->
      <nav class="flex-1 py-4">
        <NuxtLink
          v-for="link in links"
          :key="link.to"
          :to="link.to"
          class="flex items-center gap-3 mx-3 px-4 py-3 font-mono text-[10px] uppercase tracking-[0.15em]
                 transition-all duration-150"
          :class="isActive(link.to)
            ? 'bg-[#D4A017] text-black font-bold'
            : 'text-white/50 hover:text-white hover:bg-white/5'"
        >

          <!-- Dashboard: Grilla geométrica -->
          <svg v-if="link.icon === 'dashboard'" width="16" height="16" viewBox="0 0 16 16" fill="none">
            <rect x="1" y="1" width="6" height="6" stroke="currentColor" stroke-width="1.5"/>
            <rect x="9" y="1" width="6" height="6" stroke="currentColor" stroke-width="1.5"/>
            <rect x="1" y="9" width="6" height="6" stroke="currentColor" stroke-width="1.5"/>
            <rect x="9" y="9" width="6" height="6" stroke="currentColor" stroke-width="1.5"/>
          </svg>

          <!-- Productos: Caja isométrica -->
          <svg v-if="link.icon === 'products'" width="16" height="16" viewBox="0 0 16 16" fill="none">
            <path d="M8 2L14 5.5V11L8 14.5L2 11V5.5L8 2Z" stroke="currentColor" stroke-width="1.5"/>
            <path d="M2 5.5L8 9L14 5.5" stroke="currentColor" stroke-width="1.5"/>
            <path d="M8 9V14.5" stroke="currentColor" stroke-width="1.5"/>
          </svg>

          <!-- Pedidos: Documento/lista -->
          <svg v-if="link.icon === 'orders'" width="16" height="16" viewBox="0 0 16 16" fill="none">
            <rect x="3" y="1" width="10" height="14" stroke="currentColor" stroke-width="1.5"/>
            <line x1="6" y1="5" x2="10" y2="5" stroke="currentColor" stroke-width="1.2"/>
            <line x1="6" y1="8" x2="10" y2="8" stroke="currentColor" stroke-width="1.2"/>
            <line x1="6" y1="11" x2="8" y2="11" stroke="currentColor" stroke-width="1.2"/>
          </svg>

          {{ link.label }}
        </NuxtLink>
      </nav>

      <!-- Footer sidebar -->
      <div class="px-6 py-4 border-t border-white/10">
        <NuxtLink
          to="/"
          class="font-mono text-[9px] text-white/30 hover:text-[#D4A017] uppercase tracking-[0.2em] transition-colors"
        >
          ← Volver a la Tienda
        </NuxtLink>
      </div>
    </aside>

    <!-- ═══ CONTENIDO ═══ -->
    <div class="flex-1 flex flex-col min-w-0">

      <!-- Toolbar superior -->
      <header class="bg-white border-b border-borde px-6 py-3 flex items-center justify-between">
        <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em]">
          Admin · {{ links.find(l => isActive(l.to))?.label || "DEPRODECA" }}
        </span>

        <!-- Indicador online -->
        <div class="flex items-center gap-2">
          <span class="w-2 h-2 bg-exito" />
          <span class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.15em]">Online</span>
        </div>
      </header>

      <!-- Contenido de la página -->
      <main class="flex-1 p-6 md:p-8">
        <slot />
      </main>
    </div>

  </div>
</template>
