import type { AxiosInstance } from "axios";
import axios from "axios";

declare module "#app" {
  interface NuxtApp {
    $axios: AxiosInstance;
  }
}

export default defineNuxtPlugin((nuxtApp) => {
  const runtimeConfig = useRuntimeConfig();
  const axiosInstance: AxiosInstance = axios.create({
    baseURL: "http://192.168.3.28:8888",
    timeout: 300000,
  });
  axiosInstance.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    config.headers["Authorization"] = `Bearer ${token}`;
    return config;
  });

  // axiosInstance.interceptors.response.use(
  //   (response) => {
  //     if (response.status !== 200) {
  //       console.error("服务失败", response.status);
  //       return Promise.reject(new Error(`服务失败: ${response.statusText}`));
  //     }
  //     return response.data;
  //   },
  //   (err) => {
  //     console.log("服务错误", err);
  //     return Promise.reject(new Error(`服务错误: ${err.message}`));
  //   }
  // );

  return {
    provide: {
      axios: axiosInstance,
    },
  };
});
