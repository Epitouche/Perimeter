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
  <div class="flex justify-center items-center h-screen w-screen">
    <UContainer
      :ui="{
        padding: 'pt-8 pb-16 px-0',
        constrained: 'min-w-[30%] max-w-[60%]',
      }"
      class="scale-[0.75] bg-custom_color-bg_section flex flex-col justify-between items-center gap-14 rounded-custom_border_radius"
    >
      <h1 class="pb-5">Sign up</h1>
      <div class="flex flex-col gap-12 min-w-[80%] max-w-[80%] px-5">
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Email</h2>
          <UInput
            v-model="email"
            :ui="{
              placeholder: '!px-5 !py-3 font-light',
              size: { sm: 'text-5xl' },
            }"
          />
        </div>
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
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Confirm Password</h2>
          <UInput
            v-model="confirmPassword"
            type="password"
            :ui="{
              placeholder: '!px-5 !py-3 font-light',
              size: { sm: 'text-5xl' },
            }"
          />
        </div>
        <div class="flex flex-col justify-center items-center min-w-full pt-4">
          <div v-if="signUpError" class="text-red-500 text-xl pb-1">
            {{ signUpError }}
          </div>
          <UButton
            class="text-center text-[2.5rem] px-12"
            tabindex="0"
            @click="handleSignUp"
            >Sign up
          </UButton>
          <p class="text-xl">
            Already signed up?
            <ULink
              to="/login"
              class="hover:text-custom_color-text_link"
              tabindex="0"
            >
              <u>Login</u>
            </ULink>
          </p>
        </div>
      </div>
      <div class="min-w-[80%] max-w-[80%] pt-2">
        <UDivider
          size="xs"
          label="or sign up with"
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
