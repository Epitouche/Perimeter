<template>
  <div class="navbar">
    <div class="logo">
      <img src="../public/PerimeterIcon.png" alt="Logo" class="logo-image" />
      <span class="title">Perimeter</span>
    </div>

    <div class="nav-links">
      <nav>
        <ul class="nav-list">
          <li>
            <NuxtLink to="/workflow" class="nav-link">Workflow</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/myareas" class="nav-link">My Areas</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/services" class="nav-link">My Services</NuxtLink>
          </li>
        </ul>
      </nav>
    </div>

    <div class="profile">
      <div class="profile-icon"></div>
    </div>
  </div>
</template>

<style>
.navbar {
  display: flex;
  align-items: center;
  padding: 1.5rem;
  background-color: white;
}

.logo {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex-grow: 0.95;
}

.logo-image {
  height: 4rem;
  width: 4rem;
}

.title {
  font-size: 2rem;
  font-weight: 900;
}

.nav-list {
  display: flex;
  gap: 2.5rem;
}

.nav-link {
  color: black;
  text-decoration: none;
  font-size: 1.25rem;
  font-weight: 700;
}

.nav-link:hover {
  text-decoration: underline;
}

.profile {
  margin-left: auto;
}

.profile-icon {
  height: 4rem;
  width: 4rem;
  background-color: #d1d5db;
  border-radius: 50%;
  border: 2px solid black;
}

</style>
