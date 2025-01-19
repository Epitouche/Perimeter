import { handleError } from "~/utils/handleErrors";

/**
 * Fetches all services from the server.
 */
export default defineEventHandler(async () => {
  try {
    const response = await $fetch("http://server:8080/api/v1/service/info", {
      method: "GET",
    });
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
