<template>
  <nav class="bg-white shadow-lg">
    <div class="max-w-6xl mx-auto px-4">
      <div class="flex justify-between">
        <div class="flex space-x-7">
          <div>
            <a href="#" class="flex items-center py-4 px-2">
              <span class="font-semibold text-gray-500 text-lg">{{ appName }}</span>
            </a>
          </div>
          <div class="hidden md:flex items-center space-x-1">
            <a href="#"
              class="py-4 px-2 text-gray-500 font-semibold hover:text-green-500 transition duration-300">Home</a>
            <a href="#"
              class="py-4 px-2 text-gray-500 font-semibold hover:text-green-500 transition duration-300">About</a>
            <a href="#"
              class="py-4 px-2 text-gray-500 font-semibold hover:text-green-500 transition duration-300">Services</a>
            <a href="#"
              class="py-4 px-2 text-gray-500 font-semibold hover:text-green-500 transition duration-300">Contact</a>
          </div>
        </div>
        <div class="hidden md:flex items-center space-x-3">
          <NuxtLink to="/login"
            class="py-2 px-2 font-medium text-gray-500 rounded hover:bg-green-500 hover:text-white transition duration-300">
            Log
            In</NuxtLink>
          <NuxtLink to="/register"
            class="py-2 px-2 font-medium text-white bg-green-500 rounded hover:bg-green-400 transition duration-300">
            Sign
            Up</NuxtLink>
        </div>
        <div class="md:hidden flex items-center">
          <button class="outline-none mobile-menu-button">
            <svg class="w-6 h-6 text-gray-500 hover:text-green-500" fill="none" stroke-linecap="round"
              stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor">
              <path d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
    <div class="hidden mobile-menu">
      <ul class="">
        <li class="active"><a href="#" class="block text-sm px-2 py-4 text-white bg-green-500 font-semibold">Home</a>
        </li>
        <li><a href="#" class="block text-sm px-2 py-4 hover:bg-green-500 transition duration-300">About</a></li>
        <li><a href="#" class="block text-sm px-2 py-4 hover:bg-green-500 transition duration-300">Services</a></li>
        <li><a href="#" class="block text-sm px-2 py-4 hover:bg-green-500 transition duration-300">Contact</a></li>
      </ul>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
const appConfig = useAppConfig()
const appName = appConfig.AppName

const isMobileMenuOpen = ref(false);

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
  const mobileMenu = document.querySelector('.mobile-menu');
  if (mobileMenu) {
    mobileMenu.classList.toggle('hidden');
  }
};

onMounted(() => {
  const mobileMenuButton = document.querySelector('.mobile-menu-button');
  if (mobileMenuButton) {
    mobileMenuButton.addEventListener('click', toggleMobileMenu);
  }
});
</script>

<style></style>