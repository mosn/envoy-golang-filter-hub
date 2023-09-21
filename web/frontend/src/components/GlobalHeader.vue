<script setup lang="ts">
import DarkModeSwitch from './DarkModeSwitch.vue';
import LoadingIcon from './LoadingIcon.vue';
import { usePageStore } from '@/store/page'

const pageStore = usePageStore()

</script>

<template>
    <header class="special-bg">
        <section class="limit">
            <RouterLink to="/">
                <section class="logo">
                    <span>Envoy</span> Filter Hub
                </section>
            </RouterLink>
            <section class="tools">
                <div v-if="pageStore.isGlobalLoading" class="text-primary text-lg">
                    <LoadingIcon />
                </div>
                <DarkModeSwitch />
                <a href="https://github.com/NX-Official/envoy-golang-filter-hub" target="_blank" class="link-item">
                    <div class="i-mingcute-github-line"></div>
                </a>
            </section>
        </section>
    </header>
</template>

<style scoped lang="scss">
header {
    @apply w-full;
    @apply flex justify-center;
    @apply sticky top-0 left-0 right-0;
    @apply z-10;
    // @apply ring-1 ring-gray-2 dark-ring-gray-8;

    .limit {
        @apply max-w-1200px w-full h-full;
        @apply flex justify-between items-center;
        @apply p-3;
    }

    .logo {
        @apply text-2xl font-bold;
    }

    .tools {
        @apply flex items-center;
        // TODO: 不使用 gap 是为了兼容低版本浏览器
        @apply space-x-3;

        .link-item {
            @apply w-7 h-7 text-lg;
            @apply flex justify-center items-center;
            @apply text-primary rounded-full;
            @apply transition-colors;
            @apply relative;

            &::before {
                @apply content-empty absolute top-0 left-0 right-0 bottom-0;
                @apply border-1.5 border-primary/0 rounded-full;
                @apply transition transform-gpu;
            }

            &:hover {
                @apply text-white bg-primary;

                &::before {
                    @apply scale-130 border-primary/30;
                }
            }

            &:active::before {
                @apply scale-100;
            }
        }
    }
}
</style>