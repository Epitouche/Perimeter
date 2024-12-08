<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const router = useRouter();
const serviceId = route.params.service;
const token = useCookie("token");

const actions = ref<any>(null);
const error = ref<string | null>(null);

const configIsOpen = ref<{ [key: number]: boolean }>({});
const modifiedOptions = reactive<{
  [key: number]: { [key: string]: string | number };
}>({});

const fetchActions = async () => {
  try {
    error.value = null;
    actions.value = await $fetch("/api/workflow/actions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });

    actions.value.forEach((action: any) => {
      const parsedOption = JSON.parse(action.option || "{}");
      for (const key in parsedOption) {
        parsedOption[key] = isNaN(Number(parsedOption[key]))
          ? parsedOption[key]
          : Number(parsedOption[key]);
      }
      modifiedOptions[action.id] = parsedOption;
    });

    console.log("actions", actions.value);
  } catch (err) {
    error.value = "Failed to load actions";
    console.error("Error fetching actions:", err);
  }
};

onMounted(() => {
  fetchActions();
});

const openConfig = (actionId: number) => {
  configIsOpen.value[actionId] = !configIsOpen.value[actionId] || false;
};

const parseOption = (option: string) => {
  try {
    return JSON.parse(option);
  } catch (err) {
    console.error("Invalid JSON format for option:", option, err);
    return {};
  }
};

const saveOptions = (actionId: number) => {
  for (const key in modifiedOptions[actionId]) {
    const value = modifiedOptions[actionId][key];
    if (!isNaN(Number(value)) && value !== "") {
      modifiedOptions[actionId][key] = Number(value);
    }
  }

  router.push({
    name: "workflow",
    query: {
      actionId: actionId,
      actionOptions: JSON.stringify(modifiedOptions[actionId]),
    },
  });
};
</script>

<template>
  <div>
    <h1>Service Actions Page</h1>
    <div v-if="error">
      <div>Error: {{ error }}</div>
    </div>
    <div v-else-if="actions">
      <div v-for="action in actions" :key="action.id">
        <button @click="openConfig(action.id)">
          {{ action.name }}
        </button>
        <div v-if="configIsOpen[action.id]">
          <div>
            <div v-for="(value, key) in parseOption(action.option)" :key="key">
              <strong>{{ key }}:</strong>
              <input
                v-model="modifiedOptions[action.id][key]"
                :type="typeof value === 'number' ? 'number' : 'text'"
              />
            </div>
            <button @click="saveOptions(action.id)">Save</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
