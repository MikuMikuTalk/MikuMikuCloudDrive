import type { LoginRequest, LoginResponse } from "~/types/login";
import type { ApiResponse } from "~/types/api-response";
import { useNuxtApp } from "#app";

export const useAuthLogin = () => {
  const authStore = useAuthStore(); // 获取 authStore
  const notification = useNotification(); // 获取 notification

  const login = async (data: LoginRequest): Promise<boolean> => {
    const { $axios } = useNuxtApp(); // 获取 $axios 实例
    try {
      // 发送登录请求
      const response = await $axios.post("/user/login", data, {
        headers: {
          "Content-Type": "application/json",
        },
      });
      const result: ApiResponse<LoginResponse> = response.data;

      // 检查响应数据
      if (result.data == null) {
        console.error("token为空");
        notification.showError("登录失败，请重试");
        return false;
      }

      // 保存 token 和更新登录状态
      authStore.saveJwtToken(result.data);
      authStore.updateLoggedStatus(true);
      notification.showSuccess("登录成功");
      return true;
    } catch (err) {
      console.error("登录时遇到错误:", err);
      notification.showError("登录失败，请检查用户名和密码");
      return false;
    }
  };

  return {
    login,
  };
};
