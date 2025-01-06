import { GoogleSignin, statusCodes } from '@react-native-google-signin/google-signin';
import { Alert } from 'react-native';
import { handleCallback } from './Callback';

async function HandleGoogleLogin(setToken: any, navigation: any, ipAddress: string, login: boolean = false) {
  try {
    // Check if Google Play services are available
    await GoogleSignin.hasPlayServices();

    // Sign in the user
    const userInfo = await GoogleSignin.signIn();

    console.log('Google Sign-In Success:', userInfo);

    // Get the tokens (ID token and Access token)
    const { idToken, accessToken } = userInfo;

    let data;
    if (login) {
      // If logging in, send the tokens to your backend for further processing
      data = await handleCallback(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, { userInfo });
    } else {
      // If not logging in, just set the token
      setToken(accessToken);
    }

    // Handle the response from your backend or the token result
    if (data?.error) {
      console.error('Error:', data.error);
    } else {
      setToken(data?.token || accessToken);

      if (login) {
        // Navigate to the 'AreaView' screen after a successful login
        navigation.navigate('AreaView');
      }
    }

  } catch (error: any) {
    if (error.code === statusCodes.SIGN_IN_CANCELLED) {
      console.log('User cancelled the login flow');
    } else if (error.code === statusCodes.IN_PROGRESS) {
      console.log('Signing in');
    } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
      console.log('Play services not available');
    } else {
      console.log('Google Sign-In Error:', error);
      Alert.alert('Error', error.message);
    }
  }
}

export { HandleGoogleLogin };
