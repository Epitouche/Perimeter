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

/**
 * Provides the application context to its children components.
 *
 * @param {AppProviderProps} props - The properties for the AppProvider component.
 * @param {React.ReactNode} props.children - The child components that will have access to the context.
 *
 * @returns {JSX.Element} The AppContext provider with the given context value.
 *
 * @remarks
 * This component uses React's `useState` to manage the state of `ipAddress`, `token`, and `service`.
 * It also uses `React.useMemo` to memoize the context value, ensuring that the context is only updated
 * when one of the dependencies (`ipAddress`, `token`, `service`) changes.
 */
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
