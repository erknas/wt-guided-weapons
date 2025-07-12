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
	top: 1px;
	left: 1px;
	z-index: 100;
}

.category-selector {
	min-width: 170px;
	padding: 8px;
	font-size: 15px;
	outline: none !important;
	background: #dbdbdb;
	color: black;
	font-size: small;
}

</style>
