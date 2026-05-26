<!--
  ═══════════════════════════════════════════════════════════════
  PerfilSidebar.vue — Sidebar de navegación brutalista
  Bloques rectangulares duros, sin curvas. Ícono geométrico
  único para cada sección. Activo = fondo negro + texto blanco.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
const links = [
  {
    to: "/perfil",
    label: "Datos Personales",
    icon: "user",
  },
  {
    to: "/perfil/rangos",
    label: "Mi Rango",
    icon: "rank",
  },
  {
    to: "/pedidos",
    label: "Mis Pedidos",
    icon: "orders",
  },
] as const

const route = useRoute()
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
  <nav class="flex flex-col gap-0 border border-borde">
    <NuxtLink
      v-for="link in links"
      :key="link.to"
      :to="link.to"
      class="group flex items-center gap-3 px-5 py-4 font-mono text-[11px] uppercase tracking-[0.15em]
             transition-colors duration-150 min-h-[52px]"
      :class="isActive(link.to)
        ? 'bg-texto text-white font-bold'
        : 'bg-white text-texto-muted hover:text-texto hover:bg-fondo'"
    >

      <!-- ÍCONO: Persona geométrica (user) -->
      <svg v-if="link.icon === 'user'" width="18" height="18" viewBox="0 0 18 18" fill="none"
           :class="isActive(link.to) ? 'text-white' : 'text-texto-muted'">
        <circle cx="9" cy="6" r="3" stroke="currentColor" stroke-width="1.5"/>
        <path d="M3 16C3 13.2 5.7 10.5 9 10.5C12.3 10.5 15 13.2 15 16"
              stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
      </svg>

      <!-- ÍCONO: Corona/rombo geométrica (rank) -->
      <svg v-if="link.icon === 'rank'" width="18" height="18" viewBox="0 0 18 18" fill="none"
           :class="isActive(link.to) ? 'text-white' : 'text-texto-muted'">
        <path d="M4 14L9 4L14 14L9 10L4 14Z" stroke="currentColor" stroke-width="1.5"
              stroke-linecap="square" stroke-linejoin="miter"/>
        <rect x="6" y="13" width="6" height="2" fill="currentColor"/>
      </svg>

      <!-- ÍCONO: Caja/paquete geométrica (orders) -->
      <svg v-if="link.icon === 'orders'" width="18" height="18" viewBox="0 0 18 18" fill="none"
           :class="isActive(link.to) ? 'text-white' : 'text-texto-muted'">
        <rect x="2" y="5" width="14" height="12" stroke="currentColor" stroke-width="1.5"/>
        <path d="M2 5L9 1L16 5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
        <line x1="9" y1="1" x2="9" y2="10" stroke="currentColor" stroke-width="1" stroke-dasharray="1.5 1.5"/>
      </svg>

      <span>{{ link.label }}</span>
    </NuxtLink>

    <!-- Cerrar Sesión -->
    <button
      class="group flex items-center gap-3 px-5 py-4 font-mono text-[11px] uppercase tracking-[0.15em]
             transition-colors duration-150 min-h-[52px] w-full text-left
             bg-white text-error hover:bg-error hover:text-white mt-auto border-t border-borde"
      @click="cerrarSesion"
    >
      <!-- ÍCONO: Salida/puerta geométrica -->
      <svg width="18" height="18" viewBox="0 0 18 18" fill="none"
           class="text-error group-hover:text-white">
        <path d="M7 3H3V15H7" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
        <path d="M11 13L15 9L11 5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
        <line x1="15" y1="9" x2="6" y2="9" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
      </svg>
      <span>Cerrar Sesión</span>
    </button>
  </nav>
</template>
