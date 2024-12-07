<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const isLoading = ref(true);
</script>

<template>
  <div>
    <LoadingScreen v-if="isLoading" :timeout="500" />
  </div>
</template>
