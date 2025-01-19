<script lang="ts" setup>
import type { ServiceResponse } from "~/interfaces/serviceResponse";

const tokenCookie = useCookie("token");
const errorMessage = ref<string | null>(null);
const username = ref<string>("");
const infosConnection = ref<ServiceResponse | null>(null);

/**
 * Load service connection information
 */
async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);
      if (
        infosConnection.value &&
        infosConnection.value.user &&
        infosConnection.value.user.username
      ) {
        username.value = infosConnection.value.user.username;
      }
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

/**
 * Clear token and logout
 */
const clearTokenAndLogout = () => {
  const tokenCookie = useCookie("token");
  tokenCookie.value = null;
};

/**
 * Items for the dropdown
 */
const items = [
  [
    {
      label: "Name",
      slot: "name",
    },
    {
      label: "Settings",
      slot: "settings",
    },
    {
      label: "Logout",
      slot: "logout",
    },
  ],
];

/**
 * When the component is mounted, load the service connection information
 */
onMounted(() => {
  loadConnectionInfos();
});
</script>

<template>
  <UContainer
    :ui="{ padding: '!p-4 max-sm:!p-2', constrained: 'min-w-screen' }"
    class="flex flex-row justify-between items-center bg-custom_color-bg_section"
  >
    <div class="max-md:hidden flex flex-row items-center gap-5">
      <img
        src="../public/PerimeterIcon.png"
        alt="perimeter-icon"
        style="width: 18%; height: 18%"
      >
      <h5>Perimeter</h5>
    </div>

    <div
      class="flex flex-row justify-around max-md:justify-between text-center max-md:pl-2 items-center gap-8 max-md:gap-0 max-md:w-full"
    >
      <NuxtLink to="/myareas" class="hover:underline">
        <h6>My Areas</h6>
      </NuxtLink>
      <NuxtLink to="/workflow" class="hover:underline">
        <h6>Workflow</h6>
      </NuxtLink>
      <NuxtLink to="/myservices" class="hover:underline">
        <h6>My Services</h6>
      </NuxtLink>

      <UDropdown
        :items="items"
        :popper="{ placement: 'bottom', arrow: true }"
        :ui="{ item: { padding: '!p-4 max-sm:!p-2' } }"
      >
        <UAvatar
          icon="i-bytesize-user"
          :ui="{
            size: {
              sm: '!h-fit !w-fit !py-[20%] !px-0 max-lg:!px-5 max-sm:!py-[3%] max-sm:!px-3',
            },
            icon: {
              size: {
                sm: '!w-[3.3vw] !h-[3.3vh] max-sm:!w-[7vw] max-sm:!h-[7vh]',
              },
            },
          }"
        />

        <template #name>
          <p class="w-full text-black max-sm:pt-2">{{ username }}</p>
        </template>

        <template #settings>
          <NuxtLink
            to="/settings"
            class="w-full flex flex-row justify-start items-center gap-2 max-sm:gap-4"
            tabindex="0"
          >
            <UIcon
              name="i-bytesize-settings"
              class="text-black w-[2vw] h-[2vh] max-sm:w-[6vw] max-sm:h-[6vh]"
            />
            <p class="text-black">Settings</p>
          </NuxtLink>
        </template>

        <template #logout>
          <NuxtLink
            to="/login"
            class="w-full self-center"
            tabindex="0"
            @keydown.enter="clearTokenAndLogout"
          >
            <UButton
              class="flex items-center gap-2 py-2 px-4 w-full text-base font-bold rounded-custom_border_radius cursor-pointer self-center bg-custom_color-bg_section logout-button"
              tabindex="-1"
              @click="clearTokenAndLogout"
            >
              <Icon name="bytesize:sign-out" class="text-white h-5 w-5" />
              Logout
            </UButton>
          </NuxtLink>
        </template>
      </UDropdown>
    </div>
  </UContainer>
</template>

<style scoped>
.logout-button {
  background-color: #ff0000;
}

.logout-button:hover {
  background-color: #dc2626;
}

[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
