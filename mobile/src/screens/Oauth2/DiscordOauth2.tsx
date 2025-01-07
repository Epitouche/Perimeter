import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { DISCORD_CLIENT_ID, DISCORD_SECRET } from '@env';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleDiscordLogin(
  setToken: any,
  navigation: any,
  ipAddress: string,
  login: boolean = false,
) {
  const config: AuthConfiguration = {
    clientId: DISCORD_CLIENT_ID,
    clientSecret: DISCORD_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['identify', 'email', 'messages.read'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://discord.com/oauth2/authorize',
      tokenEndpoint: 'https://discord.com/api/oauth2/token',
    },
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    let data;
    if (login) {
      data = await handleCallback(
        `http://${ipAddress}:8080/api/v1/discord/auth/callback/mobile`,
        result,
      );
    } else {
      data = await handleCallback(
        `http://${ipAddress}:8080/api/v1/discord/auth/callback`,
        result,
      );
    }

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

export { HandleDiscordLogin };
