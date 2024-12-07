<template>
  <NuxtLayout>
    <div
      class="relative block w-full border-0 rounded-m text-gray-900 text-center py-2"
    >
      <template v-if="error.statusCode === 404">
        <h1 style="font-size: 6rem; font-weight: bold">404!</h1>
        <p>Sorry, that page doesn't exist.</p>
      </template>
      <template v-else>
        <h1 style="font-size: 6rem; font-weight: bold">Dang!</h1>
        <p>
          <strong>{{ error.message }}</strong>
        </p>
        <p>It looks like something broke.</p>
        <p>Sorry about that.</p>
      </template>
      <p>
        Go back to your
        <NuxtLink to="/myareas" class="font-bold underline">Area</NuxtLink>
      </p>
    </div>
  </NuxtLayout>
</template>

<script setup>
const error = useError();
</script>
