<script setup lang="ts">

const props = withDefaults(
  defineProps<{
    message: string
    type?: "success" | "error" | "warning" | "info"
    show: boolean
    duration?: number
  }>(),
  {
    type: "info",
    duration: 4000,
  },
)

const emit = defineEmits<{
  close: []
}>()

// Auto-dismiss
if (props.show && props.duration > 0) {
  setTimeout(() => emit("close"), props.duration)
}

const config = computed(() => {
  const map = {
    success: {
      bg: "bg-status-success/10 border-status-success",
      text: "text-status-success",
      icon: "M5 13l4 4L19 7",
    },
    error: {
      bg: "bg-status-error/10 border-status-error",
      text: "text-status-error",
      icon: "M18 6 6 18M6 6l12 12",
    },
    warning: {
      bg: "bg-status-warning/10 border-status-warning",
      text: "text-status-warning",
      icon: "M12 9v4M12 17h.01",
    },
    info: {
      bg: "bg-brand-secondary/10 border-brand-secondary",
      text: "text-brand-secondary",
      icon: "M12 16v-4M12 8h.01",
    },
  }
  return map[props.type]
})
</script>

<template>
  <Teleport to="body">
    <Transition name="toast">
      <div
        v-if="show"
        :class="[
          'fixed bottom-6 right-6 z-toast flex items-center gap-3 px-5 py-3 rounded-lg border-2 shadow-lg',
          config.bg,
        ]"
        role="alert"
      >
        <!-- Icono -->
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="20" height="20"
          viewBox="0 0 24 24"
          fill="none"
          :stroke="config.text"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          :class="config.text"
          aria-hidden="true"
        >
          <path :d="config.icon" />
        </svg>

        <span class="font-body text-small font-medium text-text-primary">{{ message }}</span>

        <button
          class="ml-2 p-1 rounded hover:bg-black/10 transition-colors min-h-[44px] min-w-[44px] flex items-center justify-center"
          aria-label="Cerrar notificación"
          @click="emit('close')"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
        </button>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.toast-enter-active {
  transition: all 400ms ease-out;
}
.toast-leave-active {
  transition: all 250ms ease-in;
}
.toast-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.95);
}
.toast-leave-to {
  opacity: 0;
  transform: translateY(10px) scale(0.95);
}
</style>
