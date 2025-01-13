<script setup lang="ts">
import type { LocationQueryValue } from "vue-router";
import type { Area } from "@/interfaces/areas";

const props = defineProps<{
  areas: Area[];
}>();

const token = useCookie("token");
const errorMessage = ref<string | null>(null);
const router = useRouter();
const route = useRoute();

const getQueryParam = (
  param: LocationQueryValue | LocationQueryValue[] | undefined,
): string | null => {
  if (Array.isArray(param)) {
    return param.length > 0 ? String(param[0]) : null;
  }
  return param ? String(param) : null;
};

const areaId = getQueryParam(route.query.areaId);
const typeName = getQueryParam(route.query.typeName);
const keyString = getQueryParam(route.query.keyString);
const value = getQueryParam(route.query.value);

const emit = defineEmits(["refreshAreas"]);

const areaIdNumber = areaId ? Number(areaId) : null;
const valueNumber = value ? Number(value) : null;

if (areaIdNumber !== null && isNaN(areaIdNumber)) {
  console.error("Invalid areaId:", areaId);
}
if (valueNumber !== null && isNaN(valueNumber)) {
  console.error("Invalid value:", value);
}

const componentKey = ref(0);

const areaIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.areas.map((area) => [area.id, false])),
);

const areaIsEnabled = (areaId: number) => {
  const areaIndex = props.areas.findIndex((area) => area.id === areaId);
  if (areaIndex === -1) {
    console.error("Area not found");
    return false;
  }
  return props.areas[areaIndex].enable;
};

const confirmDeletionIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.areas.map((area) => [area.id, false])),
);

const toggleAreaModal = (areaId: number) => {
  areaIsOpen[areaId] = !areaIsOpen[areaId];
};

const toggleAreaEnableSwitch = async (areaId: number) => {
  const areaIndex = props.areas.findIndex((area) => area.id === areaId);
  if (areaIndex === -1) {
    console.error("Area not found");
    return;
  }

  const updatedArea = JSON.parse(
    JSON.stringify(props.areas[areaIndex]),
  ) as Area;
  updatedArea.enable = !updatedArea.enable;

  console.log("updatedArea after toggling enable:", updatedArea);

  try {
    errorMessage.value = null;

    const response = await $fetch("/api/area/update", {
      method: "POST",
      body: {
        token: token.value,
        area: updatedArea,
      },
    });

    console.log("response:", response);

    emit("refreshAreas");
  } catch (error) {
    console.log("error:", error);
    errorMessage.value = handleErrorStatus(error);
    alert("Failed to update enable/disable status");
  }
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
          areaId: areaId,
        },
      });
      console.log("response:", response);
      emit("refreshAreas");
    } catch (error: unknown) {
      console.log("error:", error);
      errorMessage.value = handleErrorStatus(error);
      if (errorMessage.value === "An unknown error occurred") {
        console.error("An unknown error occurred", error);
      }
    }
    toggleConfirmDeletionModal(areaId);
    componentKey.value += 1;
    return;
  }
  toggleConfirmDeletionModal(areaId);
  toggleAreaModal(areaId);
  return;
};

const cancelDeletion = (areaId: number) => {
  toggleConfirmDeletionModal(areaId);
  toggleAreaModal(areaId);
};

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}

const updateAreaValue = async (
  areaId: number,
  typeName: string,
  keyString: string,
  value: string | number,
) => {
  console.log(
    "areaId:",
    areaId,
    "typeName:",
    typeName,
    "keyString:",
    keyString,
    "value:",
    value,
  );

  const areaIndex = props.areas.findIndex((area) => area.id === areaId);
  if (areaIndex === -1) {
    console.error("Area not found");
    return;
  }

  const updatedArea = JSON.parse(
    JSON.stringify(props.areas[areaIndex]),
  ) as Area;

  if (typeName === "action") {
    if (!updatedArea.action_option) {
      updatedArea.action_option = {};
    }

    (updatedArea.action_option as { [key: string]: string | number })[
      keyString
    ] =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;
    (updatedArea.action.option as { [key: string]: string | number })[
      keyString
    ] =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;

    console.log("After updating action_option:", updatedArea.action_option);
  } else if (typeName === "reaction") {
    if (!updatedArea.reaction_option) {
      updatedArea.reaction_option = {};
    }

    (updatedArea.reaction_option as { [key: string]: string | number })[
      keyString
    ] =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;
    (updatedArea.reaction.option as { [key: string]: string | number })[
      keyString
    ] =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;

    console.log("After updating reaction_option:", updatedArea.reaction_option);
  } else {
    console.error("Invalid typeName:", typeName);
    return;
  }

  console.log("Final updatedArea:", updatedArea);

  try {
    errorMessage.value = null;

    const response = await $fetch("/api/area/update", {
      method: "POST",
      body: {
        token: token.value,
        area: updatedArea,
      },
    });

    console.log("response:", response);
    emit("refreshAreas");
  } catch (error) {
    console.log("error:", error);
    errorMessage.value = handleErrorStatus(error);
  }

  router.push("myareas");
};

onMounted(() => {
  console.log("areas in AreaCardContainer", props.areas);
});

if (areaIdNumber !== null && valueNumber !== null) {
  updateAreaValue(areaIdNumber, typeName!, keyString!, valueNumber);
}
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
          {{ formatName(area.title) }}
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
          width: 'w-1/2',
        }"
      >
        <div
          class="flex flex-col gap-14 font-semibold text-white rounded-custom_border_radius pl-20 pr-12 py-10 w-full"
          :style="{ backgroundColor: area.action.service.color }"
        >
          <div>
            <div class="flex flex-row justify-between items-center w-full">
              <div class="flex flex-row items-center gap-3">
                <UToggle
                  size="xl"
                  :model-value="areaIsEnabled(area.id)"
                  @update:model-value="toggleAreaEnableSwitch(area.id)"
                />
                <div v-if="areaIsEnabled(area.id)" class="text-xl">
                  <p>Enabled</p>
                </div>
                <div v-else class="text-xl">
                  <p>Disabled</p>
                </div>
              </div>
              <UButton
                variant="ghost"
                class="self-end w-fit"
                @click="toggleAreaModal(area.id)"
              >
                <UIcon name="i-bytesize-close" class="w-12 h-12 text-white" />
              </UButton>
            </div>

            <h2 class="text-6xl text-center w-full">
              <b>{{ area.title }}</b>
            </h2>
          </div>

          <UpdateAreaOptions
            :area-id="area.id"
            type-name="action"
            :color="area.action.service.color"
            :type="area.action"
            @update-area-value="updateAreaValue"
          />
          <UpdateAreaOptions
            :area-id="area.id"
            type-name="reaction"
            :color="area.action.service.color"
            :type="area.reaction"
            @update-area-value="updateAreaValue"
          />

          <div>
            <p class="self-start text-5xl pb-2"><b>Description</b>:</p>
            <p class="text-4xl">{{ area.description }}</p>
          </div>

          <div class="flex flex-row justify-end items-center gap-5">
            <UTooltip text="Edit" class="self-end w-fit">
              <UButton
                variant="ghost"
                class="hover_underline_animation items-end w-fit p-0 pb-1"
                @click="onEdit(area.id)"
              >
                <UIcon name="i-bytesize-edit" class="w-11 h-11 text-white" />
              </UButton>
            </UTooltip>

            <UTooltip text="Delete" class="self-end w-fit">
              <UButton
                variant="ghost"
                class="hover_underline_animation items-end w-fit p-0 pb-1"
                @click="onDelete(area.id)"
              >
                <UIcon name="i-bytesize-trash" class="w-12 h-12 text-white" />
              </UButton>
            </UTooltip>
          </div>
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
  transition: transform 0.1s ease-out;
  transform-origin: bottom center;
}

.hover_underline_animation:hover::after {
  transform: scaleX(0.9);
  transform-origin: bottom center;
}
</style>
