<script setup lang="ts">
const props = defineProps<{
  pedidoId: string | number
  estado: string
  productos: Array<{ id?: number | string; cantidad?: number }>
}>()

const emit = defineEmits<{
  cancelar: []
  repetir: [productos: Array<{ id?: number | string; cantidad?: number }>]
}>()

const confirmarCancelacion = ref(false)

function clickCancelar() {
  if (!confirmarCancelacion.value) {
    confirmarCancelacion.value = true
    return
  }
  emit("cancelar")
  confirmarCancelacion.value = false
}
</script>

<template>
  <section class="flex flex-wrap gap-0">
    <button
      v-if="estado === 'pendiente'"
      class="inline-flex items-center gap-2 border border-[#DC2626] px-4 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#DC2626] hover:bg-[#DC2626] hover:text-white transition-colors duration-200"
      @click="clickCancelar"
    >
      <svg width="14" height="14" viewBox="0 0 14 14" fill="none"><circle cx="7" cy="7" r="6" stroke="currentColor"/><path d="M4 4L10 10M10 4L4 10" stroke="currentColor"/></svg>
      {{ confirmarCancelacion ? "Confirmar cancelacion?" : "Cancelar Pedido" }}
    </button>

    <button
      class="inline-flex items-center gap-2 border border-[#D4A017] px-4 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#D4A017] hover:bg-[#D4A017] hover:text-black transition-colors duration-200"
      :class="estado === 'pendiente' ? 'border-l-0' : ''"
      @click="emit('repetir', productos)"
    >
      <svg width="14" height="14" viewBox="0 0 14 14" fill="none"><path d="M12 7A5 5 0 1 1 10.7 3.6" stroke="currentColor"/><path d="M12 2.8V6.2H8.6" stroke="currentColor"/></svg>
      Repetir Pedido
    </button>

    <NuxtLink
      to="/pedidos"
      class="inline-flex items-center gap-2 border border-borde px-4 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-texto-muted hover:text-texto hover:bg-fondo transition-colors duration-200 border-l-0"
    >
      <svg width="14" height="14" viewBox="0 0 14 14" fill="none"><path d="M12 7H2M2 7L6 3M2 7L6 11" stroke="currentColor"/></svg>
      Volver a Mis Pedidos
    </NuxtLink>
  </section>
</template>
