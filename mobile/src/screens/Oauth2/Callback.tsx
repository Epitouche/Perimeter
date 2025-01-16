/**
 * Handles the OAuth2 callback by sending the authentication result to the specified callback URL.
 *
 * @param {string} callbackUrl - The URL to which the authentication result should be sent.
 * @param {any} authResult - The authentication result to be sent to the callback URL.
 * @param {string} [bearer=''] - Optional bearer token for authorization.
 * @returns {Promise<any>} - A promise that resolves to the response data from the callback URL.
 */
export async function handleCallback(
  callbackUrl: string,
  authResult: any,
  bearer: string = '',
) {
  const result = await fetch(callbackUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      ...(bearer && { Authorization: `Bearer ${bearer}` }),
    },
    body: JSON.stringify(authResult),
  });
  const data = await result.json();
  console.log('data: ', data);
  return data;
}
