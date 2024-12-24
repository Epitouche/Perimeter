export const handleTokenStatus = (statusCode?: number, message?: string) => {
  const token = useCookie("token");
  if (statusCode === 401) {
    token.value = null;
    navigateTo("/login");
    return "401 Unauthorized access.";
  } else {
    return `Error: ${message}`;
  }
};
