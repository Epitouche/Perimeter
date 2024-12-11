export default defineEventHandler(async (event) => {
    try {
      const params = await readBody(event);
      if (!params.token || !params.serviceId) {
        throw createError({
          statusCode: 400,
          message: "Missing parameters: code, state, or service",
        });
      }
      const response = await $fetch(
        `http://server:8080/api/v1/service/info/${params.serviceId}`,
        {
          method: "GET",
          headers: {
            Authorization: "Bearer " + params.token,
          },
        },
      );
      return response;
    } catch (error) {
      console.error("Error fetching services:", error);
      console.log("Error fetching services:", error);
    }
  });
  