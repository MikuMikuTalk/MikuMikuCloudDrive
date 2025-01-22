import type {
  AxiosInstance,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";
import axios from "axios";

declare module "#app" {
  interface NuxtApp {
    $axios: AxiosInstance;
  }
}
const handleRequestHeader = (config: InternalAxiosRequestConfig<any>) => {
  const token = localStorage.getItem("token");
  config.headers["Authorization"] = token;
  console.log("jwtToken获取成功: ", token);
};
const handleResponse = (response: AxiosResponse) => {};
const handleNetworkError = (errStatus: number): string => {
  // HTTP 错误码参考 https://blog.meowrain.cn/api/i/2025/01/22/bTJO3P1737530752562639385.avif
  let errMessage: string = "";
  if (errStatus) {
    switch (errStatus) {
      case 400:
        errMessage = "错误的请求";
        break;
      case 401:
        errMessage = "未授权，请重新登录";
        break;
      case 403:
        errMessage = "拒绝访问";
        break;
      case 404:
        errMessage = "请求错误，未找到该资源";
        break;
      case 405:
        errMessage = "请求方法未允许";
        break;
      case 408:
        errMessage = "请求超时";
        break;
      case 500:
        errMessage = "服务器错误";
        break;
      case 501:
        errMessage = "网络未实现";
        break;
      case 502:
        errMessage = "网络错误";
        break;
      case 503:
        errMessage = "服务不可用";
        break;
      case 504:
        errMessage = "网络超时";
        break;
      case 505:
        errMessage = "http版本不支持该请求";
        break;
      default:
        errMessage = `其他连接错误 --${errStatus}`;
    }
  } else {
    errMessage = "无法连接到服务器！";
  }
  return errMessage;
};
export default defineNuxtPlugin((nuxtApp) => {
  //获取运行时配置
  const app = useAppConfig();
  const axiosInstance: AxiosInstance = axios.create({
    baseURL: app.Api.baseUrl,
    timeout: app.Api.timeout,
  });
  axiosInstance.interceptors.request.use(
    (config: InternalAxiosRequestConfig<any>) => {
      handleRequestHeader(config);
      return config;
    }
  );
  axiosInstance.interceptors.response.use(
    (response: AxiosResponse) => {
      //对响应数据做处理
      return response.data;
    },
    (error: any) => {
      let errMessage: string = "";
      //处理响应错误
      if (error.response) {
        errMessage = handleNetworkError(error.response.status);
        console.error(errMessage);
      } else if (error.request) {
        errMessage = handleNetworkError(0); //无法连接到服务器
        // 请求已经发出，但是没有收到响应
      }
      return Promise.reject(new Error(errMessage));
    }
  );
  return {
    provide: {
      axios: axiosInstance,
    },
  };
});
