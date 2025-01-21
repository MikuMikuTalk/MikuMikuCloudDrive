<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <h2 class="text-2xl font-bold mb-6 text-center">Register</h2>
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
            Username
          </label>
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none  focus:ring-blue-500 focus:border-blue-500"
            id="username" type="text" placeholder="Username" v-model="registerRequest.username" />
        </div>
        <!-- <div class="mb-4">
          <label for="email" class="block text-gray-700 text-sm font-bold mb-2">
            Email
          </label>
          <input type="email"
            class="shadow border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none  focus:ring-blue-500 focus:border-blue-500"
            placeholder="Email">
        </div> -->
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
            Password
          </label>
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            id="password" type="password" placeholder="Password" v-model="registerRequest.password" />
        </div>
        <div class="flex items-center justify-center">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none"
            type="submit">
            Sign Up
          </button>
        </div>
        <div class="w-full flex justify-end">
          <NuxtLink class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800 "
            to="/login">
            Login
          </NuxtLink>
        </div>

      </form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useAuthRegister } from '~/composables/api/auth/register';

import type { RegisterRequest } from '~/types/register';
definePageMeta({
  layout: "auth"
})
const registerRequest = ref<RegisterRequest>({
  username: "",
  password: ""
})
const { register } = useAuthRegister()
const { navigateToLoginPage } = useNavigation()
const handleLogin = async () => {
  // Handle register logic here
  console.log("register with", registerRequest.value.username);

  const result = await register(registerRequest.value)

  if (result) {
    setTimeout(() => {
      navigateToLoginPage()
    }, 3000);
  }
  console.log('Logging in with:', registerRequest.value.username, registerRequest.value.password);

};
</script>

<style></style>
