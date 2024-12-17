import React, {useState, useEffect, useContext} from 'react';
import {View, Text, ActivityIndicator, StyleSheet} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../../Navigation/navigate';
import {AppContext} from '../../context/AppContext';
import { SpotifyOauthCallback } from './GoogleOauth2';
import { GoogleOauthCallback } from './SpotifyOauth2';
import { GithubOauthCallback } from './GithubOauth2';

type Props = NativeStackScreenProps<RootStackParamList, 'authRedirect'>;

const AuthRedirectScreen: React.FC<Props> = ({navigation, route}) => {
  const [isLoading, setIsLoading] = useState(true);
  const {service} = useContext(AppContext);
  const code = route.params?.code || '';

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
      navigation.goBack();
    }, 5000);

    return () => clearTimeout(timer);
  }, [navigation]);

  if (code) {
    if (service == 'Spotify') {
      SpotifyOauthCallback(code, navigation);
    } else if (service == 'Github') {
      GithubOauthCallback(code, navigation);
    } else if (service == 'Google') {
      GoogleOauthCallback(code, navigation);
    }
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
