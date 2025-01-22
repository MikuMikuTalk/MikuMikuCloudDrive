import Swal from "sweetalert2";
export const useNotification = () => {
  const showSuccess = (message: string, timeout: number = 2000) => {
    Swal.fire({
      icon: "success",
      title: "Success",
      text: message,
      timer: timeout,
      timerProgressBar: true, // 显示进度条
    });
  };

  const showError = (message: string, timeout: number = 2000) => {
    Swal.fire({
      icon: "error",
      title: "Error",
      text: message,
      timer: timeout,
      timerProgressBar: true, // 显示进度条
    });
  };

  return {
    showSuccess,
    showError,
  };
};
