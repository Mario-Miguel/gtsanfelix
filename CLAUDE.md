# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Backend (Go)

```bash
cd backend
go mod tidy          # download dependencies (first time)
go run main.go       # start API server at http://localhost:8080
go build -o server . # compile binary
```

Requires **Go 1.22+** (uses standard-library method routing and `PathValue`).
Mock credentials: `admin@telonabierto.com` / `admin123`.

## Frontend Commands

```bash
pnpm dev          # Dev server at http://localhost:5173
pnpm build        # Type-check + production build
pnpm type-check   # vue-tsc type checking only
pnpm lint         # oxlint + eslint (auto-fixes)
pnpm test:unit    # Vitest unit tests
pnpm test:unit -- --run src/__tests__/App.spec.ts  # Single test file
pnpm test:e2e:dev # Cypress e2e against dev server
```

## Architecture

**Stack:** Vue 3 (Composition API + `<script setup>`) · Vite · TypeScript · Tailwind CSS v4 · Pinia · Vue Router · `@nuxt/ui` component library (used in a plain Vite/Vue app, not Nuxt).

**Entry flow:** `index.html` → `src/main.ts` → `src/App.vue` → `RouterView`.

**Routing:** All routes live in `src/router/index.ts`. Page components live in `src/components/` (not a separate `pages/` directory).

| Route | Component |
|---|---|
| `/` | `HomeView.vue` |
| `/about` | `AboutView.vue` |
| `/repertoire` | `RepertoireView.vue` |
| `/calendar` | `CalendarView.vue` |
| `/contact` | `ContactView.vue` |
| `/admin/login` | `AdminLoginView.vue` |
| `/admin` | `AdminView.vue` *(requires JWT)* |

**Layout:** Two distinct layouts exist. Public routes use `App.vue`'s fixed top navbar (`h-[73px]`) + full-width `<main class="pt-[73px]">` — page components own full-width sections, do not add a constraining wrapper. The `/admin` route bypasses the top navbar entirely (`App.vue` checks `route.path.startsWith('/admin')`); `AdminView.vue` self-contains a sidebar + content layout using `h-screen overflow-hidden flex`.

**Styling:** Tailwind v4 with `@nuxt/ui` imported in `src/assets/main.css`. Custom color palette (dark red theater theme, background `#221010`, primary red `#d41111`) is defined as CSS variables in `main.css`. The `tailwind.config.cjs` maps these to Tailwind color tokens (`primary-*`, `secondary-*`, etc.). Use inline Tailwind classes; scoped `<style>` blocks are rare.

**Icons:** Google Material Symbols loaded via CDN in `index.html`. Use `<span class="material-symbols-outlined">icon_name</span>`.

**State:** Pinia store at `src/stores/counter.ts` (currently a scaffold). Add domain stores there.

**Path alias:** `@` → `src/` (configured in `vite.config.ts`).

**API client:** `src/composables/useAdminApi.ts` exports typed async functions (`playsApi`, `membersApi`, `performancesApi`, `login`). JWT is stored in `localStorage` as `admin_token`. The navigation guard in `router/index.ts` redirects unauthenticated visits to `/admin` → `/admin/login`.

**Backend structure (`backend/`):**
```
main.go          — HTTP server, route registration (Go 1.22 mux with method+path patterns)
auth/auth.go     — JWT generation and validation (HS256)
store/store.go   — Thread-safe in-memory store, seeded with mock data
models/models.go — Shared types: Play, Member, Performance, User
handlers/        — One file per resource + helpers.go (jsonOK/jsonError)
middleware/      — CORS (allows localhost:5173) + RequireAuth JWT guard
```
Write endpoints (`POST`, `PUT`, `DELETE`) require `Authorization: Bearer <token>`. GET endpoints are public.

**Tests:** Unit tests with Vitest + `@vue/test-utils` in `src/__tests__/`. E2e with Cypress in `cypress/e2e/`.

## Known pitfalls

**v-calendar `DatePicker`:** No usar `@dayclick` para llamar a `target.blur()` — interrumpe el procesamiento interno y provoca `TypeError: Cannot read properties of undefined (reading 'dayIndex')`. Usar siempre `:masks="{ modelValue: 'YYYY-MM-DD' }"` junto con `v-model.string`; sin esta máscara el valor se guarda en ISO (`2024-01-15T00:00:00.000Z`) y las funciones de formateo que hacen `dateStr + 'T00:00:00'` producen fechas inválidas. Usar slot `#default` con `inputValue` + `inputEvents` para input personalizado.
