<template>
  <div class="min-h-screen bg-[#160808] flex items-center justify-center px-4">
    <div class="w-full max-w-sm">
      <!-- Logo -->
      <div class="text-center mb-8">
        <span class="material-symbols-outlined text-red-600 text-5xl">theater_comedy</span>
        <h1 class="text-2xl font-bold mt-3">G.T. San Félix de Valdesoto</h1>
        <p class="text-gray-500 text-sm mt-1">Panel de Administración</p>
      </div>

      <!-- Card -->
      <div class="bg-[#1a0a0a] border border-red-900/30 rounded-lg p-8">
        <h2 class="text-lg font-semibold mb-6">Iniciar Sesión</h2>

        <form @submit.prevent="handleLogin" class="flex flex-col gap-4">
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Email</label>
            <input
              v-model="email"
              type="email"
              required
              autocomplete="username"
              placeholder="admin@telonabierto.com"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-3 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors"
            />
          </div>

          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Contraseña</label>
            <input
              v-model="password"
              type="password"
              required
              autocomplete="current-password"
              placeholder="••••••••"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-3 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors"
            />
          </div>

          <!-- Error -->
          <div
            v-if="error"
            class="bg-red-900/20 border border-red-800/40 text-red-400 rounded px-4 py-3 text-sm flex items-center gap-2"
          >
            <span class="material-symbols-outlined text-sm">error</span>
            {{ error }}
          </div>

          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-red-700 hover:bg-red-600 disabled:opacity-50 disabled:cursor-not-allowed py-3 rounded font-medium text-sm transition-colors flex items-center justify-center gap-2 mt-2"
          >
            <span v-if="loading" class="material-symbols-outlined text-sm animate-spin">progress_activity</span>
            {{ loading ? 'Entrando...' : 'Entrar' }}
          </button>
        </form>
      </div>

      <p class="text-center text-xs text-gray-700 mt-6">
        Credenciales de prueba: <span class="text-gray-500">admin@telonabierto.com / admin123</span>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login } from '../composables/useAdminApi'

const router = useRouter()
const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

async function handleLogin() {
  loading.value = true
  error.value = ''
  try {
    const res = await login(email.value, password.value)
    localStorage.setItem('admin_token', res.token)
    localStorage.setItem('admin_user', JSON.stringify(res.user))
    router.push('/admin')
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Error de conexión con el servidor'
  } finally {
    loading.value = false
  }
}
</script>
