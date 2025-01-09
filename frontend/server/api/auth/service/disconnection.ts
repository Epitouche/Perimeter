import { handleErrorStatus } from "~/utils/handleErrorStatus";

export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.authorization || !params.tokenId ) {
    throw createError({
      statusCode: 400,
      message: "Missing parameters: token or tokenId",
    });
  }

  try {
    console.log("Id is : ", params.tokenId, " and token : ", params.authorization);
    const response = await $fetch(`http://server:8080/api/v1/token`,
      {
        method: "DELETE",
        body: {
          id: params.tokenId,
        },
        headers: {
          Authorization: params.authorization ? `Bearer  ${params.authorization}` : "",
        },
      },
    );
    console.log("Deleting ? : ", response);
    return response;
  }  catch (error: unknown) {
    const errorMessage = handleErrorStatus(error);
    if (errorMessage === "An unknown error occurred") {
      console.error("An unknown error occurred", error);
    }
    throw new Error(errorMessage || "Failed to fetch service connection info.");
  }
});
