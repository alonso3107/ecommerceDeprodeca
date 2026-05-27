<!--
  ═══════════════════════════════════════════════════════════════
  HomeCategorias.vue — Grid de categorías de producto
  Swiss Modernism 2.0 × Brutalista: celdas geométricas, hover
  invertido, tipografía estructural.
  ═══════════════════════════════════════════════════════════════
-->
<script setup lang="ts">
const categorias = [
  { 
    nombre: "Bebidas", 
    slug: "bebidas", 
    count: "120+",
    desc: "Gaseosas, aguas, jugos y energizantes",
    svgPath: "M12 4 L28 20 L44 4 M12 12 L28 28 L44 12",
  },
  { 
    nombre: "Snacks", 
    slug: "snacks", 
    count: "85+",
    desc: "Galletas, chocolates, chips y frutos secos",
    svgPath: "M8 24 H48 M28 4 V44 M16 12 L40 36 M40 12 L16 36",
  },
  { 
    nombre: "Lácteos", 
    slug: "lacteos", 
    count: "60+",
    desc: "Leches, yogures, quesos y mantequillas",
    svgPath: "M28 8 A20 20 0 1 1 28 48 A20 20 0 1 1 28 8 M28 18 A10 10 0 1 0 28 38 A10 10 0 1 0 28 18 M28 25 A3 3 0 1 0 28 31 A3 3 0 1 0 28 25",
  },
  { 
    nombre: "Abarrotes", 
    slug: "abarrotes", 
    count: "150+",
    desc: "Arroz, azúcar, aceite, menestras y conservas",
    svgPath: "M8 16 H48 M8 28 H48 M8 40 H48 M20 8 V48 M36 8 V48",
  },
  { 
    nombre: "Limpieza", 
    slug: "limpieza", 
    count: "45+",
    desc: "Detergentes, desinfectantes y artículos de aseo",
    svgPath: "M16 8 L40 8 L48 24 L40 48 L16 48 L8 24 Z M28 8 V48 M22 24 H34",
  },
  { 
    nombre: "Cuidado Personal", 
    slug: "cuidado-personal", 
    count: "40+",
    desc: "Shampoo, jabón, desodorante y cuidado bucal",
    svgPath: "M20 6 A8 14 0 1 0 36 6 M28 20 A12 12 0 1 0 28 44 A12 12 0 1 0 28 20 M20 44 H36",
  },
] as const
</script>

<template>
  <section class="relative py-28 md:py-36 bg-[#FFFFFF] border-t border-[#E5E5E5] overflow-hidden">
    
    <!-- SVG geométrico de fondo — líneas diagonales Swiss -->
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute -right-20 top-0 h-full w-[50%] opacity-[0.04]"
      viewBox="0 0 600 800"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M100 0L600 500V800L0 300V0Z" stroke="#D4A017" stroke-width="1" />
      <path d="M300 0L600 300V800L200 800L0 600" stroke="#D4A017" stroke-width="1" />
      <path d="M150 0L600 450" stroke="#D4A017" stroke-width="1" />
      <path d="M0 150L450 600" stroke="#D4A017" stroke-width="1" />
      <path d="M0 400L200 600" stroke="#D4A017" stroke-width="1" />
    </svg>

    <div class="relative max-w-[1280px] mx-auto px-6 md:px-8">
      
      <!-- Header Swiss -->
      <div class="flex flex-col md:flex-row md:items-end justify-between mb-16 md:mb-20 gap-6">
        <div>
          <p class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.3em] mb-5">
            02 · Categorías
          </p>
          <h2 class="font-display text-[clamp(2.5rem,6vw,5rem)] text-[#171717] uppercase leading-[0.92]">
            Explorá<span class="text-[#D4A017]">.</span>
          </h2>
          <div class="mt-5 h-[2px] w-24 bg-[#D4A017]" />
        </div>
        <NuxtLink
          to="/catalogo"
          class="font-mono text-[11px] text-[#666666] hover:text-[#D4A017] uppercase tracking-[0.2em] transition-colors duration-300 flex items-center gap-2 group/link"
        >
          <span>Ver todas las categorías</span>
          <span class="text-[#D4A017] group-hover/link:translate-x-1 transition-transform duration-300">→</span>
        </NuxtLink>
      </div>

      <!-- Grid 6 columnas — gap-px con hover invertido -->
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 bg-[#E5E5E5] gap-px">
        
        <NuxtLink
          v-for="cat in categorias"
          :key="cat.slug"
          :to="`/catalogo?categoria=${cat.slug}`"
          class="group/cat relative bg-[#FAFAFA] flex flex-col justify-between p-6 md:p-7 min-h-[200px] transition-all duration-500 ease-out hover:bg-[#171717] overflow-hidden"
        >
          
          <!-- SVG geométrico decorativo -->
          <svg
            aria-hidden="true"
            class="absolute bottom-3 right-3 w-16 h-16 opacity-[0.12] group-hover/cat:opacity-[0.25] transition-opacity duration-500"
            viewBox="0 0 56 56"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path :d="cat.svgPath" stroke="currentColor" stroke-width="1.5" class="text-[#D4A017] group-hover/cat:text-[#D4A017]" />
          </svg>

          <!-- Contenido superior -->
          <div class="relative z-10">
            <span class="font-display text-[17px] font-semibold text-[#171717] uppercase tracking-wide leading-tight group-hover/cat:text-[#FAFAFA] transition-colors duration-500">
              {{ cat.nombre }}
            </span>
            <p class="mt-2 font-body text-[12px] text-[#999999] leading-relaxed group-hover/cat:text-[#D4A017]/70 transition-colors duration-500 line-clamp-2">
              {{ cat.desc }}
            </p>
          </div>

          <!-- Contador + línea inferior -->
          <div class="relative z-10">
            <span class="font-mono text-[11px] text-[#999999] tracking-[0.15em] group-hover/cat:text-[#FAFAFA]/60 transition-colors duration-500">
              {{ cat.count }} productos
            </span>
          </div>

          <!-- Barra inferior dorada que aparece en hover -->
          <div class="absolute bottom-0 left-0 right-0 h-[3px] bg-[#D4A017] transform scale-x-0 group-hover/cat:scale-x-100 transition-transform duration-500 ease-out origin-left" />

        </NuxtLink>
        
      </div>

    </div>
  </section>
</template>
