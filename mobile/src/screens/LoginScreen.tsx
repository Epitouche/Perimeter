import React, {useState, useContext, useEffect} from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  Image,
} from 'react-native';
import 'url-search-params-polyfill';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../Navigation/navigate';
import {AppContext} from '../context/AppContext';
import {
  GoogleSignin,
  isErrorWithCode,
} from '@react-native-google-signin/google-signin';
import {HandleSpotifyLogin} from './Oauth2/OAuth2';

type Props = NativeStackScreenProps<RootStackParamList, 'Login'>;

const LoginScreen: React.FC<Props> = ({navigation, route}) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errors, setErrors] = useState({username: '', password: ''});
  const {ipAddress, setToken} = useContext(AppContext);

  useEffect(() => {
    GoogleSignin.configure({
      webClientId:
        '616333423597-9gfqlpsa6l7520sbgrk8th02apobie8m.apps.googleusercontent.com',
      offlineAccess: true,
      forceCodeForRefreshToken: true,
    });
  });

  const signIn = async () => {
    try {
      await GoogleSignin.hasPlayServices();
      const userInfo = await GoogleSignin.signIn();

      const idToken = userInfo.data?.idToken;
      const userEmail = userInfo.data?.user.email;
      const userUsername = userInfo.data?.user.name;
      console.log('SUIIIIIIIIIIIIIIII:', userInfo.data);
      const resp = await fetch(
        `http://${ipAddress}:8080/api/v1/gmail/auth/callback/mobile`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            token: String(idToken),
            username: userUsername,
            email: userEmail,
          }),
        },
      );
      if (resp.status === 200) {
        type ResponseType = {
          token: string;
        };
        const token: ResponseType = await resp.json();
        setToken(token.token);
        navigation.navigate('AreaView');
      }
    } catch (error: any) {
      if (isErrorWithCode(error)) {
        console.error(
          'Error with code:',
          error.code,
          '\n\terror message:',
          error.message,
        );
      }
    }
  };

  const handleLogin = async () => {
    let hasError = false;
    const newErrors = {username: '', password: ''};

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
            body: JSON.stringify({username, password}),
          },
        );

        if (response.ok) {
          const data = await response.json();
          setToken(data.token);
          navigation.navigate('AreaView');
        } else {
          console.error('Error:', response.status);
        }
      } catch (error) {
        console.error('Error:', error);
      }
    }
  };

  const switchToSignup = () => {
    console.log('Switch to signup');
    navigation.navigate('SignUp');
  };

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
        <TouchableOpacity onPress={signIn}>
          <Image
            source={{uri: 'https://img.icons8.com/color/48/google-logo.png'}}
            style={styles.socialIcon}
          />
        </TouchableOpacity>
        <Image
          source={{uri: 'https://img.icons8.com/ios-glyphs/50/github.png'}}
          style={styles.socialIcon}
        />
        <TouchableOpacity onPress={() => {console.log("tamere"); HandleSpotifyLogin()}}>
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
