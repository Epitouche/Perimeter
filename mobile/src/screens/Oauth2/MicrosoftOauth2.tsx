import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';
import { MICROSOFT_MOBILE_ID } from '@env'

async function HandleMicrosoftLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
) {
  const config: AuthConfiguration = {
    clientId: MICROSOFT_MOBILE_ID,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['Mail.ReadWrite', 'User.Read', 'Mail.Send', 'offline_access'],
    serviceConfiguration: {
      authorizationEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/authorize',
      tokenEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/token',
    },
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    let data;
    if (login) {
      data = await handleCallback(
        `http://${ipAddress}:8080/api/v1/microsoft/auth/callback/mobile`,
        result,
      );
    } else {
      data = await handleCallback(
        `http://${ipAddress}:8080/api/v1/microsoft/auth/callback`,
        result,
      );
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

export { HandleMicrosoftLogin };
