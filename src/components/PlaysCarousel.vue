<template>
  <div v-if="visiblePlays.length === 0" class="text-center py-12 text-slate-600">
    <span class="material-symbols-outlined text-5xl mb-3 block">theater_comedy</span>
    <p>No hay producciones en este momento.</p>
  </div>

  <div v-else class="relative">
    <div class="overflow-hidden rounded-lg">
      <Transition :name="transitionName" mode="out-in">
        <div v-if="currentPlay" :key="currentIndex"
          class="grid md:grid-cols-2 gap-0 bg-white border border-[#C2B280]/10 rounded-lg overflow-hidden">
          <!-- Image -->
          <div class="h-72 md:h-auto bg-white flex items-center justify-center relative overflow-hidden">
            <img v-if="currentPlay.imageUrl" :src="currentPlay.imageUrl" :alt="currentPlay.title"
              class="w-full h-full object-cover" />
            <span v-else class="material-symbols-outlined text-[#8e7a52] text-8xl">theater_comedy</span>
            <div class="absolute inset-0 bg-gradient-to-r from-transparent to-[#f8f6f6] hidden md:block" />
          </div>

          <!-- Info -->
          <div class="p-10 flex flex-col justify-center">
            <div class="flex items-center gap-2 mb-4">
              <span class="text-xs bg-[#C2B280]/20 text-[#8e7a52] px-3 py-1 rounded-full">
                {{ currentPlay.genre }}
              </span>
              <span v-if="currentPlay.active"
                class="text-xs bg-green-700/20 px-3 py-1 rounded-full flex items-center gap-1">
                <span class="w-1.5 h-1.5 rounded-full bg-green-400 inline-block" />
                En cartelera
              </span>
              <span v-else class="text-xs bg-gray-800/20 px-3 py-1 rounded-full flex items-center gap-1">
                <span class="w-1.5 h-1.5 rounded-full bg-gray-600 inline-block" />
                Archivo
              </span>
            </div>
            <h3 class="text-3xl font-bold mb-2">{{ currentPlay.title }}</h3>
            <p class="text-slate-600 italic mb-6">{{ currentPlay.author }}</p>
            <p v-if="currentPlay.summary" class="text-slate-600 leading-relaxed mb-6">
              {{ currentPlay.summary }}
            </p>
            <div class="flex items-center gap-4 text-sm text-slate-600 mb-8">
              <span class="flex items-center gap-1.5">
                <span class="material-symbols-outlined text-sm">schedule</span>
                {{ currentPlay.duration }}
              </span>
            </div>
            <RouterLink v-if="currentPlay.active" to="/contact"
              class="self-start border border-[#C2B280]/50 hover:bg-[#C2B280] text-[#C2B280] hover:text-slate-900 px-6 py-2.5 rounded transition-colors flex items-center gap-2 text-sm">
              <span class="material-symbols-outlined text-sm">event_available</span>
              Solicitar Función
            </RouterLink>
          </div>
        </div>
      </Transition>
    </div>

    <!-- Controls -->
    <div v-if="visiblePlays.length > 1" class="flex items-center justify-center gap-4 mt-6">
      <button @click="prev"
        class="w-9 h-9 rounded-full border border-[#C2B280]/20 flex items-center justify-center text-slate-600 hover:text-slate-900 hover:border-[#C2B280] transition-colors">
        <span class="material-symbols-outlined text-sm">chevron_left</span>
      </button>

      <div class="flex gap-2">
        <button v-for="(_, i) in visiblePlays" :key="i" @click="goTo(i)"
          :class="['w-2 h-2 rounded-full transition-colors', i === currentIndex ? 'bg-[#C2B280]' : 'bg-[#C2B280]/20 hover:bg-[#C2B280]']" />
      </div>

      <button @click="next"
        class="w-9 h-9 rounded-full border border-[#C2B280]/20 flex items-center justify-center text-slate-600 hover:text-slate-900 hover:border-[#C2B280] transition-colors">
        <span class="material-symbols-outlined text-sm">chevron_right</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { RouterLink } from 'vue-router'
import type { Play } from '../composables/useAdminApi'

const props = defineProps<{
  plays: Play[]
  showInactive?: boolean
}>()

const visiblePlays = computed(() =>
  props.showInactive ? props.plays : props.plays.filter((p) => p.active),
)

const currentIndex = ref(0)
const transitionName = ref('slide-left')

// Reset index when the list changes (e.g. data loads)
watch(visiblePlays, () => { currentIndex.value = 0 })

const currentPlay = computed(() => visiblePlays.value[currentIndex.value] ?? null)

function goTo(i: number) {
  transitionName.value = i > currentIndex.value ? 'slide-left' : 'slide-right'
  currentIndex.value = i
}

function next() {
  transitionName.value = 'slide-left'
  currentIndex.value = (currentIndex.value + 1) % visiblePlays.value.length
}

function prev() {
  transitionName.value = 'slide-right'
  currentIndex.value = (currentIndex.value - 1 + visiblePlays.value.length) % visiblePlays.value.length
}
</script>

<style scoped>
.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: all 0.3s ease;
}

.slide-left-enter-from {
  opacity: 0;
  transform: translateX(40px);
}

.slide-left-leave-to {
  opacity: 0;
  transform: translateX(-40px);
}

.slide-right-enter-from {
  opacity: 0;
  transform: translateX(-40px);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translateX(40px);
}
</style>
