import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HomeScreen from '../screens/HomeScreen';
import LoginScreen from '../screens/LoginScreen';
import SignupScreen from '../screens/SignupScreen';
import AreaView from '../screens/AreaView';
import ServicesScreen from '../screens/ServiceScreen';
import AuthRedirectScreen from '../screens/Oauth2/authRedirect';
import WorkflowScreen from '../screens/AreaCreation/Workflow';
import AddActionScreen from '../screens/AreaCreation/AddAction';
import SelectActionScreen from '../screens/AreaCreation/SelectAction';
import WorkflowReactionScreen from '../screens/AreaCreation/WorkflowReaction';
import AddReactionScreen from '../screens/AreaCreation/AddReaction';
import SelectReactionScreen from '../screens/AreaCreation/SelectReaction';
import SettingsScreen from '../screens/SettingsScreen';
import AreaDetailsScreen from '../screens/AreaDetails/AreaDetails';
import ValidateAreaScreen from '../screens/AreaCreation/ValidateArea';

export type RootStackParamList = {
  Home: undefined;
  Login: undefined;
  SignUp: undefined;
  AreaView: undefined;
  ServicesScreen: undefined;
  authRedirect: { code: string };
  WorkflowScreen: undefined;
  AddActionScreen: undefined;
  SelectActionScreen: { serviceId: number };
  WorkflowReactionScreen: {
    actionId: number;
    actionOptions: { [key: string]: string };
  };
  AddReactionScreen: {
    actionId: number;
    actionOptions: { [key: string]: string };
  };
  SelectReactionScreen: {
    actionId: number;
    actionOptions: { [key: string]: string };
    serviceId: number;
  };
  ValidateAreaScreen: {
    actionId: number;
    actionOptions: { [key: string]: string };
    reactionId: number;
    reactionOptions: { [key: string]: string };
  };
  SettingsScreen: undefined;
  AreaDetails: { area: any };
};

const Stack = createNativeStackNavigator<RootStackParamList>();

/**
 * Navigation component that sets up the navigation structure for the application.
 * 
 * This component uses React Navigation to define a stack navigator with various screens.
 * It also configures deep linking for the application.
 * 
 * @returns {JSX.Element} The NavigationContainer component with the defined stack navigator.
 * 
 * @example
 * ```tsx
 * import Navigation from './Navigation';
 * 
 * const App = () => {
 *   return (
 *     <Navigation />
 *   );
 * };
 * 
 * export default App;
 * ```
 * 
 * @component
 * @example
 * ```tsx
 * <Navigation />
 * ```
 */
const Navigation = () => {
  const linking = {
    prefixes: ['com.perimeter-epitech://'],
    config: {
      screens: {
        authRedirect: 'oauthredirect',
      },
    },
  };

  return (
    <NavigationContainer linking={linking}>
      <Stack.Navigator
        initialRouteName="Home"
        screenOptions={{ headerShown: false }}>
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen
          name="Login"
          component={LoginScreen}
          options={{
            gestureEnabled: false,
          }}
        />
        <Stack.Screen
          name="SignUp"
          component={SignupScreen}
          options={{
            gestureEnabled: false,
          }}
        />
        <Stack.Screen name="AreaView" component={AreaView} />
        <Stack.Screen name="ServicesScreen" component={ServicesScreen} />
        <Stack.Screen name="authRedirect" component={AuthRedirectScreen} />
        <Stack.Screen name="WorkflowScreen" component={WorkflowScreen} />
        <Stack.Screen name="AddActionScreen" component={AddActionScreen} />
        <Stack.Screen
          name="SelectActionScreen"
          component={SelectActionScreen}
        />
        <Stack.Screen
          name="WorkflowReactionScreen"
          component={WorkflowReactionScreen}
        />
        <Stack.Screen name="AddReactionScreen" component={AddReactionScreen} />
        <Stack.Screen
          name="SelectReactionScreen"
          component={SelectReactionScreen}
        />
        <Stack.Screen
          name="ValidateAreaScreen"
          component={ValidateAreaScreen}
        />
        <Stack.Screen name="SettingsScreen" component={SettingsScreen} />
        <Stack.Screen name="AreaDetails" component={AreaDetailsScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default Navigation;
