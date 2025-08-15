<script setup>
import { defineAsyncComponent, computed, ref, nextTick, watch } from 'vue'

const props = defineProps({
  weapons: {
    type: Array,
    required: true
  },
  category: String,
  highlightedWeapon: String
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

const scrollToWeapon = async () => {
  if (!props.highlightedWeapon) {
    return
  }

  await nextTick()

  const weaponCell = document.querySelector(`[data-weapon-name="${props.highlightedWeapon}"]`)

  if (weaponCell) {
    weaponCell.classList.add('highlighted-weapon')

    weaponCell.scrollIntoView({
      behavior: 'smooth',
      block: 'center',
      inline: 'center'
    })
    
    setTimeout(() => {
      weaponCell.classList.remove('highlighted-weapon')
    }, 3000)
  }
}

watch([() => props.highlightedWeapon, () => props.category], async () => {
  if (props.highlightedWeapon) {
    setTimeout(scrollToWeapon, 100)
  }
}, { immediate: true })
</script>

<template>
  <div v-if="error" class="error-message">{{ error }}</div>
  <component 
    v-else
    :is="currentTable" 
    :weapons="weapons" 
    :highlighted-weapon="highlightedWeapon" 
  />
</template>
