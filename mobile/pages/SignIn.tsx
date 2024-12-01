import {View, Text, Button, TextInput} from 'react-native';
import {NavigationProp} from '@react-navigation/native';

export const SignInScreen = ({
  navigation,
}: {
  readonly navigation: NavigationProp<any>;
}) => {
  return (
    <View style={{backgroundColor: '#DFDCDC', width: '100%', height: '100%'}}>
      <View
        style={{
          flex: 1,
          margin: 25,
          padding: 10,
          backgroundColor: 'rgb(255, 255, 255)',
          borderRadius: 18,
        }}>
        <Text style={{fontSize: 64, textAlign: 'center', marginBottom: 24}}>
          Log In
        </Text>
        <View
          id="usernameInput"
          style={{paddingLeft: 12, paddingRight: 12, paddingBottom: 10}}>
          <Text style={{fontSize: 24}}>Enter username</Text>
          <TextInput
            style={{
              borderColor: '#000000',
              borderRadius: 30,
              padding: 13,
              marginTop: 10,
              marginBottom: 10,
              color: '#000000',
              backgroundColor: '#DFDCDC',
              fontSize: 16,
            }}
            placeholder="Username"></TextInput>
        </View>

        <View id="passwordInput" style={{paddingLeft: 12, paddingRight: 12}}>
          <Text style={{fontSize: 24}}>Enter password</Text>
          <TextInput
            style={{
              borderColor: '#000000',
              borderRadius: 30,
              padding: 13,
              marginTop: 10,
              marginBottom: 10,
              color: '#000000',
              backgroundColor: '#DFDCDC',
              fontSize: 16,
            }}
            placeholder="Password"></TextInput>
          <Text
            id="forgotPassword"
            style={{fontSize: 12, color: '#005EFF', textDecorationLine: 'underline}}
            onPress={() => navigation.navigate('SignUp')}> {/* TODO change Url to corect one once done */}
            Forgot Password ?
          </Text>
        </View>
      </View>
    </View>
  );
};

// {
/* <Button
  title="Go to Sign Up"
  onPress={() => navigation.navigate('SignUp')}
/> */
// }
