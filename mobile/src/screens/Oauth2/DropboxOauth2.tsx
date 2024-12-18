import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';
import { AuthConfiguration, authorize } from 'react-native-app-auth';
import { DROPBOX_CLIENT_ID, DROPBOX_SECRET } from '@env';

async function HandleDropboxLogin(setToken: any, navigation: any) {
  const config: AuthConfiguration = {
    clientId: DROPBOX_CLIENT_ID,
    clientSecret: DROPBOX_SECRET,
    redirectUrl: 'com.perimeter-epitech://oauthredirect',
    scopes: ['files.metadata.read', 'files.content.read'],
    serviceConfiguration: {
      authorizationEndpoint: 'https://www.dropbox.com/oauth2/authorize',
      tokenEndpoint: 'https://api.dropboxapi.com/oauth2/token',
    },
  };

  try {
    const result = await authorize(config);
    console.log('result', result);
    setToken(result.accessToken);
    navigation.navigate('AreaView');
  } catch (error) {
    console.error('Failed to log in to Dropbox, ', error);
  }
}

async function DropboxOauthCallback(codeGoogle: string, navigation: any) {
  const { ipAddress, token, setToken } = useContext(AppContext);
  const response = await fetch(
    `http://${ipAddress}:8080/api/v1/dropbox/auth/callback/mobile`,
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

export { DropboxOauthCallback, HandleDropboxLogin };
