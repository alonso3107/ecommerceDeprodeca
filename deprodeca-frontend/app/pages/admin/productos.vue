<!--
  ═══════════════════════════════════════════════════════════════
  admin/productos.vue — Gestion de Productos · DEPRODECA
  CRUD completo conectado a la API. Modal brutalista para
  crear/editar. Dropdown de categorias desde backend.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "admim", middleware: "admin" })

const config = useRuntimeConfig()
const productos = ref<any[]>([])
const categorias = ref<any[]>([])
const loading = ref(true)
const busqueda = ref("")
const categoriaFiltro = ref("")
const pagina = ref(1)
const totalPaginas = ref(1)

// ─── Modal ──────────────────────────────────────────────
const modalAbierto = ref(false)
const editandoId = ref<number | null>(null)
const saving = ref(false)
const form = ref({ nombre: "", categoria_id: 0, precio: 0, stock: 0, unidad: "", descripcion: "", imagen_url: "" })
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

function getToken() {
  return import.meta.client ? localStorage.getItem("deprodeca_token") : null
}

// ─── Cargar datos ───────────────────────────────────────
async function cargarProductos() {
  loading.value = true
  const token = getToken()
  if (!token) return
  try {
    const params = new URLSearchParams({ pagina: String(pagina.value), limite: "20" })
    if (busqueda.value) params.set("q", busqueda.value)
    if (categoriaFiltro.value) params.set("categoria_id", categoriaFiltro.value)

    const res = await $fetch<{ success: boolean; data: any[]; total?: number }>(
      `${config.public.apiBase}/productos?${params}`,
      { headers: { Authorization: `Bearer ${token}` } },
    )
    if (res.success) {
      productos.value = res.data
      totalPaginas.value = Math.ceil((res.total || res.data.length) / 20)
    }
  } catch (_) { /* fallback */ }
  finally { loading.value = false }
}

async function cargarCategorias() {
  try {
    const res = await $fetch<{ success: boolean; data: any[] }>(
      `${config.public.apiBase}/categorias`,
    )
    if (res.success) categorias.value = res.data
  } catch (_) { /* fallback */ }
}

// ─── Modal ──────────────────────────────────────────────
function abrirNuevo() {
  editandoId.value = null
  form.value = { nombre: "", categoria_id: 0, precio: 0, stock: 0, unidad: "", descripcion: "", imagen_url: "" }
  modalAbierto.value = true
}

async function abrirEditar(p: any) {
  editandoId.value = p.id
  form.value = {
    nombre: p.nombre || "",
    categoria_id: p.categoria_id || 0,
    precio: p.precio || 0,
    stock: p.stock || 0,
    unidad: p.unidad || "",
    descripcion: p.descripcion || "",
    imagen_url: p.imagen_url || "",
  }
  modalAbierto.value = true
}

function cerrarModal() {
  modalAbierto.value = false
  editandoId.value = null
}

async function guardarProducto() {
  saving.value = true
  const token = getToken()
  try {
    const body = {
      nombre: form.value.nombre,
      categoria_id: form.value.categoria_id,
      precio: Number(form.value.precio),
      stock: Number(form.value.stock),
      unidad: form.value.unidad,
      descripcion: form.value.descripcion,
      imagen_url: form.value.imagen_url,
    }

    if (editandoId.value) {
      await $fetch(`${config.public.apiBase}/admin/productos/${editandoId.value}`, {
        method: "PUT",
        headers: { Authorization: `Bearer ${token}`, "Content-Type": "application/json" },
        body,
      })
      mostrarToast("Producto actualizado", "success")
    } else {
      await $fetch(`${config.public.apiBase}/admin/productos`, {
        method: "POST",
        headers: { Authorization: `Bearer ${token}`, "Content-Type": "application/json" },
        body,
      })
      mostrarToast("Producto creado", "success")
    }
    cerrarModal()
    await cargarProductos()
  } catch (e: any) {
    mostrarToast(e?.data?.message || "Error al guardar producto", "error")
  } finally { saving.value = false }
}

async function eliminarProducto(p: any) {
  if (!confirm(`Eliminar "${p.nombre}"? Esta accion no se puede deshacer.`)) return
  const token = getToken()
  try {
    await $fetch(`${config.public.apiBase}/admin/productos/${p.id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${token}` },
    })
    mostrarToast("Producto eliminado", "success")
    await cargarProductos()
  } catch (e: any) {
    mostrarToast(e?.data?.message || "Error al eliminar", "error")
  }
}

function mostrarToast(msg: string, type: "success" | "error") {
  toastMsg.value = msg
  toastType.value = type
  showToast.value = true
  setTimeout(() => (showToast.value = false), 3500)
}

// ─── Utilidades ────────────────────────────────────────
function formatearPrecio(p: number) {
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(p)
}

function stockClass(s: number) {
  if (s > 50) return "bg-exito"
  if (s > 10) return "bg-advertencia"
  return "bg-error"
}

function nombreCategoria(id: number) {
  return categorias.value.find((c: any) => c.id === id)?.nombre || "—"
}

// ─── Lifecycle ──────────────────────────────────────────
onMounted(async () => {
  await cargarCategorias()
  await cargarProductos()
})

watch([busqueda, categoriaFiltro], () => {
  pagina.value = 1
  cargarProductos()
})
</script>

<template>
  <div class="page-enter">

    <!-- Encabezado -->
    <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 mb-8">
      <div>
        <p class="font-mono text-xs text-texto-muted uppercase tracking-[0.3em] mb-2">
          ─── Admin
        </p>
        <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
          Productos<span class="text-[#D4A017]">.</span>
        </h1>
      </div>

      <button
        class="bg-texto text-white font-mono text-xs uppercase tracking-[0.12em] px-5 py-3.5
               hover:bg-[#D4A017] hover:text-black transition-colors duration-200
               min-h-[48px] flex items-center gap-2 self-start cursor-pointer"
        @click="abrirNuevo"
      >
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect x="7" width="2" height="16" fill="currentColor"/>
          <rect y="7" width="16" height="2" fill="currentColor"/>
        </svg>
        Nuevo Producto
      </button>
    </div>

    <!-- Barra de herramientas -->
    <div class="flex flex-col sm:flex-row gap-0 mb-0">
      <div class="flex flex-1 gap-0">
        <input
          v-model="busqueda"
          type="search"
          placeholder="Buscar producto..."
          class="flex-1 border border-borde px-4 py-3 font-body text-sm text-texto
                 bg-white placeholder:text-texto-muted
                 focus:border-[#D4A017] focus:outline-none
                 transition-colors duration-200 min-h-[48px]"
        />
        <button
          class="border border-l-0 border-borde bg-texto text-white px-6 py-3
                 font-mono text-xs uppercase tracking-[0.12em]
                 hover:bg-[#D4A017] hover:text-black transition-colors min-h-[48px] cursor-pointer"
          @click="cargarProductos"
        >
          Filtrar
        </button>
      </div>

      <select
        v-model="categoriaFiltro"
        class="border border-l-0 sm:border-l-0 border-t-0 sm:border-t border-borde px-4 py-3
               font-mono text-xs uppercase tracking-[0.1em]
               text-texto bg-white focus:border-[#D4A017] focus:outline-none
               min-h-[48px] cursor-pointer"
      >
        <option value="">Todas las categorias</option>
        <option v-for="cat in categorias" :key="cat.id" :value="cat.id">
          {{ cat.nombre }}
        </option>
      </select>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-24 gap-3">
      <div class="w-10 h-10 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-sm text-texto-muted uppercase tracking-[0.15em]">Cargando productos...</span>
    </div>

    <!-- Tabla -->
    <div v-else class="border border-borde border-t-0">
      <!-- Header -->
      <div class="hidden md:grid grid-cols-[60px_1.5fr_1fr_120px_100px_90px] gap-4 px-6 py-4
                  bg-texto text-white font-mono text-xs uppercase tracking-[0.15em]">
        <span>ID</span>
        <span>Producto</span>
        <span>Categoria</span>
        <span>Precio</span>
        <span>Stock</span>
        <span>Acciones</span>
      </div>

      <!-- Vacio -->
      <div v-if="productos.length === 0" class="px-6 py-20 text-center border-t border-borde">
        <p class="font-display text-display-md text-texto/10 uppercase mb-3">Sin productos</p>
        <p class="font-body text-sm text-texto-muted">
          No hay productos registrados. Crea el primero con "Nuevo Producto".
        </p>
      </div>

      <!-- Filas -->
      <div
        v-for="(p, i) in productos"
        :key="p.id"
        class="grid grid-cols-1 md:grid-cols-[60px_1.5fr_1fr_120px_100px_90px] gap-3 md:gap-4 px-6 py-4
               items-center border-t border-borde transition-colors duration-200"
        :class="i % 2 === 1 ? 'bg-fondo/50' : 'bg-white'"
      >
        <!-- ID -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">ID</span>
          <p class="font-mono text-sm text-texto-muted tracking-wider">#{{ p.id }}</p>
        </div>

        <!-- Nombre -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Producto</span>
          <p class="font-body text-sm text-texto font-medium truncate">{{ p.nombre }}</p>
          <p class="font-mono text-xs text-texto-muted uppercase tracking-wider mt-0.5">/ {{ p.unidad || 'unidad' }}</p>
        </div>

        <!-- Categoria -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Categoria</span>
          <span class="font-mono text-xs text-texto-muted uppercase tracking-[0.1em]">
            {{ nombreCategoria(p.categoria_id) }}
          </span>
        </div>

        <!-- Precio -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Precio</span>
          <p class="font-body text-sm font-bold text-texto">{{ formatearPrecio(p.precio) }}</p>
        </div>

        <!-- Stock -->
        <div>
          <span class="font-mono text-xs text-texto-muted md:hidden uppercase tracking-wider">Stock</span>
          <div class="flex items-center gap-2">
            <span class="w-2.5 h-2.5" :class="stockClass(p.stock)" />
            <span class="font-mono text-sm text-texto font-bold">{{ p.stock }}</span>
          </div>
        </div>

        <!-- Acciones -->
        <div class="flex gap-0">
          <button
            class="border border-borde px-3 py-2 font-mono text-xs uppercase tracking-[0.1em]
                   text-texto-muted hover:text-texto hover:border-texto/30 transition-colors cursor-pointer"
            title="Editar"
            @click="abrirEditar(p)"
          >
            <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
              <path d="M10 1L13 4L5 12H2V9L10 1Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="miter"/>
            </svg>
          </button>
          <button
            class="border border-l-0 border-borde px-3 py-2 font-mono text-xs uppercase tracking-[0.1em]
                   text-texto-muted hover:text-error hover:border-error/30 transition-colors cursor-pointer"
            title="Eliminar"
            @click="eliminarProducto(p)"
          >
            <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
              <path d="M3.5 3.5L10.5 10.5M10.5 3.5L3.5 10.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Paginacion -->
    <div class="flex justify-between items-center mt-6">
      <span class="font-mono text-xs text-texto-muted uppercase tracking-[0.12em]">
        {{ productos.length }} productos
      </span>
      <div class="flex gap-0">
        <button
          class="border border-borde px-4 py-2.5 font-mono text-xs uppercase tracking-[0.12em]
                 text-texto-muted hover:text-texto hover:border-texto/30 transition-colors min-h-[44px] cursor-pointer
                 disabled:opacity-30"
          :disabled="pagina <= 1"
          @click="pagina--; cargarProductos()"
        >
          Anterior
        </button>
        <span class="border-t border-b border-borde px-4 py-2.5 font-mono text-xs text-texto
                     uppercase tracking-[0.12em] font-bold min-h-[44px] flex items-center">
          {{ pagina }} / {{ totalPaginas }}
        </span>
        <button
          class="border border-borde px-4 py-2.5 font-mono text-xs uppercase tracking-[0.12em]
                 text-texto-muted hover:text-texto hover:border-texto/30 transition-colors min-h-[44px] cursor-pointer
                 disabled:opacity-30"
          :disabled="pagina >= totalPaginas"
          @click="pagina++; cargarProductos()"
        >
          Siguiente
        </button>
      </div>
    </div>

    <!-- ═══ MODAL: Nuevo/Editar Producto ═══ -->
    <Teleport to="body">
      <Transition name="modal">
        <div
          v-if="modalAbierto"
          class="fixed inset-0 z-50 flex items-center justify-center p-4"
        >
          <!-- Overlay -->
          <div class="absolute inset-0 bg-black/50" @click="cerrarModal" />

          <!-- Panel -->
          <div class="relative bg-white border border-borde w-full max-w-lg max-h-[90vh] overflow-y-auto">
            <!-- Header -->
            <div class="px-6 py-4 bg-texto text-white flex items-center justify-between sticky top-0">
              <h2 class="font-mono text-sm uppercase tracking-[0.15em] font-bold">
                {{ editandoId ? 'Editar Producto' : 'Nuevo Producto' }}
              </h2>
              <button class="text-white/50 hover:text-white transition-colors cursor-pointer" @click="cerrarModal">
                <svg width="18" height="18" viewBox="0 0 18 18" fill="none">
                  <path d="M3 3L15 15M15 3L3 15" stroke="currentColor" stroke-width="2" stroke-linecap="square"/>
                </svg>
              </button>
            </div>

            <!-- Formulario -->
            <form class="p-6 space-y-5" @submit.prevent="guardarProducto">
              <!-- Nombre -->
              <div>
                <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Nombre *</label>
                <input
                  v-model="form.nombre"
                  required
                  class="w-full border border-borde px-4 py-3 font-body text-sm text-texto
                         bg-white placeholder:text-texto-muted
                         focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]"
                  placeholder="Ej: Milo 400g"
                />
              </div>

              <!-- Categoria -->
              <div>
                <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Categoria *</label>
                <select
                  v-model="form.categoria_id"
                  required
                  class="w-full border border-borde px-4 py-3 font-mono text-xs uppercase tracking-[0.1em]
                         text-texto bg-white focus:border-[#D4A017] focus:outline-none
                         transition-colors min-h-[48px] cursor-pointer"
                >
                  <option :value="0" disabled>Seleccionar categoria</option>
                  <option v-for="cat in categorias" :key="cat.id" :value="cat.id">
                    {{ cat.nombre }}
                  </option>
                </select>
              </div>

              <!-- Precio + Stock -->
              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Precio (S/) *</label>
                  <input
                    v-model.number="form.precio"
                    type="number" step="0.01" min="0" required
                    class="w-full border border-borde px-4 py-3 font-mono text-sm text-texto
                           bg-white placeholder:text-texto-muted
                           focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]"
                    placeholder="12.50"
                  />
                </div>
                <div>
                  <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Stock *</label>
                  <input
                    v-model.number="form.stock"
                    type="number" min="0" required
                    class="w-full border border-borde px-4 py-3 font-mono text-sm text-texto
                           bg-white placeholder:text-texto-muted
                           focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]"
                    placeholder="100"
                  />
                </div>
              </div>

              <!-- Unidad + Imagen URL -->
              <div class="grid grid-cols-2 gap-3">
                <div>
                  <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Unidad</label>
                  <input
                    v-model="form.unidad"
                    class="w-full border border-borde px-4 py-3 font-body text-sm text-texto
                           bg-white placeholder:text-texto-muted
                           focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]"
                    placeholder="unidad / pack / caja"
                  />
                </div>
                <div>
                  <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Imagen URL</label>
                  <input
                    v-model="form.imagen_url"
                    class="w-full border border-borde px-4 py-3 font-body text-sm text-texto
                           bg-white placeholder:text-texto-muted
                           focus:border-[#D4A017] focus:outline-none transition-colors min-h-[48px]"
                    placeholder="https://..."
                  />
                </div>
              </div>

              <!-- Descripcion -->
              <div>
                <label class="block font-mono text-xs text-texto-muted uppercase tracking-[0.12em] mb-2">Descripcion</label>
                <textarea
                  v-model="form.descripcion"
                  rows="3"
                  class="w-full border border-borde px-4 py-3 font-body text-sm text-texto
                         bg-white placeholder:text-texto-muted resize-none
                         focus:border-[#D4A017] focus:outline-none transition-colors"
                  placeholder="Descripcion breve del producto..."
                />
              </div>

              <!-- Acciones -->
              <div class="flex gap-0 pt-2">
                <button
                  type="submit"
                  :disabled="saving"
                  class="flex-1 bg-texto text-white font-mono text-xs uppercase tracking-[0.12em]
                         px-6 py-3.5 hover:bg-[#D4A017] hover:text-black
                         transition-colors duration-200 disabled:opacity-50
                         min-h-[52px] flex items-center justify-center gap-2 cursor-pointer"
                >
                  <span v-if="saving" class="w-4 h-4 border-2 border-white border-t-transparent animate-spin" />
                  <span v-else>{{ editandoId ? 'Guardar Cambios' : 'Crear Producto' }}</span>
                </button>
                <button
                  type="button"
                  class="border border-l-0 border-borde font-mono text-xs uppercase tracking-[0.12em]
                         px-6 py-3.5 text-texto-muted hover:text-texto hover:bg-fondo
                         transition-colors duration-200 min-h-[52px] cursor-pointer"
                  @click="cerrarModal"
                >
                  Cancelar
                </button>
              </div>
            </form>
          </div>
        </div>
      </Transition>
    </Teleport>

    <!-- Toast -->
    <Toast
      :message="toastMsg"
      :type="toastType"
      :show="showToast"
      @close="showToast = false"
    />
  </div>
</template>

<style scoped>
.modal-enter-active { transition: all 250ms ease-out; }
.modal-leave-active { transition: all 150ms ease-in; }
.modal-enter-from { opacity: 0; }
.modal-enter-from > div:last-child { transform: translateY(20px); opacity: 0; }
.modal-leave-to { opacity: 0; }
.modal-leave-to > div:last-child { transform: translateY(20px); opacity: 0; }
</style>
