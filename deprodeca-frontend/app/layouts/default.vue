<script setup lang="ts">
// Layout principal con Navbar Hungry Tiger + Footer
// Feature-Sliced Design: solo layout, sin lógica de negocio.
</script>

<template>
  <div class="min-h-screen flex flex-col bg-white">
    <Navbar />
    <main class="flex-1">
      <slot />
    </main>
    <Footer />
  </div>
</template>
