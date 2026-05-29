<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    src?: string
    nombre: string
    size?: number
  }>(),
  {
    src: "",
    size: 96,
  },
)

const pathGestalt = computed(() => {
  const s = props.size
  const s075 = s * 0.75
  const s085 = s * 0.85
  const s07 = s * 0.7
  const s05 = s * 0.5
  const s03 = s * 0.3
  const s015 = s * 0.15

  return `M 0,0 L ${s},0 L ${s},${s075} C ${s085},${s085} ${s07},${s} ${s05},${s} C ${s03},${s} ${s015},${s085} 0,${s075} Z`
})

const iniciales = computed(() => {
  const nombre = (props.nombre || "").trim()
  if (!nombre) return "--"
  const partes = nombre.split(/\s+/).slice(0, 2)
  return partes.map((p) => p[0]?.toUpperCase() || "").join("")
})
</script>

<template>
  <svg :width="size" :height="size" :viewBox="`0 0 ${size} ${size}`" role="img" :aria-label="`Avatar de ${nombre || 'usuario'}`">
    <defs>
      <clipPath id="avatarClip">
        <path :d="pathGestalt" />
      </clipPath>
    </defs>

    <rect :width="size" :height="size" fill="#D4A017" fill-opacity="0.3" clip-path="url(#avatarClip)" />

    <image
      v-if="src"
      :href="src"
      :width="size"
      :height="size"
      preserveAspectRatio="xMidYMid slice"
      clip-path="url(#avatarClip)"
    />

    <text
      v-else
      x="50%"
      y="50%"
      text-anchor="middle"
      dominant-baseline="middle"
      class="font-display"
      :font-size="Math.round(size * 0.3)"
      fill="#171717"
    >
      {{ iniciales }}
    </text>

    <path :d="pathGestalt" stroke="#D4A017" stroke-width="2" fill="none" />
  </svg>
</template>
