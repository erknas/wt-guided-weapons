<script setup>
import { WEAPONS_CATEGORIES } from '../consts/weaponsCategories';

defineProps({
	modelValue: String
})

defineEmits(['update:modelValue'])

const categories = WEAPONS_CATEGORIES
</script>

<template>
	<div class="selector-container">
		<select :value="modelValue" @change="$emit('update:modelValue', $event.target.value)" class="category-selector">
			<option v-for="category in categories" :value="category.value" :key="category.value" :class="`option-${category.value}`">
				{{ category.label }}
			</option>
		</select>
	</div>
</template>

<style scoped>
.selector-container {
	position: fixed;
	top: 20px;
	left: 0px;
	z-index: 1000;
}

.category-selector {
	min-width: 170px;
	padding: 8px;
	font-size: 15px;
}
</style>