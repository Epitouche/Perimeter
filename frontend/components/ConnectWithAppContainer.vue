<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import { fetchServices } from "~/utils/fetchServices";

const services = ref<ServiceInfo[]>([]);
const errorMessage = ref<string | null>(null);

onMounted(() => {
  loadServices();
});

const loadServices = async () => {
  try {
    errorMessage.value = null;
    services.value = await fetchServices();
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading services:", error);
  }
};

const filteredServices = computed(() =>
  services.value
    .filter((service) => service.oauth)
    .map((service) => ({
      name: service.name,
    })),
);
</script>

<template>
  <UContainer
    :ui="{ padding: 'px-0' }"
    class="bg-custom_color-bg_section min-w-full flex flex-wrap justify-evenly"
  >
    <ServiceList styling="button" :apps="filteredServices" />
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
}
</style>
