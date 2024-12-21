import React, {useContext, useEffect, useState} from 'react';
import {
  View,
  Text,
  FlatList,
  TouchableOpacity,
  StyleSheet,
  Linking,
  ActivityIndicator,
} from 'react-native';
import Svg, {Path} from 'react-native-svg';
import BottomNavBar from './NavBar';
import {AppContext} from '../context/AppContext';
import { HandleSpotifyLogin } from './Oauth2/SpotifyOauth2';
import { HandleGoogleLogin } from './Oauth2/GoogleOauth2';
import { HandleGithubLogin } from './Oauth2/GithubOauth2';
import { HandleDropboxLogin } from './Oauth2/DropboxOauth2';

const ServicesScreen = (navigation: any) => {
  const [services, setServices] = useState([]);
  const [loading, setLoading] = useState(true);
  const { ipAddress, token, setToken} = useContext(AppContext);

  const serviceIcons = {
    spotify: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M12 0a12 12 0 1012 12A12 12 0 0012 0zm5.33 17.46a.87.87 0 01-1.2.29 12.73 12.73 0 00-5.77-1.66 12.92 12.92 0 00-5.31 1.28.86.86 0 01-.82-1.53 14.68 14.68 0 016.13-1.46 14.41 14.41 0 016.53 1.86.87.87 0 01.29 1.22zm1.79-3.26a1.07 1.07 0 01-1.47.35 18.18 18.18 0 00-6.65-2.16 17.86 17.86 0 00-6.94 1.31 1.08 1.08 0 11-.84-2A20.77 20.77 0 0112 11a20.88 20.88 0 017.71 2.47 1.07 1.07 0 01.41 1.45zm.36-3.29a1.31 1.31 0 01-1.78.44A22.75 22.75 0 0012 9.66a22.92 22.92 0 00-8.68 1.7 1.31 1.31 0 01-1-.06 1.31 1.31 0 01.69-2.48A25.36 25.36 0 0112 7.34a25.37 25.37 0 019.62 2.06 1.32 1.32 0 01.44 1.86z"
        />
      </Svg>
    ),
    timer: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M15.07 1H8.93V2.5H15.07V1ZM12 7A8 8 0 1 0 20 15 8 8 0 0 0 12 7ZM18.93 14H12V8.93A6.12 6.12 0 0 1 18.93 14Z"
        />
      </Svg>
    ),
    gmail: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M12 13L3.5 6.5V18H20.5V6.5L12 13ZM3 4H21A1 1 0 0 1 22 5V19A1 1 0 0 1 21 20H3A1 1 0 0 1 2 19V5A1 1 0 0 1 3 4ZM20.5 5H3.5L12 10.5L20.5 5Z"
        />
      </Svg>
    ),
    dropbox: props => (
      <Svg viewBox="0 0 24 24" {...props}>
        <Path
          fill="white"
          d="M6.5 3L12 7.5 17.5 3 12 0 6.5 3zM0 6.5L6.5 11 12 6.5 5.5 2 0 6.5zM12 6.5L17.5 11 24 6.5 18.5 2 12 6.5zM0 13L6.5 17.5 12 13 5.5 8.5 0 13zM12 13L17.5 17.5 24 13 18.5 8.5 12 13zM6.5 18.5L12 24 17.5 18.5 12 14 6.5 18.5z"
        />
      </Svg>
    ),
  };

  useEffect(() => {
    const fetchServices = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/service/info`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const data = await response.json();
        setServices(data);
      } catch (error) {
        console.error('Error fetching services:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchServices();
  }, []);

  const handleUrl = (event: any) => {
    console.log('Redirect URL:', event.url);
    if (event.url) {
      const url = new URL(event.url).searchParams
      const token = url.get('token')
      const code = url.get('code')
      const error = url.get('error')

      if (code) {
        console.log('Received auth code:', code);
      } else if (error) {
        console.error('OAuth error:', error);
      } else if (token) {
        console.log('Received token:', token);
        setToken(token);
      }
    }
  }
  Linking.addEventListener('url', handleUrl);

  function connectService(service: string) {
    switch (service) {
      case 'spotify':
        HandleSpotifyLogin(setToken, navigation, ipAddress);
        break;
      case 'gmail':
        HandleGoogleLogin(setToken, navigation);
        break;
      case 'dropbox':
        HandleDropboxLogin(setToken, navigation);
        break;
      case 'github':
        HandleGithubLogin(setToken, navigation);
        break;
      default:
        break;
    }
  }

  const renderService = ({item}: {item: any}) => (
    <TouchableOpacity
      style={[styles.serviceButton, {backgroundColor: '#2196F3'}]} // Default color for unknown services
      onPress={() => connectService(item.name)}>
      {serviceIcons[item.name]?.({width: 36, height: 36}) || (
        <Text style={styles.unknownIcon}>?</Text>
      )}
      <Text style={styles.serviceText}>{item.name}</Text>
    </TouchableOpacity>
  );

  if (loading) {
    return (
      <View style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="#0000ff" />
        <Text>Loading services...</Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Services</Text>
      <View style={styles.separator} />
      <FlatList
        data={services}
        renderItem={renderService}
        keyExtractor={item => item.id.toString()}
        numColumns={2}
        contentContainerStyle={styles.list}
      />
      <BottomNavBar navigation={navigation} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    paddingHorizontal: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginTop: 20,
    textAlign: 'center',
  },
  separator: {
    height: 1,
    backgroundColor: '#ccc',
    marginVertical: 10,
  },
  list: {
    justifyContent: 'center',
    paddingVertical: 20,
  },
  serviceButton: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    margin: 10,
    height: 100,
    borderRadius: 10,
  },
  serviceText: {
    marginTop: 8,
    color: 'white',
    fontWeight: 'bold',
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  unknownIcon: {
    fontSize: 18,
    color: 'white',
  },
});

export default ServicesScreen;
