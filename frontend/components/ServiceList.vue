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
  <UContainer
    class="flex flex-wrap gap-5 justify-center p-4 bg-white rounded-lg mx-auto"
  >
    <UButton
      v-for="(app, index) in apps"
      :key="index"
      :icon="app.icon"
      class="app_button flex flex-col items-center justify-center w-[15rem] h-[15rem] rounded-lg transition-transform hover:scale-105"
      :style="{ backgroundColor: app.color }"
      @click="handleClick(app.name)"
    >
      <span class="text-3xl text-white font-bold mt-auto">{{ app.name }}</span>
    </UButton>
  </UContainer>
</template>

<style scoped>
:deep(.app_button span) {
  height: 6rem;
  width: 6rem;
  color: white;
}
</style>
