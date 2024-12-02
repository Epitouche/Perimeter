<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps<{
  apps: string[]
}>();

const handleClick = (label: string) => {
  console.log("app icon clicked");
};

</script>

<template>
  <UContainer :ui="{ padding: 'px-0' }" class="bgbg-custom_color-bg_section min-w-full flex flex-wrap justify-between">
    <UButton variant="ghost" v-for="(app, index) in apps" :key="index" @click="handleClick(app)" :icon="app" class="app_button basis-1/3 flex justify-center" />
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 5rem;
  width: 5rem;
}
</style>
