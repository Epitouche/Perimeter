<script setup lang="ts">
defineProps<{
  services: any[];
}>();
</script>

<template>
  <UContainer :ui="{ padding: 'px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center flex-wrap w-full">
    <div v-for="service in services" :key="service.id">
      <NuxtLink 
      :to="{
        name: 'workflow-actions-service',
        params: { service: service.id },
      }">
        <UContainer 
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
          :class="[`bg-custom_color-${service.name}`, 'basis-1/4 flex flex-col justify-end items-center flex-wrap gap-10 text-white font-extrabold text-6xl p-8 rounded-custom_border_radius w-[5em] h-[4.5em]']">
          <UIcon :name="`my-icons:white-${service.name}`" />
          <h2 class="capitalize">{{ service.name }}</h2>
        </UContainer>
      </NuxtLink>
    </div>
  </UContainer>
</template>
