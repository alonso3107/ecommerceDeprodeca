// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },

  modules: ["@pinia/nuxt", "@vueuse/nuxt"],

  components: [
    { path: "~/shared/components", pathPrefix: false },
    { path: "~/features/home/components", pathPrefix: false },
    { path: "~/features/catalogo/components", pathPrefix: false },
    { path: "~/features/contacto/components", pathPrefix: false },
    { path: "~/features/checkout/components", pathPrefix: false },
    { path: "~/features/perfil/components", pathPrefix: false },
    { path: "~/features/gamificacion/components", pathPrefix: false },
  ],

  css: ["~/assets/styles/main.css"],

  vite: {
    plugins: [tailwindcss()],
  },

  app: {
    head: {
      title: "DEPRODECA",
      htmlAttrs: { lang: "es" },
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          name: "description",
          content:
            "Tienda online B2B para proveedores de bodega de DEPRODECA, filial de Nestlé Perú.",
        },
      ],
      link: [
        { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
        { rel: "preconnect", href: "https://fonts.googleapis.com" },
        {
          rel: "preconnect",
          href: "https://fonts.gstatic.com",
          crossorigin: "",
        },
        {
          rel: "stylesheet",
          href: "https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;700&family=Nunito+Sans:wght@300;400;500;600;700&family=Rubik:wght@400;500;600;700;800;900&display=swap",
        },
      ],
    },
  },

  runtimeConfig: {
    public: {
      apiBase:
        process.env.NUXT_PUBLIC_API_BASE || "http://localhost:8080/api/v1",
      supabaseUrl:
        process.env.NUXT_PUBLIC_SUPABASE_URL || process.env.SUPABASE_URL || "",
      supabaseAnonKey:
        process.env.NUXT_PUBLIC_SUPABASE_ANON_KEY ||
        process.env.SUPABASE_ANON_KEY ||
        "",
    },
  },

  typescript: {
    strict: true,
  },
});
