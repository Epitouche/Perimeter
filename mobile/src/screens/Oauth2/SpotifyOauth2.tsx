import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';
import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { SPOTIFY_CLIENT_ID, SPOTIFY_SECRET } from '@env';

async function HandleSpotifyLogin(setToken: any, navigation: any) {
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
    if (error.message != 'User cancelled flow') {
      console.error('Failed to log in to Spotify', error);
      Alert.alert('Error', error.message);
    }
  }
}

async function GoogleOauthCallback(codeGoogle: string, navigation: any) {
  const { ipAddress, token, setToken } = useContext(AppContext);
  const response = await fetch(
    `http://${ipAddress}:8080/api/v1/google/auth/callback/mobile`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ codeGoogle }),
    },
  );
  console.log('response: ', response);
  const data = await response.json();
  if (data.error) {
    console.error(data.error);
    navigation.goBack();
  } else {
    setToken(data.accessToken);
    console.log('data: ', data);
    if (data.accessToken !== '') {
      navigation.navigate('AreaView');
    } else {
      console.error('Error: no token');
    }
  }
}

export { GoogleOauthCallback, HandleSpotifyLogin };
