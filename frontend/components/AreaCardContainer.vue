<script setup lang="ts">
import type { Area } from "@/interfaces/areas";

const props = defineProps<{
  areas: Area[];
}>();

const token = useCookie("token");
const errorMessage = ref<string | null>(null);

const componentKey = ref(0);

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

const onDelete = async (areaId: number) => {
  if (confirmDeletionIsOpen[areaId]) {
    try {
      errorMessage.value = null;
      const response = await $fetch("/api/area/delete", {
        method: "POST",
        body: {
          token: token.value,
          id: areaId,
        },
      });
      console.log("response:", response);
    } catch (error: unknown) {
      console.log("error:", error);
      errorMessage.value = handleErrorStatus(error);
      if (errorMessage.value === "An unknown error occurred") {
        console.error("An unknown error occurred", error);
      }
    }
    componentKey.value += 1;
    // toggleConfirmDeletionModal(areaId);
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
  console.log("areas in AreaCardContainer", props.areas);
});
</script>

<template>
  <UContainer
    :key="componentKey"
    :ui="{ padding: '!px-0', constrained: 'max-w-full max-h-full' }"
    class="flex flex-row justify-center items-center gap-10 flex-wrap py-5 w-full h-full"
  >
    <div v-for="area in areas" :key="area.id">
      <UContainer
        :ui="{ padding: 'px-0', constrained: 'max-w-none' }"
        class="flex flex-col justify-center items-center text-white font-extrabold text-6xl rounded-custom_border_radius w-[5em] h-[4.5em]"
        :style="{ backgroundColor: area.action.service.color }"
        @click="toggleAreaModal(area.id)"
      >
        <h2
          class="clamp-2-lines capitalize text-4xl text-center break-words pb-2 w-full"
        >
          {{ formatName(area.action.name) }}
        </h2>
        <div class="grid place-items-center h-36 relative w-full">
          <img
            :src="area.action.service.icon"
            :alt="area.action.service.name"
            class="w-24 h-24 p-0 absolute top-1 left-12"
          />
          <img
            :src="area.reaction.service.icon"
            :alt="area.reaction.service.name"
            class="w-24 h-24 p-0 absolute bottom-0 right-12"
          />
        </div>
      </UContainer>
      <UModal
        v-model="areaIsOpen[area.id]"
        :ui="{
          width: 'w-2/5',
        }"
      >
        <div
          class="flex flex-col gap-14 font-semibold text-white rounded-custom_border_radius pl-20 pr-12 py-10 w-full"
          :style="{ backgroundColor: area.action.service.color }"
        >
          <div class="flex flex-row justify-between pb-2 w-full">
            <h2 class="text-6xl text-center w-full"><b>Temp title</b></h2>
            <UButton
              variant="ghost"
              class="self-end w-fit"
              @click="toggleAreaModal(area.id)"
            >
              <UIcon name="i-bytesize-close" class="w-12 h-12 text-white" />
            </UButton>
          </div>
          <div
            class="capitalize self-start flex flex-row items-center text-5xl gap-5"
          >
            <img
              :src="area.action.service.icon"
              :alt="area.action.service.name"
              class="w-16 h-16 p-0"
            />
            <p>
              <b>{{ area.action.service.name }}</b
              >: {{ formatName(area.action.name) }}
            </p>
          </div>
          <div
            class="capitalize self-start flex flex-row items-center text-5xl gap-5"
          >
            <img
              :src="area.reaction.service.icon"
              :alt="area.reaction.service.name"
              class="w-16 h-16 p-0"
            />
            <p>
              <b>{{ area.reaction.service.name }}</b
              >:
              {{ formatName(area.reaction.name) }}
            </p>
          </div>
          <p class="self-start text-5xl">
            <b>Description</b>: Desc will go here
          </p>
          <UTooltip text="Delete" class="self-end w-fit">
            <UButton
              variant="ghost"
              class="hover_underline_animation items-end w-fit p-0 pb-1"
              @click="onDelete(areaId)"
            >
              <UIcon name="i-bytesize-trash" class="w-12 h-12 text-white" />
            </UButton>
          </UTooltip>
        </div>
      </UModal>
      <UModal
        v-model="confirmDeletionIsOpen[area.id]"
        :ui="{
          base: 'relative text-left rtl:text-right flex flex-col gap-10 p-10 border-custom_border_width',
        }"
        :style="{ borderColor: area.action.service.color }"
      >
        <h2 class="text-4xl font-semibold">
          Are you sure you want to delete this area?
        </h2>
        <p class="text-2xl">This action cannot be undone!</p>
        <div class="flex flex-row justify-end items-center gap-5 pt-5">
          <UButton
            class="bg-opacity-0 border-custom_border_width text-2xl font-semibold py-3 px-5"
            :style="{
              borderColor: area.action.service.color,
              color: area.action.service.color,
            }"
            @click="cancelDeletion(area.id)"
            >Cancel</UButton
          >
          <UButton
            class="text-white text-2xl font-semibold py-3 px-5"
            :style="{ backgroundColor: area.action.service.color }"
            @click="onDelete(areaId)"
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
  transition: transform 0.1s ease-out;
  transform-origin: bottom center;
}

.hover_underline_animation:hover::after {
  transform: scaleX(0.9);
  transform-origin: bottom center;
}
</style>
