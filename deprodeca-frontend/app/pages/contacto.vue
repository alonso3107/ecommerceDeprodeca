<!--
  ═══════════════════════════════════════════════════════════════
  contacto.vue — Página de Contacto DEPRODECA
  Brutalista × Cyberpunk Sutil: bordes duros, acentos neón dorado,
  glitch tipográfico controlado, geometría asimétrica.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
definePageMeta({ layout: "default" })

const supabase = useSupabase()

// ─── Estado del formulario ──────────────────────────────
const formulario = reactive({
  nombre: "",
  email: "",
  telefono: "",
  motivo: "",
  mensaje: "",
})

const enviando = ref(false)
const enviado = ref(false)
const error = ref("")

const motivoOptions = [
  "Quiero registrarme como bodega",
  "Consulta sobre pedidos",
  "Soporte técnico",
  "Devolución o reclamo",
  "Otro",
] as const

const infoContacto = [
  {
    titulo: "WhatsApp",
    valor: "+51 900 000 000",
    desc: "Respuesta en menos de 2 horas",
    link: "https://wa.me/51900000000",
    svgPath: "M24 8A16 16 0 0 0 8 24c0 2.8.7 5.5 2 7.9L8 40l8.4-2.2A15.9 15.9 0 1 0 24 8zm-4.5 22.5l-1.3-1.3c-2.2-2.2-3.6-5.2-3.6-8.5s1.4-6.3 3.6-8.5l1.3-1.3 1.3 1.3c2.2 2.2 3.6 5.2 3.6 8.5s-1.4 6.3-3.6 8.5l-1.3 1.3zM19.5 20c.5 0 .9.4.9.9v.2c0 .5-.4.9-.9.9s-.9-.4-.9-.9v-.2c0-.5.4-.9.9-.9z",
  },
  {
    titulo: "Email",
    valor: "contacto@deprodeca.pe",
    desc: "Respondemos en 24h hábiles",
    link: "mailto:contacto@deprodeca.pe",
    svgPath: "M8 12h32v16a4 4 0 0 1-4 4H12a4 4 0 0 1-4-4V12zm4-4h24a4 4 0 0 1 4 4H8a4 4 0 0 1 4-4zm12 16l12-8v12l-12-8zm0 0l-12 8V16l12 8z",
  },
  {
    titulo: "Ubicación",
    valor: "Ica, Perú",
    desc: "Cobertura: Ica, Chincha, Pisco, Nazca",
    link: null,
    svgPath: "M24 4c-5.5 0-10 4.5-10 10 0 7.7 10 20 10 20s10-12.3 10-20c0-5.5-4.5-10-10-10zm0 14a4 4 0 1 1 0-8 4 4 0 0 1 0 8z",
  },
] as const

const horarios = [
  { dia: "Lunes — Viernes", hora: "8:00 — 18:00" },
  { dia: "Sábado", hora: "9:00 — 13:00" },
  { dia: "Domingo", hora: "Cerrado" },
] as const

// ─── Enviar formulario a Supabase ─────────────────────
async function enviarFormulario() {
  if (enviando.value) return
  
  // Validación básica
  if (!formulario.nombre.trim() || !formulario.email.trim() || !formulario.mensaje.trim()) {
    error.value = "Completá nombre, email y mensaje como mínimo."
    return
  }

  enviando.value = true
  error.value = ""

  try {
    const { error: err } = await supabase
      .from("consultas")
      .insert({
        nombre: formulario.nombre.trim(),
        email: formulario.email.trim(),
        telefono: formulario.telefono.trim(),
        motivo: formulario.motivo || "No especificado",
        mensaje: formulario.mensaje.trim(),
      })

    if (err) throw err

    enviado.value = true
  } catch (e: any) {
    error.value = e.message || "Error al enviar. Intentá de nuevo."
  } finally {
    enviando.value = false
  }
}
</script>

<template>
  <div class="page-enter">
    
    <!-- ═══ HERO — Título glitch + SVG cyberpunk ═══ -->
    <section class="relative py-28 md:py-36 bg-[#171717] border-b-4 border-[#D4A017] overflow-hidden">
      
      <!-- SVG cyberpunk de fondo — líneas diagonales + cuadrícula -->
      <svg
        aria-hidden="true"
        class="pointer-events-none absolute inset-0 h-full w-full opacity-[0.07]"
        viewBox="0 0 1440 600"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <!-- Grid -->
        <defs>
          <pattern id="cyber-grid" x="0" y="0" width="80" height="80" patternUnits="userSpaceOnUse">
            <path d="M80 0V80M0 80H80" stroke="#D4A017" stroke-width="0.5" />
          </pattern>
        </defs>
        <rect width="1440" height="600" fill="url(#cyber-grid)" />
        
        <!-- Líneas diagonales -->
        <path d="M-200 0L400 600" stroke="#D4A017" stroke-width="1.5" />
        <path d="M0 0L600 600" stroke="#D4A017" stroke-width="1" />
        <path d="M300 0L900 600" stroke="#D4A017" stroke-width="0.8" />
        <path d="M600 0L1200 600" stroke="#D4A017" stroke-width="0.5" />
        <path d="M900 0L1500 600" stroke="#D4A017" stroke-width="0.8" />
        <path d="M1100 0L1700 600" stroke="#D4A017" stroke-width="1" />
        
        <!-- Círculo central -->
        <circle cx="720" cy="300" r="180" stroke="#D4A017" stroke-width="1" />
        <circle cx="720" cy="300" r="100" stroke="#D4A017" stroke-width="2" opacity="0.5" />
      </svg>

      <!-- Elemento decorativo esquina -->
      <div class="absolute top-0 right-0 w-64 h-64 border-l border-b border-[#D4A017]/20" aria-hidden="true" />

      <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
        <div class="max-w-[680px]">
          <p class="font-mono text-[11px] text-[#D4A017] uppercase tracking-[0.3em] mb-5">
            ─── Contacto
          </p>
          <h1 class="font-display text-[clamp(3rem,8vw,7rem)] text-[#FAFAFA] uppercase leading-[0.88] tracking-[-0.02em]">
            Hablemos<span class="text-[#D4A017]">.</span>
          </h1>
          <div class="mt-5 h-[2px] w-32 bg-[#D4A017]" />
          <p class="mt-6 font-body text-[18px] text-[#FAFAFA]/50 max-w-[500px] leading-relaxed">
            ¿Tenés una bodega y querés comprar directo de Nestlé? 
            Escribinos. Respondemos rápido, sin formularios eternos.
          </p>
        </div>
      </div>
    </section>

    <!-- ═══ FORMULARIO + INFO — Layout asimétrico ═══ -->
    <section class="py-24 md:py-32 bg-[#FFFFFF]">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        
        <div class="grid grid-cols-1 lg:grid-cols-[5fr_4fr] gap-0 border border-[#E5E5E5] bg-[#E5E5E5]">
          
          <!-- COLUMNA IZQUIERDA · Formulario brutalista -->
          <div class="bg-[#FAFAFA] p-8 md:p-12 lg:p-14">
            
            <div class="mb-10">
              <p class="font-mono text-[10px] text-[#D4A017] uppercase tracking-[0.25em] mb-3">
                ─── Formulario
              </p>
              <h2 class="font-display text-[clamp(1.8rem,4vw,2.5rem)] text-[#171717] uppercase leading-[0.95]">
                Envianos tu<br />consulta<span class="text-[#D4A017]">.</span>
              </h2>
            </div>

            <form class="space-y-6" @submit.prevent="enviarFormulario">
              
              <!-- Mensaje de éxito -->
              <div v-if="enviado" class="border border-[#D4A017] bg-[#D4A017]/5 p-6">
                <p class="font-display text-[18px] text-[#D4A017] uppercase tracking-wide mb-2">
                  ¡Mensaje enviado!
                </p>
                <p class="font-body text-[14px] text-[#666666] leading-relaxed">
                  Gracias por escribirnos. Te respondemos en menos de 24 horas.
                </p>
              </div>

              <!-- Mensaje de error -->
              <div v-if="error" class="border border-[#DC2626] bg-[#DC2626]/5 p-4">
                <p class="font-mono text-[12px] text-[#DC2626]">{{ error }}</p>
              </div>

              <!-- Fila: Nombre + Email -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div>
                  <label class="block font-mono text-[10px] text-[#666666] uppercase tracking-[0.15em] mb-2">
                    Nombre completo *
                  </label>
                  <input
                    v-model="formulario.nombre"
                    type="text"
                    placeholder="Ej: María López"
                    class="w-full bg-[#FFFFFF] border border-[#D1D1D1] px-4 py-3.5 font-body text-[15px] text-[#171717] placeholder:text-[#BBBBBB] outline-none focus:border-[#D4A017] transition-colors duration-300"
                  />
                </div>
                <div>
                  <label class="block font-mono text-[10px] text-[#666666] uppercase tracking-[0.15em] mb-2">
                    Email *
                  </label>
                  <input
                    v-model="formulario.email"
                    type="email"
                    placeholder="maria@bodega.pe"
                    class="w-full bg-[#FFFFFF] border border-[#D1D1D1] px-4 py-3.5 font-body text-[15px] text-[#171717] placeholder:text-[#BBBBBB] outline-none focus:border-[#D4A017] transition-colors duration-300"
                  />
                </div>
              </div>

              <!-- Teléfono -->
              <div>
                <label class="block font-mono text-[10px] text-[#666666] uppercase tracking-[0.15em] mb-2">
                  Teléfono (WhatsApp)
                </label>
                <input
                  v-model="formulario.telefono"
                  type="tel"
                  placeholder="+51 900 000 000"
                  class="w-full bg-[#FFFFFF] border border-[#D1D1D1] px-4 py-3.5 font-body text-[15px] text-[#171717] placeholder:text-[#BBBBBB] outline-none focus:border-[#D4A017] transition-colors duration-300"
                />
              </div>

              <!-- Motivo -->
              <div>
                <label class="block font-mono text-[10px] text-[#666666] uppercase tracking-[0.15em] mb-2">
                  Motivo de contacto
                </label>
                <select
                  v-model="formulario.motivo"
                  class="w-full bg-[#FFFFFF] border border-[#D1D1D1] px-4 py-3.5 font-body text-[15px] text-[#171717] outline-none focus:border-[#D4A017] transition-colors duration-300 appearance-none cursor-pointer"
                  style="background-image: url('data:image/svg+xml,%3Csvg xmlns=%22http://www.w3.org/2000/svg%22 width=%2212%22 height=%2212%22%3E%3Cpath d=%22M2 4l4 4 4-4%22 stroke=%22%23171717%22 stroke-width=%222%22 fill=%22none%22/%3E%3C/svg%3E'); background-repeat: no-repeat; background-position: right 16px center;"
                >
                  <option value="" disabled>Seleccioná un motivo</option>
                  <option v-for="m in motivoOptions" :key="m" :value="m">{{ m }}</option>
                </select>
              </div>

              <!-- Mensaje -->
              <div>
                <label class="block font-mono text-[10px] text-[#666666] uppercase tracking-[0.15em] mb-2">
                  Mensaje *
                </label>
                <textarea
                  v-model="formulario.mensaje"
                  rows="5"
                  placeholder="Contanos sobre tu bodega y cómo podemos ayudarte..."
                  class="w-full bg-[#FFFFFF] border border-[#D1D1D1] px-4 py-3.5 font-body text-[15px] text-[#171717] placeholder:text-[#BBBBBB] outline-none focus:border-[#D4A017] transition-colors duration-300 resize-none"
                />
              </div>

              <!-- Botón submit — negro sólido brutalista -->
              <button
                type="submit"
                :disabled="enviando || enviado"
                class="w-full bg-[#171717] text-[#FAFAFA] font-display text-[16px] uppercase tracking-[0.08em] py-4 hover:bg-[#D4A017] hover:text-[#171717] transition-all duration-300 active:scale-[0.99] flex items-center justify-center gap-3 group disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span v-if="enviando" class="w-5 h-5 border-2 border-[#FAFAFA] border-t-transparent animate-spin" />
                <span v-else-if="enviado">Mensaje enviado ✓</span>
                <template v-else>
                  <span>Enviar mensaje</span>
                  <span class="group-hover:translate-x-1 transition-transform duration-300">→</span>
                </template>
              </button>

              <p class="font-mono text-[9px] text-[#AAAAAA] uppercase tracking-[0.1em] text-center">
                Respondemos en menos de 24 horas · Sin spam
              </p>

            </form>
          </div>

          <!-- COLUMNA DERECHA · Info de contacto + decoración cyberpunk -->
          <div class="relative bg-[#171717] p-8 md:p-12 lg:p-14 flex flex-col justify-between overflow-hidden">
            
            <!-- SVG cyberpunk decorativo -->
            <svg
              aria-hidden="true"
              class="absolute -bottom-20 -right-20 w-80 h-80 opacity-[0.06]"
              viewBox="0 0 400 400"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <circle cx="200" cy="200" r="180" stroke="#D4A017" stroke-width="2" />
              <circle cx="200" cy="200" r="130" stroke="#D4A017" stroke-width="1" />
              <circle cx="200" cy="200" r="80" stroke="#D4A017" stroke-width="3" />
              <path d="M20 200H380" stroke="#D4A017" stroke-width="1" />
              <path d="M200 20V380" stroke="#D4A017" stroke-width="1" />
              <path d="M73 73L327 327" stroke="#D4A017" stroke-width="1" />
              <path d="M327 73L73 327" stroke="#D4A017" stroke-width="1" />
            </svg>

            <!-- Info -->
            <div class="relative">
              <p class="font-mono text-[10px] text-[#D4A017]/50 uppercase tracking-[0.25em] mb-8">
                ─── Canales directos
              </p>

              <div class="space-y-8">
                <div
                  v-for="info in infoContacto"
                  :key="info.titulo"
                  class="flex items-start gap-4 group"
                >
                  <!-- SVG icon -->
                  <svg
                    class="w-6 h-6 shrink-0 mt-0.5"
                    viewBox="0 0 48 48"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path :d="info.svgPath" stroke="#D4A017" stroke-width="2" fill="none" />
                  </svg>

                  <div>
                    <h4 class="font-mono text-[10px] text-[#D4A017] uppercase tracking-[0.2em] mb-1">
                      {{ info.titulo }}
                    </h4>
                    <component
                      :is="info.link ? 'a' : 'span'"
                      :href="info.link"
                      :class="info.link ? 'hover:text-[#D4A017] transition-colors duration-300 cursor-pointer' : ''"
                      class="block font-display text-[18px] text-[#FAFAFA] font-semibold mb-1"
                    >
                      {{ info.valor }}
                    </component>
                    <p class="font-body text-[13px] text-[#FAFAFA]/40 leading-relaxed">
                      {{ info.desc }}
                    </p>
                  </div>
                </div>
              </div>
            </div>

            <!-- Horarios -->
            <div class="relative mt-12 pt-8 border-t border-[#FAFAFA]/10">
              <p class="font-mono text-[10px] text-[#D4A017]/50 uppercase tracking-[0.2em] mb-4">
                ─── Horarios
              </p>
              <div class="space-y-2.5">
                <div
                  v-for="h in horarios"
                  :key="h.dia"
                  class="flex items-center justify-between font-mono text-[12px]"
                >
                  <span class="text-[#FAFAFA]/50">{{ h.dia }}</span>
                  <span class="text-[#FAFAFA]/80 tracking-[0.1em]">{{ h.hora }}</span>
                </div>
              </div>
            </div>

            <!-- Detalle decorativo inferior -->
            <div class="relative mt-12 flex gap-1.5">
              <div class="h-1 w-8 bg-[#D4A017]/40" />
              <div class="h-1 w-4 bg-[#D4A017]/60" />
              <div class="h-1 w-12 bg-[#D4A017]/20" />
              <div class="h-1 w-2 bg-[#D4A017]/80" />
            </div>

          </div>

        </div>

      </div>
    </section>

  </div>
</template>
