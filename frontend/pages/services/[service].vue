<script setup lang="ts">
interface ApiResponse {
  token?: string;
}

const isLoading = ref(true);
const errorMessage = ref<string | null>(null);
const token = useCookie("token");
const loadingPath = "/myservices";

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
        setTimeout(() => reject(new Error("Timeout")), 5000),
      ),
    ]);

    token.value = response.token;
    navigateTo("/workflow");
  } catch (error) {
    if (error instanceof Error) {
      showError(`Failed to connect to service: ${error.message}`);
    } else {
      showError("Failed to connect to service: An unknown error occurred.");
      console.error("Unexpected error:", error);
    }
  } finally {
    isLoading.value = false;
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
    <LoadingScreen v-if="isLoading" :path="loadingPath" :timeout="3000" />

    <div
      v-if="errorMessage"
      class="text-red-600 font-bold mt-2 text-lg text-center"
    >
      {{ errorMessage }}
    </div>
  </div>
</template>

<style scoped></style>
