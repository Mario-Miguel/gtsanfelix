import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', name: 'Home', component: () => import('../components/HomeView.vue') },
  { path: '/about', name: 'About', component: () => import('../components/AboutView.vue') },
  { path: '/repertoire', name: 'Repertoire', component: () => import('../components/RepertoireView.vue') },
  { path: '/calendar', name: 'Calendar', component: () => import('../components/CalendarView.vue') },
  { path: '/contact', name: 'Contact', component: () => import('../components/ContactView.vue') },
  { path: '/admin/login', name: 'AdminLogin', component: () => import('../components/AdminLoginView.vue') },
  { path: '/admin', name: 'Admin', component: () => import('../components/AdminView.vue'), meta: { requiresAuth: true } },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

// Navigation guard: redirect to login if no JWT token
router.beforeEach((to) => {
  if (to.meta.requiresAuth && !localStorage.getItem('admin_token')) {
    return '/admin/login'
  }
})

export default router
