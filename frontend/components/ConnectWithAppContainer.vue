<script setup lang="ts">

const props = defineProps<{
  apps: string[]
}>();

interface OAuthLink {
  authentication_url: string;
}

const authApiCall = async (label: string) => {
  try {
    const response = await $fetch<OAuthLink>('/api/auth/service/redirect', {
      method: 'POST',
      body: {
        link: label,
      },
    });
    navigateTo(response.authentication_url, { external: true})
    console.log(response.authentication_url)
    return response;
  } catch (err: any) {
    console.log(err.message);
    throw err;
  }
};

const handleClick = (label: string) => {
  if (label == 'i-logos-spotify-icon') {
    console.log('Spotify icon clicked');
    const spotifyApiLink = 'http://server:8080/api/v1/spotify/auth/';
    authApiCall(spotifyApiLink);
  } else if (label === 'i-logos-google-icon') {
    console.log('Google icon clicked');
    const gmailApiLink = 'http://server:8080/api/v1/gmail/auth/';
    authApiCall(gmailApiLink);
  } else {
    console.log(`${label} unknown icon clicked`);
  }
};

</script>

<template>
  <UContainer :ui="{ padding: 'px-0' }" class="bgbg-custom_color-bg_section min-w-full flex flex-wrap justify-between">
    <UButton variant="ghost" v-for="(app, index) in apps" :key="index" @click="handleClick(app)" :icon="app" class="app_button basis-1/3 flex justify-center" />
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
}
</style>
