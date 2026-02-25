<template>
  <div class="items-table">
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Description</th>
          <th>Price</th>
          <th>SKU</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>{{ item.id }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
          <td>${{ item.price.toFixed(2) }}</td>
          <td>{{ item.sku }}</td>
          <td>
            <button @click="onEdit(item)" class="btn btn-primary">Edit</button>
            <button @click="onDelete(item.id)" class="btn btn-danger">Delete</button>
          </td>
        </tr>
        <tr v-if="items.length === 0">
          <td colspan="6" class="text-center">No items found</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import type { Item } from '@/types/Item'

defineProps<{
  items: Item[]
}>()

const emit = defineEmits<{
  edit: [item: Item]
  delete: [id: number]
}>()

function onEdit(item: Item) {
  emit('edit', item)
}

function onDelete(id: number) {
  if (confirm('Are you sure you want to delete this item?')) {
    emit('delete', id)
  }
}
</script>

<style scoped>
.items-table {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
}

th, td {
  padding: 0.75rem;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

th {
  background-color: #f5f5f5;
  font-weight: 600;
}

.text-center {
  text-align: center;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.875rem;
  margin-right: 0.5rem;
}

.btn:last-child {
  margin-right: 0;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-danger {
  background-color: #dc3545;
  color: white;
}

.btn-danger:hover {
  background-color: #a71d2a;
}
</style>
