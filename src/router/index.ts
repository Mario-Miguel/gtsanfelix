import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', name: 'Home', component: () => import('../components/HomeView.vue') },
  { path: '/about', name: 'About', component: () => import('../components/AboutView.vue') },
  { path: '/contact', name: 'Contact', component: () => import('../components/ContactView.vue') },
  {
    path: '/calendar',
    name: 'Calendar',
    component: () => import('../components/CalendarView.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
