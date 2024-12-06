export default defineEventHandler(async () => {
  try {
    const response = await $fetch(`http://localhost:8080/api/v1/service/info`, {
      method: 'GET',
    });
    return response;
  } catch (error: any) {
    console.error('Error fetching services:', error);
    throw createError({
      statusCode: error?.response?.status || 500,
      statusMessage: error?.response?.statusText || 'Failed to fetch services',
      data: error?.response?.data || { message: 'An unknown error occurred' },
    });
  }
});
