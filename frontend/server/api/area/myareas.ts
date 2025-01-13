import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (!params.token) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters",
      });
    }

    const response = await $fetch(`http://server:8080/api/v1/area`, {
      method: "GET",
      headers: {
        Authorization: "Bearer " + params.token,
      },
    });
    return response;
  } catch (error: unknown) {
    console.error(error);
    handleError(error);
  }
});
