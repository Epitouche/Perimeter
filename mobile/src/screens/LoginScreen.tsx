import React, { useState, useContext, useEffect } from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
} from 'react-native';
import 'url-search-params-polyfill';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';
import { AppContext } from '../context/AppContext';
import { HandleGithubLogin } from './Oauth2/GithubOauth2';
import { HandleMicrosoftLogin } from './Oauth2/MicrosoftOauth2';
import { HandleSpotifyLogin } from './Oauth2/SpotifyOauth2';
import { HandleGoogleLogin } from './Oauth2/GoogleOauth2';
import { SvgFromUri } from 'react-native-svg';

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen: React.FC<Props> = ({ navigation }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [services, setServices] = useState([
    {
      color: '',
      created_at: '',
      description: '',
      icon: '',
      id: 0,
      name: '',
      oauth: false,
      update_at: '',
    },
  ]);
  const [errors, setErrors] = useState({ username: '', password: '' });
  const { ipAddress, token, setToken, setService } = useContext(AppContext);

  const handleLogin = async () => {
    let hasError = false;

    const newErrors = { username: '', password: '' };

    if (!username) {
      newErrors.username = 'Username is required';
      hasError = true;
    }
    if (!password) {
      newErrors.password = 'Password is required';
      hasError = true;
    }

    setErrors(newErrors);

    if (!hasError) {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/user/login`,
          {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
          },
        );

        const data = await response.json();
        setToken(data.token);
        navigation.navigate('AreaView');
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error:', error);
      }
    }
  };

  const switchToSignup = () => {
    console.log('Switch to signup');
    navigation.navigate('SignUp');
  };

  useEffect(() => {
    const fetchServices = async () => {
      const serviceResponse = await fetch(
        `http://${ipAddress}:8080/api/v1/service/info`,
        {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        },
      );

      const serviceData = await serviceResponse.json();
      setServices(serviceData);
      console.log(serviceData.filter(service => service.oauth));
    };

    fetchServices();
  }, [ipAddress]);

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Log in</Text>

      <TextInput
        style={styles.input}
        placeholder="Enter username"
        placeholderTextColor="#aaa"
        value={username}
        onChangeText={text => setUsername(text)}
      />
      {errors.username ? (
        <Text style={styles.errorText}>{errors.username}</Text>
      ) : null}

      <TextInput
        style={styles.input}
        placeholder="Enter password"
        placeholderTextColor="#aaa"
        secureTextEntry
        value={password}
        onChangeText={text => setPassword(text)}
      />
      {errors.password ? (
        <Text style={styles.errorText}>{errors.password}</Text>
      ) : null}

      <TouchableOpacity>
        <Text style={styles.forgotPassword}>Forgot password?</Text>
      </TouchableOpacity>

      <TouchableOpacity style={styles.loginButton} onPress={handleLogin}>
        <Text style={styles.loginButtonText}>Log in</Text>
      </TouchableOpacity>

      <View style={styles.signUpContainer}>
        <Text style={styles.newText}>New?</Text>
        <TouchableOpacity onPress={switchToSignup}>
          <Text style={styles.signUpText}>Sign Up</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.dividerContainer}>
        <View style={styles.divider} />
        <Text style={styles.orText}>or log in with</Text>
        <View style={styles.divider} />
      </View>

      <View style={styles.socialIconsContainer}>
        {services
          .filter(service => service.oauth)
          .map(service => (
            <TouchableOpacity
              style={{
                backgroundColor: service.color,
                borderRadius: 8,
                marginHorizontal: 10,
                padding: 5,
              }}
              key={service.name}
              onPress={() => {
                setService(service.name);
                switch (service.name) {
                  case 'Github':
                    HandleGithubLogin(setToken, navigation, ipAddress, true);
                    break;
                  case 'Microsoft':
                    HandleMicrosoftLogin(setToken, navigation, ipAddress, true);
                    break;
                  case 'Spotify':
                    HandleSpotifyLogin(setToken, navigation, ipAddress, true);
                    break;
                  case 'Google':
                    HandleGoogleLogin(setToken, navigation, ipAddress, true);
                    break;
                  default:
                    break;
                }
              }}>
              <SvgFromUri uri={service.icon} width={50} height={50} />
            </TouchableOpacity>
          ))}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 20,
    backgroundColor: '#f9f9f9',
  },
  header: {
    fontSize: 32,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  input: {
    width: '100%',
    padding: 12,
    borderRadius: 20,
    borderWidth: 1,
    borderColor: '#ccc',
    marginBottom: 10,
    backgroundColor: '#fff',
  },
  forgotPassword: {
    alignSelf: 'flex-end',
    color: '#007BFF',
    marginBottom: 20,
  },
  loginButton: {
    width: '100%',
    backgroundColor: '#000',
    padding: 12,
    borderRadius: 20,
    alignItems: 'center',
    marginBottom: 20,
  },
  loginButtonText: {
    color: '#fff',
    fontWeight: 'bold',
  },
  signUpContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 20,
  },
  newText: {
    marginRight: 5,
    color: '#555',
  },
  signUpText: {
    color: '#007BFF',
    fontWeight: 'bold',
  },
  dividerContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginVertical: 20,
    width: '100%',
  },
  divider: {
    flex: 1,
    height: 1,
    backgroundColor: '#ccc',
  },
  orText: {
    marginHorizontal: 10,
    color: '#555',
  },
  socialIconsContainer: {
    flexDirection: 'row',
    justifyContent: 'center',
    marginTop: 10,
  },
  socialIcon: {
    width: 40,
    height: 40,
    marginHorizontal: 10,
  },
  errorText: {
    color: 'red',
    fontSize: 12,
    alignSelf: 'flex-start',
    marginBottom: 10,
  },
});

export default LoginScreen;
