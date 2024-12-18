<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const token = useCookie("token");

const error = ref<string | null>(null);
const services = ref<any[]>([]);
const filteredServices = ref<any[]>([]);
const isLoading = ref(true);

const searchQuery = ref<string>("");

const fetchServices = async () => {
  try {
    error.value = null;
    const result = await $fetch<any[]>("/api/workflow/services", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    services.value = result;
    filteredServices.value = result;
    console.log("services", services.value);
  } catch (error) {
    filteredServices.value = [];
    console.error("Error fetching services:", error);
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
  <div class="p-10">
    <BackButton link="/workflow" :is-white="false" />
    <h1
      class="flex justify-center w-full text-8xl font-custom_weight_title pb-20"
    >
      Add a reaction
    </h1>
    <UContainer
      :ui="{ base: 'mx-auto' }"
      class="flex flex-col justify-center items-center gap-16 w-full h-full !p-0"
    >
      <SearchBar v-model:search-query="searchQuery" class="!w-1/3" />
      <div v-if="error">Error: {{ error }}</div>
      <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
      <div
        v-else-if="filteredServices.length"
        class="flex flex-row justify-evenly items-center w-full"
      >
        <ServiceCardContainer :services="filteredServices" />
      </div>
    </UContainer>
  </div>
</template>

<style scoped></style>
