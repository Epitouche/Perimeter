<script setup lang="ts">
import type { ServiceInfo } from "@/interfaces/serviceinfo";
import type { Token, ServiceResponse } from "~/interfaces/serviceResponse";

const props = defineProps<{
  type: string;
  services: ServiceInfo[];
}>();

const tokenCookie = useCookie("token");
const errorMessage = ref<string | null>(null);
const serviceConnected = ref<string[]>([]);
const tokens = ref<Token[]>([]);
const infosConnection = ref<ServiceResponse | null>(null);

const hover = reactive<{ [key: number]: boolean }>(
  Object.fromEntries(props.services.map((service) => [service.id, false])),
);

const isLongText = (text: string): boolean => text.length > 8;

onMounted(() => {
  loadConnectionInfos();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);

      if (infosConnection.value) {
        tokens.value = infosConnection.value.tokens;

        serviceConnected.value = tokens.value.map(
          (token) => token.service.name,
        );
      }
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

const isServiceConnectedOrInvalid = (appName: string): boolean => {
  const matchingService = props.services.find(
    (service) => service.name.toLowerCase() === appName.toLowerCase(),
  );

  if (
    serviceConnected.value.includes(appName) ||
    (matchingService && matchingService.oauth === false)
  ) {
    return true;
  }
  return false;
};

function formatName(name: string): string {
  return name.replace(/([a-z])([A-Z])/g, "$1 $2");
}
</script>

<template>
  <UContainer
    :ui="{ padding: '!px-0', constrained: 'max-w-full' }"
    class="flex flex-row justify-evenly items-center gap-10 flex-wrap w-full"
  >
    <div v-for="service in services" :key="service.id">
      <NuxtLink
        v-if="isServiceConnectedOrInvalid(service.name)"
        :to="{
          name: `workflow-${type}-service`,
          params: { service: service.id },
        }"
      >
        <UContainer
          :ui="{ padding: '!px-0 !py-5', constrained: 'max-w-none' }"
          class="custom_card flex flex-col !gap-0 text-white overflow-hidden"
          :style="{ backgroundColor: service.color }"
          @mouseenter="hover[service.id] = true"
          @mouseleave="hover[service.id] = false"
        >
          <img
            v-if="!hover[service.id]"
            :src="service.icon"
            :alt="service.name"
            style="width: 45%"
          />
          <img
            v-else-if="hover[service.id] && !isLongText(service.name)"
            :src="service.icon"
            :alt="service.name"
            style="width: 45%"
          />
          <h5
            class="clamp-1-line capitalize text-center break-words w-full hover-expand-text"
          >
            {{ formatName(service.name) }}
          </h5>
        </UContainer>
      </NuxtLink>
      <UContainer
        v-else
        :ui="{ padding: '!px-0 !py-5', constrained: 'max-w-none' }"
        class="custom_card flex flex-col !gap-0 text-white overflow-hidden opacity-40 cursor-not-allowed"
        :style="{ backgroundColor: service.color }"
      >
        <img
          v-if="!hover[service.id]"
          :src="service.icon"
          :alt="service.name"
          style="width: 45%"
        />
        <img
          v-else-if="hover[service.id] && !isLongText(service.name)"
          :src="service.icon"
          :alt="service.name"
          style="width: 45%"
        />
        <h5
          class="clamp-1-line capitalize text-center break-words w-full hover-expand-text"
        >
          {{ formatName(service.name) }}
        </h5>
      </UContainer>
    </div>
  </UContainer>
</template>

<style scoped>
@font-face {
  font-family: "ellipsis-font";
  src: local("DejaVu Sans");
  unicode-range: U+2026;
  size-adjust: 0%;
}

.clamp-1-line {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
  word-break: break-word;
  white-space: normal;
  transition: all 1s ease-in-out;
}

.hover-expand-text:hover {
  -webkit-line-clamp: unset;
  line-clamp: unset;
  overflow: visible;
  white-space: normal;
}
</style>
