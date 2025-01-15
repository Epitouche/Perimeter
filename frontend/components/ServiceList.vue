<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { Token, ServiceResponse } from "~/interfaces/serviceResponse";
import { fetchServices } from "~/utils/fetchServices";
import { handleClick } from "~/utils/authUtils";
import { servicesConnectionInfos } from "~/utils/fetchServicesConnectionInfos.js";

const props = defineProps<{
  styling: string;
  apps: {
    name: string;
  }[];
}>();

const tokenCookie = useCookie("token");
const errorMessage = ref<string | null>(null);
const isLoading = ref(true);
const isPopupVisible = ref(false);
const services = ref<ServiceInfo[]>([]);
const serviceConnected = ref<string[]>([]);
const tokens = ref<Token[]>([]);
const infosConnection = ref<ServiceResponse | null>(null);
const selectedService = ref<string | null>(null);

const isVisible = ref(false);
const focusDiv = ref<HTMLElement | null>(null);

onMounted(() => {
  loadConnectionInfos();
  loadServices();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);

      if (infosConnection.value) {
        tokens.value = infosConnection.value.tokens;

        serviceConnected.value = tokens.value.map(
          (token) => token.service.name,
        );
      }

      isLoading.value = false;
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
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
  const message = isConnected ? `Disconnect` : `Connect`;
  return message;
};

const isServiceConnectedOrInvalid = (appName: string): boolean => {
  const matchingService = services.value.find(
    (service) => service.name.toLowerCase() === appName.toLowerCase(),
  );

  if (
    serviceConnected.value.includes(appName) ||
    (matchingService && matchingService.oauth === false)
  ) {
    return true;
  }
  return false;
};

const getServiceDetails = (appName: string) =>
  serviceDetails.value.find((service) => service.name === appName);

const onClick = (label: string) => {
  if (isServiceConnectedOrInvalid(label)) {
    selectedService.value = label;
    isPopupVisible.value = true;
    isVisible.value = !isVisible.value;

    if (isVisible.value && focusDiv.value) {
      focusDiv.value.focus();
    }
  } else {
    executeHandleClick(label);
  }
};

const confirmAction = async () => {
  if (!selectedService.value) return;
  await executeHandleClick(selectedService.value);
  isPopupVisible.value = false;
  selectedService.value = null;
};

const executeHandleClick = async (label: string) => {
  try {
    const response = await handleClick(
      label,
      services,
      tokens,
      tokenCookie.value || undefined,
    );
    if (response) {
      loadConnectionInfos();
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error executing handleClick:", error);
  }
};

const cancelAction = () => {
  isPopupVisible.value = false;
  selectedService.value = null;
};

const hover = reactive<{ [key: string]: boolean }>(
  Object.fromEntries(props.apps.map((app) => [app.name, false])),
);
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full max-h-full' }"
    class="flex flex-row justify-center items-center gap-10 flex-wrap py-5 w-full h-full rounded-custom_border_radius"
  >
    <div v-for="(app, index) in apps" :key="index">
      <UContainer
        v-if="styling === 'card'"
        :ui="{ padding: '!px-0', constrained: 'max-w-none' }"
        class="custom_card button_shadow !justify-between !gap-0 overflow-hidden"
        tabindex="0"
        :style="{
          backgroundColor: getServiceDetails(app.name)?.color || '#ccc',
        }"
        @click="onClick(app.name)"
      >
        <h3
          class="clamp-1-line break-words text-center pt-4 text-white w-full hover-expand-text"
        >
          {{ app.name }}
        </h3>
        <img
          v-if="getServiceDetails(app.name)?.icon"
          :src="getServiceDetails(app.name)?.icon"
          alt=""
          class="icon_circle"
        />
        <UButton
          :class="[
            isServiceConnectedOrInvalid(app.name)
              ? 'bg-black text-white'
              : 'bg-white text-black',
            'w-full min-h-[5vh] max-h-[5vh] !rounded-t-none',
          ]"
          @click="onClick(app.name)"
        >
          <h4 class="w-full">{{ getServiceStateText(app.name) }}</h4>
        </UButton>
      </UContainer>

      <UContainer
        v-if="styling === 'button'"
        :ui="{ padding: '!px-0', constrained: 'max-w-none' }"
        class="custom_card_circle"
        tabindex="0"
        :style="{
          backgroundColor: getServiceDetails(app.name)?.color || '#ccc',
        }"
        @mouseenter="hover[app.name] = true"
        @mouseleave="hover[app.name] = false"
        @click="onClick(app.name)"
      >
        <img
          v-if="getServiceDetails(app.name)?.icon && !hover[app.name]"
          :src="getServiceDetails(app.name)?.icon"
          alt=""
          class="icon_circle"
        />
        <UButton
          v-if="hover[app.name]"
          variant="ghost"
          @click="onClick(app.name)"
        >
          <p
            class="clamp-1-line break-words text-center text-white py-2 w-full hover-expand-text"
          >
            {{ app.name }}
          </p>
        </UButton>
      </UContainer>
    </div>

    <div
      v-if="isPopupVisible"
      ref="focusDiv"
      class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50"
    >
      <div
        class="bg-white p-10 border-custom_border_width rounded-custom_border_radius shadow-lg max-w-md w-full"
      >
        <h2 class="mb-2">
          Are you sure you want to disconnect from this service?
        </h2>
        <p class="text-2xl mb-5">This action cannot be undone!</p>
        <div class="flex flex-row justify-end items-center gap-5 pt-5">
          <UButton
            class="text-black border-2 border-black bg-opacity-0 text-2xl font-semibold py-3 px-5"
            tabindex="0"
            @click="cancelAction"
            >No</UButton
          >
          <UButton
            class="text-red-600 border-2 border-red-600 bg-opacity-0 text-2xl font-semibold py-3 px-5"
            tabindex="0"
            @click="confirmAction"
            >Yes</UButton
          >
        </div>
      </div>
    </div>
  </UContainer>
</template>

<style scoped>
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

[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}

.button_shadow {
  box-shadow: 0 5px 10px rgba(0, 0, 0, 0.2);
}
</style>
