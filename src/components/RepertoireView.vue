<template>
  <div class="bg-[#f8f6f6] text-slate-900">
    <!-- Hero -->
    <section class="relative py-20 bg-[#C2B280]/10 border-b border-[#C2B280]/20 overflow-hidden">
      <div class="absolute inset-0 opacity-10 bg-gradient-to-br from-[#8e7a52] to-transparent" />
      <div class="relative z-10 max-w-7xl mx-auto px-6">
        <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">Obres</p>
        <h1 class="text-4xl md:text-5xl font-bold mb-4">El nuesu repertoriu</h1>
        <p class="text-slate-600 max-w-xl leading-relaxed">
          Estes son les nueses obres que ponemos y pusimos en escena.
        </p>
      </div>
    </section>

    <!-- Content -->
    <section class="max-w-7xl mx-auto px-6 py-12">

      <ApiState :loading="loading" :error="error">
        <!-- Grid view -->
        <div v-if="gridView === 'grid'" class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <article v-for="play in plays" :key="play.id"
            class="bg-white border border-[#C2B280]/10 rounded-lg overflow-hidden hover:border-[#C2B280]/40 transition-colors group">
            <div class="relative overflow-hidden h-52 bg-[#C2B280]/5">
              <img v-if="play.imageUrl" :src="play.imageUrl" :alt="play.title"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
              <div v-else class="w-full h-full flex items-center justify-center">
                <span class="material-symbols-outlined text-[#8e7a52] text-6xl">theater_comedy</span>
              </div>
              <div class="absolute inset-0 bg-gradient-to-t from-[#f8f6f6] to-transparent opacity-70" />
            </div>
            <div class="p-5">
              <h3 class="font-bold text-lg mb-1">{{ play.title }}</h3>
              <p class="text-slate-500 text-sm italic mb-3">{{ play.author }}</p>
              <div class="flex items-center gap-4 text-xs text-slate-500 mb-4">
                <span class="flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">schedule</span>
                  {{ play.duration }}
                </span>
              </div>
              <RouterLink to="/contact"
                class="w-full border border-[#C2B280]/50 hover:bg-[#C2B280] text-[#C2B280] hover:text-slate-900 text-sm py-2 rounded transition-colors flex items-center justify-center gap-2">
                <span class="material-symbols-outlined text-sm">event_available</span>
                Solicitar Función
              </RouterLink>
            </div>
          </article>
        </div>

        <!-- List view -->
        <div v-else class="flex flex-col gap-4">
          <article v-for="play in plays" :key="play.id"
            class="bg-white border border-[#C2B280]/10 rounded-lg overflow-hidden flex hover:border-[#C2B280]/40 transition-colors">
            <div class="w-32 h-28 bg-[#C2B280]/5 flex-shrink-0 flex items-center justify-center overflow-hidden">
              <img v-if="play.imageUrl" :src="play.imageUrl" :alt="play.title" class="w-full h-full object-cover" />
              <span v-else class="material-symbols-outlined text-[#8e7a52] text-4xl">theater_comedy</span>
            </div>
            <div class="p-5 flex items-center justify-between flex-1 flex-wrap gap-4">
              <div>
                <span class="text-xs bg-[#C2B280]/20 text-[#8e7a52] px-2 py-0.5 rounded-full">{{ play.genre }}</span>
                <h3 class="font-bold mt-1">{{ play.title }}</h3>
                <p class="text-slate-500 text-sm italic">{{ play.author }}</p>
              </div>
              <div class="flex items-center gap-6 text-xs text-slate-500">
                <span class="flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">schedule</span>
                  {{ play.duration }}
                </span>
              </div>
              <RouterLink to="/contact"
                class="border border-[#C2B280]/50 hover:bg-[#C2B280] text-[#C2B280] hover:text-slate-900 text-sm px-4 py-2 rounded transition-colors flex items-center gap-1">
                <span class="material-symbols-outlined text-sm">event_available</span>
                Solicitar
              </RouterLink>
            </div>
          </article>
        </div>

        <!-- Empty state -->
        <div v-if="plays?.length === 0 && !loading" class="text-center py-20 text-slate-600">
          <span class="material-symbols-outlined text-5xl mb-3 block">search_off</span>
          <p>No hay obras en esta categoría.</p>
        </div>
      </ApiState>
    </section>

    <!-- CTA -->
    <section class="bg-[#C2B280]/10 border-t border-[#C2B280]/20 py-16">
      <div class="max-w-2xl mx-auto px-6 text-center">
        <h2 class="text-2xl font-bold mb-4">Quies contratanos?</h2>
        <p class="text-slate-600 mb-8">
          Llevamos les nueses obres a cases de cultura, teatros y eventos privaos. Contáctanos pa más información.
        </p>
        <RouterLink to="/contact"
          class="bg-[#C2B280] hover:bg-[#aa9668] text-slate-900 px-8 py-3 rounded font-medium transition-colors inline-block">
          Solicitar Información
        </RouterLink>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { playsApi, type Play } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'

const { data: plays, loading, error } = useApiRequest(() => playsApi.list(), {
  initialData: [] as Play[],
})
const gridView = ref<'grid' | 'list'>('grid')

</script>
