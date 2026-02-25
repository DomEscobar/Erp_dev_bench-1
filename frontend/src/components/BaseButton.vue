<template>
  <button class="base-button">
    <slot>Submit</slot>
  </button>
</template>

<style scoped>
.base-button {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.875rem;
  background: #007bff;
  color: white;
}

.base-button:hover {
  background: #0056b3;
}

.base-button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}
</style>
