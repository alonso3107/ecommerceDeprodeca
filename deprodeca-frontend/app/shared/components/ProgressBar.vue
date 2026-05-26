<script setup lang="ts">

const props = withDefaults(
  defineProps<{
    valor: number
    maximo: number
    rango: "bronce" | "plata" | "oro" | "platino"
    mostrarEtiqueta?: boolean
  }>(),
  {
    mostrarEtiqueta: true,
  },
)

const porcentaje = computed(() => Math.min((props.valor / props.maximo) * 100, 100))

const colores = {
  bronce: "bg-amber-600",
  plata: "bg-gray-400",
  oro: "bg-yellow-500",
  platino: "bg-gradient-to-r from-cyan-400 to-purple-500",
}

const etiquetas = {
  bronce: "S/ 0 — S/ 4,999",
  plata: "S/ 5,000 — S/ 19,999",
  oro: "S/ 20,000 — S/ 49,999",
  platino: "S/ 50,000+",
}
</script>

<template>
  <div class="w-full">
    <!-- Barra de progreso -->
    <div
      class="w-full h-3 bg-surface rounded-full overflow-hidden"
      role="progressbar"
      :aria-valuenow="valor"
      :aria-valuemin="0"
      :aria-valuemax="maximo"
    >
      <div
        :class="[colores[rango], 'h-full rounded-full transition-all duration-700 ease-out']"
        :style="{ width: `${porcentaje}%` }"
      />
    </div>

    <!-- Etiqueta de progreso -->
    <div v-if="mostrarEtiqueta" class="flex justify-between mt-2 text-caption text-text-muted font-body">
      <span>S/ {{ valor.toLocaleString("es-PE") }}</span>
      <span>{{ etiquetas[rango] }}</span>
    </div>
  </div>
</template>
