import React, { useEffect } from 'react';
import {
  GoogleSignin,
  isErrorWithCode,
} from '@react-native-google-signin/google-signin';

const GoogleOauth2 = () => {
  useEffect(() => {
    // Configure Google Sign-In on mount
    GoogleSignin.configure({
      webClientId:
        '616333423597-9gfqlpsa6l7520sbgrk8th02apobie8m.apps.googleusercontent.com',
      offlineAccess: true,
      forceCodeForRefreshToken: true,
    });
  }, []); // Empty dependency array means this runs once when the component is mounted

  return null; // Or you can return some UI, e.g., a button
}

const googleSign = async (navigation: any, setToken: (token: string) => void, ipAddress: string) => {
  console.log('1');
  try {
    await GoogleSignin.hasPlayServices();
    console.log('2');
    const userInfo = await GoogleSignin.signIn();
    console.log('3');
    const idToken = userInfo.data?.idToken;
    const userEmail = userInfo.data?.user.email;
    const userUsername = userInfo.data?.user.name;
    console.log('4');

    const resp = await fetch(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        token: String(idToken),
        username: userUsername,
        email: userEmail,
      }),
    });

    console.log('5');
    if (resp.status === 200) {
      type ResponseType = {
        token: string;
      };
      const token: ResponseType = await resp.json();
      setToken(token.token);
      navigation.navigate('AreaView');
    }
  } catch (error: any) {
    if (isErrorWithCode(error)) {
      console.error('Error with code:', error.code, '\n\terror message:', error.message);
    }
  }
};

export { GoogleOauth2, googleSign };
