<script setup lang="ts">
defineProps<{
  searchQuery: string;
}>();

const _emit = defineEmits(["update:searchQuery"]);
</script>

<template>
  <div
    class="flex flex-row justify-between min-w-full max-w-full border-custom_border_width border-custom_color-border rounded-custom_border_radius bg-custom_color-input"
  >
    <UInput
      :model-value="searchQuery"
      variant="none"
      :ui="{ size: { sm: 'text-4xl' }, base: 'border-none' }"
      class="rounded-l-custom_border_radius overflow-hidden"
      @input="$emit('update:searchQuery', $event.target.value)"
    />
    <div 
      class="flex justify-center items-center bg-black rounded-r-custom_border_radius"
      style="width: 20%; min-width: 1vw; max-width: 3vw;">
      <UIcon name="i-bytesize-search" class="round_icon text-white" />
    </div>
  </div>
</template>

<style scoped>
.round_icon {
  width: 50%;
  height: 50%;
}
</style>
