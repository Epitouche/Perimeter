<script lang="ts" setup>
import type { Area } from "@/interfaces/areas";
import { handleErrorStatus } from "../utils/handleErrorStatus.js";

definePageMeta({
  middleware: "auth",
});

const token = useCookie("token");
const errorMessage = ref<string | null>(null);
const isLoading = ref(true);

const areas = ref<Area[]>([]);
const filteredAreas = ref<Area[]>([]);

const searchQuery = ref<string>("");

const dateSort = ref(false);

const fetchAreas = async () => {
  try {
    errorMessage.value = null;
    const result = await $fetch<Area[]>("/api/area/myareas", {
      method: "POST",
      body: {
        token: token.value,
      },
    });
    areas.value = result;
    filteredAreas.value = result;
    //console.log("filteredAreas: ", filteredAreas.value);
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);

    if (errorMessage.value === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    filteredAreas.value = [];
  } finally {
    isLoading.value = false;
  }
};

watch(searchQuery, (newQuery) => {
  const lowerQuery = newQuery.toLowerCase();
  filteredAreas.value = areas.value.filter(
    (area) =>
      area.action.service.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.service.name.toLowerCase().includes(lowerQuery) ||
      area.action.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.name.toLowerCase().includes(lowerQuery),
  );
});

watch(dateSort, (newSort) => {
  const sortFn = (a: Area, b: Area) =>
    newSort
      ? new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
      : new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime();
  filteredAreas.value.sort(sortFn);
});

const items = [
  [
    {
      label: "Date",
      slot: "date",
    },
  ],
];

onMounted(() => {
  fetchAreas();
});
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-5 w-full">
    <h1>My Areas</h1>
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div
      v-else
      class="flex flex-col justify-center items-center gap-10 w-[90%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
      tabindex="0"
    >
      <div class="flex flex-row justify-between items-center w-full px-5 pt-1">
        <div class="w-1/4">
        <SearchBar
          v-model:search-query="searchQuery"
          tabindex="0"
        />
      </div>
        <UDropdown
          :items="items"
          :popper="{ placement: 'bottom' }"
          tabindex="0"
        >
          <UIcon
            name="i-bytesize-filter"
            class="text-black w-10 h-10 p-0 pb-1"
          />
          <template #date="{ item }">
            <div class="flex flex-row justify-evenly items-center w-full">
              <h3>Latest</h3>
              <UTooltip
                :text="`Sort by ${item.label}`"
                :popper="{ placement: 'top' }"
                class="w-fit"
              >
                <UToggle v-model="dateSort" />
              </UTooltip>
              <h3>Oldest</h3>
            </div>
          </template>
        </UDropdown>
      </div>
      <div v-if="isLoading" class="text-xl font-semibold">
        <p>Loading...</p>
      </div>
      <div v-else-if="errorMessage">
        <p>Error: {{ errorMessage }}</p>
      </div>
      <div v-else-if="filteredAreas.length === 0" class="w-full">
        <p>No areas found, create some!</p>
      </div>
      <div v-else class="w-[95%] overflow-y-scroll max-h-[64vh]">
        <AreaCardContainer :areas="filteredAreas" @refresh-areas="fetchAreas" />
      </div>
    </div>
  </div>
</template>

<style scoped>
p {
  font-size: 2.5rem;
  line-height: 2.5rem;
  color: black;
  font-weight: 600;
  text-align: center;
  padding-top: 2rem;
  padding-bottom: 2rem;
}
</style>
