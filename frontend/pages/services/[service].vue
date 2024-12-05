<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, navigateTo } from 'nuxt/app';

interface ApiResponse {
    token: string;
}

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);

onMounted(() => {
  connectToService();
})

async function connectToService() {
  const route = useRoute();
  const code = route.query['code'];
  const state = route.query['state'];

  if (!code || !state) {
    showError('Missing parameters: code or state');
    return;
  }

  try {
    console.log(`Connecting to service with code=${code} and state=${state}...`);
    const response = await $fetch<ApiResponse>('/api/auth/service/connection', {
      method: 'POST',
      body: {
        service: route.params.service,
        code: code as string,
        state: state as string,
      },
    });
    new Promise((_, reject) => setTimeout(() => reject(new Error('Timeout')), 5000)),
    console.log("Service connected  : ", response);

  } catch (error:any) {
    showError(`Failed to connect to service: ${error.message}`);
  } finally {
    isLoading.value = false;
  }
}

function showError(message: string) {
  errorMessage.value = message;
  console.error(message);

  setTimeout(() => {
    navigateTo('/error');
  }, 3000);
}

</script>

<template>
  <div>
    <div v-if="isLoading">Loading...</div>

    <div v-if="errorMessage" class="error">
      {{ errorMessage }}
    </div>
  </div>
</template>

<style>
.error {
  color: red;
  font-weight: bold;
  margin-top: 10px;
}
</style>

