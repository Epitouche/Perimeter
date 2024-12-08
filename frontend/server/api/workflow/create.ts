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
    let actionString: string = "{}";
    let reactionString: string = "{}";

    if (typeof params.actionOptions === "object" && params.actionOptions !== null) {
      const newTimeOptions = {
        hour: params.actionOptions.hour ? Number(params.actionOptions.hour) : 0,
        minute: params.actionOptions.minute ? Number(params.actionOptions.minute) : 0
      };

      actionString = JSON.stringify(newTimeOptions);
    }
    if (typeof params.reactionOptions === "object" && params.reactionOptions !== null) {
      reactionString = `{ ` + Object.entries(params.reactionOptions)
        .map(([key, value]) => {
          if (typeof value === "string") {
            return `"${key}": "${value}"`;
          }
          return `"${key}": ${value}`;
        })
        .join(', ') + ` }`;
    }
    console.log(params.reactionOptions)
    console.log("Formatted time without braces:", actionString);
    console.log("reactionOptions.value:", reactionString);

    const body = {
      action_option: actionString,
      action_id: Number(params.actionId),
      reaction_option: reactionString,
      reaction_id: Number(params.reactionId),
    };
    console.log("Body is : ", body)

    const response = await $fetch(`http://server:8080/api/v1/area`, {
      method: "POST",
      headers: {
        Authorization: "Bearer " + params.token,
      },
      body,
    });
    return response;
  } catch (error) {
    console.error("Error fetching services:", error);
    console.log("Error fetching services:", error);
  }
});
