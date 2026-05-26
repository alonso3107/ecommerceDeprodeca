<!--
  ═══════════════════════════════════════════════════════════════
  auth.vue — Layout para login/registro
  Brutalismo puro: fondo blanco, sin adornos, formulario centrado
  con bordes duros. Sin navbar ni footer.
  ═══════════════════════════════════════════════════════════════
-->
<template>
  <div class="min-h-screen flex items-center justify-center bg-fondo p-4">
    <!--
      Contenedor centrado con borde duro y decoración geométrica
      en las esquinas superiores como firma brutalista.
    -->
    <div class="w-full max-w-md relative">

      <!-- Detalles técnicos en esquinas (firma brutalista) -->
      <div class="absolute -top-2 -left-2 w-4 h-4 border-l-2 border-t-2 border-[#D4A017]" aria-hidden="true" />
      <div class="absolute -top-2 -right-2 w-4 h-4 border-r-2 border-t-2 border-[#D4A017]" aria-hidden="true" />
      <div class="absolute -bottom-2 -left-2 w-4 h-4 border-l-2 border-b-2 border-[#D4A017]" aria-hidden="true" />
      <div class="absolute -bottom-2 -right-2 w-4 h-4 border-r-2 border-b-2 border-[#D4A017]" aria-hidden="true" />

      <slot />
    </div>
  </div>
</template>
