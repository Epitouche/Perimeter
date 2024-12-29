import { GoogleSignin } from '@react-native-google-signin/google-signin';
import { Alert } from 'react-native';
import { GMAIL_MOBILE_CLIENT_ID, GMAIL_SECRET } from '@env';
import { handleCallback } from './Callback';

async function HandleGoogleLogin(setToken: any, navigation: any, ipAddress: string, login: boolean = false) {
  GoogleSignin.configure({
    webClientId: GMAIL_MOBILE_CLIENT_ID,
    offlineAccess: true,
    forceCodeForRefreshToken: true,
  });

  try {
    // Sign in the user
    await GoogleSignin.hasPlayServices(); // Ensure Play Services are available
    const userInfo = await GoogleSignin.signIn();

    const { idToken } = userInfo;
    let data;

    if (login) {
      // Handle login callback with your backend
      data = await handleCallback(`http://${ipAddress}:8080/api/v1/google/auth/callback/mobile`, { idToken });
    } else {
      // Directly set token for client-side use
      setToken(idToken);
    }

    if (data?.error) {
      console.error(data.error);
    } else {
      setToken(data?.token || idToken);
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

export { HandleGoogleLogin };
