<script setup lang="ts">
defineProps<{
  link: string;
  isWhite: boolean;
}>();
</script>

<template>
  <UButton
    variant="ghost"
    :to="link"
    :class="[
      'border-custom_border_width flex flex-row justify-between items-center max-w-[7vw] pr-4 py-3',
      isWhite ? '!border-white text-white' : '!border-black text-black',
    ]"
  >
    <UIcon name="i-bytesize-chevron-left" class="w-[2.5vw] h-[2.5vh]" />
    <h6>Back</h6>
  </UButton>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
