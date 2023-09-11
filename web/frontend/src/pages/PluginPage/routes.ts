import type { RouteRecordRaw } from 'vue-router'

export const routes: RouteRecordRaw[] = [
  {
    path: '/plugin/:id',
    component: () => import('./PluginPage.vue'),
    meta: {
      title: 'Plugin'
    }
  }
]
