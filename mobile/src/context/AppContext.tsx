import React, { createContext, useState, ReactNode } from 'react';

interface AppContextProps {
  ipAddress: string;
  setIpAddress: (ip: string) => void;
  token: string;
  setToken: (token: string) => void;
  service: string;
  setService: (service: string) => void;
}

const AppContext = createContext<AppContextProps>({
  ipAddress: '',
  setIpAddress: () => {},
  token: '',
  setToken: () => {},
  service: '',
  setService: () => {},
});

interface AppProviderProps {
  readonly children: ReactNode;
}

export function AppProvider({ children }: AppProviderProps) {
  const [ipAddress, setIpAddress] = useState('');
  const [token, setToken] = useState('');
  const [service, setService] = useState('');

  const contextValue = React.useMemo(
    () => ({
      ipAddress,
      setIpAddress,
      token,
      setToken,
      service,
      setService,
    }),
    [ipAddress, token, service],
  );

  return (
    <AppContext.Provider value={contextValue}>{children}</AppContext.Provider>
  );
}

export { AppContext };
export default AppProvider;
