import React, { useContext } from 'react';
import {
  View,
  Text,
  TextInput,
  StyleSheet,
  TouchableOpacity,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../Navigation/navigate';
import { AppContext } from '../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'Home'>;

/**
 * HomeScreen component allows the user to enter an IP address and navigate to the login screen.
 *
 * @component
 * @param {object} props - The properties object.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @returns {React.FC<Props>} A React functional component.
 *
 * @example
 * return (
 *   <HomeScreen navigation={navigation} />
 * );
 */

/**
 * Context hook to get and set the IP address.
 * @constant
 * @type {object}
 * @property {string} ipAddress - The current IP address.
 * @property {function} setIpAddress - Function to update the IP address.
 */

/**
 * Renders the main container view.
 * @constant
 * @type {object}
 */

/**
 * Renders the instruction text for entering the IP address.
 * @constant
 * @type {object}
 * @property {string} accessibilityHint - Accessibility hint for the instruction text.
 */

/**
 * Renders the input field for entering the IP address.
 * @constant
 * @type {object}
 * @property {object} style - The style object for the input field.
 * @property {string} placeholder - Placeholder text for the input field.
 * @property {string} value - The current value of the input field.
 * @property {function} onChangeText - Function to handle text change in the input field.
 * @property {string} keyboardType - The type of keyboard to display.
 * @property {string} accessibilityHint - Accessibility hint for the input field.
 */

/**
 * Renders the connect button to set the IP address and navigate to the login screen.
 * @constant
 * @type {object}
 * @property {string} title - The title of the button.
 * @property {function} onPress - Function to handle button press.
 * @property {string} accessibilityHint - Accessibility hint for the button.
 */
const HomeScreen: React.FC<Props> = ({ navigation }) => {
  const { ipAddress, setIpAddress } = useContext(AppContext);

  return (
    <View style={styles.container}>
      <Text accessibilityHint="Instruction to enter the IP address">
        Enter the IP address to ping:
      </Text>
      <TextInput
        style={styles.input}
        placeholder="Enter IP address"
        value={ipAddress}
        onChangeText={setIpAddress}
        keyboardType="numeric"
        accessibilityHint="Input field for IP address"
      />
      <TouchableOpacity
        onPress={() => {
          navigation.navigate('Login');
        }}
        accessibilityLabel="Connect"
        accessibilityHint="Navigates to the login screen"
        style={{ backgroundColor: 'blue', padding: 10, borderRadius: 5 }}>
        <Text style={{ color: 'white' }}>Connect</Text>
      </TouchableOpacity>
    </View>
  );
};

export default HomeScreen;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 16,
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    width: '100%',
    marginVertical: 10,
    paddingHorizontal: 10,
  },
  responseContainer: {
    marginTop: 20,
    padding: 10,
    backgroundColor: '#f0f0f0',
    borderRadius: 5,
    width: '100%',
  },
  responseText: {
    fontWeight: 'bold',
    marginBottom: 5,
  },
});
