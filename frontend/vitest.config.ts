import { defineConfig } from 'vitest/config';
import Vue from '@vitejs/plugin-vue';
import { fileURLToPath } from 'url';

export default defineConfig({
  plugins: [Vue()],
  test: {
    environment: 'happy-dom',
    setupFiles: './tests/setup.ts',
    alias: {
      '@': fileURLToPath(new URL('./', import.meta.url)),
    },
    globals: true,
  },
});
