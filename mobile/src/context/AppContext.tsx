import React, {createContext, useState, ReactNode} from 'react';

interface AppContextProps {
    ipAddress: string;
    setIpAddress: (ip: string) => void;
    token: string;
    setToken: (token: string) => void;
    codeVerifier: string;
    setCodeVerifier: (code: string) => void;
}

const AppContext = createContext<AppContextProps>({ ipAddress: '', setIpAddress: () => {}, token: '', setToken: () => {}, codeVerifier: '', setCodeVerifier: () => {} });

interface AppProviderProps { children: ReactNode }

export function AppProvider({ children }: AppProviderProps) {
    const [ipAddress, setIpAddress] = useState('');
    const [token, setToken] = useState('');
    const [codeVerifier, setCodeVerifier] = useState('');

    return (
        <AppContext.Provider value={{ ipAddress, setIpAddress, token, setToken, codeVerifier, setCodeVerifier }}>
            {children}
        </AppContext.Provider>
    );
}

export {AppContext};
export default AppProvider;
