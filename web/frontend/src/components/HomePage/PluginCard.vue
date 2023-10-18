<script setup lang="ts">
import { type PluginItem } from '@/apis/plugin';
defineProps<{
    plugin: PluginItem
}>()
</script>

<template>
    <RouterLink class="plugin-item" :to="`/plugin/${plugin.path_name}`">
        <div class="flex justify-between items-center">
            <p class="name">{{ plugin.name }}</p>
            <p class="version">{{ plugin.version }}</p>
        </div>
        <p class="description">{{ plugin.description }}</p>
    </RouterLink>
</template>

<style scoped lang="scss">
.plugin-item {
    @apply p-3;
    @apply bg-white dark-bg-black;
    @apply border-1.5 border-gray-2 dark-border-gray-7 rounded-xl;
    @apply cursor-pointer transition;

    .name {
        @apply font-bold text-lg text-primary;
        @apply truncate;
    }

    .version {
        @apply g-capsule;
        @apply text-xs;
        @apply truncate;
    }

    .description {
        @apply mt-1;
        @apply text-sm text-gray-4 dark-text-gray-5;
        @apply overflow-hidden text-ellipsis;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;
        display: -webkit-box;
    }

    &:hover {
        @apply bg-primary/3;
    }
}
</style>