<script setup lang="ts">
import SearchBar from "@/components/SearchBar.vue";
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import { fetchServices } from "~/utils/fetchServices";

definePageMeta({
  middleware: "auth",
});

const errorMessage = ref<string | null>(null);
const services = ref<ServiceInfo[]>([]);
const searchQuery = ref("");

onMounted(() => {
  loadServices();
});

const loadServices = async () => {
  try {
    errorMessage.value = null;
    services.value = await fetchServices();
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading services:", error);
  }
};

const filteredApps = computed(() => {
  return services.value.filter((app) => app.name.includes(searchQuery.value));
});
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-10 w-full">
    <h1 class="text-custom_size_title font-custom_weight_title">My Services</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div
      v-else
      class="flex flex-col justify-center items-start gap-10 w-[90%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
    >
      <div class="flex flex-row justify-between items-center w-full px-5 pt-1">
        <SearchBar
          v-model:search-query="searchQuery"
          class="!w-1/4"
          tabindex="0"
        />
      </div>

      <UContainer
        class="flex flex-wrap gap-5 justify-center p-4 bg-white rounded-lg w-full mx-auto"
      >
        <ServiceList :apps="filteredApps" />
      </UContainer>
    </div>
  </div>
</template>

<style scoped></style>
