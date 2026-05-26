<script setup lang="ts">
defineProps<{
  show: boolean
  title?: string
  size?: "sm" | "md" | "lg"
}>()

const emit = defineEmits<{
  close: []
}>()

const sizes = {
  sm: "max-w-md",
  md: "max-w-lg",
  lg: "max-w-2xl",
}
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="show"
        class="fixed inset-0 z-modal flex items-center justify-center p-4"
        role="dialog"
        aria-modal="true"
        :aria-label="title || 'Modal'"
        @click.self="emit('close')"
      >
        <!-- Overlay -->
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" />

        <!-- Contenido -->
        <div
          :class="[
            'relative bg-white rounded-xl border border-border-base shadow-lg w-full',
            sizes[size || 'md'],
            'max-h-[90vh] overflow-y-auto',
          ]"
        >
          <!-- Header -->
          <div
            v-if="title || $slots.header"
            class="flex items-center justify-between px-6 py-4 border-b border-border-base"
          >
            <h2 class="font-body text-heading font-bold text-text-primary">
              <slot name="header">{{ title }}</slot>
            </h2>
            <button
              class="p-1.5 rounded-lg hover:bg-surface transition-colors duration-300 min-h-[44px] min-w-[44px] flex items-center justify-center"
              aria-label="Cerrar modal"
              @click="emit('close')"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
            </button>
          </div>

          <!-- Body -->
          <div class="px-6 py-4">
            <slot />
          </div>

          <!-- Footer -->
          <div
            v-if="$slots.footer"
            class="flex justify-end gap-3 px-6 py-4 border-t border-border-base"
          >
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 300ms ease;
}
.modal-enter-active > div:nth-child(2),
.modal-leave-active > div:nth-child(2) {
  transition: transform 300ms ease, opacity 300ms ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-from > div:nth-child(2) {
  transform: scale(0.95) translateY(10px);
  opacity: 0;
}
.modal-leave-to > div:nth-child(2) {
  transform: scale(0.95) translateY(10px);
  opacity: 0;
}
</style>
