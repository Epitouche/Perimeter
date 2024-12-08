export default defineEventHandler(async (event) => {
  try {
    const params = await readBody(event);
    if (
      !params.token ||
      !params.actionOptions ||
      !params.actionId ||
      !params.reactionOptions ||
      !params.reactionId
    ) {
      throw createError({
        statusCode: 400,
        message: "Missing parameters",
      });
    }
    console.log(params.actionId);
    console.log(params.reactionId);
    console.log(params.actionOptions);
    console.log(params.reactionOptions);
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
      },
    });
    return response;
  } catch (error) {
    console.error("Error fetching services:", error);
  }
});
