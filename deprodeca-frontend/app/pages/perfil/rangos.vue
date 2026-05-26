<script setup lang="ts">

definePageMeta({
  layout: "default",
  middleware: "auth",
})

const config = useRuntimeConfig()
const gamificacion = ref<any>(null)
const leaderboard = ref<any[]>([])
const loading = ref(true)
const loadingLB = ref(true)
const wsConnected = ref(false)

let ws: WebSocket | null = null

const rangoActual = computed(() => gamificacion.value?.rango || "bronce")

const umbralSiguiente = computed(() => {
  const umbrales: Record<string, number> = { bronce: 5000, plata: 20000, oro: 50000, platino: 50000 }
  return umbrales[rangoActual.value] || 5000
})

const progreso = computed(() => {
  if (!gamificacion.value) return 0
  return Math.min((gamificacion.value.volumen_total / umbralSiguiente.value) * 100, 100)
})

onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return

  // Carga inicial via HTTP
  try {
    const [gr, lr] = await Promise.all([
      $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/gamificacion/rango`, {
        headers: { Authorization: `Bearer ${token}` },
      }),
      $fetch<{ success: boolean; data: any[] }>(`${config.public.apiBase}/gamificacion/leaderboard?limite=20`, {
        headers: { Authorization: `Bearer ${token}` },
      }),
    ])
    if (gr.success) gamificacion.value = gr.data
    if (lr.success) leaderboard.value = lr.data
  } catch (_) {
    // continuar aunque falle HTTP
  } finally {
    loading.value = false
    loadingLB.value = false
  }

  // WebSocket para actualizaciones en tiempo real
  conectarWebSocket(token)
})

onUnmounted(() => {
  if (ws) {
    ws.close()
    ws = null
  }
})

function conectarWebSocket(token: string) {
  const wsUrl = config.public.apiBase
    .replace("http://", "ws://")
    .replace("https://", "wss://")
    .replace("/api/v1", "") + `/ws/rangos?token=${token}`

  ws = new WebSocket(wsUrl)

  ws.onopen = () => {
    wsConnected.value = true
  }

  ws.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data)

      if (msg.type === "estado_actual") {
        gamificacion.value = msg.data
      } else if (msg.type === "cambio_rango") {
        gamificacion.value = msg.data
        cargarLeaderboard(token)
      } else if (msg.type === "leaderboard") {
        leaderboard.value = msg.data
      }
    } catch (_) {
      // ignorar mensajes malformados
    }
  }

  ws.onclose = () => {
    wsConnected.value = false
    setTimeout(() => {
      if (import.meta.client && localStorage.getItem("deprodeca_token")) {
        conectarWebSocket(token)
      }
    }, 5000)
  }

  ws.onerror = () => {
    wsConnected.value = false
  }
}

async function cargarLeaderboard(token: string) {
  try {
    const res = await $fetch<{ success: boolean; data: any[] }>(
      `${config.public.apiBase}/gamificacion/leaderboard?limite=20`,
      { headers: { Authorization: `Bearer ${token}` } },
    )
    if (res.success) leaderboard.value = res.data
  } catch (_) {
    // mantener datos anteriores
  }
}

function formatearMonto(monto: number) {
  return new Intl.NumberFormat("es-PE", { style: "currency", currency: "PEN" }).format(monto)
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-4 md:px-6 py-10">
    <h1 class="font-display text-display-lg text-text-primary uppercase mb-8">Mi Rango</h1>

    <div v-if="loading" class="flex justify-center py-20">
      <Spinner size="lg" label="Cargando gamificación..." />
    </div>

    <template v-else-if="gamificacion">
      <div class="grid grid-cols-1 lg:grid-cols-[1fr_400px] gap-8">
        <!-- Estado actual -->
        <div>
          <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-8 mb-6">
            <div class="flex items-center gap-4 mb-6">
              <RangoBadge :rango="rangoActual" size="lg" />
              <div>
                <h2 class="font-display text-display-md text-text-primary uppercase leading-none">
                  Rango {{ gamificacion.rango }}
                </h2>
                <p class="font-body text-small text-text-muted mt-1">
                  {{ gamificacion.puntos_totales.toLocaleString("es-PE") }} puntos acumulados
                </p>
              </div>
            </div>

            <!-- Progreso al siguiente rango -->
            <div v-if="rangoActual !== 'platino'">
              <p class="font-body text-small text-text-muted mb-3">
                Progreso a {{ rangoActual === "bronce" ? "Plata" : rangoActual === "plata" ? "Oro" : "Platino" }}
              </p>
              <ProgressBar
                :valor="gamificacion.volumen_total"
                :maximo="umbralSiguiente"
                :rango="rangoActual"
              />
              <p class="mt-2 font-body text-caption text-text-muted text-right">
                {{ formatearMonto(gamificacion.volumen_total) }} / {{ formatearMonto(umbralSiguiente) }}
              </p>
            </div>
            <div v-else class="text-center py-4">
              <span class="text-4xl" aria-hidden="true">💎</span>
              <p class="font-display text-display-md text-text-primary mt-2 uppercase">¡Rango Máximo!</p>
              <p class="font-body text-small text-text-muted mt-1">Eres Platino — el nivel más alto de DEPRODECA.</p>
            </div>
          </div>

          <!-- Beneficios por rango -->
          <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-8">
            <h3 class="font-body text-subheading font-bold text-text-primary mb-4">Beneficios de tu Rango</h3>
            <div class="space-y-3">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-status-success/20 flex items-center justify-center flex-shrink-0">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-status-success"><path d="M5 13l4 4L19 7"/></svg>
                </div>
                <p class="font-body text-small text-text-primary">Precios exclusivos B2B</p>
              </div>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-status-success/20 flex items-center justify-center flex-shrink-0">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-status-success"><path d="M5 13l4 4L19 7"/></svg>
                </div>
                <p class="font-body text-small text-text-primary">Envío gratis a partir de S/ 500</p>
              </div>
              <div
                v-if="rangoActual === 'plata' || rangoActual === 'oro' || rangoActual === 'platino'"
                class="flex items-center gap-3"
              >
                <div class="w-10 h-10 rounded-xl bg-status-success/20 flex items-center justify-center flex-shrink-0">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-status-success"><path d="M5 13l4 4L19 7"/></svg>
                </div>
                <p class="font-body text-small text-text-primary">Atención prioritaria en pedidos</p>
              </div>
              <div
                v-if="rangoActual === 'oro' || rangoActual === 'platino'"
                class="flex items-center gap-3"
              >
                <div class="w-10 h-10 rounded-xl bg-status-success/20 flex items-center justify-center flex-shrink-0">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-status-success"><path d="M5 13l4 4L19 7"/></svg>
                </div>
                <p class="font-body text-small text-text-primary">Descuento 5% en volumen</p>
              </div>
              <div
                v-if="rangoActual === 'platino'"
                class="flex items-center gap-3"
              >
                <div class="w-10 h-10 rounded-xl bg-status-success/20 flex items-center justify-center flex-shrink-0">
                  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-status-success"><path d="M5 13l4 4L19 7"/></svg>
                </div>
                <p class="font-body text-small text-text-primary">Acceso anticipado a nuevos productos</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Leaderboard -->
        <div>
          <div class="bg-white rounded-2xl border-2 border-border-base shadow-sm p-6 lg:sticky lg:top-[88px]">
            <h3 class="font-body text-subheading font-bold text-text-primary mb-4 flex items-center gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-brand-primary"><path d="M6 9H4.5a2.5 2.5 0 010-5C7 4 7 7 7 7"/><path d="M18 9h1.5a2.5 2.5 0 000-5C17 4 17 7 17 7"/><path d="M4 22h16"/><path d="M10 14.66V17c0 .55-.47.98-.97 1.21C7.85 18.75 7 20.24 7 22"/><path d="M14 14.66V17c0 .55.47.98.97 1.21C16.15 18.75 17 20.24 17 22"/><path d="M18 2H6v7a6 6 0 0012 0V2Z"/></svg>
              Leaderboard
              <span v-if="wsConnected" class="w-2 h-2 rounded-full bg-status-success inline-block" title="Conectado en vivo" />
            </h3>

            <div v-if="loadingLB" class="py-8">
              <Spinner size="sm" label="Cargando ranking..." />
            </div>

            <div v-else-if="leaderboard.length === 0" class="text-center py-8">
              <p class="font-body text-small text-text-muted">Sin datos aún</p>
            </div>

            <div v-else class="space-y-2">
              <div
                v-for="entry in leaderboard"
                :key="entry.usuario_id"
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg hover:bg-surface transition-colors"
              >
                <span
                  class="w-8 h-8 rounded-full flex items-center justify-center font-display text-small font-bold flex-shrink-0"
                  :class="{
                    'bg-brand-primary text-text-primary': entry.posicion === 1,
                    'bg-gray-200 text-text-muted': entry.posicion === 2,
                    'bg-amber-100 text-amber-700': entry.posicion === 3,
                    'bg-surface text-text-muted': entry.posicion > 3,
                  }"
                >
                  {{ entry.posicion }}
                </span>

                <div class="flex-1 min-w-0">
                  <p class="font-body text-small font-semibold text-text-primary truncate">
                    {{ entry.usuario_empresa || entry.usuario_nombre }}
                  </p>
                </div>

                <RangoBadge :rango="entry.rango" size="sm" :show-label="false" />

                <span class="font-body text-caption text-text-muted w-16 text-right">
                  {{ entry.puntos_totales.toLocaleString("es-PE") }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>

    <!-- Fallback si no hay datos -->
    <div v-else class="text-center py-20">
      <div class="w-20 h-20 mx-auto mb-4 rounded-full bg-surface flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-text-muted"><path d="M12 15l-4.24-4.24a6 6 0 0112.48 0L20 15"/></svg>
      </div>
      <h3 class="font-body text-subheading font-bold text-text-primary mb-2">Sin datos de gamificación</h3>
      <p class="font-body text-small text-text-muted">Realiza tu primera compra para empezar a acumular puntos.</p>
    </div>
  </div>
</template>
