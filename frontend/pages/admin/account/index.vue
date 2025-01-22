<template>
  <div class="flex flex-col items-center justify-center p-5 text-center min-h-screen bg-black text-white">
    <!-- 大头像 -->
    <div class="mb-5 relative inline-block avatar-container">
      <img
        class="w-[100px] h-[100px] sm:w-[120px] sm:h-[120px] md:w-[150px] md:h-[150px] rounded-full border-4 border-white shadow-[0_4px_8px_rgba(255,255,255,0.2)] relative z-10 transition-transform duration-300 ease-in-out"
        src="https://avatars.githubusercontent.com/u/107172084?v=4" alt="My Avatar">
    </div>

    <!-- 昵称 -->
    <div class="text-[24px] sm:text-[26px] md:text-[28px] font-bold text-white mb-2.5">MeowRain</div>

    <!-- 个人社交平台 -->
    <div class="flex gap-3 sm:gap-4 md:gap-5 mb-5">
      <a href="https://github.com/meowrain" target="_blank"
        class="flex items-center gap-1 sm:gap-1.25 text-white text-sm sm:text-base no-underline transition-colors duration-200 ease-in-out hover:text-blue-500">
        <img class="w-5 h-5 sm:w-6 sm:h-6 invert" src="https://cdn-icons-png.flaticon.com/512/25/25231.png"
          alt="GitHub">
        GitHub
      </a>
      <a href="https://twitter.com/meowrain" target="_blank"
        class="flex items-center gap-1 sm:gap-1.25 text-white text-sm sm:text-base no-underline transition-colors duration-200 ease-in-out hover:text-blue-500">
        <img class="w-5 h-5 sm:w-6 sm:h-6 invert" src="https://cdn-icons-png.flaticon.com/512/25/25347.png"
          alt="Twitter">
        Twitter
      </a>
      <a href="https://weibo.com/meowrain" target="_blank"
        class="flex items-center gap-1 sm:gap-1.25 text-white text-sm sm:text-base no-underline transition-colors duration-200 ease-in-out hover:text-blue-500">
        <img class="w-5 h-5 sm:w-6 sm:h-6 invert" src="https://cdn-icons-png.flaticon.com/512/23/23931.png" alt="Weibo">
        Weibo
      </a>
    </div>

    <!-- 所在学校 -->
    <div class="text-base sm:text-lg text-gray-300 mb-5">山西农业大学 - 计算机科学与技术专业</div>

    <!-- 经历 -->
    <div
      class="bg-gray-900 p-4 sm:p-5 rounded-lg shadow-[0_4px_8px_rgba(255,255,255,0.1)] w-full max-w-[90%] sm:max-w-[400px] mb-5">
      <h3 class="text-[20px] sm:text-[22px] text-white mb-2.5">经历</h3>
      <ul class="list-disc pl-5 text-left">
        <li class="text-sm sm:text-base text-gray-300 mb-2">2021年 - 加入某实验室，参与前端开发项目</li>
        <li class="text-sm sm:text-base text-gray-300 mb-2">2022年 - 获得校级优秀学生奖</li>
        <li class="text-sm sm:text-base text-gray-300 mb-2">2023年 - 实习于某互联网公司，负责前端开发</li>
      </ul>
    </div>

    <!-- 欢迎语 -->
    <div class="text-xl sm:text-2xl font-bold text-white mt-5 mb-10">欢迎来到我的个人页面！</div>
  </div>
</template>

<script lang="ts" setup>
definePageMeta({
  layout: "admin"
})
</script>

<style scoped>
.avatar-container::before {
  content: '';
  position: absolute;
  top: -8px;
  left: -8px;
  right: -8px;
  bottom: -8px;
  border-radius: 50%;
  background: conic-gradient(from 0deg,
      #ff0000,
      #ff7f00,
      #ffff00,
      #00ff00,
      #0000ff,
      #4b0082,
      #9400d3,
      #ff0000);
  z-index: 0;
  opacity: 0;
  transition: opacity 0.3s ease, transform 0.3s ease;
  transform: scale(0.9);
}

.avatar-container:hover::before {
  opacity: 1;
  transform: scale(1);
  animation: rotateRainbow 1.5s linear infinite;
}

.avatar-container:hover .avatar {
  transform: scale(1.1);
}

@keyframes rotateRainbow {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
</style>