<script setup lang="ts">
defineProps<{
  link: string;
  color: string;
}>();
</script>
<!-- :class="[`text-${color}`, `!border-${color}`, 'bg-opacity-0 border-custom_border_width px-0 py-1 w-fit']" -->
<template>
  <UButton
    :to="link"
    class="text-black !border-black bg-opacity-0 border-custom_border_width px-0 py-1 w-fit"
  >
    <UIcon name="heroicons:chevron-left-16-solid" class="w-[4em] h-[4em]" />
    <div class="text-4xl font-bold pr-5">Back</div>
  </UButton>
</template>

<style scoped></style>
