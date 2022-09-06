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
      // 管理员
      // 管理员列表
      {
        path: "/monitor/admin/list",
        name: "/monitor/admin/list",
        component: () => import("@/views/admin/admin.vue"),
        meta: {
          title: "管理员列表"
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
      },
      // 客户端机器列表
      {
        path: "/monitor/machine/list",
        name: "/monitor/machine/list",
        component: () => import("@/views/machine/machine.vue"),
      },
      {
        path: "/monitor/machine/dial",
        name: "/monitor/machine/dial",
        component: () => import("@/views/machine/dial.vue"),
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
