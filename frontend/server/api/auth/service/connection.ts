import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.code || !params.service) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: code, or service",
    });
  }

  try {
    const response = await $fetch(
      `http://server:8080/api/v1/${params.service}/auth/callback`,
      {
        method: "POST",
        body: {
          code: params.code,
        },
        headers: {
          Authorization: params.authorization ? `${params.authorization}` : "",
        },
      },
    );
    return response;
  } catch (error: unknown) {
    handleError(error);
  }
});
