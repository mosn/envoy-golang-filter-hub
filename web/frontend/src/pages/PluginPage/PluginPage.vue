<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import LoadingIcon from '@/components/LoadingIcon.vue';
import { type PluginData, getPlugin } from '@/apis/plugin'
import { usePageStore } from '@/store/page';
import { useRoute } from 'vue-router';
import DownloadBtn from '@/components/PluginPage/DownloadBtn.vue';
import TabView, { type TabOption } from '@/components/TabView.vue';
import VersionsView from './views/VersionsView.vue';
import HtmlView from './views/HtmlView.vue';

const pageStore = usePageStore()

const loading = computed({
    get() {
        return pageStore.isGlobalLoading
    },
    set(value) {
        pageStore.isGlobalLoading = value
    }
})
const error = ref(true)
const plugin = ref<PluginData>()
const { id } = useRoute().params

const fetchPlugin = async () => {
    loading.value = true
    error.value = false

    const result = await getPlugin(id as string)
    if (!result) {
        error.value = true
    } else {
        result.versions.sort((a, b) => {
            let verA = a.version.split('.').map(v => parseInt(v))
            let verB = b.version.split('.').map(v => parseInt(v))
            for (let i = 0; i < Math.max(verA.length, verB.length); i++) {
                if (verA[i] === undefined) {
                    return 1
                } else if (verB[i] === undefined) {
                    return -1
                } else if (verA[i] > verB[i]) {
                    return -1
                } else if (verA[i] < verB[i]) {
                    return 1
                }
            }
            return 0
        })
        result.versions.forEach(version => {
            version.downloads.sort((a, b) => {
                return a.type.localeCompare(b.type)
            })
        })
        plugin.value = result
        document.title = `${result.name} | Envoy Hub`
    }
    loading.value = false
}

onMounted(() => {
    fetchPlugin()
})

const tabViewOptions = computed(() => {
    const options: TabOption[] = []

    if (plugin.value?.overview) {
        options.push({
            icon: 'i-mingcute-document-2-line',
            label: 'Overview',
            name: 'overview'
        })
    }
    if (plugin.value?.config) {
        options.push({
            icon: 'i-mingcute-tool-line',
            label: 'Config',
            name: 'config'
        })
    }
    if (plugin.value?.changelog) {
        options.push({
            icon: 'i-mingcute-paper-line',
            label: 'Changelog',
            name: 'changelog'
        })
    }
    if (plugin.value?.versions) {
        options.push({
            icon: 'i-mingcute-tag-2-line',
            label: `${plugin.value.versions.length} Versions`,
            name: 'versions'
        })
    }
    return options
})
</script>

<template>
    <div v-if="error">
        <div class="flex justify-center items-center flex-col h-[calc(100dvh-98px)]">
            <div class="i-mingcute-close-circle-line text-4xl text-red-7"></div>
            <p class="text-red-7 text-lg">Network Error</p>
            <button class="g-btn mt-10" @click="fetchPlugin()">
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
    <div v-else-if="plugin">
        <div class="flex items-start">
            <div class="space-y-2 flex-grow mb-2">
                <h1 class="font-bold text-2xl md-text-3xl">{{ plugin.name }}</h1>
                <div class="space-x-2 text-sm">
                    <span class="g-capsule">{{ plugin.version }}</span>
                    <span class="g-capsule">#{{ plugin.category }}</span>
                </div>
            </div>
            <DownloadBtn v-if="plugin.versions.length > 0" :downloads="plugin.versions[0].downloads" />
        </div>
        <p class="text-gray-5">{{ plugin.description }}</p>
        <TabView :tabs="tabViewOptions" class="mt-5">
            <template v-if="plugin.overview" #view-overview>
                <HtmlView :html="plugin.overview" />
            </template>
            <template v-if="plugin.config" #view-config>
                <HtmlView :html="plugin.config" />
            </template>
            <template v-if="plugin.changelog" #view-changelog>
                <HtmlView :html="plugin.changelog" />
            </template>
            <template #view-versions>
                <VersionsView :versions="plugin.versions" />
            </template>
        </TabView>
    </div>
</template>

<style scoped lang="scss"></style>