<script setup lang="ts">

definePageMeta({
  middleware: 'auth',
});

const apps = [
  {
    name: 'Spotify',
    color: '#1DC000',
    icon: 'i-logos-spotify-icon',
    apiLink: 'http://server:8080/api/v1/spotify/auth/',
  },
  {
    name: 'Google',
    color: '#E60000',
    icon: 'i-logos-google-icon',
    apiLink: 'http://server:8080/api/v1/gmail/auth/',
  },
  {
    name: 'Time',
    color: '#BB00FF',
    icon: 'i-logos-dropbox-icon',
    apiLink: 'http://server:8080/api/v1/dropbox/auth/',
  },
];
</script>

<template>
  <NavigationBar />
  <div class="services-page py-8 text-center font-sans">
    <h1 class="mb-4 text-3xl font-bold">My Services</h1>
    <ServiceList :apps="apps" />
  </div>
</template>

<style scoped>
.services-page {
  padding: 2rem;
  text-align: center;
  font-family: Arial, sans-serif;
}

h1 {
  margin-bottom: 1rem;
  font-size: 2rem;
}
</style>
