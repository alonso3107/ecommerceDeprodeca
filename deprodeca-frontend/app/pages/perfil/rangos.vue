<script setup lang="ts">
definePageMeta({ layout: "default", middleware: "auth" })

await navigateTo("/perfil", { replace: true })
</script>

<template>
  <div class="flex items-center justify-center min-h-[60vh]">
    <p class="font-mono text-[11px] text-texto-muted uppercase tracking-[0.2em]">
      Redirigiendo...
    </p>
  </div>
</template>
