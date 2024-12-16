import React, {useState, useEffect, useContext} from 'react';
import {View, Text, ActivityIndicator, StyleSheet} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../../Navigation/navigate';
import {AppContext} from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'authRedirect'>;

const AuthRedirectScreen: React.FC<Props> = ({navigation, route}) => {
  const [isLoading, setIsLoading] = useState(true);
  const {ipAddress, token, setToken, codeVerifier} = useContext(AppContext);
  const code = route.params?.code || '';

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
      navigation.goBack();
    }, 5000);

    return () => clearTimeout(timer);
  }, [navigation]);

  async function oauthCallback(codeSpotify: string) {
    const response = await fetch(
      `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({codeSpotify, code_verifier: codeVerifier}),
      },
    );
    console.log('response: ', response);
    const data = await response.json();
    if (data.error) {
      console.error(data.error);
      navigation.goBack();
    } else {
      setToken(data.token);
      console.log('data: ', data);
      if (data.token !== '') {
        navigation.navigate('AreaView');
      } else {
        console.error('Error: no token');
      }
    }
  }

  if (code) {
    oauthCallback(code);
  }

  if (isLoading) {
    return (
      <View style={styles.container}>
        <ActivityIndicator size="large" color="#6200EE" />
        <Text style={styles.loadingText}>Loading...</Text>
      </View>
    );
  }

  return null;
};

export default AuthRedirectScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5F5F5',
  },
  loadingText: {
    marginTop: 20,
    fontSize: 16,
    color: '#6200EE',
  },
});
