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

onMounted(() => {
  getServiceInfo();
  console.log("typeName:", props.typeName);
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
  <div 
  v-if="isSelected" :class="[
    'flex flex-row justify-evenly items-center py-12 px-12 gap-10 rounded-3xl w-full',
    isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
  ]" :style="{ backgroundColor: serviceInfo ? serviceInfo.color : 'black' }">
    <img 
    :src="serviceInfo ? `${serviceInfo.icon}` : ''" :alt="serviceInfo ? `${serviceInfo.name}` : ''"
      class="w-16 h-16 p-0">
    <h2 
    :class="['text-white text-7xl font-custom_weight_title', isDisabled ? 'text-opacity-50' : 'text-opacity-100',]">
      {{ typeName }}
    </h2>
  </div>

  <div 
  v-else :class="[
    'flex flex-row justify-evenly items-center bg-black py-12 px-12 gap-10 rounded-3xl w-full',
    isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
  ]">
    <h2 
    :class="[
      'text-white text-8xl font-custom_weight_title',
      isDisabled ? 'text-opacity-50' : 'text-opacity-100',
    ]">
      {{ title }}
    </h2>
    <UButton :disabled="isDisabled" :to="link" :ui="{ rounded: 'rounded-2xl' }" :class="[
      'text-black bg-white text-5xl font-extrabold px-10 py-3',
      isDisabled ? '!text-opacity-60' : 'text-opacity-100',
    ]">Add</UButton>
  </div>
</template>

<style scoped></style>
