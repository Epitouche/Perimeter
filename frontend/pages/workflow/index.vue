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
const creationPopup = ref(false);
const isLoading = ref(false);
const title = ref<string>("");
const description = ref<string>("");

const validateCreation = () => {
  creationPopup.value = !creationPopup.value;
};

const onCreate = async () => {
  //console.log("actionId:", websiteStore.actionId);
  //console.log("actionOptions:", websiteStore.actionOptions);
  //console.log("reactionId:", websiteStore.reactionId);
  //console.log("reactionOptions:", websiteStore.reactionOptions);
  //console.log("title: ", title.value);
  //console.log("description: ", description.value);

  creationPopup.value = false;
  error.value = null;

  try {
    await $fetch("/api/workflow/create", {
      method: "POST",
      body: {
        token: token.value,
        actionOptions: websiteStore.actionOptions,
        actionId: websiteStore.actionId,
        reactionOptions: websiteStore.reactionOptions,
        reactionId: websiteStore.reactionId,
        title: title.value,
        description: description.value,
      },
    });
    createdMessage.value = "Workflow created successfully!";
    showPageContent.value = false;
    setTimeout(() => {
      createdMessage.value = null;
      showPageContent.value = true;
    }, 500);
    websiteStore.resetWorkflowPage();
    router.push("/workflow");
  } catch (error: unknown) {
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

function validateOptions(
  options: Record<string, unknown>,
): Record<string, unknown> {
  return Object.fromEntries(
    Object.entries(options).map(([key, value]) => {
      if (typeof value === "string" && !isNaN(Number(value))) {
        return [key, Number(value)];
      }
      return [key, value];
    }),
  );
}

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

  if (actionId) {
    websiteStore.actionId = actionId;
    const actionName = getQueryParam(route.query.actionName);
    if (actionName) {
      websiteStore.actionName = actionName;
    }
    const actionOptionsString = getQueryParam(route.query.actionOptions);
    let actionOptions = {};

    if (actionOptionsString) {
      try {
        actionOptions = JSON.parse(actionOptionsString);
      } catch (err) {
        console.error(
          "Failed to parse actionOptions:",
          actionOptionsString,
          err,
        );
        actionOptions = {};
      }
    }
    actionOptions = validateOptions(actionOptions);
    websiteStore.actionOptions = actionOptions;
    websiteStore.actionServiceId = getQueryParam(route.query.actionServiceId);
    websiteStore.onActionSelected();
  }

  const reactionId = getQueryParam(route.query.reactionId);

  if (reactionId) {
    websiteStore.reactionId = reactionId;
    const reactionName = getQueryParam(route.query.reactionName);
    if (reactionName) {
      websiteStore.reactionName = reactionName;
    }
    const reactionOptionsString = getQueryParam(route.query.reactionOptions);
    let reactionOptions = {};

    if (reactionOptionsString) {
      try {
        reactionOptions = JSON.parse(reactionOptionsString);
      } catch (err) {
        console.error(
          "Failed to parse reactionOptions:",
          reactionOptionsString,
          err,
        );
        reactionOptions = {};
      }
    }
    reactionOptions = validateOptions(reactionOptions);
    websiteStore.reactionOptions = reactionOptions;
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
    <div
      v-if="createdMessage"
      class="flex justify-center items-center text-7xl font-bold h-screen w-screen"
    >
      {{ createdMessage }}
    </div>
    <div v-if="showPageContent">
      <div v-if="websiteStore.showNavBar" class="pb-10">
        <NavigationBar />
      </div>
      <div v-if="websiteStore.showCancelButton" class="pt-24 pl-28">
        <UButton
          class="bg-white text-custom_color-text text-4xl font-bold px-7 py-3 !border-custom_border_width border-custom_color-border"
          @click="onCancel()"
          >Cancel</UButton
        >
      </div>

      <div class="flex flex-col justify-center items-center gap-10">
        <h1 class="text-custom_size_title font-custom_weight_title pb-5">
          Workflow
        </h1>
        <div v-if="isLoading" class="text-xl font-semibold">Loading...</div>
        <div class="flex flex-col justify-center items-center">
          <ReActionButton
            title="Action"
            link="/workflow/actions"
            :is-disabled="false"
            :is-selected="websiteStore.actionIsSelected"
            :service-id="Number(websiteStore.actionServiceId)"
            :type-name="websiteStore.actionName"
          />
          <div
            :class="[
              'bg-black min-w-4 min-h-28',
              websiteStore.reactionButtonisDisabled
                ? 'bg-opacity-60'
                : 'bg-opacity-100',
            ]"
          />
          <ReActionButton
            title="Reaction"
            link="/workflow/reactions"
            :is-disabled="websiteStore.reactionButtonisDisabled"
            :is-selected="websiteStore.reactionIsSelected"
            :service-id="Number(websiteStore.reactionServiceId)"
            :type-name="websiteStore.reactionName"
          />
        </div>
        <div v-if="websiteStore.showCreateButton" class="pt-10">
          <UButton
            class="text-5xl font-bold px-8 py-4"
            @click="validateCreation"
            >Create</UButton
          >
        </div>
        <div
          v-if="creationPopup"
          class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50"
        >
          <div
            class="bg-white p-10 border-custom_border_width rounded-custom_border_radius shadow-lg max-w-md w-full"
          >
            <h2 class="text-4xl font-semibold mb-2">
              You're about to create a new area
            </h2>
            <UInput
              v-model="title"
              :ui="{
                placeholder: '!px-5 !py-3 font-light',
                size: { sm: 'text-3xl' },
              }"
              placeholder="Title"
              class="flex-1 bg-white text-black pb-4 rounded-full transition-colors duration-300"
            />
            <UInput
              v-model="description"
              :ui="{
                placeholder: '!px-5 !py-3 font-light',
                size: { sm: 'text-3xl' },
              }"
              placeholder="Description"
              class="flex-1 bg-white text-black rounded-full transition-colors duration-300"
            />
            <div class="flex flex-row justify-end items-center gap-5 pt-5">
              <UButton
                class="text-red-600 border-2 border-red-600 bg-opacity-0 text-2xl font-semibold py-3 px-5"
                @click="validateCreation"
                >Cancel</UButton
              >
              <UButton
                class="text-black border-2 border-black bg-opacity-0 text-2xl font-semibold py-3 px-5"
                @click="onCreate"
                >Submit</UButton
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
