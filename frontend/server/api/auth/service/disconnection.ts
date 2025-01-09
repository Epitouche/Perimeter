export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.token) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token",
    });
  }

  try {
    const response = await $fetch(`http://server:8080/api/v1/token`, {
      method: "DELETE",
      body: {
        code: params.tokenId,
      },
      headers: {
        Authorization: params.authorization ? `${params.authorization}` : "",
      },
    });
    console.log("Deleting ? : ", response);
    return response;
  } catch (error) {
    console.log("Error :", error);
  }
});
