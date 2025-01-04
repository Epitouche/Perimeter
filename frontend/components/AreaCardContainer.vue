<script setup lang="ts">
import type { Area } from "@/interfaces/areas";

defineProps<{
  areas: Area[];
}>();

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full max-h-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap py-5 w-full h-full"
  >
    <div v-for="area in areas" :key="area.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        :class="[
          `bg-custom_color-${area.action.service.name}`,
          'flex justify-center items-center text-white font-extrabold text-6xl rounded-custom_border_radius w-[5em] h-[4.5em]',
        ]"
      >
        <h2
          class="clamp-2-lines capitalize text-5xl text-center break-words w-full"
        >
          {{ formatName(area.action.name) }}
        </h2>
      </UContainer>
    </div>
  </UContainer>
</template>

<style scoped></style>
