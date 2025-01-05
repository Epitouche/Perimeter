<script setup lang="ts">
import type { ServiceInfo } from "~/interfaces/serviceinfo";
import type { OAuthLink } from "~/interfaces/authLink";

defineProps<{
  apps: {
    name: string;
    icon: string;
  }[];
}>();

const tokenCookie = useCookie("token");
const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const services = ref<ServiceInfo[]>([]);

let serviceConnected: string[] = [];

onMounted(() => {
  servicesConnectionInfos();
  fetchServices();
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
      serviceConnected = tokens.map((token) => token.service.name);
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

const fetchServices = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<ServiceInfo[]>("/api/workflow/services", {
      method: "POST",
      body: {
        token: tokenCookie.value,
      },
    });
    services.value = result;
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
  }
};

const authApiCall = async (label: string) => {
  try {
    const response = await $fetch<OAuthLink>("/api/auth/service/redirect", {
      method: "POST",
      body: {
        link: label,
      },
    });
    navigateTo(response.authentication_url, { external: true });
    return response;
  } catch (err) {
    if (err instanceof Error) {
      console.error(err.message);
    } else {
      console.error("Unexpected error:", err);
    }
    throw err;
  }
};

const handleClick = (label: string) => {
  const normalizedLabel = label.toLowerCase();
  const serviceNames = services.value.map((service) =>
    service.name.toLowerCase(),
  );
  if (
    serviceConnected.map((name) => name.toLowerCase()).includes(normalizedLabel)
  ) {
    //disconnectService(label);
  } else {
    const apiLink = `http://server:8080/api/v1/${normalizedLabel}/auth/`;

    if (serviceNames.includes(normalizedLabel)) {
      authApiCall(apiLink);
    } else {
      console.log(`${label} unknown icon clicked`);
    }
  }
};

const getServiceStateText = (appName: string) => {
  if (["openweathermap", "timer"].includes(appName.toLowerCase())) {
    return `Automatically connected`;
  }

  const isConnected = serviceConnected.includes(appName.toLowerCase());
  const message = isConnected ? `Disconnect ${appName}` : `Connect ${appName}`;
  return message;
};

const isSpecialCase = (appName: string) => {
  return ["openweathermap", "timer"].includes(appName.toLowerCase());
};
</script>

<template>
  <div class="flex flex-wrap gap-5 justify-center">
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :icon="app.icon"
      :class="[
        `bg-custom_color-${app.name}`,
        `app_button flex flex-col items-center justify-start relative w-[15rem] h-[15rem] rounded-[25%] overflow-hidden transition-transform hover:scale-105`,
      ]"
      @click="handleClick(app.name)"
    >
      <span class="text-3xl font-bold text-white mt-auto mb-[2.25rem]">{{
        app.name
      }}</span>

      <div
        v-if="!isLoading"
        class="absolute bottom-0 w-full h-[3rem] flex items-center justify-center text-2x1 font-bold"
        :class="{
          'bg-black text-white':
            isSpecialCase(app.name) ||
            serviceConnected.includes(app.name.toLowerCase()),
          'bg-white text-black':
            !isSpecialCase(app.name) &&
            !serviceConnected.includes(app.name.toLowerCase()),
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
