import React, {createContext, useState, ReactNode} from 'react';

interface AppContextProps {
    ipAddress: string;
    setIpAddress: (ip: string) => void;
    token: string;
    setToken: (token: string) => void;
    serviceConnecting: string;
    setServiceConnecting: (service: string) => void;
}

const AppContext = createContext<AppContextProps>({ ipAddress: '', setIpAddress: () => {}, token: '', setToken: () => {}, serviceConnecting: '', setServiceConnecting: () => {} });

interface AppProviderProps { children: ReactNode }

export function AppProvider({ children }: AppProviderProps) {
    const [ipAddress, setIpAddress] = useState('');
    const [token, setToken] = useState('');
    const [serviceConnecting, setServiceConnecting] = useState('');

    return (
        <AppContext.Provider value={{ ipAddress, setIpAddress, token, setToken, serviceConnecting, setServiceConnecting }}>
            {children}
        </AppContext.Provider>
    );
}

export {AppContext};
export default AppProvider;