<script setup lang="ts">

const props = withDefaults(
  defineProps<{
    modelValue: string
    label: string
    type?: "text" | "email" | "password" | "number" | "tel" | "search"
    placeholder?: string
    error?: string
    hint?: string
    required?: boolean
    disabled?: boolean
    icon?: "search" | "user" | "lock" | "none"
    fullWidth?: boolean
  }>(),
  {
    type: "text",
    placeholder: "",
    error: "",
    hint: "",
    required: false,
    disabled: false,
    icon: "none",
    fullWidth: true,
  },
)

const emit = defineEmits<{
  "update:modelValue": [value: string]
}>()

const id = useId()

const inputClasses = computed(() =>
  [
    "w-full rounded-lg border-2 bg-white px-4 py-2.5 font-body text-body text-text-primary",
    "placeholder:text-text-muted",
    "transition-all duration-300 ease-out",
    "min-h-[44px]", // Touch target
    props.error
      ? "border-status-error focus:border-status-error focus:ring-2 focus:ring-status-error/30"
      : "border-border-base focus:border-brand-secondary focus:ring-2 focus:ring-brand-secondary/20",
    props.disabled ? "bg-surface opacity-60 cursor-not-allowed" : "",
    props.icon !== "none" && props.icon === "search" ? "pl-10" : "",
  ].join(" "),
)

const iconComponent = computed(() => {
  switch (props.icon) {
    case "search": return "lucide-search"
    case "user": return "lucide-user"
    case "lock": return "lucide-lock"
    default: return null
  }
})
</script>

<template>
  <div :class="fullWidth ? 'w-full' : ''">
    <!-- Label -->
    <label
      v-if="label"
      :for="id"
      class="block mb-2 font-body text-small font-semibold text-text-primary"
    >
      {{ label }}
      <span v-if="required" class="text-status-error ml-0.5" aria-hidden="true">*</span>
    </label>

    <!-- Input wrapper -->
    <div class="relative">
      <!-- Icono izquierdo (Lucide inline SVG) -->
      <span
        v-if="icon === 'search'"
        class="absolute left-3 top-1/2 -translate-y-1/2 text-text-muted pointer-events-none"
        aria-hidden="true"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
      </span>

      <input
        :id="id"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :aria-invalid="!!error"
        :aria-describedby="error ? `${id}-error` : hint ? `${id}-hint` : undefined"
        :class="inputClasses"
        @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      />
    </div>

    <!-- Mensaje de error -->
    <p
      v-if="error"
      :id="`${id}-error`"
      class="mt-1.5 text-small text-status-error font-medium"
      role="alert"
    >
      {{ error }}
    </p>

    <!-- Hint -->
    <p
      v-else-if="hint"
      :id="`${id}-hint`"
      class="mt-1.5 text-caption text-text-muted"
    >
      {{ hint }}
    </p>
  </div>
</template>
