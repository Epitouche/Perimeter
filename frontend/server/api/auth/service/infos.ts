import { handleError } from "~/utils/handleErrors";

/**
 * Fetches a users information from the server.
 */
export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  try {
    const response = await $fetch(`http://server:8080/api/v1/user/info/all`, {
      method: "GET",
      headers: {
        Authorization: params.authorization
          ? `Bearer ${params.authorization}`
          : "",
      },
    });
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
