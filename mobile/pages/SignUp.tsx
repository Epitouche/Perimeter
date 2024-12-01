import {View, Text, Button} from 'react-native';
import {NavigationProp} from '@react-navigation/native';

export const SignUpScreen = ({
  navigation,
}: {
  readonly navigation: NavigationProp<any>;
}) => {
  return (
    <View style={{flex: 1, justifyContent: 'center', alignItems: 'center'}}>
      <Text>Sign Up</Text>
      <Button
        title="Go to Sign In"
        onPress={() => navigation.navigate('SignIn')}
      />
    </View>
  );
};
