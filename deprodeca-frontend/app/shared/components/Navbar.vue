     1|<script setup lang="ts">
     2|import { ref, computed, onMounted, onUnmounted } from "vue"
     3|
     4|// ─── Estado local ─────────────────────────────────────
     5|const searchQuery = ref("")
     6|const mobileMenuOpen = ref(false)
     7|const categoriasOpen = ref(false)
     8|const cartCount = ref(0)
     9|
    10|// Leer carrito del localStorage
    11|function actualizarCarrito() {
    12|  if (import.meta.client) {
    13|    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    14|    cartCount.value = carrito.reduce((sum: number, i: any) => sum + i.cantidad, 0)
    15|  }
    16|}
    17|
    18|onMounted(() => {
    19|  actualizarCarrito()
    20|  // Escuchar cambios en localStorage desde otras pestañas
    21|  window.addEventListener("storage", actualizarCarrito)
    22|  // Escuchar evento custom (cuando se modifica en la misma pestaña)
    23|  window.addEventListener("carrito-actualizado", actualizarCarrito)
    24|})
    25|
    26|onUnmounted(() => {
    27|  window.removeEventListener("storage", actualizarCarrito)
    28|  window.removeEventListener("carrito-actualizado", actualizarCarrito)
    29|})
    30|
    31|// Categorías del mega menú
    32|const categorias = [
    33|  { nombre: "Bebidas", slug: "bebidas", icono: "M9 20H4V9h5v11zm11 0h-5v-5h5v5zm0-6h-5V4h5v10z" },
    34|  { nombre: "Snacks", slug: "snacks", icono: "M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5" },
    35|  { nombre: "Lácteos", slug: "lacteos", icono: "M20 9V7a2 2 0 00-2-2h-4a2 2 0 00-2 2v2M4 9V7a2 2 0 012-2h4a2 2 0 012 2v2" },
    36|  { nombre: "Abarrotes", slug: "abarrotes", icono: "M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2z" },
    37|  { nombre: "Limpieza", slug: "limpieza", icono: "M12 2L2 7l10 5 10-5-10-5z" },
    38|  { nombre: "Cuidado Personal", slug: "cuidado-personal", icono: "M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z" },
    39|]
    40|
    41|function buscar() {
    42|  if (searchQuery.value.trim()) {
    43|    navigateTo(`/catalogo?q=${encodeURIComponent(searchQuery.value.trim())}`)
    44|  }
    45|}
    46|</script>
    47|
    48|<template>
    49|  <header class="sticky top-0 z-sticky bg-white border-b-2 border-border-base">
    50|    <div class="max-w-[1280px] mx-auto px-4 md:px-6">
    51|      <div class="flex items-center justify-between h-[72px] gap-4">
    52|
    53|        <!-- ─── LOGO ─────────────────────────────────── -->
    54|        <NuxtLink to="/" class="flex items-center gap-2 flex-shrink-0">
    55|          <span class="font-display text-display-md text-text-primary leading-none tracking-tight">
    56|            DEPRO<span class="text-brand-primary">DECA</span>
    57|          </span>
    58|        </NuxtLink>
    59|
    60|        <!-- ─── MEGA MENÚ CATEGORÍAS (Desktop) ────────── -->
    61|        <div class="hidden lg:block relative">
    62|          <button
    63|            class="flex items-center gap-1.5 px-3 py-2 font-body text-small font-semibold text-text-primary hover:text-brand-primary transition-colors duration-300 min-h-[44px]"
    64|            @click="categoriasOpen = !categoriasOpen"
    65|            @blur="setTimeout(() => categoriasOpen = false, 200)"
    66|          >
    67|            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
    68|            Categorías
    69|          </button>
    70|
    71|          <!-- Dropdown mega menú -->
    72|          <Transition name="mega">
    73|            <div
    74|              v-if="categoriasOpen"
    75|              class="absolute top-full left-0 mt-2 w-[560px] bg-white border-2 border-border-base rounded-xl shadow-lg p-4 grid grid-cols-3 gap-2"
    76|            >
    77|              <NuxtLink
    78|                v-for="cat in categorias"
    79|                :key="cat.slug"
    80|                :to="`/catalogo?categoria=${cat.slug}`"
    81|                class="flex items-center gap-2 px-3 py-2.5 rounded-lg hover:bg-surface transition-colors duration-300 font-body text-small font-medium text-text-primary hover:text-brand-secondary"
    82|                @click="categoriasOpen = false"
    83|              >
    84|                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path :d="cat.icono"/></svg>
    85|                {{ cat.nombre }}
    86|              </NuxtLink>
    87|            </div>
    88|          </Transition>
    89|        </div>
    90|
    91|        <!-- ─── SEARCH BAR ────────────────────────────── -->
    92|        <form
    93|          class="hidden md:flex flex-1 max-w-[480px] relative"
    94|          @submit.prevent="buscar"
    95|        >
    96|          <input
    97|            v-model="searchQuery"
    98|            type="search"
    99|            placeholder="Buscar productos, marcas..."
   100|            class="w-full rounded-lg border-2 border-border-base bg-surface px-4 py-2.5 pl-10 font-body text-body text-text-primary placeholder:text-text-muted focus:border-brand-primary focus:ring-2 focus:ring-brand-primary/20 transition-all duration-300 min-h-[44px]"
   101|          />
   102|          <svg
   103|            class="absolute left-3 top-1/2 -translate-y-1/2 text-text-muted pointer-events-none"
   104|            xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"
   105|            aria-hidden="true"
   106|          >
   107|            <circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/>
   108|          </svg>
   109|          <button
   110|            type="submit"
   111|            class="absolute right-1 top-1/2 -translate-y-1/2 px-3 py-1.5 rounded-md bg-brand-primary text-text-primary font-body text-small font-semibold hover:shadow-brand transition-all duration-300 min-h-[36px]"
   112|          >
   113|            Buscar
   114|          </button>
   115|        </form>
   116|
   117|        <!-- ─── ACCIONES (Desktop) ────────────────────── -->
   118|        <div class="hidden md:flex items-center gap-1">
   119|          <!-- Carrito -->
   120|          <NuxtLink
   121|            to="/carrito"
   122|            class="relative p-2.5 rounded-lg hover:bg-surface transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center"
   123|            aria-label="Carrito de compras"
   124|          >
   125|            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 002 1.58h9.78a2 2 0 001.95-1.57l1.65-7.43H5.12"/></svg>
   126|            <span class="absolute -top-0.5 -right-0.5 bg-brand-primary text-text-primary text-[10px] font-display font-bold rounded-full h-5 w-5 flex items-center justify-center">{{ cartCount }}</span>
   127|          </NuxtLink>
   128|
   129|          <!-- Perfil / Login -->
   130|          <NuxtLink
   131|            to="/auth/login"
   132|            class="p-2.5 rounded-lg hover:bg-surface transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center"
   133|            aria-label="Iniciar sesión"
   134|          >
   135|            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21v-2a4 4 0 00-4-4H9a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
   136|          </NuxtLink>
   137|        </div>
   138|
   139|        <!-- ─── HAMBURGER (Mobile) ────────────────────── -->
   140|        <button
   141|          class="lg:hidden p-2.5 rounded-lg hover:bg-surface transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center"
   142|          aria-label="Menú"
   143|          @click="mobileMenuOpen = !mobileMenuOpen"
   144|        >
   145|          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
   146|        </button>
   147|      </div>
   148|
   149|      <!-- ─── MOBILE MENU ─────────────────────────────── -->
   150|      <Transition name="slide">
   151|        <div
   152|          v-if="mobileMenuOpen"
   153|          class="lg:hidden border-t border-border-base bg-white pb-4"
   154|        >
   155|          <!-- Search mobile -->
   156|          <form class="px-4 pt-4" @submit.prevent="buscar(); mobileMenuOpen = false">
   157|            <div class="relative">
   158|              <input
   159|                v-model="searchQuery"
   160|                type="search"
   161|                placeholder="Buscar productos..."
   162|                class="w-full rounded-lg border-2 border-border-base bg-surface px-4 py-2.5 pl-10 font-body text-body text-text-primary min-h-[44px]"
   163|              />
   164|              <svg class="absolute left-3 top-1/2 -translate-y-1/2 text-text-muted" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
   165|            </div>
   166|          </form>
   167|
   168|          <!-- Categorías mobile -->
   169|          <nav class="px-4 pt-3 space-y-1">
   170|            <NuxtLink
   171|              v-for="cat in categorias"
   172|              :key="cat.slug"
   173|              :to="`/catalogo?categoria=${cat.slug}`"
   174|              class="flex items-center gap-2 px-3 py-2.5 rounded-lg hover:bg-surface transition-colors font-body text-body text-text-primary min-h-[44px]"
   175|              @click="mobileMenuOpen = false"
   176|            >
   177|              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path :d="cat.icono"/></svg>
   178|              {{ cat.nombre }}
   179|            </NuxtLink>
   180|          </nav>
   181|
   182|          <!-- Acciones mobile -->
   183|          <div class="flex gap-2 px-4 pt-3">
   184|            <NuxtLink
   185|              to="/carrito"
   186|              class="flex items-center gap-2 px-4 py-2.5 rounded-lg bg-surface hover:bg-border-base transition-colors font-body text-small font-semibold text-text-primary min-h-[44px]"
   187|              @click="mobileMenuOpen = false"
   188|            >
   189|              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 002 1.58h9.78a2 2 0 001.95-1.57l1.65-7.43H5.12"/></svg>
   190|              Carrito
   191|            </NuxtLink>
   192|            <NuxtLink
   193|              to="/auth/login"
   194|              class="flex items-center gap-2 px-4 py-2.5 rounded-lg bg-surface hover:bg-border-base transition-colors font-body text-small font-semibold text-text-primary min-h-[44px]"
   195|              @click="mobileMenuOpen = false"
   196|            >
   197|              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21v-2a4 4 0 00-4-4H9a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
   198|              Iniciar Sesión
   199|            </NuxtLink>
   200|          </div>
   201|        </div>
   202|      </Transition>
   203|    </div>
   204|  </header>
   205|</template>
   206|
   207|<style scoped>
   208|.mega-enter-active,
   209|.mega-leave-active {
   210|  transition: opacity 200ms ease, transform 200ms ease;
   211|}
   212|.mega-enter-from,
   213|.mega-leave-to {
   214|  opacity: 0;
   215|  transform: translateY(-8px);
   216|}
   217|
   218|.slide-enter-active,
   219|.slide-leave-active {
   220|  transition: all 300ms ease;
   221|}
   222|.slide-enter-from,
   223|.slide-leave-to {
   224|  opacity: 0;
   225|  max-height: 0;
   226|}
   227|.slide-enter-to,
   228|.slide-leave-from {
   229|  opacity: 1;
   230|  max-height: 500px;
   231|}
   232|</style>
   233|