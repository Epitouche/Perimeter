<script setup lang="ts">
import SearchBar from "@/components/SearchBar.vue";
import type { ServiceInfo } from "~/interfaces/serviceinfo";

definePageMeta({
  middleware: "auth",
});

const errorMessage = ref<string | null>(null);
const services = ref<ServiceInfo[]>([]);
const tokenCookie = useCookie("token");

onMounted(() => {
  fetchServices();
});

const fetchServices = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<ServiceInfo[]>("/api/workflow/services", {
      method: "POST",
      body: {
        token: tokenCookie.value,
      },
    });
    services.value = result;
    console.log("services", services.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
  }
};

const searchQuery = ref("");

const apps = computed(() =>
  services.value.map((service) => ({
    name: service.name,
    icon: `my-icons:white-${service.name}`,
  })),
);

const filteredApps = computed(() => {
  return apps.value.filter((app) =>
    app.name.toLowerCase().includes(searchQuery.value.toLowerCase()),
  );
});
</script>

<template>
  <div class="py-8 text-center font-sans w-full">
    <h1 class="text-custom_size_title font-custom_weight_title pb-5">
      My Services
    </h1>

    <div class="flex justify-center mb-4">
      <SearchBar v-model:search-query="searchQuery" />
    </div>

    <UContainer
      class="flex flex-wrap gap-5 justify-center p-4 bg-white rounded-lg w-full mx-auto"
    >
      <ServiceList :apps="filteredApps" />
    </UContainer>
  </div>
</template>

<style scoped></style>
