---
name: Duplication hotspots audit (March 2026)
description: Identified reusability opportunities across src/components/ — fetch boilerplate, date utilities, CRUD modal pattern, admin auth
type: project
---

Audit completed 2026-03-16. Key findings:

**Why:** App has grown to 7 view components; 3 public views independently duplicate API fetch + loading + error state boilerplate. AdminView has 3 nearly-identical CRUD modal blocks. Date helpers are copy-pasted between CalendarView and AdminView.

**How to apply:** When adding new views or CRUD resources, always reach for the composables listed below rather than writing fetch/loading/error state inline again.

## Recommended extractions (not yet implemented as of audit date)

1. `useApiRequest<T>` — Generic fetch wrapper returning `{ data, loading, error, execute }`. Eliminates the onMounted try/catch/finally pattern repeated in AboutView, CalendarView, RepertoireView, and AdminView.

2. `useCrud<T>` — Wraps a resource API object (list/create/update/delete) with `items`, `saving`, `modalError`, `save()`, `remove()` and optimistic list updates. Eliminates the three parallel CRUD blocks in AdminView (plays, performances, members).

3. `src/utils/dates.ts` — Pure utility: `parseDate`, `monthLabel`, `dayLabel`, `formatDate`. Currently duplicated verbatim between CalendarView and AdminView.

4. `useAdminAuth` composable — Centralises `adminUser` (from localStorage) and `logout()` currently inlined in AdminView. Single-concern composable, also makes logout testable.

## Files known to be duplication sources
- src/components/AdminView.vue — monolithic, 774 lines, contains all CRUD for plays/performances/members plus 3 modals
- src/components/RepertoireView.vue — independent fetch of playsApi.list()
- src/components/CalendarView.vue — independent fetch of performancesApi.list(), owns date helpers
- src/components/AboutView.vue — independent fetch of membersApi.list()
