     1|<script setup lang="ts">
     2|// Layout admin — protegido por middleware/admin.ts
     3|</script>
     4|
     5|<template>
     6|  <div class="min-h-screen flex bg-surface">
     7|    <!-- Sidebar admin -->
     8|    <aside class="w-64 bg-text-primary text-white p-6 hidden lg:block">
     9|      <span class="font-display text-display-md text-brand-primary leading-none">
    10|        DEPRO<span class="text-white">DECA</span>
    11|      </span>
    12|      <p class="mt-1 font-body text-caption text-gray-400">Panel de Administración</p>
    13|
    14|      <nav class="mt-8 space-y-2">
    15|        <NuxtLink
    16|          to="/admin/dashboard"
    17|          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
    18|        >
    19|          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="9" x="3" y="3" rx="1"/><rect width="7" height="5" x="14" y="3" rx="1"/><rect width="7" height="9" x="14" y="12" rx="1"/><rect width="7" height="5" x="3" y="16" rx="1"/></svg>
    20|          Dashboard
    21|        </NuxtLink>
    22|        <NuxtLink
    23|          to="/admin/productos"
    24|          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
    25|        >
    26|          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 4.27 9 5.15"/><path d="M21 8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16Z"/><path d="m3.3 7 8.7 5 8.7-5"/><path d="M12 22V12"/></svg>
    27|          Productos
    28|        </NuxtLink>
    29|        <NuxtLink
    30|          to="/admin/pedidos"
    31|          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
    32|        >
    33|          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M6 2 3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4Z"/><path d="M3 6h18"/><path d="M16 10a4 4 0 01-8 0"/></svg>
    34|          Pedidos
    35|        </NuxtLink>
    36|      </nav>
    37|    </aside>
    38|
    39|    <!-- Contenido admin -->
    40|    <div class="flex-1 p-6">
    41|      <slot />
    42|    </div>
    43|  </div>
    44|</template>
    45|