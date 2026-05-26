<script setup lang="ts">
import { ref, computed, onMounted } from "vue"

definePageMeta({
  layout: "default",
})

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

onMounted(async () => {
  try {
    const res = await $fetch<{ success: boolean; data: any }>(
      `${config.public.apiBase}/productos/${slug.value}`,
    )
    if (res.success) producto.value = res.data
  } catch (_) {
    producto.value = null
  } finally {
    loading.value = false
  }
})

function formatearPrecio(precio: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency",
    currency: "PEN",
  }).format(precio)
}

function agregarAlCarrito() {
  if (!producto.value) return
  agregando.value = true

  if (import.meta.client) {
    const carrito = JSON.parse(localStorage.getItem("deprodeca_carrito") || "[]")
    const idx = carrito.findIndex((item: any) => item.id === producto.value.id)
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

  toastMsg.value = `${cantidad.value} x ${producto.value.nombre} agregado al carrito`
  toastType.value = "success"
  showToast.value = true
  setTimeout(() => (showToast.value = false), 3000)
  agregando.value = false
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    <!-- Loading -->
    <div v-if="loading" class="flex justify-center py-20">
      <Spinner size="lg" label="Cargando producto..." />
    </div>

    <!-- 404: Producto no encontrado -->
    <div
      v-else-if="producto === null"
      class="text-center py-20"
    >
      <h2 class="font-display text-display-md text-text-primary mb-4">
        Producto no encontrado
      </h2>
      <Button variant="primary" @click="navigateTo('/catalogo')">
        Volver al Catálogo
      </Button>
    </div>

    <!-- Producto encontrado -->
    <template v-else>
      <!-- Breadcrumb -->
      <nav class="mb-6 font-body text-small text-text-muted">
        <NuxtLink
          to="/catalogo"
          class="hover:text-brand-primary transition-colors"
        >
          Catálogo
        </NuxtLink>
        <span class="mx-2">/</span>
        <NuxtLink
          :to="`/catalogo?categoria=${producto.categoria_slug}`"
          class="hover:text-brand-primary transition-colors"
        >
          {{ producto.categoria_nombre }}
        </NuxtLink>
        <span class="mx-2">/</span>
        <span class="text-text-primary font-semibold">{{ producto.nombre }}</span>
      </nav>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
        <!-- Imagen del producto -->
        <div class="relative rounded-2xl overflow-hidden bg-surface border-2 border-border-base aspect-square">
          <img
            :src="producto.imagen_url || 'https://images.unsplash.com/photo-1543168256-418811576931?w=600&q=80'"
            :alt="`Producto: ${producto.nombre}`"
            class="w-full h-full object-cover"
          />
          <span
            v-if="producto.stock <= 10 && producto.stock > 0"
            class="absolute top-4 right-4 px-3 py-1.5 rounded-full bg-status-warning text-white font-body text-caption font-semibold"
          >
            &iexcl;&Uacute;ltimas {{ producto.stock }}!
          </span>
        </div>

        <!-- Info del producto -->
        <div class="flex flex-col justify-center">
          <span class="inline-block px-3 py-1 rounded-full bg-brand-primary/20 text-text-primary font-body text-caption font-semibold mb-3 w-fit">
            {{ producto.categoria_nombre }}
          </span>

          <h1 class="font-display text-display-lg text-text-primary uppercase mb-3 leading-tight">
            {{ producto.nombre }}
          </h1>

          <p class="font-body text-body text-text-muted mb-6 leading-relaxed">
            {{ producto.descripcion || "Producto de alta calidad para tu bodega." }}
          </p>

          <!-- Precio -->
          <div class="flex items-baseline gap-2 mb-6">
            <span class="font-display text-display-xl text-text-primary leading-none">
              {{ formatearPrecio(producto.precio) }}
            </span>
            <span class="font-body text-body text-text-muted">/ {{ producto.unidad }}</span>
          </div>

          <!-- Stock -->
          <p class="font-body text-small mb-6">
            <span v-if="producto.stock > 20" class="text-status-success font-semibold">
              {{ producto.stock }} en stock
            </span>
            <span v-else-if="producto.stock > 0" class="text-status-warning font-semibold">
              Solo quedan {{ producto.stock }}
            </span>
            <span v-else class="text-status-error font-semibold">Agotado</span>
          </p>

          <!-- Cantidad + Agregar al carrito -->
          <div v-if="producto.stock > 0" class="flex items-center gap-4">
            <div class="flex items-center border-2 border-border-base rounded-lg">
              <button
                class="px-3 py-2 font-body text-body hover:bg-surface transition-colors min-h-[44px] min-w-[44px]"
                :disabled="cantidad <= 1"
                @click="cantidad--"
              >
                &minus;
              </button>
              <span class="px-4 py-2 font-body text-body font-semibold min-w-[48px] text-center">
                {{ cantidad }}
              </span>
              <button
                class="px-3 py-2 font-body text-body hover:bg-surface transition-colors min-h-[44px] min-w-[44px]"
                :disabled="cantidad >= producto.stock"
                @click="cantidad++"
              >
                +
              </button>
            </div>

            <Button
              variant="primary"
              size="lg"
              :loading="agregando"
              @click="agregarAlCarrito"
            >
              Agregar al Carrito
            </Button>
          </div>

          <Button
            v-else
            variant="outline"
            size="lg"
            disabled
          >
            Producto Agotado
          </Button>
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
