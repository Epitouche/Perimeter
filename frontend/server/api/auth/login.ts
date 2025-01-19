/**
 * Server-side handler for the login endpoint.
 */
export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const response = await $fetch("http://server:8080/api/v1/user/login", {
    method: "POST",
    body: {
      username: body.username,
      password: body.password,
    },
  });

  return response;
});
