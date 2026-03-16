<template>
  <div class="bg-[#221010] text-white">
    <!-- Hero -->
    <section class="bg-[#1a0a0a] py-20 border-b border-red-900/30">
      <div class="max-w-2xl mx-auto px-6 text-center">
        <p class="text-red-400 text-sm font-semibold uppercase tracking-widest mb-4">Contacto</p>
        <h1 class="text-4xl font-bold mb-4">Estamos para escucharte</h1>
        <p class="text-gray-400 leading-relaxed">
          ¿Tienes preguntas sobre nuestras obras, quieres colaborar o necesitas información para prensa? Escríbenos y te
          responderemos pronto.
        </p>
      </div>
    </section>

    <!-- Main content -->
    <section class="max-w-7xl mx-auto px-6 py-16">
      <div class="grid lg:grid-cols-5 gap-12">
        <!-- Form -->
        <div class="lg:col-span-3">
          <h2 class="text-2xl font-bold mb-7">Envíanos un mensaje</h2>
          <form @submit.prevent="handleSubmit" class="flex flex-col gap-5">
            <div class="grid sm:grid-cols-2 gap-5">
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-2">Nombre completo</label>
                <input v-model="form.name" type="text" placeholder="Tu nombre"
                  class="w-full bg-[#2d1515] border border-red-900/30 rounded px-4 py-3 text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-2">Correo electrónico</label>
                <input v-model="form.email" type="email" placeholder="tu@email.com"
                  class="w-full bg-[#2d1515] border border-red-900/30 rounded px-4 py-3 text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Asunto</label>
              <select v-model="form.subject"
                class="w-full bg-[#2d1515] border border-red-900/30 rounded px-4 py-3 text-white focus:outline-none focus:border-red-600 transition-colors">
                <option value="" disabled>Selecciona un asunto</option>
                <option>Información general</option>
                <option>Prensa y Medios</option>
                <option>Otros</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Mensaje</label>
              <textarea v-model="form.message" rows="5" placeholder="Escribe tu mensaje aquí..."
                class="w-full bg-[#2d1515] border border-red-900/30 rounded px-4 py-3 text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors resize-none" />
            </div>
            <div v-if="submitted"
              class="bg-green-900/30 border border-green-700/40 text-green-400 rounded px-4 py-3 text-sm flex items-center gap-2">
              <span class="material-symbols-outlined text-sm">check_circle</span>
              Mensaje enviado. Te responderemos en breve.
            </div>
            <button type="submit"
              class="bg-red-700 hover:bg-red-600 text-white font-medium px-8 py-3 rounded transition-colors flex items-center gap-2 self-start">
              <span class="material-symbols-outlined text-sm">send</span>
              Enviar Mensaje
            </button>
          </form>
        </div>

        <!-- Contact info -->
        <div class="lg:col-span-2 flex flex-col gap-6">
          <div class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-6">
            <h3 class="font-bold mb-5">Información de Contacto</h3>
            <div class="flex flex-col gap-5">
              <div class="flex items-start gap-4">
                <span class="material-symbols-outlined text-red-500 mt-0.5">mail</span>
                <div>
                  <p class="text-xs text-gray-500 mb-0.5 uppercase tracking-wider">Email</p>
                  <a href="mailto:hola@telonabierto.com" class="text-white hover:text-red-400 transition-colors">
                    gtvaldesoto@gmail.com
                  </a>
                </div>
              </div>
              <div class="flex items-start gap-4">
                <span class="material-symbols-outlined text-red-500 mt-0.5">phone</span>
                <div>
                  <p class="text-xs text-gray-500 mb-0.5 uppercase tracking-wider">Teléfono</p>
                  <a href="tel:+34900123456" class="text-white hover:text-red-400 transition-colors">
                    +34 606 58 54 12
                  </a>
                </div>
              </div>
            </div>
          </div>

          <!-- Social -->
          <div class="bg-[#2d1515]/60 border border-red-900/20 rounded-lg p-6">
            <h3 class="font-bold mb-4">Síguenos</h3>
            <div class="flex flex-col gap-3">
              <a v-for="social in socials" :key="social.name" href="#"
                class="flex items-center gap-3 text-gray-400 hover:text-white transition-colors group">
                <span class="material-symbols-outlined text-red-500 group-hover:text-red-400">{{ social.icon }}</span>
                <div>
                  <p class="text-sm font-medium text-white">{{ social.name }}</p>
                  <p class="text-xs text-gray-500">{{ social.handle }}</p>
                </div>
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Map placeholder -->
    <section class="h-64 bg-[#1a0a0a] border-t border-red-900/30 flex items-center justify-center">
      <div class="text-center">
        <span class="material-symbols-outlined text-red-700 text-5xl mb-3 block">map</span>
        <p class="text-gray-600 text-sm">Valdesoto · Asturies</p>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

const submitted = ref(false)

const form = reactive({
  name: '',
  email: '',
  subject: '',
  message: '',
})

function handleSubmit() {
  submitted.value = true
  Object.assign(form, { name: '', email: '', subject: '', message: '' })
  setTimeout(() => {
    submitted.value = false
  }, 5000)
}

const socials = [
  { name: 'Instagram', handle: '@gtsanfelixvaldesoto', icon: 'photo_camera' },
  { name: 'Facebook', handle: 'G.T. San Félix de Valdesoto', icon: 'thumb_up' },
  { name: 'TikTok', handle: '@teatru.valdesoto', icon: 'music_note' },
]
</script>
