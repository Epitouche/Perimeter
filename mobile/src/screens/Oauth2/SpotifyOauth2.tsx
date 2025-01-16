import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

/**
 * Handles the Spotify login process using OAuth2.
 *
 * @param {Function} setToken - Function to set the token received from Spotify.
 * @param {Object} navigation - Navigation object to navigate between screens.
 * @param {string} ipAddress - IP address of the server handling the callback.
 * @param {boolean} [login=true] - Flag to determine if the user should be navigated to the AreaView screen after login.
 * @param {string} [bearerToken=''] - Optional bearer token for authentication.
 * @returns {Promise<void>} - A promise that resolves when the login process is complete.
 * @throws {Error} - Throws an error if the login process fails.
 */
async function HandleSpotifyLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = true,
  bearerToken: string = '',
) {
  const config: AuthConfiguration = {
    clientId: SPOTIFY_CLIENT_ID,
    clientSecret: SPOTIFY_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: [
      'user-read-email',
      'playlist-modify-public',
      'user-modify-playback-state',
      'user-read-playback-state',
    ],
    serviceConfiguration: {
      authorizationEndpoint: 'https://accounts.spotify.com/authorize',
      tokenEndpoint: 'https://accounts.spotify.com/api/token',
    },
  };

  try {
    const result = await authorize(config);
    let data = await handleCallback(
      `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
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

export { HandleSpotifyLogin };
