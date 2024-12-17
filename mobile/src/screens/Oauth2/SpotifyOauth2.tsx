import { useContext } from 'react';
import { AppContext } from '../../context/AppContext';

async function GoogleOauthCallback(codeGoogle: string, navigation: any) {
    const { ipAddress, token, setToken, } = useContext(AppContext);
    const response = await fetch(
        `http://${ipAddress}:8080/api/v1/google/auth/callback/mobile`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({ codeGoogle }),
        },
    );
    console.log('response: ', response);
    const data = await response.json();
    if (data.error) {
        console.error(data.error);
        navigation.goBack();
    } else {
        setToken(data.accessToken);
        console.log('data: ', data);
        if (data.accessToken !== '') {
            navigation.navigate('AreaView');
        } else {
            console.error('Error: no token');
        }
    }
}

export { GoogleOauthCallback }
