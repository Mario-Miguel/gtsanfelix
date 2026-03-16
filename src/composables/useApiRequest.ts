import { ref, onMounted, type Ref } from 'vue'

export function useApiRequest<T>(
  fn: () => Promise<T>,
  options: { immediate?: boolean; initialData?: T } = {},
): { data: Ref<T | undefined>; loading: Ref<boolean>; error: Ref<string | null>; execute: () => Promise<void> } {
  const { immediate = true, initialData = undefined } = options

  const data = ref<T | undefined>(initialData) as Ref<T | undefined>
  const loading = ref(immediate)
  const error = ref<string | null>(null)

  async function execute() {
    loading.value = true
    error.value = null
    try {
      data.value = await fn()
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Error desconocido'
    } finally {
      loading.value = false
    }
  }

  if (immediate) {
    onMounted(execute)
  }

  return { data, loading, error, execute }
}
