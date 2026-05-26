     1|<script setup lang="ts">
     2|const año = new Date().getFullYear()
     3|</script>
     4|
     5|<template>
     6|  <footer class="bg-text-primary text-white mt-16">
     7|    <div class="max-w-[1280px] mx-auto px-4 md:px-6 py-16">
     8|      <div class="grid grid-cols-1 md:grid-cols-4 gap-10">
     9|
    10|        <!-- ─── Columna 1: Marca ────────────────────── -->
    11|        <div class="md:col-span-1">
    12|          <span class="font-display text-display-md text-brand-primary leading-none">
    13|            DEPRO<span class="text-white">DECA</span>
    14|          </span>
    15|          <p class="mt-3 font-body text-small text-gray-400 leading-relaxed">
    16|            Tienda online B2B para proveedores de bodega. Filial de Nestlé Perú.
    17|          </p>
    18|        </div>
    19|
    20|        <!-- ─── Columna 2: Enlaces ──────────────────── -->
    21|        <div>
    22|          <h4 class="font-body text-small font-bold text-white mb-4 uppercase tracking-wide">Navegación</h4>
    23|          <ul class="space-y-2.5">
    24|            <li><NuxtLink to="/catalogo" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Catálogo</NuxtLink></li>
    25|            <li><NuxtLink to="/catalogo?categoria=bebidas" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Bebidas</NuxtLink></li>
    26|            <li><NuxtLink to="/catalogo?categoria=snacks" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Snacks</NuxtLink></li>
    27|            <li><NuxtLink to="/catalogo?categoria=abarrotes" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Abarrotes</NuxtLink></li>
    28|          </ul>
    29|        </div>
    30|
    31|        <!-- ─── Columna 3: Soporte ──────────────────── -->
    32|        <div>
    33|          <h4 class="font-body text-small font-bold text-white mb-4 uppercase tracking-wide">Soporte</h4>
    34|          <ul class="space-y-2.5">
    35|            <li><NuxtLink to="/auth/login" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Mi Cuenta</NuxtLink></li>
    36|            <li><NuxtLink to="/pedidos" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Mis Pedidos</NuxtLink></li>
    37|            <li><a href="mailto:soporte@deprodeca.pe" class="font-body text-small text-gray-400 hover:text-brand-primary transition-colors duration-300">Contacto</a></li>
    38|            <li><span class="font-body text-small text-gray-400">Lima, Perú</span></li>
    39|          </ul>
    40|        </div>
    41|
    42|        <!-- ─── Columna 4: Gamificación ─────────────── -->
    43|        <div>
    44|          <h4 class="font-body text-small font-bold text-white mb-4 uppercase tracking-wide">Programa de Rangos</h4>
    45|          <ul class="space-y-2">
    46|            <li class="flex items-center gap-2 font-body text-small text-gray-400">
    47|              <span class="w-2 h-2 rounded-full bg-amber-600" /> Bronce
    48|            </li>
    49|            <li class="flex items-center gap-2 font-body text-small text-gray-400">
    50|              <span class="w-2 h-2 rounded-full bg-gray-400" /> Plata
    51|            </li>
    52|            <li class="flex items-center gap-2 font-body text-small text-gray-400">
    53|              <span class="w-2 h-2 rounded-full bg-yellow-500" /> Oro
    54|            </li>
    55|            <li class="flex items-center gap-2 font-body text-small text-gray-400">
    56|              <span class="w-2 h-2 rounded-full bg-purple-500" /> Platino
    57|            </li>
    58|          </ul>
    59|        </div>
    60|      </div>
    61|
    62|      <!-- ─── Línea inferior ──────────────────────────── -->
    63|      <div class="mt-12 pt-6 border-t border-gray-800 flex flex-col md:flex-row items-center justify-between gap-4">
    64|        <p class="font-body text-caption text-gray-500">
    65|          &copy; {{ año }} DEPRODECA — Nestlé Perú. Todos los derechos reservados.
    66|        </p>
    67|        <p class="font-body text-caption text-gray-500">
    68|          Hecho con ❤️ en Ica, Perú
    69|        </p>
    70|      </div>
    71|    </div>
    72|  </footer>
    73|</template>
    74|