import React from 'react';
import AppProvider from './src/context/AppContext';
import Navigation from './src/Navigation/navigate';

const App: React.FC = () => {

  return (
    <AppProvider>
      <Navigation />
    </AppProvider>
  );
};

export default App;
