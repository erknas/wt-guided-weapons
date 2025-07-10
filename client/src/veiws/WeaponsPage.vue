<script setup>
import { ref, onMounted } from 'vue'
import CategorySelector from '../components/CategorySelector.vue';
import WeaponsLoader from '../components/WeaponsLoader.vue';
import { useWeaponsApi } from '../composables/useWeaponsApi';

const selectedCategory = ref('aam-ir-rear-aspect')
const { weapons, loading, error, fetchWeaponsByCategory } = useWeaponsApi()

const handleCategoryChange = async (category) => {
	selectedCategory.value = category
	await fetchWeaponsByCategory(category)
}

onMounted(() => fetchWeaponsByCategory(selectedCategory.value))
</script>

<template>
	<div>
		<div v-if="loading">Loading data...</div>
		<div v-else-if="error" class="error">{{ error }}</div>

		<CategorySelector 
			v-model="selectedCategory"
			@update:modelValue="handleCategoryChange"
		/>

		<WeaponsLoader
			v-if="!loading && !error"
			:weapons="weapons"
			:category="selectedCategory"
		/>
	</div>
</template>