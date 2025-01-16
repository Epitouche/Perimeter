import React, { useState, useEffect } from 'react';
import { View, Text, ActivityIndicator, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'authRedirect'>;

/**
 * AuthRedirectScreen component handles the OAuth2 authentication redirect.
 * It displays a loading screen for 5 seconds before navigating back to the previous screen.
 *
 * @component
 * @param {object} props - The component props.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @param {object} props.route - The route object provided by React Navigation.
 * @returns {React.ReactElement|null} The rendered component or null.
 *
 * @example
 * <AuthRedirectScreen navigation={navigation} route={route} />
 */
const AuthRedirectScreen: React.FC<Props> = ({ navigation, route }) => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    /**
     * Sets a timeout to execute a function after 5 seconds.
     * The function will set the `isLoading` state to false and navigate back to the previous screen.
     *
     * @remarks
     * This timeout is used to simulate a delay, possibly for an authentication process or loading state.
     *
     * @param setIsLoading - Function to update the loading state.
     * @param navigation - Navigation object to handle screen transitions.
     *
     * @returns A timeout ID that can be used to clear the timeout if needed.
     */
    const timer = setTimeout(() => {
      setIsLoading(false);
      navigation.goBack();
    }, 5000);

    return () => clearTimeout(timer);
  }, [navigation]);

  if (isLoading) {
    return (
      <View style={styles.container} accessibilityHint="Loading screen">
        <ActivityIndicator
          size="large"
          color="#6200EE"
          accessibilityHint="Loading indicator"
        />
        <Text style={styles.loadingText} accessibilityHint="Loading text">
          Loading...
        </Text>
      </View>
    );
  }

  return null;
};

export default AuthRedirectScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#F5F5F5',
  },
  loadingText: {
    marginTop: 20,
    fontSize: 16,
    color: '#6200EE',
  },
});
