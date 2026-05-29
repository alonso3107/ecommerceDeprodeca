<script setup lang="ts">
const props = defineProps<{
  categorias: Array<{ slug: string; nombre: string }>
  categoriaActiva: string
  buscando: boolean
  queryTexto: string
  resultados: number
}>()

const emit = defineEmits<{
  filtrar: [slug: string]
  buscar: [query: string]
}>()

const query = ref(props.queryTexto || "")

watch(
  () => props.queryTexto,
  (valor) => {
    query.value = valor || ""
  },
)

function enviarBusqueda() {
  emit("buscar", query.value.trim())
}
</script>

<template>
  <section class="mb-10">
    <form class="flex flex-col md:flex-row gap-3" @submit.prevent="enviarBusqueda">
      <input
        v-model="query"
        type="search"
        placeholder="Buscar producto premium..."
        class="flex-1 border border-stone bg-blanco px-5 py-4 font-body text-[16px] text-negro placeholder:text-stone-oscuro focus:outline-none focus:border-dorado"
      />
      <button type="submit" class="bg-negro text-blanco border border-negro font-mono text-[11px] uppercase tracking-[0.3em] px-8 py-4 hover:bg-dorado hover:text-negro transition-colors duration-300">
        Buscar
      </button>
    </form>

    <div class="mt-6 flex flex-col md:flex-row md:items-center md:justify-between gap-4">
      <div class="flex flex-wrap gap-3">
        <button
          class="border px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] transition-colors duration-300"
          :class="categoriaActiva ? 'border-stone text-stone-oscuro hover:border-negro hover:text-negro' : 'bg-negro text-blanco border-negro font-bold'"
          @click="emit('filtrar', '')"
        >
          Todo
        </button>
        <button
          v-for="cat in categorias"
          :key="cat.slug"
          class="border px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] transition-colors duration-300"
          :class="categoriaActiva === cat.slug ? 'bg-negro text-blanco border-negro font-bold' : 'border-stone text-stone-oscuro hover:border-negro hover:text-negro'"
          @click="emit('filtrar', cat.slug)"
        >
          {{ cat.nombre }}
        </button>
      </div>
      <p class="font-mono text-[11px] text-stone-oscuro uppercase tracking-[0.2em]">{{ resultados }} productos</p>
    </div>

    <div v-if="categoriaActiva || buscando" class="mt-4 flex items-center gap-3">
      <span class="border border-stone bg-crema px-4 py-2 font-mono text-[11px] uppercase tracking-[0.18em] text-stone-oscuro">
        Filtro activo: {{ categoriaActiva || queryTexto }}
      </span>
      <button class="border border-stone px-3 py-2 font-mono text-[10px] uppercase tracking-[0.2em] text-stone-oscuro hover:border-negro hover:text-negro transition-colors duration-300" @click="emit('filtrar', ''); emit('buscar', '')">
        Limpiar
      </button>
    </div>
  </section>
</template>
