import type { ApiResponse } from "~/types/api-response";
import { useNuxtApp } from "#app";
import type { RegisterRequest, RegisterResponse } from "~/types/register";

export const useAuthRegister = () => {
  const notification = useNotification();
  const register = async (data: RegisterRequest): Promise<boolean> => {
    const { $axios } = useNuxtApp(); //获取axios实例
    try {
      // 发送注册请求
      const response = await $axios.post("/user/register", data, {
        headers: {
          "Content-Type": "application/json",
        },
      });
      const result: ApiResponse<RegisterResponse> = response.data;
      if (result.code !== 200) {
        notification.showError("注册失败，请重试");
        return false;
      }
      notification.showSuccess("注册成功，即将跳转");
      return true;
    } catch (err) {
      console.error("注册时遇到错误:", err);
      notification.showError("注册失败");
      return false;
    }
  };
  return {
    register,
  };
};
