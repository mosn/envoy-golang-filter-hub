import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

interface WrapperedRouteConfig {
  routes?: RouteRecordRaw[]
}

const routeModules = import.meta.glob<WrapperedRouteConfig>('../pages/**/routes.ts', { eager: true });

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home'
  }
]

for (const routeModule of Object.values(routeModules)) {
  if (routeModule['routes']) {
    routes.push(...routeModule['routes'])
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.afterEach((to, from, failure) => {
  if (!failure) {
    window.document.title = `${to.meta.title} | Envoy Hub` || 'Envoy Hub'
  }
})

export default router
