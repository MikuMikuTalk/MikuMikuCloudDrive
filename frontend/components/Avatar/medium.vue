<template>
  <img :src="props.url" alt="User"
    class="relative inline-block h-[58px] w-[58px] !rounded-full object-cover object-center border border-slate-800 p-0.5" />
</template>

<script lang="ts" setup>
const props = defineProps({
  url: {
    type: String,
    default: "https://avatars.githubusercontent.com/u/107172084?v=4"
  }
})
</script>

<style></style>