<script setup lang="ts">
interface ApiResponse {
  token?: string;
}

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const token = useCookie("token");

onMounted(() => {
  connectToService();
});

async function connectToService() {
  const route = useRoute();
  const code = route.query["code"];
  const state = route.query["state"];
  const tokenCookie = useCookie("token");

  if (!code || !state) {
    showError("Missing parameters: code or state");
    return;
  }

  try {
    const response = await Promise.race([
      $fetch<ApiResponse>("/api/auth/service/connection", {
        method: "POST",
        body: {
          service: route.params.service,
          code: code as string,
          state: state as string,
          authorization: tokenCookie.value ? `Bearer ${tokenCookie.value}` : "",
        },
      }),
      new Promise<never>((_, reject) =>
        setTimeout(() => reject(new Error("Timeout")), 5000)
      ),
    ]);

    token.value = response.token;
    isLoading.value = false;
    navigateTo("/myservices");
  } catch (error) {
    if (error instanceof Error) {
      showError(`Failed to connect to service: ${error.message}`);
    } else {
      showError("Failed to connect to service: An unknown error occurred.");
      console.error("Unexpected error:", error);
    }
  }
}

function showError(message: string) {
  errorMessage.value = message;
  console.error(message);

  setTimeout(() => {
    navigateTo("/login");
  }, 3000);
}
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <LoadingScreen
      v-if="isLoading"
      v-model:is-loading="isLoading"
      :timeout="5000"
    />

    <div
      v-if="errorMessage"
      class="text-red-600 font-bold mt-2 text-lg text-center"
    >
      {{ errorMessage }}
    </div>
  </div>
</template>

<style scoped></style>
