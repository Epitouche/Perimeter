<script setup lang="ts">
/**
 * @description This page is used to connect to a service.
 * It is called by the service provider after the user has authorized the connection.
 *
 * @param {string} service - The service to connect to.
 */
interface ApiResponse {
  token?: string;
}

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const token = useCookie("token");

/**
 * @description Connect to the service.
 */
onMounted(() => {
  connectToService();
});

/**
 * @description Connect to the service.
 */
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

/**
 * @description Show an error message and redirect to the login page.
 * @param {string} message - The error message to show.
 */
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
