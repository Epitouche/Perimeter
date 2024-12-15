<script setup lang="ts">
defineProps<{
    link: string;
}>();
</script>

<template>
    <UButton
        :to="link"
        class="text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border px-0 py-1">
        <UIcon name="heroicons:chevron-left-16-solid" class="w-[4em] h-[4em]" />
        <div class="text-4xl font-bold pr-5">Back</div>
    </UButton>
</template>

<style scoped></style>
