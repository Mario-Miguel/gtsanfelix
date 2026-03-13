<script setup lang="ts">
import { RouterLink, RouterView, useRoute } from 'vue-router'
import { ref, computed } from 'vue'

const route = useRoute()
const menuOpen = ref(false)

const isAdmin = computed(() => route.path.startsWith('/admin'))
</script>

<template>
  <div class="min-h-screen flex flex-col bg-[#221010] text-white">
    <!-- Public top navbar — hidden on admin routes -->
    <header
      v-if="!isAdmin"
      class="fixed top-0 left-0 right-0 z-50 bg-[#1a0a0a]/95 backdrop-blur-sm border-b border-red-900/30"
    >
      <nav class="max-w-7xl mx-auto px-6 py-4 flex items-center justify-between">
        <RouterLink to="/" class="flex items-center gap-2 hover:opacity-80 transition-opacity">
          <span class="material-symbols-outlined text-red-600 text-2xl">theater_comedy</span>
          <span class="font-bold text-lg tracking-wide">G.T. San Félix de Valdesoto</span>
        </RouterLink>

        <!-- Desktop links -->
        <ul class="hidden md:flex items-center gap-7 text-sm font-medium">
          <li>
            <RouterLink to="/" class="hover:text-red-400 transition-colors" active-class="text-red-500">
              Inicio
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/about" class="hover:text-red-400 transition-colors" active-class="text-red-500">
              Compañía
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/repertoire" class="hover:text-red-400 transition-colors" active-class="text-red-500">
              Repertorio
            </RouterLink>
          </li>
          <li>
            <RouterLink to="/calendar" class="hover:text-red-400 transition-colors" active-class="text-red-500">
              Cartelera
            </RouterLink>
          </li>
          <li>
            <RouterLink
              to="/contact"
              class="bg-red-700 hover:bg-red-600 px-4 py-2 rounded transition-colors"
            >
              Contacto
            </RouterLink>
          </li>
        </ul>

        <!-- Mobile hamburger -->
        <button class="md:hidden text-white" @click="menuOpen = !menuOpen" aria-label="Menú">
          <span class="material-symbols-outlined">{{ menuOpen ? 'close' : 'menu' }}</span>
        </button>
      </nav>

      <!-- Mobile menu -->
      <div v-if="menuOpen" class="md:hidden bg-[#1a0a0a] border-t border-red-900/30 px-6 pb-4">
        <ul class="flex flex-col gap-4 pt-4 text-sm font-medium">
          <li><RouterLink to="/" @click="menuOpen = false" class="block hover:text-red-400">Inicio</RouterLink></li>
          <li><RouterLink to="/about" @click="menuOpen = false" class="block hover:text-red-400">Compañía</RouterLink></li>
          <li><RouterLink to="/repertoire" @click="menuOpen = false" class="block hover:text-red-400">Repertorio</RouterLink></li>
          <li><RouterLink to="/calendar" @click="menuOpen = false" class="block hover:text-red-400">Cartelera</RouterLink></li>
          <li><RouterLink to="/contact" @click="menuOpen = false" class="block hover:text-red-400">Contacto</RouterLink></li>
        </ul>
      </div>
    </header>

    <!-- Page content -->
    <main :class="['flex-1', !isAdmin && 'pt-[73px]']">
      <RouterView />
    </main>
  </div>
</template>
