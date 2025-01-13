import { Alert } from 'react-native';
import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { DROPBOX_CLIENT_ID, DROPBOX_SECRET } from '@env';
import { handleCallback } from './Callback';

async function HandleDropboxLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
  bearerToken: string = '',
) {
  const config: AuthConfiguration = {
    clientId: DROPBOX_CLIENT_ID,
    clientSecret: DROPBOX_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['account_info.read', 'profile', 'email', 'openid'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://www.dropbox.com/oauth2/authorize',
      tokenEndpoint: 'https://api.dropboxapi.com/oauth2/token',
    },
  };

  try {
    const result = await authorize(config);
    let data = await handleCallback(
      `http://${ipAddress}:8080/api/v1/dropbox/auth/callback/mobile`,
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

export { HandleDropboxLogin };
