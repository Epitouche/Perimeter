<script setup lang="ts">
const props = defineProps<{
  title: string;
  link: string;
  isDisabled: boolean;
  isSelected: boolean;
  serviceId: number | null;
  typeName: string;
}>();

const { serviceId } = toRefs(props);
const { isSelected } = toRefs(props);

const token = useCookie("token");

const serviceInfo = ref<{ name: string; color: string; icon: string } | null>(
  null,
);
const error = ref<string | null>(null);

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

function formatName(name: string): string {
  console.log("typeName: ", props.typeName); ///////////////
  console.log("typeName.length: ", props.typeName.length); ///////////
  return name
    .replace(/^action_/, "")
    .replace(/_/g, " ")
    .replace(/([a-z])([A-Z])/g, "$1 $2");
}

const isLongText = (text: string): boolean => text.length > 12;

function countWords(text: string) {
  return text.trim().split(/\s+/).length;
}

onMounted(() => {
  getServiceInfo();
});

watch(
  isSelected,
  (newValue) => {
    if (newValue) {
      getServiceInfo();
    }
  },
  { immediate: true },
);
</script>

<template>
  <UContainer
    v-if="isSelected"
    :ui="{ padding: '!px-4 !py-4', constrained: '!min-w-none !min-h-none' }"
    :class="[
      'flex flex-row justify-evenly items-center gap-2 rounded-3xl w-[30vw] h-[14vh] max-lg:w-[45vw] max-md:w-[60vw] max-sm:w-[70vw] max-lg:h-[12vh] max-md:h-[10vh] max-sm:h-[9vh]',
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
      'flex flex-row justify-evenly items-center bg-black rounded-3xl w-[30vw] h-[14vh] max-lg:w-[45vw] max-md:w-[60vw] max-sm:w-[70vw] max-lg:h-[12vh] max-md:h-[10vh] max-sm:h-[9vh]',
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
        'text-black bg-white w-[6vw] h-[4.8vh] max-lg:w-[8vw] max-lg:h-[4vh] max-md:w-[10vw] max-md:h-[3.5vh] max-sm:w-[10vw] max-sm:h-[3vh]',
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
