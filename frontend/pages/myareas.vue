<script lang="ts" setup>
import type { Workflow } from "@/interfaces/areas";
import { handleErrorStatus } from "../utils/handleErrorStatus.js";

definePageMeta({
  middleware: "auth",
});

const token = useCookie("token");
const workflows = ref<Workflow[] | null>(null);
const errorMessage = ref<string | null>(null);

const fetchWorkflows = async () => {
  errorMessage.value = null;
  try {
    workflows.value = await $fetch("/api/myareas", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    console.log("workflows: ", workflows.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    console.log("error: ", error);
  }
};

onMounted(() => {
  fetchWorkflows();
});
</script>

<template>
  <div>
    <div>My Areas</div>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div v-else-if="workflows">
      <div v-for="workflow in workflows" :key="workflow.id">
        {{ workflow.action.name }}
        {{ workflow.reaction.name }}
      </div>
    </div>
  </div>
</template>

<style></style>
