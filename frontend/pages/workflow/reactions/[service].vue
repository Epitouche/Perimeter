<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";

interface Reaction {
  id: number;
  name: string;
  description: string;
}

definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const serviceId = route.params.service;
const token = useCookie("token");
const isLoading = ref(true);

const reactions = ref<Reaction[] | null>(null);
const errorMessage = ref<string | null>(null);

const serviceInfo = ref<ServiceInfo | null>(null);

const getServiceInfo = async () => {
  if (serviceId) {
    isLoading.value = true;
    try {
      errorMessage.value = null;
      serviceInfo.value = await $fetch("/api/servicebyid", {
        method: "POST",
        body: {
          token: token.value,
          serviceId: serviceId,
        },
      });
      console.log("serviceInfo: ", serviceInfo.value);
    } catch (error: unknown) {
      errorMessage.value = handleErrorStatus(error);

      if (errorMessage.value === "An unknown error occurred") {
        console.error("An unknown error occurred", error);
      }
      console.log("error", error);
    } finally {
      isLoading.value = false;
    }
  }
};

const fetchReactions = async () => {
  isLoading.value = true;
  try {
    errorMessage.value = null;
    reactions.value = await $fetch<Reaction[]>("/api/workflow/reactions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });

    console.log("reactions", reactions.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  getServiceInfo();
  fetchReactions();
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
      class="py-20"
      :style="{ backgroundColor: serviceInfo.color }"
    >
      <div class="px-20">
        <BackButton link="/workflow/reactions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-2">
        <h1 class="text-white">
          Add a reaction
        </h1>
        <img
          :src="serviceInfo.icon"
          :alt="serviceInfo.name"
          class="w-[8em] h-[8em]"
        >
        <h2 class="capitalize text-white text-7xl font-bold pt-8">
          {{ serviceInfo.name }}
        </h2>
      </div>
    </UContainer>
    <div v-if="errorMessage">
      <div>Error: {{ errorMessage }}</div>
    </div>
    <div v-else-if="reactions">
      <ReActionCardContainer
        type-name="reaction"
        :types="reactions"
        :service-info="serviceInfo"
      />
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
  </div>
</template>

<style scoped></style>
