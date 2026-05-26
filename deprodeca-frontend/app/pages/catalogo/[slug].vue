<!--
  ═══════════════════════════════════════════════════════════════
  catalogo/[slug].vue — Detalle de Producto · DEPRODECA
  Brutalismo aplicado a la ficha de producto:
  • Imagen rectangular, sin curvas
  • Precio masivo como elemento estructural
  • Selector de cantidad como bloques duros
  • Botón "AGREGAR" negro sólido, sin sombras
  • Breadcrumb + metadata en monospace técnico
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })

const route = useRoute()
const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const slug = computed(() => route.params.slug as string)
const producto = ref<any>(null)
const loading = ref(true)
const cantidad = ref(1)
const agregando = ref(false)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

// ─── Ciclo de vida ────────────────────────────────────────
onMounted(async () => {
  try {
    const res = await $fetch<{ success: boolean; data: any }>(
      `${config.public.apiBase}/productos/${slug.value}`
    )
    if (res.success) producto.value = res.data
  } catch (_) {
    producto.value = null
  } finally {
    loading.value = false
  }
})

// ─── Utilidades ───────────────────────────────────────────
function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}

const PLACEHOLDER = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='600' height='600'%3E%3Crect fill='%23FAFAFA' width='600' height='600'/%3E%3Ctext x='300' y='310' text-anchor='middle' font-size='20' font-family='monospace' fill='%23A3A3A3'%3EDEPRODECA%3C/text%3E%3C/svg%3E"

// ─── Carrito ──────────────────────────────────────────────
function agregarAlCarrito() {
  if (!producto.value) return
  agregando.value = true

  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    const idx = carrito.findIndex((i: any) => i.id === producto.value.id)
    if (idx >= 0) {
      carrito[idx].cantidad += cantidad.value
    } else {
      carrito.push({
        id: producto.value.id,
        nombre: producto.value.nombre,
        slug: producto.value.slug,
        precio: producto.value.precio,
        imagen_url: producto.value.imagen_url,
        unidad: producto.value.unidad,
        cantidad: cantidad.value,
      })
    }
    localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito))
    window.dispatchEvent(new Event("carrito-actualizado"))
  }

  toastMsg.value = `${cantidad.value} × ${producto.value.nombre} agregado`
  toastType.value = "success"
  showToast.value = true
  setTimeout(() => (showToast.value = false), 3000)
  agregando.value = false
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">

    <!-- ═══ LOADING ═══ -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-32 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <!-- ═══ 404 ═══ -->
    <div v-else-if="!producto" class="text-center py-32">
      <p class="font-display text-[clamp(3rem,8vw,6rem)] text-texto/10 uppercase leading-none mb-6">
        Producto No<br />Encontrado
      </p>
      <button
        class="font-mono text-[11px] uppercase tracking-[0.2em] border border-borde px-6 py-3
               text-texto-muted hover:text-[#D4A017] hover:border-[#D4A017] transition-colors duration-200"
        @click="navigateTo('/catalogo')"
      >
        Volver al Catálogo
      </button>
    </div>

    <!-- ═══ DETALLE DEL PRODUCTO ═══ -->
    <template v-else>
      <!-- Breadcrumb · Monospace técnico -->
      <nav class="mb-10 font-mono text-[11px] uppercase tracking-[0.15em] flex items-center gap-2 flex-wrap">
        <NuxtLink to="/catalogo" class="text-texto-muted hover:text-texto transition-colors">
          Catálogo
        </NuxtLink>
        <span class="text-borde">/</span>
        <span class="text-texto-muted">{{ producto.categoria_nombre }}</span>
        <span class="text-borde">/</span>
        <span class="text-texto font-bold">{{ producto.nombre }}</span>
      </nav>

      <!-- Grid producto · 2 columnas -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-0 border border-borde">

        <!-- Imagen · Rectangular, sin curvas, sin overflow-hidden redondeado -->
        <div class="relative bg-fondo border-b md:border-b-0 md:border-r border-borde aspect-square">
          <img
            :src="producto.imagen_url || 'https://images.unsplash.com/photo-1553530979-3f62c23fb101?w=600&q=80'"
            :alt="producto.nombre"
            class="w-full h-full object-cover"
            @error="(e: Event) => (e.target as HTMLImageElement).src = PLACEHOLDER"
          />

          <!-- Badge AGOTADO · overlay brutalista -->
          <div
            v-if="producto.stock === 0"
            class="absolute inset-0 bg-black/50 flex items-center justify-center"
          >
            <span class="font-display text-display-md text-white uppercase tracking-widest">
              Agotado
            </span>
          </div>
        </div>

        <!-- Info · Columna derecha -->
        <div class="flex flex-col justify-center p-8 md:p-12">

          <!-- Categoría · Monospace -->
          <span class="font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-4">
            {{ producto.categoria_nombre }}
          </span>

          <!-- Nombre · Display masivo -->
          <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95] mb-6">
            {{ producto.nombre }}
          </h1>

          <!-- Descripción -->
          <p class="font-body text-body text-texto-muted mb-8 leading-relaxed max-w-[480px]">
            {{ producto.descripcion || "Producto de calidad para tu bodega. Distribución directa Nestlé." }}
          </p>

          <!-- Precio · MASIVO, brutalista -->
          <div class="flex items-baseline gap-3 mb-8">
            <span class="font-display text-[clamp(3rem,6vw,4.5rem)] text-texto leading-none">
              {{ formatearPrecio(producto.precio) }}
            </span>
            <span class="font-mono text-sm text-texto-muted tracking-wider">
              / {{ producto.unidad }}
            </span>
          </div>

          <!-- Stock · Indicador técnico -->
          <div class="flex items-center gap-2 mb-10">
            <span
              class="w-2.5 h-2.5"
              :class="producto.stock > 20 ? 'bg-exito' : producto.stock > 0 ? 'bg-advertencia' : 'bg-error'"
              aria-hidden="true"
            />
            <span class="font-mono text-[11px] uppercase tracking-[0.15em]"
                  :class="producto.stock > 20 ? 'text-exito' : producto.stock > 0 ? 'text-advertencia' : 'text-error'">
              {{ producto.stock > 20
                  ? `${producto.stock} en stock`
                  : producto.stock > 0
                    ? `Solo ${producto.stock} quedan`
                    : 'Agotado'
              }}
            </span>
          </div>

          <!-- Acciones · Cantidad + Botón Agregar -->
          <div v-if="producto.stock > 0" class="flex items-stretch gap-0">

            <!-- Selector de cantidad · Bloques duros, sin rounded -->
            <div class="flex border border-borde">
              <button
                class="px-5 py-4 font-mono text-lg hover:bg-fondo transition-colors
                       disabled:opacity-30 disabled:cursor-not-allowed
                       min-h-[56px] min-w-[52px] flex items-center justify-center"
                :disabled="cantidad <= 1"
                @click="cantidad--"
              >
                −
              </button>
              <span class="px-6 py-4 font-mono text-body font-bold text-texto
                           border-x border-borde min-w-[56px] flex items-center justify-center">
                {{ cantidad }}
              </span>
              <button
                class="px-5 py-4 font-mono text-lg hover:bg-fondo transition-colors
                       disabled:opacity-30 disabled:cursor-not-allowed
                       min-h-[56px] min-w-[52px] flex items-center justify-center"
                :disabled="cantidad >= producto.stock"
                @click="cantidad++"
              >
                +
              </button>
            </div>

            <!-- Botón AGREGAR · Negro sólido, sin sombras, brutalista -->
            <button
              class="flex-1 bg-texto text-white font-display text-heading uppercase tracking-[0.05em]
                     px-8 py-4 hover:bg-[#D4A017] hover:text-black
                     transition-colors duration-200
                     disabled:opacity-50 disabled:cursor-not-allowed
                     min-h-[56px] flex items-center justify-center gap-2"
              :disabled="agregando"
              @click="agregarAlCarrito"
            >
              <span v-if="agregando" class="w-5 h-5 border-2 border-white border-t-transparent animate-spin" />
              <span v-else>Agregar al Carrito</span>
            </button>

          </div>

          <!-- Botón AGOTADO -->
          <button
            v-else
            disabled
            class="bg-borde text-texto-muted font-display text-heading uppercase tracking-[0.05em]
                   px-8 py-4 min-h-[56px] flex items-center justify-center"
          >
            Agotado
          </button>

        </div>
      </div>
    </template>

    <!-- Toast -->
    <Toast
      :message="toastMsg"
      :type="toastType"
      :show="showToast"
      @close="showToast = false"
    />
  </div>
</template>
