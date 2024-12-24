<script setup lang="ts">
defineProps<{
  apps: {
    name: string;
    color: string;
    icon: string;
  }[];
}>();

interface OAuthLink {
  authentication_url: string;
}

const tokenCookie = useCookie("token");
const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
let serviceNames: string[] = [];

onMounted(() => {
  servicesConnectionInfos();
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
        response as { tokens: Array<{ service_id: { name: string } }> }
      ).tokens;
      serviceNames = tokens.map((token) => token.service_id.name);
      //console.log("Service Names Updated:", serviceNames);
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

const authApiCall = async (label: string) => {
  try {
    const response = await $fetch<OAuthLink>("/api/auth/service/redirect", {
      method: "POST",
      body: {
        link: label,
      },
    });
    navigateTo(response.authentication_url, { external: true });
    //console.log(response.authentication_url);
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
  if (
    serviceNames.map((name) => name.toLowerCase()).includes(normalizedLabel)
  ) {
    //disconnectService(label);
  } else {
    const apiLink =
      normalizedLabel === "spotify"
        ? "http://server:8080/api/v1/spotify/auth/"
        : normalizedLabel === "gmail"
          ? "http://server:8080/api/v1/gmail/auth/"
          : null;

    if (apiLink) {
      authApiCall(apiLink);
    } else {
      console.log(`${label} unknown icon clicked`);
    }
  }
};

const getServiceStateText = (appName: string) => {
  const isConnected = serviceNames.includes(appName.toLowerCase());
  const message = isConnected ? `Disconnect ${appName}` : `Connect ${appName}`;
  console.log(`Rendering state for ${appName}: ${message}`);
  return message;
};
</script>

<template>
  <UContainer
    class="flex flex-wrap gap-5 justify-center p-4 bg-white rounded-lg mx-auto"
  >
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :icon="app.icon"
      class="app_button flex flex-col items-center justify-start relative w-[15rem] h-[15rem] rounded-[25%] overflow-hidden transition-transform hover:scale-105"
      :style="{ backgroundColor: app.color }"
      @click="handleClick(app.name)"
    >
      <span class="text-3xl font-bold text-white mt-auto mb-[2.25rem]">{{
        app.name
      }}</span>

      <div
        v-if="!isLoading"
        class="absolute bottom-0 w-full h-[3rem] flex items-center justify-center text-2x1 font-bold"
        :class="{
          'bg-black text-white': serviceNames.includes(app.name.toLowerCase()),
          'bg-white text-black': !serviceNames.includes(app.name.toLowerCase()),
        }"
      >
        {{ getServiceStateText(app.name) }}
      </div>
    </UButton>
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
  color: white;
}
</style>
