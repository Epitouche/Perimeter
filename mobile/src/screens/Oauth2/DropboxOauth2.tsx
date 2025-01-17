import { Alert } from 'react-native';
import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { DROPBOX_CLIENT_ID, DROPBOX_SECRET } from '@env';
import { handleCallback } from './Callback';

/**
 * Handles the Dropbox login process using OAuth2.
 *
 * @param {Function} setToken - Function to set the authentication token.
 * @param {Object} navigation - Navigation object to navigate between screens.
 * @param {string} ipAddress - IP address of the server to handle the callback.
 * @param {boolean} [login=false] - Flag to determine if the user is logging in.
 * @param {string} [bearerToken=''] - Bearer token for authentication.
 *
 * @returns {Promise<void>} - A promise that resolves when the login process is complete.
 *
 * @throws {Error} - Throws an error if the login process fails.
 */
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
    scopes: [],
    serviceConfiguration: {
      authorizationEndpoint: 'https://www.dropbox.com/oauth2/authorize',
      tokenEndpoint: `https://www.dropbox.com/oauth2/token`,
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
