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
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-10 w-full">
    <h1 class="text-custom_size_title font-custom_weight_title">Settings</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div
      v-else
      class="flex flex-col justify-center items-center gap-10 w-[60%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
    >
      <div class="flex flex-col justify-center items-center gap-8 w-full px-5">
        <EditableInput v-model="username" name="Username" />
        <EditableInput v-model="email" name="Email" />
      </div>
    </div>
  </div>
</template>

<style></style>
