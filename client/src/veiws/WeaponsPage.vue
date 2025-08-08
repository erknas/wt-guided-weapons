<script setup>
import { ref, watch } from 'vue'
import { useWeaponsApi } from '../composables/useWeaponsApi';
import CategorySelector from '../components/CategorySelector.vue';
import WeaponsLoader from '../components/WeaponsLoader.vue';
import WeaponSearcher from '../components/WeaponSearcher.vue';

const selectedCategory = ref('aam-ir-rear-aspect')
const selectedWeaponName = ref(null)
const { weapons, loading, error, fetchWeaponsByCategory } = useWeaponsApi()

const handleWeaponSelected = (selectedWeapon) => {
	selectedCategory.value = selectedWeapon.category
	selectedWeaponName.value = selectedWeapon.name
}

const handleCategoryChange = (newCategory) => {
	selectedCategory.value = newCategory
	selectedWeaponName.value = null
}

watch(selectedCategory, async (newCategory) => {
  if (newCategory) {
    await fetchWeaponsByCategory(newCategory)
  }
}, { immediate: true })
</script>

<template>
	<div>
		<div v-if="loading">Loading data...</div>
		<div v-else-if="error" class="error-message">{{ error }}</div>
		<WeaponSearcher 
			@weapon-selected="handleWeaponSelected" 
		/>

		<CategorySelector 
			:model-value="selectedCategory"
			@update:model-value="handleCategoryChange"
		/>

		<WeaponsLoader
			v-if="!loading && !error"
			:weapons="weapons"
			:category="selectedCategory"
			:highlighted-weapon="selectedWeaponName"
		/>
	</div>
</template>
