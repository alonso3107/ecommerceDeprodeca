<script setup lang="ts">
const searchQuery = ref("")
const mobileMenuOpen = ref(false)
const categoriasOpen = ref(false)
const cartCount = ref(0)

function actualizarCarrito() {
  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    cartCount.value = carrito.reduce((sum: number, i: any) => sum + i.cantidad, 0)
  }
}

onMounted(() => {
  actualizarCarrito()
  window.addEventListener("storage", actualizarCarrito)
  window.addEventListener("carrito-actualizado", actualizarCarrito)
})

onUnmounted(() => {
  window.removeEventListener("storage", actualizarCarrito)
  window.removeEventListener("carrito-actualizado", actualizarCarrito)
})

const categorias = [
  { nombre: "Bebidas", slug: "bebidas", desc: "Milo, Nescafé, Nestea" },
  { nombre: "Snacks", slug: "snacks", desc: "Sublime, Princesa, galletas" },
  { nombre: "Lácteos", slug: "lacteos", desc: "Ideal, Yogu Yogu, Gloria" },
  { nombre: "Abarrotes", slug: "abarrotes", desc: "Maggi, Doña Gusta, menestras" },
  { nombre: "Limpieza", slug: "limpieza", desc: "Detergente, lejía, limpiadores" },
  { nombre: "Cuidado Personal", slug: "cuidado-personal", desc: "Shampoo, jabón, desodorante" },
]

function buscar() {
  if (searchQuery.value.trim()) {
    categoriasOpen.value = false
    navigateTo(`/catalogo?q=${encodeURIComponent(searchQuery.value.trim())}`)
  }
}

function closeDropdown() { categoriasOpen.value = false }
</script>

<template>
  <header class="sticky top-0 z-sticky bg-white/80 backdrop-blur-lg border-b border-borde">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="flex items-center justify-between h-[68px] gap-4">

        <!-- Logo -->
        <NuxtLink to="/" class="flex items-center gap-2 flex-shrink-0" @click="closeDropdown">
          <span class="font-display text-display-md text-texto leading-none tracking-tight">
            depro<span class="text-texto-muted font-normal">deca</span>
          </span>
        </NuxtLink>

        <!-- Categorías Dropdown (desktop) -->
        <div class="hidden lg:block relative">
          <button
            class="flex items-center gap-2 px-4 py-2.5 rounded-xl font-body text-small font-semibold text-texto hover:bg-fondo transition-all duration-300 min-h-[44px]"
            :class="categoriasOpen ? 'bg-fondo' : ''"
            @click="categoriasOpen = !categoriasOpen"
          >
            Categorías
            <svg
              xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
              :class="['transition-transform duration-300', categoriasOpen ? 'rotate-180' : '']"
            ><path d="m6 9 6 6 6-6"/></svg>
          </button>

          <!-- Dropdown panel -->
          <Transition name="dropdown">
            <div
              v-if="categoriasOpen"
              class="absolute top-full left-0 mt-3 w-[540px] bg-white rounded-2xl shadow-xl border border-borde p-4 grid grid-cols-2 gap-1 z-dropdown"
              @mouseleave="closeDropdown"
            >
              <NuxtLink
                v-for="cat in categorias"
                :key="cat.slug"
                :to="`/catalogo?categoria=${cat.slug}`"
                class="flex flex-col gap-0.5 px-4 py-3 rounded-xl hover:bg-fondo transition-colors duration-200 group"
                @click="closeDropdown"
              >
                <span class="font-body text-small font-semibold text-texto group-hover:text-[#D4A017] transition-colors">
                  {{ cat.nombre }}
                </span>
                <span class="font-body text-caption text-texto-muted">{{ cat.desc }}</span>
              </NuxtLink>
            </div>
          </Transition>
        </div>

        <!-- Spacer -->
        <div class="flex-1 hidden md:block" />

        <!-- Search (desktop) -->
        <form class="hidden md:flex relative max-w-[300px]" @submit.prevent="buscar">
          <input
            v-model="searchQuery"
            type="search"
            placeholder="Buscar..."
            class="w-full rounded-xl bg-fondo border border-borde px-4 py-2.5 pl-9 font-body text-small text-texto placeholder:text-texto-muted focus:border-texto/20 focus:ring-2 focus:ring-texto/5 transition-all duration-300 min-h-[40px]"
          />
          <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-texto-muted pointer-events-none" xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
        </form>

        <!-- Actions -->
        <div class="flex items-center gap-1">
          <NuxtLink to="/carrito" class="relative p-2 rounded-xl hover:bg-fondo transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center" aria-label="Carrito" @click="closeDropdown">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 002 1.58h9.78a2 2 0 001.95-1.57l1.65-7.43H5.12"/></svg>
            <span v-if="cartCount > 0" class="absolute -top-0.5 -right-0.5 bg-texto text-white text-[10px] font-display font-bold rounded-full h-5 w-5 flex items-center justify-center">{{ cartCount }}</span>
          </NuxtLink>

          <NuxtLink to="/auth/login" class="p-2 rounded-xl hover:bg-fondo transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center" aria-label="Perfil" @click="closeDropdown">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21v-2a4 4 0 00-4-4H9a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
          </NuxtLink>

          <!-- Hamburger (mobile) -->
          <button class="lg:hidden p-2 rounded-xl hover:bg-fondo transition-colors min-h-[44px] min-w-[44px] flex items-center justify-center" aria-label="Menú" @click="mobileMenuOpen = !mobileMenuOpen">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
          </button>
        </div>
      </div>

      <!-- Mobile menu -->
      <Transition name="slide">
        <div v-if="mobileMenuOpen" class="lg:hidden border-t border-borde bg-white pb-4">
          <form class="px-4 pt-4" @submit.prevent="buscar(); mobileMenuOpen = false">
            <input v-model="searchQuery" type="search" placeholder="Buscar productos..." class="w-full rounded-xl bg-fondo border border-borde px-4 py-3 font-body text-body text-texto min-h-[44px]" />
          </form>
          <nav class="px-4 pt-3 space-y-1">
            <NuxtLink v-for="cat in categorias" :key="cat.slug" :to="`/catalogo?categoria=${cat.slug}`" class="flex flex-col px-3 py-2.5 rounded-lg hover:bg-fondo transition-colors min-h-[44px]" @click="mobileMenuOpen = false">
              <span class="font-body text-body font-semibold text-texto">{{ cat.nombre }}</span>
              <span class="font-body text-caption text-texto-muted">{{ cat.desc }}</span>
            </NuxtLink>
          </nav>
        </div>
      </Transition>
    </div>
  </header>
</template>

<style scoped>
.dropdown-enter-active { transition: all 250ms ease-out; }
.dropdown-leave-active { transition: all 150ms ease-in; }
.dropdown-enter-from { opacity: 0; transform: translateY(-8px) scale(0.98); }
.dropdown-leave-to { opacity: 0; transform: translateY(-8px) scale(0.98); }

.slide-enter-active, .slide-leave-active { transition: all 300ms ease; }
.slide-enter-from, .slide-leave-to { opacity: 0; max-height: 0; }
.slide-enter-to, .slide-leave-from { opacity: 1; max-height: 500px; }
</style>
