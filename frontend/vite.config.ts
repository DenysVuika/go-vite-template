import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  base: './', // 👈 relative paths so Go can serve from root
  build: {
    outDir: 'dist',        // default, but just to be explicit
    emptyOutDir: true,     // cleans before build
  },
})
