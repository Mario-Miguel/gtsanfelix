/**
 * Typed API client for the Go backend (http://localhost:8080).
 * All write operations require a valid JWT stored in localStorage.
 */

import { TOKEN_KEY } from './useAdminAuth'

const BASE = 'http://localhost:8080/api'

// ── Types ─────────────────────────────────────────────────────────────────

export interface Play {
  id: number
  title: string
  author: string
  genre: string
  duration: string
  imageUrl?: string
  active: boolean
  summary: string
}

export interface Member {
  id: number
  name: string
  role: string
  email: string
  quote?: string
}

export interface Performance {
  id: number
  playId: number
  playTitle: string
  date: string
  time: string
  venue: string
}

export interface LoginResponse {
  token: string
  user: { id: number; email: string; name: string; role: string }
}

// ── Helpers ───────────────────────────────────────────────────────────────

function authHeaders(): Record<string, string> {
  const token = localStorage.getItem(TOKEN_KEY)
  return {
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  }
}

async function request<T>(url: string, init?: RequestInit): Promise<T> {
  const res = await fetch(url, init)
  if (res.status === 204) return undefined as T
  const data = await res.json()
  if (!res.ok) throw new Error(data.error ?? 'Error desconocido')
  return data as T
}

// ── Auth ──────────────────────────────────────────────────────────────────

export function login(email: string, password: string): Promise<LoginResponse> {
  return request(`${BASE}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  })
}

// ── Plays ─────────────────────────────────────────────────────────────────

export const playsApi = {
  list: () => request<Play[]>(`${BASE}/plays`),
  get: (id: number) => request<Play>(`${BASE}/plays/${id}`),
  create: (data: Omit<Play, 'id'>) =>
    request<Play>(`${BASE}/plays`, { method: 'POST', headers: authHeaders(), body: JSON.stringify(data) }),
  update: (id: number, data: Omit<Play, 'id'>) =>
    request<Play>(`${BASE}/plays/${id}`, { method: 'PUT', headers: authHeaders(), body: JSON.stringify(data) }),
  delete: (id: number) =>
    request<void>(`${BASE}/plays/${id}`, { method: 'DELETE', headers: authHeaders() }),
}

// ── Members ───────────────────────────────────────────────────────────────

export const membersApi = {
  list: () => request<Member[]>(`${BASE}/members`),
  get: (id: number) => request<Member>(`${BASE}/members/${id}`),
  create: (data: Omit<Member, 'id'>) =>
    request<Member>(`${BASE}/members`, { method: 'POST', headers: authHeaders(), body: JSON.stringify(data) }),
  update: (id: number, data: Omit<Member, 'id'>) =>
    request<Member>(`${BASE}/members/${id}`, { method: 'PUT', headers: authHeaders(), body: JSON.stringify(data) }),
  delete: (id: number) =>
    request<void>(`${BASE}/members/${id}`, { method: 'DELETE', headers: authHeaders() }),
}

// ── Performances ──────────────────────────────────────────────────────────

export const performancesApi = {
  list: () => request<Performance[]>(`${BASE}/performances`),
  get: (id: number) => request<Performance>(`${BASE}/performances/${id}`),
  create: (data: Omit<Performance, 'id'>) =>
    request<Performance>(`${BASE}/performances`, { method: 'POST', headers: authHeaders(), body: JSON.stringify(data) }),
  update: (id: number, data: Omit<Performance, 'id'>) =>
    request<Performance>(`${BASE}/performances/${id}`, { method: 'PUT', headers: authHeaders(), body: JSON.stringify(data) }),
  delete: (id: number) =>
    request<void>(`${BASE}/performances/${id}`, { method: 'DELETE', headers: authHeaders() }),
}
