<script setup lang="ts">
import services from '~/server/api/workflow/services';

const props = defineProps<{
  title: string;
  link: string;
  isDisabled: boolean;
  isSelected: boolean;
  canDelete: boolean;
  serviceId: number | null;
}>();

const { serviceId } = toRefs(props);

const token = useCookie("token");

const serviceInfo = ref<{ name: string } | null>(null);
const error = ref<string | null>(null);

const getServiceInfo = async () => {
  if (serviceId) {
    try {
      error.value = null;
      serviceInfo.value = await $fetch("/api/servicebyid", {
        method: "POST",
        body: {
          token: token.value,
          serviceId: serviceId,
        },
      });
      console.log("services", serviceInfo.value);
    } catch (error) {
      console.error("Error fetching services:", error);
    }
  }
};

onMounted(() => {
  getServiceInfo();
});

</script>

<template>
  <div v-if="isSelected" :class="[
    'flex flex-row justify-evenly py-12 px-12 gap-10 rounded-3xl w-full', `bg-custom_color-${serviceInfo?.name}`, isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
  ]">
    <!-- <UIcon :name="" /> -->
    <h2 :class="[
      'text-white text-8xl font-custom_weight_title',
      isDisabled ? 'text-opacity-50' : 'text-opacity-100',
    ]">
      {{ serviceInfo ? serviceInfo.name : '' }}
    </h2>
    <div v-if="canDelete" class="flex items-end">
      <UButton class="bg-opacity-0 text-4xl font-bold px-10 py-3"><u>Delete</u></UButton>
    </div>
  </div>
  <div v-else :class="[
    'flex flex-row justify-evenly items-center bg-black py-12 px-12 gap-10 rounded-3xl w-full',
    isDisabled ? 'bg-opacity-60' : 'bg-opacity-100',
  ]">
    <h2 :class="[
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
