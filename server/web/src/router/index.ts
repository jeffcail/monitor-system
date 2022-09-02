import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Login from '@/views/login/login.vue';
import Index from '@/views/home/index.vue';


const routes: Array<RouteRecordRaw> = [
  {
    path: '/monitor',
    name: '/monitor/index',
    component: Index,
    redirect: '/monitor/board',
    meta: {
      title: "首页看板"
    },
    children: [
      // 首页看板
      {
        path: "/monitor/board",
        name: "/monitor/board",
        component: () => import("@/views/board/board.vue"),
        meta: {
          title: '首页看板',
        }
      },
      // 服务检测列表
      {
        path: "/monitor/serve/list",
        name: "/monitor/serve/list",
        component: () => import("@/views/serve/serve.vue"),
        meta: {
          title: "服务检测"
        }
      }
    ]
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
