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
