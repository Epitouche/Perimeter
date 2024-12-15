<script setup lang="ts"></script>

<template>
  <div
    class="border-custom_border_width border-custom_color-border rounded-custom_border_radius bg-custom_color-input flex flex-row justify-between w-fit">
      <UInput :ui="{ base: 'w-full border-none' }" />
    <div
      class="bg-black border-custom_border_width border-custom_color-border rounded-custom_border_radius flex justify-center items-center px-1">
      <UIcon name="i-bytesize-search" class="text-white w-6 h-6" />
    </div>
  </div>
</template>
