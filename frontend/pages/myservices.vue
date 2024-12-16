<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const apps = [
  {
    name: "Spotify",
    color: "#1DC000",
    icon: "i-logos-spotify-icon",
  },
  {
    name: "Gmail",
    color: "#E60000",
    icon: "i-logos-google-icon",
  },
  {
    name: "Time",
    color: "#BB00FF",
    icon: "i-ic-outline-timelapse",
  },
];
</script>

<template>
  <div>
    <div class="py-8 text-center font-sans">
      <h1 class="text-custom_size_title font-custom_weight_title pb-5">My Services</h1>
      <ServiceList :apps="apps" />
    </div>
  </div>
</template>

<style scoped></style>
