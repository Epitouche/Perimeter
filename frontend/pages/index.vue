<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

/**
 * @description This page is used to download the client APK file.
 */
onMounted(() => {
  navigateTo("/workflow");
});
</script>
