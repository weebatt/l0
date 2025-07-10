import { createRouter, createWebHistory } from 'vue-router'
import OrderView from '@/views/OrderView.vue'

const routes = [
  {
    path: '/order/:order_uid?',
    name: 'Order',
    component: OrderView
  },
  {
    path: '/',
    redirect: '/order'
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router