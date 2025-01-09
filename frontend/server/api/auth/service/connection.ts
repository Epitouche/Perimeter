import { handleErrorStatus } from "~/utils/handleErrorStatus";

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
    //console.log(response);
    return response;
  } catch (error: unknown) {
    const errorMessage = handleErrorStatus(error);
    if (errorMessage === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    throw new Error(errorMessage || "Failed to fetch service connection info.");
  }
});
