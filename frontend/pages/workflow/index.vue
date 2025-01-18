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
const title = ref<string>();
const description = ref<string>();
const refreshRate = ref<number>();

const validateCreation = () => {
  creationPopup.value = !creationPopup.value;
};

const onCreate = async () => {
  if (!title.value || !description.value || !refreshRate.value) {
    alert("Please fill out all fields");
    return;
  }

  if (isNaN(refreshRate.value)) {
    alert("Refresh rate must be a number");
    return;
  }

  creationPopup.value = false;
  error.value = null;

  try {
    await $fetch("/api/workflow/create", {
      method: "POST",
      body: {
        token: token.value,
        actionId: websiteStore.actionId,
        actionOptions: websiteStore.actionOptions,
        refreshRate: refreshRate.value,
        description: description.value,
        reactionId: websiteStore.reactionId,
        reactionOptions: websiteStore.reactionOptions,
        title: title.value,
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
      class="flex justify-center items-center h-screen w-screen"
    >
      <h1 class="text-center">{{ createdMessage }}</h1>
    </div>
    <div v-if="showPageContent">
      <div v-if="websiteStore.showNavBar" class="pb-10">
        <NavigationBar />
      </div>
      <div v-if="websiteStore.showCancelButton" class="pt-10 pl-10 max-sm:pt-3 max-sm:pl-2">
        <UButton
          class="bg-white text-custom_color-text p-4 max-lg:py-2 !border-custom_border_width border-custom_color-border"
          tabindex="0"
          @click="onCancel()"
          ><h6>Cancel</h6></UButton
        >
      </div>

      <div class="flex flex-col justify-center items-center gap-10 max-sm:pt-5">
        <h1 class="pb-5">Workflow</h1>
        <div v-if="isLoading"><h4>Loading...</h4></div>
        <div class="flex flex-col justify-center items-center w-full">
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
            class="px-8 py-7 max-lg:py-4"
            tabindex="0"
            @click="validateCreation"
            ><h5>Create</h5></UButton
          >
        </div>
        <div
          v-if="creationPopup"
          class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50"
        >
          <div
            class="flex flex-col justify-center items-center gap-10 bg-white px-14 py-10 max-lg:px-12 max-md:px-8 max-sm:px-2 border-custom_border_width rounded-custom_border_radius shadow-lg w-fit max-w-[90%]"
          >
            <h3 class="text-center">You're about to<br >create a new area!</h3>
            <div class="flex flex-col gap-1 min-w-[90%] max-w-[95%]">
              <h6 class="px-5">Title</h6>
              <UInput
                v-model="title"
                :ui="{
                  placeholder: '!px-5 !py-3 font-light',
                  size: {
                    sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl',
                  },
                }"
                placeholder="Title"
                class="flex-1 bg-white text-black rounded-full transition-colors duration-300"
              />
            </div>
            <div class="flex flex-col gap-1 min-w-[90%] max-w-[95%]">
              <h6 class="px-5">Description</h6>
              <UInput
                v-model="description"
                :ui="{
                  placeholder: '!px-5 !py-3 font-light',
                  size: {
                    sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl',
                  },
                }"
                placeholder="Description"
                class="flex-1 bg-white text-black rounded-full transition-colors duration-300"
              />
            </div>
            <div class="flex flex-col gap-1 min-w-[90%] max-w-[95%]">
              <h6 class="px-5">Refresh Rate</h6>
              <UInput
                v-model="refreshRate"
                :ui="{
                  placeholder: '!px-5 !py-3 font-light',
                  size: {
                    sm: 'text-4xl max-lg:text-3xl max-md:text-2xl max-sm:text-xl',
                  },
                }"
                placeholder="Ex: 0"
                class="flex-1 bg-white text-black rounded-full transition-colors duration-300"
              />
            </div>
            <div class="flex flex-row justify-end items-center gap-5 pt-5">
              <UButton
                class="text-red-600 border-custom_border_width !border-red-600 bg-opacity-0 py-5 px-6 max-lg:py-4 max-md:py-3 max-sm:py-2"
                tabindex="0"
                @click="validateCreation"
                ><h6>Cancel</h6></UButton
              >
              <UButton
                class="text-black border-custom_border_width !border-black bg-opacity-0 py-5 px-6 max-lg:py-4 max-md:py-3 max-sm:py-2"
                tabindex="0"
                @click="onCreate"
                ><h6>Submit</h6></UButton
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
