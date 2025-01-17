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
import { HandleDropboxLogin } from './Oauth2/DropboxOauth2';
import { SvgFromUri } from 'react-native-svg';

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

/**
 * LoginScreen component allows users to log in using their username and password,
 * or through various OAuth services.
 *
 * @param {Props} props - The props for the LoginScreen component.
 * @param {object} props.navigation - The navigation object used to navigate between screens.
 *
 * @returns {JSX.Element} The rendered LoginScreen component.
 */

/**
 * Handles the login process by validating the input fields and making a POST request
 * to the login API endpoint. If successful, navigates to the AreaView screen.
 *
 * @async
 * @function handleLogin
 * @returns {Promise<void>}
 */

/**
 * Switches the screen to the SignUp screen.
 *
 * @function switchToSignup
 * @returns {void}
 */

/**
 * Fetches the available services from the API and sets the services state.
 *
 * @async
 * @function fetchServices
 * @returns {Promise<void>}
 */

/**
 * @typedef {Object} Service
 * @property {string} color - The color associated with the service.
 * @property {string} created_at - The creation date of the service.
 * @property {string} description - The description of the service.
 * @property {string} icon - The icon URL of the service.
 * @property {number} id - The unique identifier of the service.
 * @property {string} name - The name of the service.
 * @property {boolean} oauth - Indicates if the service supports OAuth.
 * @property {string} update_at - The last update date of the service.
 */

/**
 * @typedef {Object} Errors
 * @property {string} username - The error message for the username field.
 * @property {string} password - The error message for the password field.
 */
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
  const [errors, setErrors] = useState("");
  const { ipAddress, token, setToken, setService } = useContext(AppContext);

  /**
   * Handles the login process by validating the input fields and making a POST request to the login API.
   * If the username or password fields are empty, it sets the appropriate error messages.
   * If the fields are valid, it sends a login request to the server.
   * On successful login, it sets the token and navigates to the 'AreaView' screen.
   * If there is an error during the login process, it handles the error accordingly.
   *
   * @async
   * @function handleLogin
   * @returns {Promise<void>} A promise that resolves when the login process is complete.
   */
  const handleLogin = async () => {
    let hasError = false;
    let newErrors = "";

    if (!username) {
      console.log("Username is required");
      newErrors += " Username";
      hasError = true;
    }
    if (!password) {
      console.log("Password is required");
      newErrors += " Password";
      hasError = true;
    }

    if (!username || !password) {
      newErrors += " is required"
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
        if (response.ok) {
          navigation.navigate('AreaView');
        } else if (response.status == 400 || response.status == 409) {
          console.log("error")
          setErrors(data.error)
        }
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
    /**
     * Fetches services from the API and updates the state with the fetched data.
     *
     * @async
     * @function fetchServices
     * @returns {Promise<void>} A promise that resolves when the services are fetched and state is updated.
     * @throws {Error} Throws an error if the fetch request fails.
     *
     * @example
     * fetchServices();
     */
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
        accessibilityHint="Enter your username"
      />

      <TextInput
        style={styles.input}
        placeholder="Enter password"
        placeholderTextColor="#aaa"
        secureTextEntry
        value={password}
        onChangeText={text => setPassword(text)}
        accessibilityHint="Enter your password"
      />
      {errors != "" ? (
              <Text style={styles.errorText}>{errors}</Text>
      ) : null}

      <TouchableOpacity
        style={styles.loginButton}
        onPress={handleLogin}
        accessibilityHint="Tap to log in">
        <Text style={styles.loginButtonText}>Log in</Text>
      </TouchableOpacity>

      <View style={styles.signUpContainer}>
        <Text style={styles.newText}>New?</Text>
        <TouchableOpacity
          onPress={switchToSignup}
          accessibilityHint="Tap to sign up">
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
                  case 'Dropbox':
                    HandleDropboxLogin(setToken, navigation, ipAddress, true);
                    break;
                  default:
                    break;
                }
              }}
              accessibilityHint={`Tap to log in with ${service.name}`}>
              <SvgFromUri
                uri={service.icon}
                width={50}
                height={50}
                accessibilityLabel={`Connect with ${service.name}`}
              />
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
    color: '#001DDA',
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
    color: '#001DDA',
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
    color: '#E60000',
    fontSize: 12,
    alignSelf: 'flex-start',
    marginBottom: 10,
  },
});

export default LoginScreen;
