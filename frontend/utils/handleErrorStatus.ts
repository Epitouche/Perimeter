export const handleErrorStatus = (error: unknown): string => {
  const token = useCookie("token");

  if (typeof error === "object" && error !== null) {
    const statusCode = (error as { statusCode?: number }).statusCode;
    const message =
      (error as { message?: string }).message || "An error occurred";


    if (statusCode === 400) {
      return "Invalid Credentials"
    } else if (statusCode === 401) {
      token.value = null;
      navigateTo("/login");
      return "Unauthorized access.";
    } else if (statusCode === 409) {
      return "Email already exist."
    } else {
      return `Error: ${message}`;
    }
  }

  return "An unknown error occurred";
};
