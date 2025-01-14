<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "guest",
});

const username = ref("");
const password = ref("");

const token = useCookie("token");
const loginError = ref<string | null>(null);

interface RegisterResponse {
  token: string;
  message?: string;
}

const handleLogin = async () => {
  try {
    loginError.value = null;

    if (!username.value || !password.value) {
      loginError.value = "Please fill out all fields.";
      return;
    }

    const response = await $fetch<RegisterResponse>("/api/auth/login", {
      method: "POST",
      body: {
        username: username.value,
        password: password.value,
      },
    });

    if (response.token) {
      token.value = response.token;
    }
    //console.log("Login successful:", response);
    navigateTo("/myareas");
  } catch (error) {
    if (error && typeof error === "object" && "data" in error) {
      const typedError = error as {
        status?: number;
        data?: { message?: string };
      };
      if (typedError.status === 401) {
        loginError.value = "Login failed. Please try again.";
      }
      console.error("Login error:", typedError);
    } else {
      loginError.value = "An unknown error occurred.";
      console.error("Unexpected error:", error);
    }
  }
};
</script>

<template>
  <div class="flex justify-center items-center h-screen w-screen">
    <UContainer
      :ui="{
        padding: 'pt-8 pb-16 px-0',
        constrained: 'min-w-[30%] max-w-[60%]',
      }"
      class="scale-[0.75] bg-custom_color-bg_section flex flex-col justify-between items-center gap-14 rounded-custom_border_radius"
    >
      <h1
        class="pb-5"
      >
        Log in
      </h1>
      <div class="flex flex-col gap-12 min-w-[80%] max-w-[80%] px-5">
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Username</h2>
          <UInput
            v-model="username"
            :ui="{
              placeholder: '!px-5 !py-3 font-light',
              size: { sm: 'text-5xl' },
            }"
          />
        </div>
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Password</h2>
          <UInput
            v-model="password"
            type="password"
            :ui="{
              placeholder: '!px-5 !py-3 font-light',
              size: { sm: 'text-5xl' },
            }"
          />
        </div>
        <div class="flex flex-col justify-center items-center min-w-full pt-4">
          <div v-if="loginError" class="text-red-500 text-xl pb-1">
            {{ loginError }}
          </div>
          <UButton
            class="text-center text-[2.5rem] px-12"
            tabindex="0"
            @click="handleLogin"
            >Log in</UButton
          >
          <p class="text-xl">
            New?
            <ULink
              to="/signup"
              class="hover:text-custom_color-text_link"
              tabindex="0"
            >
              <u>Sign Up</u>
            </ULink>
          </p>
        </div>
      </div>
      <div class="min-w-[80%] max-w-[80%] pt-2">
        <UDivider
          size="xs"
          label="or log in with"
          :ui="{ label: 'text-custom_color-text_other text-xl' }"
        />
      </div>
      <ConnectWithAppContainer />
    </UContainer>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
