import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (
      !params.token ||
      !params.actionId ||
      !params.actionOptions ||
      !params.refreshRate ||
      !params.description ||
      !params.reactionId ||
      !params.reactionOptions ||
      !params.title
    ) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters",
      });
    }
    const response = await $fetch(`http://server:8080/api/v1/area`, {
      method: "POST",
      headers: {
        Authorization: "Bearer " + params.token,
      },
      body: {
        action_id: Number(params.actionId),
        action_option: params.actionOptions,
        action_refresh_rate: Number(params.refreshRate),
        description: params.description,
        reaction_id: Number(params.reactionId),
        reaction_option: params.reactionOptions,
        title: params.title,
      },
    });
    return response;
  } catch (error: unknown) {
    console.log("error", error);
    handleError(error);
  }
});
