<script setup lang="ts">
// Layout admin — protegido por middleware/admin.ts
</script>

<template>
  <div class="min-h-screen flex bg-surface">
    <!-- Sidebar admin -->
    <aside class="w-64 bg-text-primary text-white p-6 hidden lg:block">
      <span class="font-display text-display-md text-brand-primary leading-none">
        DEPRO<span class="text-white">DECA</span>
      </span>
      <p class="mt-1 font-body text-caption text-gray-400">Panel de Administración</p>

      <nav class="mt-8 space-y-2">
        <NuxtLink
          to="/admin/dashboard"
          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="9" x="3" y="3" rx="1"/><rect width="7" height="5" x="14" y="3" rx="1"/><rect width="7" height="9" x="14" y="12" rx="1"/><rect width="7" height="5" x="3" y="16" rx="1"/></svg>
          Dashboard
        </NuxtLink>
        <NuxtLink
          to="/admin/productos"
          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="m7.5 4.27 9 5.15"/><path d="M21 8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16Z"/><path d="m3.3 7 8.7 5 8.7-5"/><path d="M12 22V12"/></svg>
          Productos
        </NuxtLink>
        <NuxtLink
          to="/admin/pedidos"
          class="flex items-center gap-2 px-3 py-2.5 rounded-lg font-body text-small text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M6 2 3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4Z"/><path d="M3 6h18"/><path d="M16 10a4 4 0 01-8 0"/></svg>
          Pedidos
        </NuxtLink>
      </nav>
    </aside>

    <!-- Contenido admin -->
    <div class="flex-1 p-6">
      <slot />
    </div>
  </div>
</template>
