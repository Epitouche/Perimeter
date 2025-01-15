// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  telemetry: { enabled: false },

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  modules: [
    "@nuxt/ui",
    "@nuxt/icon",
    "@nuxt/eslint",
    "@pinia/nuxt",
    "@nuxtjs/google-fonts",
    ...(process.env.NODE_ENV != "production" ? ["@nuxt/test-utils/module"] : []),
  ],

  css: ["~/assets/css/main.css"],

  app: {
    head: {
      titleTemplate: "Perimeter",
      link: [{ rel: "icon", type: "image/x-icon", href: "/PerimeterIcon.png" }],
    },
  },

  googleFonts: {
    families: {
      "Kantumruy Pro": true,
    },
  },
});
