<script setup lang="ts">
import { useIntersectionObserver, usePreferredReducedMotion, useTransition } from "@vueuse/core"

const beneficios = [
  {
    titulo: "Precios B2B",
    desc: "Sin intermediarios. Precios directos de Nestle para tu bodega.",
    cifra: "-30%",
    etiqueta: "vs retail",
  },
  {
    titulo: "Pedidos al Mayor",
    desc: "Compra por caja o unidad con stock estable para reposicion continua.",
    cifra: "500+",
    etiqueta: "productos",
  },
  {
    titulo: "Entrega 24-48h",
    desc: "Despacho rapido en Ica con seguimiento desde salida de almacen.",
    cifra: "48h",
    etiqueta: "maximo",
  },
  {
    titulo: "Gamificacion",
    desc: "Sube de rango y desbloquea beneficios acumulables por compra.",
    cifra: "4",
    etiqueta: "rangos",
  },
] as const

const sectionRef = ref<HTMLElement | null>(null)
const visible = ref(false)
const reducedMotion = usePreferredReducedMotion()

const cifraTargets = [30, 500, 48, 4]
const cifraValores = cifraTargets.map(() => ref(0))
const cifraTransiciones = cifraValores.map((v) =>
  useTransition(v, { duration: 1200 }),
)

useIntersectionObserver(
  sectionRef,
  ([entry]) => {
    if (entry?.isIntersecting && !visible.value) {
      visible.value = true
      cifraValores.forEach((valor, index) => {
        setTimeout(() => {
          valor.value = cifraTargets[index] || 0
        }, index * 100)
      })
    }
  },
  { threshold: 0.15 },
)

const cifrasMostradas = computed(() => {
  if (reducedMotion.value === "reduce") return ["-30%", "500+", "48h", "4"]
  return [
    `-${Math.round(cifraTransiciones[0]?.value || 0)}%`,
    `${Math.round(cifraTransiciones[1]?.value || 0)}+`,
    `${Math.round(cifraTransiciones[2]?.value || 0)}h`,
    `${Math.round(cifraTransiciones[3]?.value || 0)}`,
  ]
})
</script>

<template>
  <section ref="sectionRef" class="py-24 md:py-32 bg-crema border-t border-stone">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="mb-14 md:mb-16">
        <p class="font-mono text-[11px] uppercase text-stone-oscuro tracking-[0.2em] mb-4">Valor para tu bodega</p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-negro leading-[0.92]">
          Por que DEPRODECA<span class="text-dorado">.</span>
        </h2>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-[2fr_1fr_1fr] gap-6">
        <article
          class="bg-blanco border border-stone p-8 md:p-10 lg:row-span-2 hover:border-dorado transition-colors duration-300 flex flex-col justify-center"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
          :style="{ transition: 'opacity 500ms ease-out, transform 500ms ease-out', transitionDelay: '0ms' }"
        >
          <p class="font-display text-[clamp(4rem,8vw,7rem)] font-black text-dorado leading-none text-center">
            {{ cifrasMostradas[0] }}
          </p>
          <p class="font-mono text-[11px] text-stone-oscuro uppercase tracking-[0.2em] text-center mt-2">
            {{ beneficios[0].etiqueta }}
          </p>
          <div class="w-16 border-t border-stone/40 mx-auto my-5" />
          <h3 class="font-display text-[22px] font-bold text-negro text-center">
            {{ beneficios[0].titulo }}
          </h3>
          <p class="font-body text-[15px] text-stone-oscuro leading-relaxed text-center mt-3 max-w-[320px] mx-auto">
            {{ beneficios[0].desc }}
          </p>
        </article>

        <article
          v-for="(b, index) in beneficios.slice(1)"
          :key="b.titulo"
          class="bg-blanco border border-stone p-5 md:p-6 hover:border-dorado transition-colors duration-300 flex items-center gap-4"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
          :style="{ transition: 'opacity 500ms ease-out, transform 500ms ease-out', transitionDelay: `${(index + 1) * 80}ms` }"
        >
          <div class="shrink-0 text-right min-w-[70px]">
            <p class="font-display text-[clamp(2rem,3.5vw,2.8rem)] font-bold text-dorado leading-none">
              {{ cifrasMostradas[index + 1] }}
            </p>
            <p class="font-mono text-[9px] text-stone-oscuro uppercase tracking-[0.15em] mt-1">
              {{ b.etiqueta }}
            </p>
          </div>

          <div class="w-px h-10 bg-stone/30 shrink-0" />

          <div class="min-w-0">
            <h3 class="font-body text-[14px] md:text-[15px] font-bold text-negro leading-tight">
              {{ b.titulo }}
            </h3>
            <p class="font-body text-[12px] text-stone-oscuro leading-snug mt-1 line-clamp-2">
              {{ b.desc }}
            </p>
          </div>
        </article>
      </div>
    </div>
  </section>
</template>
