import type { AxiosInstance } from "axios";
import axios from "axios";

export const instanceOfAxios: AxiosInstance = axios.create({
  baseURL: "",
  timeout: 300000,
});

// 请求拦截器
instanceOfAxios.interceptors.request.use();
