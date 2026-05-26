<script setup lang="ts">
const año = new Date().getFullYear()
</script>

<template>
  <footer class="border-t border-borde bg-white mt-24">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8 py-20">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-12">

        <div>
          <span class="font-display text-display-md text-texto leading-none">
            depro<span class="text-texto-muted font-normal">deca</span>
          </span>
          <p class="mt-3 font-body text-small text-texto-muted leading-relaxed max-w-[240px]">
            Tienda B2B para proveedores de bodega. Filial Nestlé Perú.
          </p>
        </div>

        <div>
          <h4 class="font-body text-caption font-bold text-texto uppercase tracking-wider mb-5">Navegación</h4>
          <ul class="space-y-3">
            <li><NuxtLink to="/catalogo" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Catálogo</NuxtLink></li>
            <li><NuxtLink to="/catalogo?categoria=bebidas" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Bebidas</NuxtLink></li>
            <li><NuxtLink to="/catalogo?categoria=abarrotes" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Abarrotes</NuxtLink></li>
            <li><NuxtLink to="/catalogo?categoria=snacks" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Snacks</NuxtLink></li>
          </ul>
        </div>

        <div>
          <h4 class="font-body text-caption font-bold text-texto uppercase tracking-wider mb-5">Cuenta</h4>
          <ul class="space-y-3">
            <li><NuxtLink to="/auth/login" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Iniciar Sesión</NuxtLink></li>
            <li><NuxtLink to="/pedidos" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Mis Pedidos</NuxtLink></li>
            <li><NuxtLink to="/perfil/rangos" class="font-body text-small text-texto-muted hover:text-texto transition-colors">Mi Rango</NuxtLink></li>
          </ul>
        </div>

        <div>
          <h4 class="font-body text-caption font-bold text-texto uppercase tracking-wider mb-5">Rangos</h4>
          <ul class="space-y-2">
            <li class="flex items-center gap-2 font-body text-small text-texto-muted">
              <span class="w-2 h-2 rounded-full bg-amber-600" /> Bronce
            </li>
            <li class="flex items-center gap-2 font-body text-small text-texto-muted">
              <span class="w-2 h-2 rounded-full bg-gray-400" /> Plata
            </li>
            <li class="flex items-center gap-2 font-body text-small text-texto-muted">
              <span class="w-2 h-2 rounded-full bg-yellow-500" /> Oro
            </li>
            <li class="flex items-center gap-2 font-body text-small text-texto-muted">
              <span class="w-2 h-2 rounded-full bg-purple-500" /> Platino
            </li>
          </ul>
        </div>
      </div>

      <div class="mt-16 pt-6 border-t border-borde flex flex-col md:flex-row items-center justify-between gap-3">
        <p class="font-body text-caption text-texto-muted">&copy; {{ año }} DEPRODECA — Nestlé Perú</p>
        <p class="font-body text-caption text-texto-muted">Ica, Perú</p>
      </div>
    </div>
  </footer>
</template>
