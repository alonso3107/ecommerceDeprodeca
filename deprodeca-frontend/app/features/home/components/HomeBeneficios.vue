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
  <section ref="sectionRef" class="py-24 md:py-32 bg-[#FAFAF9] border-t border-[#C5BFB5]">
    <div class="max-w-[1280px] mx-auto px-6 md:px-8">
      <div class="mb-14 md:mb-16">
        <p class="font-mono text-[11px] uppercase text-[#5C554D] tracking-[0.2em] mb-4">Valor para tu bodega</p>
        <h2 class="font-display text-[clamp(2.5rem,6vw,4rem)] font-bold text-[#1C1917] leading-[0.92]">Por que DEPRODECA</h2>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <article
          v-for="(b, index) in beneficios"
          :key="b.titulo"
          class="bg-white border border-[#C5BFB5] p-8 hover:border-[#A16207] transition-colors duration-300"
          :class="(visible || reducedMotion === 'reduce') ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6'"
          :style="{ transition: 'opacity 500ms ease-out, transform 500ms ease-out', transitionDelay: `${index * 100}ms` }"
        >
          <p class="font-display text-[clamp(2.5rem,5vw,3.5rem)] font-bold text-[#A16207] leading-none">{{ cifrasMostradas[index] }}</p>
          <p class="font-mono text-[10px] text-[#5C554D] uppercase tracking-[0.15em] mt-2">{{ b.etiqueta }}</p>
          <h3 class="font-body text-[16px] font-bold text-[#1C1917] mt-4">{{ b.titulo }}</h3>
          <p class="font-body text-[14px] text-[#5C554D] leading-relaxed mt-2">{{ b.desc }}</p>
        </article>
      </div>
    </div>
  </section>
</template>
