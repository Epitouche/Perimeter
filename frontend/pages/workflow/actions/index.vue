<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  // middleware: 'auth'
  middleware: 'guest' // temporary
});

</script>

<template>
  <h1>
    Connected services for Actions
  </h1>
  <NuxtLink :to="{ name: 'workflow-actions-service', params: { service: 1 } }"> ServiceName Actions </NuxtLink>
</template>

<style scoped></style>
