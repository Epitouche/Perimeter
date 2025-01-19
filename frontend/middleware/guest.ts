/**
 * Middleware to redirect authenticated users to the dashboard
 */
export default defineNuxtRouteMiddleware(() => {
  const authToken = useCookie("token");

  if (authToken.value) {
    return navigateTo("/myareas");
  }
});
