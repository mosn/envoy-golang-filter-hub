<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import LoadingIcon from '@/components/LoadingIcon.vue';
import CategoriesSidebar from '@/components/HomePage/CategoriesSidebar.vue';
import { type PluginItem, getPluginList } from '@/apis/plugin'
import SearchBar from '@/components/HomePage/SearchBar.vue';
import PluginCard from '@/components/HomePage/PluginCard.vue';

const loading = ref(true)
const error = ref(false)
const plugins = ref<PluginItem[]>([])

const categories = computed(() => {
    const categories = new Set<string>()
    plugins.value.forEach(plugin => {
        categories.add(plugin.category)
    })
    return Array.from(categories)
})

const currentCategory = ref('All')

const currentSearchKeyword = ref('')

const filteredPlugins = computed(() => {
    return plugins.value.filter(plugin => {
        if (currentCategory.value === 'All') {
            return true
        } else {
            return plugin.category === currentCategory.value
        }
    }).filter(plugin => {
        if (currentSearchKeyword.value === '') {
            return true
        } else {
            let key = currentSearchKeyword.value.toLowerCase();
            return plugin.name.toLowerCase().includes(key) || plugin.description.toLowerCase().includes(key)
        }
    })
})

const fetchPluginList = async () => {
    loading.value = true
    error.value = false

    const result = await getPluginList()
    if (!result) {
        error.value = true
    } else {
        plugins.value = result
    }
    loading.value = false
}

onMounted(() => {
    fetchPluginList()
})
</script>

<template>
    <div v-if="error">
        <div class="flex justify-center items-center flex-col h-[calc(100dvh-98px)]">
            <div class="i-mingcute-close-circle-line text-4xl text-red-7"></div>
            <p class="text-red-7 text-lg">Network Error</p>
            <button class="g-btn mt-10" @click="fetchPluginList()">
                <div class="i-mingcute-refresh-1-line"></div>
                Retry
            </button>
        </div>
    </div>
    <div v-else-if="loading">
        <div class="flex justify-center items-center h-[calc(100dvh-98px)]">
            <LoadingIcon class="text-black dark:text-white text-2xl" />
        </div>
    </div>
    <div class="main-wrapper" v-else>
        <CategoriesSidebar v-model:currentCategory="currentCategory" :categories="categories" />
        <div class="w-full">
            <SearchBar v-model="currentSearchKeyword" />
            <div class="plugin-list" aria-live="polite">
                <span v-if="filteredPlugins.length === 0" class="text-center col-span-full text-gray-5">No Result</span>
                <template v-else>
                    <template v-for="plugin of filteredPlugins" :key="plugin.path_name">
                        <PluginCard :plugin="plugin" />
                    </template>
                </template>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.main-wrapper {
    @apply flex gap-10;

    .plugin-list {
        @apply grid grid-cols-1 sm-grid-cols-2 md-grid-cols-3 gap-5 w-full;
    }
}
</style>