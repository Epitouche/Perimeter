<script lang="ts" setup>
definePageMeta({
  middleware: "auth",
});

const token = useCookie("token");

const workflows = ref<any>(null);
const error = ref<string | null>(null);

const fetchWorkflows = async () => {
  try {
    error.value = null;
    workflows.value = await $fetch("/api/myareas", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    console.log("workflows: ", workflows.value);
  } catch (err: any) {
    console.error("Error fetching workflows:", err);
    console.log("Error fetching workflows:", err);
    if (err.data) {
      error.value = err.data.message || "Failed to fetch workflows";
    } else {
      error.value = "Unexpected error occurred";
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
    <div v-if="error" class="alert alert-danger">{{ error }}</div>
    <div v-else-if="workflows">
      <div v-for="workflow in workflows" :key="workflow.id">
        {{ workflow.action_id.name }}
        {{ workflow.reaction_id.name }}
      </div>
    </div>
  </div>
</template>

<style></style>
