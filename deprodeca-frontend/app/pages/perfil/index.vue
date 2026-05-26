<!--
  ═══════════════════════════════════════════════════════════════
  perfil/index.vue — Mi Perfil · DEPRODECA
  Brutalismo aplicado a la gestión de perfil:
  • Sidebar modular (PerfilSidebar)
  • Datos en tabla con labels monospace
  • Formulario con inputs rectangulares, sin curvas
  • Botones toggle: Ver/Editar como bloques
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })

const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const perfil = ref<any>(null)
const editando = ref(false)
const loading = ref(true)
const saving = ref(false)
const toastMsg = ref("")
const toastType = ref<"success" | "error">("success")
const showToast = ref(false)
const form = ref({ nombre: "", empresa: "", telefono: "" })

// ─── Ciclo de vida ────────────────────────────────────────
onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const res = await $fetch<{ success: boolean; data: any }>(
      `${config.public.apiBase}/usuarios/perfil`,
      { headers: { Authorization: `Bearer ${token}` } },
    )
    if (res.success) {
      perfil.value = res.data
      form.value = {
        nombre: res.data.nombre,
        empresa: res.data.empresa,
        telefono: res.data.telefono,
      }
    }
  } catch (_) {
    /* fallback */
  } finally {
    loading.value = false
  }
})

// ─── Guardar perfil ───────────────────────────────────────
async function guardarPerfil() {
  saving.value = true
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  try {
    await $fetch(`${config.public.apiBase}/usuarios/perfil`, {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: form.value,
    })
    perfil.value = { ...perfil.value, ...form.value }
    toastMsg.value = "Perfil actualizado"
    toastType.value = "success"
  } catch (e: any) {
    toastMsg.value = e?.data?.message || "Error"
    toastType.value = "error"
  } finally {
    saving.value = false
    editando.value = false
    showToast.value = true
    setTimeout(() => (showToast.value = false), 3000)
  }
}
</script>

<template>
  <div class="page-enter max-w-[960px] mx-auto px-6 md:px-8 py-12">

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Perfil
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Mi Perfil<span class="text-[#D4A017]">.</span>
      </h1>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <!-- Contenido -->
    <template v-else-if="perfil">
      <div class="grid grid-cols-1 md:grid-cols-[220px_1fr] gap-0">

        <!-- Sidebar -->
        <PerfilSidebar />

        <!-- Panel de datos -->
        <div class="border border-borde border-l-0 bg-white p-8">

          <!-- Header panel -->
          <div class="flex items-center justify-between mb-8 border-b border-borde pb-6">
            <div class="flex items-center gap-3">
              <!-- Ícono: Persona geométrica -->
              <svg width="22" height="22" viewBox="0 0 22 22" fill="none" class="text-[#D4A017]">
                <circle cx="11" cy="7" r="4" stroke="currentColor" stroke-width="1.5"/>
                <path d="M3 21C3 16.6 6.6 13 11 13C15.4 13 19 16.6 19 21"
                      stroke="currentColor" stroke-width="1.5" stroke-linecap="square"/>
              </svg>
              <h2 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
                Datos Personales
              </h2>
            </div>

            <!-- Toggle Ver/Editar -->
            <button
              v-if="!editando"
              class="font-mono text-[10px] uppercase tracking-[0.2em] border border-borde px-4 py-2
                     text-texto-muted hover:text-texto hover:border-texto/30 transition-colors duration-200"
              @click="editando = true"
            >
              Editar
            </button>
          </div>

          <!-- MODO VER · Tabla de datos -->
          <div v-if="!editando">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-0 border border-borde">
              <div class="border-b border-r border-borde p-4">
                <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mb-1">Nombre</p>
                <p class="font-body text-body font-bold text-texto">{{ perfil.nombre }}</p>
              </div>
              <div class="border-b border-borde p-4">
                <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mb-1">Empresa</p>
                <p class="font-body text-body font-bold text-texto">{{ perfil.empresa }}</p>
              </div>
              <div class="border-r border-borde p-4">
                <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mb-1">RUC</p>
                <p class="font-body text-body font-bold text-texto">{{ perfil.ruc }}</p>
              </div>
              <div class="p-4">
                <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mb-1">Email</p>
                <p class="font-body text-body font-bold text-texto">{{ perfil.email }}</p>
              </div>
              <div class="col-span-1 sm:col-span-2 border-t border-borde p-4">
                <p class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em] mb-1">Teléfono</p>
                <p class="font-body text-body font-bold text-texto">{{ perfil.telefono || "—" }}</p>
              </div>
            </div>
          </div>

          <!-- MODO EDITAR · Formulario brutalista -->
          <form v-else @submit.prevent="guardarPerfil" class="space-y-6">
            <div>
              <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2">
                Nombre
              </label>
              <input
                v-model="form.nombre"
                required
                class="w-full border border-borde px-4 py-3 font-body text-body text-texto
                       bg-white placeholder:text-texto-muted
                       focus:border-[#D4A017] focus:outline-none
                       transition-colors duration-200 min-h-[48px]"
              />
            </div>
            <div>
              <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2">
                Empresa
              </label>
              <input
                v-model="form.empresa"
                required
                class="w-full border border-borde px-4 py-3 font-body text-body text-texto
                       bg-white placeholder:text-texto-muted
                       focus:border-[#D4A017] focus:outline-none
                       transition-colors duration-200 min-h-[48px]"
              />
            </div>
            <div>
              <label class="block font-mono text-[10px] text-texto-muted uppercase tracking-[0.2em] mb-2">
                Teléfono
              </label>
              <input
                v-model="form.telefono"
                type="tel"
                class="w-full border border-borde px-4 py-3 font-body text-body text-texto
                       bg-white placeholder:text-texto-muted
                       focus:border-[#D4A017] focus:outline-none
                       transition-colors duration-200 min-h-[48px]"
              />
            </div>

            <!-- Acciones -->
            <div class="flex gap-0 pt-2">
              <button
                type="submit"
                :disabled="saving"
                class="bg-texto text-white font-mono text-[11px] uppercase tracking-[0.15em]
                       px-8 py-3 hover:bg-[#D4A017] hover:text-black
                       transition-colors duration-200 disabled:opacity-50
                       min-h-[48px] flex items-center gap-2"
              >
                <span v-if="saving" class="w-4 h-4 border-2 border-white border-t-transparent animate-spin" />
                <span v-else>Guardar</span>
              </button>
              <button
                type="button"
                class="border border-l-0 border-borde font-mono text-[11px] uppercase tracking-[0.15em]
                       px-8 py-3 text-texto-muted hover:text-texto hover:bg-fondo
                       transition-colors duration-200 min-h-[48px]"
                @click="editando = false"
              >
                Cancelar
              </button>
            </div>
          </form>

        </div>
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
