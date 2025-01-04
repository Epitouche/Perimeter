<script lang="ts" setup>
import type { Area } from "@/interfaces/areas";
import { handleErrorStatus } from "../utils/handleErrorStatus.js";

definePageMeta({
  middleware: "auth",
});

const token = useCookie("token");
const errorMessage = ref<string | null>(null);
const isLoading = ref(true);

const areas = ref<Area[]>([]);
const filteredAreas = ref<Area[]>([]);

const searchQuery = ref<string>("");

const fetchAreas = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<Area[]>("/api/myareas", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    areas.value = result;
    filteredAreas.value = result;
    console.log("areas: ", areas.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    filteredAreas.value = [];
  } finally {
    isLoading.value = false;
  }
};

watch(searchQuery, (newQuery) => {
  const lowerQuery = newQuery.toLowerCase();
  filteredAreas.value = areas.value.filter(
    (area) =>
      area.action.service.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.service.name.toLowerCase().includes(lowerQuery) ||
      area.action.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.name.toLowerCase().includes(lowerQuery),
  );
});

onMounted(() => {
  fetchAreas();
});
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-10 w-full">
    <h1 class="text-custom_size_title font-custom_weight_title">My Areas</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div
      v-else
      class="flex flex-col justify-center items-start gap-12 w-[90%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
    >
      <div class="flex flex-row justify-between items-center w-full">
        <SearchBar v-model:search-query="searchQuery" class="!w-1/4" />
        <div>filter</div>
      </div>
      <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>
      <div v-else-if="errorMessage">Error: {{ errorMessage }}</div>
      <div
        v-else-if="filteredAreas.length"
        class="w-full overflow-y-scroll max-h-[64vh]"
      >
        <AreaCardContainer :areas="filteredAreas" />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
