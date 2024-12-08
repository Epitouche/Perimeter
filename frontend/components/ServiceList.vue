<script setup lang="ts">
defineProps<{
  apps: {
    name: string;
    color: string;
    icon: string;
  }[];
}>();

interface ServicesInfos {
  AllInfos: string;
}

const tokenCookie = useCookie("token");

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

      const serviceNames = tokens.map(token => token.service_id.name);
      console.log("Service Names:", serviceNames);
      return serviceNames;
    } else {
      console.warn("Response does not contain valid tokens.");
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


const handleClick = (name: string) => {
  console.log(`${name} clicked`);
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
