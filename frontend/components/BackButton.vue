<script setup lang="ts">
/**
 * The page to navigate to when the button is clicked.
 */
defineProps<{
  link: string; // The link to the page
  isWhite: boolean; // Whether the button is white
}>();
</script>

<template>
  <UButton
    variant="ghost"
    :to="link"
    :class="[
      'border-custom_border_width flex items-center w-fit px-6 py-4 max-lg:py-3 max-md:py-2 max-md:px-4',
      isWhite ? '!border-white text-white' : '!border-black text-black',
    ]"
  >
    <h6>Back</h6>
  </UButton>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
