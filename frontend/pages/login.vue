<script setup lang="ts">
const username = ref('')
const password = ref('')
</script>

<template>
  <div class="flex items-center justify-center h-screen w-screen">
    <UContainer :ui="{ padding: 'p-10', constrained: 'min-w-[30%] max-w-[80%]' }"
      class="bg-custom_color-bg_section flex flex-col items-center gap-10 rounded-custom_border_radius">
      <h1 class="text-custom_size_title font-custom_weight_connection_title">Log in</h1>
      <div class="flex flex-col gap-12 min-w-[80%] max-w-[80%]">
        <div class="flex flex-col">
          <h2 class="text-3xl px-5">Username</h2>
          <UInput v-model="username" :ui="{ placeholder: '!px-5 font-light', size: { sm: 'text-5xl' } }" />
        </div>
        <div class="flex flex-col">
          <h2 class="text-3xl px-5">Password</h2>
          <UInput type="password" v-model="password"
            :ui="{ placeholder: '!px-5 font-light', size: { sm: 'text-5xl' } }" />
          <ULink to="/forgotpassword" class="text-xl text-custom_color-text_link self-end px-5">Forgot password?</ULink>
        </div>
        <div class="flex flex-col justify-center items-center min-w-full">
          <UButton class="text-center text-[2.5rem] px-7">Log in</UButton>
          <p class="text-xl">New? <ULink to="/signup" class="hover:text-custom_color-text_link"><u>Sign Up</u></ULink>
          </p>
        </div>
      </div>
      <div>
        <!-- <ConnectWithApp /> -->
      </div>
    </UContainer>
  </div>
</template>

<style scoped></style>