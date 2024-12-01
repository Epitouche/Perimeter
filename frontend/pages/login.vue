<script setup lang="ts">
const username = ref('')
const password = ref('')

const apps = ref<string[]>(['i-logos-google-icon', 'i-logos-google-icon', 'i-logos-google-icon']);

const loginError = ref<string | null>(null);

const handleLogin = async () => {
  try {
    loginError.value = null;

    const response = await $fetch('http://localhost:8080/api/v1/auth/login', {
      method: 'POST',
      body: {
        username: username.value,
        password: password.value,
      },
    });
    console.log('Login successful:', response);
  } catch (error: any) {
    console.error('Login failed:', error);
    loginError.value = error?.data?.message || 'Login failed. Please try again.';
  }
};
</script>

<template>
  <div class="flex justify-center items-center h-screen w-screen">
    <UContainer :ui="{ padding: 'pt-8 pb-16 px-0', constrained: 'min-w-[30%] max-w-[80%]' }"
      class="bg-custom_color-bg_section flex flex-col justify-between items-center gap-14 rounded-custom_border_radius">
      <h1 class="text-custom_size_title font-custom_weight_connection_title pb-5">Log in</h1>
      <div class="flex flex-col gap-12 min-w-[80%] max-w-[80%] px-5">
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Username</h2>
          <UInput v-model="username" :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Password</h2>
          <UInput type="password" v-model="password"
            :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
          <ULink to="/forgotpassword" class="text-xl text-custom_color-text_link self-end px-5">Forgot password?</ULink>
        </div>
        <div class="flex flex-col justify-center items-center min-w-full pt-4">
          <UButton class="text-center text-[2.5rem] px-12">Log in</UButton>
          <p class="text-xl">New? <ULink to="/signup" class="hover:text-custom_color-text_link"><u>Sign Up</u></ULink>
          </p>
        </div>
      </div>
      <div class="min-w-[80%] max-w-[80%] pt-2">
        <UDivider size="xs" label="or log in with" :ui="{ label: 'text-custom_color-text_other text-xl' }" />
      </div>
      <ConnectWithAppContainer :apps="apps" />
    </UContainer>
  </div>
</template>

<style scoped></style>