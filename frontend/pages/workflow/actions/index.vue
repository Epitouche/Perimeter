<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const token = useCookie("token");

const services = ref<any>(null);
const error = ref<string | null>(null);

const fetchServices = async () => {
  try {
    error.value = null;
    services.value = await $fetch("/api/workflow/services", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    console.log("services", services.value);
  } catch (error) {
    console.error("Error fetching services:", error);
  }
};

onMounted(() => {
  fetchServices();
});
</script>

<template>
  <div class="flex flex-col p-10">
    <div class="flex flex-row items-center">
      <BackButton link="/workflow" />
      <h1 class="w-full flex justify-center text-8xl font-custom_weight_title">Add an action</h1>
    </div>
    <UContainer class="">
      <SearchBar />
      <div v-if="error">Error: {{ error }}</div>
      <div v-else-if="services">
        <div v-for="service in services" :key="service.id">
          <NuxtLink
          :to="{
            name: 'workflow-actions-service',
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