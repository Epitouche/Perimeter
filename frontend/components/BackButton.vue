<script setup lang="ts">
defineProps<{
  link: string;
  isWhite: boolean;
}>();
</script>

<template>
  <UButton
    :to="link"
    :class="[
      'bg-opacity-0 border-custom_border_width px-0 py-1 w-fit',
      isWhite ? '!border-white text-white' : '!border-black text-black',
    ]"
  >
    <UIcon name="heroicons:chevron-left-16-solid" class="w-[4em] h-[4em]" />
    <div class="text-4xl font-bold pr-5">Back</div>
  </UButton>
</template>

<style scoped></style>
