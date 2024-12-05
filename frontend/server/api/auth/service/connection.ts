export default defineEventHandler(async (event) => {
  const params = await readBody(event);
  if (!params.code || !params.service) {
    throw createError({
      statusCode: 400,
      message: 'Missing parameters: code, state, or service',
    });
  }

  const response = await $fetch(`http://server:8080/api/v1/${params.service}/auth/callback`, {
    method: 'POST',
    body: {
      code: params.code,
    },
  });
  return response;
});
