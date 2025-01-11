<script lang="ts">
import { ref, watch } from "vue";

export default {
  name: "EditableInput",
  props: {
    name: {
      type: String,
      required: true,
    },
    modelValue: {
      type: String,
      default: '',
    },
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const isEditable = ref(false);
    const inputValue = ref(props.modelValue);

    watch(
      () => props.modelValue,
      (newValue) => {
        inputValue.value = newValue;
      }
    );

    const toggleEdit = () => {
      isEditable.value = !isEditable.value;
      if (!isEditable.value) {
        emit("update:modelValue", inputValue.value);
      }
    };

    return {
      isEditable,
      inputValue,
      toggleEdit,
    };
  },
};
</script>

<template>
  <div class="flex flex-col w-[50%]">
    <h2 class="text-xl px-5">{{ name }}</h2>
    <div class="flex items-center gap-3 justify-center">
      <UInput
        v-model="inputValue"
        :disabled="!isEditable"
        :ui="{
          placeholder: '!px-5 !py-3 font-light',
          size: { sm: 'text-3xl' },
        }"
        :class="{
          'bg-gray-100 text-gray-500 cursor-not-allowed rounded-full': !isEditable,
          'bg-white text-black rounded-full': isEditable,
        }"
        class="flex-1 transition-colors duration-300"
      />
      <UButton
        class="h-12 bg-white border-2 border-black rounded-full flex items-center justify-center shadow-lg text-black font-semibold transition-colors duration-300"
        @click="toggleEdit"
      >
        <Icon
          :name="isEditable ? 'bytesize:edit' : 'mdi:lock'"
          class="text-black text-xl"
        />
      </UButton>
    </div>
  </div>
</template>

<style scoped>
</style>
