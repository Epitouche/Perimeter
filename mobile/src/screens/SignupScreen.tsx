import React, { useContext, useState, useEffect } from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  Linking,
  Alert,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';
import { AppContext } from '../context/AppContext';
import { HandleGithubLogin } from './Oauth2/GithubOauth2';
import { HandleMicrosoftLogin } from './Oauth2/MicrosoftOauth2';
import { HandleSpotifyLogin } from './Oauth2/SpotifyOauth2';
import { HandleGoogleLogin } from './Oauth2/GoogleOauth2';
import { HandleDropboxLogin } from './Oauth2/DropboxOauth2';
import { SvgFromUri } from 'react-native-svg';

type Props = NativeStackScreenProps<RootStackParamList, 'SignUp'>;

/**
 * SignupScreen component allows users to sign up by providing their username, email, and password.
 * It also provides options to sign up using various OAuth services.
 *
 * @param {Props} props - The props for the SignupScreen component.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @param {object} props.route - The route object provided by React Navigation.
 *
 * @returns {JSX.Element} The rendered SignupScreen component.
 */

/**
 * Handles the URL event for OAuth redirection.
 *
 * @param {object} event - The event object containing the URL.
 */

/**
 * Handles the signup process by validating the input fields and making a POST request to the server.
 * If the signup is successful, navigates to the Login screen.
 * If there are errors, displays appropriate error messages.
 */

/**
 * Switches to the Login screen.
 */

/**
 * Fetches the available services from the server and updates the state.
 * This effect runs once when the component mounts.
 */

/**
 * Renders the SignupScreen component.
 *
 * @returns {JSX.Element} The rendered SignupScreen component.
 */
const SignupScreen: React.FC<Props> = ({ navigation, route }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
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
   * Handles the URL event, extracting and logging the authorization code or error from the URL.
   *
   * @param {any} event - The event object containing the URL to be handled.
   * @returns {void}
   */
  const handleUrl = (event: any) => {
    console.log('Redirect URL:', event.url);
    if (event.url) {
      const url = new URL(event.url).searchParams;
      const code = url.get('code');
      const error = url.get('error');

      if (code) {
        console.log('Received auth code:', code);
      } else if (error) {
        console.error('OAuth error:', error);
      }
    }
  };
  Linking.addEventListener('url', handleUrl);

  /**
   * Handles the signup process by validating user input and sending a registration request to the server.
   *
   * @async
   * @function handleSignup
   * @returns {Promise<void>} - A promise that resolves when the signup process is complete.
   *
   * @description
   * This function performs the following steps:
   * 1. Validates the `username`, `password`, and `email` fields.
   * 2. Sets error messages if any of the fields are empty.
   * 3. If there are no validation errors, sends a POST request to the server to register the user.
   * 4. If the registration is successful, navigates to the Login screen and shows a success alert.
   * 5. If the registration fails, logs the error and shows an error alert.
   * 6. If a network error occurs, logs the error.
   *
   * @throws Will throw an error if the fetch request fails.
   */
  const handleSignup = async () => {
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

    if (!email) {
      newErrors += " Email";
      hasError = true;
    }

    if (!email || !username || !password) {
      newErrors += " is required"
    }

    setErrors(newErrors);

    if (!hasError) {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/user/register`,
          {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email, username, password }),
          },
        );
        const data = await response.json();
        console.log('Data:', data);

        if (response.ok) {
          navigation.navigate('Login');
          Alert.alert('Successfully registered, please login now');
        } else if (response.status == 400 || response.status == 409) {
          console.log("error")
          setErrors(data.error)
        } else {
          console.error('Error:', response.status, ' | ', response.statusText);
          Alert.alert(
            'Error registering, please try again (Username or email might already be taken)',
          );
        }
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error', error);
      }
    }
  };

  const switchToLogin = () => {
    console.log('Switch to login');
    navigation.navigate('Login');
  };

  useEffect(() => {
    /**
     * Fetches services from the API and updates the state with the fetched data.
     *
     * This function makes a GET request to the `/api/v1/service/info` endpoint using the provided
     * `ipAddress` and `token` for authorization. The response is then parsed as JSON and used to
     * update the `services` state. Additionally, it logs the services that have OAuth enabled.
     *
     * @async
     * @function fetchServices
     * @returns {Promise<void>} A promise that resolves when the services have been fetched and the state has been updated.
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
      <Text style={styles.header}>Sign up</Text>

      <TextInput
        style={styles.input}
        placeholder="Enter username"
        placeholderTextColor="#aaa"
        value={username}
        inputMode="text"
        onChangeText={text => setUsername(text)}
        accessibilityHint="Enter your username"
      />

      <TextInput
        style={styles.input}
        placeholder="Enter email"
        placeholderTextColor="#aaa"
        value={email}
        inputMode="email"
        onChangeText={text => setEmail(text)}
        accessibilityHint="Enter your email address"
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
        style={styles.registerButton}
        onPress={handleSignup}
        accessibilityHint="Tap to sign up">
        <Text style={styles.signupButtonText}>Sign up</Text>
      </TouchableOpacity>

      <View style={styles.signUpContainer}>
        <Text style={styles.alreadySignUpText}>Already sign up ?</Text>
        <TouchableOpacity
          onPress={switchToLogin}
          accessibilityHint="Tap to switch to login screen">
          <Text style={styles.loginText}>Login</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.dividerContainer}>
        <View style={styles.divider} />
        <Text style={styles.orText}>or sign up with</Text>
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
              accessibilityHint={`Tap to sign up with ${service.name}`}>
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
  registerButton: {
    width: '100%',
    backgroundColor: '#000',
    padding: 12,
    borderRadius: 20,
    alignItems: 'center',
    marginBottom: 20,
  },
  signupButtonText: {
    color: '#fff',
    fontWeight: 'bold',
  },
  signUpContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 20,
  },
  alreadySignUpText: {
    marginRight: 5,
    color: '#555',
  },
  loginText: {
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

export default SignupScreen;
