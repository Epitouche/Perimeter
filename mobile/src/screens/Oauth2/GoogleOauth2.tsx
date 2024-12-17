import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';
import {AuthConfiguration, authorize} from 'react-native-app-auth';
import { GMAIL_CLIENT_ID } from '@env';

async function HandleGoogleLogin(setToken: any, navigation: any) {
    const config: AuthConfiguration =
    {
        clientId: GMAIL_CLIENT_ID,
        redirectUrl: 'com.perimeter-epitech://oauthredirect',
        scopes: ['profile', 'email'],
        serviceConfiguration: {
            authorizationEndpoint: 'https://accounts.google.com/o/oauth2/auth',
            tokenEndpoint: 'https://accounts.google.com/o/oauth2/token',
        },
    };

    try {
        const result = await authorize(config);
        console.log('result', result);
        setToken(result.accessToken);
        navigation.navigate('AreaView');
    } catch (error) {
        console.error('Failed to log in to Google', error);
    }
}

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

export { SpotifyOauthCallback, HandleGoogleLogin };