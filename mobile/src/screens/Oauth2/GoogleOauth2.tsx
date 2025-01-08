import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';
import { GMAIL_MOBILE_CLIENT_ID } from '@env';

async function HandleGoogleLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
) {
  const config: AuthConfiguration = {
    issuer: 'https://accounts.google.com',
    clientId: `${GMAIL_MOBILE_CLIENT_ID}.apps.googleusercontent.com`,
    redirectUrl: `com.googleusercontent.apps.${GMAIL_MOBILE_CLIENT_ID}:/oauth2redirect/google`,
    scopes: ['openid', 'profile'],
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    let data;
    if (login) {
      data = await handleCallback(
        `http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`,
        result,
      );
    } else {
      setToken(result.accessToken);
      // TODO: call route when loging in from myServices page (waiting for back to be done)
    }
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

export { HandleGoogleLogin };
