<template>
  <div>

    <div class="bg-luotianyi bg-cover bg-center" :style="{ height: `calc(100vh - 60px)` }">
      <div class="relative z-10 flex items-center justify-center h-full text-white">
        <h1 class="text-4xl font-bold">Welcome to MikuMikuCloudDrive</h1>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
definePageMeta({
  middleware: ["auth"]
})
</script>

<style></style>