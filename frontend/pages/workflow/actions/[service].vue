<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";

/**
 * @description This page is used to add an action to a service.
 */

definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

/**
 * @description This type is used to define the structure of an action.
 * @property {number} id
 * @property {string} name
 * @property {string} description
 * @property {string} [option]
 */
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

/**
 * @description This function fetches the service information.
 */
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

/**
 * @description This function handles the error status.
 * @param {unknown} error
 * @returns {string}
 */
onMounted(() => {
  getServiceInfo();
  fetchActions();
});
</script>

<template>
  <div class="flex flex-col gap-16">
    <div v-if="errorMessage">
      <UAlert
        color="red"
        variant="solid"
        title="ERROR"
        :description="errorMessage"
        class="justify-center items-center gap-5 w-[15%]"
      />
    </div>
    <div v-else-if="isLoading"><h2>Loading...</h2></div>
    <UContainer
      v-else-if="serviceInfo"
      :ui="{ constrained: 'max-w-none' }"
      class="pt-10 max-sm:pt-3"
      :style="{ backgroundColor: serviceInfo.color }"
    >
      <div class="px-10 -mb-5 max-sm:mb-0 max-sm:px-0 max-sm:pb-5">
        <BackButton link="/workflow/actions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-0 pb-5">
        <h1 class="text-white">Add an action</h1>
        <img
          :src="serviceInfo.icon"
          :alt="serviceInfo.name"
          class="w-[12vw] h-[12vh] max-sm:w-[25vw] max-sm:h-[25vh] max-sm:-my-10"
        />
        <h2 class="capitalize text-white">
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
    <div v-else-if="isLoading"><h2>Loading...</h2></div>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
