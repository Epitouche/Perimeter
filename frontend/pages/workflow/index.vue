<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  // middleware: 'auth'
  middleware: 'guest' // temporary
});

</script>

<template>
  <h1>
    Workflow Page
  </h1>
  <UButton to="/workflow/actions">Actions</UButton>
  <UButton to="/workflow/reactions">Reactions</UButton>
</template>

<style scoped></style>
