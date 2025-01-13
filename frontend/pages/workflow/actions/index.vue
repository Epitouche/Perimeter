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
  <div class="py-20">
    <div class="px-20">
      <BackButton link="/workflow" :is-white="false" />
    </div>
    <h1
      class="flex justify-center w-full text-8xl font-custom_weight_title pb-20"
    >
      Add an action
    </h1>
    <UContainer
      :ui="{ base: 'mx-auto' }"
      class="flex flex-col justify-center items-center gap-16 w-full h-full !p-0"
    >
      <SearchBar v-model:search-query="searchQuery" class="!w-1/3" />
      <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>
      <div v-else-if="errorMessage">Error: {{ errorMessage }}</div>
      <div
        v-else-if="filteredServices.length"
        class="flex flex-row justify-evenly items-center w-full"
      >
        <ServiceCardContainer type="actions" :services="filteredServices" />
      </div>
    </UContainer>
  </div>
</template>

<style scoped></style>
