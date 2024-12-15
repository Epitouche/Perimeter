<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const token = useCookie("token");

const error = ref<string | null>(null);
const services = ref<any[]>([]);
const filteredServices = ref<any[]>([]);

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
  }
};

watch(searchQuery, (newQuery) => {
  filteredServices.value = services.value.filter((service) =>
    service.name.toLowerCase().includes(newQuery.toLowerCase())
  );
});

onMounted(() => {
  fetchServices();
});
</script>

<template>
  <div class="p-10">
    <BackButton link="/workflow" color="black" />
    <h1 class="flex justify-center w-full text-8xl font-custom_weight_title pb-20">Add a reaction</h1>
    <UContainer :ui="{ base: 'mx-auto' }" class="flex flex-col justify-center items-center gap-10 w-full h-full !p-0">
      <SearchBar v-model:search-query="searchQuery" class="!w-1/3" />
      <div v-if="error">Error: {{ error }}</div>
      <div v-else-if="filteredServices.length" class="flex flex-row justify-evenly items-center flex-wrap w-full">
        <div v-for="service in filteredServices" :key="service.id">
          <NuxtLink :to="{
            name: 'workflow-reactions-service',
            params: { service: service.id },
          }">
            {{ service.name }}
          </NuxtLink>
        </div>
      </div>
    </UContainer>
  </div>
</template>

<style scoped></style>