<script setup lang="ts">
import SearchBar from "@/components/SearchBar.vue";
definePageMeta({
  middleware: "auth",
});

const apps = [
  {
    name: "Spotify",
    color: "#1DC000",
    icon: "my-icons:white-spotify",
  },
  {
    name: "Gmail",
    color: "#E60000",
    icon: "my-icons:white-gmail",
  },
];

const searchQuery = ref("");

const filteredApps = computed(() => {
  return apps.filter((app) =>
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
