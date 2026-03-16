<template>
  <div class="bg-[#221010] text-white">
    <!-- Hero -->
    <section class="bg-[#1a0a0a] py-16 border-b border-red-900/30">
      <div class="max-w-7xl mx-auto px-6">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-3">Programación</p>
        <h1 class="text-4xl font-bold mb-3">Cartelera de Funciones</h1>
        <p class="text-gray-400 max-w-xl">
          Toda la programación de G.T. San Félix de Valdesoto. Consulta fechas, sedes y reserva tu entrada.
        </p>
      </div>
    </section>

    <!-- Filters bar -->
    <section class="bg-[#1a0a0a] border-b border-red-900/20 sticky top-[73px] z-40">
      <div class="max-w-7xl mx-auto px-6 py-3 flex items-center gap-4 flex-wrap">
        <div class="flex items-center gap-2 ml-auto text-sm">
          <span class="material-symbols-outlined text-red-400 text-sm">filter_list</span>
          <select
            v-model="filterVenue"
            class="bg-[#2d1515] border border-red-900/30 rounded px-3 py-2 text-sm text-gray-300 focus:outline-none focus:border-red-600"
          >
            <option value="">Todas las sedes</option>
            <option v-for="v in uniqueVenues" :key="v">{{ v }}</option>
          </select>
        </div>
      </div>
    </section>

    <ApiState :loading="loading" :error="error" py="py-24">
      <div class="max-w-7xl mx-auto px-6 py-12">

        <!-- Today's shows -->
        <div v-if="todayPerfs.length > 0" class="mb-14">
          <h2 class="text-xl font-bold mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-red-500">today</span>
            Hoy — {{ todayLabel }}
          </h2>
          <div class="grid md:grid-cols-2 gap-6">
            <article
              v-for="perf in todayPerfs"
              :key="perf.id"
              class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-6 hover:border-red-700/40 transition-colors"
            >
              <div class="flex items-center justify-between mb-4">
                <span class="bg-red-700/30 text-red-400 text-xs px-3 py-1 rounded-full">HOY</span>
                <span v-if="perf.sold > 75" class="text-xs text-yellow-500 flex items-center gap-1">
                  <span class="material-symbols-outlined text-xs">local_fire_department</span>
                  Pocas entradas
                </span>
              </div>
              <h3 class="font-bold text-xl mb-3">{{ perf.playTitle }}</h3>
              <div class="flex flex-wrap gap-4 text-sm text-gray-400 mb-5">
                <span class="flex items-center gap-1.5">
                  <span class="material-symbols-outlined text-sm">schedule</span>
                  {{ perf.time }} hrs
                </span>
                <span class="flex items-center gap-1.5">
                  <span class="material-symbols-outlined text-sm">location_on</span>
                  {{ perf.venue }}
                </span>
                <span v-if="perf.price" class="flex items-center gap-1.5">
                  <span class="material-symbols-outlined text-sm">payments</span>
                  {{ perf.price }}
                </span>
              </div>
              <div class="flex items-center justify-between">
                <div class="flex-1 mr-4">
                  <div class="flex justify-between text-xs text-gray-600 mb-1">
                    <span>Ocupación</span>
                    <span>{{ perf.sold }}%</span>
                  </div>
                  <div class="bg-red-900/20 rounded-full h-1.5">
                    <div class="bg-red-600 h-1.5 rounded-full transition-all" :style="{ width: perf.sold + '%' }" />
                  </div>
                </div>
                <RouterLink to="/contact" class="bg-red-700 hover:bg-red-600 text-white text-sm px-4 py-2 rounded transition-colors flex-shrink-0">
                  Reservar
                </RouterLink>
              </div>
            </article>
          </div>
        </div>

        <!-- Upcoming performances -->
        <div>
          <h2 class="text-xl font-bold mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-red-500">event_upcoming</span>
            {{ todayPerfs.length > 0 ? 'Próximos Días' : 'Todas las Funciones' }}
          </h2>

          <div v-if="upcomingPerfs.length > 0" class="flex flex-col gap-4">
            <article
              v-for="perf in upcomingPerfs"
              :key="perf.id"
              class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg px-6 py-4 flex items-center gap-6 hover:border-red-700/40 transition-colors"
            >
              <div class="text-center bg-red-700/20 text-red-400 rounded px-4 py-2 min-w-[60px] flex-shrink-0">
                <p class="text-xs font-medium">{{ monthLabel(perf.date) }}</p>
                <p class="text-2xl font-bold leading-none">{{ dayLabel(perf.date) }}</p>
              </div>
              <div class="flex-1 min-w-0">
                <h4 class="font-semibold truncate">{{ perf.playTitle }}</h4>
                <p class="text-gray-500 text-sm flex items-center gap-3 mt-1 flex-wrap">
                  <span class="flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">location_on</span>
                    {{ perf.venue }}
                  </span>
                  <span class="flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">schedule</span>
                    {{ perf.time }} hrs
                  </span>
                  <span v-if="perf.price" class="flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">payments</span>
                    {{ perf.price }}
                  </span>
                </p>
              </div>
              <!-- Sold bar (compact) -->
              <div class="hidden sm:flex items-center gap-2 min-w-[80px]">
                <div class="w-16 bg-red-900/20 rounded-full h-1.5">
                  <div class="bg-red-600 h-1.5 rounded-full" :style="{ width: perf.sold + '%' }" />
                </div>
                <span class="text-xs text-gray-600">{{ perf.sold }}%</span>
              </div>
              <RouterLink
                to="/contact"
                class="text-red-400 hover:text-red-300 text-sm font-medium transition-colors flex items-center gap-1 flex-shrink-0"
              >
                Reservar <span class="material-symbols-outlined text-sm">chevron_right</span>
              </RouterLink>
            </article>
          </div>

          <div v-else class="text-center py-16 text-gray-600">
            <span class="material-symbols-outlined text-5xl mb-3 block">event_busy</span>
            <p>No hay funciones programadas próximamente.</p>
          </div>
        </div>
      </div>

      <!-- Venues -->
      <section v-if="uniqueVenues.length > 0" class="bg-[#1a0a0a] border-t border-red-900/30 py-12">
        <div class="max-w-7xl mx-auto px-6">
          <h2 class="text-xl font-bold mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-red-500">theater_comedy</span>
            Sedes
          </h2>
          <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div
              v-for="venue in uniqueVenues"
              :key="venue"
              class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-5 flex items-center gap-4"
            >
              <span class="material-symbols-outlined text-red-500 text-3xl">theater_comedy</span>
              <div>
                <h4 class="font-semibold">{{ venue }}</h4>
                <p class="text-gray-500 text-sm">{{ venueCounts[venue] }} función{{ venueCounts[venue] !== 1 ? 'es' : '' }} programada{{ venueCounts[venue] !== 1 ? 's' : '' }}</p>
              </div>
            </div>
          </div>
        </div>
      </section>
    </ApiState>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { RouterLink } from 'vue-router'
import { performancesApi, type Performance } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'

const { data: performances, loading, error } = useApiRequest(() => performancesApi.list(), {
  initialData: [] as Performance[],
})
const filterVenue = ref('')

// ── Date helpers ──────────────────────────────────────────────────────────
const MONTHS = ['ENE', 'FEB', 'MAR', 'ABR', 'MAY', 'JUN', 'JUL', 'AGO', 'SEP', 'OCT', 'NOV', 'DIC']

function parseDate(dateStr: string) {
  // Append T00:00:00 to avoid timezone shifts
  return new Date(dateStr + 'T00:00:00')
}

function monthLabel(dateStr: string) {
  return MONTHS[parseDate(dateStr).getMonth()]
}

function dayLabel(dateStr: string) {
  return parseDate(dateStr).getDate()
}

const todayLabel = computed(() => {
  return new Date().toLocaleDateString('es-ES', { weekday: 'long', day: 'numeric', month: 'long' })
})

function isToday(dateStr: string) {
  const today = new Date()
  const d = parseDate(dateStr)
  return d.getFullYear() === today.getFullYear() &&
    d.getMonth() === today.getMonth() &&
    d.getDate() === today.getDate()
}

function isFuture(dateStr: string) {
  return parseDate(dateStr) > new Date()
}

// ── Derived data ──────────────────────────────────────────────────────────
const filtered = computed(() => {
  let list = performances.value ?? []
  if (filterVenue.value) list = list.filter((p) => p.venue === filterVenue.value)
  return [...list].sort((a, b) => parseDate(a.date).getTime() - parseDate(b.date).getTime())
})

const todayPerfs = computed(() => filtered.value.filter((p) => isToday(p.date)))
const upcomingPerfs = computed(() => filtered.value.filter((p) => isFuture(p.date)))

const uniqueVenues = computed(() => [...new Set((performances.value ?? []).map((p) => p.venue).filter(Boolean))])

const venueCounts = computed(() => {
  const counts: Record<string, number> = {}
  ;(performances.value ?? []).forEach((p) => {
    counts[p.venue] = (counts[p.venue] ?? 0) + 1
  })
  return counts
})
</script>
