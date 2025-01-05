import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    const response = await $fetch("http://server:8080/api/v1/service/info", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + params.token,
      },
    });
    //console.log("Services fetched successfully", response);
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
