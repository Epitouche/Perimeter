export default defineEventHandler(async (event) => {
  interface OAuthLink {
    authentication_url: string;
  }

  const body = await readBody(event);

  const response = await fetch(body.label);
  if (!response.ok) {
    throw new Error(`API Error: ${response.statusText}`);
  }

  const data: OAuthLink = await response.json();

  if (!data.authentication_url || !isValidUrl(data.authentication_url)) {
    throw new Error('Invalid authentication_url: Expected a valid URL');
  }

  return data;
});

const isValidUrl = (url: string) => {
  try {
    new URL(url);
    return true;
  } catch (err) {
    return false;
  }
};
