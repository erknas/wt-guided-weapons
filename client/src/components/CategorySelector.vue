<script setup>
import { WEAPONS_CATEGORIES } from '../consts/weaponsCategories';

defineProps({
	modelValue: String
})

defineEmits(['update:modelValue'])

const categories = WEAPONS_CATEGORIES
</script>

<template>
	<select :value="modelValue" @change="$emit('update:modelValue', $event.target.value)">
		<option v-for="category in categories" :value="category.value" :key="category.value">
			{{ category.label }}
		</option>
	</select>
</template>