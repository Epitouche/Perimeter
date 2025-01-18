<script setup lang="ts">
import type { LocationQueryValue } from "vue-router";
import type { Area } from "@/interfaces/areas";
import type { AreaResult } from "@/interfaces/areaResult";

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

const areaIdNumber = Number(areaId);
const valueNumber = value ? Number(value) : null;

if (areaIdNumber !== null && isNaN(areaIdNumber)) {
  console.error("Invalid areaId:", areaId);
}
if (valueNumber !== null && isNaN(valueNumber)) {
  console.error("Invalid value:", value);
}

const componentKey = ref(0);
const focusDiv = ref<HTMLElement | null>(null);
const selectedAreaData = ref<{ date: string; result: string }[] | null>(null);

const areaIsOpen = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.areas.map((area) => [area.id, false])),
);

const editAreaIsOpen = reactive<{ [key: number]: boolean }>(
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
  if (areaIsOpen[areaId]) {
    fetchAreaResult(areaId);
  }
};

const toggleEditArea = (areaId: number) => {
  console.log("Main Area Edit Toggle: ", editAreaIsOpen[areaId]); /////////////////:
  editAreaIsOpen[areaId] = !editAreaIsOpen[areaId];
  if (
    editAreaIsOpen[areaId] &&
    !state[areaId]?.title &&
    !state[areaId]?.description &&
    !state[areaId]?.action_refresh_rate
  ) {
    const area = props.areas.find((a) => a.id === areaId);
    if (area) {
      state[areaId] = {
        title: area.title,
        description: area.description,
        action_refresh_rate: area.action_refresh_rate,
      };
    }
  }
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
  return name
    .replace(/^action_/, "")
    .replace(/_/g, " ")
    .replace(/([a-z])([A-Z])/g, "$1 $2");
}

const fetchAreaResult = async (areaId: number) => {
  if (token.value) {
    try {
      errorMessage.value = null;
      const response = await $fetch<AreaResult[]>("/api/area/result", {
        method: "POST",
        body: {
          token: token.value,
          areaId: areaId,
        },
      });
      if (response && response.length > 0) {
        const combinedData = response.map((item) => ({
          date: formatDate(item.created_at),
          result: item.result,
        }));

        selectedAreaData.value = combinedData;
      } else {
        selectedAreaData.value = null;
      }
    } catch (error) {
      errorMessage.value = handleErrorStatus(error);
      console.error(errorMessage.value);
    }
  }
};

function formatDate(isoDate: string): string {
  const date = new Date(isoDate);
  const day = String(date.getDate()).padStart(2, "0");
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const year = date.getFullYear();
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  const seconds = String(date.getSeconds()).padStart(2, "0");

  return `${day}-${month}-${year} ${hours}:${minutes}:${seconds}`;
}

const updateAreaValue = async (
  areaId: number,
  typeName: string | null,
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

  if (typeName) {
    const targetOptionKey = `${typeName}_option` as keyof Area;

    if (!(targetOptionKey in updatedArea)) {
      (updatedArea[targetOptionKey] as { [key: string]: string | number }) = {};
    }

    (updatedArea[targetOptionKey] as { [key: string]: string | number })[
      keyString
    ] =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;
  } else {
    const targetOtherKey = `${keyString}` as keyof Area;
    (updatedArea[targetOtherKey] as string | number) =
      typeof value === "string" && !isNaN(Number(value))
        ? Number(value)
        : value;
  }

  console.log("Final updatedArea:", updatedArea);

  try {
    errorMessage.value = null;

    await $fetch("/api/area/update", {
      method: "POST",
      body: {
        token: token.value,
        area: updatedArea,
      },
    });

    emit("refreshAreas");
  } catch (error) {
    errorMessage.value = handleErrorStatus(error);
  }

  router.push("myareas");
  // toggleEditArea(areaId);
};

const state = reactive<
  Record<number, Pick<Area, "title" | "description" | "action_refresh_rate">>
>({});

const filteredState = (areaId: number) => {
  const areaState = state[areaId] || {};
  return Object.entries(areaState)
    .filter(([key]) =>
      ["title", "description", "action_refresh_rate"].includes(key),
    )
    .reduce(
      (obj, [key, value]) => {
        obj[key] = value;
        return obj;
      },
      {} as Record<string, string | number>,
    );
};

const isValidKey = (
  key: string,
): key is "title" | "description" | "action_refresh_rate" => {
  return (
    key === "title" || key === "description" || key === "action_refresh_rate"
  );
};

onMounted(() => {
  console.log("areas in AreaCardContainer", props.areas);

  props.areas.forEach((area) => {
    state[area.id] = {
      title: area.title,
      description: area.description,
      action_refresh_rate: area.action_refresh_rate,
    };
  });
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
        :ui="{ padding: '!px-0 !py-4', constrained: 'max-w-none' }"
        class="custom_card text-white"
        :style="{ backgroundColor: area.action.service.color }"
        tabindex="0"
        @click="toggleAreaModal(area.id)"
        @keydown.enter="toggleAreaModal(area.id)"
      >
        <h6 class="clamp-1-line overflow-hidden w-full text-center px-4">
          {{ formatName(area.title) }}
        </h6>

        <div
          class="flex flex-col justify-center items-center gap-0 w-full px-2"
        >
          <img
            :src="area.action.service.icon"
            :alt="area.action.service.name"
            class="w-[35%] h-[35%] mr-[38%] -mb-[14%]"
          >
          <img
            :src="area.reaction.service.icon"
            :alt="area.reaction.service.name"
            class="w-[35%] h-[35%] ml-[38%]"
          >
        </div>
      </UContainer>
      <UModal
        ref="focusDiv"
        v-model="areaIsOpen[area.id]"
        tabindex="0"
        :ui="{
          width:
            'min-w-[50%] max-w-[70%] max-lg:max-w-[75%] max-md:max-w-[85%] max-sm:max-w-[95%]',
        }"
      >
        <div
          class="flex flex-col gap-12 font-semibold text-white rounded-custom_border_radius pl-20 pr-16 py-10 max-lg:pl-10 max-lg:pr-10 max-lg:py-5 w-full max-h-[90vh]"
          :style="{ backgroundColor: area.action.service.color }"
        >
          <div>
            <div
              class="flex flex-row justify-between items-center w-full overflow-y-auto px-1"
            >
              <div class="flex flex-row items-center gap-3">
                <UToggle
                  :model-value="areaIsEnabled(area.id)"
                  size="xl"
                  tabindex="0"
                  @update:model-value="toggleAreaEnableSwitch(area.id)"
                  @keydown.enter="toggleAreaEnableSwitch(area.id)"
                />
                <div v-if="areaIsEnabled(area.id)" class="text-2xl">
                  <p>Enabled</p>
                </div>
                <div v-else class="text-2xl">
                  <p>Disabled</p>
                </div>
              </div>
              <UButton
                variant="ghost"
                class="self-end w-fit"
                tabindex="-1"
                @click="toggleAreaModal(area.id)"
                @keydown.enter="toggleAreaModal(area.id)"
              >
                <UIcon
                  name="i-bytesize-close"
                  class="w-[3.5vw] h-[3.5vh] text-white"
                />
              </UButton>
            </div>

            <h2 class="capitalize text-center w-full">
              {{ area.title }}
            </h2>
          </div>

          <div class="flex flex-col gap-10">
            <UpdateAreaOptions
              :area-id="area.id"
              type-name="action"
              :color="area.action.service.color"
              :type="area.action"
              :type-options="area.action_option"
              @update-area-value="updateAreaValue"
            />
            <UpdateAreaOptions
              :area-id="area.id"
              type-name="reaction"
              :color="area.action.service.color"
              :type="area.reaction"
              :type-options="area.reaction_option"
              @update-area-value="updateAreaValue"
            />
          </div>

          <div class="scrollbar-hidden w-full overflow-x-scroll max-h-[10vh]">
            <h5 class="self-start whitespace-nowrap">Description:</h5>
            <h6 class="pl-10 whitespace-nowrap">{{ area.description }}</h6>
          </div>

          <UContainer
            :ui="{ padding: '!px-0', constrained: 'max-w-none' }"
            class="scrollable-element w-full bg-custom_color-bg_section overflow-y-scroll min-h-[10vh] rounded-lg text-black"
          >
            <div>
              <h5
                v-if="!selectedAreaData || selectedAreaData.length === 0"
                class="px-1"
              >
                No Result
              </h5>
              <ul v-else>
                <li v-for="(item, index) in selectedAreaData" :key="index">
                  <span>{{ item.date }}</span> - <span>{{ item.result }}</span>
                </li>
              </ul>
            </div>
          </UContainer>

          <div class="flex flex-row justify-end items-center gap-2">
            <UTooltip text="Edit" class="self-end w-fit">
              <UButton
                variant="ghost"
                class="hover_underline_animation items-end w-fit p-0 pb-1"
                tabindex="0"
                @click="toggleEditArea(area.id)"
                @keydown.enter="toggleEditArea(area.id)"
              >
                <UIcon
                  name="i-bytesize-edit"
                  class="w-[3.5vw] h-[3.5vh] p-0 text-white"
                />
              </UButton>
            </UTooltip>

            <USlideover v-model="editAreaIsOpen[area.id]">
              <UForm
                :state="state[area.id]"
                class="flex flex-col justify-center items-center gap-5 py-10 bg-custom_color-bg_section"
              >
                <UFormGroup
                  v-for="(value, key) in filteredState(area.id)"
                  :key="key"
                  :label="formatName(key)"
                  :name="key"
                  :ui="{ label: { base: 'capitalize text-xl pl-3' } }"
                >
                  <div class="flex flex-row justify-center items-center gap-3">
                    <UInput
                      v-model="
                        state[area.id][
                          key as keyof Pick<Area, 'title' | 'description'>
                        ]
                      "
                      :ui="{
                        placeholder: '!px-5 !py-2 font-light',
                        size: { sm: 'text-lg' },
                      }"
                      :placeholder="key + '...'"
                    />
                    <UButton
                      tabindex="0"
                      @click="
                        isValidKey(key) &&
                        state[area.id][key] !==
                          props.areas.find((a) => a.id === area.id)?.[key] &&
                        updateAreaValue(area.id, null, key, state[area.id][key])
                      "
                    >
                      <UIcon name="i-bytesize-checkmark" />
                    </UButton>
                  </div>
                </UFormGroup>
              </UForm>
            </USlideover>

            <UTooltip text="Delete" class="self-end w-fit">
              <UButton
                variant="ghost"
                class="hover_underline_animation items-end w-fit p-0 pb-1"
                tabindex="0"
                @click="onDelete(area.id)"
              >
                <UIcon
                  name="i-bytesize-trash"
                  class="w-[4vw] h-[4vh] p-0 text-white"
                />
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
        <h2>Are you sure you want to delete this area?</h2>
        <p class="text-2xl">This action cannot be undone!</p>
        <div class="flex flex-row justify-end items-center gap-5 pt-5">
          <UButton
            class="bg-opacity-0 border-custom_border_width text-2xl font-semibold py-3 px-5"
            :style="{
              borderColor: area.action.service.color,
              color: area.action.service.color,
            }"
            tabindex="0"
            @click="cancelDeletion(area.id)"
            >Cancel</UButton
          >
          <UButton
            class="text-white text-2xl font-semibold py-3 px-5"
            :style="{ backgroundColor: area.action.service.color }"
            tabindex="0"
            @click="onDelete(area.id)"
            >Delete
          </UButton>
        </div>
      </UModal>
    </div>
  </UContainer>
</template>

<style scoped>
.clamp-1-line {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: normal;
  transition: all 1s ease-in-out;
}

/* .hover-expand-text:hover {
  -webkit-line-clamp: unset;
  line-clamp: unset;
  overflow: visible;
  white-space: normal;
} */

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

[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}

.scrollable-element {
  scrollbar-width: thick;
  scrollbar-color: black rgba(255, 255, 255, 0.2);
}

.scrollbar-hidden::-webkit-scrollbar {
  display: none;
}
</style>
