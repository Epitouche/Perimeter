<script setup lang="ts">
/**
 * The ReActionButton component is a reusable component that displays which action and reactions were chosen by the user.
 */
const props = defineProps<{
  title: string; // The title of the type of button
  link: string; // The link to the page
  isDisabled: boolean; // Whether the button is disabled
  isSelected: boolean; // Whether the action or reaction has been selected
  serviceId: number | null; // The ID of the service
  typeName: string; // The name of the action or reaction
}>();

const { serviceId } = toRefs(props);
const { isSelected } = toRefs(props);

const token = useCookie("token");

/**
 * The serviceInfo object contains information about the service.
 */
const serviceInfo = ref<{ name: string; color: string; icon: string } | null>(
  null
);

const error = ref<string | null>(null);

/**
 * The getServiceInfo function fetches the service information.
 */
const getServiceInfo = async () => {
  if (serviceId.value) {
    try {
      error.value = null;
      serviceInfo.value = await $fetch("/api/servicebyid", {
        method: "POST",
        body: {
          token: token.value,
          serviceId: serviceId.value,
        },
      });
      console.log("services", serviceInfo.value);
    } catch (err) {
      console.error("Error fetching services:", err);
    }
  }
};

/**
 * The formatName function formats the name of the action or reaction.
 */
function formatName(name: string): string {
  return name
    .replace(/^action_/, "")
    .replace(/_/g, " ")
    .replace(/([a-z])([A-Z])/g, "$1 $2");
}

/**
 * The isLongText function checks if the text is longer than 12 characters.
 */
const isLongText = (text: string): boolean => text.length > 12;

/**
 * The countWords function counts the number of words in the text.
 */
function countWords(text: string) {
  return text.trim().split(/\s+/).length;
}

/**
 * When the component is mounted, the getServiceInfo function is called.
 */
onMounted(() => {
  getServiceInfo();
});

/**
 * The watch function watches for changes in the isSelected prop.
 */
watch(
  isSelected,
  (newValue) => {
    if (newValue) {
      getServiceInfo();
    }
  },
  { immediate: true }
);
</script>

<template>
  <UContainer
    v-if="isSelected"
    :ui="{ padding: '!px-4 !py-4', constrained: '!min-w-none !min-h-none' }"
    :class="[
      'flex flex-row justify-evenly items-center gap-2 rounded-3xl w-[30vw] h-[14vh] max-lg:w-[45vw] max-md:w-[60vw] max-sm:w-[75vw] max-lg:h-[12vh] max-md:h-[12vh]',
      isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
    ]"
    :style="{ backgroundColor: serviceInfo ? serviceInfo.color : 'black' }"
    tabindex="-1"
  >
    <img
      :src="serviceInfo ? `${serviceInfo.icon}` : ''"
      :alt="serviceInfo ? `${serviceInfo.name}` : ''"
      class="p-0"
      :style="{ width: isLongText(typeName) ? '20%' : '15%' }"
    />
    <h3
      :class="[
        'text-white text-center break-words whitespace-normal leading-[100%]',
        isDisabled ? 'text-opacity-50' : 'text-opacity-100',
        isLongText(typeName) && countWords(formatName(typeName)) === 2
          ? 'w-min'
          : '',
      ]"
    >
      {{ formatName(typeName) }}
    </h3>
  </UContainer>

  <UContainer
    v-else
    :ui="{ padding: '!px-0 !py-0', constrained: '!min-w-none !min-h-none' }"
    :class="[
      'flex flex-row justify-evenly items-center bg-black rounded-3xl w-[30vw] h-[14vh] max-lg:w-[45vw] max-md:w-[60vw] max-sm:w-[75vw] max-lg:h-[12vh] max-md:h-[12vh]',
      isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
    ]"
  >
    <h2
      :class="[
        'text-white',
        isDisabled ? 'text-opacity-50' : 'text-opacity-100',
      ]"
    >
      {{ title }}
    </h2>
    <UButton
      :disabled="isDisabled"
      :to="link"
      :ui="{ rounded: 'rounded-2xl max-lg:rounded-xl' }"
      :class="[
        'text-black bg-white px-5 py-5 max-lg:py-4 max-md:py-3 max-sm:py-2 max-md:px-4',
        isDisabled ? '!text-opacity-60' : 'text-opacity-100',
      ]"
    >
      <h5 class="text-center w-full">Add</h5>
    </UButton>
  </UContainer>
</template>

<style scoped>
[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
