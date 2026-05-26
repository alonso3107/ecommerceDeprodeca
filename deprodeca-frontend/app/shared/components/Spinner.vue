<script setup lang="ts">
withDefaults(
  defineProps<{
    size?: "sm" | "md" | "lg"
    label?: string
  }>(),
  {
    size: "md",
    label: "Cargando...",
  },
)

const sizes = {
  sm: "h-5 w-5 border-2",
  md: "h-8 w-8 border-2",
  lg: "h-12 w-12 border-[3px]",
}
</script>

<template>
  <div class="flex flex-col items-center gap-3" role="status" :aria-label="label">
    <div
      :class="[
        sizes[size],
        'rounded-full border-brand-primary border-t-transparent animate-spin',
      ]"
    />
    <span v-if="label" class="text-small text-text-muted font-body">{{ label }}</span>
  </div>
</template>
