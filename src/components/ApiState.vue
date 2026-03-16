<template>
  <div v-if="loading" :class="`flex justify-center ${py}`">
    <span class="material-symbols-outlined text-red-600 text-4xl animate-spin">progress_activity</span>
  </div>
  <div v-else-if="error" :class="`text-center ${py} text-gray-500`">
    <span class="material-symbols-outlined text-5xl mb-3 block text-red-800">wifi_off</span>
    <p>{{ error }}</p>
  </div>
  <slot v-else />
</template>

<script setup lang="ts">
defineProps<{
  loading: boolean
  error: string | null
  py?: string
}>()
</script>
