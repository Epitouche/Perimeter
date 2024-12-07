<script lang="ts" setup>
definePageMeta({
  middleware: "auth",
});

const menuOpen = ref(false);

function toggleMenu() {
  menuOpen.value = !menuOpen.value;
}

const clearTokenAndLogout = () => {
  const tokenCookie = useCookie("token");
  tokenCookie.value = null;
};
</script>

<template>
  <div class="flex items-center p-[1.5em] bg-custom_color-bg_section">
    <div class="flex items-center gap-[1em] grow-[0.95]">
      <img
        src="../public/PerimeterIcon.png"
        alt="perimeter-icon"
        class="h-[4em] w-[4em]"
      />
      <span class="font-black text-[2em]">Perimeter</span>
    </div>

    <div class="nav-links">
      <nav>
        <ul class="flex gap-[2.5em]">
          <li>
            <NuxtLink to="/workflow" class="nav-link">Workflow</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/myareas" class="nav-link">My Areas</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/myservices" class="nav-link">My Services</NuxtLink>
          </li>
        </ul>
      </nav>
    </div>

    <div class="ml-auto relative">
      <button
        class="h-[4em] w-[4em] bg-gray-300 rounded-full border-custom_border_width border-black cursor-pointer"
        @click="toggleMenu"
      ></button>
      <div
        v-if="menuOpen"
        class="absolute top-full mt-4 right-0 p-4 rounded shadow-md flex flex-col gap-4 min-w-[200px] z-[1000] bg-custom_color-bg_section"
      >
        <div class="menu-header flex items-center justify-between gap-[1em]">
          <span class="font-[400] text-[1em]">Username</span>
          <div
            class="h-[3em] w-[3em] bg-gray-300 rounded-full border-custom_border_width border-black"
          ></div>
        </div>

        <NuxtLink to="/settings" class="nav-link">Settings</NuxtLink>

        <UButton
          class="flex items-center gap-2 py-2 px-4 text-base font-bold rounded-custom_border_radius cursor-pointer bg-custom_color-bg_section logout-button"
          @click="clearTokenAndLogout"
        >
          <svg
            class="w-[1em] h-[1em]"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 32 32"
          >
            <path
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M28 16H8m12-8l8 8l-8 8m-9 4H3V4h8"
            />
          </svg>
          <NuxtLink to="/login">Logout</NuxtLink>
        </UButton>
      </div>
    </div>
  </div>
</template>

<style>
.nav-link {
  color: black;
  text-decoration: none;
  font-size: 1.25rem;
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
</style>
