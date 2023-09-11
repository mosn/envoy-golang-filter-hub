<script setup lang="ts">
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue';
import markdownStyle from '@/assets/markdown-style.scss?inline'
import { ShadowStyle } from 'vue-shadow-dom';
import { usePageStore } from '@/store/page';
import { markdownFallbackUrl } from '@/utils/env';

const pageStore = usePageStore()

const props = defineProps<{
    html: string
}>()

const shadowRef = ref<{
    shadow_root: ShadowRoot
}>()

watch(() => props.html, async (html) => {
    await nextTick()
    shadowRef.value?.shadow_root.querySelectorAll<HTMLAnchorElement>('a[href^="#"]').forEach((a) => {
        a.addEventListener('click', function (e) {
            e.preventDefault()
            location.hash = a.hash
        })
    })
    shadowRef.value?.shadow_root.querySelectorAll<HTMLAnchorElement>('a[href^="http"]').forEach((a) => {
        a.addEventListener('click', function (e) {
            e.preventDefault()
            window.open(a.href, '_blank')
        })
    })
    shadowRef.value?.shadow_root.querySelectorAll<HTMLAnchorElement>('a[href^="/"]').forEach((a) => {
        a.addEventListener('click', function (e) {
            e.preventDefault()
            const href = a.getAttribute('href')
            const url = new URL(href!, markdownFallbackUrl)
            window.open(url, '_blank')
        })
    })
    shadowRef.value?.shadow_root.querySelectorAll<HTMLAnchorElement>('a[href^="."]').forEach((a) => {
        a.addEventListener('click', function (e) {
            e.preventDefault()
            const href = a.getAttribute('href')
            const url = new URL(href!, markdownFallbackUrl)
            window.open(url, '_blank')
        })
    })
}, { immediate: true })

let isHashChanged = false;


onMounted(async () => {
    await nextTick()
    hashchange()
    window.addEventListener('hashchange', hashchange)
})

onUnmounted(function () {
    window.removeEventListener('hashchange', hashchange)
    if (isHashChanged) {
        location.hash = ''
    }
})

function hashchange() {
    const OFFSET = 56 + 44 + 10
    let hash: string

    try {
        hash = decodeURIComponent(location.hash.slice(1)).toLowerCase()
    } catch {
        return
    }

    if (hash) {
        isHashChanged = true
    } else if(!isHashChanged) {
        setTimeout(() => {
            window.scrollTo({
                top: 0,
                behavior: 'smooth'
            })
        }, 0)
    }

    const name = 'user-content-' + hash
    const target =
        shadowRef.value?.shadow_root.getElementById(name)

    if (target) {
        const top = target.offsetTop - OFFSET
        window.setTimeout(function () {
            window.scrollTo({
                top,
                behavior: 'smooth'
            })
        }, 0)
    }
}
</script>

<template>
    <div class="html-view">
        <shadow-root abstract ref="shadowRef">
            <ShadowStyle>
                {{ markdownStyle }}
            </ShadowStyle>
            <div :class="{
                'dark': pageStore.colorModeState === 'dark'
            }">
                <article class="markdown-body" v-html="html"></article>
            </div>
        </shadow-root>
    </div>
</template>

<style scoped>
.html-view {
    overflow-x: hidden;
    @apply -mx-3;
}
</style>