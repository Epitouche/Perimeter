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
    console.log("Services loaded:", services.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading services:", error);
  }
};

const filteredApps = computed(() => {
  console.log("Search Query:", searchQuery.value);
  console.log("Services:", services.value);
  return services.value.filter((app) =>
    app.name.toLowerCase().includes(searchQuery.value.toLowerCase()),
  );
});
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-5 w-full">
    <h1>My Services</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div v-else
      class="flex flex-col justify-center items-center gap-10 w-[90%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
      tabindex="0">
      <div class="self-start w-1/4 px-5 pt-1">
        <SearchBar v-model:search-query="searchQuery" tabindex="0" />
      </div>
      <div class="w-[95%] overflow-y-scroll max-h-[64vh]">
        <ServiceList styling="card" :apps="filteredApps" />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
