<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";
import type { Type } from "@/interfaces/type";

const props = defineProps<{
  typeName: string;
  types: Type[];
  serviceInfo: ServiceInfo | null;
}>();

const router = useRouter();

const configIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.types.map((type) => [type.id, false])),
);

const state = reactive<{ [key: number]: Record<string, string | number> }>(
  Object.fromEntries(
    props.types.map((type) => [
      type.id,
      typeof type.option === "string"
        ? JSON.parse(type.option)
        : type.option || {},
    ]),
  ),
);

const initialState = reactive<{
  [key: number]: Record<string, string | number>;
}>(JSON.parse(JSON.stringify(state)));

const openConfig = (typeId: number) => {
  if (configIsOpen[typeId]) {
    state[typeId] = JSON.parse(JSON.stringify(initialState[typeId]));
  }

  configIsOpen[typeId] = !configIsOpen[typeId];
};

const onSubmit = (typeId: number, typeTitle: string) => {
  const modifiedOptions = { ...state[typeId] };

  const hasInvalidTypes = Object.entries(modifiedOptions).some(
    ([key, value]) => {
      const expectedType = fieldTypes[typeId][key];
      return typeof value !== expectedType;
    },
  );

  if (hasInvalidTypes) {
    alert("Some fields have invalid types! Please correct them.");
    return;
  }

  router.push({
    name: "workflow",
    query: {
      [`${props.typeName}Id`]: typeId,
      [`${props.typeName}Options`]: JSON.stringify(modifiedOptions),
      [`${props.typeName}ServiceId`]: props.serviceInfo?.id,
      [`${props.typeName}Name`]: typeTitle,
    },
  });
};

const fieldTypes = reactive<{ [key: number]: Record<string, string> }>(
  Object.fromEntries(
    props.types.map((type) => [
      type.id,
      Object.fromEntries(
        Object.keys(state[type.id]).map((key) => {
          const value = state[type.id][key];
          return [key, typeof value];
        }),
      ),
    ]),
  ),
);

function formatString(str: string): string {
  return str.replace(/([a-z])([A-Z])/g, "$1 $2");
}

onMounted(() => {
  console.log("types", props.types);
});
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <div v-for="type in props.types" :key="type.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        class="custom_card flex flex-col justify-evenly items-center gap-4 text-white"
        :style="{ backgroundColor: props.serviceInfo?.color || 'black' }"
        tabindex="0"
        @click="openConfig(type.id)"
        @keydown.space="openConfig(type.id)"
      >
        <h4 class="clamp-2-lines capitalize text-center break-words w-full">
          {{ formatString(type.name) }}
        </h4>
      </UContainer>

      <UModal
        v-model="configIsOpen[type.id]"
        :ui="{
          base: 'relative text-left rtl:text-right flex flex-col p-10 border-custom_border_width', width: 'w-fit'
        }"
        :style="{ borderColor: props.serviceInfo?.color || 'black' }"
      >
        <template #default>
          <UForm
            :state="state[type.id]"
            class="flex flex-col gap-12 p-5 bg-custom_color-bg_section"
            @submit.prevent="onSubmit(type.id, type.name)"
          >
            <h2 class="text-center pb-2">
              {{ formatString(type.name) }}
            </h2>
            <h5 class="text-center -mt-6">
              {{ type.description }}
            </h5>

            <UFormGroup
              v-for="(value, key) in state[type.id]"
              :key="key"
              :label="key"
              :name="key"
              :ui="{ label: { base: 'capitalize text-2xl pl-5 font-semibold' } }"
            >
              <UInput
                v-model="state[type.id][key] as string | number | undefined"
                :type="
                  fieldTypes[type.id][key] === 'number' ? 'number' : 'text'
                "
                :ui="{
                  placeholder: '!px-5 !py-3 font-light',
                  size: { sm: 'text-3xl' },
                }"
                :placeholder="key + '...'"
              />
            </UFormGroup>

            <div class="flex flex-row justify-evenly gap-4 pt-4">
              <UButton
                class="text-3xl font-semibold px-5 py-3 text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border"
                tabindex="0"
                @click="openConfig(type.id)"
              >
                Cancel
              </UButton>
              <UButton
                type="submit"
                class="text-3xl font-semibold px-5 py-3"
                tabindex="0"
              >
                Submit
              </UButton>
            </div>
          </UForm>
        </template>
      </UModal>
    </div>
  </UContainer>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
