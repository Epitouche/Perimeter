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
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

const handleSubmit = async () => {
  const userData = {
    username: username.value,
    email: email.value,
  };

  console.log("User Data : ", userData);
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
        class="text-black bg-custom_color-bg_section border-2 border-black items-right text-2xl font-semibold py-3 px-5"
        @click="handleSubmit"
        >Submit
      </UButton>
    </div>
  </div>
</template>

<style></style>
