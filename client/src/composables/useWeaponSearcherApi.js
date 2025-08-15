import { computed, ref, watch } from "vue";

export function useSearchApi() {
  const query = ref("");
  const results = ref([]);
  const loading = ref(false);
  const error = ref("");
  const lastSearchQuery = ref("");

  const searchAPI = async (searchQuery) => {
    const response = await fetch(
      `/api/weapons/search/${encodeURIComponent(searchQuery)}`
    );
    if (!response.ok) {
      const errorData = await response.json();
      const errorMessage = errorData.message;
      throw new Error(errorMessage);
    }
    return await response.json();
  };

  const debounce = (func, delay) => {
    let timeout;
    return (...args) => {
      clearTimeout(timeout);
      timeout = setTimeout(() => func.apply(null, args), delay);
    };
  };

  const search = async (searchQuery) => {
    const trimmedQuery = searchQuery.trim();

    if (!trimmedQuery) {
      results.value = [];
      lastSearchQuery.value = "";
      return;
    }

    if (trimmedQuery === lastSearchQuery.value) {
      return;
    }

    loading.value = true;
    error.value = "";

    try {
      const data = await searchAPI(trimmedQuery);
      results.value = data.results;
      lastSearchQuery.value = trimmedQuery;
    } catch (err) {
      error.value = err.message;
      results.value = [];
    } finally {
      loading.value = false;
    }
  };

  const debouncedSearch = debounce(search, 500);

  watch(query, (newQuery) => {
    debouncedSearch(newQuery);
  });

  const clearSearch = () => {
    query.value = "";
    results.value = [];
    error.value = "";
  };

  return {
    query,
    loading,
    error,
    results,

    search,
    clearSearch,
    debouncedSearch,
  };
}
