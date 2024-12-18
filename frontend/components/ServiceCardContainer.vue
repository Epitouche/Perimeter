<script setup lang="ts">
defineProps<{
  type: string;
  services: any[];
}>();

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, '$1 $2');
}
</script>

<template>
  <UContainer 
  :ui="{ padding: 'px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full">
    <div v-for="service in services" :key="service.id">
      <NuxtLink 
      :to="{
        name: `workflow-${type}-service`,
        params: { service: service.id },
      }">
        <UContainer 
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }" :class="[
          `bg-custom_color-${service.name}`,
          'basis-1/4 flex flex-col justify-end items-center gap-10 text-white font-extrabold text-6xl p-8 rounded-custom_border_radius w-[5em] h-[4.5em]'
        ]">
          <UIcon :name="`my-icons:white-${service.name}`" />
          <h2 class="clamp-2-lines capitalize text-5xl text-center break-words w-full">
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

.clamp-2-lines {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: clip;
  word-break: break-word;
  white-space: normal;
  font-family: ellipsis-font;
}
</style>
