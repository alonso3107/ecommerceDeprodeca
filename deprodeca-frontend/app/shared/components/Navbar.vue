<script setup lang="ts">
const searchQuery = ref("")
const mobileMenuOpen = ref(false)
const categoriasOpen = ref(false)
const cartCount = ref(0)
const autenticado = ref(false)
const esAdmin = ref(false)

function verificarAuth() {
  if (import.meta.client) {
    const token = localStorage.getItem("deprodeca_token")
    autenticado.value = !!token
    if (token) {
      try {
        const user = JSON.parse(localStorage.getItem("deprodeca_user") || "{}")
        esAdmin.value = user.rol === "admin"
      } catch { esAdmin.value = false }
    } else {
      esAdmin.value = false
    }
  }
}

function actualizarCarrito() {
  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    cartCount.value = carrito.reduce((sum: number, i: any) => sum + i.cantidad, 0)
  }
}

onMounted(() => {
  verificarAuth()
  actualizarCarrito()
  window.addEventListener("storage", actualizarCarrito)
  window.addEventListener("carrito-actualizado", actualizarCarrito)
  window.addEventListener("auth-cambiado", verificarAuth)
})

onUnmounted(() => {
  window.removeEventListener("storage", actualizarCarrito)
  window.removeEventListener("carrito-actualizado", actualizarCarrito)
  window.removeEventListener("auth-cambiado", verificarAuth)
})

const categorias = [
  { nombre: "Bebidas", slug: "bebidas", desc: "Milo, Nescafé, Nestea" },
  { nombre: "Snacks", slug: "snacks", desc: "Sublime, Princesa, galletas" },
  { nombre: "Lácteos", slug: "lacteos", desc: "Ideal, Yogu Yogu, Gloria" },
  { nombre: "Abarrotes", slug: "abarrotes", desc: "Maggi, Doña Gusta, menestras" },
  { nombre: "Limpieza", slug: "limpieza", desc: "Detergente, lejía, limpiadores" },
  { nombre: "Cuidado Personal", slug: "cuidado-personal", desc: "Shampoo, jabón, desodorante" },
]

// ─── Dropdown posicionamiento (Teleport al body) ──────
const catBtnRef = ref<HTMLElement | null>(null)
const dropdownStyle = ref<Record<string, string>>({})

function abrirCategorias() {
  if (!catBtnRef.value) return
  const rect = catBtnRef.value.getBoundingClientRect()
  dropdownStyle.value = {
    position: "fixed",
    top: `${rect.bottom + 12}px`,
    left: `${rect.left}px`,
    width: "540px",
    zIndex: "100",
  }
  categoriasOpen.value = true
  // Cerrar con clic fuera
  setTimeout(() => document.addEventListener("click", cerrarCategoriasFuera), 0)
}

function cerrarCategoriasFuera(e: MouseEvent) {
  if (catBtnRef.value && !catBtnRef.value.contains(e.target as Node)) {
    categoriasOpen.value = false
    document.removeEventListener("click", cerrarCategoriasFuera)
  }
}

function closeDropdown() {
  categoriasOpen.value = false
  document.removeEventListener("click", cerrarCategoriasFuera)
}

function cerrarSesion() {
  if (import.meta.client) {
    localStorage.removeItem("deprodeca_token")
    localStorage.removeItem("deprodeca_user")
    window.dispatchEvent(new Event("auth-cambiado"))
  }
  closeDropdown()
  navigateTo("/")
}

function buscar() {
  if (searchQuery.value.trim()) {
    closeDropdown()
    navigateTo(`/catalogo?q=${encodeURIComponent(searchQuery.value.trim())}`)
  }
}

</script>

<template>
  <header class="sticky top-0 z-50 h-[68px] bg-[#FAFAF9] border-b border-[#D6D3D1]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="flex items-center justify-between h-[68px] gap-4">

        <!-- Logo -->
        <NuxtLink to="/" class="flex items-center gap-2 flex-shrink-0" @click="closeDropdown">
          <svg width="18" height="18" viewBox="0 0 18 18" fill="none" aria-hidden="true">
            <path d="M9 1.5L16.5 9L9 16.5L1.5 9L9 1.5Z" stroke="#A16207" stroke-width="1.5" stroke-linejoin="miter"/>
          </svg>
          <span class="font-display font-bold text-xl text-[#1C1917] leading-none tracking-tight">
            DEPRODECA
          </span>
        </NuxtLink>

        <!-- Categorías Dropdown (desktop) -->
        <div class="hidden lg:block">
          <button
            ref="catBtnRef"
            class="flex items-center gap-2 px-3 py-2 font-body text-[14px] font-medium text-[#1C1917] hover:text-[#A16207] transition-colors duration-200 min-h-[44px]"
            @click="categoriasOpen ? closeDropdown() : abrirCategorias()"
          >
            Categorías
            <svg
              width="14" height="14" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"
              :class="['transition-transform duration-200', categoriasOpen ? 'rotate-180' : '']"
            ><path d="M2 5L7 10L12 5"/></svg>
          </button>
        </div>

        <!-- Dropdown panel (Teleport al body para evitar stacking context del header) -->
        <Teleport to="body">
          <Transition name="dropdown">
            <div
              v-if="categoriasOpen"
              :style="dropdownStyle"
              class="bg-[#FAFAF9] border border-[#D6D3D1] p-6 grid grid-cols-2 gap-2"
              @mouseleave="closeDropdown"
            >
              <NuxtLink
                v-for="cat in categorias"
                :key="cat.slug"
                :to="`/catalogo?categoria=${cat.slug}`"
                class="flex flex-col gap-0.5 px-3 py-2.5 hover:bg-[#F5F0E8] transition-colors duration-200"
                @click="closeDropdown"
              >
                <span class="font-body text-[14px] font-semibold text-[#1C1917]">
                  {{ cat.nombre }}
                </span>
                <span class="font-mono text-[10px] text-[#78716C] uppercase tracking-[0.1em]">{{ cat.desc }}</span>
              </NuxtLink>
            </div>
          </Transition>
        </Teleport>

        <!-- Spacer -->
        <div class="flex-1 hidden md:block" />

        <!-- Search (desktop) -->
        <form class="hidden md:flex relative max-w-[220px] w-full" @submit.prevent="buscar">
          <svg class="absolute left-1 top-1/2 -translate-y-1/2 text-[#1C1917] pointer-events-none" width="14" height="14" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter" aria-hidden="true"><circle cx="6" cy="6" r="4.5"/><path d="M9.5 9.5L13 13"/></svg>
          <input
            v-model="searchQuery"
            type="search"
            placeholder="Buscar productos..."
            class="w-full border-0 border-b border-[#D6D3D1] bg-transparent py-2 px-1 pl-6 font-body text-[14px] text-[#1C1917] placeholder:text-[#78716C] focus:border-b-[#A16207] focus:outline-none transition-colors duration-200 min-h-[40px]"
          />
        </form>

        <!-- Actions -->
        <div class="flex items-center gap-1">
          <!-- Admin Dashboard (solo admin) -->
          <NuxtLink
            v-if="esAdmin"
            to="/admin/dashboard"
            class="font-mono text-[10px] uppercase tracking-[0.15em] px-3 py-1.5 border border-[#A16207] text-[#A16207] hover:bg-[#A16207] hover:text-[#FAFAF9] transition-colors duration-200 min-h-[44px] flex items-center"
            aria-label="Dashboard Admin"
            @click="closeDropdown"
          >
            ADMIN
          </NuxtLink>

          <NuxtLink to="/carrito" class="relative w-[44px] h-[44px] flex items-center justify-center text-[#1C1917] hover:text-[#A16207] transition-colors duration-200" aria-label="Carrito" @click="closeDropdown">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"><path d="M3 6H17L15.5 15H4.5L3 6Z"/><path d="M6 6V4.8C6 3.8 6.8 3 7.8 3H12.2C13.2 3 14 3.8 14 4.8V6"/></svg>
            <span v-if="cartCount > 0" class="absolute -top-1 -right-1 w-4 h-4 bg-[#A16207] text-[#FAFAF9] font-mono text-[9px] font-bold flex items-center justify-center">{{ cartCount }}</span>
          </NuxtLink>

          <!-- Perfil (condicional: logueado → /perfil, no logueado → /auth/login) -->
          <NuxtLink
            :to="autenticado ? '/perfil' : '/auth/login'"
            class="relative w-[44px] h-[44px] flex items-center justify-center text-[#1C1917] hover:text-[#A16207] transition-colors duration-200"
            :aria-label="autenticado ? 'Perfil' : 'Iniciar sesión'"
            @click="closeDropdown"
          >
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"><circle cx="10" cy="6" r="3.2"/><path d="M3.5 16.2C3.5 13.8 6.3 12 10 12C13.7 12 16.5 13.8 16.5 16.2"/></svg>
            <!-- Indicador de sesión activa -->
            <span v-if="autenticado" class="absolute top-0 right-0 w-[6px] h-[6px] bg-[#A16207]" />
          </NuxtLink>

          <!-- Cerrar Sesión (desktop, solo visible con sesión) -->
            <button
              v-if="autenticado"
              class="hidden sm:flex items-center w-[44px] h-[44px] justify-center text-[#DC2626] hover:text-[#A16207] transition-colors duration-200"
              aria-label="Cerrar sesión"
              title="Cerrar sesión"
              @click="cerrarSesion"
            >
              <svg width="18" height="18" viewBox="0 0 18 18" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"><path d="M2.5 2.5H7.5V15.5H2.5Z"/><path d="M10 9H16"/><path d="M13 6L16 9L13 12"/></svg>
            </button>

          <!-- Hamburger (mobile) -->
          <button class="lg:hidden w-[44px] h-[44px] flex items-center justify-center text-[#1C1917] hover:text-[#A16207] transition-colors duration-200" aria-label="Menú" @click="mobileMenuOpen = !mobileMenuOpen">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"><line x1="3" y1="5" x2="17" y2="5"/><line x1="3" y1="10" x2="17" y2="10"/><line x1="3" y1="15" x2="17" y2="15"/></svg>
          </button>
        </div>
      </div>

      <!-- Mobile menu -->
      <Transition name="slide">
        <div v-if="mobileMenuOpen" class="lg:hidden border-t border-[#D6D3D1] bg-[#FAFAF9]">
          <form class="px-4 pt-4" @submit.prevent="buscar(); mobileMenuOpen = false">
            <input v-model="searchQuery" type="search" placeholder="Buscar productos..." class="w-full py-3 border-0 border-b border-[#D6D3D1] bg-transparent px-1 font-body text-[14px] text-[#1C1917] focus:outline-none focus:border-b-[#A16207]" />
          </form>
          <nav class="pt-3">
            <NuxtLink v-for="cat in categorias" :key="cat.slug" :to="`/catalogo?categoria=${cat.slug}`" class="block font-body text-[14px] text-[#1C1917] py-3 px-4 border-b border-[#D6D3D1]/40 hover:bg-[#F5F0E8] transition-colors duration-200" @click="mobileMenuOpen = false">
              {{ cat.nombre }}
            </NuxtLink>

            <!-- Separador + Cerrar Sesión (mobile) -->
            <template v-if="autenticado">
              <div class="border-t border-[#D6D3D1]" />
              <button
                class="flex items-center gap-3 w-full px-4 py-3 font-mono text-[10px] uppercase tracking-[0.15em] text-[#DC2626] hover:bg-[#F5F0E8] transition-colors duration-200 min-h-[44px]"
                @click="cerrarSesion(); mobileMenuOpen = false"
              >
                <svg width="16" height="16" viewBox="0 0 18 18" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"><path d="M2.5 2.5H7.5V15.5H2.5Z"/><path d="M10 9H16"/><path d="M13 6L16 9L13 12"/></svg>
                Cerrar Sesión
              </button>
            </template>
          </nav>
        </div>
      </Transition>
    </div>
  </header>
</template>

<style scoped>
.dropdown-enter-active { transition: opacity 180ms ease-out, transform 180ms ease-out; }
.dropdown-leave-active { transition: opacity 140ms ease-in, transform 140ms ease-in; }
.dropdown-enter-from { opacity: 0; transform: translateY(-4px); }
.dropdown-leave-to { opacity: 0; transform: translateY(-4px); }

.slide-enter-active, .slide-leave-active { transition: all 300ms ease; }
.slide-enter-from, .slide-leave-to { opacity: 0; max-height: 0; }
.slide-enter-to, .slide-leave-from { opacity: 1; max-height: 500px; }
</style>
