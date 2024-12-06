<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'auth'
});

const showNavBar = ref(true);
const showCancelButton = ref(false);
const reactionButtonisDisabled = ref(true);
const showCreateButton = ref(false);

const onActionSelected = () => {
  showNavBar.value = false;
  showCancelButton.value = true;
  reactionButtonisDisabled.value = false;
};

const onReactionSelected = () => {
  showCreateButton.value = true;
};

const setWorkflowPageDefault = () => {
  showNavBar.value = true;
  showCancelButton.value = false;
  reactionButtonisDisabled.value = true;
};

</script>

<template>
  <div>
    <div v-if="showNavBar" class="pb-10">
      <NavigationBar />
    </div>
    <div v-if="showCancelButton" class="pt-24 pl-28">
      <UButton
        class="bg-white text-custom_color-text text-4xl font-bold px-7 py-3 !border-custom_border_width border-custom_color-border"
        @click="setWorkflowPageDefault">Cancel</UButton>
    </div>

    <div class="flex flex-col justify-center items-center gap-10 ">
      <h1 class="text-custom_size_title font-custom_weight_title pb-5">Workflow</h1>
      <div class="flex flex-col justify-center items-center w-[28%]">
        <ReActionButton title="Action" link="/workflow/actions" :is-disabled="false" />
        <div :class="['bg-black min-w-4 min-h-28', reactionButtonisDisabled ? 'bg-opacity-60' : 'bg-opacity-100']" />
        <ReActionButton title="Reaction" link="/workflow/reactions" :is-disabled="reactionButtonisDisabled" />
      </div>
      <div v-if="showCreateButton" class="pt-10">
        <UButton class="text-5xl font-bold px-8 py-4">Create</UButton>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
