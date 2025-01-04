import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { GMAIL_MOBILE_CLIENT_ID, GMAIL_SECRET } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleGoogleLogin(setToken: any, navigation: any, ipAddress: string, login: boolean = false) {
  const config: AuthConfiguration = {
    clientId: GMAIL_MOBILE_CLIENT_ID,
    clientSecret: GMAIL_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['profile', 'email'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://accounts.google.com/o/oauth2/auth',
      tokenEndpoint: 'https://accounts.google.com/o/oauth2/token',
    },
  };

  try {
    const result = await authorize(config);
    // console.log('result', result);
    let data;
    if (login) {
      data = await handleCallback(`http://${ipAddress}:8080/api/v1/google/auth/callback/mobile`, result);
    } else {
      setToken(result.accessToken);
      // TODO: call route when loging in from myServices page (waiting for back to be done)
    }
    if (data.error) {
      console.error(data.error);
    } else {
      setToken(data.token);
      if (login)
        {navigation.navigate('AreaView');}
    }
  } catch (error) {
    if ((error as Error).message != 'User cancelled flow') {
      console.error('Failed to log in', error);
      Alert.alert('Error', (error as Error).message);
    }
  }
}

export { HandleGoogleLogin };
