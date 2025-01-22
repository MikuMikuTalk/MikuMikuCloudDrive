import type { ApiResponse } from "~/types/api-response";
import type {
  GetUserInfoRequest,
  GetUserInfoResponse,
} from "~/types/user_info";
import { useNuxtApp } from "#app";
export const useUserInfo = () => {
  const userInfoStore = useUserStore();
  const getUserInfo = async (data: GetUserInfoRequest): Promise<boolean> => {
    const { $axios } = useNuxtApp();
    try {
      const response: ApiResponse<GetUserInfoResponse> = await $axios.get(
        "/user/info",
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      if (response.code != 200) {
        return false;
      }
      // 存储用户信息到store里面
      userInfoStore.updateUserInfo(response.data);
      return true;
    } catch (err) {
      console.error("获取用户信息时遇到错误:", err);
      return false;
    }
  };
  return {
    getUserInfo,
  };
};
