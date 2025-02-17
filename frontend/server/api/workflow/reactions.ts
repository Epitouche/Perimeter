import { handleError } from "~/utils/handleErrors";

/**
 * Fetches a services reactions from the server
 */
export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (!params.token || !params.service) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters: code, state, or service",
      });
    }
    const response = await $fetch(
      `http://server:8080/api/v1/reaction/info/${params.service}`,
      {
        method: "GET",
        headers: {
          Authorization: "Bearer " + params.token,
        },
      }
    );
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
