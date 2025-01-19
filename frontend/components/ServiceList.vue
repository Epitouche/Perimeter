<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { Token, ServiceResponse } from "~/interfaces/serviceResponse";
import { fetchServices } from "~/utils/fetchServices";
import { handleClick } from "~/utils/authUtils";
import { servicesConnectionInfos } from "~/utils/fetchServicesConnectionInfos.js";

/**
 * The lit of services to be displayed with the type of styling.
 */
const props = defineProps<{
  styling: string; // The type of styling to be used
  apps: {
    // The list of services to be displayed
    name: string; // The name of the service
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

/**
 * Load the services connection infos for the user.
 */
async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);

      if (infosConnection.value) {
        tokens.value = infosConnection.value.tokens;

        serviceConnected.value = tokens.value.map(
          (token) => token.service.name
        );
      }

      isLoading.value = false;
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

/**
 * Load the services from the backend.
 */
const loadServices = async () => {
  try {
    errorMessage.value = null;
    services.value = await fetchServices();
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading services:", error);
  }
};

/**
 * Get the service details for the services.
 */
const serviceDetails = computed(() =>
  services.value.map((service) => ({
    name: service.name,
    color: service.color,
    icon: service.icon,
    oauth: service.oauth,
  }))
);

/**
 * Get the state text for the service.
 */
const getServiceStateText = (appName: string) => {
  const matchingService = services.value.find(
    (service) => service.name === appName && !service.oauth
  );

  if (matchingService) {
    return "Disconnect";
  }

  const isConnected = serviceConnected.value.includes(appName);
  const message = isConnected ? "Disconnect" : "Connect";
  return message;
};

/**
 * Check if the service is connected or invalid.
 */
const isServiceConnectedOrInvalid = (appName: string): boolean => {
  const matchingService = services.value.find(
    (service) => service.name.toLowerCase() === appName.toLowerCase()
  );

  if (
    serviceConnected.value.includes(appName) ||
    (matchingService && matchingService.oauth === false)
  ) {
    return true;
  }
  return false;
};

/**
 * Get the service details for the service.
 */
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

/**
 * Confirm the action to be executed
 */
const confirmAction = async () => {
  if (!selectedService.value) return;
  await executeHandleClick(selectedService.value);
  isPopupVisible.value = false;
  selectedService.value = null;
};

/**
 * Execute the handle click action.
 */
const executeHandleClick = async (label: string) => {
  try {
    const response = await handleClick(
      label,
      services,
      tokens,
      tokenCookie.value || undefined
    );
    if (response) {
      loadConnectionInfos();
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error executing handleClick:", error);
  }
};

/**
 * If the action is canceled, close the popup.
 */
const cancelAction = () => {
  isPopupVisible.value = false;
  selectedService.value = null;
};

/**
 * Hover state for the service.
 */
const hover = reactive<{ [key: string]: boolean }>(
  Object.fromEntries(props.apps.map((app) => [app.name, false]))
);

/**
 * Check if the text is longer than 10 characters.
 */
const isLongText = (text: string): boolean => text.length > 10;

/**
 * Format the name of the service.
 */
function formatName(name: string): string {
  return name
    .replace(/^action_/, "")
    .replace(/_/g, " ")
    .replace(/([a-z])([A-Z])/g, "$1 $2");
}

/**
 * Check if the device is a touch device.
 */
function isTouchDevice() {
  return window.matchMedia("(pointer: coarse)").matches;
}

/**
 * When the component is mounted, load the connection information and services list.
 */
onMounted(() => {
  loadConnectionInfos();
  loadServices();
});
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full max-h-full' }"
    class="flex flex-row justify-around items-center gap-10 flex-wrap py-5 w-full h-full rounded-custom_border_radius"
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
        @mouseenter="hover[app.name] = true"
        @mouseleave="hover[app.name] = false"
      >
        <h5
          class="clamp-1-line break-words text-center pt-4 -m-1 text-white w-full hover-expand-text"
        >
          {{ formatName(app.name) }}
        </h5>
        <img
          v-if="getServiceDetails(app.name)?.icon && !hover[app.name]"
          :src="getServiceDetails(app.name)?.icon"
          alt=""
          class="pb-2"
          style="width: 40%; min-width: 1vw; max-width: 8vw"
        />
        <img
          v-else-if="
            getServiceDetails(app.name)?.icon &&
            hover[app.name] &&
            !isLongText(app.name)
          "
          :src="getServiceDetails(app.name)?.icon"
          alt=""
          class="pb-2"
          style="width: 40%; min-width: 1vw; max-width: 8vw"
        />

        <UButton
          v-if="!isLoading"
          :class="[
            isServiceConnectedOrInvalid(app.name)
              ? 'bg-black text-white'
              : 'bg-white text-black',
            'py-[8%] !rounded-t-none !rounded-b-[25%]',
          ]"
          style="min-width: 100%; height: fit-content"
          @click="onClick(app.name)"
        >
          <h6 class="w-full">{{ getServiceStateText(app.name) }}</h6>
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
      <p v-if="isTouchDevice() && styling === 'button'" class="text-center">
        {{ formatName(app.name) }}
      </p>
    </div>

    <div
      v-if="isPopupVisible && styling == 'card'"
      ref="focusDiv"
      class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50"
    >
      <div
        class="flex flex-col justify-between items-center gap-10 bg-custom_color-bg_section p-10 border-custom_border_width rounded-custom_border_radius shadow-lg w-fit h-fit"
        :style="{
          borderColor: selectedService
            ? getServiceDetails(selectedService)?.color || '#ccc'
            : '#ccc',
        }"
      >
        <h4>
          Are you sure you want to <br />
          disconnect from this service?
        </h4>
        <h6>This action cannot be undone!</h6>
        <div class="flex flex-row justify-center items-center gap-5 w-full">
          <UButton
            class="text-black !border-custom_border_width border-black bg-opacity-0 font-semibold px-[4%] py-[1.5%]"
            tabindex="0"
            @click="cancelAction"
          >
            <h6>No</h6>
          </UButton>
          <UButton
            class="text-red-600 !border-custom_border_width border-red-600 bg-opacity-0 font-semibold px-[4%] py-[1.5%]"
            tabindex="0"
            @click="confirmAction"
          >
            <h6>Yes</h6>
          </UButton>
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
