<script setup lang="ts">
import type { Area } from "@/interfaces/areas";

defineProps<{
  type: Area["action"] | Area["reaction"];
}>();

const items = [
  [
    {
      label: "Update",
      slot: "update",
    },
  ],
];

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}
</script>

<template>
  <UDropdown :items="items" :popper="{ placement: 'bottom' }" class="w-fit">
    <div
      class="capitalize self-start flex flex-row items-center gap-5 border-custom_border_width border-white rounded-custom_border_radius w-full py-2 px-4"
    >
      <img
        :src="type.service.icon"
        :alt="type.service.name"
        class="w-16 h-16 p-0"
      >
      <p class="text-5xl self-start">
        <b>{{ type.service.name }}</b
        >:
      </p>
      <p class="text-4xl self-center">{{ formatName(type.name) }}</p>
      <UButton variant="ghost" class="w-fit">
        <UIcon name="i-bytesize-chevron-bottom" class="w-12 h-12 text-white" />
      </UButton>
    </div>
    <template #update="{ item }">
      <div class="text-black text-5xl">
        {{ item.label }}
      </div>
    </template>
  </UDropdown>
</template>
