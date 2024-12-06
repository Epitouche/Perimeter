<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'auth'
});

const route = useRoute();

const serviceId = route.params.service;

</script>

<template>
  <div>
    <h1>
      Service Actions Page
    </h1>
    <p>You are viewing service with ID: {{ serviceId }}</p>
  </div>
</template>

<style scoped></style>
