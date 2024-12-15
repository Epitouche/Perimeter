<script setup lang="ts">
defineProps<{
  searchQuery: string;
}>();

const _emit = defineEmits(["update:searchQuery"]);

</script>

<template>
  <div
    class="border-custom_border_width border-custom_color-border rounded-custom_border_radius bg-custom_color-input flex flex-row justify-between w-fit">
      <UInput :model-value="searchQuery" variant="none" :ui="{ size: { sm: 'text-4xl' }, base: 'w-full border-none' }" class="w-full" @input="$emit('update:searchQuery', $event.target.value)" />
    <div
      class="bg-black border-8 border-custom_color-border rounded-full flex justify-center items-center px-2">
      <UIcon name="i-bytesize-search" class="text-white w-6 h-6" />
    </div>
  </div>
</template>
