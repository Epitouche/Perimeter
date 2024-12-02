export default defineNuxtRouteMiddleware(() => {
    const authToken = localStorage.getItem('authToken') !== null;

    if (!authToken) {
        return navigateTo('/login');
    }
})