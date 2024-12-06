import React, { useState, useEffect, useContext } from 'react';
import {
  View,
  Text,
  ActivityIndicator,
  StyleSheet,
} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../App';
import { AppContext } from '../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'authRedirect'>;

const AuthRedirectScreen: React.FC<Props> = ({ navigation, route }) => {
  const [isLoading, setIsLoading] = useState(true);
  const {ipAddress, setToken} = useContext(AppContext);
  const code = route.params?.code || '';

  useEffect(() => {

    const timer = setTimeout(() => {
      setIsLoading(false);
      navigation.goBack();
    }, 3000);

    return () => clearTimeout(timer);
  }, [navigation]);

  async function oauthCallback(code: string) {
    const response = await fetch(`gttp://${ipAddress}/spotify/auth/callback`, // change it to be modular with route name (change spotify with something else...)
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ code }),
      }
    ) 
    const data = await response.json()
    if (data.error) {
      console.error(data.error)
      navigation.goBack();
    }
    setToken(data.token)
    navigation.navigate("AreaView")
    console.log("data: ", data)
  }

  if (code) {
    console.log("previous page: ", navigation.getParent())
    oauthCallback(code)
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
