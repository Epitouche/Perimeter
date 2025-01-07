<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import { fetchServices } from "~/utils/fetchServices"
import { handleClick } from "~/utils/authUtils";

defineProps<{
  apps: {
    name: string;
  }[];
}>();

const tokenCookie = useCookie("token");
const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const services = ref<ServiceInfo[]>([]);
const serviceConnected = ref<string[]>([]);

onMounted(() => {
  servicesConnectionInfos();
  loadServices();
});

async function servicesConnectionInfos() {
  try {
    const response = await $fetch("/api/auth/service/infos", {
      method: "POST",
      body: {
        authorization: tokenCookie.value,
      },
    });

    if (
      typeof response === "object" &&
      response !== null &&
      "tokens" in response &&
      Array.isArray((response as { tokens: unknown }).tokens)
    ) {
      const tokens = (
        response as { tokens: Array<{ service: { name: string } }> }
      ).tokens;
      serviceConnected.value = tokens.map((token) => token.service.name);
      console.log("tokens: ", tokens);
      console.log("serviceConnected", serviceConnected);
      isLoading.value = false;
    } else {
      console.error("Response does not contain valid tokens.");
      return [];
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
  }
}

const loadServices = async () => {
  try {
    errorMessage.value = null;
      services.value = await fetchServices();
      console.log("services", services.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading services:", error);
  }
};

const serviceDetails = computed(() =>
  services.value.map((service) => ({
    name: service.name,
    color: service.color,
    icon: service.icon,
    oauth: service.oauth,
  })),
);

const getServiceStateText = (appName: string) => {
  const isConnected = serviceConnected.value.includes(appName);
  const message = isConnected ? `Disconnect ${appName}` : `Connect ${appName}`;
  return message;
};

const getServiceDetails = (appName: string) =>
  serviceDetails.value.find(
    (service) => service.name === appName,
  );

const onClick = (label: string) => {
  handleClick(label, services, serviceConnected);
};

</script>

<template>
  <div class="flex flex-wrap gap-5 justify-center">
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :style="{ backgroundColor: getServiceDetails(app.name)?.color || '#ccc' }"
      :class="[
        `app_button flex flex-col items-center justify-start relative w-[15rem] h-[15rem] rounded-[25%] overflow-hidden transition-transform hover:scale-105`,
      ]"
      @click="onClick(app.name)"
    >
      <img
        v-if="getServiceDetails(app.name)?.icon"
        :src="getServiceDetails(app.name)?.icon"
        alt=""
        class="w-20 h-20 mt-4"
      >
      
      <span class="text-3xl font-bold text-white mt-auto mb-[2.25rem]">{{
        app.name
      }}</span>

      <div
        v-if="!isLoading"
        class="absolute bottom-0 w-full h-[3rem] flex items-center justify-center text-2x1 font-bold"
        :class="{
          'bg-black text-white':
            serviceConnected.includes(app.name),
          'bg-white text-black':
            !serviceConnected.includes(app.name),
        }"
      >
        {{ getServiceStateText(app.name) }}
      </div>
    </UButton>
  </div>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
  color: white;
}
</style>
