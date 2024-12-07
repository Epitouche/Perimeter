<script setup lang="ts">
defineProps<{
  title: string,
  link: string,
  isDisabled: boolean,
  isSelected: boolean
}>();

</script>

<template>
    <div v-if="isSelected" :class="['flex flex-row justify-evenly items-center bg-black py-12 rounded-3xl w-full', isDisabled ? 'bg-opacity-60' : 'bg-opacity-100']">
        <h2 :class="['text-white text-8xl font-custom_weight_title', isDisabled ? 'text-opacity-50' : 'text-opacity-100']">{{ title }}</h2> <!-- Change later -->
    </div>
    <div v-else :class="['flex flex-row justify-evenly items-center bg-black py-12 rounded-3xl w-full', isDisabled ? 'bg-opacity-60' : 'bg-opacity-100']">
        <h2 :class="['text-white text-8xl font-custom_weight_title', isDisabled ? 'text-opacity-50' : 'text-opacity-100']">{{ title }}</h2>
        <UButton :disabled="isDisabled" :to="link" :ui="{ rounded: 'rounded-2xl' }" :class="['text-black bg-white text-5xl font-extrabold px-10 py-3', isDisabled ? '!text-opacity-60' : 'text-opacity-100']">Add</UButton>
    </div>
</template>

<style scoped></style>
