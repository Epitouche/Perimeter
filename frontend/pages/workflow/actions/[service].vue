<script setup lang="ts">
definePageMeta({
  layout: 'nonavbar',
  middleware: 'auth'
});

const route = useRoute();
const serviceId = route.params.service;
const token = useCookie('token');

const actions = ref<any>(null);
const error = ref<string | null>(null);

const fetchActions = async () => {
  try {
    error.value = null;
    actions.value = await $fetch('/api/workflow/actions', {
      method: 'POST',
      body: {
        token: token.value,
        service: serviceId,
      },
    });
    console.log('actions', actions.value);
  } catch (error) {
    console.error('Error fetching actions:', error);
  }
};

onMounted(() => {
  fetchActions();
});

const configIsOpen = ref<{ [key: number]: boolean }>({});
const openConfig = (actionId: number) => {
  configIsOpen.value[actionId] = !configIsOpen.value[actionId] || false;
};

const parseOption = (option: string) => {
  try {
    return JSON.parse(option);
  } catch (err) {
    console.error('Invalid JSON format for option:', option, err);
    return {}; // Return an empty object to avoid breaking the UI
  }
};

</script>

<template>
  <div>
    <h1>
      Service Actions Page
    </h1>
    <div v-if="error">
      <div>Error: {{ error }}</div>
    </div>
    <div v-else-if="actions">
      <div v-for="action in actions" :key="action.id" class="mb-4">
        <button @click="openConfig(action.id)">
          {{ action.name }}
        </button>
        <div v-if="configIsOpen[action.id]" class="config-panel">
          <div>
            <h3>Options:</h3>
            <!-- <div>{{ action.option }}</div> -->
            <div v-for="(value, key) in parseOption(action.option)" :key="key">
              <strong>{{ key }}:</strong> {{ value }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
