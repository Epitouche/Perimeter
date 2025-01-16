<script lang="ts" setup>
import type { ServiceResponse } from "~/interfaces/serviceResponse";

const tokenCookie = useCookie("token");
const errorMessage = ref<string | null>(null);
const username = ref<string>("");
const infosConnection = ref<ServiceResponse | null>(null);

onMounted(() => {
  loadConnectionInfos();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);
      if (infosConnection.value && infosConnection.value.user && infosConnection.value.user.username) {
        username.value = infosConnection.value.user.username;
      }
    }
  } catch (error: unknown) {
    errorMessage.value = handleErrorStatus(error);
    console.error("Error loading connections infos:", error);
  }
}

// function toggleMenu() {
//   menuOpen.value = !menuOpen.value;
// }

const clearTokenAndLogout = () => {
  const tokenCookie = useCookie("token");
  tokenCookie.value = null;
};

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
</script>

<template>
  <UContainer :ui="{ padding: '!p-3', constrained: 'min-w-screen' }"
    class="flex flex-row justify-between items-center bg-custom_color-bg_section">
    <div class="flex flex-row items-center gap-5">
      <img src="../public/PerimeterIcon.png" alt="perimeter-icon" class="nav_icon">
      <h5>Perimeter</h5>
    </div>

    <div class="flex flex-row justify-evenly items-center gap-5">
      <NuxtLink to="/myareas" class="nav-link">
        <h6>My Areas</h6>
      </NuxtLink>
      <NuxtLink to="/workflow" class="nav-link">
        <h6>Workflow</h6>
      </NuxtLink>
      <NuxtLink to="/myservices" class="nav-link">
        <h6>My Services</h6>
      </NuxtLink>

      <UDropdown :items="items" :popper="{ placement: 'bottom', arrow: true }">
        <div
          class="nav_profile_circle border-black border-custom_border_width rounded-full flex justify-center items-center"
          tabindex="0">
          <Icon name="bytesize:user" class="nav_profile text-black" />
        </div>
        <template #name>
          <p class="w-full self-center text-black">{{ username }}</p>
        </template>
        <template #settings>
          <NuxtLink to="/settings" class="nav-link" tabindex="0">Settings</NuxtLink>
        </template>
        <template #logout>
          <NuxtLink to="/login" tabindex="0" @keydown.enter="clearTokenAndLogout">
            <UButton
              class="flex items-center gap-2 py-2 px-4 text-base font-bold rounded-custom_border_radius cursor-pointer bg-custom_color-bg_section logout-button"
              tabindex="-1" @click="clearTokenAndLogout">
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
.nav_icon {
  width: fit-content;
  height: 6vh;
}

.nav_profile_circle {
  width: 5vh;
  height: 5vh;
}

.nav_profile {
  width: 90%;
  height: 90%;
}

.logout-button {
  background-color: #ef4444;
}

.logout-button:hover {
  background-color: #dc2626;
}

[tabindex="0"]:focus {
  outline: 2px solid #007bff;
  outline-offset: 2px;
}
</style>
