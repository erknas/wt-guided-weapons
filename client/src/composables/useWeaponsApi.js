import { ref } from "vue";

export function useWeaponsApi() {
  const weapons = ref([]);
  const loading = ref(false);
  const error = ref(null);

  const fetchWeaponsByCategory = async (category) => {
    try {
      loading.value = true;
      const response = await fetch(
        `/api/weapons/${encodeURIComponent(category)}`
      );
      const data = await response.json();
      weapons.value = data.weapons;
    } catch (err) {
      error.value = err.message;
    } finally {
      loading.value = false;
    }
  };

  return {
    weapons,
    loading,
    error,
    fetchWeaponsByCategory,
  };
}
