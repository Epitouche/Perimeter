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
    const response = await $fetch(`http://server:8080/api/v1/area`, {
      method: "POST",
      headers: {
        Authorization: "Bearer " + params.token,
      },
      body: {
        actionOptions: params.actionOptions,
        actionId: params.actionId,
        reactionOptions: params.reactionOptions,
        reactionId: params.reactionId,
      },
    });
    return response;
  } catch (error) {
    console.error("Error fetching services:", error);
    console.log("Error fetching services:", error);
  }
});
