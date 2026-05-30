<script setup lang="ts">
definePageMeta({ layout: "default" })

import { useContactoForm } from "~/features/contacto/composables/useContactoForm"

const {
  formulario,
  formError,
  submitted,
  submitting,
  copied,
  isReady,
  canSubmit,
  contactCards,
  horarios,
  responseStats,
  motivoOptions,
  whatsappUrl,
  copyPhoneNumber,
  markTouched,
  sanitizePhone,
  fieldStatus,
  enviarFormulario,
} = useContactoForm()
</script>

<template>
  <div class="contact-shell relative overflow-hidden bg-[#F4ECDD] text-[#181614]">
    <div class="contact-grain pointer-events-none absolute inset-0" aria-hidden="true" />
    <div class="pointer-events-none absolute inset-x-0 top-0 h-[32rem] bg-[radial-gradient(circle_at_top,_rgba(212,160,23,0.20),_transparent_58%)]" aria-hidden="true" />

    <main class="relative mx-auto max-w-[1440px] px-4 pb-16 pt-6 sm:px-6 md:px-8 md:pb-28 md:pt-10">
      <section class="grid gap-5 lg:grid-cols-[1.15fr_.85fr]">
        <ContactoHero />
        <ContactoSignalPanel :response-stats="responseStats" :is-ready="isReady" />
      </section>

      <section class="mt-5 grid gap-5 lg:grid-cols-[.9fr_1.1fr]">
        <div class="space-y-5">
          <ContactoDirectChannels
            :contact-cards="contactCards"
            :copied="copied"
            @copy-phone="copyPhoneNumber"
          />
          <ContactoHoursCard :horarios="horarios" />
        </div>

        <ContactoForm
          v-model:nombre="formulario.nombre"
          v-model:email="formulario.email"
          v-model:telefono="formulario.telefono"
          v-model:motivo="formulario.motivo"
          v-model:mensaje="formulario.mensaje"
          :motivo-options="motivoOptions"
          :form-error="formError"
          :submitted="submitted"
          :submitting="submitting"
          :can-submit="canSubmit"
          :whatsapp-url="whatsappUrl"
          :field-status="fieldStatus"
          @submit-form="enviarFormulario"
          @touch-field="markTouched"
          @sanitize-phone="sanitizePhone"
        />
      </section>
    </main>
  </div>
</template>

<style>
.contact-shell {
  min-height: 100dvh;
}

.contact-grain {
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.18), rgba(255, 255, 255, 0.18)),
    radial-gradient(rgba(24, 22, 20, 0.05) 0.75px, transparent 0.75px);
  background-size: 100% 100%, 6px 6px;
  mix-blend-mode: multiply;
  opacity: 0.24;
}

.panel-appear {
  opacity: 0;
  transform: translateY(18px) scale(0.985);
  transition: opacity 600ms ease, transform 600ms ease;
}

.panel-appear.is-visible {
  opacity: 1;
  transform: translateY(0) scale(1);
}

.panel-appear.delay-1 { transition-delay: 100ms; }
.panel-appear.delay-2 { transition-delay: 180ms; }
.panel-appear.delay-3 { transition-delay: 260ms; }

.signal-orbit { perspective: 1200px; }

.orbit {
  position: absolute;
  inset: 50% auto auto 50%;
  border-radius: 999px;
  border: 1px solid rgba(215, 178, 90, 0.24);
  transform: translate(-50%, -50%);
}

.orbit-1 { width: 100%; height: 100%; animation: pulse-ring 7s ease-in-out infinite; }
.orbit-2 { width: 72%; height: 72%; animation: pulse-ring 7s ease-in-out infinite 1.2s; }
.orbit-3 { width: 44%; height: 44%; animation: pulse-ring 7s ease-in-out infinite 2.4s; }

.orbit-core {
  position: absolute;
  inset: 50% auto auto 50%;
  display: flex;
  min-height: 132px;
  min-width: 132px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  border: 1px solid rgba(255, 249, 241, 0.1);
  background: rgba(255, 249, 241, 0.05);
  box-shadow: inset 0 0 0 1px rgba(255, 249, 241, 0.05), 0 24px 80px -40px rgba(0, 0, 0, 0.55);
  backdrop-filter: blur(12px);
  transform: translate(-50%, -50%);
}

.orbit-node {
  position: absolute;
  width: 13px;
  height: 13px;
  border-radius: 999px;
  background: #d7b25a;
  box-shadow: 0 0 0 8px rgba(215, 178, 90, 0.14);
}

.node-a { top: 12%; left: 20%; animation: drift-a 8s ease-in-out infinite; }
.node-b { top: 20%; right: 16%; animation: drift-b 9s ease-in-out infinite; }
.node-c { right: 24%; bottom: 15%; animation: drift-c 10s ease-in-out infinite; }

@keyframes pulse-ring {
  0%, 100% { transform: translate(-50%, -50%) scale(0.98); opacity: 0.55; }
  50% { transform: translate(-50%, -50%) scale(1.02); opacity: 1; }
}

@keyframes drift-a {
  0%, 100% { transform: translate(0, 0); }
  50% { transform: translate(10px, 12px); }
}

@keyframes drift-b {
  0%, 100% { transform: translate(0, 0); }
  50% { transform: translate(-8px, 10px); }
}

@keyframes drift-c {
  0%, 100% { transform: translate(0, 0); }
  50% { transform: translate(-10px, -8px); }
}

@media (prefers-reduced-motion: reduce) {
  .panel-appear,
  .orbit-1,
  .orbit-2,
  .orbit-3,
  .node-a,
  .node-b,
  .node-c {
    transition: none !important;
    animation: none !important;
  }

  .panel-appear {
    opacity: 1;
    transform: none;
  }
}
</style>
