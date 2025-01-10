<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

onMounted(() => {
  navigateTo("/workflow");
});

</script>
