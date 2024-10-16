import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';

import { resolve } from 'node:path';
import inputPlugin from '@macropygia/vite-plugin-glob-input';
import suidPlugin from '@suid/vite-plugin';

// @ts-expect-error
import handlebars from 'vite-plugin-handlebars';

export default defineConfig({
  build: {
    outDir: '.assets',
  },
  plugins: [
    solid(),
    suidPlugin(),
    inputPlugin({
      patterns: ['./pages/*.html'],
    }),
    handlebars({
      partialDirectory: resolve(__dirname, './src/partials'),
    }),
  ],
  css: {
    modules: {
      localsConvention: 'dashes',
    },
    preprocessorOptions: {
      scss: {
        api: 'modern-compiler',
      },
    },
  },
  resolve: {
    alias: {
      '@src': resolve(__dirname, './src'),
      '@partial': resolve(__dirname, './src/partials'),
    },
  },
});
