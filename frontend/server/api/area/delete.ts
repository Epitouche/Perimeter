import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (!params.token || !params.areaId) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters",
      });
    }
    const response = await $fetch(`http://server:8080/api/v1/area/`, {
      method: "DELETE",
      headers: {
        Authorization: "Bearer " + params.token,
      },
      body: {
        id: params.areaId,
      },
    });
    return response;
  } catch (error: unknown) {
    console.log("error: ", error);
    console.error(error);
    handleError(error);
  }
});
