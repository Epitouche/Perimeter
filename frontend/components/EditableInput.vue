<script lang="ts">
import { ref, watch } from "vue";

export default {
  name: "EditableInput",
  /**
   * The name of the input field
   * @type {String}
   * @required
   * The value of the input field
   * @type {String}
   * @default ""
   */
  props: {
    name: {
      type: String,
      required: true,
    },
    modelValue: {
      type: String,
      default: "",
    },
  },
  /**
   * emits the updated value of the input field
   */
  emits: ["update:modelValue"],
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
      isEditable.value = false;
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
  <div class="flex flex-col w-full">
    <h5 class="px-5">{{ name }}</h5>
    <div class="flex items-center gap-4 justify-center">
      <UInput
        v-model="inputValue"
        :disabled="!isEditable"
        :ui="{
          placeholder: '!px-5 !py-3 max-md:!py-1 font-light',
          size: {
            sm: 'text-3xl max-lg:text-2xl max-md:text-xl max-sm:text-md',
          },
        }"
        :class="{
          'bg-gray-100 text-gray-500 cursor-not-allowed rounded-full':
            !isEditable,
          'bg-white text-black rounded-full': isEditable,
        }"
        class="flex-1 transition-colors duration-300"
      />
    </div>
  </div>
</template>

<style scoped></style>
