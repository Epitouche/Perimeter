<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { Token } from "~/interfaces/serviceResponse";
import { fetchServices } from "~/utils/fetchServices";
import { handleClick } from "~/utils/authUtils";
import { servicesConnectionInfos } from "~/utils/fetchServicesConnectionInfos.js";

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
const tokens = ref<Token[]>([]);

onMounted(() => {
  loadConnectionInfos();
  loadServices();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      tokens.value = await servicesConnectionInfos(tokenCookie.value);
      serviceConnected.value = tokens.value.map((token) => token.service.name);
      isLoading.value = false;
    }
  } catch (error) {
    console.error("Error loading tokens:", error);
  }
}

const loadServices = async () => {
  try {
    errorMessage.value = null;
    services.value = await fetchServices();
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
  const matchingService = services.value.find(
    (service) => service.name === appName && !service.oauth,
  );

  if (matchingService) {
    return "Automatically connected";
  }
  const isConnected = serviceConnected.value.includes(appName);
  const message = isConnected ? `Disconnect ${appName}` : `Connect ${appName}`;
  return message;
};

const isServiceConnectedOrInvalid = (appName: string): boolean => {
  const matchingService = services.value.find((service) => service.name.toLowerCase() === appName.toLowerCase(),);

  if (serviceConnected.value.includes(appName) || (matchingService && matchingService.oauth === false) ) {
    return true;
  }
  return false;
};

const getServiceDetails = (appName: string) =>
  serviceDetails.value.find((service) => service.name === appName);

const onClick = (label: string) => {
  if (tokenCookie.value) {
    handleClick(label, services, tokens, tokenCookie.value);
  } else {
    handleClick(label, services, undefined, undefined);
  }
};
</script>

<template>
  <div class="flex flex-wrap gap-5 justify-center">
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :style="{ backgroundColor: getServiceDetails(app.name)?.color || '#ccc' }"
      :class="[
        `flex flex-col items-center justify-start relative p-5 w-[15rem] h-[15rem] font-extrabold rounded-custom_border_radius overflow-hidden transition-transform hover:scale-105`,
      ]"
      @click="onClick(app.name)"
    >
      <img
        v-if="getServiceDetails(app.name)?.icon"
        :src="getServiceDetails(app.name)?.icon"
        alt="service_icon"
        class="w-20 h-20"
      >

      <span
        class="clamp-1-line p-8 text-2xl text-center break-words w-full hover-expand-text"
        >{{ app.name }}</span
      >

      <div
        v-if="!isLoading"
        class="absolute bottom-0 w-full h-[3rem] flex items-center justify-center text-2x1 font-bold"
        :class="{
          'bg-black text-white': isServiceConnectedOrInvalid(app.name),
          'bg-white text-black': !isServiceConnectedOrInvalid(app.name),
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

.clamp-1-line {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: normal;
  transition: all 1s ease-in-out;
}

.hover-expand-text:hover {
  -webkit-line-clamp: unset;
  line-clamp: unset;
  overflow: visible;
  white-space: normal;
}
</style>
