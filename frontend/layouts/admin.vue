<template>
  <div class="flex min-h-screen">
    <!-- 侧边栏 -->
    <SideBar />

    <!-- 内容区域 -->
    <main class="flex-1 p-6 bg-gray-100">
      <slot />
    </main>
  </div>
</template>

<script lang="ts" setup>

</script>