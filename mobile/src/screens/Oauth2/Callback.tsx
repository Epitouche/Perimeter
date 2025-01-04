export async function handleCallback(callbackUrl: string, authResult: any) {
    const result = await fetch(
        callbackUrl,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(authResult),
        }
    );
    const data = await result.json();
    console.log('data: ', data);
    return data;
}
