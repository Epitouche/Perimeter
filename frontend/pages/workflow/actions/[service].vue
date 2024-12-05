<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  // middleware: 'auth'
  middleware: 'guest' // temporary
});

const route = useRoute();

const serviceId = route.params.service;

</script>

<template>
  <h1>
    Service Actions Page
  </h1>
  <p>You are viewing service with ID: {{ serviceId }}</p>
</template>

<style scoped></style>
