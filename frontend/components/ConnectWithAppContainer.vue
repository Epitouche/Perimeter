<script setup lang="ts">
defineProps<{
  apps: string[];
}>();

interface OAuthLink {
  authentication_url: string;
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
  if (label == "my-icons:color-spotify") {
    console.log("Spotify icon clicked");
    const spotifyApiLink = "http://server:8080/api/v1/spotify/auth/";
    authApiCall(spotifyApiLink);
  } else if (label === "my-icons:color-google") {
    console.log("Google icon clicked");
    const gmailApiLink = "http://server:8080/api/v1/gmail/auth/";
    authApiCall(gmailApiLink);
  } else {
    console.log(`${label} unknown icon clicked`);
  }
};
</script>

<template>
  <UContainer
    :ui="{ padding: 'px-0' }"
    class="bg-custom_color-bg_section min-w-full flex flex-wrap justify-evenly"
  >
    <UButton
      variant="ghost"
      v-for="(app, index) in apps"
      :key="index"
      @click="handleClick(app)"
      :icon="app"
      class="app_button basis-1/3 flex justify-center"
    />
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
}
</style>
