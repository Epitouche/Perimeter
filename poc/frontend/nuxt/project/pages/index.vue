<script setup lang="ts">
definePageMeta({
    layout: 'default'
});

</script>

<template>
    <div class="flex flex-col justify-center items-center h-screen gap-28">
        <h1 class="text-9xl font-serif font-medium">Want to know what kind of cat you are?</h1>
        <UButton to="/quiz" target="_self" class="text-6xl font-serif font-semibold px-6 py-5 rounded-xl">Take Quiz!</UButton>
    </div>
</template>
