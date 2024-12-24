export const handleErrorStatus = (error: unknown): string => {
  const token = useCookie("token");

  if (typeof error === "object" && error !== null) {
    const statusCode = (error as { statusCode?: number }).statusCode;
    const message = (error as { message?: string }).message || "An error occurred";

    if (statusCode === 401) {
      token.value = null;
      navigateTo("/login");
      return "401 Unauthorized access.";
    } else {
      return `Error: ${message}`;
    }
  }

  return "An unknown error occurred";
};
