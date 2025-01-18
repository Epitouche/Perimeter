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

// const state = reactive<{ [key: number]: Record<string, string | number> }>(
//   Object.fromEntries(
//     props.types.map((type) => [
//       type.id,
//       typeof type.option === "string"
//         ? JSON.parse(type.option)
//         : type.option || {},
//     ]),
//   ),
// );

const state = reactive<{
  [key: number]: Record<string, string | number | undefined>;
}>(
  Object.fromEntries(
    props.types.map((type) => [
      type.id,
      typeof type.option === "string"
        ? JSON.parse(type.option)
        : Object.keys(type.option || {}).reduce(
            (acc, key) => {
              acc[key] = undefined;
              return acc;
            },
            {} as Record<string, undefined>,
          ),
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
          return [
            key,
            value === undefined || value === null ? "string" : typeof value,
          ];
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
        :ui="{ padding: '!px-0 !py-6', constrained: 'max-w-none' }"
        class="custom_card flex flex-col justify-evenly items-center gap-4 text-white"
        :style="{ backgroundColor: props.serviceInfo?.color || 'black' }"
        tabindex="0"
        @click="openConfig(type.id)"
        @keydown.space="openConfig(type.id)"
      >
        <h4 class="clamp-1-line p-2 capitalize text-center break-words w-full">
          {{ formatString(type.name) }}
        </h4>
      </UContainer>

      <UModal
        v-model="configIsOpen[type.id]"
        :ui="{
          base: 'relative text-left rtl:text-right flex flex-col justify-center items-center p-10 max-lg:p-9 max-md:p-8 max-sm:p-6 border-custom_border_width',
          width: 'w-fit max-w-[95%]',
        }"
        :style="{ borderColor: props.serviceInfo?.color || 'black' }"
      >
        <template #default>
          <UForm
            :state="state[type.id]"
            class="flex flex-col gap-12 p-5 max-lg:p-4 max-md:p-3 max-sm:p-2 w-full bg-custom_color-bg_section"
            @submit.prevent="onSubmit(type.id, type.name)"
          >
            <h2 class="text-center">
              {{ formatString(type.name) }}
            </h2>
            <h6 class="text-center -mt-6 flex-wrap">
              {{ type.description }}
            </h6>

            <UFormGroup
              v-for="(value, key) in state[type.id]"
              :key="key"
              :label="key"
              :name="key"
              :ui="{
                label: {
                  base: 'capitalize text-3xl max-lg:text-2xl max-md:text-xl max-sm:text-lg px-5 font-semibold',
                },
              }"
              class="self-center min-w-[85%] max-lg:min-w-[90%] max-md:min-w-[95%] max-sm:min-w-full"
            >
              <UInput
                v-model="state[type.id][key]"
                :type="
                  fieldTypes[type.id][key] === 'number' ? 'number' : 'text'
                "
                :ui="{
                  placeholder: '!px-5 !py-3 max-lg:!py-2 font-light',
                  size: {
                    sm: 'text-3xl max-lg:text-2xl max-md:text-xl max-sm:text-lg',
                  },
                }"
                :placeholder="
                  typeof type.option === 'object' &&
                  type.option !== null &&
                  (type.option as Record<string, any>)[key]
                    ? 'Ex: ' + String((type.option as Record<string, any>)[key])
                    : 'Ex: ' + key
                "
                :value="
                  state[type.id][key] === undefined
                    ? ''
                    : 'Ex: ' + state[type.id][key]
                "
              />
            </UFormGroup>

            <div class="flex flex-row justify-evenly gap-4 pt-4">
              <UButton
                class="font-semibold px-6 py-5 max-lg:py-3 max-md:py-2 max-sm:py-1 text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border"
                tabindex="0"
                @click="openConfig(type.id)"
              >
                <h6>Cancel</h6>
              </UButton>
              <UButton
                type="submit"
                class="font-semibold px-6 py-5 max-lg:py-3 max-md:py-2 max-sm:py-1"
                tabindex="0"
              >
                <h6>Submit</h6>
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
