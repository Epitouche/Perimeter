<script lang="ts" setup>
import type { ServiceResponse } from "~/interfaces/serviceResponse";

const tokenCookie = useCookie("token");
const errorMessage = ref<string | null>(null);
const menuOpen = ref(false);
const username = ref<string>("");
const infosConnection = ref<ServiceResponse | null>(null);

onMounted(() => {
  loadConnectionInfos();
});

async function loadConnectionInfos() {
  try {
    if (tokenCookie.value) {
      infosConnection.value = await servicesConnectionInfos(tokenCookie.value);
      username.value = infosConnection.value.user.username;
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
  ],
];
</script>

<template>
  <div class="flex items-center p-[1.5em] bg-custom_color-bg_section">
    <div class="flex items-center gap-[1em] grow-[0.95]">
      <img
        src="../public/PerimeterIcon.png"
        alt="perimeter-icon"
        class="h-[4em] w-[4em]"
      >
      <span class="font-black text-[2.5em]">Perimeter</span>
    </div>

    <div class="nav-links">
      <nav>
        <ul class="flex gap-[2.5em]">
          <li>
            <NuxtLink to="/myareas" class="nav-link">My Areas</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/workflow" class="nav-link">Workflow</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/myservices" class="nav-link">My Services</NuxtLink>
          </li>
        </ul>
      </nav>
    </div>

    <UDropdown :items="items" :popper="{ placement: 'bottom', arrow: true }">
      <UButton
        class="flex items-center justify-center bg-white h-14 w-14 shadow-lg rounded-full cursor-pointer"
        tabindex="0"
      >
        <Icon name="bytesize:user" class="text-black h-14 w-14" />
      </UButton>
      <div
        v-if="menuOpen"
        class="absolute top-full mt-4 right-0 p-4 rounded shadow-md flex flex-col gap-4 min-w-[200px] z-[1000] bg-custom_color-bg_section"
      >
        <div class="menu-header flex items-center justify-between gap-[1em]">
          <span class="font-[400] text-[1em]"> {{ username }}</span>
        </div>

        <NuxtLink to="/settings" class="nav-link" tabindex="0"
          >Settings</NuxtLink
        >
        <NuxtLink to="/login" tabindex="0" @keydown.enter="clearTokenAndLogout">
          <UButton
            class="flex items-center gap-2 py-2 px-4 text-base font-bold rounded-custom_border_radius cursor-pointer bg-custom_color-bg_section logout-button"
            tabindex="-1"
            @click="clearTokenAndLogout"
          >
            <Icon name="bytesize:sign-out" class="text-white h-5 w-5" />
            Logout
          </UButton>
        </NuxtLink>
      </div>
    </UDropdown>
  </div>
</template>

<style>
.nav-link {
  color: black;
  text-decoration: none;
  font-size: 1.5rem;
  font-weight: 700;
}

.nav-link:hover {
  text-decoration: underline;
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
