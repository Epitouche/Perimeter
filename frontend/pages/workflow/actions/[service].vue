<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";

definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

interface ActionType {
  id: number;
  name: string;
  description: string;
  option?: string;
}

const route = useRoute();
const serviceId = route.params.service as string;
const token = useCookie("token");

const isLoading = ref(true);
const actions = ref<ActionType[] | null>(null);
const errorMessage = ref<string | null>(null);
const serviceInfo = ref<ServiceInfo | null>(null);

const getServiceInfo = async () => {
  if (!serviceId) return;

  isLoading.value = true;
  try {
    errorMessage.value = null;
    serviceInfo.value = await $fetch<ServiceInfo>("/api/servicebyid", {
      method: "POST",
      body: {
        token: token.value,
        serviceId,
      },
    });
    console.log("serviceInfo: ", serviceInfo.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error fetching services:", errorMessage);
  } finally {
    isLoading.value = false;
  }
};

const fetchActions = async () => {
  isLoading.value = true;
  try {
    errorMessage.value = null;
    actions.value = await $fetch<ActionType[]>("/api/workflow/actions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error fetching actions:", errorMessage);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  getServiceInfo();
  fetchActions();
});
</script>

<template>
  <div class="flex flex-col gap-20">
    <div v-if="errorMessage">
      <div>Error: {{ errorMessage }}</div>
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
    <UContainer
      v-else-if="serviceInfo"
      :ui="{ constrained: 'max-w-none' }"
      class="pt-16 pb-16"
      :style="{ backgroundColor: serviceInfo.color }"
    >
      <div class="px-20">
        <BackButton link="/workflow/actions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-8">
        <h1 class="text-white">
          Add an action
        </h1>
        <img
          :src="serviceInfo.icon"
          :alt="serviceInfo.name"
          class="w-[8em] h-[8em]"
        >
        <h2 class="capitalize text-white text-7xl font-bold">
          {{ serviceInfo.name }}
        </h2>
      </div>
    </UContainer>
    <div v-if="actions">
      <ReActionCardContainer
        type-name="action"
        :types="actions"
        :service-info="serviceInfo"
      />
    </div>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
