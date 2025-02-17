import React, { useContext, useEffect } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<
  RootStackParamList,
  'WorkflowReactionScreen'
>;

/**
 * WorkflowReactionScreen component is responsible for displaying the workflow reaction screen.
 * It fetches and displays the service information based on the provided actionId.
 *
 * @param {object} props - The props object.
 * @param {object} props.navigation - The navigation object used for navigating between screens.
 * @param {object} props.route - The route object containing route parameters.
 * @param {string} props.route.params.actionId - The ID of the action to fetch information for.
 * @param {object} props.route.params.actionOptions - Additional options for the action.
 *
 * @returns {JSX.Element} The rendered component.
 *
 * @component
 *
 * @example
 * return (
 *   <WorkflowReactionScreen navigation={navigation} route={route} />
 * )
 */
const WorkflowReactionScreen = ({ navigation, route }: Props) => {
  const { actionId, actionOptions } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [name, setName] = React.useState('');
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

  const [service, setService] = React.useState<Service | null>(null);

  useEffect(() => {
    /**
     * Fetches service information from the API and updates the state with the retrieved data.
     *
     * @async
     * @function getService
     * @throws Will navigate to the 'Login' screen if the response status is 401 (Unauthorized).
     * @throws Will log an error message to the console if there is an error during the fetch operation.
     *
     * @returns {Promise<void>} A promise that resolves when the service information has been fetched and the state has been updated.
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

        let data = await response.json();
        setName(data.name);
        setService(data.service);
      } catch (error) {
        if ((error as any).code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching service:', error);
      }
    };

    getService();
  }, [ipAddress, token]);

  return (
    <View style={styles.container}>
      <Text
        style={styles.title}
        accessibilityLabel="Add Area Title"
        accessibilityHint="Displays the title of the screen">
        Add Area
      </Text>
      <View
        style={[
          {
            flexDirection: 'row',
            backgroundColor: service ? service.color : 'black',
            alignItems: 'center',
            justifyContent: 'space-between',
            padding: 15,
            borderRadius: 8,
            width: '80%',
            marginBottom: 10,
          },
        ]}>
        <Text
          style={styles.boxText}
          accessibilityLabel={`action Name: ${name}`}
          accessibilityHint="Displays the name of the action">
          {name}
        </Text>
      </View>
      <View
        style={styles.line}
        accessibilityLabel="Separator Line"
        accessibilityHint="Separates the service name and reaction sections"
      />
      <View style={styles.actionBox}>
        <Text
          style={styles.boxText}
          accessibilityLabel="Reaction Label"
          accessibilityHint="Displays the label for the reaction section">
          Reaction
        </Text>
        <TouchableOpacity
          style={styles.addButton}
          onPress={() =>
            navigation.navigate('AddReactionScreen', {
              actionId,
              actionOptions,
            })
          }
          accessibilityLabel="Add Reaction Button"
          accessibilityHint="Navigates to the Add Reaction screen">
          <Text style={styles.addTextDisabled}>Add</Text>
        </TouchableOpacity>
      </View>
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
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: 'black',
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
});

export default WorkflowReactionScreen;
