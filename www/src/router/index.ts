import { createRouter, createWebHashHistory } from 'vue-router'
import { useAuth } from '@/composables/useAuth'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/overview',
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { public: true },
    },
    {
      path: '/overview',
      name: 'overview',
      component: () => import('@/views/OverviewView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/heating',
      name: 'heating',
      component: () => import('@/views/ControlView.vue'),
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach(async (to) => {
  const { isLoggedIn, checkSession } = useAuth()

  if (to.meta.requiresAuth) {
    if (!isLoggedIn.value) {
      const ok = await checkSession()
      if (!ok) {
        return { name: 'login' }
      }
    }
  }

  if (to.meta.public && isLoggedIn.value) {
    return { name: 'overview' }
  }
})

export default router
