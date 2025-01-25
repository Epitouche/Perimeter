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

/**
 * @description Fetches the areas from the server.
 */
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

/**
 * @description Watches the search query and filters the areas based on the query.
 *
 * @param {string} newQuery - The new search query.
 */
watch(searchQuery, (newQuery) => {
  const lowerQuery = newQuery.toLowerCase();
  filteredAreas.value = areas.value.filter(
    (area) =>
      area.action.service.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.service.name.toLowerCase().includes(lowerQuery) ||
      area.action.name.toLowerCase().includes(lowerQuery) ||
      area.reaction.name.toLowerCase().includes(lowerQuery)
  );
});

/**
 * @description Watches the date sort and sorts the areas based on the date.
 *
 * @param {boolean} newSort - The new date sort.
 */
watch(dateSort, (newSort) => {
  const sortFn = (a: Area, b: Area) =>
    newSort
      ? new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
      : new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime();
  filteredAreas.value.sort(sortFn);
});

/**
 * @description The items for the dropdown.
 */
const items = [
  [
    {
      label: "Date",
      slot: "date",
    },
  ],
];

/**
 * @description The items for the dropdown.
 */
onMounted(() => {
  fetchAreas();
});
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-5 w-full">
    <h1>My Areas</h1>
    <UAlert v-if="errorMessage" color="red" variant="solid" title="ERROR":description="errorMessage" class="justify-center items-center gap-5 w-[15%]"/>
    <div
      v-else
      class="flex flex-col justify-center items-center gap-10 w-[90%] h-full p-10 rounded-custom_border_radius bg-custom_color-bg_section"
      tabindex="0"
    >
      <div class="flex flex-row justify-between items-center w-full">
        <div class="min-w-[25%] max-w-[90%] px-5">
          <SearchBar v-model:search-query="searchQuery" tabindex="0" />
        </div>
        <UDropdown
          :items="items"
          :popper="{ placement: 'bottom' }"
          tabindex="0"
        >
          <UIcon name="i-bytesize-filter" class="text-black w-[4vw] h-[4vh]" />
          <template #date="{ item }">
            <div class="flex flex-row justify-evenly items-center w-full h-fit">
              <p style="font-size: 100%">Latest</p>
              <UTooltip
                :text="`Sort by ${item.label}`"
                :popper="{ placement: 'top' }"
                class=""
              >
                <UToggle v-model="dateSort" />
              </UTooltip>
              <p style="font-size: 100%">Oldest</p>
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
      <div
        v-else
        class="w-[95%] overflow-y-scroll max-h-[64vh] max-lg:max-h-[55vh]"
      >
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
