export default defineEventHandler(async (event) => {
  interface OAuthToken {
    token: string;
  }

  const body = await readBody(event);
  console.log(`Link is : ${body.link}`)

  try {
    console.log(`Link is : ${body.link}`)
    const response = await fetch(body.link);
    if (!response.ok) {
      throw new Error(`API Error: ${response.statusText}`);
    }
    
    const data: OAuthToken = await response.json();
    
    if (!data.token) {
      throw new Error('Token unavailable');
    }
    return data;
  } catch(error) {
    throw new Error(`${error}`)
  }

});
