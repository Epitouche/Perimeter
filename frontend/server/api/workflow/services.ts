import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async () => {
  try {
    const response = await $fetch("http://server:8080/api/v1/service/info", {
      method: "GET",
    });
    //console.log("Services fetched successfully", response);
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
