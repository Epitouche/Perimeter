<script setup lang="ts">
const props = defineProps<{
  timeout: number;
}>();

const timedOut = ref(false);

onMounted(() => {
  setTimeout(() => {
    timedOut.value = true;
    navigateTo("/myareas");
  }, props.timeout);
});
</script>

<template>
  <div v-if="!timedOut" class="loading-screen">
    <div class="loader"/>
    <p>Chargement...</p>
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
  width: 50px;
  height: 50px;
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
