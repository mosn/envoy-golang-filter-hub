import { defineConfig } from 'unocss'
import presetUno from 'unocss/preset-uno'
import transformerVariantGroup from '@unocss/transformer-variant-group'
import transformerDirectives from '@unocss/transformer-directives'
import presetIcons from '@unocss/preset-icons'

export default defineConfig({
  presets: [
    presetUno({
      dark: 'class'
    }),
    presetIcons({
      collections: {
        mingcute: () => import('@iconify-json/mingcute/icons.json').then((i) => i.default)
      }
    })
  ],
  transformers: [transformerVariantGroup(), transformerDirectives()],
  theme: {
    colors: {
      primary: {
        light: '#e8effb',
        DEFAULT: '#0366d6',
        dark: '#071f40'
      }
    }
  },
  shortcuts: {
    'g-capsule':
      'px-2 lh-1.1em flex-shrink-0 max-w-1/3 text-primary/80 border-1.5 border-primary/30 rounded-full whitespace-nowrap',
    'g-btn':
      'transition bg-transparent border-primary/20 border-1.5 text-primary px-4 py-1 rounded-md transform-gpu flex items-center gap-1 hover:(bg-primary text-white) active:(scale-90)'
  }
})
