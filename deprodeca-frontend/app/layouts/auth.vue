<script setup lang="ts">
// Layout para páginas de login/registro.
// Sin navbar ni footer — solo el formulario centrado.
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-surface p-4">
    <div class="w-full max-w-md">
      <slot />
    </div>
  </div>
</template>
