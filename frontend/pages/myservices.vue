<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const apps = [
  {
    name: "Spotify",
    color: "#1DC000",
    icon: "my-icons:white-spotify",
  },
  {
    name: "Gmail",
    color: "#E60000",
    icon: "my-icons:white-gmail",
  },
  {
    name: "Time",
    color: "#BB00FF",
    icon: "my-icons:white-timer",
  },
];
</script>

<template>
  <div>
    <div class="py-8 text-center font-sans">
      <h1 class="text-custom_size_title font-custom_weight_title pb-5">
        My Services
      </h1>
      <ServiceList :apps="apps" />
    </div>
  </div>
</template>

<style scoped></style>
