/**
 * Redirects the user to the OAuth provider's authentication page.
 */
export default defineEventHandler(async (event) => {
  interface OAuthLink {
    authentication_url: string;
  }
  const body = await readBody(event);

  const data = await $fetch<OAuthLink>(body.link);

  if (!data.authentication_url) {
    throw new Error("Invalid authentication_url: Expected a valid URL");
  }

  return data;
});
