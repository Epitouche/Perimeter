<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const serviceId = route.params.service;
const token = useCookie("token");
const isLoading = ref(true);
const actions = ref<any>(null);
const errorMessage = ref<string | null>(null);

const serviceInfo = ref<{ name: string } | null>(null);

const getServiceInfo = async () => {
  if (serviceId) {
    isLoading.value = true;
    try {
      errorMessage.value = null;
      serviceInfo.value = await $fetch("/api/servicebyid", {
        method: "POST",
        body: {
          token: token.value,
          serviceId: serviceId,
        },
      });
      // console.log("services", serviceInfo.value);
    } catch (error: unknown) {
      errorMessage.value = handleErrorStatus(error);

      if (errorMessage.value === "An unknown error occurred") {
        console.error("An unknown error occurred", error);
      }
    } finally {
      isLoading.value = false;
    }
  }
};

const fetchActions = async () => {
  isLoading.value = true;
  try {
    errorMessage.value = null;
    actions.value = await $fetch("/api/workflow/actions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });

    console.log("actions", actions.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  getServiceInfo();
  fetchActions();
});
</script>

<template>
  <div class="flex flex-col gap-20">
    <div v-if="errorMessage">
      <div>Error: {{ errorMessage }}</div>
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
    <UContainer
      v-else-if="serviceInfo"
      :ui="{ constrained: 'max-w-none' }"
      :class="[`bg-custom_color-${serviceInfo.name}`, 'pt-16 pb-16']"
    >
      <div class="px-20">
        <BackButton link="/workflow/actions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-8">
        <h1 class="text-8xl text-white font-custom_weight_title">
          Add an action
        </h1>
        <UIcon
          :name="`my-icons:white-${serviceInfo.name}`"
          class="w-[8em] h-[8em]"
        />
        <h2 class="capitalize text-white text-7xl font-bold">
          {{ serviceInfo.name }}
        </h2>
      </div>
    </UContainer>
    <div v-if="errorMessage">
      <div>Error: {{ errorMessage }}</div>
    </div>
    <div v-else-if="actions">
      <ReActionCardContainer
        type-name="action"
        :types="actions"
        :service-info="serviceInfo"
      />
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
  </div>
</template>

<style scoped></style>
