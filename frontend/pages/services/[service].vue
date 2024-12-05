<script setup lang="ts">

const route = useRoute();

onMounted(async () => {
  const service = route.params.service;

  try {
    console.log(`Connecting to ${service}...`);
    await new Promise((resolve) => setTimeout(resolve, 5000));
  } catch (error) {
    console.error(`Failed to connect to ${service}:`, error);
    return navigateTo('/error')
  }
});
</script>

<template>
  <div>
    <LoadingScreen :timeout="5000"/>
  </div>
</template>
