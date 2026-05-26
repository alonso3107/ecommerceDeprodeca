<script setup lang="ts">
definePageMeta({ layout: "default" })

const parallaxY = ref(0)
onMounted(() => { window.addEventListener("scroll", () => { parallaxY.value = window.scrollY * 0.15 }) })
onUnmounted(() => { window.removeEventListener("scroll", () => {}) })

const categoriasDestacadas = [
  { nombre: "Bebidas", slug: "bebidas" },
  { nombre: "Snacks", slug: "snacks" },
  { nombre: "Lácteos", slug: "lacteos" },
  { nombre: "Abarrotes", slug: "abarrotes" },
  { nombre: "Limpieza", slug: "limpieza" },
  { nombre: "Cuidado Personal", slug: "cuidado-personal" },
]

const rangos = [
  { nombre: "Bronce", minimo: "S/ 0", icono: "🥉", beneficio: "Precios exclusivos B2B" },
  { nombre: "Plata", minimo: "S/ 5,000", icono: "🥈", beneficio: "Envío gratis + prioridad" },
  { nombre: "Oro", minimo: "S/ 20,000", icono: "🥇", beneficio: "5% descuento en volumen" },
  { nombre: "Platino", minimo: "S/ 50,000", icono: "💎", beneficio: "Acceso anticipado + soporte VIP" },
]

const beneficios = [
  { titulo: "Precios B2B", desc: "Sin intermediarios. Precios directos de Nestlé para tu bodega.", icono: "M12 2v20M17 5H9.5a3.5 3.5 0 000 7h5a3.5 3.5 0 010 7H6" },
  { titulo: "Pedidos al Mayor", desc: "Comprá por caja o unidad. Stock siempre disponible.", icono: "M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" },
  { titulo: "Entrega 24-48h", desc: "Despacho rápido a todo Ica. Seguimiento en vivo.", icono: "M13 10V3L4 14h7v7l9-11h-7z" },
  { titulo: "Gamificación", desc: "Subí de Bronce a Platino. Cada compra te da puntos y beneficios.", icono: "M12 15l-4.24-4.24a6 6 0 0112.48 0L20 15" },
]
</script>

<template>
  <div>

    <!-- ═══════════════════════════════════════════════════
         HERO MAXIMALISTA — Dorado masivo, parallax, datos
         ═══════════════════════════════════════════════════ -->
    <section class="relative bg-[#D4A017] overflow-hidden min-h-[90vh] flex items-center">
      <div class="absolute inset-0 opacity-[0.04] bg-[radial-gradient(circle_at_50%_50%,#000_1px,transparent_1px)] bg-[length:20px_20px]" aria-hidden="true" />

      <div class="relative w-full max-w-[1280px] mx-auto px-6 md:px-8 py-24 md:py-0">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 lg:gap-24 items-center">

          <!-- LEFT: Texto masivo -->
          <div class="hero-text max-w-[560px]">
            <span class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-black/20 text-white font-body text-caption font-bold tracking-wider uppercase mb-8 backdrop-blur-sm">
              <span class="w-2 h-2 rounded-full bg-white animate-pulse" />
              Filial Oficial — Nestlé Perú
            </span>

            <h1 class="font-display text-[clamp(4rem,10vw,7rem)] text-white leading-[0.88] tracking-[-0.03em] uppercase mb-3">
              Tu bodega
            </h1>
            <h2 class="font-display text-[clamp(3.5rem,8vw,5.5rem)] text-black/20 leading-[0.9] tracking-[-0.02em] uppercase mb-8">
              a otro nivel
            </h2>

            <p class="font-body text-body text-white/85 max-w-[420px] mb-10 leading-relaxed text-[1.05rem]">
              La plataforma B2B que conecta tu bodega directamente con Nestlé.
              Precios exclusivos, pedidos al por mayor, entregas en 24-48h y un
              sistema de rangos que premia tu crecimiento.
            </p>

            <div class="flex flex-col sm:flex-row gap-4 mb-16">
              <button class="inline-flex items-center justify-center gap-2 px-8 py-4 rounded-2xl bg-black text-white font-body text-body font-bold hover:bg-black/90 transition-all duration-500 hover:-translate-y-[2px] hover:shadow-xl active:scale-[0.98] min-h-[54px]" @click="navigateTo('/catalogo')">
                Explorar Catálogo
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
              </button>
              <button class="inline-flex items-center justify-center gap-2 px-8 py-4 rounded-2xl bg-transparent text-black font-body text-body font-bold border-2 border-black/20 hover:border-black hover:bg-black/5 transition-all duration-500 min-h-[54px]" @click="navigateTo('/auth/registro')">
                Crear Cuenta Gratis
              </button>
            </div>

            <div class="flex gap-12">
              <div><p class="font-display text-[3rem] text-black leading-none">500<span class="text-white">+</span></p><p class="font-body text-caption text-black/60 mt-1 uppercase tracking-wider">Productos</p></div>
              <div><p class="font-display text-[3rem] text-black leading-none">200<span class="text-white">+</span></p><p class="font-body text-caption text-black/60 mt-1 uppercase tracking-wider">Bodegas</p></div>
              <div><p class="font-display text-[3rem] text-black leading-none">4</p><p class="font-body text-caption text-black/60 mt-1 uppercase tracking-wider">Rangos</p></div>
            </div>
          </div>

          <!-- RIGHT: Imagen con parallax -->
          <div class="relative hidden lg:block">
            <div class="relative rounded-3xl overflow-hidden shadow-2xl" :style="{ transform: `translateY(${parallaxY}px)` }">
              <img src="https://images.unsplash.com/photo-1543168256-418811576931?w=800&q=80" alt="Productos Nestlé para bodegas" class="w-full h-[620px] object-cover" />
              <div class="absolute inset-0 bg-gradient-to-t from-black/30 to-transparent" />
            </div>
            <div class="absolute -bottom-8 -left-8 bg-black rounded-2xl shadow-2xl p-5 flex items-center gap-4">
              <div class="w-14 h-14 rounded-xl bg-[#D4A017]/20 flex items-center justify-center flex-shrink-0">
                <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#D4A017" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20M17 5H9.5a3.5 3.5 0 000 7h5a3.5 3.5 0 010 7H6"/></svg>
              </div>
              <div><p class="font-body text-caption text-white/50">Precio mayorista</p><p class="font-display text-heading text-white leading-none">S/ 1.20</p></div>
            </div>
          </div>
        </div>
      </div>

      <div class="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 opacity-40">
        <span class="font-body text-caption text-black/50 uppercase tracking-[0.3em]">Scroll</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-black/50 animate-bounce"><path d="M12 5v14"/><path d="m19 12-7 7-7-7"/></svg>
      </div>
    </section>

    <!-- ═══ BENEFICIOS ═══ -->
    <section class="py-24 bg-white border-b border-borde">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        <p class="font-body text-caption text-texto-muted uppercase tracking-[0.2em] mb-4 text-center">¿Por qué DEPRODECA?</p>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mt-12">
          <div v-for="b in beneficios" :key="b.titulo" class="p-8 rounded-2xl bg-fondo shadow-xs hover-lift text-center group">
            <div class="w-14 h-14 rounded-xl bg-[#D4A017]/10 flex items-center justify-center mx-auto mb-5">
              <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="#D4A017" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path :d="b.icono" /></svg>
            </div>
            <h3 class="font-body text-body font-bold text-texto mb-2">{{ b.titulo }}</h3>
            <p class="font-body text-small text-texto-muted leading-relaxed">{{ b.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- ═══ CATEGORÍAS ═══ -->
    <section class="py-28 bg-fondo">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        <div class="flex flex-col md:flex-row md:items-end justify-between mb-14 gap-4">
          <div>
            <p class="font-body text-caption text-texto-muted uppercase tracking-[0.2em] mb-3">Explorá</p>
            <h2 class="font-display text-display-lg text-texto uppercase">Categorías</h2>
          </div>
          <NuxtLink to="/catalogo" class="font-body text-small text-texto-muted hover:text-[#D4A017] transition-colors flex items-center gap-1">Ver todas <span>&rarr;</span></NuxtLink>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
          <NuxtLink v-for="cat in categoriasDestacadas" :key="cat.slug" :to="`/catalogo?categoria=${cat.slug}`" class="group flex flex-col items-center justify-center gap-4 p-8 rounded-2xl bg-white shadow-xs hover-lift text-center min-h-[140px]">
            <span class="font-body text-body font-semibold text-texto group-hover:text-[#D4A017] transition-colors duration-300">{{ cat.nombre }}</span>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- ═══ GAMIFICACIÓN ═══ -->
    <section class="py-28 bg-white">
      <div class="max-w-[1280px] mx-auto px-6 md:px-8">
        <div class="mb-14">
          <p class="font-body text-caption text-texto-muted uppercase tracking-[0.2em] mb-3">Programa de Lealtad</p>
          <h2 class="font-display text-display-lg text-texto uppercase mb-4">Crece con Nosotros</h2>
          <p class="font-body text-body text-texto-muted max-w-lg">Cada compra suma puntos. Subí de rango y desbloqueá beneficios exclusivos.</p>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div v-for="rango in rangos" :key="rango.nombre" class="flex flex-col items-center p-8 rounded-2xl bg-fondo shadow-xs hover-lift text-center">
            <span class="text-5xl mb-4" aria-hidden="true">{{ rango.icono }}</span>
            <h3 class="font-display text-heading text-texto uppercase mb-1">{{ rango.nombre }}</h3>
            <p class="font-body text-caption text-texto-muted mb-2">Desde {{ rango.minimo }}</p>
            <p class="font-body text-caption text-[#D4A017] font-semibold">{{ rango.beneficio }}</p>
          </div>
        </div>
        <div class="mt-10 text-center">
          <button class="inline-flex items-center gap-2 px-6 py-3 rounded-xl border border-borde font-body text-small font-semibold text-texto hover:border-[#D4A017] hover:text-[#D4A017] transition-all duration-300" @click="navigateTo('/perfil/rangos')">
            Ver todos los beneficios &rarr;
          </button>
        </div>
      </div>
    </section>

    <!-- ═══ CTA FINAL ═══ -->
    <section class="py-36 bg-texto relative overflow-hidden">
      <div class="absolute inset-0 opacity-[0.03] bg-[radial-gradient(circle_at_50%_50%,#fff_1px,transparent_1px)] bg-[length:20px_20px]" />
      <div class="relative max-w-[1280px] mx-auto px-6 md:px-8 text-center">
        <h2 class="font-display text-display-lg text-white uppercase mb-4">Empezá a comprar hoy</h2>
        <p class="font-body text-body text-white/50 max-w-md mx-auto mb-12">Registrate gratis. Sin costo de inscripción. Accedé a precios exclusivos.</p>
        <button class="inline-flex items-center gap-2 px-8 py-4 rounded-2xl bg-[#D4A017] text-black font-body text-body font-bold hover:bg-[#C49014] transition-all duration-500 hover:-translate-y-[2px] hover:shadow-xl active:scale-[0.98] min-h-[54px]" @click="navigateTo('/auth/registro')">
          Crear Cuenta Gratis
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
        </button>
      </div>
    </section>

  </div>
</template>

<style scoped>
.hero-text { animation: heroIn 900ms ease both; }
@keyframes heroIn { from { opacity: 0; transform: translateY(30px); } to { opacity: 1; transform: translateY(0); } }
</style>
