<script lang="ts" setup>
definePageMeta({
  middleware: "auth",
});

// const token = useCookie("token");

// const workflows = ref<any>(null);
// const error = ref<string | null>(null);

// const fetchWorkflows = async () => {
//   try {
//     error.value = null;
//     workflows.value = await $fetch("/api/workflow/workflows", {
//       method: "POST",
//       body: {
//         token: token.value,
//       },
//     });

//     console.log("workflows", workflows);
//   } catch (err) {
//     console.error("Error fetching workflows:", err);
//   }
// };

</script>

<template>
  <div>My Areas</div>
</template>

<style></style>
