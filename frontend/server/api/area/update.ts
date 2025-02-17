import { handleError } from "~/utils/handleErrors";

/**
 * Sends an area with updated values to the backend
 */
export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (!params.token || !params.area) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters",
      });
    }
    const response = await $fetch(`http://server:8080/api/v1/area/`, {
      method: "PUT",
      headers: {
        Authorization: "Bearer " + params.token,
      },
      body: params.area,
    });
    return response;
  } catch (error: unknown) {
    console.log("error: ", error);
    console.error(error);
    handleError(error);
  }
});
