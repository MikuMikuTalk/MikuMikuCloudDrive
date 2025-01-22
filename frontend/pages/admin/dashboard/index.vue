<template>
  <div class="min-h-screen bg-gray-100">
    <!-- 顶部导航栏 -->
    <div class="bg-white shadow-md p-4">
      <div class="flex justify-between items-center">
        <div class="text-xl font-bold">Dashboard</div>
        <div class="flex items-center space-x-4">
          <button class="bg-gray-200 p-2 rounded-full hover:bg-gray-300">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
            </svg>
          </button>
          <div class="relative">
            <AvatarWithDropDown />
          </div>
        </div>
      </div>
    </div>

    <!-- 内容区域 -->
    <div class="p-4">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div class="bg-white p-6 rounded-lg shadow-md">
          <h3 class="text-lg font-semibold">Total Users</h3>
          <p class="text-2xl font-bold">1,234</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md">
          <h3 class="text-lg font-semibold">Total Revenue</h3>
          <p class="text-2xl font-bold">$12,345</p>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md">
          <h3 class="text-lg font-semibold">Total Orders</h3>
          <p class="text-2xl font-bold">567</p>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts" setup>
definePageMeta({
  layout: "admin"
})
</script>

<style></style>