import { computed, ref, watch } from "vue";

export function useSearchApi() {
  const query = ref("");
  const results = ref({});
  const loading = ref(false);
  const error = ref("");

  const searchAPI = async (searchQuery) => {
    const response = await fetch(
      `/api/weapons/search/${encodeURIComponent(searchQuery)}`
    );
    if (!response.ok) {
      throw new Error("Network error");
    }
    return await response.json();
  };

  const resultsArray = computed(() => {
    return Object.entries(results.value);
  });

  const debounce = (func, delay) => {
    let timeout;
    return (...args) => {
      clearTimeout(timeout);
      timeout = setTimeout(() => func.apply(null, args), delay);
    };
  };

  const search = async (searchQuery) => {
    if (!searchQuery.trim()) {
      results.value = {};
      return;
    }

    loading.value = true;
    error.value = "";

    try {
      const data = await searchAPI(searchQuery);
      results.value = data;
    } catch (err) {
      error.value = "Search error";
      console.error("Search error:", err);
      results.value = {};
    } finally {
      loading.value = false;
    }
  };

  const debouncedSearch = debounce(search, 300);

  watch(query, (newQuery) => {
    debouncedSearch(newQuery);
  });

  const clearSearch = () => {
    query.value = "";
    results.value = {};
    error.value = "";
  };

  return {
    query,
    loading,
    error,
    resultsArray,

    search,
    clearSearch,
    debouncedSearch,
  };
}
