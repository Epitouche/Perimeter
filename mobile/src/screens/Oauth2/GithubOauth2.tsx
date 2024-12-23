import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { GITHUB_SECRET, GITHUB_CLIENT_ID } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleGithubLogin(setToken: any, navigation: any, ipAddress: string, login: boolean = false) {
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
      // console.log('result', result);
      let data;
      if (login) {
        data = await handleCallback(`http://${ipAddress}:8080/api/v1/github/auth/callback/mobile`, result);
      } else {
        setToken(result.accessToken);
        // TODO: call route when loging in from myServices page (waiting for back to be done)
      }
      if (data.error) {
        console.error(data.error);
      } else {
        setToken(data.token);
        if (login)
          navigation.navigate('AreaView');
      }
    } catch (error) {
      if ((error as Error).message != 'User cancelled flow') {
        console.error('Failed to log in', error);
        Alert.alert('Error', (error as Error).message);
      }
    }
}

export { HandleGithubLogin };
