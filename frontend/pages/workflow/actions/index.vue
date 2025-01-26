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

/**
 * @description Fetches all the services from the backend
 */
const fetchServices = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<ServiceInfo[]>("/api/workflow/services", {
      method: "GET",
    });
    services.value = result;
    filteredServices.value = result;
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

/**
 * @description Watches the search query and filters the services based on the query
 * @param searchQuery - The search query
 */
watch(searchQuery, (newQuery) => {
  filteredServices.value = services.value.filter((service) =>
    service.name.toLowerCase().includes(newQuery.toLowerCase())
  );
});

/**
 * @description Fetches all the services when the component is mounted
 */
onMounted(() => {
  fetchServices();
});
</script>

<template>
  <div class="py-10 max-sm:py-3">
    <div class="px-10 max-sm:px-2 max-sm:pb-5">
      <BackButton link="/workflow" :is-white="false" />
    </div>
    <div
      class="flex flex-col justify-between items-center gap-10 max-sm:gap-5 w-full"
    >
      <h1>Add an action</h1>
      <UContainer
        :ui="{ base: 'mx-auto' }"
        class="flex flex-col justify-center items-center gap-16 max-sm:gap-10 w-[80%] max-lg:w-[85%] max-md:w-[90%] max-sm:w-[95%] h-full !p-0"
      >
        <div class="min-w-1/3 max-w-[95%]">
          <SearchBar v-model:search-query="searchQuery" tabindex="0" />
        </div>
        <div
          v-if="isLoading"
          class="flex justify-center items-center w-full h-full"
        >
          <h3>Loading...</h3>
        </div>
        <div v-else-if="errorMessage">
          <UAlert
            color="red"
            variant="solid"
            title="ERROR"
            :description="errorMessage"
            class="justify-center items-center gap-5 w-[15%]"
          />
        </div>
        <div
          v-else-if="filteredServices.length"
          class="flex flex-row justify-evenly items-center w-full"
        >
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
