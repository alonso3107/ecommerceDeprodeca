<!--
  ═══════════════════════════════════════════════════════════════
  auth.vue — Layout para login/registro · Split-Screen Brutalista
  Izquierda: panel negro con branding, patrón geométrico, tagline
  Derecha: formulario blanco con bordes duros y navegación
  ═══════════════════════════════════════════════════════════════
-->
<template>
  <div class="min-h-screen flex bg-fondo">

    <!-- ═══ PANEL IZQUIERDO · Branding brutalista ═══ -->
    <div class="hidden lg:flex w-[42%] bg-texto relative overflow-hidden flex-col justify-between p-12">

      <!-- Patrón geométrico de fondo (grid de puntos) -->
      <div class="absolute inset-0 opacity-[0.04]"
           style="background-image: radial-gradient(circle, #fff 1px, transparent 1px); background-size: 24px 24px;"
           aria-hidden="true" />

      <!-- Diagonal decorativa -->
      <div class="absolute -top-20 -right-20 w-80 h-80 border border-[#D4A017]/20 rotate-45" aria-hidden="true" />
      <div class="absolute -bottom-32 -left-32 w-96 h-96 border border-[#D4A017]/10 rotate-45" aria-hidden="true" />

      <!-- Logo grande -->
      <div class="relative z-10">
        <NuxtLink to="/" class="inline-block mb-6">
          <span class="font-display text-[clamp(3rem,5vw,4rem)] text-white leading-none tracking-tight">
            DEPRO<span class="text-[#D4A017]">DECA</span>
          </span>
        </NuxtLink>
      </div>

      <!-- Tagline brutalista -->
      <div class="relative z-10">
        <div class="border-l-2 border-[#D4A017] pl-5 mb-6">
          <p class="font-display text-display-md text-white uppercase leading-[1.1] mb-3">
            Tu bodega<br />a otro nivel<span class="text-[#D4A017]">.</span>
          </p>
          <p class="font-mono text-[11px] text-white/40 uppercase tracking-[0.2em] leading-relaxed">
            Filial Oficial Nestlé Perú<br />
            Distribución B2B · Mayorista
          </p>
        </div>

        <!-- Métricas rápidas -->
        <div class="flex gap-8 mt-10">
          <div>
            <p class="font-display text-[2rem] text-[#D4A017] leading-none">500<span class="text-white/30">+</span></p>
            <p class="font-mono text-[9px] text-white/30 uppercase tracking-[0.2em] mt-1">Productos</p>
          </div>
          <div>
            <p class="font-display text-[2rem] text-[#D4A017] leading-none">200<span class="text-white/30">+</span></p>
            <p class="font-mono text-[9px] text-white/30 uppercase tracking-[0.2em] mt-1">Bodegas</p>
          </div>
          <div>
            <p class="font-display text-[2rem] text-[#D4A017] leading-none">4</p>
            <p class="font-mono text-[9px] text-white/30 uppercase tracking-[0.2em] mt-1">Rangos</p>
          </div>
        </div>
      </div>

      <!-- Footer del panel -->
      <p class="relative z-10 font-mono text-[9px] text-white/15 uppercase tracking-[0.3em]">
        © 2026 DEPRODECA — Ica, Perú
      </p>
    </div>

    <!-- ═══ PANEL DERECHO · Formulario ═══ -->
    <div class="flex-1 flex items-center justify-center p-6 md:p-12 relative">

      <!-- Link volver al inicio (top-left) -->
      <NuxtLink
        to="/"
        class="absolute top-6 left-6 md:top-8 md:left-8 font-mono text-[10px] text-texto-muted
               hover:text-[#D4A017] uppercase tracking-[0.2em] transition-colors duration-200
               flex items-center gap-1.5"
      >
        <!-- Flecha geométrica -->
        <svg width="12" height="10" viewBox="0 0 12 10" fill="none">
          <path d="M11 5H1M1 5L4.5 1.5M1 5L4.5 8.5" stroke="currentColor" stroke-width="1.5"
                stroke-linecap="square" stroke-linejoin="miter"/>
        </svg>
        Volver al Inicio
      </NuxtLink>

      <!-- Marcas de esquina -->
      <div class="absolute top-6 right-6 md:top-8 md:right-8 w-4 h-4 border-r-2 border-t-2 border-[#D4A017]" aria-hidden="true" />
      <div class="absolute bottom-6 right-6 md:bottom-8 md:right-8 w-4 h-4 border-r-2 border-b-2 border-[#D4A017]" aria-hidden="true" />

      <!-- Logo móvil (solo en pantallas chicas) -->
      <div class="lg:hidden absolute top-6 left-1/2 -translate-x-1/2">
        <NuxtLink to="/">
          <span class="font-display text-display-md text-texto leading-none tracking-tight">
            depro<span class="text-texto-muted font-normal">deca</span>
          </span>
        </NuxtLink>
      </div>

      <!-- Contenido del formulario -->
      <div class="w-full max-w-[420px]">
        <slot />
      </div>
    </div>

  </div>
</template>
