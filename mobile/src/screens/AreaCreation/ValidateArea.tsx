import React, { useContext, useEffect } from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  StyleSheet,
  TextInput,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'ValidateAreaScreen'>;

/**
 * ValidateAreaScreen component is responsible for rendering the screen where users can validate and save an area.
 * It fetches action and reaction services based on the provided actionId and reactionId from the route parameters.
 * Users can input a title, description, and refresh timer for the area and save it.
 *
 * @param {object} props - The props object.
 * @param {object} props.navigation - The navigation object used for navigating between screens.
 * @param {object} props.route - The route object containing parameters passed to this screen.
 * @param {string} props.route.params.actionId - The ID of the action to fetch information for.
 * @param {object} props.route.params.actionOptions - The options for the action.
 * @param {string} props.route.params.reactionId - The ID of the reaction to fetch information for.
 * @param {object} props.route.params.reactionOptions - The options for the reaction.
 *
 * @returns {JSX.Element} The rendered component.
 */
const ValidateAreaScreen: React.FC<Props> = ({ navigation, route }) => {
  const { actionId, actionOptions, reactionId, reactionOptions } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  interface Service {
    color: string;
    created_at: string;
    description: string;
    icon: string;
    id: number;
    name: string;
    oauth: boolean;
    update_at: string;
  }

  const [actionService, setActionService] = React.useState<Service | null>(
    null,
  );
  const [reactionService, setReactionService] = React.useState<Service | null>(
    null,
  );
  const [actionName, setActionName] = React.useState('');
  const [reactionName, setReactionName] = React.useState('');
  const [title, setTitle] = React.useState('');
  const [description, setDescription] = React.useState('');
  const [refreshTimer, setRefreshTimer] = React.useState('');

  useEffect(() => {
    /**
     * Fetches action and reaction data from the server and updates the state with the retrieved information.
     * 
     * @async
     * @function getService
     * @throws Will navigate to the 'Login' screen if the response status is 401 (Unauthorized).
     * @throws Will log an error message to the console if there is an error during the fetch operation.
     * 
     * @returns {Promise<void>} A promise that resolves when the data has been successfully fetched and the state has been updated.
     */
    const getService = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/action/info/action/${actionId}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const res = await fetch(
          `http://${ipAddress}:8080/api/v1/reaction/info/reaction/${reactionId}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          },
        );

        const actionData = await response.json();
        const reactionData = await res.json();
        console.log('reactionData', reactionData);
        setActionName(actionData.name);
        setReactionName(reactionData.name);
        setActionService(actionData.service);
        setReactionService(reactionData.service);
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching service:', error);
      }
    };

    getService();
  }, [token, ipAddress, actionId, reactionId]);

  /**
   * Handles the save button press event by sending a POST request to the server
   * to save the area details. If successful, navigates to the 'AreaView' screen.
   * If there is an error and the error code is 401, navigates to the 'Login' screen.
   *
   * @async
   * @function saveButtonPressed
   * @returns {Promise<void>} A promise that resolves when the area is saved and navigation occurs.
   * @throws Will log an error message if the area saving fails.
   */
  const saveButtonPressed = async () => {
    try {
      console.log(parseInt(refreshTimer));
      let data = await fetch(`http://${ipAddress}:8080/api/v1/area`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          action_id: actionId,
          action_option: actionOptions,
          action_refresh_rate: parseInt(refreshTimer),
          description,
          reaction_id: reactionId,
          reaction_option: reactionOptions,
          title,
        }),
      });
      let res = await data.json();
      navigation.navigate('AreaView');
    } catch (error) {
      if (error.code === 401) {
        navigation.navigate('Login');
      }
      console.error('Error saving area:', error);
    }
  };

  return (
    <View style={styles.container}>
      <View style={{ width: '80%', marginBottom: 20 }}>
        <Text style={styles.label}>Title</Text>
        <TextInput
          style={styles.input}
          placeholder="Enter title"
          onChangeText={text => setTitle(text)}
          accessibilityLabel="Title Input"
          accessibilityHint="Enter the title for the area"
        />
      </View>
      <View style={{ width: '80%', marginBottom: 20 }}>
        <Text style={styles.label}>Description</Text>
        <TextInput
          style={styles.input}
          placeholder="Enter description"
          onChangeText={text => setDescription(text)}
          accessibilityLabel="Description Input"
          accessibilityHint="Enter the description for the area"
        />
      </View>
      <View style={{ width: '80%', marginBottom: 20 }}>
        <Text style={styles.label}>Refresh Timer (in seconds)</Text>
        <TextInput
          style={styles.input}
          placeholder="Enter refresh timer"
          keyboardType="numeric"
          onChangeText={text => setRefreshTimer(text)}
          accessibilityLabel="Refresh Timer Input"
          accessibilityHint="Enter the refresh timer in seconds"
        />
      </View>
      <Text style={styles.title}>Add Area</Text>
      <View
        style={[
          styles.actionBox,
          { backgroundColor: actionService?.color, borderRadius: 8 },
        ]}>
        <Text
          style={styles.boxText}
          accessibilityLabel="Action Name"
          accessibilityHint="Displays the name of the action">
          {actionName}
        </Text>
      </View>
      <View style={styles.line} />
      <View
        style={[
          styles.actionBox,
          { backgroundColor: reactionService?.color, borderRadius: 8 },
        ]}>
        <Text
          style={styles.boxText}
          accessibilityLabel="Reaction Name"
          accessibilityHint="Displays the name of the reaction">
          {reactionName}
        </Text>
      </View>
      <TouchableOpacity
        style={styles.saveButton}
        onPress={() => {
          saveButtonPressed();
        }}
        accessibilityLabel="Save Button"
        accessibilityHint="Press to save the area">
        <Text style={styles.addText}>Save</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#fff',
  },
  title: {
    fontSize: 24,
    marginBottom: 20,
  },
  actionBox: {
    flexDirection: 'row',
    backgroundColor: 'black',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 15,
    borderRadius: 8,
    width: '80%',
    marginBottom: 10,
  },
  boxText: {
    color: '#fff',
    fontSize: 18,
  },
  addButton: {
    backgroundColor: 'white',
    paddingVertical: 5,
    paddingHorizontal: 15,
    borderRadius: 5,
  },
  addButtonDisabled: {
    paddingVertical: 5,
    paddingHorizontal: 15,
    borderRadius: 5,
  },
  addText: {
    color: 'black',
    fontSize: 16,
    fontWeight: 'bold',
  },
  addTextDisabled: {
    color: 'gray',
    fontSize: 16,
    fontWeight: 'bold',
  },
  line: {
    width: 2,
    height: 20,
    backgroundColor: 'black',
    marginVertical: 10,
  },
  saveButton: {
    backgroundColor: '#1DC000',
    paddingVertical: 10,
    paddingHorizontal: 20,
    borderRadius: 8,
    marginTop: 20,
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    width: '100%',
    marginVertical: 5,
    paddingHorizontal: 10,
    color: 'black',
  },
  label: {
    fontSize: 16,
    marginBottom: 2,
  },
});

export default ValidateAreaScreen;
