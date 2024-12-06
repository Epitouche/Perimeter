import React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import HomeScreen from './screens/HomeScreen';
import LoginScreen from './screens/LoginScreen';
import SignupScreen from './screens/SignupScreen';
import AreaView from './screens/AreaView';
import ServicesScreen from './screens/ServiceScreen';
import AuthRedirectScreen from './screens/authRedirect';
import AppProvider from './context/AppContext';

export type RootStackParamList = {
  Home: undefined;
  Login: undefined;
  SignUp: undefined;
  AreaView: undefined;
  ServicesScreen: undefined;
  authRedirect: {code: string};
};

const Stack = createNativeStackNavigator<RootStackParamList>();

const App: React.FC = () => {
  const linking = {
    prefixes: ['com.area://'],
    config: {
      screens: {
        authRedirect: 'oauthredirect',
      },
    },
  };

  return (
    <AppProvider>
      <NavigationContainer linking={linking}>
        <Stack.Navigator
          initialRouteName="Home"
          screenOptions={{headerShown: false}}>
          <Stack.Screen name="Home" component={HomeScreen} />
          <Stack.Screen name="Login" component={LoginScreen} />
          <Stack.Screen name="SignUp" component={SignupScreen} />
          <Stack.Screen name="AreaView" component={AreaView} />
          <Stack.Screen name="ServicesScreen" component={ServicesScreen} />
          <Stack.Screen name="authRedirect" component={AuthRedirectScreen} />
        </Stack.Navigator>
      </NavigationContainer>
    </AppProvider>
  );
};

export default App;
