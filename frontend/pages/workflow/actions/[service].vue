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
    actions.value.forEach((action: any) => {
      modifiedOptions[action.id] = JSON.parse(action.option || '{}');
    });
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
  console.log('option', option);
  try {
    return JSON.parse(option);
  } catch (err) {
    console.error('Invalid JSON format for option:', option, err);
    return {};
  }
};

const modifiedOptions = reactive<{ [key: number]: { [key: string]: string } }>({});
const saveOptions = async (actionId: number) => {
  try {
    const payload = {
      token: token.value,
      service: serviceId,
      actionId: actionId,
      options: modifiedOptions[actionId],
    };

    const response = await $fetch('/api/workflow/actions/save', {
      method: 'POST',
      body: payload,
    });

    console.log('Options saved successfully:', response);
    alert('Options saved successfully!');
  } catch (err) {
    console.error('Error saving options:', err);
    alert('Failed to save options. Please try again.');
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
            <div v-for="(value, key) in parseOption(action.option)" :key="key">
              <strong>{{ key }}:</strong>
              <input v-model="modifiedOptions[action.id][key]" type="text" />
            </div>
            <button @click="saveOptions(action.id)" class="save-button">Save</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
