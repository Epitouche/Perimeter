<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";
import type { Type } from "@/interfaces/type";

/**
 * @description This component is used to display the different types of actions and reactions that can be performed by the user.
 */

/**
 * Which type of action/reaction is being displayed, and its corresponding information and service information.
 */
const props = defineProps<{
  typeName: string; // "action" or "reaction"
  types: Type[]; // List of actions/reactions
  serviceInfo: ServiceInfo | null; // Information about the service
}>();

const router = useRouter();

/**
 * State for the options configuration modal of each action/reaction.
 */
const configIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.types.map((type) => [type.id, false]))
);

/**
 * State for the options configuration of each action/reaction.
 */
const state = reactive<{
  [key: number]: Record<string, string | number | undefined>;
}>(
  Object.fromEntries(
    props.types.map((type) => [
      type.id,
      typeof type.option === "string"
        ? JSON.parse(type.option)
        : Object.keys(type.option || {}).reduce((acc, key) => {
            acc[key] = undefined;
            return acc;
          }, {} as Record<string, undefined>),
    ])
  )
);

/**
 * Toggles the options configuration modal of an action/reaction.
 * @param typeId - The ID of the action/reaction being toggled.
 */
const toggleConfig = (typeId: number) => {
  if (configIsOpen[typeId]) {
    Object.keys(state[typeId]).forEach((key) => {
      state[typeId][key] = undefined;
    });
  }

  configIsOpen[typeId] = !configIsOpen[typeId];
};

/**
 * Submits the options configuration of an action/reaction by redirecting them to the workflow page.
 * @param typeId - The ID of the action/reaction being submitted.
 * @param typeTitle - The title of the action/reaction being submitted.
 */
const onSubmit = (typeId: number, typeTitle: string) => {
  const modifiedOptions = { ...state[typeId] };

  const hasInvalidTypes = Object.entries(modifiedOptions).some(
    ([key, value]) => {
      const expectedType = fieldTypes[typeId][key];
      return typeof value !== expectedType;
    }
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

/**
 * The types of each field in the options configuration of each action/reaction.
 */
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
        })
      ),
    ])
  )
);

/**
 * Formats a string by adding spaces between camel case words.
 * @param str - The string to be formatted.
 * @returns The formatted string.
 */
function formatString(str: string): string {
  return str.replace(/([a-z])([A-Z])/g, "$1 $2");
}

/**
 * Checks if a string has a word with more than 8 characters.
 * @param text - The text to be checked.
 * @returns Whether the text has a word with more than 8 characters.
 */
const hasLongWord = (text: string): boolean => {
  const words = formatString(text).split(" ");
  for (const word of words) {
    if (word.length > 8) {
      return true;
    }
  }
  return false;
};

/**
 * Counts the number of words in a text.
 * @param text - The text to be counted.
 * @returns The number of words in the text.
 */
function countWords(text: string) {
  return text.trim().split(/\s+/).length;
}
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <div v-for="type in props.types" :key="type.id">
      <UContainer
        :ui="{ padding: '!px-0 !py-6', constrained: 'max-w-none' }"
        :class="[
          'custom_card flex flex-col justify-evenly items-center gap-4 text-white',
        ]"
        :style="{ backgroundColor: props.serviceInfo?.color || 'black' }"
        tabindex="0"
        @click="toggleConfig(type.id)"
        @keydown.space="toggleConfig(type.id)"
      >
        <!-- , countWords(formatString(type.name)) > 3 ? '!w-[14vw]' : '' -->
        <h4
          :class="[
            'p-2 capitalize text-center break-words w-full',
            hasLongWord(type.name) && countWords(formatString(type.name)) > 1
              ? 'leading-[120%]'
              : '',
            !hasLongWord(type.name) && countWords(formatString(type.name)) > 2
              ? 'leading-[120%]'
              : '',
          ]"
        >
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
            class="flex flex-col gap-12 max-sm:gap-8 p-5 max-lg:p-4 max-md:p-3 max-sm:p-2 w-full bg-custom_color-bg_section"
            @submit.prevent="onSubmit(type.id, type.name)"
          >
            <h2 class="text-center">
              {{ formatString(type.name) }}
            </h2>
            <h6
              class="text-center self-center -mt-6 max-w-[80%] max-lg:max-w-[85%] max-md:max-w-[90%] max-sm:max-w-[95%]"
            >
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
              class="self-center min-w-[80%] max-lg:min-w-[85%] max-md:min-w-[90%] max-sm:min-w-[95%]"
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
                class="px-6 py-5 max-lg:py-3 max-md:py-2 text-custom_color-text bg-opacity-0 border-custom_border_width !border-custom_color-border"
                tabindex="0"
                @click="toggleConfig(type.id)"
              >
                <h6>Cancel</h6>
              </UButton>
              <UButton
                type="submit"
                class="px-6 py-5 max-lg:py-3 max-md:py-2"
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
