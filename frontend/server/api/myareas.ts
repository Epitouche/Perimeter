import type { ErrorResponse } from "~/interfaces/error";

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
    let statusCode = 500;
    let message = "An unknown error occurred";

    if (error instanceof Error) {
      if ("statusCode" in error) {
        statusCode = (error as ErrorResponse).statusCode || statusCode;
      }
      message = error.message;
    } else if (
      typeof error === "object" &&
      error !== null &&
      "statusCode" in error
    ) {
      statusCode = (error as ErrorResponse).statusCode || statusCode;
    }
    throw createError({
      statusCode,
      message,
    });
  }
});
