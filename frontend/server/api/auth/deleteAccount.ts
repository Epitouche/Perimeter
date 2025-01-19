import { handleError } from "~/utils/handleErrors";

/**
 * Sends a DELETE request to the server to delete the users account.
 */
export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.authorization) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token",
    });
  }

  try {
    await $fetch(`http://server:8080/api/v1/user/info`, {
      method: "DELETE",
      headers: {
        Authorization: params.authorization
          ? `Bearer ${params.authorization}`
          : "",
      },
    });
    return true;
  } catch (error: unknown) {
    handleError(error);
  }
});
