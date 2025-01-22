import { defineStore } from "pinia";
import type { GetUserInfoResponse } from "~/types/user_info";

export const useUserStore = defineStore("user", () => {
  // 状态
  //用户信息状态
  const userInfo = ref<GetUserInfoResponse | null>(null);
  function updateUserInfo(data: GetUserInfoResponse) {
    userInfo.value = data;
    console.info("用户信息更新成功：", data);
  }
  return {
    userInfo,
    updateUserInfo,
  };
});
