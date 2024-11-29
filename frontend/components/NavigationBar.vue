<template>
  <UHorizontalNavigation :links="links" class="border-b border-gray-900 dark:border-gray-900" />
  <div>
  </div>
</template>

<script lang="ts" setup>
  const links = [{
    label: 'Workflow',
    icon: 'i-heroicons-home',
    to: '/workflow'
  }, {
    label: 'My Area',
    icon: 'i-heroicons-home',
    to: '/myareas'
  }, {
    label: 'My Services',
    icon: 'i-heroicons-chart-bar',
    to: '/services'
  }]
</script>

<style>
</style>