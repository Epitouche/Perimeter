import React, { useContext, useEffect, useState } from 'react';
import {
  View,
  Text,
  FlatList,
  TouchableOpacity,
  StyleSheet,
  Linking,
  ActivityIndicator,
} from 'react-native';
import BottomNavBar from './NavBar';
import { AppContext } from '../context/AppContext';
import { HandleSpotifyLogin } from './Oauth2/SpotifyOauth2';
import { HandleMicrosoftLogin } from './Oauth2/MicrosoftOauth2';
import { HandleGithubLogin } from './Oauth2/GithubOauth2';
import { HandleDropboxLogin } from './Oauth2/DropboxOauth2';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';
import { SvgFromUri } from 'react-native-svg';
import { HandleGoogleLogin } from './Oauth2/GoogleOauth2';

type Props = NativeStackScreenProps<RootStackParamList, 'ServicesScreen'>;

/**
 * ServicesScreen component displays a list of services that the user can connect to or disconnect from.
 *
 * @param {object} props - The component props.
 * @param {object} props.navigation - The navigation object used for navigating between screens.
 *
 * @returns {JSX.Element} The rendered component.
 */
const ServicesScreen = ({ navigation }: { navigation: any }) => {
  /**
   * State to store the list of available services.
   */
  const [services, setServices] = useState([]);

  /**
   * State to store the list of connected services.
   */
  const [connectedServices, setConnectedServices] = useState<string[]>([]);

  /**
   * State to manage the loading state of the component.
   */
  const [loading, setLoading] = useState(true);

  /**
   * Context values from AppContext.
   */
  const { ipAddress, token, setToken } = useContext(AppContext);

  /**
   * Handles the disconnection of a service.
   *
   * @param {string} id - The ID of the service to disconnect.
   * @param {string} name - The name of the service to disconnect.
   */
  const handleDisconnect = async (id: string, name: string) => {
    try {
      await fetch(`http://${ipAddress}:8080/api/v1/token`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ id }),
      });
      setConnectedServices(connectedServices.filter(s => s !== name));
    } catch (error) {
      if ((error as any).code === 401) {
        navigation.navigate('Login');
      }
      console.error('Error disconnecting:', error);
    }
  };

  /**
   * Connects to a specified service.
   *
   * @param {string} service - The name of the service to connect to.
   */
  function connectService(service: string) {
    console.log('Connecting to:', service);
    switch (service) {
      case 'Spotify':
        HandleSpotifyLogin(setToken, navigation, ipAddress, false, token);
        break;
      case 'Google':
        HandleGoogleLogin(setToken, navigation, ipAddress, false, token);
        break;
      case 'Dropbox':
        HandleDropboxLogin(setToken, navigation, ipAddress, false, token);
        break;
      case 'Github':
        HandleGithubLogin(setToken, navigation, ipAddress, false, token);
        break;
      case 'Microsoft':
        HandleMicrosoftLogin(setToken, navigation, ipAddress, false, token);
        break;
      default:
        break;
    }
  }

  /**
   * Fetches the list of services and connected services from the API.
   */
  useEffect(() => {
    const fetchServices = async () => {
      try {
        const serviceResponse = await fetch(
          `http://${ipAddress}:8080/api/v1/service/info`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const userResponse = await fetch(
          `http://${ipAddress}:8080/api/v1/user/info/all`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const serviceData = await serviceResponse.json();
        const userData = await userResponse.json();

        setServices(serviceData);
        const connected = userData.tokens.map(
          (token: { service: { name: string } }) => token.service.name,
        );
        setConnectedServices(connected);
      } catch (error) {
        if ((error as any).code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching services:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchServices();
  }, []);

  /**
   * Handles URL events for OAuth redirection.
   *
   * @param {object} event - The URL event object.
   */
  const handleUrl = (event: any) => {
    console.log('Redirect URL:', event.url);
    if (event.url) {
      const url = new URL(event.url).searchParams;
      const token = url.get('token');
      const code = url.get('code');
      const error = url.get('error');

      if (code) {
        console.log('Received auth code:', code);
      } else if (error) {
        console.error('OAuth error:', error);
      } else if (token) {
        console.log('Received token:', token);
        setToken(token);
      }
    }
  };
  Linking.addEventListener('url', handleUrl);

  /**
   * Renders a single service item.
   *
   * @param {object} item - The service item to render.
   *
   * @returns {JSX.Element} The rendered service item.
   */
  const renderService = ({ item }: { item: any }) => {
    const isConnected = connectedServices.includes(item.name);

    return (
      <TouchableOpacity
        accessibilityLabel={`${!isConnected ? 'Connect' : 'Disconnect'} from ${
          item.name
        }`}
        accessibilityHint={`Double tap to ${
          !isConnected ? 'connect to' : 'disconnect from'
        } ${item.name}`}
        style={[styles.serviceButton, { backgroundColor: item.color }]}
        onPress={
          !isConnected
            ? () => connectService(item.name)
            : () => handleDisconnect(item.id, item.name)
        }>
        <SvgFromUri uri={item.icon} width={50} height={50} />
        <Text style={styles.serviceText}>{item.name}</Text>
        <Text style={styles.serviceText}>
          {isConnected || !item.oauth ? 'Connected' : 'Not Connected'}
        </Text>
      </TouchableOpacity>
    );
  };

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
    height: 120,
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
