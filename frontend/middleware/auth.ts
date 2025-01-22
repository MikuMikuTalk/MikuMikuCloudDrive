export default defineNuxtRouteMiddleware((to, from) => {
  const notification = useNotification();
  if (import.meta.client) {
    const token = localStorage.getItem("token");

    if (token == null) {
      // 如果 token 不存在且目标页面不是登录页，重定向到登录页
      if (to.path !== "/login") {
        notification.showError("您还没有登录，即将跳转到登录页");
        setTimeout(() => {
          return navigateTo("/login");
        }, 2000);
      }
    } else {
      // 如果 token 存在且目标页面是登录页，显示提示并重定向到首页
      if (to.path === "/login") {
        // 显示 Toast 提示

        notification.showSuccess("您已经登录，正在跳转到主页...");
        // 延迟 2 秒后跳转到首页
        setTimeout(() => {
          return navigateTo("/");
        }, 2000);
      }
    }
  }
});
