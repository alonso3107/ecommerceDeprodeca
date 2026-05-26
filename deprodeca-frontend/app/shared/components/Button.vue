<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    variant?: "primary" | "secondary" | "outline" | "ghost"
    size?: "sm" | "md" | "lg"
    disabled?: boolean
    loading?: boolean
    type?: "button" | "submit"
    fullWidth?: boolean
  }>(),
  {
    variant: "primary",
    size: "md",
    disabled: false,
    loading: false,
    type: "button",
    fullWidth: false,
  },
)

const emit = defineEmits<{ click: [e: MouseEvent] }>()
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    :class="[
      'inline-flex items-center justify-center gap-2 font-body font-semibold',
      'transition-all duration-500 ease-out',
      'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-acento',
      'disabled:opacity-40 disabled:cursor-not-allowed',
      'active:scale-[0.98]',
      'min-h-[44px]',
      {
        primary: 'bg-acento text-white hover:bg-texto/90 shadow-xs hover:shadow-md rounded-xl',
        secondary: 'bg-texto text-white hover:bg-texto/80 shadow-xs hover:shadow-md rounded-xl',
        outline: 'bg-transparent text-texto border-2 border-borde hover:border-texto/30 hover:bg-fondo rounded-xl',
        ghost: 'bg-transparent text-texto-muted hover:text-texto hover:bg-fondo rounded-xl',
      }[variant],
      {
        sm: 'px-4 py-1.5 text-small rounded-lg',
        md: 'px-6 py-2.5 text-body',
        lg: 'px-8 py-3.5 text-subheading',
      }[size],
      fullWidth ? 'w-full' : '',
    ]"
    @click="emit('click', $event)"
  >
    <svg
      v-if="loading"
      class="animate-spin h-5 w-5"
      xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
    >
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
    </svg>
    <slot />
  </button>
</template>
