<script setup lang="ts">
import { computed } from "vue"

const props = withDefaults(
  defineProps<{
    rango: "bronce" | "plata" | "oro" | "platino"
    size?: "sm" | "md" | "lg"
    showLabel?: boolean
  }>(),
  {
    size: "md",
    showLabel: true,
  },
)

const config = computed(() => {
  const map = {
    bronce: {
      color: "bg-amber-600",
      text: "text-amber-600",
      border: "border-amber-600",
      glow: "shadow-amber-600/30",
      icon: "🥉",
    },
    plata: {
      color: "bg-gray-400",
      text: "text-gray-400",
      border: "border-gray-400",
      glow: "shadow-gray-400/30",
      icon: "🥈",
    },
    oro: {
      color: "bg-yellow-500",
      text: "text-yellow-500",
      border: "border-yellow-500",
      glow: "shadow-yellow-500/30",
      icon: "🥇",
    },
    platino: {
      color: "bg-gradient-to-r from-cyan-400 to-purple-500",
      text: "text-purple-500",
      border: "border-purple-500",
      glow: "shadow-purple-500/30",
      icon: "💎",
    },
  }
  return map[props.rango]
})

const sizes = {
  sm: "px-2 py-0.5 text-caption gap-1",
  md: "px-3 py-1 text-small gap-1.5",
  lg: "px-4 py-1.5 text-body gap-2",
}
</script>

<template>
  <span
    :class="[
      'inline-flex items-center rounded-full border-2 font-display font-bold uppercase tracking-wide',
      config.border,
      sizes[size],
      `shadow-sm ${config.glow}`,
    ]"
    :aria-label="`Rango ${rango}`"
  >
    <span aria-hidden="true" class="text-sm leading-none">{{ config.icon }}</span>
    <span v-if="showLabel" :class="config.text">{{ rango }}</span>
  </span>
</template>
