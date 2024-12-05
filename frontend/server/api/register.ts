export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const response = await $fetch('http://localhost:8080/api/v1/auth/register', {
    method: 'POST',
    body: {
      email: body.email,
      username: body.username,
      password: body.password,
    },
  });

  return response;
});