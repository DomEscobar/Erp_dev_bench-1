<template>
  <div class="items-page">
    <h1>Items Management</h1>

    <div v-if="store.loading" class="loading">Loading items...</div>
    <div v-else-if="store.error" class="error">{{ store.error }}</div>
    <div v-else>
      <ItemsTable
        :items="store.items"
        @edit="onEdit"
        @delete="onDelete"
      />

      <div class="add-item-section">
        <h2>{{ isModalOpen ? 'Edit Item' : 'Add New Item' }}</h2>
        <ItemForm
          :item="isModalOpen ? selectedItem : null"
          @submit="onFormSubmit"
          @cancel="onCancel"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useItemsStore } from '@/stores/items'
import ItemsTable from '@/components/ItemsTable.vue'
import ItemForm from '@/components/ItemForm.vue'
import type { Item } from '@/types/Item'

const store = useItemsStore()
store.fetchItems()

const isModalOpen = ref(false)
const selectedItem = ref<Item | null>(null)

function onEdit(item: Item) {
  selectedItem.value = item
  isModalOpen.value = true
}

function onDelete(id: number) {
  store.deleteItem(id)
}

function onFormSubmit(data: Omit<Item, 'id' | 'createdAt' | 'updatedAt' | 'deletedAt'>) {
  if (isModalOpen.value && selectedItem.value) {
    store.updateItem(selectedItem.value.id, data)
    isModalOpen.value = false
    selectedItem.value = null
  } else {
    store.createItem(data)
  }
}

function onCancel() {
  isModalOpen.value = false
  selectedItem.value = null
}
</script>

<style scoped>
.items-page {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

h1 {
  margin-bottom: 1.5rem;
}

h2 {
  margin-bottom: 1rem;
  font-size: 1.25rem;
}

.add-item-section {
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid #ddd;
}

.loading,
.error {
  padding: 1rem;
  text-align: center;
  font-size: 1.1rem;
}

.error {
  color: #e53e3e;
}
</style>
