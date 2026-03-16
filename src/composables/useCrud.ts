import { ref, type Ref } from 'vue'

interface CrudApi<T> {
  list: () => Promise<T[]>
  create: (data: Omit<T, 'id'>) => Promise<T>
  update: (id: number, data: Omit<T, 'id'>) => Promise<T>
  delete: (id: number) => Promise<void>
}

export function useCrud<T extends { id: number }>(
  api: CrudApi<T>,
  options: { prepend?: boolean } = {},
): {
  items: Ref<T[]>
  saving: Ref<boolean>
  modalError: Ref<string>
  save: (formData: Omit<T, 'id'>, editingId?: number) => Promise<boolean>
  remove: (id: number) => Promise<void>
} {
  const { prepend = false } = options

  const items = ref<T[]>([]) as Ref<T[]>
  const saving = ref(false)
  const modalError = ref('')

  async function save(formData: Omit<T, 'id'>, editingId?: number): Promise<boolean> {
    saving.value = true
    modalError.value = ''
    try {
      if (editingId !== undefined) {
        const updated = await api.update(editingId, formData)
        const idx = items.value.findIndex((x) => x.id === editingId)
        if (idx !== -1) items.value[idx] = updated
      } else {
        const created = await api.create(formData)
        if (prepend) items.value.unshift(created)
        else items.value.push(created)
      }
      return true
    } catch (e) {
      modalError.value = e instanceof Error ? e.message : 'Error al guardar'
      return false
    } finally {
      saving.value = false
    }
  }

  async function remove(id: number): Promise<void> {
    await api.delete(id)
    items.value = items.value.filter((x) => x.id !== id)
  }

  return { items, saving, modalError, save, remove }
}
