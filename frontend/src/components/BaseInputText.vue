<script setup lang="ts">
defineProps({
  type: {
    type: String,
    default: 'text',
    // Only allow types that essentially just render text boxes
    validator: (value: string) =>
      ['email', 'number', 'password', 'search', 'tel', 'text', 'url'].includes(value)
  }
})

const model = defineModel()
</script>

<template>
  <input v-model="model" :type="type" class="base-input" />
</template>

<style scoped>
.base-input {
  display: block;
  width: 100%;
  padding: 0.5rem 0.75rem;
  margin-bottom: 1rem;
  line-height: 1;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.base-input:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}
</style>
