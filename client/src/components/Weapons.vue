<script setup>
import { onMounted, ref } from "vue";
import { useWeaponsApi } from "../composables/useWeaponsApi"

const { weapons, loading, error, fetchWeaponsByCategory } = useWeaponsApi()

const selectedCategory = ref("aam-ir-rear-aspect")

const handleCategoryChange = () => {
  fetchWeaponsByCategory(selectedCategory.value)
}

onMounted(() => {
  fetchWeaponsByCategory(selectedCategory.value)
})
</script>

<template>
  <div>
    <h1>Guided Weapons</h1>

    <div>
      <select 
        v-model="selectedCategory" 
        @change="handleCategoryChange"
      >
        <option value="aam-ir-rear-aspect">AAM (IR rear-aspect)</option>
        <option value="aam-ir-all-aspect">AAM (IR all-aspect)</option>
        <option value="aam-ir-heli">AAM (IR heli)</option>
        <option value="aam-sarh">AAM (SARH)</option>
        <option value="aam-arh">AAM (ARH)</option>
        <option value="aam-manual">AAM (Manual)</option>
        <option value="agm-tv">AGM (TV)</option>
        <option value="agm-ir">AGM (IR)</option>
        <option value="agm-gnss">AGM (GNSS)</option>
        <option value="agm-salh">AGM (SALH)</option>
        <option value="agm-losbr">AGM (LOSBR)</option>
        <option value="agm-saclos">AGM (SACLOS)</option>
        <option value="agm-mclos">AGM (MCLOS)</option>
        <option value="gbu-tv">GBU (TV)</option>
        <option value="gbu-ir">GBU (IR)</option>
        <option value="gbu-gnss">GBU (GNSS)</option>
        <option value="gbu-salh">GBU (SALH)</option>
        <option value="sam-arh">SAM (ARH)</option>
        <option value="sam-ir">SAM (IR)</option>
        <option value="sam-ir-optical">SAM (IR/OPTICAL)</option>
        <option value="sam-losbr">SAM (LOSBR)</option>
        <option value="sam-saclos">SAM (SACLOS)</option>
        <option value="atgm-ir">ATGM (IR)</option>
        <option value="atgm-losbr">ATGM (LOSBR)</option>
        <option value="atgm-saclos">ATGM (SACLOS)</option>
        <option value="atgm-mclos">ATGM (MCLOS)</option>
        <option value="ashm-arh">AShM (ARH)</option>
        <option value="ashm-saclos">AShM (SACLOS)</option>
        <option value="sam-ir-naval">SAM (IR) (naval)</option>
        <option value="sam-saclos-naval">SAM (SACLOS) (naval)</option>
      </select>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-if="error" class="error">{{ error }}</div>
    
    <div v-if="weapons.length">
      <ul>
        <li v-for="weapon in weapons" :key="weapon.id">
          {{ weapon.name }}
        </li>
      </ul>
    </div>
    <div v-else>Weapons not found</div>
  </div>
</template>