import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleSpotifyLogin(setToken: any, navigation: any, ipAddress: string, login: boolean = false) {
  const config: AuthConfiguration = {
    clientId: SPOTIFY_CLIENT_ID,
    clientSecret: SPOTIFY_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['user-read-email', 'playlist-modify-public'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://accounts.spotify.com/authorize',
      tokenEndpoint: 'https://accounts.spotify.com/api/token',
    },
  };

  try {
    const result = await authorize(config);
    // console.log('result', result);
    let data;
    if (login) {
      data = await handleCallback(`http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`, result);
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

export { HandleSpotifyLogin };
