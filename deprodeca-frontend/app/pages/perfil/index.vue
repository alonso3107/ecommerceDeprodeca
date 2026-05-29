<script setup lang="ts">
import PerfilCabecera from "~/features/perfil/components/PerfilCabecera.vue"
import PerfilStats from "~/features/perfil/components/PerfilStats.vue"
import PerfilRangoSection from "~/features/perfil/components/PerfilRangoSection.vue"
import PerfilPedidosRecientes from "~/features/perfil/components/PerfilPedidosRecientes.vue"

definePageMeta({ layout: "default", middleware: "auth" })

const config = useRuntimeConfig()
const perfil = ref<any>(null)
const gamificacion = ref<any>(null)
const pedidos = ref<any[]>([])
const leaderboard = ref<any[]>([])
const loading = ref(true)
const editando = ref(false)
const saving = ref(false)
const form = ref({ nombre: "", empresa: "", telefono: "" })

const wsConnected = ref(false)

function tokenAuth() {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  return token ? { Authorization: `Bearer ${token}` } : null
}

const statsComputados = computed(() => {
  const rango = (gamificacion.value?.rango || "bronce").toLowerCase()
  const descuentoMap: Record<string, number> = { bronce: 0, plata: 2, oro: 5, platino: 10 }
  return {
    puntos: gamificacion.value?.puntos_totales || 0,
    pedidos: pedidos.value.length,
    volumen: gamificacion.value?.volumen_total || 0,
    descuento: descuentoMap[rango] ?? 0,
  }
})

onMounted(async () => {
  const headers = tokenAuth()
  if (!headers) {
    loading.value = false
    return
  }
  try {
    const [perfilRes, rangoRes, pedidosRes] = await Promise.all([
      $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/usuarios/perfil`, { headers }),
      $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/gamificacion/rango`, { headers }),
      $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/pedidos`, { headers }),
    ])

    if (perfilRes?.success) {
      perfil.value = perfilRes.data
      form.value = {
        nombre: perfilRes.data?.nombre || "",
        empresa: perfilRes.data?.empresa || "",
        telefono: perfilRes.data?.telefono || "",
      }
    }
    if (rangoRes?.success) gamificacion.value = rangoRes.data
    if (pedidosRes?.success) pedidos.value = pedidosRes.data || []

    try {
      const lbRes = await $fetch<{ success: boolean; data: any[] }>(
        `${config.public.apiBase}/gamificacion/leaderboard?limite=10`,
        { headers },
      )
      if (lbRes?.success) leaderboard.value = lbRes.data || []
    } catch {
      leaderboard.value = []
    }
  } finally {
    loading.value = false
  }
})

async function guardarPerfil() {
  const headers = tokenAuth()
  if (!headers) return
  saving.value = true
  try {
    await $fetch(`${config.public.apiBase}/usuarios/perfil`, {
      method: "PUT",
      headers: { ...headers, "Content-Type": "application/json" },
      body: form.value,
    })
    perfil.value = { ...perfil.value, ...form.value }
    editando.value = false
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <main class="max-w-[1100px] mx-auto px-6 md:px-8 py-12">
    <div v-if="loading" class="flex flex-col items-center justify-center min-h-[50vh] gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <p class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto-muted">Cargando</p>
    </div>

    <template v-else-if="perfil">
      <section class="mb-12">
        <PerfilCabecera :perfil="perfil" :gamificacion="gamificacion" />
      </section>

      <section class="mb-12">
        <PerfilStats :stats="statsComputados" />
      </section>

      <section class="mb-12">
        <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-texto-muted mb-3">--- Datos Personales</p>
        <div class="border border-borde bg-white p-6 md:p-8">
          <div class="flex items-center justify-between mb-6 border-b border-borde pb-4">
            <h2 class="font-display text-display-sm uppercase text-texto">Mi Perfil</h2>
            <button v-if="!editando" class="border border-borde px-4 py-2 font-mono text-[10px] uppercase tracking-[0.15em] hover:bg-fondo transition duration-500" @click="editando = true">Editar</button>
          </div>

          <div v-if="!editando" class="grid grid-cols-1 sm:grid-cols-2 border border-borde">
            <div class="p-4 border-b sm:border-r border-borde"><p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-1">Nombre</p><p class="font-body text-texto">{{ perfil.nombre }}</p></div>
            <div class="p-4 border-b border-borde"><p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-1">Empresa</p><p class="font-body text-texto">{{ perfil.empresa }}</p></div>
            <div class="p-4 border-b sm:border-b-0 sm:border-r border-borde"><p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-1">RUC</p><p class="font-body text-texto">{{ perfil.ruc }}</p></div>
            <div class="p-4"><p class="font-mono text-[9px] uppercase text-texto-muted tracking-[0.15em] mb-1">Email</p><p class="font-body text-texto">{{ perfil.email }}</p></div>
          </div>

          <form v-else @submit.prevent="guardarPerfil" class="space-y-4">
            <div><label class="block font-mono text-[10px] uppercase tracking-[0.15em] text-texto-muted mb-2">Nombre</label><input v-model="form.nombre" class="w-full border border-borde bg-white px-4 py-3 font-body focus:border-[#D4A017] focus:outline-none" required /></div>
            <div><label class="block font-mono text-[10px] uppercase tracking-[0.15em] text-texto-muted mb-2">Empresa</label><input v-model="form.empresa" class="w-full border border-borde bg-white px-4 py-3 font-body focus:border-[#D4A017] focus:outline-none" required /></div>
            <div><label class="block font-mono text-[10px] uppercase tracking-[0.15em] text-texto-muted mb-2">Telefono</label><input v-model="form.telefono" type="tel" class="w-full border border-borde bg-white px-4 py-3 font-body focus:border-[#D4A017] focus:outline-none" /></div>
            <div class="flex gap-0 pt-1">
              <button type="submit" :disabled="saving" class="bg-texto text-white px-6 py-3 font-mono text-[10px] uppercase tracking-[0.15em] hover:scale-[1.02] transition duration-500 disabled:opacity-60">{{ saving ? "Guardando" : "Guardar" }}</button>
              <button type="button" class="border border-borde border-l-0 px-6 py-3 font-mono text-[10px] uppercase tracking-[0.15em] text-texto-muted hover:bg-fondo transition duration-500" @click="editando = false">Cancelar</button>
            </div>
          </form>
        </div>
      </section>

      <section class="mb-12">
        <PerfilRangoSection :gamificacion="gamificacion" :leaderboard="leaderboard" :ws-connected="wsConnected" />
      </section>

      <section class="mb-12">
        <PerfilPedidosRecientes :pedidos="pedidos.slice(0, 5)" />
      </section>
    </template>
  </main>
</template>
