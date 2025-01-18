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
      'border-custom_border_width flex flex-row justify-between items-center gap-2 w-fit p-4',
      isWhite ? '!border-white text-white' : '!border-black text-black',
    ]"
  >
    <UIcon name="i-bytesize-chevron-left" class="w-[2vh] h-[2vh]" />
    <h6>Back</h6>
  </UButton>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
