<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <header class="backdrop-blur-sm border-b bg-primary-500 text-white">
      <nav class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <div class="text-lg font-semibold">G.T. Valdesoto</div>
        <ul class="flex items-center gap-4 text-sm">
          <li><RouterLink to="/">Home</RouterLink></li>
          <li><RouterLink to="/about">About</RouterLink></li>
          <li><RouterLink to="/contact">Contact</RouterLink></li>
          <li><RouterLink to="/calendar">Calendar</RouterLink></li>
        </ul>
      </nav>
    </header>

    <main class="flex-1 min-h-0 overflow-auto bg-secondary-100">
      <div class="max-w-7xl mx-auto p-6">
        <RouterView />
      </div>
    </main>
  </div>
</template>

<style scoped>
/* Ensure the layout fills the viewport and that the scrolling region behaves correctly */
.min-h-0 {
  min-height: 0;
}
</style>
