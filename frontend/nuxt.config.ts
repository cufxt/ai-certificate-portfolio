// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',

  devtools: { enabled: true },

  modules: ['@nuxt/eslint', '@pinia/nuxt'],

  typescript: {
    strict: true,
    typeCheck: false,
  },

  vite: {
    plugins: [tailwindcss()],
  },

  css: ['~/assets/css/main.css'],

  components: [
    {
      path: '~/components',
      ignore: ['**/ui/**'],
    },
    {
      path: '~/components/ui',
      extensions: ['vue'],
    },
  ],
})
