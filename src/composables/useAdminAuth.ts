import { computed } from 'vue'
import { useRouter } from 'vue-router'

export const TOKEN_KEY = 'admin_token'
export const USER_KEY = 'admin_user'

export interface AdminUser {
  id: number
  email: string
  name: string
  role: string
}

export function useAdminAuth() {
  const router = useRouter()

  const adminUser = computed<AdminUser | null>(() => {
    const raw = localStorage.getItem(USER_KEY)
    return raw ? (JSON.parse(raw) as AdminUser) : null
  })

  function setSession(token: string, user: AdminUser) {
    localStorage.setItem(TOKEN_KEY, token)
    localStorage.setItem(USER_KEY, JSON.stringify(user))
  }

  function logout() {
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(USER_KEY)
    router.push('/admin/login')
  }

  return { adminUser, setSession, logout }
}
