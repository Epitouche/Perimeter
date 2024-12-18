<script setup lang="ts">
import servicebyid from "~/server/api/servicebyid";

definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const router = useRouter();
const serviceId = route.params.service;
const token = useCookie("token");
const isLoading = ref(true);
const actions = ref<any>(null);
const error = ref<string | null>(null);

const configIsOpen = ref<{ [key: number]: boolean }>({});
const modifiedOptions = reactive<{
  [key: number]: { [key: string]: string | number };
}>({});

const serviceInfo = ref<{ name: string } | null>(null);

const getServiceInfo = async () => {
  if (serviceId) {
    isLoading.value = true;
    try {
      error.value = null;
      serviceInfo.value = await $fetch("/api/servicebyid", {
        method: "POST",
        body: {
          token: token.value,
          serviceId: serviceId,
        },
      });
      // console.log("services", serviceInfo.value);
    } catch (err) {
      console.error("Error fetching services:", err);
    } finally {
      isLoading.value = false;
    }
  }
};

const fetchActions = async () => {
  isLoading.value = true;
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
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  getServiceInfo();
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
      actionServiceId: serviceId,
    },
  });
};
</script>

<template>
  <div>
    <div v-if="error">
      <div>Error: {{ error }}</div>
    </div>
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
    <UContainer
      v-else-if="serviceInfo"
      :ui="{ constrained: 'max-w-none' }"
      :class="[`bg-custom_color-${serviceInfo.name}`, 'py-20']"
    >
      <div class="px-20">
        <BackButton link="/workflow/actions" :is-white="true" />
      </div>
      <div class="flex flex-col justify-center items-center gap-2">
        <h1 class="text-8xl text-white font-custom_weight_title">
          Add an action
        </h1>
        <UIcon
          :name="`my-icons:white-${serviceInfo.name}`"
          class="w-[9em] h-[9em]"
        />
        <h2 class="capitalize text-white text-7xl font-bold pt-8">
          {{ serviceInfo.name }}
        </h2>
      </div>
    </UContainer>
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
    <div v-else-if="isLoading" class="text-xl font-semibold">Loading...</div>
  </div>
</template>

<style scoped></style>
