import React, { useState, useEffect } from 'react';
import { View, Text, ActivityIndicator, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';

type Props = NativeStackScreenProps<RootStackParamList, 'authRedirect'>;

const AuthRedirectScreen: React.FC<Props> = ({ navigation, route }) => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
      navigation.goBack();
    }, 5000);

    return () => clearTimeout(timer);
  }, [navigation]);

  if (isLoading) {
    return (
      <View style={styles.container} accessibilityHint="Loading screen">
        <ActivityIndicator size="large" color="#6200EE" accessibilityHint="Loading indicator" />
        <Text style={styles.loadingText} accessibilityHint="Loading text">Loading...</Text>
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
