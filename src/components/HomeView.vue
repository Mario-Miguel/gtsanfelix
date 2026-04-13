<template>
  <div class="bg-[#f8f6f6] text-slate-900">
    <!-- Hero -->
    <section class="relative h-[90vh] min-h-[600px] flex items-center justify-center overflow-hidden">
      <img
        src="https://lh3.googleusercontent.com/aida-public/AB6AXuAW9aqP2FW6v3L_643dwCFFyeHF1yCTuoGE8gfOqYBmZVs5BVKP37Aprvx_SURthZHCPOZaRKMKaOtAnXWyV8H9KGkNtkWdcRMQhVqCSuc0ar9-R9i8Trc3GEWEblN5ogeXcGPbvpsNMrOoIgYy3jk9NeDkxx-lHRNunGBq18hIFr-v-7ZtaFZrYzjhdKc0B-w7LMuTvcacgNjRo06OQaezXpG_Ebrp1fbg3ethE55CQCWAUFtS1ICrfv71z_l9slcvQL8Jf_6V9fY"
        alt="Teatro" class="absolute inset-0 w-full h-full object-cover" />
      <div class="absolute inset-0 bg-gradient-to-b from-[#f8f6f6]/70 via-[#f8f6f6]/50 to-[#f8f6f6]" />
      <div class="relative z-10 text-center px-6 max-w-3xl mx-auto">
        <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-4">Valdesoto · Desde 2008</p>
        <h1 class="text-5xl md:text-6xl font-bold leading-tight mb-6">
          Donde la pasión<br />se encuentra con el escenario
        </h1>
        <p class="text-slate-700 text-lg mb-8 max-w-xl mx-auto">
          Vive la magia del teatro amateur y únete a nuestra comunidad creativa.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <RouterLink to="/calendar"
            class="bg-[#C2B280] hover:bg-[#aa9668] text-slate-900 px-7 py-3 rounded font-medium transition-colors">
            Ver Cartelera
          </RouterLink>
          <RouterLink to="/about"
            class="border border-white/30 hover:border-white px-7 py-3 rounded font-medium transition-colors">
            Conoce la Compañía
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Stats bar -->
    <section class="bg-[#C2B280]/10 border-y border-[#C2B280]/20 py-10">
      <div class="max-w-5xl mx-auto px-6 grid grid-cols-2 gap-6 text-center">
        <div>
          <p class="text-4xl font-bold text-[#C2B280]">{{ activeYears }}+</p>
          <p class="text-slate-600 text-sm mt-1">Años en escena</p>
        </div>
        <ApiState :loading="playsLoading" :error="playsError">
          <div>
            <p class="text-4xl font-bold text-[#C2B280]">{{ plays?.length }}</p>
            <p class="text-slate-600 text-sm mt-1">Obras</p>
          </div>
        </ApiState>
      </div>
    </section>

    <!-- Obres de Teatru Actuales -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="flex items-center gap-3 mb-10">
        <span class="material-symbols-outlined text-[#C2B280]">star</span>
        <h2 class="text-2xl font-bold">Obres de Teatru Actuales</h2>
      </div>
      <ApiState :loading="playsLoading" :error="playsError">
        <PlaysCarousel :plays="plays ?? []" />
      </ApiState>
    </section>

    <!-- Próximas Funciones -->
    <section class="bg-[#C2B280]/10 py-20">
      <div class="max-w-7xl mx-auto px-6">
        <div class="flex items-center justify-between mb-10">
          <div class="flex items-center gap-3">
            <span class="material-symbols-outlined text-[#C2B280]">event</span>
            <h2 class="text-2xl font-bold">Próximes Funciones</h2>
          </div>
          <RouterLink to="/calendar"
            class="text-[#C2B280] hover:text-[#aa9668] text-sm transition-colors flex items-center gap-1">
            Ver todas <span class="material-symbols-outlined text-sm">chevron_right</span>
          </RouterLink>
        </div>
        <ApiState :loading="performancesLoading" :error="performancesError">
          <div v-if="upcomingShows?.length === 0" class="text-center py-12 text-slate-600">
            <span class="material-symbols-outlined text-5xl mb-3 block">theater_comedy</span>
            <p>No hay funciones a la vista.</p>
          </div>
          <div v-else class="grid md:grid-cols-3 gap-6">
            <div v-for="show in upcomingShows" :key="show.playId"
              class="bg-white border border-[#C2B280]/10 rounded-lg p-6 hover:border-[#C2B280]/40 transition-colors">
              <div class="flex items-center gap-3 mb-4">
                <div class="bg-[#C2B280]/20 text-[#C2B280] rounded px-3 py-2 text-center min-w-[52px]">
                  <p class="text-xs font-medium">{{ monthLabel(show.date) }}</p>
                  <p class="text-2xl font-bold leading-none">{{ dayLabel(show.date) }}</p>
                </div>
                <div>
                  <p class="text-xs text-slate-500">{{ weekdayLabel(show.date) }}</p>
                  <p class="text-sm font-medium">{{ show.time }}</p>
                </div>
              </div>
              <h4 class="font-semibold mb-1">{{ show.playTitle }}</h4>
              <p class="text-slate-600 text-xs flex items-center gap-1 mt-2">
                <span class="material-symbols-outlined text-xs">location_on</span>
                {{ show.venue }}
              </p>
            </div>
          </div>
        </ApiState>
      </div>
    </section>

    <!-- About teaser -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="grid md:grid-cols-2 gap-12 items-center">
        <div>
          <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">Quién somos</p>
          <h2 class="text-3xl font-bold mb-5">Dende'l 2000</h2>
          <p class="text-slate-600 leading-relaxed mb-6">
            Con más de venticincu años faciendo representaciones teatrales basaes prauticamente nel teatru costumista
            asturianu, intentamos reflexar en ca puesta n&apos;escena la vida cotidiana del pueblu asturianu.
          </p>
          <RouterLink to="/about"
            class="text-[#C2B280] hover:text-[#aa9668] font-medium flex items-center gap-1 transition-colors">
            Conoz la nuesa hestoria
            <span class="material-symbols-outlined text-base">arrow_forward</span>
          </RouterLink>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div v-for="pillar in pillars" :key="pillar.title" class="bg-white border border-[#C2B280]/10 rounded-lg p-5">
            <span class="material-symbols-outlined text-[#C2B280] mb-3 block">{{ pillar.icon }}</span>
            <h4 class="font-semibold text-sm mb-1">{{ pillar.title }}</h4>
            <p class="text-slate-500 text-xs">{{ pillar.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-white border-t border-[#C2B280]/20 py-12">
      <div class="max-w-7xl mx-auto px-6">
        <div class="grid md:grid-cols-4 gap-8 mb-10">
          <div class="md:col-span-2">
            <div class="flex items-center gap-2 mb-4">
              <span class="material-symbols-outlined text-[#C2B280]">theater_comedy</span>
              <span class="font-bold text-lg">G.T. San Félix de Valdesoto</span>
            </div>

          </div>
          <div>
            <h4 class="font-semibold mb-4 text-xs uppercase tracking-wider text-slate-600">Navegación</h4>
            <ul class="flex flex-col gap-2 text-sm text-slate-500">
              <li>
                <RouterLink to="/" class="hover:text-slate-900 transition-colors">Iniciu</RouterLink>
              </li>
              <li>
                <RouterLink to="/about" class="hover:text-slate-900 transition-colors">Compañía</RouterLink>
              </li>
              <li>
                <RouterLink to="/repertoire" class="hover:text-slate-900 transition-colors">Repertoriu</RouterLink>
              </li>
              <li>
                <RouterLink to="/calendar" class="hover:text-slate-900 transition-colors">Cartelera</RouterLink>
              </li>
            </ul>
          </div>
          <div>
            <h4 class="font-semibold mb-4 text-xs uppercase tracking-wider text-slate-600">Contactu</h4>
            <ul class="flex flex-col gap-2 text-sm text-slate-500">
              <li class="flex items-center gap-2">
                <span class="material-symbols-outlined text-sm">mail</span> gteatrusanfelix@hotmail.com
              </li>
              <li class="flex items-center gap-2">
                <span class="material-symbols-outlined text-sm">phone</span> +34 606 58 54 12
              </li>
              <li class="flex items-center gap-2">
                <span class="material-symbols-outlined text-sm">location_on</span> Valdesoto, Asturies
              </li>
            </ul>
          </div>
        </div>
        <div
          class="border-t border-[#C2B280]/10 pt-6 flex flex-col md:flex-row items-center justify-between gap-4 text-xs text-slate-600">
          <p>© 2026 G.T. San Félix de Valdesoto. Todos los derechos reservados.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { playsApi, type Play, performancesApi, type Performance } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'
import PlaysCarousel from './PlaysCarousel.vue'
import { computed } from 'vue'


const MONTHS = ['ENE', 'FEB', 'MAR', 'ABR', 'MAY', 'JUN', 'JUL', 'AGO', 'SEP', 'OCT', 'NOV', 'DIC']
function monthLabel(dateStr: string) {
  const d = new Date(dateStr)
  return MONTHS[d.getMonth()] ?? ''
}
function dayLabel(dateStr: string) {
  return new Date(dateStr).getDate()
}
function weekdayLabel(dateStr: string) {
  return new Date(dateStr).toLocaleDateString('es-ES', { weekday: 'long' })
}

const activeYears = new Date().getFullYear() - new Date('2000-01-01').getFullYear()

const { data: plays, loading: playsLoading, error: playsError } = useApiRequest(() => playsApi.list(), {
  initialData: [] as Play[],
})

const { data: performances, loading: performancesLoading, error: performancesError } = useApiRequest(() => performancesApi.list(), {
  initialData: [] as Performance[],
})

const upcomingShows = computed(() => performances.value?.filter((p) => new Date(p.date) > new Date()).slice(0, 3))


const pillars = [
  { icon: 'favorite', title: 'Puro Amateurismo', desc: 'Arte creado desde la pasión, no la comercialidad.' },
  { icon: 'groups', title: 'Comunidad Abierta', desc: 'Bienvenidos actores de todos los niveles.' },
  { icon: 'public', title: 'Impacto Social', desc: 'El teatro como herramienta de diálogo.' },
  { icon: 'workspace_premium', title: 'Excelencia', desc: 'Producción de calidad profesional.' },
]
</script>
