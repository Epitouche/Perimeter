<script setup lang="ts">

interface ApiResponse {
  token?: string;
}

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const token = useCookie('token');

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
    console.log("Testing API")
    const response = await $fetch<ApiResponse>('/api/auth/service/connection', {
      method: 'POST',
      body: {
        service: route.params.service,
        code: code as string,
        state: state as string,
      },
    });
    new Promise((_, reject) => setTimeout(() => reject(new Error('Timeout')), 5000)),
    console.log("Testing answer API")
    console.log("Service token is  : ", response.token);
    token.value = response.token;
    navigateTo('/');
  } catch (error: any) {
    showError(`Failed to connect to service: ${error.message}`);
  } finally {
    isLoading.value = false;
  }
}

function showError(message: string) {
  errorMessage.value = message;
  console.error(message);

  setTimeout(() => {
    navigateTo('/login');
  }, 3000);
}

</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>

    <div v-if="errorMessage" class="text-red-600 font-bold mt-2 text-lg text-center">
      {{ errorMessage }}
    </div>
  </div>
</template>

<style scoped></style>
