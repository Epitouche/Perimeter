import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';
import { MICROSOFT_CLIENT_ID } from '@env';

/**
 * Handles Microsoft OAuth2 login process.
 *
 * @param {Function} setToken - Function to set the authentication token.
 * @param {Object} navigation - Navigation object to navigate between screens.
 * @param {string} ipAddress - IP address of the server to handle the callback.
 * @param {boolean} [login=false] - Flag to indicate if the user is logging in.
 * @param {string} [bearerToken=''] - Bearer token for authentication.
 *
 * @returns {Promise<void>} - A promise that resolves when the login process is complete.
 *
 * @throws {Error} - Throws an error if the login process fails.
 */
async function HandleMicrosoftLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
  bearerToken: string = '',
) {
  const config: AuthConfiguration = {
    clientId: MICROSOFT_CLIENT_ID,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: [
      'Mail.ReadWrite',
      'Mail.Read',
      'User.Read',
      'Mail.Send',
      'offline_access',
      'calendars.Read',
      'calendars.ReadWrite',
    ],
    serviceConfiguration: {
      authorizationEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/authorize',
      tokenEndpoint:
        'https://login.microsoftonline.com/common/oauth2/v2.0/token',
    },
  };

  try {
    const result = await authorize(config);
    let data = await handleCallback(
      `http://${ipAddress}:8080/api/v1/microsoft/auth/callback/mobile`,
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

export { HandleMicrosoftLogin };
