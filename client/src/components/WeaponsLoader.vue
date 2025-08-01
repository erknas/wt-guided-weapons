<script setup>
import { defineAsyncComponent, computed, ref } from 'vue'

const props = defineProps({
  weapons: {
    type: Array,
    required: true
  },
  category: String
})

const error = ref(null)

const getTableName = (category) => {
  error.value = null
  return category
    .split('-')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join('')
}

const currentTable = computed(() => {
  const componentName = getTableName(props.category)
  return defineAsyncComponent(() => 
    import(`./WeaponsTables/${componentName}.vue`)
    .catch((err) => {
      error.value = 'Category not found'
      console.error('failed to load component:', err)
      return
    })
  )
})
</script>

<template>
  <div v-if="error" class="error-message">{{ error }}</div>
  <component :is="currentTable" :weapons="weapons" v-else />
</template>

<style scoped>
.error-message {
  color: #ff4444;
  padding: 1rem;
  background: #ffeeee;
  border: 1px solid #ffcccc;
  border-radius: 4px;
  margin: 1rem 0;
}
</style>