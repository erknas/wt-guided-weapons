<script setup>
import { ref, watch } from 'vue';
import { useSearchApi } from '../composables/useWeaponSearcherApi';

const {
  query,
  loading,
  error,
  results,
  clearSearch
} = useSearchApi()
const isOpen = ref(false)

const handleInputFocus = () => {
  if (query.value && results.value.length > 0) {
    setTimeout(() => {
      isOpen.value = true
    }, 50)
  }
}

const handleClickOutside = () => {
  setTimeout(() => {
    isOpen.value = false
  }, 50)
}

watch(results, (newResults) => {
  const hasResults = Array.isArray(newResults) && newResults.length > 0
  const hasQuery = query.value?.trim?.() !== ''
  
  isOpen.value = hasResults && hasQuery
}, { 
  immediate: false
})
</script>

<template>
  <div class="search-input-container">
    <input
      v-model="query"
      type="text"
      placeholder="Search"
      class="search-input"
      @focus="handleInputFocus"
      @blur="handleClickOutside"
    >
    <button
      v-if="query"
      @click="clearSearch"
      class="clear-button"
      type="button"
    >
      ✕
    </button>
    <div v-if="error" class="error-message">
      {{ error }}
    </div>
    <div v-if="isOpen && results.length > 0" class="dropdown-container">
      <div class="search-results">
        <div
          v-for="weapon in results"
          :key="weapon.name"
          class="weapon-item"
        >
          <div class="weapon-info">
            <span class="weapon-name">{{ weapon.name }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.search-input-container {
  position: fixed;
  top: 1px;
  left: 50%;
  transform: translateX(-50%);
  width: 300px;
  z-index: 1000;
}

.search-input {
  position: relative;
  height: 40px;
  width: 100%;
  padding: 0 40px 0 15px;
  color: black; 
  background: #dbdbdb;
  border: 1px solid gray; 
  font-family: Arial, Helvetica, sans-serif;
  font-size: 15px;
  outline: none;
  box-sizing: border-box;
  transition: border-color 0.2s ease;
}

.search-input:focus {
  border-color: black;
}

.clear-button {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  font-size: 18px;
  color: #666;
  cursor: pointer;
  padding: 4px;
}

.clear-button:hover {
  color: black;
}

.clear-button:focus:not(:focus-visible) {
  outline: none;
}

.error-message {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  padding: 10px 15px;
  background: #ffe6e6;
  border: 1px solid #ff9999;
  border-top: none;
  border-radius: 0 0 8px 8px;
  color: #cc0000;
  font-size: 14px;
}

.dropdown-container {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: #dbdbdb;
  border: 1px solid gray;
  border-top: none;
  overflow: hidden;
  z-index: 1001;
  text-align: left;
}

.search-results {
  max-height: 400px;
  overflow-y: auto;
  overflow-x: hidden;
}

.search-results::-webkit-scrollbar {
  width: 6px;
}

.search-results::-webkit-scrollbar-track {
  background: gray;
}

.search-results::-webkit-scrollbar-thumb {
  background: black;
}

.weapon-item {
  padding: 12px 15px;
  cursor: pointer;
  transition: background-color 0.2s ease;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  align-items: center;
  border-bottom: 1px solid gray;
  height: 20px;
}

.weapon-item:hover {
  background-color: #f8f9fa;
  font-weight: bold;
}

.weapon-item:last-child {
  border-bottom: none;
}

.weapon-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  width: 100%;
}

.weapon-name {
  color: black; 
  font-size: 15px;
}

.weapon-category {
  color: #666;
  font-size: 12px;
  background: #e9ecef;
  padding: 2px 6px;
  border-radius: 3px;
  display: inline-block;
  width: fit-content;
}

@media (max-width: 480px) {
  .search-input-container {
    width: 90%;
    left: 5%;
    transform: none;
  }
  
  .search-results {
    max-height: 250px;
  }
  
  .weapon-item {
    padding: 10px 12px;
  }
}
</style>