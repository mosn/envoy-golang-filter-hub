import '@unocss/reset/tailwind.css'
import './assets/main.scss'
import 'virtual:uno.css'
import '@/components/Modal'

import { createApp } from 'vue'
import shadow from 'vue-shadow-dom'

import App from './App.vue'
import router from '@/router'
import { createPinia } from 'pinia'

const app = createApp(App)

app.use(createPinia())
app.use(shadow)
app.use(router)

app.mount('#app')
