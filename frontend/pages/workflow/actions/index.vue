<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";

definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);

const services = ref<ServiceInfo[]>([]);
const filteredServices = ref<ServiceInfo[]>([]);

const searchQuery = ref<string>("");

const fetchServices = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<ServiceInfo[]>("/api/workflow/services", {
      method: "GET",
    });
    services.value = result;
    filteredServices.value = result;
    //console.log("services", services.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    filteredServices.value = [];
  } finally {
    isLoading.value = false;
  }
};

watch(searchQuery, (newQuery) => {
  filteredServices.value = services.value.filter((service) =>
    service.name.toLowerCase().includes(newQuery.toLowerCase()),
  );
});

onMounted(() => {
  fetchServices();
});
</script>

<template>
  <div class="py-10">
    <div class="px-10">
      <BackButton link="/workflow" :is-white="false" />
    </div>
    <div class="flex flex-col justify-between items-center gap-10 w-full">
      <h1>Add an action</h1>
      <UContainer
:ui="{ base: 'mx-auto' }"
        class="flex flex-col justify-center items-center gap-16 w-[80%] h-full !p-0">
        <div class="w-1/3">
          <SearchBar v-model:search-query="searchQuery" tabindex="0" />
        </div>
        <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>
        <div v-else-if="errorMessage">Error: {{ errorMessage }}</div>
        <div v-else-if="filteredServices.length" class="flex flex-row justify-evenly items-center w-full">
          <ServiceCardContainer type="actions" :services="filteredServices" />
        </div>
      </UContainer>
    </div>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
