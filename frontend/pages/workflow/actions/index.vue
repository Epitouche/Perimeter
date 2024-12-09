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
  <div>
    <h1>Connected services for Actions</h1>
    <UButton to="/workflow">Back</UButton>
    <div v-if="error">Error: {{ error }}</div>
    <div v-else-if="services">
      <div v-for="service in services" :key="service.id">
        <NuxtLink
          :to="{
            name: 'workflow-actions-service',
            params: { service: service.id },
          }"
        >
          {{ service.name }}
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<style scoped></style>