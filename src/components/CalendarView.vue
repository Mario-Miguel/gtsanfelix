<template>
  <div class="bg-[#f8f6f6] text-slate-900">
    <!-- Hero -->
    <section class="bg-[#C2B280]/10 py-16 border-b border-[#C2B280]/20">
      <div class="max-w-7xl mx-auto px-6">
        <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">Programación</p>
        <h1 class="text-4xl font-bold mb-3">Cartelera de Funciones</h1>
        <p class="text-slate-600 max-w-xl">
          Toda la programación de G.T. San Félix de Valdesoto. Consulta fechas, sedes y reserva tu entrada.
        </p>
      </div>
    </section>

    <!-- Filters bar -->
    <section class="bg-[#C2B280]/10 border-b border-[#C2B280]/10 sticky top-[73px] z-40">
      <div class="max-w-7xl mx-auto px-6 py-3 flex items-center gap-4 flex-wrap">
        <div class="flex items-center gap-2 ml-auto text-sm">
          <span class="material-symbols-outlined text-[#C2B280] text-sm">filter_list</span>
          <select v-model="filterVenue"
            class="bg-white border border-[#C2B280]/20 rounded px-3 py-2 text-sm text-slate-700 focus:outline-none focus:border-[#C2B280]">
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
            <span class="material-symbols-outlined text-[#C2B280]">today</span>
            Hoy — {{ todayLabel }}
          </h2>
          <div class="grid md:grid-cols-2 gap-6">
            <article v-for="perf in todayPerfs" :key="perf.id"
              class="bg-white border border-[#C2B280]/10 rounded-lg p-6 hover:border-[#C2B280]/40 transition-colors">
              <div class="flex items-center justify-between mb-4">
                <span class="bg-[#C2B280]/20 text-[#8e7a52] text-xs px-3 py-1 rounded-full">HOY</span>
              </div>
              <h3 class="font-bold text-xl mb-3">{{ perf.playTitle }}</h3>
              <div class="flex flex-wrap gap-4 text-sm text-slate-600 mb-5">
                <span class="flex items-center gap-1.5">
                  <span class="material-symbols-outlined text-sm">schedule</span>
                  {{ perf.time }} hrs
                </span>
                <span class="flex items-center gap-1.5">
                  <span class="material-symbols-outlined text-sm">location_on</span>
                  {{ perf.venue }}
                </span>
              </div>
            </article>
          </div>
        </div>

        <!-- Upcoming performances -->
        <div>
          <h2 class="text-xl font-bold mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-[#C2B280]">event_upcoming</span>
            {{ todayPerfs.length > 0 ? 'Próximos Días' : 'Todas las Funciones' }}
          </h2>

          <div v-if="upcomingPerfs.length > 0" class="flex flex-col gap-4">
            <article v-for="perf in upcomingPerfs" :key="perf.id"
              class="bg-white border border-[#C2B280]/10 rounded-lg px-6 py-4 flex items-center gap-6 hover:border-[#C2B280]/40 transition-colors">
              <div class="text-center bg-[#C2B280]/20 text-[#C2B280] rounded px-4 py-2 min-w-[60px] flex-shrink-0">
                <p class="text-xs font-medium">{{ monthLabel(perf.date) }}</p>
                <p class="text-2xl font-bold leading-none">{{ dayLabel(perf.date) }}</p>
              </div>
              <div class="flex-1 min-w-0">
                <h4 class="font-semibold truncate">{{ perf.playTitle }}</h4>
                <p class="text-slate-500 text-sm flex items-center gap-3 mt-1 flex-wrap">
                  <span class="flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">location_on</span>
                    {{ perf.venue }}
                  </span>
                  <span class="flex items-center gap-1">
                    <span class="material-symbols-outlined text-xs">schedule</span>
                    {{ perf.time }} hrs
                  </span>
                </p>
              </div>
            </article>
          </div>

          <div v-else class="text-center py-16 text-slate-600">
            <span class="material-symbols-outlined text-5xl mb-3 block">event_busy</span>
            <p>No hay funciones programadas próximamente.</p>
          </div>
        </div>
      </div>

      <!-- Venues -->
      <section v-if="uniqueVenues.length > 0" class="bg-[#C2B280]/10 border-t border-[#C2B280]/20 py-12">
        <div class="max-w-7xl mx-auto px-6">
          <h2 class="text-xl font-bold mb-6 flex items-center gap-2">
            <span class="material-symbols-outlined text-[#C2B280]">theater_comedy</span>
            Sedes
          </h2>
          <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div v-for="venue in uniqueVenues" :key="venue"
              class="bg-white border border-[#C2B280]/10 rounded-lg p-5 flex items-center gap-4">
              <span class="material-symbols-outlined text-[#C2B280] text-3xl">theater_comedy</span>
              <div>
                <h4 class="font-semibold">{{ venue }}</h4>
                <p class="text-slate-500 text-sm">{{ venueCounts[venue] }} función{{ venueCounts[venue] !== 1 ? 'es' : ''
                  }} programada{{ venueCounts[venue] !== 1 ? 's' : '' }}</p>
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
    ; (performances.value ?? []).forEach((p) => {
      counts[p.venue] = (counts[p.venue] ?? 0) + 1
    })
  return counts
})
</script>
