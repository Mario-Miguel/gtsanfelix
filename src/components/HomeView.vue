<template>
  <div class="bg-[#221010] text-white">
    <!-- Hero -->
    <section class="relative h-[90vh] min-h-[600px] flex items-center justify-center overflow-hidden">
      <img
        src="https://lh3.googleusercontent.com/aida-public/AB6AXuAW9aqP2FW6v3L_643dwCFFyeHF1yCTuoGE8gfOqYBmZVs5BVKP37Aprvx_SURthZHCPOZaRKMKaOtAnXWyV8H9KGkNtkWdcRMQhVqCSuc0ar9-R9i8Trc3GEWEblN5ogeXcGPbvpsNMrOoIgYy3jk9NeDkxx-lHRNunGBq18hIFr-v-7ZtaFZrYzjhdKc0B-w7LMuTvcacgNjRo06OQaezXpG_Ebrp1fbg3ethE55CQCWAUFtS1ICrfv71z_l9slcvQL8Jf_6V9fY"
        alt="Teatro" class="absolute inset-0 w-full h-full object-cover" />
      <div class="absolute inset-0 bg-gradient-to-b from-[#221010]/70 via-[#221010]/50 to-[#221010]" />
      <div class="relative z-10 text-center px-6 max-w-3xl mx-auto">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-4">Valdesoto · Desde 2008</p>
        <h1 class="text-5xl md:text-6xl font-bold leading-tight mb-6">
          Donde la pasión<br />se encuentra con el escenario
        </h1>
        <p class="text-gray-300 text-lg mb-8 max-w-xl mx-auto">
          Vive la magia del teatro amateur y únete a nuestra comunidad creativa.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <RouterLink to="/calendar"
            class="bg-red-700 hover:bg-red-600 px-7 py-3 rounded font-medium transition-colors">
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
    <section class="bg-[#1a0a0a] border-y border-red-900/30 py-10">
      <div class="max-w-5xl mx-auto px-6 grid grid-cols-2 gap-6 text-center">
        <div>
          <p class="text-4xl font-bold text-red-500">16+</p>
          <p class="text-gray-400 text-sm mt-1">Años en escena</p>
        </div>
        <div>
          <p class="text-4xl font-bold text-red-500">24</p>
          <p class="text-gray-400 text-sm mt-1">Producciones</p>
        </div>
      </div>
    </section>

    <!-- Producciones Actuales -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="flex items-center gap-3 mb-10">
        <span class="material-symbols-outlined text-red-500">star</span>
        <h2 class="text-2xl font-bold">Producciones Actuales</h2>
      </div>
      <ApiState :loading="playsLoading" :error="playsError">
        <PlaysCarousel :plays="plays ?? []" />
      </ApiState>
    </section>

    <!-- Próximas Funciones -->
    <section class="bg-[#1a0a0a] py-20">
      <div class="max-w-7xl mx-auto px-6">
        <div class="flex items-center justify-between mb-10">
          <div class="flex items-center gap-3">
            <span class="material-symbols-outlined text-red-500">event</span>
            <h2 class="text-2xl font-bold">Próximas Funciones</h2>
          </div>
          <RouterLink to="/calendar"
            class="text-red-400 hover:text-red-300 text-sm transition-colors flex items-center gap-1">
            Ver todas <span class="material-symbols-outlined text-sm">chevron_right</span>
          </RouterLink>
        </div>
        <div class="grid md:grid-cols-3 gap-6">
          <div v-for="show in upcomingShows" :key="show.title"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-6 hover:border-red-700/40 transition-colors">
            <div class="flex items-center gap-3 mb-4">
              <div class="bg-red-700/20 text-red-400 rounded px-3 py-2 text-center min-w-[52px]">
                <p class="text-xs font-medium">{{ show.month }}</p>
                <p class="text-2xl font-bold leading-none">{{ show.day }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-500">{{ show.weekday }}</p>
                <p class="text-sm font-medium">{{ show.time }}</p>
              </div>
            </div>
            <h4 class="font-semibold mb-1">{{ show.title }}</h4>
            <p class="text-gray-500 text-sm mb-1 italic">{{ show.author }}</p>
            <p class="text-gray-600 text-xs flex items-center gap-1 mt-2">
              <span class="material-symbols-outlined text-xs">location_on</span>
              {{ show.venue }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- About teaser -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="grid md:grid-cols-2 gap-12 items-center">
        <div>
          <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-3">Quiénes somos</p>
          <h2 class="text-3xl font-bold mb-5">Fundada en 2008 con una sola misión</h2>
          <p class="text-gray-400 leading-relaxed mb-6">
            G.T. San Félix de Valdesoto nació para ofrecer un espacio profesional donde actores, técnicos y escritores
            aficionados puedan explorar su arte. Más de 16 años creando magia sobre el escenario en Asturies.
          </p>
          <RouterLink to="/about"
            class="text-red-400 hover:text-red-300 font-medium flex items-center gap-1 transition-colors">
            Conoce nuestra historia
            <span class="material-symbols-outlined text-base">arrow_forward</span>
          </RouterLink>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div v-for="pillar in pillars" :key="pillar.title"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-5">
            <span class="material-symbols-outlined text-red-500 mb-3 block">{{ pillar.icon }}</span>
            <h4 class="font-semibold text-sm mb-1">{{ pillar.title }}</h4>
            <p class="text-gray-500 text-xs">{{ pillar.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-[#0f0505] border-t border-red-900/30 py-12">
      <div class="max-w-7xl mx-auto px-6">
        <div class="grid md:grid-cols-4 gap-8 mb-10">
          <div class="md:col-span-2">
            <div class="flex items-center gap-2 mb-4">
              <span class="material-symbols-outlined text-red-600">theater_comedy</span>
              <span class="font-bold text-lg">G.T. San Félix de Valdesoto</span>
            </div>
            <p class="text-gray-500 text-sm leading-relaxed max-w-xs">
              Llevamos la magia del teatro a cada rincón. Compañía dedicada a la excelencia artística y la divulgación
              cultural en Asturies.
            </p>
          </div>
          <div>
            <h4 class="font-semibold mb-4 text-xs uppercase tracking-wider text-gray-400">Navegación</h4>
            <ul class="flex flex-col gap-2 text-sm text-gray-500">
              <li>
                <RouterLink to="/" class="hover:text-white transition-colors">Inicio</RouterLink>
              </li>
              <li>
                <RouterLink to="/about" class="hover:text-white transition-colors">Compañía</RouterLink>
              </li>
              <li>
                <RouterLink to="/repertoire" class="hover:text-white transition-colors">Repertorio</RouterLink>
              </li>
              <li>
                <RouterLink to="/calendar" class="hover:text-white transition-colors">Cartelera</RouterLink>
              </li>
            </ul>
          </div>
          <div>
            <h4 class="font-semibold mb-4 text-xs uppercase tracking-wider text-gray-400">Contacto</h4>
            <ul class="flex flex-col gap-2 text-sm text-gray-500">
              <li class="flex items-center gap-2">
                <span class="material-symbols-outlined text-sm">mail</span> gtvaldesoto@gmail.com
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
          class="border-t border-red-900/20 pt-6 flex flex-col md:flex-row items-center justify-between gap-4 text-xs text-gray-600">
          <p>© 2024 G.T. San Félix de Valdesoto. Todos los derechos reservados.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { playsApi, type Play } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'
import PlaysCarousel from './PlaysCarousel.vue'

const { data: plays, loading: playsLoading, error: playsError } = useApiRequest(() => playsApi.list(), {
  initialData: [] as Play[],
})

const upcomingShows = [
  {
    month: 'OCT',
    day: '15',
    weekday: 'Domingo',
    time: '19:00 hrs',
    title: 'Esperando a Godot',
    author: 'Samuel Beckett',
    venue: 'Teatro Principal',
  },
  {
    month: 'NOV',
    day: '2',
    weekday: 'Sábado',
    time: '20:30 hrs',
    title: 'Bodas de Sangre',
    author: 'Federico García Lorca',
    venue: 'Teatro La Latina',
  },
  {
    month: 'DIC',
    day: '10',
    weekday: 'Miércoles',
    time: '20:00 hrs',
    title: 'Midsummer Night',
    author: 'William Shakespeare',
    venue: 'Teatro del Barrio',
  },
]

const pillars = [
  { icon: 'favorite', title: 'Puro Amateurismo', desc: 'Arte creado desde la pasión, no la comercialidad.' },
  { icon: 'groups', title: 'Comunidad Abierta', desc: 'Bienvenidos actores de todos los niveles.' },
  { icon: 'public', title: 'Impacto Social', desc: 'El teatro como herramienta de diálogo.' },
  { icon: 'workspace_premium', title: 'Excelencia', desc: 'Producción de calidad profesional.' },
]
</script>
