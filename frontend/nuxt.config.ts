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
  modules: ["@nuxt/ui", "@nuxt/icon", "@nuxt/eslint", "@pinia/nuxt", "@nuxt/test-utils/module"],

  icon: {
    customCollections: [
      {
        prefix: "my-icons",
        dir: "./assets/my-icons",
      },
    ],
  },

  app: {
    head: {
      titleTemplate: "Perimeter",
      link: [{ rel: "icon", type: "image/x-icon", href: "/PerimeterIcon.png" }],
    },
  },
});
