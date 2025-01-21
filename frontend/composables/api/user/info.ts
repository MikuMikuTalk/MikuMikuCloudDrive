import type { ApiResponse } from "~/types/api-response";
import type {
  GetUserInfoRequest,
  GetUserInfoResponse,
} from "~/types/user_info";

import { useNuxtApp } from "#app";
const { $axios } = useNuxtApp();
export function useUserInfo(
  data: GetUserInfoRequest
): Promise<ApiResponse<GetUserInfoResponse>> {
  return $axios.post("/user/info", {
    data,
  });
}
