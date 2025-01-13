import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { GITHUB_MOBILE_SECRET, GITHUB_MOBILE_CLIENT_ID } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleGithubLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
  bearerToken: string = '',
) {
  const config: AuthConfiguration = {
    clientId: GITHUB_MOBILE_CLIENT_ID,
    clientSecret: GITHUB_MOBILE_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['user', 'repo', 'user:email'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://github.com/login/oauth/authorize',
      tokenEndpoint: 'https://github.com/login/oauth/access_token',
    },
  };

  try {
    const result = await authorize(config);
    let data = await handleCallback(
      `http://${ipAddress}:8080/api/v1/github/auth/callback/mobile`,
      result,
      bearerToken,
    );
    if (data.error) {
      console.error(data.error);
    } else {
      setToken(data.token);
      if (login) {
        navigation.navigate('AreaView');
      }
    }
  } catch (error) {
    if ((error as Error).message !== 'User cancelled flow') {
      console.error('Failed to log in', error);
      Alert.alert('Error', (error as Error).message);
    }
  }
}

export { HandleGithubLogin };
