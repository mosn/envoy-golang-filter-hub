import type { RouteRecordRaw } from 'vue-router';

export const routes: RouteRecordRaw[] = [
    {
        path: '/home',
        alias: '/',
        component: () => import('./HomePage.vue'),
        meta: {
            title: 'Plugins',
        }
    }
]
