<script setup lang="ts">
import { type PluginVersionItem } from "@/apis/plugin"
import { onMounted } from "vue";
defineProps<{
    versions: PluginVersionItem[]
}>()

const formatDate = (date: string) => {
    const parsedDate = new Date(date)
    return parsedDate.toLocaleDateString()
}

const shortenHash = (hash: string) => {
    return hash.slice(0, 7)
}

onMounted(() => {
    setTimeout(() => {
        window.scrollTo({
            top: 0,
            behavior: 'smooth'
        })
    }, 0)
})
</script>

<template>
    <div class="table-wrapper">
        <table>
            <thead class="special-bg">
                <tr>
                    <th class="text-left"><span>Version</span></th>
                    <th class="text-left"><span>Last Update</span></th>
                    <th class="text-right"><span>Downloads</span></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="version in versions" :key="version.version">
                    <td class="text-left flex gap-x-2 flex-col sm:(flex-row items-center)"><span class="opacity-70">{{
                        version.version }}</span>
                        <a class="commit-link" :href="version.commit_url" target="_blank" rel="noopener noreferrer">{{
                            shortenHash(version.commit_hash) }}</a>
                    </td>
                    <td class="text-left"><span class="opacity-70">{{ formatDate(version.last_update) }}</span></td>
                    <td class="text-right space-x-3">
                        <a class="download-link" v-for="download of version.downloads" :key="download.type"
                            :href="download.url" target="_blank" rel="noopener noreferrer">{{ download.type }}</a>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<style scoped lang="scss">
.table-wrapper {
    @apply mt-5 max-w-screen;
}

table {
    @apply border-collapse w-full;
}

thead {
    @apply sticky top-[calc(56px+44px)];
    @apply z-7;
    @apply text-lg;

    span {
        @apply block py-1;
        @apply border-b-1.5 border-gray-2 dark-border-gray-8/70;
    }
}



tbody {
    font-family: Consolas, monaco, monospace;

    td {
        @apply py-2;
        @apply border-b-1.5 border-gray-1 dark-border-gray-8/50;

        &:not(:last-child) {
            @apply pr-2;
        }
    }

    .commit-link {
        @apply text-sm;
        @apply text-gray-4 hover-text-gray-5;
        @apply dark-text-gray-5 dark-hover-text-gray-4;
        @apply hover-underline;
    }
}

th,
td {
    @apply whitespace-nowrap;
}

.download-link {
    @apply text-sm;
    @apply hover-underline;
    @apply text-primary;
}
</style>