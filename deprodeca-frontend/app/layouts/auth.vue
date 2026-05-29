<template>
  <div class="min-h-screen flex bg-blanco">

    <div class="hidden lg:flex w-[42%] bg-crema border-r border-stone relative overflow-hidden flex-col justify-between p-12">
      <svg class="absolute -right-20 -top-24 text-dorado/5" width="280" height="280" viewBox="0 0 280 280" fill="none" aria-hidden="true">
        <circle cx="140" cy="140" r="110" stroke="currentColor" stroke-width="20" />
      </svg>

      <div class="relative z-10">
        <NuxtLink to="/" class="inline-block mb-6">
          <span class="font-display text-[clamp(3rem,5vw,4rem)] text-dorado leading-none tracking-tight">
            DEPRO<span class="text-negro">DECA</span>
          </span>
        </NuxtLink>

        <div class="border border-dorado p-4 mb-8">
          <div class="flex items-center gap-2 mb-2">
            <svg width="10" height="10" viewBox="0 0 10 10" class="text-dorado" fill="none" aria-hidden="true">
              <rect x="5" y="0" width="7" height="7" transform="rotate(45 5 0)" fill="currentColor"/>
            </svg>
            <p class="font-mono text-[10px] text-stone-oscuro uppercase tracking-[0.15em]">Filial Oficial Nestlé Perú</p>
          </div>
          <p class="font-mono text-[9px] text-stone-oscuro/50 uppercase tracking-[0.14em]">Distribución B2B Mayorista</p>
        </div>

        <div class="mb-8">
          <p class="font-display text-display-md text-negro uppercase leading-[1.1] mb-4">Tu bodega a otro nivel.</p>
          <div class="w-16 border-t border-stone/30 mb-4" aria-hidden="true" />
          <div class="space-y-2">
            <p class="font-body text-stone-oscuro italic leading-relaxed">
              <svg width="14" height="12" viewBox="0 0 14 12" class="inline-block mr-1 text-stone-oscuro" fill="none" aria-hidden="true">
                <path d="M6 1H2L1 5V11H7V5H4L6 1ZM13 1H9L8 5V11H14V5H11L13 1Z" fill="currentColor"/>
              </svg>
              El mejor servicio de distribución en todo Ica.
            </p>
            <p class="font-mono text-[9px] text-stone-oscuro/50 uppercase tracking-[0.16em]">Carlos, Bodega El Sol</p>
          </div>
        </div>

        <div class="flex gap-8 mt-10">
          <div>
            <svg width="14" height="14" viewBox="0 0 14 14" class="text-stone-oscuro/40 mb-1" fill="none" aria-hidden="true">
              <circle cx="7" cy="7" r="5.5" stroke="currentColor" />
            </svg>
            <p class="font-display text-[2rem] text-dorado leading-none">500+</p>
            <p class="font-mono text-[9px] text-stone-oscuro/50 uppercase tracking-[0.2em] mt-1">Productos</p>
          </div>
          <div>
            <svg width="14" height="14" viewBox="0 0 14 14" class="text-stone-oscuro/40 mb-1" fill="none" aria-hidden="true">
              <rect x="7" y="0.5" width="9.2" height="9.2" transform="rotate(45 7 0.5)" stroke="currentColor" />
            </svg>
            <p class="font-display text-[2rem] text-dorado leading-none">200+</p>
            <p class="font-mono text-[9px] text-stone-oscuro/50 uppercase tracking-[0.2em] mt-1">Bodegas</p>
          </div>
          <div>
            <svg width="14" height="14" viewBox="0 0 14 14" class="text-stone-oscuro/40 mb-1" fill="none" aria-hidden="true">
              <path d="M4 1.5H10L13 7L10 12.5H4L1 7L4 1.5Z" stroke="currentColor" />
            </svg>
            <p class="font-display text-[2rem] text-dorado leading-none">4</p>
            <p class="font-mono text-[9px] text-stone-oscuro/50 uppercase tracking-[0.2em] mt-1">Rangos</p>
          </div>
        </div>
      </div>

      <p class="relative z-10 font-mono text-[9px] text-stone-oscuro/30 uppercase tracking-[0.3em]">
        © 2026 DEPRODECA — Ica, Perú
      </p>
    </div>

    <div class="flex-1 flex items-center justify-center p-6 md:p-12 relative bg-blanco">

      <NuxtLink
        to="/"
        class="absolute top-6 left-6 md:top-8 md:left-8 font-mono text-[10px] text-stone-oscuro
               hover:text-dorado uppercase tracking-[0.2em] transition-colors duration-200
               flex items-center gap-1.5"
      >
        <svg width="12" height="10" viewBox="0 0 12 10" fill="none">
          <path d="M11 5H1M1 5L4.5 1.5M1 5L4.5 8.5" stroke="currentColor" stroke-width="1.5"
                stroke-linecap="square" stroke-linejoin="miter"/>
        </svg>
        Volver al Inicio
      </NuxtLink>

      <div class="w-full max-w-[420px]">
        <slot />
      </div>
    </div>

  </div>
</template>
