import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';

async function SpotifyOauthCallback(codeSpotify: string, navigation: any) {
    const { ipAddress, token, setToken, codeVerifier } = useContext(AppContext);
    const response = await fetch(
        `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({ codeSpotify, code_verifier: codeVerifier }),
        },
    );
    console.log('response: ', response);
    const data = await response.json();
    if (data.error) {
        console.error(data.error);
        navigation.goBack();
    } else {
        setToken(data.accessToken);
        console.log('data: ', data);
        if (data.accessToken !== '') {
            navigation.navigate('AreaView');
        } else {
            console.error('Error: no token');
        }
    }
}

export { SpotifyOauthCallback }
