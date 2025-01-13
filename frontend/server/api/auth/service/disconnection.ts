import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.authorization || !params.tokenId) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token or tokenId",
    });
  }

  try {
    await $fetch(`http://server:8080/api/v1/token`, {
      method: "DELETE",
      body: {
        id: params.tokenId,
      },
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
