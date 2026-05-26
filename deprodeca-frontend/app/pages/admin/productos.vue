<!--
  ═══════════════════════════════════════════════════════════════
  admin/productos.vue — Gestión de Productos · DEPRODECA
  Brutalismo aplicado a CRUD: tabla de productos con acciones
  como bloques, búsqueda con input rectangular, paginación dura.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "admim", middleware: "admin" })

// ─── Estado ───────────────────────────────────────────────
const busqueda = ref("")
const categoriaFiltro = ref("")

// ─── Datos mock ───────────────────────────────────────────
const productos = [
  { id: 1, nombre: "Milo 400g", categoria: "Bebidas", precio: 12.50, stock: 250, unidad: "unidad" },
  { id: 2, nombre: "Nescafé Tradición 200g", categoria: "Bebidas", precio: 18.90, stock: 180, unidad: "unidad" },
  { id: 3, nombre: "Sublime Clásico 30g", categoria: "Snacks", precio: 1.50, stock: 500, unidad: "unidad" },
  { id: 4, nombre: "Galleta Morocha 6pk", categoria: "Snacks", precio: 4.20, stock: 320, unidad: "pack" },
  { id: 5, nombre: "Leche Ideal 400ml", categoria: "Lácteos", precio: 5.80, stock: 150, unidad: "unidad" },
  { id: 6, nombre: "Aceite Capri 1L", categoria: "Abarrotes", precio: 9.90, stock: 400, unidad: "unidad" },
]

function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(p)
}
</script>

<template>
  <div class="page-enter">

    <!-- Encabezado -->
    <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 mb-10">
      <div>
        <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
          ─── Admin
        </p>
        <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
          Productos<span class="text-[#D4A017]">.</span>
        </h1>
      </div>

      <!-- Botón Nuevo Producto -->
      <button
        class="bg-texto text-white font-mono text-[10px] uppercase tracking-[0.15em] px-5 py-3
               hover:bg-[#D4A017] hover:text-black transition-colors duration-200
               min-h-[48px] flex items-center gap-2 self-start"
      >
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
          <rect x="6" width="2" height="14" fill="currentColor"/>
          <rect y="6" width="14" height="2" fill="currentColor"/>
        </svg>
        Nuevo Producto
      </button>
    </div>

    <!-- ═══ BARRA DE HERRAMIENTAS ═══ -->
    <div class="flex flex-col sm:flex-row gap-0 mb-0">
      <!-- Búsqueda -->
      <div class="flex flex-1 gap-0">
        <input
          v-model="busqueda"
          type="search"
          placeholder="Buscar producto..."
          class="flex-1 border border-borde px-4 py-3 font-body text-body text-texto
                 bg-white placeholder:text-texto-muted
                 focus:border-[#D4A017] focus:outline-none
                 transition-colors duration-200 min-h-[48px]"
        />
        <button
          class="border border-l-0 border-borde bg-texto text-white px-6 py-3
                 font-mono text-[10px] uppercase tracking-[0.15em]
                 hover:bg-[#D4A017] hover:text-black transition-colors min-h-[48px]"
        >
          Buscar
        </button>
      </div>

      <!-- Filtro categoría -->
      <select
        v-model="categoriaFiltro"
        class="border border-l-0 border-borde px-4 py-3 font-mono text-[10px] uppercase tracking-[0.15em]
               text-texto bg-white focus:border-[#D4A017] focus:outline-none
               min-h-[48px]"
      >
        <option value="">Todas</option>
        <option value="bebidas">Bebidas</option>
        <option value="snacks">Snacks</option>
        <option value="lacteos">Lácteos</option>
        <option value="abarrotes">Abarrotes</option>
        <option value="limpieza">Limpieza</option>
      </select>
    </div>

    <!-- ═══ TABLA DE PRODUCTOS ═══ -->
    <div class="border border-borde border-t-0">
      <!-- Header -->
      <div class="grid grid-cols-[60px_1.5fr_1fr_120px_100px_100px] gap-4 px-6 py-4
                  bg-texto text-white font-mono text-[9px] uppercase tracking-[0.2em]">
        <span>ID</span>
        <span>Producto</span>
        <span>Categoría</span>
        <span>Precio</span>
        <span>Stock</span>
        <span>Acciones</span>
      </div>

      <!-- Filas -->
      <div
        v-for="(p, i) in productos"
        :key="p.id"
        class="grid grid-cols-[60px_1.5fr_1fr_120px_100px_100px] gap-4 px-6 py-4 items-center border-t border-borde
               hover:bg-[#D4A017]/[0.02] transition-colors duration-200"
        :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
      >
        <!-- ID -->
        <span class="font-mono text-[11px] text-texto-muted tracking-wider">#{{ p.id }}</span>

        <!-- Nombre + unidad -->
        <div>
          <p class="font-body text-small font-bold text-texto truncate">{{ p.nombre }}</p>
          <p class="font-mono text-[9px] text-texto-muted uppercase tracking-wider mt-0.5">/ {{ p.unidad }}</p>
        </div>

        <!-- Categoría -->
        <span class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.15em]">{{ p.categoria }}</span>

        <!-- Precio -->
        <span class="font-display text-heading text-texto">{{ formatearPrecio(p.precio) }}</span>

        <!-- Stock · Indicador geométrico -->
        <div class="flex items-center gap-2">
          <span class="w-2 h-2" :class="p.stock > 50 ? 'bg-exito' : p.stock > 10 ? 'bg-advertencia' : 'bg-error'" />
          <span class="font-mono text-[11px] text-texto font-bold">{{ p.stock }}</span>
        </div>

        <!-- Acciones · Bloques duros -->
        <div class="flex gap-0">
          <button
            class="border border-borde px-2.5 py-1.5 font-mono text-[9px] uppercase tracking-[0.1em]
                   text-texto-muted hover:text-texto hover:border-texto/30 transition-colors"
            title="Editar"
          >
            <!-- Ícono: Lápiz geométrico -->
            <svg width="12" height="12" viewBox="0 0 12 12" fill="none">
              <path d="M9 1L11 3L4 10H2V8L9 1Z" stroke="currentColor" stroke-width="1.2" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
          </button>
          <button
            class="border border-l-0 border-borde px-2.5 py-1.5 font-mono text-[9px] uppercase tracking-[0.1em]
                   text-texto-muted hover:text-error hover:border-error/30 transition-colors"
            title="Eliminar"
          >
            <!-- Ícono: X en rombo -->
            <svg width="12" height="12" viewBox="0 0 12 12" fill="none">
              <path d="M3 3L9 9M9 3L3 9" stroke="currentColor" stroke-width="1.2" stroke-linecap="square"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Paginación -->
    <div class="flex justify-between items-center mt-6">
      <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.15em]">
        {{ productos.length }} productos
      </span>
      <div class="flex gap-0">
        <button class="border border-borde px-4 py-2 font-mono text-[10px] uppercase tracking-[0.15em]
                       text-texto-muted hover:text-texto hover:border-texto/30 transition-colors min-h-[40px]">
          Anterior
        </button>
        <span class="border-t border-b border-borde px-4 py-2 font-mono text-[10px] text-texto
                     uppercase tracking-[0.15em] font-bold min-h-[40px] flex items-center">
          1 / 1
        </span>
        <button class="border border-borde px-4 py-2 font-mono text-[10px] uppercase tracking-[0.15em]
                       text-texto-muted hover:text-texto hover:border-texto/30 transition-colors min-h-[40px]">
          Siguiente
        </button>
      </div>
    </div>

  </div>
</template>
