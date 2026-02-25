<template>
  <form @submit.prevent="onSubmit" class="item-form">
    <div class="form-group">
      <label for="name">Name *</label>
      <BaseInputText
        id="name"
        v-model="formData.name"
        :disabled="isEditing"
        required
      />
    </div>

    <div class="form-group">
      <label for="description">Description</label>
      <textarea
        id="description"
        v-model="formData.description"
        rows="3"
        class="textarea"
      ></textarea>
    </div>

    <div class="form-group">
      <label for="price">Price *</label>
      <input
        type="number"
        id="price"
        v-model.number="formData.price"
        step="0.01"
        min="0"
        required
        class="input"
      />
    </div>

    <div class="form-group">
      <label for="sku">SKU *</label>
      <BaseInputText
        id="sku"
        v-model="formData.sku"
        required
      />
    </div>

    <div class="form-group">
      <label for="categoryId">Category ID</label>
      <input
        type="number"
        id="categoryId"
        v-model.number="formData.categoryId"
        min="0"
        class="input"
      />
    </div>

    <div class="form-actions">
      <button type="submit" class="btn btn-primary">
        {{ isEditing ? 'Update' : 'Create' }}
      </button>
      <button type="button" class="btn btn-secondary" @click="onCancel">
        Cancel
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import type { Item } from '@/types/Item'
import BaseInputText from '@/components/BaseInputText.vue'

const props = defineProps<{
  item?: Item | null
}>()

const emit = defineEmits<{
  submit: [item: Omit<Item, 'id' | 'createdAt' | 'updatedAt' | 'deletedAt'>]
  cancel: []
}>()

const formData = ref({
  name: '',
  description: '',
  price: 0,
  sku: '',
  categoryId: 0
})

const isEditing = computed(() => props.item !== null && props.item !== undefined)

watch(() => props.item, (newItem) => {
  if (newItem) {
    formData.value = {
      name: newItem.name,
      description: newItem.description,
      price: newItem.price,
      sku: newItem.sku,
      categoryId: newItem.categoryId
    }
  } else {
    resetForm()
  }
}, { immediate: true })

function resetForm() {
  formData.value = {
    name: '',
    description: '',
    price: 0,
    sku: '',
    categoryId: 0
  }
}

function onSubmit() {
  emit('submit', { ...formData.value })
}

function onCancel() {
  emit('cancel')
}
</script>

<style scoped>
.item-form {
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

label {
  font-weight: 600;
  font-size: 0.9rem;
}

.input,
.textarea {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.input:focus,
.textarea:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.form-actions {
  display: flex;
  gap: 0.5rem;
  margin-top: 1rem;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.875rem;
}

.btn-primary {
  background-color: #007bff;
  color: white;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-secondary {
  background-color: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background-color: #545b62;
}
</style>
