<script setup lang="ts">
import { computed } from "vue"

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

const emit = defineEmits<{
  click: [e: MouseEvent]
}>()

const baseClasses = computed(() => {
  const variants = {
    primary:
      "bg-brand-primary text-text-primary border-2 border-brand-primary hover:shadow-brand",
    secondary:
      "bg-brand-secondary text-white border-2 border-brand-secondary hover:brightness-110",
    outline:
      "bg-transparent text-text-primary border-2 border-border-base hover:border-brand-primary hover:text-brand-primary",
    ghost:
      "bg-transparent text-text-muted border-2 border-transparent hover:text-text-primary hover:bg-surface",
  }

  const sizes = {
    sm: "px-4 py-1.5 text-small",
    md: "px-6 py-2.5 text-body",
    lg: "px-8 py-3.5 text-subheading",
  }

  return [
    "inline-flex items-center justify-center gap-2 rounded-lg font-body font-semibold",
    "transition-all duration-500 ease-out",
    "focus-visible:outline-2 focus-visible:outline-brand-secondary focus-visible:outline-offset-2",
    "disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:shadow-none",
    "active:scale-[0.98]",
    "min-h-[44px]", // Touch target mínimo
    variants[props.variant],
    sizes[props.size],
    props.fullWidth ? "w-full" : "",
  ].join(" ")
})
</script>

<template>
  <button
    :type="type"
    :class="baseClasses"
    :disabled="disabled || loading"
    @click="emit('click', $event)"
  >
    <!-- Spinner de carga -->
    <svg
      v-if="loading"
      class="animate-spin h-5 w-5"
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle
        class="opacity-25"
        cx="12" cy="12" r="10"
        stroke="currentColor"
        stroke-width="4"
      />
      <path
        class="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
      />
    </svg>
    <slot />
  </button>
</template>
