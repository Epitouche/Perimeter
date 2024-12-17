import {AuthConfiguration, authorize} from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET, GITHUB_SECRET, GITHUB_CLIENT_ID } from '@env';

export async function HandleSpotifyLogin(setToken: any, navigation: any) {

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
    navigation.navigate('AreaView');
  } catch (error) {
    console.error('Failed to log in to Spotify', error);
  }
};

export async function HandleGithubLogin(setToken: any, navigation: any) {
  const config: AuthConfiguration = {
    clientId: GITHUB_CLIENT_ID,
    clientSecret: GITHUB_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['user', 'repo'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://github.com/login/oauth/authorize',
      tokenEndpoint: 'https://github.com/login/oauth/access_token',
    },
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    setToken(result.accessToken);
    navigation.navigate('AreaView');
  } catch (error) {
    console.error('Failed to log in to GitHub', error);
  }
}