import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Login from '@/views/login/login.vue';
import Index from '@/views/home/index.vue';


const routes: Array<RouteRecordRaw> = [
  {
    path: '/monitor',
    name: '/monitor/index',
    component: Index
  },
  {
    path: '/monitor/login',
    name: '/monitor/login',
    component: Login
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
