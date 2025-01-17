<script lang="ts" setup>
import type { ServiceResponse } from "~/interfaces/serviceResponse";
import EditableInput from "~/components/EditableInput.vue";

definePageMeta({
  middleware: "auth",
});

const errorMessage = ref<string | null>(null);
const tokenCookie = useCookie("token");
const username = ref<string>("");
const email = ref<string>("");
const userId = ref<number>(0);
const infosConnection = ref<ServiceResponse | null>(null);

onMounted(() => {
  loadConnectionInfos();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);
      username.value = infosConnection.value.user.username;
      email.value = infosConnection.value.user.email;
      userId.value = infosConnection.value.user.id;
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

async function deleteAccount() {
  try {
    await $fetch("/api/auth/deleteAccount", {
      method: "POST",
      body: {
        authorization: tokenCookie.value,
      }
    });
    clearTokenAndLogout();
    navigateTo('/login');
  } catch (error) {
    throw handleErrorStatus(error);
  }
}

const clearTokenAndLogout = () => {
  const tokenCookie = useCookie("token");
  tokenCookie.value = null;
};

</script>

<template>
  <div class="flex flex-col justify-center items-center gap-10 w-full">
    <h1>Settings</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div
      v-else
      class="flex flex-col justify-center items-center gap-10 w-[60%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
    >
      <div class="flex flex-col justify-center items-center gap-8 w-full px-5">
        <EditableInput v-model="username" name="Username" />
        <EditableInput v-model="email" name="Email" />
      </div>
      <UButton 
        class="delete-button flew justify-center items-center rounded-2x1 px-10 py-5 w-[20%] font-bold text-white"
        @click=deleteAccount>
        <h6>Delete Account</h6>
      </UButton>
    </div>
  </div>
</template>

<style scoped>
.delete-button {
  background-color: #ff0000;
}

.delete-button:hover {
  background-color: #dc2626;
}

[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
