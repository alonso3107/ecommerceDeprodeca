<script setup lang="ts">
import { resolverImagenProducto } from "~/shared/utils/imagenes-productos";

definePageMeta({ layout: "default" });

const route = useRoute();
const config = useRuntimeConfig();

const productos = ref<any[]>([]);
const categorias = ref<any[]>([]);
const loading = ref(true);
const pagina = ref(1);

const categoriaSlug = computed(() => (route.query.categoria as string) || "");
const queryBusqueda = computed(() => (route.query.q as string) || "");

const productosConImagen = computed(() => {
    const imagenesUsadas = new Set<string>();

    return productos.value.map((producto, indice) => ({
        ...producto,
        imagen_resuelta: resolverImagenProducto(
            producto,
            indice,
            imagenesUsadas,
        ),
    }));
});

onMounted(async () => {
    try {
        const catRes = await $fetch<{ success: boolean; data: any[] }>(
            `${config.public.apiBase}/categorias`,
        );
        if (catRes.success) categorias.value = catRes.data;
    } catch (_) {
        /* fallback silencioso */
    }

    await cargarProductos();
});

watch(
    () => route.query,
    () => {
        pagina.value = 1;
        cargarProductos();
    },
);

async function cargarProductos() {
    loading.value = true;

    try {
        const params = new URLSearchParams();
        params.set("pagina", String(pagina.value));
        params.set("limite", "12");
        if (categoriaSlug.value)
            params.set("categoria_id", categoriaSlug.value);
        if (queryBusqueda.value) params.set("q", queryBusqueda.value);

        const endpoint = queryBusqueda.value
            ? `${config.public.apiBase}/productos/buscar?${params}`
            : `${config.public.apiBase}/productos?${params}`;

        const res = await $fetch<{ success: boolean; data: any[] }>(endpoint);
        productos.value = res.success ? res.data : [];
    } catch (_) {
        productos.value = [];
    } finally {
        loading.value = false;
    }
}

function filtrar(slug: string) {
    if (slug) navigateTo(`/catalogo?categoria=${slug}`);
    else navigateTo("/catalogo");
}

function buscar(query: string) {
    if (query) navigateTo(`/catalogo?q=${encodeURIComponent(query)}`);
    else navigateTo("/catalogo");
}
</script>

<template>
    <div class="max-w-[1280px] mx-auto px-6 md:px-8 py-12 md:py-16">
        <header
            class="relative mb-10 md:mb-14 border border-[#e5e5e5] bg-[#FAFAFA] p-6 md:p-10"
        >
            <svg
                aria-hidden="true"
                class="pointer-events-none absolute inset-y-0 right-0 h-full w-[70%] max-w-[760px] opacity-10"
                viewBox="0 0 960 420"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path d="M40 80H920" stroke="#D4A017" stroke-width="2" />
                <path
                    d="M80 120L240 40L400 120L560 40L720 120L880 40"
                    stroke="#D4A017"
                    stroke-width="2"
                />
                <path
                    d="M120 320L260 200L380 320L520 160L680 320L840 200"
                    stroke="#D4A017"
                    stroke-width="2"
                />
                <path d="M150 360H810" stroke="#D4A017" stroke-width="2" />
                <path
                    d="M300 70L360 10L420 70"
                    stroke="#D4A017"
                    stroke-width="2"
                />
                <path
                    d="M600 70L660 10L720 70"
                    stroke="#D4A017"
                    stroke-width="2"
                />
                <path
                    d="M740 290L800 230L860 290"
                    stroke="#D4A017"
                    stroke-width="2"
                />
            </svg>

            <div class="relative z-10 max-w-[780px]">
                <p
                    class="font-mono text-[11px] uppercase tracking-[0.3em] text-[#666666]"
                >
                    Catálogo / Productos
                </p>
                <h1
                    class="mt-4 font-display text-[clamp(3rem,8vw,6.5rem)] uppercase leading-[0.9] text-[#171717]"
                >
                    Productos
                </h1>
                <p
                    class="mt-4 max-w-[60ch] font-body text-[16px] leading-relaxed text-[#666666]"
                >
                    {{
                        queryBusqueda
                            ? `Resultados para "${queryBusqueda}"`
                            : "Productos de Deprodeca."
                    }}
                </p>
                <div
                    v-if="queryBusqueda || categoriaSlug"
                    class="mt-6 flex flex-wrap gap-2"
                >
                    <span
                        v-if="queryBusqueda"
                        class="border border-[#171717] bg-[#171717] px-3 py-1 font-mono text-[10px] uppercase tracking-[0.2em] text-[#FAFAFA]"
                    >
                        Búsqueda: {{ queryBusqueda }}
                    </span>
                    <span
                        v-if="categoriaSlug"
                        class="border border-[#e5e5e5] px-3 py-1 font-mono text-[10px] uppercase tracking-[0.2em] text-[#171717]"
                    >
                        Categoría: {{ categoriaSlug }}
                    </span>
                </div>
            </div>
        </header>

        <FiltroCategoria
            :categorias="categorias"
            :categoria-activa="categoriaSlug"
            :buscando="!!queryBusqueda"
            :query-texto="queryBusqueda"
            @filtrar="filtrar"
            @buscar="buscar"
        />

        <div class="mt-10">
            <div
                v-if="loading"
                class="flex flex-col items-center justify-center py-32 gap-4"
            >
                <div
                    class="w-8 h-8 border-2 border-[#171717] border-t-transparent animate-spin"
                />
                <span
                    class="font-mono text-[11px] text-[#666666] uppercase tracking-[0.2em]"
                    >Cargando productos</span
                >
            </div>

            <div
                v-else-if="productosConImagen.length === 0"
                class="border border-[#e5e5e5] bg-[#FAFAFA] px-6 py-24 text-center"
            >
                <p
                    class="font-display text-[clamp(2.5rem,7vw,5rem)] uppercase leading-none text-[#171717]"
                >
                    Sin resultados
                </p>
                <p
                    class="mx-auto mt-5 max-w-[52ch] font-body text-[16px] leading-relaxed text-[#666666]"
                >
                    No encontramos productos con esos filtros.
                </p>
                <button
                    class="mt-8 border border-[#171717] bg-[#171717] px-6 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-[#FAFAFA] transition-colors duration-200 hover:bg-[#FAFAFA] hover:text-[#171717]"
                    @click="navigateTo('/catalogo')"
                >
                    Ver todo
                </button>
            </div>

            <div
                v-else
                class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-px bg-[#e5e5e5] p-px"
            >
                <ProductCard
                    v-for="p in productosConImagen"
                    :key="p.id"
                    :slug="p.slug"
                    :nombre="p.nombre"
                    :precio="p.precio"
                    :unidad="p.unidad"
                    :imagen-url="p.imagen_resuelta"
                    :categoria="p.categoria_nombre || 'General'"
                    :stock="p.stock"
                />
            </div>

            <div
                v-if="productos.length >= 12"
                class="mt-14 flex items-center justify-center"
            >
                <div class="flex items-stretch">
                    <button
                        class="border border-[#e5e5e5] px-5 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#666666] transition-colors duration-200 hover:border-[#171717] hover:text-[#171717] disabled:cursor-not-allowed disabled:opacity-30"
                        :disabled="pagina <= 1"
                        @click="
                            pagina--;
                            cargarProductos();
                        "
                    >
                        Anterior
                    </button>

                    <span
                        class="border-y border-[#e5e5e5] px-8 py-3 font-mono text-[11px] uppercase tracking-[0.2em] text-[#171717]"
                    >
                        Página {{ pagina }}
                    </span>

                    <button
                        class="border border-[#e5e5e5] px-5 py-3 font-mono text-[11px] uppercase tracking-[0.15em] text-[#666666] transition-colors duration-200 hover:border-[#171717] hover:text-[#171717]"
                        @click="
                            pagina++;
                            cargarProductos();
                        "
                    >
                        Siguiente
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
