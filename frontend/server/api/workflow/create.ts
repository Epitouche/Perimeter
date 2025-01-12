import { handleError } from "~/utils/handleErrors";

export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (
      !params.token ||
      !params.actionOptions ||
      !params.actionId ||
      !params.reactionOptions ||
      !params.reactionId ||
      !params.title ||
      !params.description
    ) {
      console.log("params.token", params.token);
      console.log("params.actionOptions", params.actionOptions);
      console.log("params.actionId", params.actionId);
      console.log("params.reactionOptions", params.reactionOptions);
      console.log("params.reactionId", params.reactionId);
      console.log("params.title", params.title);
      console.log("params.descritpion", params.description);
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
        action_option: params.actionOptions,
        action_id: Number(params.actionId),
        reaction_option: params.reactionOptions,
        reaction_id: Number(params.reactionId),
        title: params.title,
        description: params.description,
      },
    });
    return response;
  } catch (error: unknown) {
    console.log("error", error);
    handleError(error);
  }
});
