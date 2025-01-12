<script setup lang="ts">
import type { Area } from "@/interfaces/areas";
import { defineEmits } from "vue";

const props = defineProps<{
  areaId: number;
  typeName: string;
  type: Area["action"] | Area["reaction"];
  color: string;
}>();

const router = useRouter();

const isOpen = ref(false);

const emit = defineEmits<{
  (
    event: "updateAreaValue",
    areaId: number,
    typeName: string,
    key: string,
    value: string | number,
  ): void;
}>();

const state = reactive<{ [key: number]: Record<string, string | number> }>(
  typeof props.type.option === "string"
    ? { [props.type.id]: JSON.parse(props.type.option) }
    : { [props.type.id]: props.type.option || {} },
);

const editValue = async (typeName: string, typeId: number, key: string) => {
  // Assuming state[typeId] returns an object and you want the key and its updated value
  const updatedValues = { ...state[typeId] };

  // Get the value for the edited key (this is dynamic, not always the first key)
  const updatedValue = updatedValues[key];

  console.log("editedKey: ", key); // Log the key that was edited
  console.log("updatedValue: ", updatedValue); // Log the updated value

  // Emit the event with appropriate types
  // Ensure areaId is a number and key is a valid key (should be string)
  emit("updateAreaValue", props.areaId, typeName, key, updatedValue);

  // If you want to navigate, ensure the query parameters are also typed correctly
  router.push({
    name: "myareas",
    query: {
      areaId: props.areaId.toString(), // Convert areaId to string
      typeName: typeName,
      keyString: key, // key should remain string if it's a string key in your data
      value: updatedValue, // First value from updatedValues
    },
  });
  toggleSlideover();
};

const toggleSlideover = () => {
  isOpen.value = !isOpen.value;
};

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}

onMounted(() => {
  console.log("type: ", props.type);
});
</script>

<template>
  <div
    class="capitalize self-start flex flex-row items-center gap-5 border-custom_border_width border-white rounded-custom_border_radius w-fit py-2 px-4"
  >
    <img
      :src="type.service.icon"
      :alt="type.service.name"
      class="w-16 h-16 p-0"
    >
    <h2 class="text-5xl">
      <b>{{ type.service.name }}</b
      >:
    </h2>
    <p class="text-4xl">{{ formatName(type.name) }}</p>
    <UButton
      color="white"
      :ui="{ rounded: 'rounded-full' }"
      class="w-11 h-11 shadow-2xl active:shadow-sm transition-shadow"
      @click="toggleSlideover"
    >
      <UIcon name="i-bytesize-edit" class="w-7 h-7" :style="{ color: color }" />
    </UButton>
  </div>
  <USlideover v-model="isOpen" class="">
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
