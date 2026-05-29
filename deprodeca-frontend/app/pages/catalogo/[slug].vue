<script setup lang="ts">
import { resolverImagenProducto } from "~/shared/utils/imagenes-productos"

definePageMeta({ layout: "default" })

const route = useRoute()
const config = useRuntimeConfig()

const slug = computed(() => route.params.slug as string)
const producto = ref<any>(null)
const loading = ref(true)
const cantidad = ref(1)
const agregando = ref(false)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)

const imagenProducto = computed(() => {
  if (!producto.value) return ""
  return resolverImagenProducto(producto.value, 0, new Set<string>())
})

watch(slug, () => cargarProducto())
onMounted(() => cargarProducto())

async function cargarProducto() {
  loading.value = true
  cantidad.value = 1
  try {
    const res = await $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/productos/${slug.value}`)
    producto.value = res.success ? res.data : null
  } catch {
    producto.value = null
  } finally {
    loading.value = false
  }
}

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
    minimumFractionDigits: 2,
  }).format(precio)
}

function agregarAlCarrito() {
  if (!producto.value || agregando.value) return
  agregando.value = true

  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    const indice = carrito.findIndex((item: any) => item.id === producto.value.id)

    if (indice >= 0) carrito[indice].cantidad += cantidad.value
    else {
      carrito.push({
        id: producto.value.id,
        nombre: producto.value.nombre,
        slug: producto.value.slug,
        precio: producto.value.precio,
        imagen_url: imagenProducto.value,
        unidad: producto.value.unidad,
        cantidad: cantidad.value,
      })
    }

    localStorage.setItem("deprodeca_carrito", JSON.stringify(carrito))
    window.dispatchEvent(new Event("carrito-actualizado"))
  }

  toastMsg.value = `${cantidad.value} x ${producto.value.nombre} agregado al carrito`
  toastType.value = "success"
  showToast.value = true
  window.setTimeout(() => (showToast.value = false), 3000)
  agregando.value = false
}
</script>

<template>
  <main class="max-w-[1440px] mx-auto px-6 md:px-8 py-10 md:py-14">
    <div v-if="loading" class="flex min-h-[60vh] items-center justify-center">
      <div class="flex flex-col items-center gap-4">
        <div class="h-8 w-8 animate-spin border-2 border-negro border-t-transparent" />
        <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-stone-oscuro">Cargando producto</span>
      </div>
    </div>

    <div v-else-if="!producto" class="border border-stone bg-blanco px-6 py-24 text-center">
      <p class="mx-auto max-w-[14ch] font-display text-[clamp(2.5rem,7vw,6rem)] font-black uppercase leading-[0.9] text-negro">Producto no encontrado</p>
      <button class="mt-8 border border-negro bg-negro px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-blanco transition-colors duration-300 hover:bg-dorado hover:text-negro" @click="navigateTo('/catalogo')">
        Volver al catalogo
      </button>
    </div>

    <template v-else>
      <div class="grid grid-cols-1 lg:grid-cols-2">
        <section class="relative">
          <div class="relative aspect-square bg-crema">
            <img :src="imagenProducto" :alt="producto.nombre" class="block h-full w-full object-cover" loading="eager" />
            <div v-if="producto.stock === 0" class="absolute inset-0 bg-black/60 flex items-center justify-center">
              <span class="font-display text-6xl font-black text-blanco uppercase">Agotado</span>
            </div>
          </div>
        </section>

        <section class="pt-8 lg:pt-0 lg:pl-16 lg:py-12">
          <nav class="mb-6 flex flex-wrap items-center gap-2 font-mono text-[11px] uppercase tracking-[0.2em] text-stone-oscuro">
            <NuxtLink to="/catalogo" class="hover:text-dorado transition-colors">Catalogo</NuxtLink>
            <span>/</span>
            <span>{{ producto.categoria_nombre || "General" }}</span>
          </nav>

          <span class="inline-block border border-dorado text-dorado px-4 py-1.5 font-mono text-[10px] uppercase tracking-[0.2em] mb-6">{{ producto.categoria_nombre || "General" }}</span>

          <h1 class="font-display text-[clamp(3rem,6vw,5rem)] font-black text-negro uppercase leading-[0.9] tracking-[-0.02em]">{{ producto.nombre }}</h1>

          <p class="font-body text-[17px] text-stone-oscuro leading-relaxed mt-6 max-w-[55ch]">{{ producto.descripcion || "Producto para bodega con presentacion premium y disponibilidad sujeta a stock." }}</p>

          <div class="border-t border-stone pt-8 mt-8">
            <div class="flex items-end gap-3">
              <span class="font-display text-[clamp(3.5rem,7vw,6rem)] font-black text-dorado leading-none">{{ formatearPrecio(producto.precio) }}</span>
              <span class="font-mono text-[12px] text-stone-oscuro uppercase tracking-[0.2em] pb-2">/{{ producto.unidad }}</span>
            </div>
          </div>

          <div class="mt-8 flex items-center gap-2">
            <span class="w-3 h-3" :class="producto.stock > 0 ? 'bg-exito' : 'bg-error'" />
            <span class="font-mono text-[11px] uppercase tracking-[0.2em]" :class="producto.stock > 0 ? 'text-exito' : 'text-error'">
              {{ producto.stock > 0 ? `${producto.stock} disponibles` : "Sin inventario" }}
            </span>
          </div>

          <div v-if="producto.stock > 0" class="mt-10 flex flex-col gap-4 md:flex-row md:items-stretch">
            <div class="flex w-fit border border-stone bg-blanco">
              <button class="flex min-h-[60px] min-w-[60px] items-center justify-center px-4 font-mono text-[16px] text-negro transition-colors duration-300 hover:bg-crema disabled:opacity-30" :disabled="cantidad <= 1" @click="cantidad--">-</button>
              <span class="flex min-h-[60px] min-w-[80px] items-center justify-center border-x border-stone px-5 font-mono text-[16px] text-negro">{{ cantidad }}</span>
              <button class="flex min-h-[60px] min-w-[60px] items-center justify-center px-4 font-mono text-[16px] text-negro transition-colors duration-300 hover:bg-crema disabled:opacity-30" :disabled="cantidad >= producto.stock" @click="cantidad++">+</button>
            </div>

            <button class="min-h-[60px] bg-dorado text-negro font-display text-lg font-bold uppercase tracking-[0.1em] py-5 px-12 w-full md:w-auto transition-colors duration-300 hover:bg-negro hover:text-blanco disabled:opacity-50" :disabled="agregando" @click="agregarAlCarrito">
              {{ agregando ? "Agregando" : "Agregar" }}
            </button>
          </div>

          <button v-else disabled class="mt-10 min-h-[60px] bg-stone text-stone-oscuro font-display text-lg font-bold uppercase tracking-[0.1em] py-5 px-12 w-full md:w-auto">
            Agotado
          </button>
        </section>
      </div>
    </template>

    <Toast :message="toastMsg" :type="toastType" :show="showToast" @close="showToast = false" />
  </main>
</template>
