     1|<script setup lang="ts">
     2|// ─── DEPRODECA Landing — Hero Hungry Tiger Style ──────
     3|// Feature-Sliced Design: página, sin lógica de negocio directa
     4|
     5|definePageMeta({
     6|  layout: "default",
     7|})
     8|
     9|// Categorías destacadas (datos mock — se conectarán a la API)
    10|const categoriasDestacadas = [
    11|  { nombre: "Bebidas", slug: "bebidas", color: "bg-brand-primary/20", icono: "M9 20H4V9h5v11zm11 0h-5v-5h5v5zm0-6h-5V4h5v10z" },
    12|  { nombre: "Snacks", slug: "snacks", color: "bg-status-warning/20", icono: "M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5" },
    13|  { nombre: "Lácteos", slug: "lacteos", color: "bg-brand-secondary/20", icono: "M20 9V7a2 2 0 00-2-2h-4a2 2 0 00-2 2v2M4 9V7a2 2 0 012-2h4a2 2 0 012 2v2" },
    14|  { nombre: "Abarrotes", slug: "abarrotes", color: "bg-status-success/20", icono: "M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2z" },
    15|  { nombre: "Limpieza", slug: "limpieza", color: "bg-surface", icono: "M12 2L2 7l10 5 10-5-10-5z" },
    16|  { nombre: "Cuidado Personal", slug: "cuidado-personal", color: "bg-status-error/10", icono: "M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z" },
    17|]
    18|
    19|// Rangos de gamificación
    20|const rangos = [
    21|  { nombre: "Bronce", color: "bg-amber-600", minimo: "S/ 0", icono: "🥉" },
    22|  { nombre: "Plata", color: "bg-gray-400", minimo: "S/ 5,000", icono: "🥈" },
    23|  { nombre: "Oro", color: "bg-yellow-500", minimo: "S/ 20,000", icono: "🥇" },
    24|  { nombre: "Platino", color: "bg-purple-500", minimo: "S/ 50,000", icono: "💎" },
    25|]
    26|
    27|// Beneficios
    28|const beneficios = [
    29|  { titulo: "Precios B2B", descripcion: "Precios exclusivos para dueños de bodega. Sin intermediarios.", icono: "M12 2v20M17 5H9.5a3.5 3.5 0 000 7h5a3.5 3.5 0 010 7H6" },
    30|  { titulo: "Pedidos al por Mayor", descripcion: "Compra por caja o unidad. Stock siempre disponible.", icono: "M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" },
    31|  { titulo: "Gamificación", descripcion: "Sube de rango con cada compra. Bronce → Platino.", icono: "M12 15l-4.24-4.24a6 6 0 0112.48 0L20 15M6 12a6 6 0 0112 0" },
    32|  { titulo: "Entrega Rápida", descripcion: "Despacho en 24-48h a todo Ica. Seguimiento en tiempo real.", icono: "M13 10V3L4 14h7v7l9-11h-7z" },
    33|]
    34|</script>
    35|
    36|<template>
    37|  <div class="page-enter">
    38|
    39|    <!-- ═══════════════════════════════════════════════════
    40|        HERO HUNGRY TIGER — Asimétrico, tipografía masiva,
    41|        bloques amarillos, producto desbordando
    42|       ═══════════════════════════════════════════════════ -->
    43|    <section class="relative overflow-hidden bg-white">
    44|      <div class="max-w-[1280px] mx-auto px-4 md:px-6 pt-12 pb-20 md:pt-20 md:pb-32">
    45|
    46|        <!-- ─── Composición asimétrica (Z-pattern) ────── -->
    47|        <div class="grid grid-cols-1 lg:grid-cols-[1fr_1.2fr] gap-8 lg:gap-16 items-center">
    48|
    49|          <!-- ── Bloque izquierdo: Texto + CTA ────────── -->
    50|          <div class="relative z-10">
    51|            <!-- Badge pequeño -->
    52|            <span
    53|              class="inline-block px-3 py-1.5 rounded-full bg-brand-primary/20 text-text-primary font-body text-small font-semibold mb-6"
    54|            >
    55|              Filial Oficial de Nestlé Perú
    56|            </span>
    57|
    58|            <!-- Título masivo — Big Noodle Titling -->
    59|            <h1 class="font-display text-display-xl text-text-primary leading-[1.05] tracking-[-0.02em] uppercase mb-6">
    60|              <span class="block">Proveedor</span>
    61|              <span class="block relative">
    62|                de
    63|                <span class="relative inline-block">
    64|                  Bodega
    65|                  <!-- Subrayado amarillo detrás del texto -->
    66|                  <span class="absolute -bottom-2 left-0 right-0 h-[14px] bg-brand-primary -z-10 rounded-sm" aria-hidden="true" />
    67|                </span>
    68|              </span>
    69|            </h1>
    70|
    71|            <!-- Descripción -->
    72|            <p class="font-body text-body text-text-muted max-w-md mb-8 leading-relaxed">
    73|              La plataforma B2B que conecta tu bodega directamente con los productos Nestlé.
    74|              Precios exclusivos, pedidos al por mayor y un sistema de rangos que premia tu lealtad.
    75|            </p>
    76|
    77|            <!-- CTAs -->
    78|            <div class="flex flex-col sm:flex-row gap-3">
    79|              <Button variant="primary" size="lg" @click="navigateTo('/catalogo')">
    80|                Ver Catálogo
    81|              </Button>
    82|              <Button variant="outline" size="lg" @click="navigateTo('/auth/registro')">
    83|                Crear Cuenta Gratis
    84|              </Button>
    85|            </div>
    86|
    87|            <!-- Métricas rápidas -->
    88|            <div class="flex gap-8 mt-10 pt-8 border-t border-border-base">
    89|              <div>
    90|                <p class="font-display text-display-md text-brand-primary leading-none">500+</p>
    91|                <p class="font-body text-caption text-text-muted mt-1">Productos</p>
    92|              </div>
    93|              <div>
    94|                <p class="font-display text-display-md text-brand-primary leading-none">200+</p>
    95|                <p class="font-body text-caption text-text-muted mt-1">Bodegas</p>
    96|              </div>
    97|              <div>
    98|                <p class="font-display text-display-md text-brand-primary leading-none">4</p>
    99|                <p class="font-body text-caption text-text-muted mt-1">Rangos</p>
   100|              </div>
   101|            </div>
   102|          </div>
   103|
   104|          <!-- ── Bloque derecho: Imagen desbordando (asimetría) ── -->
   105|          <div class="relative lg:-mr-16">
   106|            <!-- Fondo amarillo enorme detrás de la imagen -->
   107|            <div
   108|              class="absolute -inset-8 bg-brand-primary rounded-3xl -rotate-3 -z-0"
   109|              aria-hidden="true"
   110|            />
   111|
   112|            <!-- Imagen principal del producto/bodega -->
   113|            <div class="relative z-10 rounded-2xl overflow-hidden shadow-lg">
   114|              <img
   115|                src="https://images.unsplash.com/photo-1543168256-418811576931?w=800&q=80"
   116|                alt="Productos Nestlé en exhibición para bodegas peruanas"
   117|                class="w-full h-[480px] object-cover"
   118|              />
   119|            </div>
   120|
   121|            <!-- Producto "desbordando" flotante -->
   122|            <div
   123|              class="absolute -bottom-8 -left-8 z-20 bg-white rounded-2xl shadow-lg p-4 flex items-center gap-3 border-2 border-border-base hover:shadow-lg transition-shadow duration-500"
   124|            >
   125|              <div class="w-14 h-14 rounded-xl bg-brand-primary/20 flex items-center justify-center flex-shrink-0">
   126|                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="text-text-primary"><path d="M6 2 3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4Z"/><path d="M3 6h18"/><path d="M16 10a4 4 0 01-8 0"/></svg>
   127|              </div>
   128|              <div>
   129|                <p class="font-body text-caption text-text-muted">Precio Bodega</p>
   130|                <p class="font-display text-heading text-text-primary leading-none">S/ 1.50</p>
   131|              </div>
   132|            </div>
   133|          </div>
   134|        </div>
   135|      </div>
   136|    </section>
   137|
   138|    <!-- ═══════════════════════════════════════════════════
   139|        CATEGORÍAS DESTACADAS
   140|       ═══════════════════════════════════════════════════ -->
   141|    <section class="py-16 md:py-24 bg-surface">
   142|      <div class="max-w-[1280px] mx-auto px-4 md:px-6">
   143|        <div class="text-center mb-12">
   144|          <h2 class="font-display text-display-lg text-text-primary uppercase mb-3">
   145|            Categorías
   146|          </h2>
   147|          <p class="font-body text-body text-text-muted max-w-lg mx-auto">
   148|            Todo lo que tu bodega necesita, organizado por categorías.
   149|          </p>
   150|        </div>
   151|
   152|        <!-- Grid de categorías -->
   153|        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
   154|          <NuxtLink
   155|            v-for="cat in categoriasDestacadas"
   156|            :key="cat.slug"
   157|            :to="`/catalogo?categoria=${cat.slug}`"
   158|            :class="[
   159|              'group flex flex-col items-center gap-3 p-6 rounded-xl border-2 border-border-base hover:border-brand-primary transition-all duration-500 hover:shadow-brand',
   160|              cat.color,
   161|            ]"
   162|          >
   163|            <svg
   164|              xmlns="http://www.w3.org/2000/svg"
   165|              width="28" height="28"
   166|              viewBox="0 0 24 24"
   167|              fill="none"
   168|              stroke="currentColor"
   169|              stroke-width="1.5"
   170|              stroke-linecap="round"
   171|              stroke-linejoin="round"
   172|              class="text-text-primary group-hover:text-brand-secondary transition-colors duration-300"
   173|              aria-hidden="true"
   174|            >
   175|              <path :d="cat.icono" />
   176|            </svg>
   177|            <span class="font-body text-small font-semibold text-text-primary text-center">
   178|              {{ cat.nombre }}
   179|            </span>
   180|          </NuxtLink>
   181|        </div>
   182|      </div>
   183|    </section>
   184|
   185|    <!-- ═══════════════════════════════════════════════════
   186|        GAMIFICACIÓN — Rangos
   187|       ═══════════════════════════════════════════════════ -->
   188|    <section class="py-16 md:py-24 bg-white">
   189|      <div class="max-w-[1280px] mx-auto px-4 md:px-6">
   190|        <div class="text-center mb-12">
   191|          <h2 class="font-display text-display-lg text-text-primary uppercase mb-3">
   192|            Crece con Nosotros
   193|          </h2>
   194|          <p class="font-body text-body text-text-muted max-w-lg mx-auto">
   195|            Cada compra te acerca al siguiente rango. Más compras, más beneficios.
   196|          </p>
   197|        </div>
   198|
   199|        <!-- Rangos -->
   200|        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
   201|          <div
   202|            v-for="(rango, i) in rangos"
   203|            :key="rango.nombre"
   204|            class="relative flex flex-col items-center p-6 rounded-xl border-2 border-border-base bg-white hover:border-brand-primary hover:shadow-brand transition-all duration-500 group"
   205|          >
   206|            <span class="text-3xl mb-3" aria-hidden="true">{{ rango.icono }}</span>
   207|            <h3 class="font-display text-heading text-text-primary uppercase mb-1">{{ rango.nombre }}</h3>
   208|            <p class="font-body text-caption text-text-muted">Desde {{ rango.minimo }}</p>
   209|            <span
   210|              v-if="i < rangos.length - 1"
   211|              class="hidden md:block absolute -right-3 top-1/2 -translate-y-1/2 text-text-muted text-xl z-10"
   212|              aria-hidden="true"
   213|            >
   214|              ›
   215|            </span>
   216|          </div>
   217|        </div>
   218|
   219|        <div class="text-center mt-8">
   220|          <Button variant="secondary" size="md" @click="navigateTo('/perfil/rangos')">
   221|            Ver Beneficios por Rango
   222|          </Button>
   223|        </div>
   224|      </div>
   225|    </section>
   226|
   227|    <!-- ═══════════════════════════════════════════════════
   228|        BENEFICIOS
   229|       ═══════════════════════════════════════════════════ -->
   230|    <section class="py-16 md:py-24 bg-text-primary">
   231|      <div class="max-w-[1280px] mx-auto px-4 md:px-6">
   232|        <div class="text-center mb-12">
   233|          <h2 class="font-display text-display-lg text-white uppercase mb-3">
   234|            ¿Por qué DEPRODECA?
   235|          </h2>
   236|        </div>
   237|
   238|        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
   239|          <div
   240|            v-for="beneficio in beneficios"
   241|            :key="beneficio.titulo"
   242|            class="p-6 rounded-xl border border-gray-800 hover:border-brand-primary transition-all duration-500 group"
   243|          >
   244|            <div class="w-12 h-12 rounded-xl bg-brand-primary/10 flex items-center justify-center mb-4 group-hover:bg-brand-primary/20 transition-colors duration-300">
   245|              <svg
   246|                xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24"
   247|                fill="none" stroke="currentColor" stroke-width="1.5"
   248|                stroke-linecap="round" stroke-linejoin="round"
   249|                class="text-brand-primary"
   250|                aria-hidden="true"
   251|              >
   252|                <path :d="beneficio.icono" />
   253|              </svg>
   254|            </div>
   255|            <h3 class="font-body text-subheading font-bold text-white mb-2">{{ beneficio.titulo }}</h3>
   256|            <p class="font-body text-small text-gray-400 leading-relaxed">{{ beneficio.descripcion }}</p>
   257|          </div>
   258|        </div>
   259|      </div>
   260|    </section>
   261|
   262|    <!-- ═══════════════════════════════════════════════════
   263|        CTA FINAL
   264|       ═══════════════════════════════════════════════════ -->
   265|    <section class="py-20 bg-brand-primary">
   266|      <div class="max-w-[1280px] mx-auto px-4 md:px-6 text-center">
   267|        <h2 class="font-display text-display-lg text-text-primary uppercase mb-4">
   268|          Empieza a Comprar Hoy
   269|        </h2>
   270|        <p class="font-body text-body text-text-primary/80 max-w-md mx-auto mb-8">
   271|          Regístrate gratis y accede a precios exclusivos para tu bodega.
   272|        </p>
   273|        <div class="flex flex-col sm:flex-row justify-center gap-3">
   274|          <Button variant="secondary" size="lg" @click="navigateTo('/auth/registro')">
   275|            Crear Cuenta Gratis
   276|          </Button>
   277|          <Button variant="outline" size="lg" @click="navigateTo('/catalogo')">
   278|            Explorar Productos
   279|          </Button>
   280|        </div>
   281|      </div>
   282|    </section>
   283|
   284|  </div>
   285|</template>
   286|