import {AuthConfiguration, authorize} from 'react-native-app-auth';
import {useContext} from 'react';
import {AppContext} from '../../context/AppContext';
import pkceChallenge from 'react-native-pkce-challenge';

async function HandleSpotifyLogin() {
  const {setToken, setCodeVerifier, ipAddress} = useContext(AppContext);
  const challenge = pkceChallenge();

  setCodeVerifier(challenge.codeVerifier);

  const spotifyAuthConfig: AuthConfiguration = {
    clientId: 'a2720e8c24db49ee938e84b83d7c2da1', // Replace with env variable
    clientSecret: '9df3f1a07db44b7981036a0b04b52e51', // Replace with env variable
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['user-read-private', 'user-read-email'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://accounts.spotify.com/authorize',
      tokenEndpoint: 'https://accounts.spotify.com/api/token',
    },
  };

  try {
    const authState = await authorize(spotifyAuthConfig);
    console.log('Spotify Auth State:', authState);
    console.log('Logged into Spotify successfully!');
    const resp = await fetch(
      `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({token: authState.accessToken}),
      },
    );
    const data = await resp.json();
    setToken(data.token);
    console.log('Bearer Token:', data.token);
  } catch (error) {
    console.log('Spotify Login Error:', error);
  }
};

export {HandleSpotifyLogin};