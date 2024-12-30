import { useEffect } from 'react';
import AppProvider from './src/context/AppContext';
import Navigation from './src/Navigation/navigate';
import { GoogleSignin } from '@react-native-google-signin/google-signin';
import { GMAIL_CLIENT_ID } from '@env';

const App: React.FC = () => {
  console.log('App.tsx');
  useEffect(() => {
    GoogleSignin.configure({
      webClientId: GMAIL_CLIENT_ID,
      offlineAccess: true,
      forceCodeForRefreshToken: true,
    });
  });

  return (
    <AppProvider>
      <Navigation>
      </Navigation>
    </AppProvider>
  );
};

export default App;