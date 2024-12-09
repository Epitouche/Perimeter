<script setup lang="ts">
definePageMeta({
  layout: "nonavbar",
  middleware: "auth",
});

const route = useRoute();
const router = useRouter();
const serviceId = route.params.service;
const token = useCookie("token");

const reactions = ref<any>(null);
const error = ref<string | null>(null);

const configIsOpen = ref<{ [key: number]: boolean }>({});
const modifiedOptions = reactive<{
  [key: number]: { [key: string]: string | number };
}>({});

const fetchReactions = async () => {
  try {
    error.value = null;
    reactions.value = await $fetch("/api/workflow/reactions", {
      method: "POST",
      body: {
        token: token.value,
        service: serviceId,
      },
    });

    reactions.value.forEach((reaction: any) => {
      const parsedOption = JSON.parse(reaction.option || "{}");
      for (const key in parsedOption) {
        parsedOption[key] = isNaN(Number(parsedOption[key]))
          ? parsedOption[key]
          : Number(parsedOption[key]);
      }
      modifiedOptions[reaction.id] = parsedOption;
    });

    console.log("reactions", reactions.value);
  } catch (err) {
    error.value = "Failed to load reactions";
    console.error("Error fetching reactions:", err);
  }
};

onMounted(() => {
  fetchReactions();
});

const openConfig = (reactionId: number) => {
  configIsOpen.value[reactionId] = !configIsOpen.value[reactionId] || false;
};

const parseOption = (option: string) => {
  try {
    return JSON.parse(option);
  } catch (err) {
    console.error("Invalid JSON format for option:", option, err);
    return {};
  }
};

const saveOptions = (reactionId: number) => {
  for (const key in modifiedOptions[reactionId]) {
    const value = modifiedOptions[reactionId][key];
    if (!isNaN(Number(value)) && value !== "") {
      modifiedOptions[reactionId][key] = Number(value);
    }
  }

  router.push({
    name: "workflow",
    query: {
      reactionId: reactionId,
      reactionOptions: JSON.stringify(modifiedOptions[reactionId]),
    },
  });
};
</script>

<template>
  <div>
    <UButton to="/workflow/reactions">Back</UButton>
    <h1>Service Reactions Page</h1>
    <div v-if="error">
      <div>Error: {{ error }}</div>
    </div>
    <div v-else-if="reactions">
      <div v-for="reaction in reactions" :key="reaction.id">
        <button @click="openConfig(reaction.id)">
          {{ reaction.name }}
        </button>
        <div v-if="configIsOpen[reaction.id]">
          <div>
            <div
              v-for="(value, key) in parseOption(reaction.option)"
              :key="key"
            >
              <strong>{{ key }}:</strong>
              <input
                v-model="modifiedOptions[reaction.id][key]"
                :type="typeof value === 'number' ? 'number' : 'text'"
              />
            </div>
            <button @click="saveOptions(reaction.id)">Save</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>