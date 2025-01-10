<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";

defineProps<{
  type: string;
  services: ServiceInfo[];
}>();

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <div v-for="service in services" :key="service.id">
      <NuxtLink
        :to="{
          name: `workflow-${type}-service`,
          params: { service: service.id },
        }"
      >
        <UContainer
          :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
          class="flex flex-col justify-end items-center gap-5 text-white font-extrabold text-6xl p-8 rounded-custom_border_radius overflow-hidden w-[5em] h-[4.5em]"
          :style="{ backgroundColor: service.color }"
        >
          <img :src="service.icon" :alt="service.name" class="w-28 h-28 p-0">
          <h2
            class="clamp-1-line capitalize text-5xl text-center break-words w-full hover-expand-text"
          >
            {{ formatName(service.name) }}
          </h2>
        </UContainer>
      </NuxtLink>
    </div>
  </UContainer>
</template>

<style scoped>
@font-face {
  font-family: "ellipsis-font";
  src: local("DejaVu Sans");
  unicode-range: U+2026;
  size-adjust: 0%;
}

.clamp-1-line {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  word-break: break-word;
  white-space: normal;
  transition: all 1s ease-in-out;
}

.hover-expand-text:hover {
  -webkit-line-clamp: unset;
  line-clamp: unset;
  overflow: visible;
  white-space: normal;
}
</style>
