<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'auth'
});

</script>

<template>
  <h1>
    Connected services - Reactions page
  </h1>
  <NuxtLink :to="{ name: 'workflow-actions-service', params: { service: 1 } }"> ServiceName Actions </NuxtLink>
</template>

<style scoped></style>
