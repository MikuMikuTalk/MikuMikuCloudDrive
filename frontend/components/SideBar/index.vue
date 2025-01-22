<template>
  <!-- Sidebar -->
  <aside class="w-80 bg-gray-800 text-gray-200 flex flex-col">
    <!-- Logo -->
    <div class="p-4 text-center font-bold text-lg bg-gray-900">
      MikuMikuCloudDrive Admin
    </div>
    <!-- Navigation -->
    <nav class="flex-1 p-4">
      <ul class="space-y-4">
        <li>
          <NuxtLink to="/admin/dashboard" class="flex items-center gap-3 p-3 rounded-lg hover:bg-gray-700 transition">
            <IconDashboard />
            <span>Dashboard</span>
          </NuxtLink>
        </li>
        <li>
          <NuxtLink to="/admin/settings" class="flex items-center gap-3 p-3 rounded-lg hover:bg-gray-700 transition">
            <IconSettings />
            <span>Settings</span>
          </NuxtLink>
        </li>
        <li>
          <NuxtLink to="/admin/account" class="flex items-center gap-3 p-3 rounded-lg hover:bg-gray-700 transition">
            <IconAccount />
            <span>Account</span>
          </NuxtLink>
        </li>
        <li>
          <NuxtLink to="/" class="flex items-center gap-3 p-3 rounded-lg hover:bg-gray-700 transition">
            <IconHouse />
            <span>Main Page</span>
          </NuxtLink>
        </li>
      </ul>
    </nav>
    <!-- Footer -->
    <div class="p-4 text-center text-sm bg-gray-900">
      © 2025 MikuMikuCloudDrive
    </div>
  </aside>
</template>

<script lang="ts" setup>
const isCollapsed = ref<boolean>(false);

function toggleSidebar() {
  isCollapsed.value = !isCollapsed.value;
}
</script>

<style scoped>
.sidebar {
  height: 100vh;
  background-color: #2c3e50;
  transition: width 0.3s ease;
}
</style>