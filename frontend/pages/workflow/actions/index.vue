<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'auth'
});

const services = ref(null);
const error = ref<string | null>(null);
const isLoading = ref(true);

const fetchServices = async () => {
  try {
    error.value = null;
    isLoading.value = true;
    services.value = await $fetch('/api/workflow');
    console.log('services', services.value);
  } catch (err: any) {
    console.error('Error fetching services:', err);
    error.value = err?.message || 'Failed to fetch services';
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchServices();
});

</script>

<template>
  <div>

    <div v-if="isLoading" class="text-gray-500">Loading services...</div>
    <div v-if="error" class="text-red-500">{{ error }}</div>
    <div v-else-if="services" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
      <div
        v-for="service in services"
        :key="service.id"
        class="p-4 border rounded-lg shadow-lg bg-white hover:bg-gray-100 transition">
        <h2 class="text-xl font-semibold capitalize">{{ service.name }}</h2>
        <p class="text-gray-600">{{ service.description }}</p>
        <NuxtLink
          :to="{ name: 'workflow-actions-service-id', params: { id: service.id } }"
          class="mt-4 inline-block text-blue-500 hover:underline">
          View Actions
        </NuxtLink>
      </div>
    </div>

    <h1>
      Connected services for Actions
    </h1>
    <UButton to="/workflow">Back</UButton>
    <NuxtLink :to="{ name: 'workflow-actions-service', params: { service: 1 } }"> ServiceName Actions </NuxtLink>
  </div>
</template>

<style scoped></style>
