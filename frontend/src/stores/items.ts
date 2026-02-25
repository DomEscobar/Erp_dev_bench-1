import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Item } from '@/types/Item'

export const useItemsStore = defineStore('items', () => {
  const items = ref<Item[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

  async function fetchItems() {
    loading.value = true
    error.value = null
    try {
      const response = await fetch(`${API_URL}/items`)
      if (!response.ok) throw new Error('Failed to fetch items')
      items.value = await response.json()
    } catch (e: any) {
      error.value = e.message
    } finally {
      loading.value = false
    }
  }

  async function fetchItem(id: number) {
    const response = await fetch(`${API_URL}/items/${id}`)
    if (!response.ok) throw new Error('Failed to fetch item')
    return await response.json() as Item
  }

  async function createItem(item: Omit<Item, 'id' | 'createdAt' | 'updatedAt' | 'deletedAt'>) {
    const response = await fetch(`${API_URL}/items`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(item)
    })
    if (!response.ok) throw new Error('Failed to create item')
    const created = await response.json() as Item
    items.value.push(created)
    return created
  }

  async function updateItem(id: number, updates: Partial<Pick<Item, 'name' | 'description' | 'price' | 'sku' | 'categoryId'>>) {
    const response = await fetch(`${API_URL}/items/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updates)
    })
    if (!response.ok) throw new Error('Failed to update item')
    const updated = await response.json() as Item
    const index = items.value.findIndex(i => i.id === id)
    if (index !== -1) items.value[index] = updated
    return updated
  }

  async function deleteItem(id: number) {
    const response = await fetch(`${API_URL}/items/${id}`, {
      method: 'DELETE'
    })
    if (!response.ok) throw new Error('Failed to delete item')
    items.value = items.value.filter(i => i.id !== id)
  }

  return {
    items,
    loading,
    error,
    fetchItems,
    fetchItem,
    createItem,
    updateItem,
    deleteItem
  }
})
