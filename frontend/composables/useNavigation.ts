// composables/useNavigation.js
import { useRouter } from "vue-router";

export const useNavigation = () => {
  const router = useRouter();

  const navigateToHome = () => {
    router.push("/");
  };

  const navigateToLoginPage = () => {
    router.push("/login");
  };
  return { navigateToHome, navigateToLoginPage };
};
