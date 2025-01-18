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
  } catch (error: unknown) {
    loginError.value = handleErrorStatus(error);
  }
};
</script>

<template>
  <div class="flex justify-center items-center w-full h-screen">
    <UContainer
      :ui="{
        padding: '!px-0 !pt-8 !pb-12 max-md:!pb-8 max-md:!px-2',
        constrained:
          'min-w-[40%] max-w-[95%] max-lg:min-w-[50%] max-md:min-w-[60%] max-sm:min-w-[80%]',
      }"
      class="bg-custom_color-bg_section flex flex-col justify-between items-center gap-12 max-lg:gap-10 max-md:gap-8 max-sm:gap-6 rounded-custom_border_radius"
    >
      <h1 class="pb-2">Log in</h1>
      <div
        class="flex flex-col items-center gap-14 max-lg:gap-10 max-md:gap-8 overflow-y-scroll max-h-[60vh]"
      >
        <div class="flex flex-col min-w-[70%] max-w-[80%]">
          <h6 class="px-5">Username</h6>
          <UInput
            v-model="username"
            :ui="{
              placeholder: '!px-5 !py-4 max-lg:!py-2 max-md:!py-1 font-light',
              size: {
                sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl',
              },
            }"
          />
        </div>
        <div class="flex flex-col min-w-[70%] max-w-[80%]">
          <h6 class="px-5">Password</h6>
          <UInput
            v-model="password"
            type="password"
            :ui="{
              placeholder: '!px-5 !py-4 max-lg:!py-2 max-md:!py-1 font-light',
              size: {
                sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl',
              },
            }"
          />
        </div>
        <div
          class="flex flex-col justify-center gap-1 items-center min-w-full py-5 max-lg:py-4 max-md:py-3 max-sm:py-2"
        >
          <div v-if="loginError" class="text-red-500 text-xl pb-1">
            {{ loginError }}
          </div>
          <UButton
            class="text-center px-10 py-7 max-lg:py-5 max-sm:py-3"
            tabindex="0"
            @click="handleLogin"
          >
            <h5>Log in</h5>
          </UButton>
          <p>
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
        <div
          class="flex flex-col justify-center items-center gap-2 min-w-[85%] max-w-[85%] max-md:min-w-[80%] max-md:max-w-[80%] pt-2 max-md:pt-0"
        >
          <UDivider
            size="xs"
            label="or sign up with"
            :ui="{ label: 'text-custom_color-text_other text-xl' }"
          />
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
