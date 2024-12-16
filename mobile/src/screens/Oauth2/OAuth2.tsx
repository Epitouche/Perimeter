import {AuthConfiguration, authorize} from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';

export async function HandleSpotifyLogin(setToken: any) {

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
    console.log('result', result);
    setToken(result.accessToken);
  } catch (error) {
    console.error('Failed to log in to Spotify', error);
  }
};
