<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  // middleware: 'auth'
  middleware: 'guest' // temporary
});

const showNavBar = ref(true);
const showCancelButton = ref(false);
const reactionButtonisDisabled = ref(true);

const onActionSelected = () => {
  showNavBar.value = false;
  showCancelButton.value = true;
  reactionButtonisDisabled.value = false;
};

const setWorkflowPageDefault = () => {
  showNavBar.value = true;
  showCancelButton.value = false;
  reactionButtonisDisabled.value = true;
};

</script>

<template>
  <div v-if="showNavBar">
    <NavigationBar />
  </div>
  <h1>
    Workflow Page
  </h1>
  <UButton to="/workflow/actions">Actions</UButton>
  <UButton :disabled="reactionButtonisDisabled" to="/workflow/reactions">Reactions</UButton>
</template>

<style scoped></style>
