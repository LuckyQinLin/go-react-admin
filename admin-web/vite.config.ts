import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@assets": path.resolve(__dirname, "src/assets"),
      "@pages": path.resolve(__dirname, "src/pages"),
      "@redux ": path.resolve(__dirname, "src/redux"),
      "@images ": path.resolve(__dirname, "src/assets/images"),
    },
  },
  css: {
    preprocessorOptions: {
      less: {
        // 全局变量
        modifyVars: {},
        javascriptEnabled: true,
        additionalData: `@import "src/assets/less/var.less";`,
      },
    },
  },
  server: {
    port: 8080,
    proxy: {
      "/api": {
        target: "http://localhost:8070",
        changeOrigin: true,
      },
    },
  },
})