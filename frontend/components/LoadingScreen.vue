<script setup lang="ts">
/**
 * @description Loading screen component that shows a loading spinner for a given amount of time
 */

/**
 * @description If the loading screen should be shown, show it for the given amount of time
 */
const props = defineProps<{
  timeout: number; // The amount of time to show the loading screen for
  isLoading: boolean; // Whether the loading screen should be shown
}>();

/**
 * @emit Emits an event to update the isLoading prop
 */
const emit = defineEmits(["update:isLoading"]);

const timedOut = ref(false);

/**
 * @description After the timeout has passed, set timedOut to true and emit an event to hide the loading screen
 */
onMounted(() => {
  setTimeout(() => {
    timedOut.value = true;
    emit("update:isLoading", false);
  }, props.timeout);
});
</script>

<template>
  <div v-if="!timedOut && props.isLoading" class="loading-screen">
    <div class="loader" />
    <p class="text6x1 font-extrabold">Loading...</p>
  </div>
</template>

<style scoped>
.loading-screen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.9);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.loader {
  border: 4px solid rgba(0, 0, 0, 0.1);
  border-top: 4px solid #3498db;
  border-radius: 50%;
  width: 100px;
  height: 100px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
