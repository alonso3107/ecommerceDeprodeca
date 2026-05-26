<script setup lang="ts">
const error = useError()

function handleError() {
  clearError({ redirect: "/" })
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-white p-4">
    <div class="text-center">
      <span class="font-display text-display-xl text-text-primary leading-none">
        {{ error?.statusCode || "Error" }}
      </span>
      <h1 class="font-body text-heading text-text-primary mt-4 mb-2">
        {{ error?.statusMessage || "Algo salió mal" }}
      </h1>
      <p class="font-body text-small text-text-muted mb-8 max-w-md">
        {{ error?.message || "No pudimos encontrar lo que buscas." }}
      </p>
      <Button variant="primary" size="md" @click="handleError">
        Volver al Inicio
      </Button>
    </div>
  </div>
</template>
