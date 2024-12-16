import React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import HomeScreen from './screens/HomeScreen';
import LoginScreen from './screens/LoginScreen';
import SignupScreen from './screens/SignupScreen';
import AreaView from './screens/AreaView';
import ServicesScreen from './screens/ServiceScreen';
import AuthRedirectScreen from './screens/Oauth2/authRedirect';
import AppProvider from './context/AppContext';
import WorkflowScreen from './screens/Workflow';
import AddActionScreen from './screens/AreaCreation/AddAction';
import SelectActionScreen from './screens/AreaCreation/SelectAction';
import WorkflowReactionScreen from './screens/WorkflowReaction';
import AddReactionScreen from './screens/AreaCreation/AddReaction';
import SelectReactionScreen from './screens/AreaCreation/SelectReaction';

export type RootStackParamList = {
  Home: undefined;
  Login: undefined;
  SignUp: undefined;
  AreaView: undefined;
  ServicesScreen: undefined;
  authRedirect: {code: string};
  WorkflowScreen: undefined;
  AddActionScreen: undefined;
  SelectActionScreen: {serviceId: number};
  WorkflowReactionScreen: {actionId: number, actionOptions: { [key: string]: string }}; 
  AddReactionScreen: {actionId: number, actionOptions: { [key: string]: string }};
  SelectReactionScreen: {actionId: number, actionOptions: { [key: string]: string }, serviceId: number};
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
          <Stack.Screen name="WorkflowScreen" component={WorkflowScreen} />
          <Stack.Screen name="AddActionScreen" component={AddActionScreen} />
          <Stack.Screen name="SelectActionScreen" component={SelectActionScreen} />
          <Stack.Screen name="WorkflowReactionScreen" component={WorkflowReactionScreen} />
          <Stack.Screen name="AddReactionScreen" component={AddReactionScreen} />
          <Stack.Screen name="SelectReactionScreen" component={SelectReactionScreen} />
        </Stack.Navigator>
      </NavigationContainer>
    </AppProvider>
  );
};

export default App;