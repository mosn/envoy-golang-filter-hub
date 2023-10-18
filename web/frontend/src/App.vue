<script setup lang="ts">
import { RouterView } from 'vue-router'
import GlobalHeader from '@/components/GlobalHeader.vue'


</script>

<template>
  <GlobalHeader />
  <main>
    <div class="limit">
      <RouterView v-slot="{ Component }">
        <KeepAlive :include="['HomePage']">
          <component :is="Component" />
        </KeepAlive>
      </RouterView>
    </div>
  </main>
</template>

<style scoped lang="scss">
main {
  .limit {
    @apply max-w-1200px w-full h-full;
    @apply mx-auto;
    @apply p-3;
    @apply min-h-[calc(100svh-56px)];

    &>* {
      @apply h-full;
    }
  }
}
</style>
