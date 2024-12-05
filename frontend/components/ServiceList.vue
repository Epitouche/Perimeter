<script setup lang="ts">
defineProps<{
  apps: {
    name: string;
    color: string;
    icon: string;
  }[];
}>();


const handleClick = (name: string) => {
  console.log(`${name} clicked`);
};
</script>

<template>
  <UContainer class="flex flex-wrap gap-5 justify-center p-4 bg-white rounded-lg mx-auto">
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :icon="app.icon"
      variant="ghost"
      class="flex flex-col items-center justify-center w-[15rem] h-[15rem] rounded-lg transition-transform hover:scale-105"
      :style="{ backgroundColor: app.color }"
      @click="handleClick(app.name)"
    >
      <span class="text-white text-4xl font-bold mb-2">{{ app.name }}</span>
    </UButton>
  </UContainer>
</template>

<style scoped>
</style>
