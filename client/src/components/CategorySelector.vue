<script setup>
import { WEAPONS_CATEGORIES } from '../consts/weaponsCategories';
import { computed, ref, onMounted, onUnmounted } from 'vue'

const categories = WEAPONS_CATEGORIES
const emit = defineEmits(['update:modelValue'])
const props = defineProps ({
	modelValue: String
})
const isOpen = ref(false)
const dropdownRef = ref(null)

const selectedCategory = computed(() => {
	return categories.find(category => category.value === props.modelValue) || categories[0]
})

const toggleDropdown = () => {
	isOpen.value = !isOpen.value
}

const selectCategory = (category) => {
	emit('update:modelValue', category.value)
	isOpen.value = false
}

const handleClickOutside = (event) => {
	if (dropdownRef.value && !dropdownRef.value.contains(event.target)) {
		isOpen.value = false
	}
}

onMounted(() => {
	document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
	document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
	<div class="selector-container" ref="dropdownRef">
		<div 
			class="category-selector" 
			@click="toggleDropdown"
			:class="{ 'open': isOpen }"
		>
			<span class="selected-text">{{ selectedCategory.label }}</span>
			<span class="arrow" :class="{ 'rotated': isOpen }">▼</span>
		</div>
			<div v-if="isOpen" class="dropdown-list">
				<div
					v-for="category in categories"
					:key="category.value"
					:class="[
						'dropdown-item',
						`option-${category.value}`,
						{ 'selected': category.value === modelValue }
					]"
					@click="selectCategory(category)"
				>
					{{ category.label }}
				</div>
			</div>
	</div>
</template>

<style scoped>
.selector-container {
	position: fixed;
	top: 1px;
	left: 1px;
	z-index: 100;
	text-align: left;
	min-width: 200px;
  	font-family: Arial, Helvetica, sans-serif;
  	font-size: 15px;
	color: black;
}

.category-selector {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 8px 10px;
	background: #dbdbdb;
	cursor: pointer;
	border: 1px solid gray;
	user-select: none;
}

.selected-text {
	font-size: small;
	flex: 1;
}

.arrow {
	margin-left: 8px;
	transition: transform 0.2s ease;
	font-size: 12px;
}

.arrow.rotated {
	transform: rotate(180deg);
}

.dropdown-list {
	position: absolute;
	top: 100%;
	left: 0;
	right: 0;
	background: #dbdbdb;
	border: 1px solid gray;
	border-top: none;
	max-height: 400px;
	overflow-y: auto;
	z-index: 101;
}

.dropdown-list::-webkit-scrollbar {
	width: 6px;
}

.dropdown-list::-webkit-scrollbar-track {
	background: gray;
}

.dropdown-list::-webkit-scrollbar-thumb {
	background: black;
}

.dropdown-item {
	border-bottom: 1px solid gray;
	padding: 8px 10px;
	cursor: pointer;
	font-size: small;
}

.dropdown-item:last-child {
	border-bottom: none;
}

.dropdown-item:hover {
  	background-color: #f8f9fa;
	font-weight: bold;
}

.dropdown-item.selected {
	background: #b8b8b8;
	font-weight: 500;
}
</style>