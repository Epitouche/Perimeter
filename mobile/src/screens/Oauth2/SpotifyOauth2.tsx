import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';
import { Alert } from 'react-native';

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
    if (login) {
      const response = await fetch(
        `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            accessToken: result.accessToken,
            refreshToken: result.refreshToken,
            
          }),
        },
      );
      const data = await response.json();
      console.log('data: ', data);
      if (data.error) {
        console.error(data.error);
      } else {
        setToken(result.accessToken);
        navigation.navigate('AreaView');
      }
    } else {
      setToken(result.accessToken);
      // TODO: call route when loging in from myServices page (waiting for back to be done)
    }
  } catch (error) {
    if ((error as Error).message != 'User cancelled flow') {
      console.error('Failed to log in to Spotify', error);
      Alert.alert('Error', (error as Error).message);
    }
  }
}

export { HandleSpotifyLogin };
