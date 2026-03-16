<template>
  <div class="flex h-screen bg-[#160808] text-white overflow-hidden">
    <!-- Sidebar -->
    <aside class="w-64 flex-shrink-0 bg-[#0f0505] border-r border-red-900/30 flex flex-col">
      <div class="px-6 py-5 border-b border-red-900/30">
        <RouterLink to="/" class="flex items-center gap-2 hover:opacity-80 transition-opacity">
          <span class="material-symbols-outlined text-red-600">theater_comedy</span>
          <span class="font-bold tracking-wide">G.T. San Félix de Valdesoto</span>
        </RouterLink>
        <p class="text-xs text-red-900 mt-1 ml-7">Panel de Administración</p>
      </div>

      <nav class="flex-1 px-3 py-4 flex flex-col gap-1">
        <button v-for="item in navItems" :key="item.label" :class="[
          'flex items-center gap-3 px-4 py-2.5 rounded text-sm font-medium w-full text-left transition-colors',
          activeSection === item.section
            ? 'bg-red-700/30 text-red-400'
            : 'text-gray-400 hover:bg-red-900/20 hover:text-white',
        ]" @click="activeSection = item.section">
          <span class="material-symbols-outlined text-[18px]">{{ item.icon }}</span>
          {{ item.label }}
        </button>
      </nav>

      <!-- User + Logout -->
      <div class="px-4 py-4 border-t border-red-900/30">
        <div class="flex items-center gap-3 mb-3">
          <div class="w-8 h-8 rounded-full bg-red-800 flex items-center justify-center flex-shrink-0">
            <span class="material-symbols-outlined text-sm">person</span>
          </div>
          <div class="overflow-hidden">
            <p class="text-xs font-medium truncate">{{ adminUser?.email ?? 'admin' }}</p>
            <p class="text-xs text-gray-600">Administrador</p>
          </div>
        </div>
        <button @click="logout"
          class="flex items-center gap-2 text-xs text-gray-500 hover:text-red-400 transition-colors w-full">
          <span class="material-symbols-outlined text-sm">logout</span>
          Cerrar sesión
        </button>
      </div>
    </aside>

    <!-- Main content -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Top bar -->
      <header class="bg-[#0f0505] border-b border-red-900/30 px-8 py-4 flex items-center justify-between flex-shrink-0">
        <div>
          <h1 class="text-xl font-bold">{{ currentMeta.title }}</h1>
          <p class="text-xs text-gray-500 mt-0.5">{{ currentMeta.subtitle }}</p>
        </div>
        <div class="flex items-center gap-3">
          <button v-if="activeSection === 'dashboard' || activeSection === 'repertoire'" @click="openPlayModal()"
            class="flex items-center gap-2 bg-red-700 hover:bg-red-600 px-4 py-2 rounded text-sm font-medium transition-colors">
            <span class="material-symbols-outlined text-sm">add</span>
            Nueva Obra
          </button>
          <button v-if="activeSection === 'calendar'" @click="openPerfModal()"
            class="flex items-center gap-2 bg-red-700 hover:bg-red-600 px-4 py-2 rounded text-sm font-medium transition-colors">
            <span class="material-symbols-outlined text-sm">add</span>
            Nueva Función
          </button>
          <button v-if="activeSection === 'users'" @click="openMemberModal()"
            class="flex items-center gap-2 bg-red-700 hover:bg-red-600 px-4 py-2 rounded text-sm font-medium transition-colors">
            <span class="material-symbols-outlined text-sm">add</span>
            Nuevo Miembro
          </button>
        </div>
      </header>

      <!-- API error banner -->
      <div v-if="apiError"
        class="bg-red-900/30 border-b border-red-800/40 px-8 py-3 text-red-400 text-sm flex items-center gap-2">
        <span class="material-symbols-outlined text-sm">warning</span>
        {{ apiError }}
      </div>

      <!-- Loading overlay -->
      <div v-if="loading" class="flex-1 flex items-center justify-center">
        <span class="material-symbols-outlined text-red-600 text-4xl animate-spin">progress_activity</span>
      </div>

      <!-- Scrollable content -->
      <main v-else class="flex-1 overflow-y-auto p-8">

        <!-- ── DASHBOARD ── -->
        <template v-if="activeSection === 'dashboard'">
          <div class="grid grid-cols-2 xl:grid-cols-4 gap-5 mb-8">
            <div v-for="stat in computedStats" :key="stat.label"
              class="bg-[#1a0a0a] border border-red-900/20 rounded-lg p-5">
              <div class="flex items-center justify-between mb-3">
                <span class="material-symbols-outlined text-red-500">{{ stat.icon }}</span>
              </div>
              <p class="text-3xl font-bold mb-1">{{ stat.value }}</p>
              <p class="text-xs text-gray-500">{{ stat.label }}</p>
            </div>
          </div>

          <div class="grid xl:grid-cols-5 gap-6">
            <!-- Repertoire preview -->
            <div class="xl:col-span-3 bg-[#1a0a0a] border border-red-900/20 rounded-lg overflow-hidden">
              <div class="flex items-center justify-between px-6 py-4 border-b border-red-900/20">
                <h2 class="font-semibold text-sm">Gestión de Repertorio</h2>
                <button @click="activeSection = 'repertoire'"
                  class="text-red-400 hover:text-red-300 text-xs transition-colors">
                  Ver todo →
                </button>
              </div>
              <table class="w-full text-sm">
                <thead>
                  <tr class="text-gray-600 text-xs uppercase tracking-wider border-b border-red-900/10">
                    <th class="px-6 py-3 text-left">Título</th>
                    <th class="px-6 py-3 text-left hidden md:table-cell">Autor</th>
                    <th class="px-6 py-3 text-right">Acciones</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="play in plays.slice(0, 5)" :key="play.id"
                    class="border-b border-red-900/10 hover:bg-red-900/5 transition-colors">
                    <td class="px-6 py-3 font-medium">{{ play.title }}</td>
                    <td class="px-6 py-3 text-gray-400 hidden md:table-cell">{{ play.author }}</td>
                    <td class="px-6 py-3 text-right">
                      <div class="flex items-center justify-end gap-2">
                        <button @click="openPlayModal(play)"
                          class="p-1 text-gray-500 hover:text-white transition-colors">
                          <span class="material-symbols-outlined text-sm">edit</span>
                        </button>
                        <button @click="removePlay(play.id)"
                          class="p-1 text-gray-500 hover:text-red-400 transition-colors">
                          <span class="material-symbols-outlined text-sm">delete</span>
                        </button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>

            <!-- Upcoming performances -->
            <div class="xl:col-span-2 bg-[#1a0a0a] border border-red-900/20 rounded-lg overflow-hidden">
              <div class="flex items-center justify-between px-6 py-4 border-b border-red-900/20">
                <h2 class="font-semibold text-sm">Próximas Funciones</h2>
                <button @click="activeSection = 'calendar'"
                  class="text-red-400 hover:text-red-300 text-xs transition-colors">
                  Ver todo →
                </button>
              </div>
              <div class="divide-y divide-red-900/10">
                <div v-for="perf in performances.slice(0, 5)" :key="perf.id"
                  class="px-6 py-4 hover:bg-red-900/5 transition-colors">
                  <div class="flex items-start gap-3">
                    <div class="bg-red-700/20 text-red-400 rounded px-2 py-1 text-center min-w-[44px] flex-shrink-0">
                      <p class="text-xs leading-none">{{ monthLabel(perf.date) }}</p>
                      <p class="text-lg font-bold leading-none mt-0.5">{{ dayLabel(perf.date) }}</p>
                    </div>
                    <div class="flex-1 min-w-0">
                      <p class="font-medium text-sm truncate">{{ perf.playTitle }}</p>
                      <p class="text-xs text-gray-500 mt-0.5">{{ perf.time }} · {{ perf.venue }}</p>
                    </div>
                  </div>
                </div>
                <p v-if="performances.length === 0" class="px-6 py-8 text-center text-gray-600 text-sm">
                  No hay funciones programadas.
                </p>
              </div>
            </div>
          </div>
        </template>

        <!-- ── REPERTOIRE ── -->
        <template v-else-if="activeSection === 'repertoire'">
          <div class="bg-[#1a0a0a] border border-red-900/20 rounded-lg overflow-hidden">
            <div class="px-6 py-4 border-b border-red-900/20 flex items-center gap-4">
              <div class="flex-1 relative">
                <span
                  class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-gray-600 text-sm">search</span>
                <input v-model="searchQuery" type="text" placeholder="Buscar obra..."
                  class="w-full bg-[#2d1515] border border-red-900/20 rounded pl-9 pr-4 py-2 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
              </div>
              <select v-model="filterGenre"
                class="bg-[#2d1515] border border-red-900/20 rounded px-3 py-2 text-sm text-gray-300 focus:outline-none">
                <option value="">Todos los géneros</option>
                <option v-for="g in genres" :key="g">{{ g }}</option>
              </select>
            </div>
            <table class="w-full text-sm">
              <thead>
                <tr class="text-gray-600 text-xs uppercase tracking-wider border-b border-red-900/10">
                  <th class="px-6 py-3 text-left">Título</th>
                  <th class="px-6 py-3 text-left">Autor</th>
                  <th class="px-6 py-3 text-left">Género</th>
                  <th class="px-6 py-3 text-left">Duración</th>
                  <th class="px-6 py-3 text-right">Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="play in filteredPlays" :key="play.id"
                  class="border-b border-red-900/10 hover:bg-red-900/5 transition-colors">
                  <td class="px-6 py-3 font-medium">{{ play.title }}</td>
                  <td class="px-6 py-3 text-gray-400">{{ play.author }}</td>
                  <td class="px-6 py-3">
                    <span class="bg-red-700/20 text-red-400 text-xs px-2 py-0.5 rounded-full">{{ play.genre }}</span>
                  </td>
                  <td class="px-6 py-3 text-gray-400">{{ play.duration }}</td>
                  <td class="px-6 py-3 text-right">
                    <div class="flex items-center justify-end gap-2">
                      <button @click="openPlayModal(play)"
                        class="p-1.5 text-gray-500 hover:text-white hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">edit</span>
                      </button>
                      <button @click="removePlay(play.id)"
                        class="p-1.5 text-gray-500 hover:text-red-400 hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">delete</span>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="filteredPlays.length === 0">
                  <td colspan="5" class="px-6 py-10 text-center text-gray-600">No se encontraron obras.</td>
                </tr>
              </tbody>
            </table>
          </div>
        </template>

        <!-- ── CALENDAR ── -->
        <template v-else-if="activeSection === 'calendar'">
          <div class="bg-[#1a0a0a] border border-red-900/20 rounded-lg overflow-hidden">
            <table class="w-full text-sm">
              <thead>
                <tr class="text-gray-600 text-xs uppercase tracking-wider border-b border-red-900/10">
                  <th class="px-6 py-3 text-left">Fecha</th>
                  <th class="px-6 py-3 text-left">Obra</th>
                  <th class="px-6 py-3 text-left">Hora</th>
                  <th class="px-6 py-3 text-left">Sede</th>
                  <th class="px-6 py-3 text-left">Ocupación</th>
                  <th class="px-6 py-3 text-right">Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="perf in performances" :key="perf.id"
                  class="border-b border-red-900/10 hover:bg-red-900/5 transition-colors">
                  <td class="px-6 py-3 font-medium">{{ formatDate(perf.date) }}</td>
                  <td class="px-6 py-3">{{ perf.playTitle }}</td>
                  <td class="px-6 py-3 text-gray-400">{{ perf.time }}</td>
                  <td class="px-6 py-3 text-gray-400">{{ perf.venue }}</td>

                  <td class="px-6 py-3 text-right">
                    <div class="flex items-center justify-end gap-2">
                      <button @click="openPerfModal(perf)"
                        class="p-1.5 text-gray-500 hover:text-white hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">edit</span>
                      </button>
                      <button @click="removePerformance(perf.id)"
                        class="p-1.5 text-gray-500 hover:text-red-400 hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">delete</span>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="performances.length === 0">
                  <td colspan="6" class="px-6 py-10 text-center text-gray-600">No hay funciones programadas.</td>
                </tr>
              </tbody>
            </table>
          </div>
        </template>

        <!-- ── USERS / MEMBERS ── -->
        <template v-else-if="activeSection === 'users'">
          <div class="bg-[#1a0a0a] border border-red-900/20 rounded-lg overflow-hidden">
            <table class="w-full text-sm">
              <thead>
                <tr class="text-gray-600 text-xs uppercase tracking-wider border-b border-red-900/10">
                  <th class="px-6 py-3 text-left">Miembro</th>
                  <th class="px-6 py-3 text-left">Email</th>
                  <th class="px-6 py-3 text-left">Rol</th>
                  <th class="px-6 py-3 text-left">Estado</th>
                  <th class="px-6 py-3 text-right">Acciones</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="member in members" :key="member.id"
                  class="border-b border-red-900/10 hover:bg-red-900/5 transition-colors">
                  <td class="px-6 py-3">
                    <div class="flex items-center gap-3">
                      <div class="w-8 h-8 rounded-full bg-red-800/50 flex items-center justify-center">
                        <span class="material-symbols-outlined text-sm text-red-400">person</span>
                      </div>
                      <span class="font-medium">{{ member.name }}</span>
                    </div>
                  </td>
                  <td class="px-6 py-3 text-gray-400">{{ member.email }}</td>
                  <td class="px-6 py-3 text-gray-300">{{ member.role }}</td>
                  <td class="px-6 py-3 text-right">
                    <div class="flex items-center justify-end gap-2">
                      <button @click="openMemberModal(member)"
                        class="p-1.5 text-gray-500 hover:text-white hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">edit</span>
                      </button>
                      <button @click="removeMember(member.id)"
                        class="p-1.5 text-gray-500 hover:text-red-400 hover:bg-red-900/20 rounded transition-colors">
                        <span class="material-symbols-outlined text-sm">delete</span>
                      </button>
                    </div>
                  </td>
                </tr>
                <tr v-if="members.length === 0">
                  <td colspan="5" class="px-6 py-10 text-center text-gray-600">No hay miembros registrados.</td>
                </tr>
              </tbody>
            </table>
          </div>
        </template>

        <!-- ── SETTINGS ── -->
        <template v-else-if="activeSection === 'settings'">
          <div class="max-w-2xl flex flex-col gap-6">
            <div class="bg-[#1a0a0a] border border-red-900/20 rounded-lg p-6">
              <h2 class="font-semibold mb-5">Información de la Compañía</h2>
              <div class="flex flex-col gap-4">
                <div>
                  <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Nombre</label>
                  <input type="text" value="G.T. San Félix de Valdesoto"
                    class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors" />
                </div>
                <div>
                  <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Email de contacto</label>
                  <input type="email" value="hola@telonabierto.com"
                    class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors" />
                </div>
                <div>
                  <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Dirección</label>
                  <input type="text" value="Calle de la Comedia, 42, 28004 Madrid"
                    class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors" />
                </div>
                <button
                  class="self-start bg-red-700 hover:bg-red-600 px-5 py-2 rounded text-sm font-medium transition-colors">
                  Guardar Cambios
                </button>
              </div>
            </div>
          </div>
        </template>

      </main>
    </div>
  </div>

  <!-- ── Modal: Obra ── -->
  <Teleport to="body">
    <div v-if="showPlayModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
      @click.self="showPlayModal = false">
      <div class="bg-[#1a0a0a] border border-red-900/30 rounded-lg w-full max-w-lg mx-4 p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="font-bold text-lg">{{ editingPlay ? 'Editar Obra' : 'Nueva Obra' }}</h2>
          <button @click="showPlayModal = false" class="text-gray-500 hover:text-white transition-colors">
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>
        <form @submit.prevent="savePlay" class="flex flex-col gap-4">
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Título *</label>
            <input v-model="playForm.title" type="text" required placeholder="Título de la obra"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
          </div>
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Autor *</label>
            <input v-model="playForm.author" type="text" required placeholder="Nombre del autor"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Género</label>
              <select v-model="playForm.genre"
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors">
                <option v-for="g in genres" :key="g">{{ g }}</option>
              </select>
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Duración</label>
              <input v-model="playForm.duration" type="text" placeholder="ej. 90 min"
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
            </div>
          </div>
          <p v-if="playsModalError" class="text-red-400 text-xs">{{ playsModalError }}</p>
          <div class="flex justify-end gap-3 mt-2">
            <button type="button" @click="showPlayModal = false"
              class="px-5 py-2 text-sm text-gray-400 hover:text-white border border-red-900/30 rounded transition-colors">Cancelar</button>
            <button type="submit" :disabled="playsSaving"
              class="px-5 py-2 text-sm bg-red-700 hover:bg-red-600 disabled:opacity-50 rounded font-medium transition-colors">
              {{ playsSaving ? 'Guardando...' : editingPlay ? 'Guardar' : 'Añadir' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>

  <!-- ── Modal: Función / Performance ── -->
  <Teleport to="body">
    <div v-if="showPerfModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
      @click.self="showPerfModal = false">
      <div class="bg-[#1a0a0a] border border-red-900/30 rounded-lg w-full max-w-lg mx-4 p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="font-bold text-lg">{{ editingPerf ? 'Editar Función' : 'Nueva Función' }}</h2>
          <button @click="showPerfModal = false" class="text-gray-500 hover:text-white transition-colors">
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>
        <form @submit.prevent="savePerformance" class="flex flex-col gap-4">
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Título de la obra *</label>
            <input v-model="perfForm.playTitle" type="text" required placeholder="Título"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Fecha *</label>
              <input v-model="perfForm.date" type="date" required
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors" />
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Hora *</label>
              <input v-model="perfForm.time" type="time" required
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white focus:outline-none focus:border-red-600 transition-colors" />
            </div>
          </div>
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Sede *</label>
            <input v-model="perfForm.venue" type="text" required placeholder="Teatro Principal"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
          </div>
          <p v-if="perfsModalError" class="text-red-400 text-xs">{{ perfsModalError }}</p>
          <div class="flex justify-end gap-3 mt-2">
            <button type="button" @click="showPerfModal = false"
              class="px-5 py-2 text-sm text-gray-400 hover:text-white border border-red-900/30 rounded transition-colors">Cancelar</button>
            <button type="submit" :disabled="perfsSaving"
              class="px-5 py-2 text-sm bg-red-700 hover:bg-red-600 disabled:opacity-50 rounded font-medium transition-colors">
              {{ perfsSaving ? 'Guardando...' : editingPerf ? 'Guardar' : 'Añadir' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>

  <!-- ── Modal: Miembro ── -->
  <Teleport to="body">
    <div v-if="showMemberModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
      @click.self="showMemberModal = false">
      <div class="bg-[#1a0a0a] border border-red-900/30 rounded-lg w-full max-w-lg mx-4 p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="font-bold text-lg">{{ editingMember ? 'Editar Miembro' : 'Nuevo Miembro' }}</h2>
          <button @click="showMemberModal = false" class="text-gray-500 hover:text-white transition-colors">
            <span class="material-symbols-outlined">close</span>
          </button>
        </div>
        <form @submit.prevent="saveMember" class="flex flex-col gap-4">
          <div>
            <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Nombre *</label>
            <input v-model="memberForm.name" type="text" required placeholder="Nombre completo"
              class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Rol</label>
              <input v-model="memberForm.role" type="text" placeholder="Actor, Director..."
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1.5 uppercase tracking-wider">Email *</label>
              <input v-model="memberForm.email" type="email" required placeholder="email@ejemplo.com"
                class="w-full bg-[#2d1515] border border-red-900/20 rounded px-4 py-2.5 text-sm text-white placeholder-gray-600 focus:outline-none focus:border-red-600 transition-colors" />
            </div>
          </div>

          <label class="flex items-center gap-3 cursor-pointer">
            <input v-model="memberForm.active" type="checkbox" class="accent-red-600 w-4 h-4" />
            <span class="text-sm text-gray-300">Miembro activo</span>
          </label>
          <p v-if="membersModalError" class="text-red-400 text-xs">{{ membersModalError }}</p>
          <div class="flex justify-end gap-3 mt-2">
            <button type="button" @click="showMemberModal = false"
              class="px-5 py-2 text-sm text-gray-400 hover:text-white border border-red-900/30 rounded transition-colors">Cancelar</button>
            <button type="submit" :disabled="membersSaving"
              class="px-5 py-2 text-sm bg-red-700 hover:bg-red-600 disabled:opacity-50 rounded font-medium transition-colors">
              {{ membersSaving ? 'Guardando...' : editingMember ? 'Guardar' : 'Añadir' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import {
  playsApi, membersApi, performancesApi,
  type Play, type Member, type Performance,
} from '../composables/useAdminApi'
import { useCrud } from '../composables/useCrud'
import { useAdminAuth } from '../composables/useAdminAuth'

type Section = 'dashboard' | 'repertoire' | 'calendar' | 'users' | 'settings'

const { adminUser, logout } = useAdminAuth()

const activeSection = ref<Section>('dashboard')
const loading = ref(true)
const apiError = ref('')
const searchQuery = ref('')
const filterGenre = ref('')

// ── Navigation ────────────────────────────────────────────────────────────
const navItems: { label: string; icon: string; section: Section }[] = [
  { label: 'Dashboard', icon: 'dashboard', section: 'dashboard' },
  { label: 'Repertorio', icon: 'menu_book', section: 'repertoire' },
  { label: 'Calendario', icon: 'calendar_month', section: 'calendar' },
  { label: 'Miembros', icon: 'group', section: 'users' },
  { label: 'Ajustes', icon: 'settings', section: 'settings' },
]

const sectionMeta: Record<Section, { title: string; subtitle: string }> = {
  dashboard: { title: 'Dashboard', subtitle: 'Resumen general de la compañía' },
  repertoire: { title: 'Repertorio', subtitle: 'Gestión de obras teatrales' },
  calendar: { title: 'Calendario', subtitle: 'Programación de funciones' },
  users: { title: 'Miembros', subtitle: 'Gestión de miembros de la compañía' },
  settings: { title: 'Ajustes', subtitle: 'Configuración del sistema' },
}
const currentMeta = computed(() => sectionMeta[activeSection.value])

// ── Data (CRUD) ────────────────────────────────────────────────────────────
const {
  items: plays,
  saving: playsSaving,
  modalError: playsModalError,
  save: saveCrudPlay,
  remove: removeCrudPlay,
} = useCrud<Play>(playsApi, { prepend: true })

const {
  items: performances,
  saving: perfsSaving,
  modalError: perfsModalError,
  save: saveCrudPerf,
  remove: removeCrudPerf,
} = useCrud<Performance>(performancesApi)

const {
  items: members,
  saving: membersSaving,
  modalError: membersModalError,
  save: saveCrudMember,
  remove: removeCrudMember,
} = useCrud<Member>(membersApi)

const genres = ['Drama', 'Tragedia', 'Teatro del Absurdo', 'Clásico', 'Contemporáneo', 'Comedia']

const filteredPlays = computed(() =>
  plays.value.filter((p) => {
    const q = searchQuery.value.toLowerCase()
    return (
      (!q || p.title.toLowerCase().includes(q) || p.author.toLowerCase().includes(q)) &&
      (!filterGenre.value || p.genre === filterGenre.value)
    )
  }),
)

const computedStats = computed(() => [
  { label: 'Obras en Repertorio', value: plays.value.length, icon: 'menu_book' },
  { label: 'Funciones Programadas', value: performances.value.length, icon: 'event' },
])

async function loadData() {
  loading.value = true
  apiError.value = ''
  try {
    const [p, m, perf] = await Promise.all([playsApi.list(), membersApi.list(), performancesApi.list()])
    plays.value = p
    members.value = m
    performances.value = perf
  } catch {
    apiError.value = 'No se pudo conectar con el servidor. Asegúrate de que el backend está corriendo en :8080'
  } finally {
    loading.value = false
  }
}

onMounted(loadData)

// ── Date helpers ──────────────────────────────────────────────────────────
const MONTHS = ['ENE', 'FEB', 'MAR', 'ABR', 'MAY', 'JUN', 'JUL', 'AGO', 'SEP', 'OCT', 'NOV', 'DIC']
function monthLabel(dateStr: string) {
  const d = new Date(dateStr + 'T00:00:00')
  return MONTHS[d.getMonth()] ?? ''
}
function dayLabel(dateStr: string) {
  return new Date(dateStr + 'T00:00:00').getDate()
}
function formatDate(dateStr: string) {
  return new Date(dateStr + 'T00:00:00').toLocaleDateString('es-ES', { day: 'numeric', month: 'short', year: 'numeric' })
}

// ── Plays CRUD ─────────────────────────────────────────────────────────────
const showPlayModal = ref(false)
const editingPlay = ref<Play | null>(null)
const playForm = reactive({ title: '', author: '', genre: 'Comedia', duration: '', active: false })

function openPlayModal(play?: Play) {
  editingPlay.value = play ?? null
  Object.assign(playForm, play ? { title: play.title, author: play.author, genre: play.genre, duration: play.duration, active: play.active } : { title: '', author: '', genre: 'Comedia', duration: '', active: false })
  playsModalError.value = ''
  showPlayModal.value = true
}

async function savePlay() {
  const ok = await saveCrudPlay(playForm, editingPlay.value?.id)
  if (ok) showPlayModal.value = false
}

async function removePlay(id: number) {
  if (!confirm('¿Eliminar esta obra?')) return
  try {
    await removeCrudPlay(id)
  } catch (e) {
    apiError.value = e instanceof Error ? e.message : 'Error al eliminar'
  }
}

// ── Performances CRUD ─────────────────────────────────────────────────────
const showPerfModal = ref(false)
const editingPerf = ref<Performance | null>(null)
const perfForm = reactive({ playId: 0, playTitle: '', date: '', time: '', venue: '' })

function openPerfModal(perf?: Performance) {
  editingPerf.value = perf ?? null
  Object.assign(perfForm, perf ?? { playId: 0, playTitle: '', date: '', time: '', venue: '' })
  perfsModalError.value = ''
  showPerfModal.value = true
}

async function savePerformance() {
  const ok = await saveCrudPerf(perfForm, editingPerf.value?.id)
  if (ok) showPerfModal.value = false
}

async function removePerformance(id: number) {
  if (!confirm('¿Eliminar esta función?')) return
  try {
    await removeCrudPerf(id)
  } catch (e) {
    apiError.value = e instanceof Error ? e.message : 'Error al eliminar'
  }
}

// ── Members CRUD ──────────────────────────────────────────────────────────
const showMemberModal = ref(false)
const editingMember = ref<Member | null>(null)
const memberForm = reactive({ name: '', role: '', email: '', active: true, quote: '' })

function openMemberModal(member?: Member) {
  editingMember.value = member ?? null
  Object.assign(memberForm, member ?? { name: '', role: '', email: '', active: true, quote: '' })
  membersModalError.value = ''
  showMemberModal.value = true
}

async function saveMember() {
  const ok = await saveCrudMember(memberForm, editingMember.value?.id)
  if (ok) showMemberModal.value = false
}

async function removeMember(id: number) {
  if (!confirm('¿Eliminar este miembro?')) return
  try {
    await removeCrudMember(id)
  } catch (e) {
    apiError.value = e instanceof Error ? e.message : 'Error al eliminar'
  }
}
</script>
