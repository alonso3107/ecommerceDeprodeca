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

const imagenRespaldo = computed(() => {
  const nombre = String(producto.value?.nombre || "DEPRODECA").toUpperCase()
  return `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='1200' height='1200' viewBox='0 0 1200 1200'%3E%3Crect width='1200' height='1200' fill='%23FFFFFF'/%3E%3Cpath d='M120 220H1080' stroke='%23E5E5E5' stroke-width='2'/%3E%3Cpath d='M180 320L320 180L460 320L600 180L740 320L880 180L1020 320' stroke='%23D4A017' stroke-opacity='.12' stroke-width='10' fill='none'/%3E%3Cpath d='M140 820L300 660L460 820L620 560L780 820L940 660L1100 820' stroke='%23D4A017' stroke-opacity='.12' stroke-width='10' fill='none'/%3E%3Ctext x='120' y='1040' fill='%23171717' font-family='monospace' font-size='48' letter-spacing='6'%3E${encodeURIComponent(nombre)}%3C/text%3E%3C/svg%3E`
})

function manejarErrorImagen(evento: Event) {
  const imagen = evento.target as HTMLImageElement | null
  if (imagen && imagen.src !== imagenRespaldo.value) {
    imagen.src = imagenRespaldo.value
  }
}

watch(slug, () => {
  cargarProducto()
})

onMounted(() => {
  cargarProducto()
})

async function cargarProducto() {
  loading.value = true
  cantidad.value = 1

  try {
    const res = await $fetch<{ success: boolean; data: any }>(
      `${config.public.apiBase}/productos/${slug.value}`,
    )
    producto.value = res.success ? res.data : null
  } catch (_) {
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

    if (indice >= 0) {
      carrito[indice].cantidad += cantidad.value
    } else {
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
  window.setTimeout(() => {
    showToast.value = false
  }, 3000)

  agregando.value = false
}
</script>

<template>
  <div class="max-w-[1440px] mx-auto px-6 md:px-8 py-10 md:py-14">
    <div v-if="loading" class="flex min-h-[60vh] items-center justify-center">
      <div class="flex flex-col items-center gap-4">
        <div class="h-8 w-8 animate-spin border-2 border-[#171717] border-t-transparent" />
        <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-[#666666]">
          Cargando producto
        </span>
      </div>
    </div>

    <div v-else-if="!producto" class="border border-[#E5E5E5] bg-white px-6 py-24 text-center">
      <p class="mx-auto max-w-[14ch] font-display text-[clamp(2.5rem,7vw,6rem)] uppercase leading-[0.9] text-[#171717]">
        Producto no encontrado
      </p>
      <button
        class="mt-8 border border-[#171717] bg-[#171717] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-white transition-colors duration-200 hover:bg-[#D4A017] hover:text-black"
        @click="navigateTo('/catalogo')"
      >
        Volver al catálogo
      </button>
    </div>

    <template v-else>
      <nav class="mb-10 flex flex-wrap items-center gap-2 font-mono text-[11px] uppercase tracking-[0.2em] text-[#666666]">
        <NuxtLink to="/catalogo" class="transition-colors hover:text-[#171717]">
          Catálogo
        </NuxtLink>
        <span class="text-[#B5B5B5]">/</span>
        <span>{{ producto.categoria_nombre || "General" }}</span>
        <span class="text-[#B5B5B5]">/</span>
        <span class="text-[#171717]">
          {{ producto.nombre }}
        </span>
      </nav>

      <div class="grid grid-cols-12 gap-px border border-[#E5E5E5] bg-[#E5E5E5]">
        <section class="relative col-span-12 lg:col-span-5 bg-white">
          <div class="relative aspect-square border-b border-[#E5E5E5] bg-[#F8F8F8]">
            <img
              :src="imagenProducto"
              :alt="producto.nombre"
              class="block h-full w-full object-cover"
              loading="eager"
              @error="manejarErrorImagen"
            />

            <svg
              aria-hidden="true"
              class="pointer-events-none absolute inset-0 h-full w-full opacity-10"
              viewBox="0 0 800 800"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M64 96H736" stroke="#D4A017" stroke-width="2" />
              <path d="M112 176L224 64L336 176L448 64L560 176L672 64" stroke="#D4A017" stroke-width="2" />
              <path d="M104 592L248 448L392 592L536 392L680 592" stroke="#D4A017" stroke-width="2" />
              <path d="M144 704H656" stroke="#D4A017" stroke-width="2" />
              <path d="M520 236L596 160L672 236" stroke="#D4A017" stroke-width="2" />
              <path d="M168 248L240 176L312 248" stroke="#D4A017" stroke-width="2" />
            </svg>

            <span
              v-if="producto.stock === 0"
              class="absolute inset-0 flex items-center justify-center bg-black/50"
            >
              <span class="font-display text-[clamp(2rem,4vw,3.5rem)] uppercase tracking-[0.18em] text-white">
                Agotado
              </span>
            </span>
          </div>
        </section>

        <section class="col-span-12 lg:col-span-7 bg-white p-6 md:p-10 lg:p-12">
          <div class="max-w-[860px]">
            <span class="inline-flex border border-[#171717] bg-[#171717] px-3 py-1 font-mono text-[10px] uppercase tracking-[0.2em] text-white">
              {{ producto.categoria_nombre || "General" }}
            </span>

            <h1 class="mt-6 max-w-[12ch] font-display text-[clamp(2.5rem,6vw,5rem)] uppercase leading-[0.9] text-[#171717]">
              {{ producto.nombre }}
            </h1>

            <p class="mt-6 max-w-[65ch] font-body text-[16px] leading-relaxed text-[#666666]">
              {{ producto.descripcion || "Producto para bodega con presentación profesional y disponibilidad según stock." }}
            </p>

            <div class="relative mt-10 border-y border-[#E5E5E5] py-8">
              <svg
                aria-hidden="true"
                class="pointer-events-none absolute right-0 top-0 h-full w-[60%] opacity-[0.12]"
                viewBox="0 0 720 220"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path d="M24 176H696" stroke="#D4A017" stroke-width="2" />
                <path d="M72 140L164 48L256 140L348 48L440 140L532 48L624 140" stroke="#D4A017" stroke-width="2" />
                <path d="M184 188L276 96L368 188L460 96L552 188" stroke="#D4A017" stroke-width="2" />
                <path d="M500 72L584 24L668 72" stroke="#D4A017" stroke-width="2" />
              </svg>

              <div class="relative flex flex-wrap items-end gap-x-4 gap-y-3">
                <span class="font-display text-[clamp(3rem,6vw,5rem)] leading-none text-[#171717]">
                  {{ formatearPrecio(producto.precio) }}
                </span>
                <span class="pb-2 font-mono text-[11px] uppercase tracking-[0.2em] text-[#666666]">
                  / {{ producto.unidad }}
                </span>
              </div>
            </div>

            <div class="mt-8 flex items-center gap-2">
              <span
                class="h-2.5 w-2.5"
                :class="producto.stock > 20 ? 'bg-[#1f7a1f]' : producto.stock > 0 ? 'bg-[#D4A017]' : 'bg-[#b42318]'"
                aria-hidden="true"
              />
              <span class="font-mono text-[11px] uppercase tracking-[0.2em] text-[#666666]">
                {{ producto.stock > 20 ? `${producto.stock} en stock` : producto.stock > 0 ? `Solo ${producto.stock} disponibles` : "Sin inventario" }}
              </span>
            </div>

            <div v-if="producto.stock > 0" class="mt-10 flex flex-col gap-4 md:flex-row md:items-stretch">
              <div class="flex w-fit border border-[#E5E5E5] bg-white">
                <button
                  class="flex min-h-[56px] min-w-[56px] items-center justify-center px-4 font-mono text-[18px] text-[#171717] transition-colors duration-200 hover:bg-[#F5F5F5] disabled:cursor-not-allowed disabled:opacity-30"
                  :disabled="cantidad <= 1"
                  @click="cantidad--"
                  aria-label="Disminuir cantidad"
                >
                  −
                </button>
                <span class="flex min-h-[56px] min-w-[72px] items-center justify-center border-x border-[#E5E5E5] px-5 font-mono text-[11px] uppercase tracking-[0.2em] text-[#171717]">
                  {{ cantidad }}
                </span>
                <button
                  class="flex min-h-[56px] min-w-[56px] items-center justify-center px-4 font-mono text-[18px] text-[#171717] transition-colors duration-200 hover:bg-[#F5F5F5] disabled:cursor-not-allowed disabled:opacity-30"
                  :disabled="cantidad >= producto.stock"
                  @click="cantidad++"
                  aria-label="Aumentar cantidad"
                >
                  +
                </button>
              </div>

              <button
                class="min-h-[56px] border border-[#171717] bg-[#171717] px-8 font-mono text-[11px] uppercase tracking-[0.25em] text-white transition-colors duration-200 hover:border-[#D4A017] hover:bg-[#D4A017] hover:text-black disabled:cursor-not-allowed disabled:opacity-50"
                :disabled="agregando"
                @click="agregarAlCarrito"
              >
                <span v-if="agregando" class="inline-flex items-center gap-3">
                  <span class="h-4 w-4 animate-spin border-2 border-white border-t-transparent" />
                  Agregando
                </span>
                <span v-else>Agregar</span>
              </button>
            </div>

            <button
              v-else
              disabled
              class="mt-10 min-h-[56px] border border-[#E5E5E5] bg-white px-8 font-mono text-[11px] uppercase tracking-[0.25em] text-[#666666] opacity-70"
            >
              Agotado
            </button>
          </div>
        </section>
      </div>
    </template>

    <Toast
      :message="toastMsg"
      :type="toastType"
      :show="showToast"
      @close="showToast = false"
    />
  </div>
</template>
