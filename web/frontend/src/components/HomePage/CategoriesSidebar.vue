<script setup lang="ts">
defineProps<{
    categories: string[]
}>()

const currentCategory = defineModel<string>('currentCategory', { local: true, default: 'All' })
</script>

<template>
    <nav class="categories-wrapper">
        <header class="text-xl font-bold mb-3">Categories</header>
        <div class="flex flex-col gap-1">
            <p tabindex="0" class="category-item" v-for="category of ['All', ...categories]" :key="category"
                :class="{ active: currentCategory === category }" @click="currentCategory = category"
                :aria-current="currentCategory === category">
                {{ category }}
            </p>
        </div>
    </nav>
</template>

<style scoped lang="scss">
.categories-wrapper {
    @apply sticky top-68px self-start;
    @apply hidden sm-block;
    @apply flex-shrink-0;
    @apply select-none;
    @apply border-1.5 border-gray-2 dark-border-gray-7 pl-5 pr-10 py-3 rounded-xl;

    .category-item {
        @apply cursor-pointer relative;
        @apply transition;
        @apply py-1 px-2 mr-2;

        &::before {
            @apply content-empty absolute top-1 bottom-1 left-0 w-4px;
            @apply transition bg-transparent rounded-full;
        }

        &:hover {
            @apply text-primary;
        }

        &.active {
            @apply pointer-events-none text-primary;

            &::before {
                @apply bg-primary;
            }
        }
    }
}
</style>