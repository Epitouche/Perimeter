<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const router = useRouter();

const SHOW_NAVBAR_KEY = "workflow_showNavBar";
const SHOW_CANCEL_BUTTON_KEY = "workflow_showCancelButton";
const SHOW_CREATE_BUTTON_KEY = "workflow_showCreateButton";
const REACTION_BUTTON_DISABLED_KEY = "workflow_reactionButtonDisabled";
const ACTION_SELECTED_KEY = "workflow_actionIsSelected";
const REACTION_SELECTED_KEY = "workflow_reactionIsSelected";

const ACTION_KEY = "workflow_actionId";
const ACTION_OPTIONS_KEY = "workflow_actionOptions";
const REACTION_KEY = "workflow_reactionId";
const REACTION_OPTIONS_KEY = "workflow_reactionOptions";

const isClient = typeof window !== "undefined";

const showNavBar = ref<boolean>(true);
const showCancelButton = ref<boolean>(false);
const reactionButtonisDisabled = ref<boolean>(true);
const showCreateButton = ref<boolean>(false);
const actionIsSelected = ref<boolean>(false);
const reactionIsSelected = ref<boolean>(false);

const actionId = ref<string | null>(null);
const actionOptions = ref<any>(null);
const reactionId = ref<string | null>(null);
const reactionOptions = ref<any>(null);

const error = ref<string | null>(null);

const loadWorkflowState = () => {
  if (isClient) {
    actionIsSelected.value = JSON.parse(
      localStorage.getItem(ACTION_SELECTED_KEY) || "false",
    );
    reactionIsSelected.value = JSON.parse(
      localStorage.getItem(REACTION_SELECTED_KEY) || "false",
    );
    showNavBar.value = JSON.parse(
      localStorage.getItem(SHOW_NAVBAR_KEY) || "true",
    );
    showCancelButton.value = JSON.parse(
      localStorage.getItem(SHOW_CANCEL_BUTTON_KEY) || "false",
    );
    showCreateButton.value = JSON.parse(
      localStorage.getItem(SHOW_CREATE_BUTTON_KEY) || "false",
    );
    reactionButtonisDisabled.value = JSON.parse(
      localStorage.getItem(REACTION_BUTTON_DISABLED_KEY) || "true",
    );

    const queryActionId = Array.isArray(route.query.actionId)
      ? route.query.actionId[0]
      : route.query.actionId;
    actionId.value = queryActionId || localStorage.getItem(ACTION_KEY);

    const queryActionOptions = Array.isArray(route.query.actionOptions)
      ? route.query.actionOptions[0]
      : route.query.actionOptions;

    actionOptions.value = queryActionOptions
      ? JSON.parse(queryActionOptions as string)
      : JSON.parse(localStorage.getItem(ACTION_OPTIONS_KEY) || "null");

    const queryReactionId = Array.isArray(route.query.reactionId)
      ? route.query.reactionId[0]
      : route.query.reactionId;
    reactionId.value = queryReactionId || localStorage.getItem(REACTION_KEY);

    const queryReactionOptions = Array.isArray(route.query.reactionOptions)
      ? route.query.reactionOptions[0]
      : route.query.reactionOptions;

    reactionOptions.value = queryReactionOptions
      ? JSON.parse(queryReactionOptions as string)
      : JSON.parse(localStorage.getItem(REACTION_OPTIONS_KEY) || "null");
  }
};

const saveWorkflowState = () => {
  if (isClient) {
    localStorage.setItem(
      ACTION_SELECTED_KEY,
      JSON.stringify(actionIsSelected.value),
    );
    localStorage.setItem(
      REACTION_SELECTED_KEY,
      JSON.stringify(reactionIsSelected.value),
    );
    localStorage.setItem(SHOW_NAVBAR_KEY, JSON.stringify(showNavBar.value));
    localStorage.setItem(
      SHOW_CANCEL_BUTTON_KEY,
      JSON.stringify(showCancelButton.value),
    );
    localStorage.setItem(
      SHOW_CREATE_BUTTON_KEY,
      JSON.stringify(showCreateButton.value),
    );
    localStorage.setItem(
      REACTION_BUTTON_DISABLED_KEY,
      JSON.stringify(reactionButtonisDisabled.value),
    );

    if (actionId.value) {
      localStorage.setItem(ACTION_KEY, actionId.value);
    }
    if (actionOptions.value) {
      localStorage.setItem(
        ACTION_OPTIONS_KEY,
        JSON.stringify(actionOptions.value),
      );
    }
    if (reactionId.value) {
      localStorage.setItem(REACTION_KEY, reactionId.value);
    }
    if (reactionOptions.value) {
      localStorage.setItem(
        REACTION_OPTIONS_KEY,
        JSON.stringify(reactionOptions.value),
      );
    }
  }
};

const clearWorkflowState = () => {
  if (isClient) {
    localStorage.removeItem(ACTION_SELECTED_KEY);
    localStorage.removeItem(REACTION_SELECTED_KEY);
    localStorage.removeItem(SHOW_NAVBAR_KEY);
    localStorage.removeItem(SHOW_CANCEL_BUTTON_KEY);
    localStorage.removeItem(SHOW_CREATE_BUTTON_KEY);
    localStorage.removeItem(REACTION_BUTTON_DISABLED_KEY);

    localStorage.removeItem(ACTION_KEY);
    localStorage.removeItem(ACTION_OPTIONS_KEY);
    localStorage.removeItem(REACTION_KEY);
    localStorage.removeItem(REACTION_OPTIONS_KEY);
  }
};

const onActionSelected = () => {
  showNavBar.value = false;
  showCancelButton.value = true;
  reactionButtonisDisabled.value = false;
  actionIsSelected.value = true;
  saveWorkflowState();
};

const onReactionSelected = () => {
  showCreateButton.value = true;
  reactionIsSelected.value = true;
  saveWorkflowState();
};

const setWorkflowPageDefault = () => {
  actionId.value = null;
  actionOptions.value = null;
  reactionId.value = null;
  reactionOptions.value = null;
  showNavBar.value = true;
  showCancelButton.value = false;
  showCreateButton.value = false;
  reactionButtonisDisabled.value = true;
  actionIsSelected.value = false;
  reactionIsSelected.value = false;
  clearWorkflowState();
  router.push("/workflow");
};

const onCreate = async () => {
  try {
    error.value = null;
    await $fetch("/api/workflow/create", {
      method: "POST",
      body: {
        actionOptions: actionOptions.value,
        actionId: actionId.value,
        reactionOptions: reactionOptions.value,
        reactionId: reactionId.value,
      },
    });
    console.log("Workflow created");
    setWorkflowPageDefault();
  } catch (err) {
    console.error("Error creating workflow:", err);
    console.log("Error creating workflow:", err);
  }
};

onMounted(() => {
  loadWorkflowState();
  if (actionIsSelected.value === false && actionId.value) {
    onActionSelected();
  }
  if (reactionId.value) {
    onReactionSelected();
  }
});
</script>

<template>
  <div>
    <div v-if="showNavBar" class="pb-10">
      <NavigationBar />
    </div>
    <div v-if="showCancelButton" class="pt-24 pl-28">
      <UButton
        class="bg-white text-custom_color-text text-4xl font-bold px-7 py-3 !border-custom_border_width border-custom_color-border"
        @click="setWorkflowPageDefault"
        >Cancel</UButton
      >
    </div>

    <div class="flex flex-col justify-center items-center gap-10">
      <h1 class="text-custom_size_title font-custom_weight_title pb-5">
        Workflow
      </h1>
      <div class="flex flex-col justify-center items-center w-[28%]">
        <ReActionButton
          title="Action"
          link="/workflow/actions"
          :is-disabled="false"
          :is-selected="actionIsSelected"
        />
        <div
          :class="[
            'bg-black min-w-4 min-h-28',
            reactionButtonisDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
          ]"
        />
        <ReActionButton
          title="Reaction"
          link="/workflow/reactions"
          :is-disabled="reactionButtonisDisabled"
          :is-selected="reactionIsSelected"
        />
      </div>
      <div v-if="showCreateButton" class="pt-10">
        <UButton class="text-5xl font-bold px-8 py-4" @click="onCreate"
          >Create</UButton
        >
      </div>
    </div>
  </div>
</template>

<style scoped></style>
