<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const token = useCookie("token");

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

const ACTION_SERVICE_KEY = "workflow_actionServiceId";
const REACTION_SERVICE_KEY = "workflow_reactionServiceId";

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

const actionServiceId = ref<string | null>(null);
const reactionServiceId = ref<string | null>(null);

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

    const queryActionId = Array.isArray(route.query.actionId)
      ? route.query.actionId[0]
      : route.query.actionId;
    actionId.value = queryActionId || localStorage.getItem(ACTION_KEY);

    const queryActionOptions = Array.isArray(route.query.actionOptions)
      ? route.query.actionOptions[0]
      : route.query.actionOptions;

    actionOptions.value = queryActionOptions
      ? JSON.parse(queryActionOptions as string, (key, value) =>
          typeof value === "string" && !isNaN(Number(value))
            ? Number(value)
            : value,
        )
      : JSON.parse(
          localStorage.getItem(ACTION_OPTIONS_KEY) || "null",
          (key, value) =>
            typeof value === "string" && !isNaN(Number(value))
              ? Number(value)
              : value,
        );

    const queryReactionId = Array.isArray(route.query.reactionId)
      ? route.query.reactionId[0]
      : route.query.reactionId;
    reactionId.value = queryReactionId || localStorage.getItem(REACTION_KEY);

    const queryReactionOptions = Array.isArray(route.query.reactionOptions)
      ? route.query.reactionOptions[0]
      : route.query.reactionOptions;

    reactionOptions.value = queryReactionOptions
      ? JSON.parse(queryReactionOptions as string, (key, value) =>
          typeof value === "string" && !isNaN(Number(value))
            ? Number(value)
            : value,
        )
      : JSON.parse(
          localStorage.getItem(REACTION_OPTIONS_KEY) || "null",
          (key, value) =>
            typeof value === "string" && !isNaN(Number(value))
              ? Number(value)
              : value,
        );

    const queryActionServiceId = Array.isArray(route.query.actionServiceId)
      ? route.query.actionServiceId[0]
      : route.query.actionServiceId;
    actionServiceId.value =
      queryActionServiceId || localStorage.getItem(ACTION_SERVICE_KEY);

    const queryReactionServiceId = Array.isArray(route.query.reactionServiceId)
      ? route.query.reactionServiceId[0]
      : route.query.reactionServiceId;
    reactionServiceId.value =
      queryReactionServiceId || localStorage.getItem(REACTION_SERVICE_KEY);
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
    if (actionServiceId.value) {
      localStorage.setItem(ACTION_SERVICE_KEY, actionServiceId.value);
    }
    if (reactionServiceId.value) {
      localStorage.setItem(REACTION_SERVICE_KEY, reactionServiceId.value);
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

    localStorage.removeItem(ACTION_SERVICE_KEY);
    localStorage.removeItem(REACTION_SERVICE_KEY);
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
  reactionButtonisDisabled.value = false;
  saveWorkflowState();
};

const setWorkflowPageDefault = () => {
  actionId.value = null;
  actionOptions.value = null;
  reactionId.value = null;
  reactionOptions.value = null;
  actionServiceId.value = null;
  reactionServiceId.value = null;
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
  console.log("actionId:", actionId.value);
  console.log("actionOptions:", actionOptions.value);
  console.log("reactionId:", reactionId.value);
  console.log("reactionOptions:", reactionOptions.value);
  try {
    error.value = null;
    const response = await $fetch("/api/workflow/create", {
      method: "POST",
      body: {
        token: token.value,
        actionOptions: JSON.stringify(actionOptions.value),
        actionId: actionId.value,
        reactionOptions: JSON.stringify(reactionOptions.value),
        reactionId: reactionId.value,
      },
    });
    console.log("response:", response);
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
      <div class="flex flex-col justify-center items-center">
        <ReActionButton
          title="Action"
          link="/workflow/actions"
          :is-disabled="false"
          :is-selected="actionIsSelected"
          :service-id="Number(actionServiceId)"
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
          :service-id="Number(reactionServiceId)"
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
