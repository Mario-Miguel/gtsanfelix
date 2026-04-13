<template>
  <div class="bg-[#f8f6f6] text-slate-900">
    <!-- Hero -->
    <section class="relative py-28 bg-[#C2B280]/10 overflow-hidden">
      <div class="absolute inset-0 opacity-10 bg-gradient-to-br from-[#8e7a52] to-transparent" />
      <div class="relative z-10 text-center px-6 max-w-2xl mx-auto">
        <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-4">Nuestra Historia</p>
        <h1 class="text-4xl md:text-5xl font-bold leading-tight mb-5">La Compañía</h1>
        <p class="text-slate-600 text-lg leading-relaxed">
          Desde unos humildes comienzos en un garaje hasta convertirnos en referente del teatro amateur en Asturies.
        </p>
      </div>
    </section>

    <!-- Stats -->
    <section class="border-y border-[#C2B280]/20 bg-[#C2B280]/10 py-10">
      <div class="max-w-5xl mx-auto px-6 grid grid-cols-2 gap-6 text-center">
        <div v-for="stat in stats" :key="stat.label">
          <p class="text-4xl font-bold text-[#C2B280]">{{ stat.value }}</p>
          <p class="text-slate-600 text-sm mt-1">{{ stat.label }}</p>
        </div>
      </div>
    </section>

    <!-- Story -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="grid md:grid-cols-2 gap-12 items-center">
        <div>
          <h2 class="text-3xl font-bold mb-6">Quién Somos</h2>
          <p class="text-slate-600 leading-relaxed mb-4">
            Con más de venticincu años faciendo representaciones teatrales basaes prauticamente nel teatru costumista
            asturianu, intentamos reflexar en ca puesta n&apos;escena la vida cotidiana del pueblu asturianu.
          </p>
          <p class="text-slate-600 leading-relaxed mb-4">
            Asomando la identidá típica de les sos xentes, según los usos y costumes, ensín escaecer lóxicamente, la
            llingua propia nuesa, l&apos;asturianu; y too regáu cola nota risible, característica fundamental
            d&apos;esti xéneru.

          </p>
          <p class="text-slate-600 leading-relaxed">
            Nel planu personal, pertenecemos al grupu na actualidá once persones a parte de un elencu d&apos;ayudantes
            que mos faciliten siempre&apos;l trabayu. El grupu cuenta con decoráu propiu, fechu pa cada obra que
            tengamos en cartelu, según la llume y el soníu que necesite la mesma.

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
    <!-- <section class="bg-[#C2B280]/10 py-20">
      <div class="max-w-7xl mx-auto px-6">
        <div class="text-center mb-12">
          <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">Nuestra Filosofía</p>
          <h2 class="text-3xl font-bold">Los Pilares de G.T. San Félix de Valdesoto</h2>
        </div>
        <div class="grid md:grid-cols-3 gap-6">
          <div v-for="pillar in pillars" :key="pillar.title"
            class="bg-white border border-[#C2B280]/10 rounded-lg p-7 text-center">
            <span class="material-symbols-outlined text-[#C2B280] text-4xl mb-4 block">{{ pillar.icon }}</span>
            <h3 class="font-bold text-lg mb-3">{{ pillar.title }}</h3>
            <p class="text-slate-600 text-sm leading-relaxed">{{ pillar.desc }}</p>
          </div>
        </div>
      </div>
    </section> -->

    <!-- Active productions carousel -->
    <section class="bg-[#C2B280]/10 px-6 py-20">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12">
          <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">En Cartelera</p>
          <h2 class="text-3xl font-bold">Obres de Teatru Actuales</h2>
        </div>

        <ApiState :loading="playsLoading" :error="playsError">
          <PlaysCarousel :plays="plays ?? []" />
        </ApiState>
      </div>
    </section>

    <!-- Team -->
    <section class="max-w-7xl mx-auto px-6 py-20">
      <div class="text-center mb-12">
        <!-- <p class="text-[#C2B280] text-sm font-semibold uppercase tracking-widest mb-3">Las Personas</p> -->
        <h2 class="text-3xl font-bold">El nuesu Equipu</h2>
      </div>
      <ApiState :loading="loading" :error="error">
        <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="member in team" :key="member.name"
            class="bg-white border border-[#C2B280]/10 rounded-lg overflow-hidden text-center">
            <div class="h-44 bg-[#C2B280]/5 flex items-center justify-center">
              <span class="material-symbols-outlined text-[#8e7a52] text-6xl">person</span>
            </div>
            <div class="p-5">
              <h4 class="font-bold mb-1">{{ member.name }}</h4>
              <p class="text-[#C2B280] text-xs uppercase tracking-wide mb-3">{{ member.role }}</p>
            </div>
          </div>
        </div>
      </ApiState>
    </section>


  </div>
</template>

<script setup lang="ts">

import { membersApi, playsApi, type Member, type Play } from '../composables/useAdminApi'
import { useApiRequest } from '../composables/useApiRequest'
import ApiState from './ApiState.vue'
import PlaysCarousel from './PlaysCarousel.vue'

const { data: team, loading, error } = useApiRequest(() => membersApi.list(), {
  initialData: [] as Member[],
})

const { data: plays, loading: playsLoading, error: playsError } = useApiRequest(() => playsApi.list(), {
  initialData: [] as Play[],
})



const stats = [
  { value: '2000', label: 'Añu de fundación' },
  { value: '24', label: 'Obres de Teatru' },
]

// const pillars = [
//   {
//     icon: 'favorite',
//     title: 'Puro Amateurismo',
//     desc: 'Teatro creado desde la pasión artística, sin motivación comercial. Hacemos teatro porque lo amamos.',
//   },
//   {
//     icon: 'groups',
//     title: 'Comunidad Abierta',
//     desc: 'Damos la bienvenida a intérpretes de todos los niveles de experiencia. El escenario es para todos.',
//   },
//   {
//     icon: 'public',
//     title: 'Impacto Social',
//     desc: 'Usamos el teatro como herramienta para el diálogo comunitario y el cambio social positivo.',
//   },
// ]


</script>
