<script setup>
import { onMounted, ref } from 'vue';
import { useGetVersionApi, useUpdateWeaponsApi } from '../composables/useWeaponsApi';

const { loading: updateLoading, error: updateError, update } = useUpdateWeaponsApi();
const { versionInfo, error: versionError, loading: versionLoading, getVersion } = useGetVersionApi();

const showDropdown = ref(false);

onMounted(() => {
  getVersion();
});

const handleUpdateClick = async () => {
  await update();
  if (!updateError.value) {
    await getVersion();
  }
};
</script>

<template>
  <div class="button-wrapper">
    <div 
      class="button-container"
      @mouseenter="showDropdown = true"
      @mouseleave="showDropdown = false"
    >
      <button 
        class="btn"
        @click="handleUpdateClick"
        :disabled="updateLoading || versionLoading"
      >
        <div class="version-info">
          <span v-if="versionLoading">Loading...</span>
          <span v-else-if="versionError">Error: {{ versionError }}</span>
          <span v-else>{{ versionInfo || 'Update' }}</span>
        </div>
        <div v-if="updateLoading" class="loading-spinner">⟳</div>
      </button>
      
      <div 
        v-show="showDropdown" 
        class="dropdown-tooltip"
      >
        <div class="tooltip-content">
          <p>This is version when last changes were made in stats.</p>
          <p><strong>Click to update:</strong> You can manually fetch latest weapons stats. Application checks changes every 30 minutes.</p>
          <div v-if="updateError" class="error-message">
            Update Error: {{ updateError }}
          </div>
        </div>
        <div class="tooltip-arrow"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.button-wrapper {
  position: fixed;
  top: 1px;
  right: 5px;
  z-index: 1000;
}

.button-container {
  position: relative;
  display: inline-block;
}

button {
  background-color: #dbdbdb;
  color: black;
  padding: 8px 12px;
  border: 1px solid gray; 
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 80px;
}

button:hover {
  border-color: black;
}

button:disabled {
  background-color: #a8a8a8;
  cursor: not-allowed;
}

.version-info {
  font-size: 12px;
  flex: 1;
}

.loading-spinner {
  animation: spin 1s linear infinite;
  font-size: 14px;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.dropdown-tooltip {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 8px;
  background: white;
  border: 1px solid #ddd;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 1001;
  min-width: 250px;
  animation: fadeIn 0.2s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.tooltip-content {
  padding: 12px 16px;
  color: #333;
  font-size: 13px;
  line-height: 1.4;
}

.tooltip-content h4 {
  margin: 0 0 8px 0;
  font-size: 14px;
  font-weight: 600;
  color: #dc3545;
}

.tooltip-content p {
  margin: 0 0 6px 0;
  text-align: left;
}

.tooltip-content p:last-child {
  margin-bottom: 0;
}

.error-message {
  margin-top: 8px;
  padding: 6px 8px;
  background-color: #f8d7da;
  color: #721c24;
  border-radius: 4px;
  font-size: 12px;
}

.tooltip-arrow {
  position: absolute;
  top: -6px;
  right: 16px;
  width: 0;
  height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-bottom: 6px solid white;
}

.tooltip-arrow::before {
  content: '';
  position: absolute;
  top: -1px;
  left: -6px;
  width: 0;
  height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-bottom: 6px solid #ddd;
}
</style>