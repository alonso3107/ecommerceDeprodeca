<!--
  ═══════════════════════════════════════════════════════════════
  admim.vue — Layout Admin · DEPRODECA
  Brutalismo industrial: sidebar negro con iconos geometricos,
  contenido con bordes duros. Toolbar tecnica superior.
  Textos aumentados para legibilidad.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
const route = useRoute()

const links = [
  { to: "/admin/dashboard", label: "Dashboard", icon: "dashboard" },
  { to: "/admin/productos", label: "Productos", icon: "products" },
  { to: "/admin/pedidos", label: "Pedidos", icon: "orders" },
] as const

function isActive(path: string) { return route.path === path }

function cerrarSesion() {
  if (import.meta.client) {
    localStorage.removeItem("deprodeca_token")
    localStorage.removeItem("deprodeca_user")
    window.dispatchEvent(new Event("auth-cambiado"))
  }
  navigateTo("/")
}
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
        <p class="mt-2 font-mono text-xs text-white/30 uppercase tracking-[0.2em]">
          Panel de Administracion
        </p>
      </div>

      <!-- Navegacion -->
      <nav class="flex-1 py-6">
        <NuxtLink
          v-for="link in links"
          :key="link.to"
          :to="link.to"
          class="flex items-center gap-3 mx-4 px-5 py-3.5 font-mono text-sm uppercase tracking-[0.12em]
                 transition-all duration-150"
          :class="isActive(link.to)
            ? 'bg-[#D4A017] text-black font-bold'
            : 'text-white/50 hover:text-white hover:bg-white/5'"
        >
          <!-- Dashboard: Grilla geometrica -->
          <svg v-if="link.icon === 'dashboard'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <rect x="2" y="2" width="6" height="6" stroke="currentColor" stroke-width="1.8"/>
            <rect x="10" y="2" width="6" height="6" stroke="currentColor" stroke-width="1.8"/>
            <rect x="2" y="10" width="6" height="6" stroke="currentColor" stroke-width="1.8"/>
            <rect x="10" y="10" width="6" height="6" stroke="currentColor" stroke-width="1.8"/>
          </svg>

          <!-- Productos: Caja isometrica -->
          <svg v-if="link.icon === 'products'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <path d="M9 2L15 5.5V12L9 15.5L3 12V5.5L9 2Z" stroke="currentColor" stroke-width="1.8"/>
            <path d="M3 5.5L9 9L15 5.5" stroke="currentColor" stroke-width="1.8"/>
            <path d="M9 9V15.5" stroke="currentColor" stroke-width="1.8"/>
          </svg>

          <!-- Pedidos: Documento/lista -->
          <svg v-if="link.icon === 'orders'" width="18" height="18" viewBox="0 0 18 18" fill="none">
            <rect x="4" y="2" width="10" height="15" stroke="currentColor" stroke-width="1.8"/>
            <line x1="7" y1="6" x2="11" y2="6" stroke="currentColor" stroke-width="1.4"/>
            <line x1="7" y1="9" x2="11" y2="9" stroke="currentColor" stroke-width="1.4"/>
            <line x1="7" y1="12" x2="9" y2="12" stroke="currentColor" stroke-width="1.4"/>
          </svg>

          {{ link.label }}
        </NuxtLink>
      </nav>

      <!-- Footer sidebar -->
      <div class="px-6 py-5 border-t border-white/10 space-y-3">
        <button
          class="w-full text-left font-mono text-sm text-white/40 hover:text-error uppercase tracking-[0.12em] transition-colors flex items-center gap-2"
          @click="cerrarSesion"
        >
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
            <path d="M6 3H3V13H6" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
            <path d="M10 12L14 8L10 4" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
            <line x1="14" y1="8" x2="5" y2="8" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
          </svg>
          Cerrar Sesion
        </button>
        <NuxtLink
          to="/"
          class="block font-mono text-sm text-white/30 hover:text-[#D4A017] uppercase tracking-[0.12em] transition-colors"
        >
          Volver a la Tienda
        </NuxtLink>
      </div>
    </aside>

    <!-- ═══ CONTENIDO ═══ -->
    <div class="flex-1 flex flex-col min-w-0">

      <!-- Toolbar superior -->
      <header class="bg-white border-b border-borde px-6 py-3.5 flex items-center justify-between">
        <span class="font-mono text-sm text-texto-muted uppercase tracking-[0.15em]">
          Admin · {{ links.find(l => isActive(l.to))?.label || "DEPRODECA" }}
        </span>

        <!-- Indicador online -->
        <div class="flex items-center gap-2">
          <span class="w-2.5 h-2.5 bg-exito" />
          <span class="font-mono text-xs text-texto-muted uppercase tracking-[0.12em]">Online</span>
        </div>
      </header>

      <!-- Contenido de la pagina -->
      <main class="flex-1 p-6 md:p-8">
        <slot />
      </main>
    </div>

  </div>
</template>
