export default defineEventHandler(async (event) => {
  interface OAuthToken {
    token: string;
  }

  const params = await readBody(event);
  console.log(params)

  if (!params.code || !params.state || !params.service) {
    throw createError({
      statusCode: 400,
      message: 'Missing parameters: code, state, or service',
    });
  }

  const url = new URL(`http://server:8080/api/v1/${params.service}/auth/callback`);

  console.log(`Constructed URL: ${url}`);

  try {
    const response = await fetch(url.toString(), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        code: params.code,
        state: params.state,
      }),
    });
    console.log(response);
    //if (!response.ok) {
    //  throw createError({
    //    statusCode: response.status,
    //    message: `API Error: ${response.statusText}`,
    //  });
    //}

    const data: OAuthToken = await response.json();
    if (!data.token) {
      throw createError({
        statusCode: 400,
        message: 'Token unavailable',
      });
    }

    return data;

  } catch (error:any) {
    console.error('Error during API call:', error);
    throw createError({
      statusCode: 500,
      message: `Internal Server Error: ${error.message || error}`,
    });
  }
});
