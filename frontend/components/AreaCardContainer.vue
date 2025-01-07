<script setup lang="ts">
import type { Area } from "@/interfaces/areas";

const props = defineProps<{
  areas: Area[];
}>();

const areaIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.areas.map((area) => [area.id, false])),
);

const confirmDeletionIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.areas.map((area) => [area.id, false])),
);

const toggleAreaModal = (areaId: number) => {
  areaIsOpen[areaId] = !areaIsOpen[areaId];
};

const toggleConfirmDeletionModal = (areaId: number) => {
  confirmDeletionIsOpen[areaId] = !confirmDeletionIsOpen[areaId];
};

const onDelete = (areaId: number) => {
  if (confirmDeletionIsOpen[areaId]) {
    console.log("Delete area here", areaId);
    toggleConfirmDeletionModal(areaId);
    return;
  }
  toggleConfirmDeletionModal(areaId);
  toggleAreaModal(areaId);
};

const cancelDeletion = (areaId: number) => {
  toggleConfirmDeletionModal(areaId);
  toggleAreaModal(areaId);
};

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}

onMounted(() => {
  console.log("areas", props.areas);
});
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full max-h-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap py-5 w-full h-full"
  >
    <div v-for="area in areas" :key="area.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        :class="[
          `bg-custom_color-${area.action.service.name}`,
          'flex justify-center items-center text-white font-extrabold text-6xl rounded-custom_border_radius w-[5em] h-[4.5em]',
        ]"
        @click="toggleAreaModal(area.id)"
      >
        <h2
          class="clamp-2-lines capitalize text-5xl text-center break-words w-full"
        >
          {{ formatName(area.action.name) }}
        </h2>
      </UContainer>

      <UModal
        v-model="areaIsOpen[area.id]"
        :ui="{
          background: `bg-custom_color-${area.action.service.name}`,
          width: 'w-2/5',
          base: 'relative flex flex-col gap-14 text-5xl font-semibold text-white pl-20 pr-12 py-10',
        }"
      >
        <div class="flex flex-row justify-between pb-2 w-full">
          <h2 class="text-6xl self-center w-full">Temp title</h2>
          <UButton
            variant="ghost"
            class="self-end w-fit"
            @click="toggleAreaModal(area.id)"
          >
            <UIcon name="i-bytesize-close" class="w-10 h-10 p-0" />
          </UButton>
        </div>
        <div class="capitalize self-start flex flex-row items-center gap-5">
          <UIcon name="i-bytesize-ban" class="w-16 h-16 p-0" />
          <p>
            {{ area.action.service.name }}: {{ formatName(area.action.name) }}
          </p>
        </div>
        <div class="capitalize self-start flex flex-row items-center gap-5">
          <UIcon name="i-bytesize-ban" class="w-16 h-16 p-0" />
          <p>
            {{ area.reaction.service.name }}:
            {{ formatName(area.reaction.name) }}
          </p>
        </div>
        <p class="self-start">Description: Desc will go here</p>
        <UTooltip text="Delete" class="self-end w-fit">
          <UButton
            variant="ghost"
            class="hover_underline_animation items-end w-fit p-0 pb-1"
            @click="onDelete(area.id)"
          >
            <UIcon name="i-bytesize-trash" class="w-12 h-12 text-white" />
          </UButton>
        </UTooltip>
      </UModal>
      <UModal
        v-model="confirmDeletionIsOpen[area.id]"
        :ui="{
          base: `relative text-left rtl:text-right flex flex-col gap-10 p-10 border-custom_border_width border-custom_color-${area.action.service.name}`,
        }"
      >
        <h2 class="text-4xl font-semibold">
          Are you sure you want to delete this area?
        </h2>
        <p class="text-2xl">This action cannot be undone!</p>
        <div class="flex flex-row justify-end items-center gap-5 pt-5">
          <UButton
            :class="[
              `text-custom_color-${area.action.service.name} !border-custom_color-${area.action.service.name}`,
              'bg-opacity-0 border-custom_border_width text-2xl font-semibold py-3 px-5',
            ]"
            @click="cancelDeletion(area.id)"
            >Cancel</UButton
          >
          <UButton
            :class="[
              `bg-custom_color-${area.action.service.name}`,
              'text-white text-2xl font-semibold py-3 px-5',
            ]"
            @click="onDelete(area.id)"
            >Delete</UButton
          >
        </div>
      </UModal>
    </div>
  </UContainer>
</template>

<style scoped>
.hover_underline_animation {
  display: inline-block;
  position: relative;
}

.hover_underline_animation::after {
  content: "";
  position: absolute;
  width: 100%;
  transform: scaleX(0);
  height: 0.15em;
  bottom: 0;
  left: 0;
  background-color: white;
  transition: transform 0.2s ease-out;
  transform-origin: bottom center;
}

.hover_underline_animation:hover::after {
  transform: scaleX(0.9);
  transform-origin: bottom center;
}
</style>
