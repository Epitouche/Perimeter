<script setup lang="ts">
import type { LocationQueryValue } from "vue-router";
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const websiteStore = useWebsiteStore();
const token = useCookie("token");
const router = useRouter();
const route = useRoute();
const error = ref<string | null>(null);
const createdMessage = ref<string | null>(null);
const errorMessage = ref<string | null>(null);
const showPageContent = ref(true);
const isLoading = ref(false);

const onCreate = async () => {
  console.log("actionId:", websiteStore.actionId);
  console.log("actionOptions:", websiteStore.actionOptions);
  console.log("reactionId:", websiteStore.reactionId);
  console.log("reactionOptions:", websiteStore.reactionOptions);

  error.value = null;

  try {
    if (
      typeof websiteStore.actionOptions !== "object" ||
      typeof websiteStore.reactionOptions !== "object"
    ) {
      console.log("options are not JSON objects.");
    } else {
      console.log("options are json objects");
    }

    const response = await $fetch("/api/workflow/create", {
      method: "POST",
      body: {
        token: token.value,
        actionOptions: websiteStore.actionOptions,
        actionId: websiteStore.actionId,
        reactionOptions: websiteStore.reactionOptions,
        reactionId: websiteStore.reactionId,
      },
    });
    console.log("response:", response);
    createdMessage.value = "Workflow created successfully!";
    showPageContent.value = false;
    setTimeout(() => {
      createdMessage.value = null;
      showPageContent.value = true;
    }, 500);
    websiteStore.resetWorkflowPage();
    router.push("/workflow");
  } catch (error: unknown) {
    console.log("error:", error);
    errorMessage.value = handleErrorStatus(error);
    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    alert("An error occurred while creating the workflow, please try again");
    websiteStore.resetWorkflowPage();
    router.push("/workflow");
  }
};

const onCancel = () => {
  websiteStore.resetWorkflowPage();
  router.push("/workflow");
};

onMounted(() => {
  isLoading.value = true;
  try {
    websiteStore.loadWorkflowState();
  } catch (err) {
    console.error("Error loading services:", err);
  } finally {
    isLoading.value = false;
  }

  const getQueryParam = (
    param: LocationQueryValue | LocationQueryValue[] | undefined,
  ): string | null => {
    if (Array.isArray(param)) {
      return param.length > 0 ? String(param[0]) : null;
    }
    return param ? String(param) : null;
  };

  const actionId = getQueryParam(route.query.actionId);
  const reactionId = getQueryParam(route.query.reactionId);

  if (actionId) {
    websiteStore.actionId = actionId;
    try {
      websiteStore.actionOptions = JSON.parse(
        route.query.actionOptions ? String(route.query.actionOptions) : "{}",
      );
    } catch (err) {
      console.error(
        "Failed to parse actionOptions:",
        route.query.actionOptions,
        err,
      );
      websiteStore.actionOptions = {};
    }
    websiteStore.actionServiceId = getQueryParam(route.query.actionServiceId);
    websiteStore.onActionSelected();
  }

  if (reactionId) {
    websiteStore.reactionId = reactionId;
    try {
      websiteStore.reactionOptions = JSON.parse(
        route.query.reactionOptions
          ? String(route.query.reactionOptions)
          : "{}",
      );
    } catch (err) {
      console.error(
        "Failed to parse reactionOptions:",
        route.query.reactionOptions,
        err,
      );
      websiteStore.reactionOptions = {};
    }
    websiteStore.reactionServiceId = getQueryParam(
      route.query.reactionServiceId,
    );
    websiteStore.onReactionSelected();
  }

  const cleanUrl = window.location.pathname;
  window.history.replaceState({}, "", cleanUrl);
});
</script>

<template>
  <div>
    <div v-if="createdMessage" class="flex justify-center items-center text-7xl font-bold h-screen w-screen">
      {{ createdMessage }}
    </div>
    <div v-if="showPageContent">
      <div v-if="websiteStore.showNavBar" class="pb-10">
        <NavigationBar />
      </div>
      <div v-if="websiteStore.showCancelButton" class="pt-24 pl-28">
        <UButton
          class="bg-white text-custom_color-text text-4xl font-bold px-7 py-3 !border-custom_border_width border-custom_color-border"
          @click="onCancel()">Cancel</UButton>
      </div>

      <div class="flex flex-col justify-center items-center gap-10">
        <h1 class="text-custom_size_title font-custom_weight_title pb-5">
          Workflow
        </h1>
        <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>
        <div class="flex flex-col justify-center items-center">
          <ReActionButton
title="Action" link="/workflow/actions" :is-disabled="false"
            :is-selected="websiteStore.actionIsSelected" :service-id="Number(websiteStore.actionServiceId)" />
          <div
:class="[
            'bg-black min-w-4 min-h-28',
            websiteStore.reactionButtonisDisabled
              ? 'bg-opacity-60'
              : 'bg-opacity-100',
          ]" />
          <ReActionButton
title="Reaction" link="/workflow/reactions"
            :is-disabled="websiteStore.reactionButtonisDisabled" :is-selected="websiteStore.reactionIsSelected"
            :service-id="Number(websiteStore.reactionServiceId)" />
        </div>
        <div v-if="websiteStore.showCreateButton" class="pt-10">
          <UButton class="text-5xl font-bold px-8 py-4" @click="onCreate">Create</UButton>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
