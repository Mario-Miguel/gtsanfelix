<template>
  <div class="bg-[#221010] text-white">
    <!-- Hero -->
    <section class="relative py-28 bg-[#1a0a0a] overflow-hidden">
      <div class="absolute inset-0 opacity-10 bg-gradient-to-br from-red-900 to-transparent" />
      <div class="relative z-10 text-center px-6 max-w-2xl mx-auto">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-4">Nuestra Historia</p>
        <h1 class="text-4xl md:text-5xl font-bold leading-tight mb-5">La Compañía</h1>
        <p class="text-gray-400 text-lg leading-relaxed">
          Desde unos humildes comienzos en un garaje hasta convertirnos en referente del teatro amateur en Madrid.
        </p>
      </div>
    </section>

    <!-- Stats -->
    <section class="border-y border-red-900/30 bg-[#1a0a0a] py-10">
      <div class="max-w-5xl mx-auto px-6 grid grid-cols-3 gap-6 text-center">
        <div v-for="stat in stats" :key="stat.label">
          <p class="text-4xl font-bold text-red-500">{{ stat.value }}</p>
          <p class="text-gray-400 text-sm mt-1">{{ stat.label }}</p>
        </div>
      </div>
    </section>

    <!-- Story -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="grid md:grid-cols-2 gap-12 items-center">
        <div>
          <h2 class="text-3xl font-bold mb-6">Nuestros Comienzos</h2>
          <p class="text-gray-400 leading-relaxed mb-4">
            G.T. San Félix de Valdesoto comenzó en 2008 cuando un grupo de amigos apasionados por el teatro decidió
            ensayar en un garaje del barrio. Nadie imaginaba entonces que aquella primera función —una adaptación de
            Lorca ante treinta espectadores— sería el inicio de una compañía que llegaría a miles de personas.
          </p>
          <p class="text-gray-400 leading-relaxed mb-4">
            A lo largo de 16 años hemos producido 24 obras, desde los clásicos grecolatinos hasta el teatro
            contemporáneo más vanguardista, siempre manteniendo nuestra esencia amateur: hacer teatro por amor al arte.
          </p>
          <p class="text-gray-400 leading-relaxed">
            Hoy somos una comunidad de más de 40 personas entre actores, técnicos, diseñadores y gestores culturales que
            comparten un mismo escenario y una misma pasión.
          </p>
        </div>
        <div class="rounded-lg overflow-hidden">
          <img
            src="https://lh3.googleusercontent.com/aida-public/AB6AXuBV8Bvg-dts3EHFHR4GSpTUe6bTBVRSzsNIi7oXOHqAd2WpJ-GI-Zh12NdOLNS6Y-zSmsERowyef6l_KL1gs35Up5vyjncscUE802A50nO23G5KdzhIyF9TZahF1tiPe8SP-iurM2ZAu2VXqiqkCNDLBIMYO9QbK265AgPELcgN3u0fVW8pYXokiz9ZyhuW_o2PbHHIcZvuFqdmjFHp2Tknz3BYsU7rDbU7Y191nDcWnje63oyBW6dUoMCwApF6JOsV2mEoBAvoLUU"
            alt="Compañía" class="w-full h-80 object-cover" />
        </div>
      </div>
    </section>

    <!-- Philosophy -->
    <section class="bg-[#1a0a0a] py-20">
      <div class="max-w-7xl mx-auto px-6">
        <div class="text-center mb-12">
          <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-3">Nuestra Filosofía</p>
          <h2 class="text-3xl font-bold">Los Pilares de G.T. San Félix de Valdesoto</h2>
        </div>
        <div class="grid md:grid-cols-3 gap-6">
          <div v-for="pillar in pillars" :key="pillar.title"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-7 text-center">
            <span class="material-symbols-outlined text-red-500 text-4xl mb-4 block">{{ pillar.icon }}</span>
            <h3 class="font-bold text-lg mb-3">{{ pillar.title }}</h3>
            <p class="text-gray-400 text-sm leading-relaxed">{{ pillar.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Team -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="text-center mb-12">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-3">Las Personas</p>
        <h2 class="text-3xl font-bold">Nuestro Equipo</h2>
      </div>
      <ApiState :loading="loading" :error="error">
        <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="member in team" :key="member.name"
            class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg overflow-hidden text-center">
            <div class="h-44 bg-[#3d1a1a] flex items-center justify-center">
              <span class="material-symbols-outlined text-red-700 text-6xl">person</span>
            </div>
            <div class="p-5">
              <h4 class="font-bold mb-1">{{ member.name }}</h4>
              <p class="text-red-400 text-xs uppercase tracking-wide mb-3">{{ member.role }}</p>
            </div>
          </div>
        </div>
      </ApiState>
    </section>


  </div>
</template>

<script setup lang="ts">

import { computed } from 'vue'
import { membersApi, type Member } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'

const { data: members, loading, error } = useApiRequest(() => membersApi.list(), {
  initialData: [] as Member[],
})

const team = computed(() => (members.value ?? []).filter((m) => m.active === true))

const stats = [
  { value: '2008', label: 'Año de fundación' },
  { value: '24', label: 'Producciones' },
  { value: '15.000', label: 'Espectadores' },
]

const pillars = [
  {
    icon: 'favorite',
    title: 'Puro Amateurismo',
    desc: 'Teatro creado desde la pasión artística, sin motivación comercial. Hacemos teatro porque lo amamos.',
  },
  {
    icon: 'groups',
    title: 'Comunidad Abierta',
    desc: 'Damos la bienvenida a intérpretes de todos los niveles de experiencia. El escenario es para todos.',
  },
  {
    icon: 'public',
    title: 'Impacto Social',
    desc: 'Usamos el teatro como herramienta para el diálogo comunitario y el cambio social positivo.',
  },
]


</script>
