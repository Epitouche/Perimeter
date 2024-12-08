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

    if (typeof response === "object" && response !== null && "tokens" in response && Array.isArray((response as { tokens: unknown }).tokens)) {
      const tokens = (response as { tokens: Array<{ service_id: { name: string } }> }).tokens;
      serviceNames = tokens.map((token) => token.service_id.name);
      console.log("Service Names Updated:", serviceNames);
    } else {
      console.error("Response does not contain valid tokens.");
      return [];
    }
  } catch (error) {
    if (error instanceof Error) {
      console.error("Unexpected error:", error.message);
    } else {
      console.error("Unknown error occurred.");
    }
    return [];
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
  if (normalizedLabel === "spotify") {
    if (serviceNames.map(name => name.toLowerCase()).includes(normalizedLabel)) {
      //Disconnect
      alert("Already connected to Spotify.");
    } else {
      const spotifyApiLink = "http://server:8080/api/v1/spotify/auth/";
      authApiCall(spotifyApiLink);
    }
  } else if (normalizedLabel === "gmail") {
    if (serviceNames.map(name => name.toLowerCase()).includes(normalizedLabel)) {
      alert("Already connected to Gmail.");
    } else {
      const gmailApiLink = "http://server:8080/api/v1/gmail/auth/";
      authApiCall(gmailApiLink);
    }
  } else {
    console.log(`${label} unknown icon clicked`);
  }
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
      class="app_button flex flex-col items-center justify-center w-[15rem] h-[15rem] rounded-lg transition-transform hover:scale-105"
      :style="{ backgroundColor: app.color }"
      :disabled="serviceNames.includes(app.name)"
      @click="handleClick(app.name)"
    >
      <span class="text-3xl text-white font-bold mt-auto">{{ app.name }}</span>
    </UButton>
  </UContainer>
</template>

<style scoped>

:deep(.app_button span) {
  height: 6rem;
  width: 6rem;
  color: white;
}

</style>
