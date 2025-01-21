import Swal from "sweetalert2";
export const useNotification = () => {
  const showSuccess = (message: string) => {
    Swal.fire({
      icon: "success",
      title: "Success",
      text: message,
    });
  };

  const showError = (message: string) => {
    Swal.fire({
      icon: "error",
      title: "Error",
      text: message,
    });
  };

  return {
    showSuccess,
    showError,
  };
};
