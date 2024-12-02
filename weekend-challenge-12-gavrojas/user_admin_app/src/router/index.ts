import { createRouter, createWebHistory } from 'vue-router'
import AuthView from '../views/AuthView.vue'
import EntriesView from '../views/EntriesView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', component: AuthView }, 
    { path: '/entries', component: EntriesView }
  ]
})


export default router
