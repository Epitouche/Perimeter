<script setup lang="ts">
import type { Area } from "@/interfaces/areas";

/**
 * Which area's option to update
 */
const props = defineProps<{
  areaId: number; // Area to update (by id)
  typeName: string; // Type of area option name
  type: Area["action"] | Area["reaction"]; // Type of area option to update
  typeOptions: string | object | undefined; // Current type options for the area
  color: string; // Action service color
}>();

const router = useRouter();

const isOpen = ref(false);

/**
 * Emit event to update area option value
 */
const emit = defineEmits<{
  (
    event: "updateAreaValue",
    areaId: number,
    typeName: string,
    key: string,
    value: string | number
  ): void;
}>();

/**
 * State to hold the updated values
 */
const state = reactive<{ [key: number]: Record<string, string | number> }>(
  typeof props.typeOptions === "string"
    ? { [props.type.id]: JSON.parse(props.typeOptions) }
    : { [props.type.id]: props.typeOptions || {} }
);

/**
 * Send update information to page with function to send updated area option value to backend
 */
const editValue = async (typeName: string, typeId: number, key: string) => {
  const updatedValues = { ...state[typeId] };
  const updatedValue = updatedValues[key];

  emit("updateAreaValue", props.areaId, typeName, key, updatedValue);

  router.push({
    name: "myareas",
    query: {
      areaId: props.areaId.toString(),
      typeName: typeName,
      keyString: key,
      value: updatedValue,
    },
  });
  toggleSlideover();
};

/**
 * Toggle the edit options slideover
 */
const toggleSlideover = () => {
  isOpen.value = !isOpen.value;
};

/**
 * The countWords function counts the number of words in the text.
 */
 function countWords(text: string) {
  return text.trim().split(/\s+/).length;
}

/**
 * Format the name of the area option
 */
function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-4 !py-4', constrained: 'max-w-full' }"
    class="capitalize self-start flex flex-row justify-between items-center gap-5 max-sm:gap-2 border-custom_border_width !border-white rounded-custom_border_radius w-full"
  >
    <div
      class="flex flex-row justify-start max-sm:justify-between items-center gap-8 max-sm:gap-2 w-full"
    >
      <img
        :src="type.service.icon"
        :alt="type.service.name"
        style="width: 10%"
        class="max-sm:hidden"
      />
      <h4 class="text-center leading-[100%]">
        <b>{{ formatName(type.service.name) }}</b
        >:
      </h4>
      <h5 v-if="countWords(formatName(type.name)) < 3" class="text-center leading-[100%]">{{ formatName(type.name) }}</h5>
      <h6 v-else class="text-center leading-[100%]">{{ formatName(type.name) }}</h6>
    </div>
    <UButton
      color="white"
      :ui="{ rounded: 'rounded-full' }"
      class="w-[3.1vw] h-[5.5vh] max-lg:w-[3.4vw] max-lg:h-[5.2vh] max-md:w-[3.8vw] max-md:h-[4.9vh] max-sm:w-[4vw] max-sm:h-[2vh] shadow-2xl active:shadow-sm transition-shadow"
      @click="toggleSlideover"
    >
      <UIcon
        name="i-bytesize-edit"
        class="w-[95%] h-[95%]"
        :style="{ color: color }"
      />
    </UButton>
  </UContainer>
  <USlideover v-model="isOpen">
    <UForm
      :state="state[type.id]"
      class="flex flex-col justify-center items-center gap-5 py-10 bg-custom_color-bg_section"
    >
      <UFormGroup
        v-for="(value, key) in state[type.id]"
        :key="key"
        :label="key"
        :name="key"
        :ui="{ label: { base: 'capitalize text-xl pl-3' } }"
      >
        <div class="flex flex-row justify-center items-center gap-3">
          <UInput
            v-model="state[type.id][key] as string | number | undefined"
            :ui="{
              placeholder: '!px-5 !py-2 font-light',
              size: { sm: 'text-lg' },
            }"
            :placeholder="key + '...'"
          />
          <UButton @click="editValue(typeName, type.id, key)">
            <UIcon name="i-bytesize-checkmark" />
          </UButton>
        </div>
      </UFormGroup>
    </UForm>
  </USlideover>
</template>

<style scoped>
:deep(div.z-20.group.w-48) {
  width: fit-content;
}
</style>
