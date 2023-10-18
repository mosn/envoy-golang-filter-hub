import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useColorMode } from '@vueuse/core'
export const usePageStore = defineStore('page', () => {
  const isGlobalLoading = ref(false)

  const { store: currentColorMode, system: systemColorMode } = useColorMode({
    disableTransition: !!document.startViewTransition
  })

  const colorModeState = computed(() => {
    if (currentColorMode.value === 'auto') {
      return systemColorMode.value
    } else {
      return currentColorMode.value
    }
  }, {})

  return {
    isGlobalLoading,
    currentColorMode,
    systemColorMode,
    colorModeState
  }
})
