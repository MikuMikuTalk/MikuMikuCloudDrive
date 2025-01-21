import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", () => {
  //用户登录状态
  const isLoggedIn = ref<boolean>(false);

  // token 存储
  const jwtToken = ref<string>("");
  function saveJwtToken(token: string) {
    jwtToken.value = token;
  }
  function updateLoggedStatus(status: boolean) {
    isLoggedIn.value = status;
    localStorage.setItem("token", jwtToken.value);
    if (status === false) {
      localStorage.removeItem("token");
    }
  }
  return {
    isLoggedIn,
    jwtToken,
    saveJwtToken,
    updateLoggedStatus,
  };
});
