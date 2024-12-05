export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const response = await $fetch('http://server:8080/api/v1/auth/login', {
    method: 'POST',
    body: {
      username: body.username,
      password: body.password,
    },
  });

  return response;
});