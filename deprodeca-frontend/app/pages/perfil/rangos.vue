<!--
  ═══════════════════════════════════════════════════════════════
  perfil/rangos.vue — Mi Rango · Gamificación DEPRODECA
  Brutalismo aplicado a la gamificación:
  • Rango actual con badge geométrico (no emoji, no rounded)
  • Barra de progreso rectangular (no rounded-full)
  • Leaderboard como tabla de datos crudos
  • Indicador WebSocket como punto técnico
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })

const config = useRuntimeConfig()

// ─── Estado ───────────────────────────────────────────────
const gamificacion = ref<any>(null)
const leaderboard = ref<any[]>([])
const loading = ref(true)
const loadingLB = ref(true)
const wsConnected = ref(false)
let ws: WebSocket | null = null

// ─── Computed ─────────────────────────────────────────────
const rangoActual = computed(() => gamificacion.value?.rango || "bronce")

const umbralSiguiente = computed(() => {
  const umbrales: Record<string, number> = {
    bronce: 5000, plata: 20000, oro: 50000, platino: 50000,
  }
  return umbrales[rangoActual.value] || 5000
})

const progreso = computed(() => {
  if (!gamificacion.value) return 0
  return Math.min((gamificacion.value.volumen_total / umbralSiguiente.value) * 100, 100)
})

const rangoSiguiente = computed(() => {
  const m: Record<string, string> = { bronce: "Plata", plata: "Oro", oro: "Platino" }
  return m[rangoActual.value] || ""
})

// ─── Beneficios por rango ─────────────────────────────────
const beneficios = computed(() => {
  const b = ["Precios exclusivos B2B"]
  if (["plata", "oro", "platino"].includes(rangoActual.value))
    b.push("Envío gratis + prioridad")
  if (["oro", "platino"].includes(rangoActual.value))
    b.push("5% descuento en volumen")
  if (rangoActual.value === "platino")
    b.push("Acceso anticipado + soporte VIP")
  return b
})

// ─── Utilidades ───────────────────────────────────────────
function formatearMonto(monto: number) {
  return new Intl.NumberFormat("es-PE", {
    style: "currency", currency: "PEN", minimumFractionDigits: 0,
  }).format(monto)
}

function colorRango(r: string) {
  return {
    bronce: { bg: "#CD7F32", hex: "#CD7F32" },
    plata: { bg: "#A8A8A8", hex: "#A8A8A8" },
    oro: { bg: "#D4A017", hex: "#D4A017" },
    platino: { bg: "#7B2FBE", hex: "#7B2FBE" },
  }[r] || { bg: "#CD7F32", hex: "#CD7F32" }
}

// ─── Ciclo de vida ────────────────────────────────────────
onMounted(async () => {
  const token = import.meta.client ? localStorage.getItem("deprodeca_token") : null
  if (!token) return
  try {
    const [gr, lr] = await Promise.all([
      $fetch<{ success: boolean; data: any }>(`${config.public.apiBase}/gamificacion/rango`, {
        headers: { Authorization: `Bearer ${token}` },
      }),
      $fetch<{ success: boolean; data: any[] }>(
        `${config.public.apiBase}/gamificacion/leaderboard?limite=20`,
        { headers: { Authorization: `Bearer ${token}` } },
      ),
    ])
    if (gr.success) gamificacion.value = gr.data
    if (lr.success) leaderboard.value = lr.data
  } catch (_) { /* fallback */ }
  finally { loading.value = false; loadingLB.value = false }

  conectarWebSocket(token)
})

onUnmounted(() => {
  if (ws) { ws.close(); ws = null }
})

// ─── WebSocket ────────────────────────────────────────────
function conectarWebSocket(token: string) {
  const wsUrl = config.public.apiBase
    .replace("http://", "ws://").replace("https://", "wss://")
    .replace("/api/v1", "") + `/ws/rangos?token=${token}`

  ws = new WebSocket(wsUrl)
  ws.onopen = () => { wsConnected.value = true }
  ws.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data)
      if (msg.type === "estado_actual" || msg.type === "cambio_rango")
        gamificacion.value = msg.data
      if (msg.type === "leaderboard" || msg.type === "cambio_rango")
        leaderboard.value = msg.data
    } catch (_) { /* ignorar */ }
  }
  ws.onclose = () => {
    wsConnected.value = false
    setTimeout(() => {
      if (import.meta.client && localStorage.getItem("deprodeca_token"))
        conectarWebSocket(token)
    }, 5000)
  }
  ws.onerror = () => { wsConnected.value = false }
}
</script>

<template>
  <div class="page-enter max-w-[1280px] mx-auto px-6 md:px-8 py-12">

    <!-- Encabezado -->
    <div class="mb-10">
      <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.3em] mb-3">
        ─── Rango
      </p>
      <h1 class="font-display text-display-lg text-texto uppercase leading-[0.95]">
        Mi Rango<span class="text-[#D4A017]">.</span>
      </h1>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-20 gap-4">
      <div class="w-8 h-8 border-2 border-texto border-t-transparent animate-spin" />
      <span class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">Cargando...</span>
    </div>

    <!-- Contenido -->
    <template v-else-if="gamificacion">
      <div class="grid grid-cols-1 lg:grid-cols-[1fr_380px] gap-0 border border-borde">

        <!-- ─── COLUMNA IZQUIERDA ─────────────────────────── -->
        <div class="border-b lg:border-b-0 lg:border-r border-borde">

          <!-- Rango actual · Cabecera masiva -->
          <div class="p-8 border-b border-borde">
            <div class="flex items-center gap-4 mb-6">

              <!-- Badge geométrico del rango -->
              <div
                class="w-16 h-16 flex items-center justify-center flex-shrink-0"
                :style="{ border: `3px solid ${colorRango(rangoActual).hex}` }"
              >
                <!-- Bronce: círculo -->
                <svg v-if="rangoActual === 'bronce'" width="28" height="28" viewBox="0 0 28 28" fill="none">
                  <circle cx="14" cy="14" r="9" :stroke="colorRango(rangoActual).hex" stroke-width="2.5"/>
                </svg>
                <!-- Plata: rombo -->
                <svg v-if="rangoActual === 'plata'" width="28" height="28" viewBox="0 0 28 28" fill="none">
                  <path d="M14 3L25 14L14 25L3 14L14 3Z" :stroke="colorRango(rangoActual).hex" stroke-width="2.5" stroke-linejoin="miter"/>
                </svg>
                <!-- Oro: hexágono -->
                <svg v-if="rangoActual === 'oro'" width="28" height="28" viewBox="0 0 28 28" fill="none">
                  <path d="M14 3L24 9V19L14 25L4 19V9L14 3Z" :stroke="colorRango(rangoActual).hex" stroke-width="2.5" stroke-linejoin="miter"/>
                </svg>
                <!-- Platino: estrella 4 puntas -->
                <svg v-if="rangoActual === 'platino'" width="28" height="28" viewBox="0 0 28 28" fill="none">
                  <path d="M14 2L17 11L26 14L17 17L14 26L11 17L2 14L11 11L14 2Z"
                        :stroke="colorRango(rangoActual).hex" stroke-width="2.5" stroke-linejoin="miter"/>
                </svg>
              </div>

              <div>
                <p class="font-mono text-[10px] uppercase tracking-[0.2em] mb-1"
                   :style="{ color: colorRango(rangoActual).hex }">
                  Rango Actual
                </p>
                <h2 class="font-display text-display-lg text-texto uppercase leading-none">
                  {{ gamificacion.rango }}
                </h2>
                <p class="font-mono text-[11px] text-texto-muted mt-2 tracking-wider">
                  {{ (gamificacion.puntos_totales || 0).toLocaleString("es-PE") }} puntos
                </p>
              </div>
            </div>

            <!-- Barra de progreso · Rectangular brutalista -->
            <div v-if="rangoActual !== 'platino'">
              <div class="flex justify-between mb-2">
                <span class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">
                  Progreso a {{ rangoSiguiente }}
                </span>
                <span class="font-mono text-[9px] text-texto-muted uppercase tracking-[0.2em]">
                  {{ Math.round(progreso) }}%
                </span>
              </div>

              <!-- Barra rectangular, sin rounded-full -->
              <div class="w-full h-4 border border-borde bg-fondo" role="progressbar"
                   :aria-valuenow="gamificacion.volumen_total"
                   :aria-valuemin="0" :aria-valuemax="umbralSiguiente">
                <div class="h-full transition-all duration-700 ease-out"
                     :style="{
                       width: `${progreso}%`,
                       backgroundColor: colorRango(rangoActual).bg,
                     }" />
              </div>

              <div class="flex justify-between mt-2 font-mono text-[9px] text-texto-muted uppercase tracking-[0.15em]">
                <span>{{ formatearMonto(gamificacion.volumen_total || 0) }}</span>
                <span>{{ formatearMonto(umbralSiguiente) }}</span>
              </div>
            </div>

            <!-- Platino: rango máximo -->
            <div v-else class="py-4">
              <div class="flex items-center gap-3">
                <svg width="32" height="32" viewBox="0 0 28 28" fill="none" class="flex-shrink-0">
                  <path d="M14 2L17 11L26 14L17 17L14 26L11 17L2 14L11 11L14 2Z"
                        :stroke="colorRango('platino').hex" stroke-width="2.5" stroke-linejoin="miter"/>
                  <path d="M14 2L17 11L26 14L17 17L14 26L11 17L2 14L11 11L14 2Z"
                        :fill="colorRango('platino').hex" opacity="0.1"/>
                </svg>
                <div>
                  <p class="font-display text-display-md text-texto uppercase leading-none">
                    ¡Rango Máximo!
                  </p>
                  <p class="font-mono text-[10px] text-texto-muted mt-1">
                    Platino — el nivel más alto de DEPRODECA
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- Beneficios · Lista con íconos geométricos -->
          <div class="p-8">
            <div class="flex items-center gap-3 mb-6">
              <svg width="18" height="18" viewBox="0 0 18 18" fill="none" class="text-[#D4A017]">
                <path d="M9 1L17 9L9 17L1 9L9 1Z" stroke="currentColor" stroke-width="1.5"/>
                <path d="M6.5 9.5L8.5 11.5L12 8" stroke="currentColor" stroke-width="2"
                      stroke-linecap="square" stroke-linejoin="miter"/>
              </svg>
              <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
                Beneficios
              </h3>
            </div>

            <div class="space-y-0 border border-borde">
              <div
                v-for="(b, i) in beneficios"
                :key="b"
                class="flex items-center gap-3 px-4 py-3.5"
                :class="i > 0 ? 'border-t border-borde' : ''"
              >
                <!-- Check cuadrado geométrico -->
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none"
                     class="text-exito flex-shrink-0">
                  <rect x="1" y="1" width="14" height="14" stroke="currentColor" stroke-width="1.5"/>
                  <path d="M4.5 8.5L6.5 10.5L11.5 5.5" stroke="currentColor" stroke-width="2"
                        stroke-linecap="square" stroke-linejoin="miter"/>
                </svg>
                <span class="font-body text-small text-texto">{{ b }}</span>
              </div>
            </div>
          </div>

        </div>

        <!-- ─── COLUMNA DERECHA · Leaderboard ─────────────── -->
        <div class="bg-white p-8 lg:sticky lg:top-[84px] h-fit">

          <!-- Header -->
          <div class="flex items-center justify-between mb-6">
            <div class="flex items-center gap-2">
              <!-- Ícono: Trofeo/pódium geométrico -->
              <svg width="20" height="20" viewBox="0 0 20 20" fill="none" class="text-[#D4A017]">
                <rect x="4" y="2" width="4" height="5" stroke="currentColor" stroke-width="1.5"/>
                <rect x="12" y="2" width="4" height="5" stroke="currentColor" stroke-width="1.5"/>
                <path d="M5 7H3V12H7V7" stroke="currentColor" stroke-width="1.5"/>
                <path d="M15 7H17V12H13V7" stroke="currentColor" stroke-width="1.5"/>
                <rect x="7" y="11" width="6" height="7" stroke="currentColor" stroke-width="1.5"/>
              </svg>
              <h3 class="font-mono text-[11px] uppercase tracking-[0.2em] text-texto font-bold">
                Leaderboard
              </h3>
            </div>

            <!-- WebSocket indicator -->
            <span
              class="w-2.5 h-2.5"
              :class="wsConnected ? 'bg-exito' : 'bg-borde'"
              :title="wsConnected ? 'Conectado en vivo' : 'Desconectado'"
            />
          </div>

          <!-- Loading LB -->
          <div v-if="loadingLB" class="flex justify-center py-10">
            <div class="w-6 h-6 border-2 border-texto border-t-transparent animate-spin" />
          </div>

          <!-- Sin datos -->
          <p v-else-if="leaderboard.length === 0" class="font-mono text-[10px] text-texto-muted text-center py-10 uppercase tracking-wider">
            Sin datos aún
          </p>

          <!-- Tabla de ranking -->
          <div v-else class="border border-borde">
            <div
              v-for="(entry, i) in leaderboard"
              :key="entry.usuario_id"
              class="flex items-center gap-3 px-4 py-3"
              :class="i > 0 ? 'border-t border-borde' : ''"
            >
              <!-- Posición · Ícono geométrico según puesto -->
              <div class="w-8 h-8 flex items-center justify-center flex-shrink-0">
                <!-- #1: Corona -->
                <svg v-if="entry.posicion === 1" width="20" height="18" viewBox="0 0 20 18" fill="none" class="text-[#D4A017]">
                  <path d="M3 15L10 2L17 15L10 10L3 15Z" stroke="currentColor" stroke-width="2"
                        stroke-linecap="square" stroke-linejoin="miter"/>
                  <rect x="7" y="14" width="6" height="2" fill="currentColor"/>
                </svg>
                <!-- #2: Diamante -->
                <svg v-else-if="entry.posicion === 2" width="18" height="18" viewBox="0 0 18 18" fill="none" class="text-texto-muted">
                  <path d="M9 2L16 9L9 16L2 9L9 2Z" stroke="currentColor" stroke-width="2" stroke-linejoin="miter"/>
                </svg>
                <!-- #3: Hexágono -->
                <svg v-else-if="entry.posicion === 3" width="18" height="18" viewBox="0 0 18 18" fill="none" class="text-borde-hover">
                  <path d="M9 2L15 6V12L9 16L3 12V6L9 2Z" stroke="currentColor" stroke-width="2" stroke-linejoin="miter"/>
                </svg>
                <!-- #4+: Número -->
                <span v-else class="font-mono text-[11px] text-texto-muted font-bold">
                  {{ entry.posicion }}
                </span>
              </div>

              <!-- Nombre -->
              <div class="flex-1 min-w-0">
                <p class="font-body text-small font-bold text-texto truncate">
                  {{ entry.usuario_empresa || entry.usuario_nombre }}
                </p>
              </div>

              <!-- Badge rango mini -->
              <span class="font-mono text-[9px] uppercase tracking-wider font-bold"
                    :style="{ color: colorRango(entry.rango).hex }">
                {{ entry.rango }}
              </span>

              <!-- Puntos -->
              <span class="font-mono text-[10px] text-texto-muted w-20 text-right tracking-wider">
                {{ (entry.puntos_totales || 0).toLocaleString("es-PE") }} pts
              </span>
            </div>
          </div>
        </div>

      </div>
    </template>

    <!-- Fallback sin datos -->
    <div v-else class="text-center py-20">
      <svg class="mx-auto mb-6 text-texto/15" width="64" height="64" viewBox="0 0 64 64" fill="none">
        <path d="M32 8L48 24V40L32 56L16 40V24L32 8Z" stroke="currentColor" stroke-width="2"/>
        <path d="M32 8V56M16 24L48 24" stroke="currentColor" stroke-width="1" opacity="0.5"/>
      </svg>
      <p class="font-display text-display-md text-texto/20 uppercase mb-4">Sin Datos</p>
      <p class="font-body text-body text-texto-muted">
        Realizá tu primera compra para empezar a acumular puntos.
      </p>
    </div>

  </div>
</template>
