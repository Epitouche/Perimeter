import React, {createContext, useState, ReactNode} from 'react';

interface AppContextProps {
    ipAddress: string;
    setIpAddress: (ip: string) => void;
}

const AppContext = createContext<AppContextProps>({ ipAddress: '', setIpAddress: () => {} });

interface AppProviderProps { children: ReactNode }

export function AppProvider({ children }: AppProviderProps) {
    const [ipAddress, setIpAddress] = useState('');

    return (
        <AppContext.Provider value={{ ipAddress, setIpAddress }}>
            {children}
        </AppContext.Provider>
    );
}

export {AppContext};
export default AppProvider;