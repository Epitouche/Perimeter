import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';
import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { GITHUB_SECRET, GITHUB_CLIENT_ID } from '@env';

async function HandleGithubLogin(setToken: any, navigation: any) {
    const config: AuthConfiguration = {
        clientId: GITHUB_CLIENT_ID,
        clientSecret: GITHUB_SECRET,
        redirectUrl: 'com.perimeter-epitech://oauthredirect',
        scopes: ['user', 'repo'],
        serviceConfiguration: {
            authorizationEndpoint: 'https://github.com/login/oauth/authorize',
            tokenEndpoint: 'https://github.com/login/oauth/access_token',
        },
    };

    try {
        const result = await authorize(config);
        console.log('result', result);
        setToken(result.accessToken);
        navigation.navigate('AreaView');
    } catch (error) {
        console.error('Failed to log in to GitHub', error);
    }
}

async function GithubOauthCallback(codeGithub: string, navigation: any) {
    const { ipAddress, token, setToken } = useContext(AppContext);
    const response = await fetch(
        `http://${ipAddress}:8080/api/v1/github/auth/callback/mobile`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({ codeGithub }),
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

export { GithubOauthCallback, HandleGithubLogin };