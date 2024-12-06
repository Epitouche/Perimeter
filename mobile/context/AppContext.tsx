import React, {createContext, useState, ReactNode} from 'react';

interface AppContextProps {
    ipAddress: string;
    setIpAddress: (ip: string) => void;
    token: string;
    setToken: (token: string) => void;
}

const AppContext = createContext<AppContextProps>({ ipAddress: '', setIpAddress: () => {}, token: '', setToken: () => {} });

interface AppProviderProps { children: ReactNode }

export function AppProvider({ children }: AppProviderProps) {
    const [ipAddress, setIpAddress] = useState('');
    const [token, setToken] = useState('');

    return (
        <AppContext.Provider value={{ ipAddress, setIpAddress, token, setToken }}>
            {children}
        </AppContext.Provider>
    );
}

export {AppContext};
export default AppProvider;