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

/**
 * @description Fetches the connection infos from the server.
 */
onMounted(() => {
  loadConnectionInfos();
});

/**
 * @description Loads the connection infos from the server.
 */
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

/**
 * @description Deletes the account.
 */
async function deleteAccount() {
  try {
    await $fetch("/api/auth/deleteAccount", {
      method: "POST",
      body: {
        authorization: tokenCookie.value,
      },
    });
    clearTokenAndLogout();
    navigateTo("/login");
  } catch (error) {
    throw handleErrorStatus(error);
  }
}

/**
 * @description Clears the token and logs out the user.
 */
const clearTokenAndLogout = () => {
  const tokenCookie = useCookie("token");
  tokenCookie.value = null;
};
</script>

<template>
  <div class="flex flex-col justify-center items-center w-full">
    <h1 class="py-5">Settings</h1>
    <UAlert v-if="errorMessage" color="red" variant="solid" title="ERROR":description="errorMessage" class="justify-center items-center gap-5 w-[15%]"/>
    <div
      v-else
      class="flex flex-col justify-center items-center gap-16 min-w-[60%] max-lg:max-w-[70%] max-md:max-w-[85%] max-sm:max-w-full h-full py-10 rounded-custom_border_radius bg-custom_color-bg_section"
    >
      <div
        class="flex flex-col justify-center items-center gap-12 min-w-[50%] max-lg:max-w-[70%] max-md:max-w-[85%] max-sm:max-w-[90%] px-5"
      >
        <EditableInput v-model="username" name="Username" />
        <EditableInput v-model="email" name="Email" />
      </div>
      <UButton
        class="delete-button text-white flex flex-col justify-center items-center gap-2 max-md:gap-0 px-8 py-3 max-lg:py-1"
        @click="deleteAccount"
      >
        <p>Delete</p>
        <p>Account</p>
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
