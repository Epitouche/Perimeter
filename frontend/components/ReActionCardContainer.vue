<!-- <script setup lang="ts">
interface Type {
  id: number;
  name: string;
  description: string;
  option?: string;
}

interface ServiceInfo {
  id: number;
  name: string;
}

const props = defineProps<{
  typeName: string;
  types: Type[];
  serviceInfo: ServiceInfo | null;
}>();

const router = useRouter();

const configIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.types.map((type) => [type.id, false])),
);

const state = reactive<{ [key: number]: Record<string, unknown> }>(
  Object.fromEntries(
    props.types.map((type) => [type.id, JSON.parse(type.option || "{}")]),
  ),
);

const initialState = reactive<{ [key: number]: Record<string, unknown> }>(
  JSON.parse(JSON.stringify(state)),
);

const openConfig = (typeId: number) => {
  if (configIsOpen[typeId]) {
    state[typeId] = JSON.parse(JSON.stringify(initialState[typeId]));
  }

  configIsOpen[typeId] = !configIsOpen[typeId];
};

const onSubmit = (typeId: number) => {
  const modifiedOptions = state[typeId];

  router.push({
    name: "workflow",
    query: {
      [`${props.typeName}Id`]: typeId,
      [`${props.typeName}Options`]: JSON.stringify(modifiedOptions),
      [`${props.typeName}ServiceId`]: props.serviceInfo.id,
    },
  });
};

function formatString(str: string): string {
  return str.replace(/([a-z])([A-Z])/g, "$1 $2");
}

onMounted(() => {
  console.log("types", props.types);
});
</script>

<template>
  <UContainer
    :ui="{ padding: 'px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <div v-for="type in types" :key="type.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        :class="[
          `bg-custom_color-${serviceInfo.name}`,
          'basis-1/4 flex flex-col justify-evenly items-center gap-4 text-white font-extrabold text-6xl p-8 rounded-custom_border_radius w-[5em] h-[4.5em]',
        ]"
        @click="openConfig(type.id)"
      >
        <h2
          class="clamp-2-lines capitalize text-4xl text-center break-words w-full"
        >
          {{ formatString(type.name) }}
        </h2>
        <p class="text-2xl">{{ type.description }}</p>
      </UContainer>

      <UModal
        v-model="configIsOpen[type.id]"
        :ui="{
          base: `relative text-left rtl:text-right flex flex-col p-10 border-custom_border_width border-custom_color-${serviceInfo.name}`,
        }"
      >
        <template #default>
          <UForm
            :state="state[type.id]"
            class="flex flex-col gap-12 p-5 bg-custom_color-bg_section"
            @submit.prevent="onSubmit(type.id)"
          >
            <h2 class="text-center text-6xl font-semibold pb-2">
              {{ formatString(type.name) }}
            </h2>
            <UFormGroup
              v-for="(value, key) in state[type.id]"
              :key="key"
              :label="key"
              :name="key"
              :ui="{ label: { base: 'capitalize text-2xl' } }"
            >
              <UInput
                v-model="state[type.id][key] as string | number | undefined"
                :ui="{
                  placeholder: '!px-5 !py-3 font-light capitalize',
                  size: { sm: 'text-3xl' },
                }"
                :placeholder="key + '...'"
              />
            </UFormGroup>

            <div class="flex flex-rox justify-evenly gap-4 pt-4">
              <UButton
                class="text-3xl font-semibold px-5 py-3 text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border"
                @click="openConfig(type.id)"
              >
                Cancel
              </UButton>
              <UButton type="submit" class="text-3xl font-semibold px-5 py-3">
                Submit
              </UButton>
            </div>
          </UForm>
        </template>
      </UModal>
    </div>
  </UContainer>
</template>

<style scoped></style> -->

<script setup lang="ts">
interface Type {
  id: number;
  name: string;
  description: string;
  option?: string; // JSON string with configuration options
}

interface ServiceInfo {
  id: number;
  name: string;
}

// Props definition with types
const props = defineProps<{
  typeName: string;
  types: Type[];
  serviceInfo: ServiceInfo | null;
}>();

const router = useRouter();

// Tracks whether a configuration modal is open for each type
const configIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.types.map((type) => [type.id, false])),
);

// State object to store current form values for each type
const state = reactive<{ [key: number]: Record<string, string | number> }>(
  Object.fromEntries(
    props.types.map((type) => [type.id, JSON.parse(type.option || "{}")]),
  ),
);

// Stores initial form values to reset on cancel
const initialState = reactive<{ [key: number]: Record<string, string | number> }>(
  JSON.parse(JSON.stringify(state)),
);

// Handles opening/closing the configuration modal
const openConfig = (typeId: number) => {
  if (configIsOpen[typeId]) {
    // Reset the state if the modal is closed
    state[typeId] = JSON.parse(JSON.stringify(initialState[typeId]));
  }

  configIsOpen[typeId] = !configIsOpen[typeId];
};

// Handles form submission
const onSubmit = (typeId: number) => {
  const modifiedOptions = state[typeId];

  // Navigate to the "workflow" route with the updated parameters
  router.push({
    name: "workflow",
    query: {
      [`${props.typeName}Id`]: typeId,
      [`${props.typeName}Options`]: JSON.stringify(modifiedOptions),
      [`${props.typeName}ServiceId`]: props.serviceInfo?.id,
    },
  });
};

// Utility function to format camelCase strings into readable text
function formatString(str: string): string {
  return str.replace(/([a-z])([A-Z])/g, "$1 $2");
}

onMounted(() => {
  console.log("types", props.types);
});
</script>

<template>
  <UContainer
    :ui="{ padding: 'px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <!-- Loop through each type to display cards -->
    <div v-for="type in props.types" :key="type.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        :class="[
          `bg-custom_color-${props.serviceInfo?.name}`,
          'basis-1/4 flex flex-col justify-evenly items-center gap-4 text-white font-extrabold text-6xl p-8 rounded-custom_border_radius w-[5em] h-[4.5em]',
        ]"
        @click="openConfig(type.id)"
      >
        <h2
          class="clamp-2-lines capitalize text-4xl text-center break-words w-full"
        >
          {{ formatString(type.name) }}
        </h2>
        <p class="text-2xl">{{ type.description }}</p>
      </UContainer>

      <!-- Modal for configuration -->
      <UModal
        v-model="configIsOpen[type.id]"
        :ui="{
          base: `relative text-left rtl:text-right flex flex-col p-10 border-custom_border_width border-custom_color-${props.serviceInfo?.name}`,
        }"
      >
        <template #default>
          <UForm
            :state="state[type.id]"
            class="flex flex-col gap-12 p-5 bg-custom_color-bg_section"
            @submit.prevent="onSubmit(type.id)"
          >
            <h2 class="text-center text-6xl font-semibold pb-2">
              {{ formatString(type.name) }}
            </h2>

            <!-- Dynamically generate form inputs based on state keys -->
            <UFormGroup
              v-for="(value, key) in state[type.id]"
              :key="key"
              :label="key"
              :name="key"
              :ui="{ label: { base: 'capitalize text-2xl' } }"
            >
              <UInput
                v-model="state[type.id][key] as string | number | undefined"
                :ui="{
                  placeholder: '!px-5 !py-3 font-light capitalize',
                  size: { sm: 'text-3xl' },
                }"
                :placeholder="key + '...'"
              />
            </UFormGroup>

            <!-- Buttons for cancel and submit -->
            <div class="flex flex-row justify-evenly gap-4 pt-4">
              <UButton
                class="text-3xl font-semibold px-5 py-3 text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border"
                @click="openConfig(type.id)"
              >
                Cancel
              </UButton>
              <UButton type="submit" class="text-3xl font-semibold px-5 py-3">
                Submit
              </UButton>
            </div>
          </UForm>
        </template>
      </UModal>
    </div>
  </UContainer>
</template>

