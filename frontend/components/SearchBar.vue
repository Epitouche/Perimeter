<script setup lang="ts">
import { defineProps } from "vue";
</script>

<template>
    <UContainer :ui="{ padding: 'p-0' }">
        <div class="border-custom_border_width border-custom_color-border rounded-custom_border_radius bg-custom_color-input flex flex-row justify-between w-full">
            <div class="w-full">
                <UInput :ui="{ base: 'w-full border-none' }" />
            </div>
            <div class="bg-black border-custom_border_width border-custom_color-border rounded-custom_border_radius flex justify-center items-center px-2">
                <UIcon name="i-bytesize-search" class="text-white w-6 h-6" />
            </div>
        </div>
    </UContainer>
</template>
