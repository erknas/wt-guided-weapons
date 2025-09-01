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
      if (!response.ok) {
        throw new Error("Application error");
      }
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

export function useUpdateWeaponsApi() {
  const loading = ref(false);
  const error = ref(null);

  const updateAPI = async () => {
    const response = await fetch(`/api/update`, {
      method: "PUT",
    });
    if (!response.ok) {
      const errorData = await response.json();
      const errorMessage = errorData.message;
      throw new Error(errorMessage);
    }

    return response;
  };

  const update = async () => {
    loading.value = true;
    error.value = "";

    try {
      await updateAPI();
    } catch (err) {
      error.value = err.message;
    } finally {
      loading.value = false;
    }
  };

  return {
    loading,
    error,
    update,
  };
}

export function useGetVersionApi() {
  const versionInfo = ref("");
  const error = ref(null);

  const versionAPI = async () => {
    const response = await fetch(`/api/version`);
    if (!response.ok) {
      const errorData = await response.json();
      const errorMessage = errorData.message;
      throw new Error(errorMessage);
    }
    return await response.json();
  };

  const getVersion = async () => {
    error.value = "";

    try {
      const data = await versionAPI();
      versionInfo.value = data.version;
    } catch (err) {
      error.value = err.message;
    }
  };

  return {
    versionInfo,
    error,
    getVersion,
  };
}
