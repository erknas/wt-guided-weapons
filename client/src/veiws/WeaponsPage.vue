<script setup>
import { ref, watch } from 'vue'
import { useWeaponsApi } from '../composables/useWeaponsApi';
import CategorySelector from '../components/CategorySelector.vue';
import WeaponsLoader from '../components/WeaponsLoader.vue';
import WeaponSearcher from '../components/WeaponSearcher.vue';

const selectedCategory = ref('aam-ir-rear-aspect')
const { weapons, loading, error, fetchWeaponsByCategory } = useWeaponsApi()

watch(selectedCategory, async (newCategory) => {
  if (newCategory) {
    await fetchWeaponsByCategory(newCategory)
  }
}, { immediate: true })

</script>

<template>
	<div>
		<div v-if="loading">Loading data...</div>
		<div v-else-if="error" class="error">{{ error }}</div>

		<CategorySelector 
			v-model="selectedCategory"
		/>

		<WeaponsLoader
			v-if="!loading && !error"
			:weapons="weapons"
			:category="selectedCategory"
		/>

		<WeaponSearcher />
	</div>
</template>
