<script setup lang="ts">
type CampoFormulario = "nombre" | "email" | "telefono" | "mensaje"

defineProps<{
  nombre: string
  email: string
  telefono: string
  motivo: string
  mensaje: string
  motivoOptions: readonly string[]
  fieldStatus: (field: CampoFormulario) => 'idle' | 'valid' | 'error'
  formError: string
  submitted: boolean
  submitting: boolean
  canSubmit: boolean
  whatsappUrl: string
}>()

const emit = defineEmits<{
  "update:nombre": [value: string]
  "update:email": [value: string]
  "update:telefono": [value: string]
  "update:motivo": [value: string]
  "update:mensaje": [value: string]
  "submit-form": []
  "touch-field": [field: CampoFormulario]
  "sanitize-phone": [event: Event]
}>()
</script>

<template>
  <section class="panel-appear rounded-[32px] border border-[#E5D8C2] bg-[#FFF9F1] p-6 shadow-[0_26px_90px_-60px_rgba(24,22,20,0.55)] md:p-8 lg:p-10" :class="'is-visible delay-2'">
    <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#A16207]">Formulario</p>
    <div class="mt-4 max-w-[18ch]">
      <h2 class="font-display text-[clamp(2rem,4vw,3rem)] font-black uppercase leading-[0.92] text-[#181614]">
        Envíanos tu consulta.
      </h2>
    </div>
    <p class="mt-4 max-w-[52ch] text-[15px] leading-relaxed text-[#5B534A] md:text-[16px]">
      Respondemos de forma clara y breve. Si querés registrarte como bodega, pedir soporte o hacer seguimiento, este es el canal correcto.
    </p>

    <transition name="fade-slide">
      <div v-if="submitted" class="mt-6 rounded-[22px] border border-[#CFA24F] bg-[#FFF6DD] p-5">
        <p class="font-mono text-[11px] uppercase tracking-[0.22em] text-[#A16207]">Mensaje enviado</p>
        <p class="mt-2 text-[15px] leading-relaxed text-[#3E3428]">
          Gracias. Tu consulta quedó registrada y te responderemos lo antes posible.
        </p>
      </div>
    </transition>

    <transition name="fade-slide">
      <div v-if="formError" class="mt-6 rounded-[22px] border border-[#D92D20]/20 bg-[#FEF3F2] p-5 text-[#9B1C1C]">
        <p class="font-mono text-[11px] uppercase tracking-[0.22em]">Atención</p>
        <p class="mt-2 text-[15px] leading-relaxed">{{ formError }}</p>
      </div>
    </transition>

    <form class="mt-8 space-y-5" @submit.prevent="emit('submit-form')">
      <div class="grid gap-5 md:grid-cols-2">
        <label class="block">
          <span class="mb-2 flex items-center justify-between gap-3 font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">
            <span>Nombre *</span>
            <span v-if="fieldStatus('nombre') === 'error'" class="text-[#B42318] normal-case tracking-normal">{{ nombre.trim().length === 0 ? 'El nombre es obligatorio' : 'Mínimo 2 caracteres' }}</span>
            <span v-else-if="fieldStatus('nombre') === 'valid'" class="text-[#0E9F6E]">OK</span>
          </span>
          <input
            :value="nombre"
            type="text"
            placeholder="María López"
            class="field-input"
            :class="fieldStatus('nombre')"
            @input="emit('update:nombre', ($event.target as HTMLInputElement).value)"
            @blur="emit('touch-field', 'nombre')"
          />
        </label>

        <label class="block">
          <span class="mb-2 flex items-center justify-between gap-3 font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">
            <span>Email *</span>
            <span v-if="fieldStatus('email') === 'error'" class="text-[#B42318] normal-case tracking-normal">{{ email.trim().length === 0 ? 'El email es obligatorio' : 'Formato de email inválido' }}</span>
            <span v-else-if="fieldStatus('email') === 'valid'" class="text-[#0E9F6E]">OK</span>
          </span>
          <input
            :value="email"
            type="email"
            placeholder="maria@bodega.pe"
            class="field-input"
            :class="fieldStatus('email')"
            @input="emit('update:email', ($event.target as HTMLInputElement).value)"
            @blur="emit('touch-field', 'email')"
          />
        </label>
      </div>

      <label class="block">
        <span class="mb-2 flex items-center justify-between gap-3 font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">
          <span>Teléfono (WhatsApp)</span>
          <span v-if="fieldStatus('telefono') === 'error'" class="text-[#B42318] normal-case tracking-normal">Usa 9 dígitos, sin espacios</span>
          <span v-else-if="fieldStatus('telefono') === 'valid'" class="text-[#0E9F6E]">OK</span>
        </span>
        <div class="relative">
          <span class="pointer-events-none absolute left-4 top-1/2 -translate-y-1/2 font-mono text-[13px] text-[#7A7369]">+51</span>
          <input
            :value="telefono"
            type="tel"
            inputmode="numeric"
            maxlength="9"
            placeholder="900000000"
            class="field-input pl-14 font-mono tracking-[0.08em]"
            :class="fieldStatus('telefono')"
            @input="emit('sanitize-phone', $event)"
            @blur="emit('touch-field', 'telefono')"
          />
        </div>
      </label>

      <label class="block">
        <span class="mb-2 block font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">Motivo de contacto</span>
        <div class="relative">
          <select
            :value="motivo"
            class="field-input appearance-none pr-11"
            @change="emit('update:motivo', ($event.target as HTMLSelectElement).value)"
          >
            <option value="">Seleccioná un motivo</option>
            <option v-for="m in motivoOptions" :key="m" :value="m">{{ m }}</option>
          </select>
          <svg class="pointer-events-none absolute right-4 top-1/2 h-4 w-4 -translate-y-1/2 text-[#5B534A]" viewBox="0 0 20 20" fill="none" aria-hidden="true">
            <path d="m5 8 5 5 5-5" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </div>
      </label>

      <label class="block">
        <span class="mb-2 flex items-center justify-between gap-3 font-mono text-[11px] uppercase tracking-[0.18em] text-[#5B534A]">
          <span>Mensaje *</span>
          <span v-if="fieldStatus('mensaje') === 'error'" class="text-[#B42318] normal-case tracking-normal">{{ mensaje.trim().length === 0 ? 'El mensaje es obligatorio' : 'Mínimo 10 caracteres' }}</span>
          <span v-else-if="fieldStatus('mensaje') === 'valid'" class="text-[#0E9F6E]">OK</span>
        </span>
        <textarea
          :value="mensaje"
          rows="6"
          placeholder="Contanos qué necesitás y desde qué zona escribís."
          class="field-input min-h-[160px] resize-none py-4 leading-relaxed"
          :class="fieldStatus('mensaje')"
          @input="emit('update:mensaje', ($event.target as HTMLTextAreaElement).value)"
          @blur="emit('touch-field', 'mensaje')"
        />
      </label>

      <div class="flex flex-col gap-3 sm:flex-row">
        <button
          type="submit"
          :disabled="submitting || !canSubmit"
          class="group inline-flex min-h-[52px] w-full items-center justify-center gap-3 rounded-full bg-[#181614] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.22em] text-[#FFF9F1] transition-all duration-300 hover:-translate-y-0.5 hover:bg-[#A16207] disabled:cursor-not-allowed disabled:opacity-60"
        >
          <span v-if="submitting" class="h-4 w-4 animate-spin rounded-full border-2 border-[#FFF9F1] border-t-transparent" />
          <span v-else>{{ submitted ? 'Enviado' : 'Enviar mensaje' }}</span>
          <span aria-hidden="true" class="transition-transform duration-300 group-hover:translate-x-1">→</span>
        </button>

        <a
          :href="whatsappUrl"
          class="inline-flex min-h-[52px] items-center justify-center rounded-full border border-[#D6C7B1] px-6 py-3 text-center font-mono text-[11px] uppercase tracking-[0.22em] text-[#181614] transition-colors duration-300 hover:border-[#181614] hover:bg-[#181614] hover:text-[#FFF9F1]"
        >
          WhatsApp directo
        </a>
      </div>

      <p class="text-center font-mono text-[9px] uppercase tracking-[0.16em] text-[#7A7369]">
        Respuesta en menos de 24 horas · Sin spam
      </p>
    </form>
  </section>
</template>

<style scoped>
.field-input {
  width: 100%;
  min-height: 52px;
  border: 1px solid #d8ccb8;
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.95);
  padding: 0.95rem 1rem;
  color: #181614;
  outline: none;
  transition: border-color 180ms ease, box-shadow 180ms ease, background-color 180ms ease, transform 180ms ease;
}

.field-input::placeholder {
  color: #aba396;
}

.field-input:hover {
  border-color: #c8b498;
}

.field-input:focus {
  border-color: #a16207;
  box-shadow: 0 0 0 4px rgba(161, 98, 7, 0.12);
}

.field-input.idle {
  background: rgba(255, 255, 255, 0.95);
}

.field-input.error {
  border-color: rgba(180, 35, 24, 0.55);
  background: rgba(254, 243, 242, 0.95);
}

.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: opacity 220ms ease, transform 220ms ease;
}

.fade-slide-enter-from,
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
</style>
