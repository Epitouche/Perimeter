<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "guest",
});
const email = ref("");
const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const signUpError = ref<string | null>(null);
const token = useCookie("token");

interface RegisterResponse {
  token: string;
  message?: string;
}

const handleSignUp = async () => {
  try {
    signUpError.value = null;

    if (
      !email.value ||
      !username.value ||
      !password.value ||
      !confirmPassword.value
    ) {
      signUpError.value = "Please fill out all fields.";
      return;
    }

    if (password.value !== confirmPassword.value) {
      signUpError.value = "Passwords do not match.";
      return;
    }

    const response = await $fetch<RegisterResponse>("/api/auth/register", {
      method: "POST",
      body: {
        email: email.value,
        username: username.value,
        password: password.value,
      },
    });

    if (response.token) {
      token.value = response.token;
    }
    navigateTo("/myareas");
  } catch (error) {
    if (error && typeof error === "object" && "data" in error) {
      const typedError = error as {
        status?: number;
        data?: { message?: string };
      };

      if (typedError.status === 409) {
        signUpError.value = "Password must be at least 8 characters long";
      }
      console.error("Sign up failed:", typedError.data);
    } else if (error instanceof Error) {
      signUpError.value = error.message || "An unknown error occurred.";
      console.error("Sign up failed:", error.message);
    } else {
      signUpError.value = "An unknown error occurred.";
      console.error("Unexpected error:", error);
    }
  }
};
</script>

<template>
  <div class="flex justify-center items-center w-full h-screen">
    <UContainer :ui="{
      padding: '!px-0',
      constrained: 'min-w-[40%] max-w-[40%]',
    }"
      class="bg-custom_color-bg_section flex flex-col justify-between items-center gap-2 p-10 pb-14 rounded-custom_border_radius">
      <h1 class="pb-2">Sign up</h1>
      <div class="flex flex-col items-center gap-6 overflow-y-scroll max-h-[60vh]">
        <div class="flex flex-col max-w-[90%]">
          <h6>Email</h6>
          <UInput v-model="email" :ui="{
            placeholder: '!px-5 !py-3 font-light',
            size: { sm: 'text-4xl' },
          }" />
        </div>
        <div class="flex flex-col">
          <h6>Username</h6>
          <UInput v-model="username" :ui="{
            placeholder: '!px-5 !py-3 font-light',
            size: { sm: 'text-4xl' },
          }" />
        </div>
        <div class="flex flex-col">
          <h6>Password</h6>
          <UInput v-model="password" type="password" :ui="{
            placeholder: '!px-5 !py-3 font-light',
            size: { sm: 'text-4xl' },
          }" />
        </div>
        <div class="flex flex-col">
          <h6>Confirm Password</h6>
          <UInput v-model="confirmPassword" type="password" :ui="{
            placeholder: '!px-5 !py-3 font-light',
            size: { sm: 'text-4xl' },
          }" />
        </div>
        <div class="flex flex-col justify-center gap-1 items-center min-w-full pt-4">
          <div v-if="signUpError" class="text-red-500 text-xl pb-1">
            {{ signUpError }}
          </div>
          <UButton class="text-center text-[2.5rem] px-12" tabindex="0" @click="handleSignUp">Sign up
          </UButton>
          <p class="text-xl">
            Already signed up?
            <ULink to="/login" class="hover:text-custom_color-text_link" tabindex="0">
              <u>Login</u>
            </ULink>
          </p>
        </div>
        <div class="flex flex-col justify-center items-center gap-2 min-w-[85%] max-w-[85%] pt-2">
          <UDivider size="xs" label="or sign up with" :ui="{ label: 'text-custom_color-text_other text-xl' }" />
          <ConnectWithAppContainer />
        </div>
      </div>
    </UContainer>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
