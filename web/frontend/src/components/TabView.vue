<script setup lang="ts">
import { useUrlSearchParams } from '@vueuse/core';
import { computed, nextTick, reactive, ref, watch } from 'vue';

export interface TabOption {
    icon: string
    label: string
    name: string
}

export interface TabViewProps {
    tabs: TabOption[]
}

const props = defineProps<TabViewProps>()

const TabViewRef = ref<HTMLDivElement>()

const currentTabBound = reactive({
    x: 0,
    width: 0
})

const params = useUrlSearchParams('history')

const activeTabName = computed<string>({
    get() {
        let activeTab = params.activeTab;

        if (typeof activeTab !== 'string' || props.tabs.findIndex(tab => tab.name === activeTab) === -1) {
            activeTab = props.tabs[0].name
            params.activeTab = activeTab
        }
        return activeTab
    },
    set(value) {
        params.activeTab = value
    }
})

watch(activeTabName, async () => {
    await nextTick()
    // const firstTab = TabViewRef.value?.querySelector('.tab-item:first-child') as HTMLElement
    const currentTab = TabViewRef.value?.querySelector('.tab-item.active') as HTMLElement
    if (currentTab) {
        currentTabBound.x = currentTab.offsetLeft
        currentTabBound.width = currentTab.offsetWidth
    }
}, {
    immediate: true
})
</script>

<template>
    <div ref="TabViewRef" class="tab-view">
        <div class="tab-bar special-bg">
            <div class="tab-group">
                <div v-for="tab of tabs" :key="tab.name" class="tab-item" :class="{ active: activeTabName === tab.name }"
                    @click="activeTabName = tab.name" tabindex="0">
                    <div class="tab-btn">
                        <div class="icon" :class="tab.icon"></div>
                        <p class="label">{{ tab.label }}</p>
                    </div>
                </div>
            </div>
            <div class="tab-bar-line">
                <div class="current-tab-line" :style="{
                    '--x': `${currentTabBound.x}px`,
                    '--width': currentTabBound.width / 2
                }"></div>
            </div>
        </div>
        <div class="tab-content">
            <slot :name="`view-${activeTabName}`"></slot>
        </div>
    </div>
</template>

<style scoped lang="scss">
.tab-view {
    @apply w-full;
}

.tab-bar {
    @apply overflow-x-auto select-none sticky top-56px;
    @apply -mx-3;
    @apply z-11;

    .tab-group {
        @apply flex px-3;
    }

    .tab-item {
        @apply cursor-pointer text-lg whitespace-nowrap;
        @apply px-1 py-1.5;
        @apply transition;
        @apply text-gray-6;

        .tab-btn {
            @apply transition;
            @apply flex items-center gap-1;
            @apply px-3 py-0.5 rounded-xl hover-bg-gray-1 dark-hover-bg-gray/10;
        }

        &.active {
            @apply pointer-events-none;

            .tab-btn {
                @apply text-primary bg-primary-light dark-bg-primary-dark;
            }
        }
    }

    .tab-bar-line {
        @apply w-full h-1.5px bg-gray-2 dark-bg-gray-8/70 relative;

        .current-tab-line {
            @apply absolute left-0 bottom-0 h-2px bg-primary w-2px;
            @apply transition;
            transform-origin: left;
            transform: translateX(var(--x)) scaleX(var(--width));
        }
    }

}
</style>