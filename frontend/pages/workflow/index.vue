<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  // middleware: 'auth'
  middleware: 'guest' // temporary
});

const showNavBar = ref(true);
const showCancelButton = ref(false);
const reactionButtonisDisabled = ref(true);
const showAddButton = ref(false);

const onActionSelected = () => {
  showNavBar.value = false;
  showCancelButton.value = true;
  reactionButtonisDisabled.value = false;
};

const onReactionSelected = () => {
  showAddButton.value = true;
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
  <div v-if="showCancelButton">
    <UButton @click="setWorkflowPageDefault">Cancel</UButton>
  </div>
  <div class="flex flex-col justify-center items-center gap-10 pt-10">
    <h1 class="text-custom_size_title font-custom_weight_title pb-5">Workflow</h1>
    <div class="flex flex-col justify-center items-center w-[28%]">
      <ReActionButton title="Action" link="/workflow/actions" :is-disabled="false" />
      <div :class="['bg-black min-w-4 min-h-28', reactionButtonisDisabled ? 'bg-opacity-60' : 'bg-opacity-100']"></div>
      <ReActionButton title="Reaction" link="/workflow/reactions" :is-disabled="reactionButtonisDisabled" />
    </div>
    <div v-if="showAddButton">
      <UButton>Add</UButton>
    </div>
  </div>
</template>

<style scoped></style>
