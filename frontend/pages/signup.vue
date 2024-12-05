<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'guest'
});

const email = ref('')
const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const signUpError = ref<string | null>(null);
const token = useCookie('token')

interface RegisterResponse {
  token: string;
  message?: string;
}

const apps = ref<string[]>(['i-logos-spotify-icon', 'i-logos-google-icon']);

const handleSignUp = async () => {
  try {
    signUpError.value = null;

    if (!email.value || !username.value || !password.value || !confirmPassword.value) {
      signUpError.value = 'Please fill out all fields.';
      return;
    }

    if (password.value !== confirmPassword.value) {
      signUpError.value = 'Passwords do not match.';
      return;
    }

    const response = await $fetch<RegisterResponse>('/api/register', {
      method: 'POST',
      body: {
        email: email.value,
        username: username.value,
        password: password.value,
      },
    });

    if (response.token) {
      token.value = response.token;
      console.log('Token stored in localStorage:', response.token);
    }
    console.log('Sign up successful:', response);
    navigateTo('/myareas');
  } catch (error: any) {
    console.error('Sign up failed:', error);
    signUpError.value = error?.data?.message || 'Sign up failed. Please try again.';
  }
};

</script>

<template>
  <div class="flex justify-center items-center h-screen w-screen">
    <UContainer :ui="{ padding: 'pt-8 pb-16 px-0', constrained: 'min-w-[30%] max-w-[80%]' }"
      class="scale-[0.75] bg-custom_color-bg_section flex flex-col justify-between items-center gap-14 rounded-custom_border_radius">
      <h1 class="text-custom_size_title font-custom_weight_connection_title pb-5">Sign up</h1>
      <div class="flex flex-col gap-12 min-w-[80%] max-w-[80%] px-5">
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Email</h2>
          <UInput v-model="email" :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Username</h2>
          <UInput v-model="username" :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Password</h2>
          <UInput type="password" v-model="password"
            :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col">
          <h2 class="text-xl px-5">Confirm Password</h2>
          <UInput type="password" v-model="confirmPassword"
            :ui="{ placeholder: '!px-5 !py-3 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col justify-center items-center min-w-full pt-4">
          <div v-if="signUpError" class="text-red-500 text-xl pb-1">
            {{ signUpError }}
          </div>
          <UButton @click="handleSignUp" class="text-center text-[2.5rem] px-12">Sign up</UButton>
          <p class="text-xl">Already signed up? <ULink to="/login" class="hover:text-custom_color-text_link">
              <u>Login</u></ULink>
          </p>
        </div>
      </div>
      <div class="min-w-[80%] max-w-[80%] pt-2">
        <UDivider size="xs" label="or sign up with" :ui="{ label: 'text-custom_color-text_other text-xl' }" />
      </div>
      <ConnectWithAppContainer :apps="apps" />
    </UContainer>
  </div>
</template>

<style scoped></style>