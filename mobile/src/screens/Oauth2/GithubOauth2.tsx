import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { GITHUB_MOBILE_SECRET, GITHUB_MOBILE_CLIENT_ID } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

/**
 * Handles the GitHub login process using OAuth2.
 *
 * @param {Function} setToken - Function to set the authentication token.
 * @param {Object} navigation - Navigation object to navigate between screens.
 * @param {string} ipAddress - IP address of the server handling the OAuth callback.
 * @param {boolean} [login=false] - Flag to indicate if the user is logging in.
 * @param {string} [bearerToken=''] - Bearer token for authentication.
 *
 * @throws {Error} If the login process fails for any reason other than user cancellation.
 *
 * @returns {Promise<void>} A promise that resolves when the login process is complete.
 */
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
