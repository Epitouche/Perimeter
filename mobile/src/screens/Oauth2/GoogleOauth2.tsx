import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';
import { GOOGLE_MOBILE_CLIENT_ID } from '@env';

async function HandleGoogleLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = true,
  bearerToken: string = '',
) {
  const config: AuthConfiguration = {
    issuer: 'https://accounts.google.com',
    clientId: `${GOOGLE_MOBILE_CLIENT_ID}.apps.googleusercontent.com`,
    redirectUrl: `com.googleusercontent.apps.${GOOGLE_MOBILE_CLIENT_ID}:/oauth2redirect/google`,
    scopes: ['https://mail.google.com/', 'profile', 'email'],
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    let data = await handleCallback(
      `http://${ipAddress}:8080/api/v1/google/auth/callback/mobile`,
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

export { HandleGoogleLogin };
