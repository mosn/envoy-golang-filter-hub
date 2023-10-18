<script setup lang="ts">
const modelValue = defineModel<string>()
</script>

<template>
    <div class="search-box">
        <input v-model.trim="modelValue" type="search" placeholder="Search for plugins..." />
    </div>
</template>

<style scoped lang="scss">
.search-box {
    @apply relative mb-5;

    input {
        @apply bg-gray-1 dark-bg-gray-8 w-full;
        @apply border-1.5 border-gray-2 dark-border-gray-7 rounded-xl;
        @apply pl-5 pr-14 py-2 text-lg;
    }

    &::before {
        @apply content-empty absolute top-1/2 right-5;
        @apply i-mingcute-search-3-line;
        @apply text-gray-5 text-xl;
        @apply transition transform-gpu -translate-y-1/2;
        @apply pointer-events-none;
    }
}
</style>