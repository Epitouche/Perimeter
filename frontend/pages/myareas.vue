<script lang="ts" setup>
import type { Workflow } from '@/interfaces/areas';
import { handleTokenStatus } from '../utils/handleErrorStatus.js';

definePageMeta({
  middleware: "auth",
});

const token = useCookie("token");
const workflows = ref<Workflow[] | null>(null);;
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
    if (typeof error === 'object' && error !== null && 'statusCode' in error) {
      const statusCode = (error as { statusCode?: number }).statusCode;
      const message = (error as { message?: string }).message || 'An error occurred';
      errorMessage.value = handleTokenStatus(statusCode, message)
    } else {
      errorMessage.value = 'An unknown error occurred';
      console.error('An unknown error occurred');
    }
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
        {{ workflow.action_id.name }}
        {{ workflow.reaction_id.name }}
      </div>
    </div>
  </div>
</template>

<style></style>
