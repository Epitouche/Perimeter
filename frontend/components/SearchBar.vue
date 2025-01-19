<script setup lang="ts">
/**
 * @definition A simple search bar with specific styling
 */

/**
 * @property {string} searchQuery - The search query
 */
defineProps<{
  searchQuery: string; // The search query
}>();

/**
 * Emits the search query
 */
const _emit = defineEmits(["update:searchQuery"]);
</script>

<template>
  <div
    class="flex flex-row justify-between w-full border-custom_border_width border-custom_color-border rounded-custom_border_radius bg-custom_color-input"
  >
    <UInput
      :model-value="searchQuery"
      variant="none"
      :ui="{
        size: { sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl' },
        base: 'border-none',
      }"
      class="rounded-l-custom_border_radius overflow-hidden"
      @input="$emit('update:searchQuery', $event.target.value)"
    />
    <div
      class="flex justify-center items-center bg-black rounded-r-custom_border_radius"
      style="width: 15%"
    >
      <UIcon name="i-bytesize-search" class="w-[65%] h-[65%] text-white" />
    </div>
  </div>
</template>

<style scoped></style>
