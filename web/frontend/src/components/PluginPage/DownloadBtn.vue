<script setup lang="ts">
import { ref } from 'vue';
import { type PluginDownloadItem } from '@/apis/plugin';

defineProps<{
    downloads: PluginDownloadItem[]
}>()

const isDownloadPanelOpen = ref(false)

</script>

<template>
    <div class="download-btn-wrapper" :class="{ open: isDownloadPanelOpen }">
        <p class="download-btn" @click="isDownloadPanelOpen = !isDownloadPanelOpen">
            Download
        <div class="icon i-mingcute-down-line"></div>
        </p>
        <div class="download-panel" @click="isDownloadPanelOpen = false">
            <a v-for="download of downloads" :key="download.type" :href="download.url" target="_blank"
                rel="noopener noreferrer">
                <div class="i-mingcute-download-2-line"></div> {{ download.type }}
            </a>
        </div>
    </div>
</template>

<style scoped lang="scss">
.download-btn-wrapper {
    @apply select-none relative;

    .download-btn {
        @apply text-primary border-primary/30 border-1.5 rounded-xl;
        @apply px-2 py-1 flex items-center gap-1;
        @apply cursor-pointer transition transform-gpu;

        .icon {
            @apply transition transform-gpu;
        }

        &:hover {
            @apply bg-primary/5;
        }

        &:active {
            @apply scale-90;
        }

    }

    .download-panel {
        @apply absolute right-0 top-10 bg-white dark-bg-black;
        @apply text-primary border-primary/30 border-1.5 rounded-xl;
        @apply flex flex-col overflow-hidden;
        @apply transition transform-gpu -translate-y-1/2 opacity-0 scale-y-20 pointer-events-none;
        @apply z-9;

        a {
            @apply whitespace-nowrap text-right;
            @apply flex items-center gap-1 px-3 py-2;
            @apply underline underline-dashed underline-primary/40;

            &:hover {
                @apply bg-primary/5;
            }

            &:active {
                @apply bg-primary/10;
            }
        }
    }

    &.open {
        .icon {
            @apply rotate-180;
        }

        .download-panel {
            @apply translate-y-0 opacity-100 scale-y-100 pointer-events-auto;
        }
    }
}
</style>