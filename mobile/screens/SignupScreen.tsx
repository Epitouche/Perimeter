import React, {useContext, useState, useEffect} from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  Image,
  Linking,
} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../App';
import {authorize, AuthConfiguration} from 'react-native-app-auth';
import {AppContext} from '../context/AppContext';
import pkceChallenge from 'react-native-pkce-challenge';
import {
  GoogleSignin,
  statusCodes,
} from '@react-native-google-signin/google-signin';

type Props = NativeStackScreenProps<RootStackParamList, 'SignUp'>;

const SignupScreen: React.FC<Props> = ({navigation, route}) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [errors, setErrors] = useState({username: '', password: '', email: ''});
  const {ipAddress, setToken, setCodeVerifier} = useContext(AppContext);

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

  useEffect(() => {
    GoogleSignin.configure({
      webClientId:
        '616333423597-nh5d001itful769q51j0o0r54qbg4poq.apps.googleusercontent.com',
      offlineAccess: true,
      forceCodeForRefreshToken: true,
    });
  });

  const signUp = async () => {
    try {
      await GoogleSignin.hasPlayServices();

      const userInfo = await GoogleSignin.signIn();

      const idToken = userInfo.data?.idToken;
      console.log(userInfo);
      const resp = await fetch(
        `http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token: String(idToken) }),
      });
      const token = await resp.json();
      setToken(token.token);
      navigation.navigate('AreaView');
    } catch (error: any) {
      if (error.code === statusCodes.SIGN_IN_CANCELLED) {
        console.log('User cancelled the login flow');
        const resp = await fetch(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token: "cancelled" }),
        });
      } else if (error.code === statusCodes.IN_PROGRESS) {
        console.log('Signing in');
        const resp = await fetch(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token: "in progress" }),
        });
      } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
        console.log('Play services not available');
        const resp = await fetch(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token: "Play services not available" }),
        });
      } else {
        console.log('Some other error happened');
        const resp = await fetch(`http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token: "other" }),
        });
        console.log(error.message);
        console.log(error.code);
      }
    }
  };


  const handleSpotifyLogin = async () => {
    const challenge = pkceChallenge();

    setCodeVerifier(challenge.codeVerifier);

    const spotifyAuthConfig: AuthConfiguration = {
      clientId: 'a2720e8c24db49ee938e84b83d7c2da1', // Replace with env variable
      clientSecret: '9df3f1a07db44b7981036a0b04b52e51', // Replace with env variable
      redirectUrl: 'com.perimeter-epitech://oauthredirect',
      scopes: ['user-read-private', 'user-read-email'],
      serviceConfiguration: {
        authorizationEndpoint: 'https://accounts.spotify.com/authorize',
        tokenEndpoint: 'https://accounts.spotify.com/api/token',
      },
    };

    try {
      const authState = await authorize(spotifyAuthConfig);
      console.log('Spotify Auth State:', authState);
      console.log('Logged into Spotify successfully!');
      const resp = await fetch(
        `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({token: authState.accessToken}),
        },
      );
      const data = await resp.json();
      setToken(data.token);
      console.log('Bearer Token:', data.token);
    } catch (error) {
      console.log('Spotify Login Error:', error);
      const resp = await fetch(
        `http://${ipAddress}:8080/api/v1/spotify/auth/callback/mobile`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          // body: JSON.stringify({token: authState.accessToken}),
          body: JSON.stringify({token: "error" + error}),
        },
      );
      const data = await resp.json();
      setToken(data.token);
    }
  };

  const handleSignup = async () => {
    let hasError = false;
    const newErrors = {username: '', password: '', email: ''};

    if (!username) {
      newErrors.username = 'Username is required';
      hasError = true;
    }
    if (!password) {
      newErrors.password = 'Password is required';
      hasError = true;
    }

    if (!email) {
      newErrors.email = 'Email is required';
      hasError = true;
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
            body: JSON.stringify({email, username, password}),
          },
        );

        if (response.ok) {
          const data = await response.json();
          console.log('Data:', data);
          navigation.navigate('Login');
        } else {
          console.error('Error:', response.status, ' | ', response.statusText);
        }
      } catch (error) {
        console.error('Error', error);
      }
    }
  };

  const switchToLogin = () => {
    console.log('Switch to login');
    navigation.navigate('Login');
  };

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Sign up</Text>

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
        placeholder="Enter email"
        placeholderTextColor="#aaa"
        value={email}
        onChangeText={text => setEmail(text)}
      />
      {errors.email ? (
        <Text style={styles.errorText}>{errors.email}</Text>
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

      <TouchableOpacity style={styles.registerButton} onPress={handleSignup}>
        <Text style={styles.signupButtonText}>Sign up</Text>
      </TouchableOpacity>

      <View style={styles.signUpContainer}>
        <Text style={styles.alreadySignUpText}>Already sign up ?</Text>
        <TouchableOpacity onPress={switchToLogin}>
          <Text style={styles.loginText}>Login</Text>
        </TouchableOpacity>
      </View>

      <View style={styles.dividerContainer}>
        <View style={styles.divider} />
        <Text style={styles.orText}>or sign up with</Text>
        <View style={styles.divider} />
      </View>

      <View style={styles.socialIconsContainer}>
        <TouchableOpacity onPress={signUp}>
          <Image
            source={{uri: 'https://img.icons8.com/color/48/google-logo.png'}}
            style={styles.socialIcon}
          />
        </TouchableOpacity>
        <TouchableOpacity>
          <Image
            source={{uri: 'https://img.icons8.com/ios-glyphs/50/github.png'}}
            style={styles.socialIcon}
          />
        </TouchableOpacity>
        <TouchableOpacity onPress={handleSpotifyLogin}>
          <Image
            source={{uri: 'https://img.icons8.com/color/50/spotify.png'}}
            style={styles.socialIcon}
          />
        </TouchableOpacity>
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

export default SignupScreen;
