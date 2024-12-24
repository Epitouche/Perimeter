<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const serviceId = route.params.service;
const token = useCookie("token");
const isLoading = ref(true);
const reactions = ref<any>(null);
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

const fetchReactions = async () => {
  isLoading.value = true;
  try {
    errorMessage.value = null;
    reactions.value = await $fetch("/api/workflow/reactions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });

    console.log("reactions", reactions.value);
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
  fetchReactions();
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
      :class="[`bg-custom_color-${serviceInfo.name}`, 'py-20']"
    >
      <div class="px-20">
        <BackButton link="/workflow/reactions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-2">
        <h1 class="text-8xl text-white font-custom_weight_title">
          Add a reaction
        </h1>
        <UIcon
          :name="`my-icons:white-${serviceInfo.name}`"
          class="w-[9em] h-[9em]"
        />
        <h2 class="capitalize text-white text-7xl font-bold pt-8">
          {{ serviceInfo.name }}
        </h2>
      </div>
    </UContainer>
    <div v-if="errorMessage">
      <div>Error: {{ errorMessage }}</div>
    </div>
    <div v-else-if="reactions">
      <ReActionCardContainer
        type-name="reaction"
        :types="reactions"
        :service-info="serviceInfo"
      />
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
  </div>
</template>

<style scoped></style>
