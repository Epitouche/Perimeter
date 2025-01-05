import { handleError } from "~/utils/handleErrors";

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
    //console.log(response);
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
