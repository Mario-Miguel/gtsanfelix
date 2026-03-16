<template>
  <div class="bg-[#221010] text-white">
    <!-- Hero -->
    <section class="relative py-20 bg-[#1a0a0a] border-b border-red-900/30 overflow-hidden">
      <div class="absolute inset-0 opacity-10 bg-gradient-to-br from-red-900 to-transparent" />
      <div class="relative z-10 max-w-7xl mx-auto px-6">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-3">Obras</p>
        <h1 class="text-4xl md:text-5xl font-bold mb-4">Nuestro Repertorio</h1>
        <p class="text-gray-400 max-w-xl leading-relaxed">
          Explora nuestra selección de obras maestras listas para ser representadas. Disponibles para giras, festivales y contrataciones privadas.
        </p>
      </div>
    </section>

    <!-- Filters -->
    <section class="bg-[#1a0a0a] border-b border-red-900/20 sticky top-[73px] z-40">
      <div class="max-w-7xl mx-auto px-6 py-3 flex items-center gap-3 flex-wrap">
        <button
          v-for="cat in categories"
          :key="cat"
          :class="[
            'px-4 py-2 rounded text-sm font-medium transition-colors',
            activeCategory === cat
              ? 'bg-red-700 text-white'
              : 'text-gray-400 hover:text-white border border-transparent hover:border-red-900/40',
          ]"
          @click="activeCategory = cat"
        >
          {{ cat }}
        </button>
        <div class="ml-auto flex items-center gap-2">
          <button
            v-for="v in (['grid', 'list'] as const)"
            :key="v"
            :class="['p-2 rounded transition-colors', gridView === v ? 'bg-red-700/40 text-red-400' : 'text-gray-500 hover:text-white']"
            @click="gridView = v"
          >
            <span class="material-symbols-outlined text-sm">{{ v === 'grid' ? 'grid_view' : 'format_list_bulleted' }}</span>
          </button>
        </div>
      </div>
    </section>

    <!-- Content -->
    <section class="max-w-7xl mx-auto px-6 py-12">

      <ApiState :loading="loading" :error="error">
        <!-- Grid view -->
        <div v-if="gridView === 'grid'" class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <article
            v-for="play in filteredPlays"
            :key="play.id"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg overflow-hidden hover:border-red-700/40 transition-colors group"
          >
            <div class="relative overflow-hidden h-52 bg-[#3d1a1a]">
              <img
                v-if="play.imageUrl"
                :src="play.imageUrl"
                :alt="play.title"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
              />
              <div v-else class="w-full h-full flex items-center justify-center">
                <span class="material-symbols-outlined text-red-800 text-6xl">theater_comedy</span>
              </div>
              <div class="absolute inset-0 bg-gradient-to-t from-[#1a0a0a] to-transparent opacity-70" />
              <span class="absolute top-3 right-3 bg-red-700/80 text-white text-xs px-2 py-1 rounded-full">
                {{ play.genre }}
              </span>
            </div>
            <div class="p-5">
              <h3 class="font-bold text-lg mb-1">{{ play.title }}</h3>
              <p class="text-gray-500 text-sm italic mb-3">{{ play.author }}</p>
              <div class="flex items-center gap-4 text-xs text-gray-500 mb-4">
                <span class="flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">schedule</span>
                  {{ play.duration }}
                </span>
                <span class="flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">category</span>
                  {{ play.genre }}
                </span>
              </div>
              <RouterLink
                to="/contact"
                class="w-full border border-red-700/50 hover:bg-red-700 text-red-400 hover:text-white text-sm py-2 rounded transition-colors flex items-center justify-center gap-2"
              >
                <span class="material-symbols-outlined text-sm">event_available</span>
                Solicitar Función
              </RouterLink>
            </div>
          </article>
        </div>

        <!-- List view -->
        <div v-else class="flex flex-col gap-4">
          <article
            v-for="play in filteredPlays"
            :key="play.id"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg overflow-hidden flex hover:border-red-700/40 transition-colors"
          >
            <div class="w-32 h-28 bg-[#3d1a1a] flex-shrink-0 flex items-center justify-center overflow-hidden">
              <img v-if="play.imageUrl" :src="play.imageUrl" :alt="play.title" class="w-full h-full object-cover" />
              <span v-else class="material-symbols-outlined text-red-800 text-4xl">theater_comedy</span>
            </div>
            <div class="p-5 flex items-center justify-between flex-1 flex-wrap gap-4">
              <div>
                <span class="text-xs bg-red-700/30 text-red-400 px-2 py-0.5 rounded-full">{{ play.genre }}</span>
                <h3 class="font-bold mt-1">{{ play.title }}</h3>
                <p class="text-gray-500 text-sm italic">{{ play.author }}</p>
              </div>
              <div class="flex items-center gap-6 text-xs text-gray-500">
                <span class="flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">schedule</span>
                  {{ play.duration }}
                </span>
              </div>
              <RouterLink
                to="/contact"
                class="border border-red-700/50 hover:bg-red-700 text-red-400 hover:text-white text-sm px-4 py-2 rounded transition-colors flex items-center gap-1"
              >
                <span class="material-symbols-outlined text-sm">event_available</span>
                Solicitar
              </RouterLink>
            </div>
          </article>
        </div>

        <!-- Empty state -->
        <div v-if="filteredPlays.length === 0 && !loading" class="text-center py-20 text-gray-600">
          <span class="material-symbols-outlined text-5xl mb-3 block">search_off</span>
          <p>No hay obras en esta categoría.</p>
        </div>
      </ApiState>
    </section>

    <!-- CTA -->
    <section class="bg-[#1a0a0a] border-t border-red-900/30 py-16">
      <div class="max-w-2xl mx-auto px-6 text-center">
        <h2 class="text-2xl font-bold mb-4">¿Quieres contratarnos?</h2>
        <p class="text-gray-400 mb-8">
          Llevamos nuestras obras a festivales, centros culturales y eventos privados. Contáctanos para más información.
        </p>
        <RouterLink to="/contact" class="bg-red-700 hover:bg-red-600 px-8 py-3 rounded font-medium transition-colors inline-block">
          Solicitar Información
        </RouterLink>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { RouterLink } from 'vue-router'
import { playsApi, type Play } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'

const { data: plays, loading, error } = useApiRequest(() => playsApi.list(), {
  initialData: [] as Play[],
})
const activeCategory = ref('Todos')
const gridView = ref<'grid' | 'list'>('grid')

const categories = computed(() => {
  const genres = [...new Set((plays.value ?? []).map((p) => p.genre).filter(Boolean))]
  return ['Todos', ...genres.sort()]
})

const filteredPlays = computed(() => {
  const list = plays.value ?? []
  if (activeCategory.value === 'Todos') return list
  return list.filter((p) => p.genre === activeCategory.value)
})
</script>
