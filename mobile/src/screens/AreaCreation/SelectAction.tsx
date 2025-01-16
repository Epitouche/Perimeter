import React, { useState, useEffect, useContext } from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  ScrollView,
  ActivityIndicator,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'SelectActionScreen'>;

/**
 * SelectActionScreen component allows users to select and configure actions.
 * 
 * @component
 * @param {object} props - The component props.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @param {object} props.route - The route object provided by React Navigation.
 * @param {object} props.route.params - The parameters passed to the route.
 * @param {string} props.route.params.serviceId - The ID of the service to fetch actions for.
 * 
 * @returns {JSX.Element} The rendered component.
 * 
 * @example
 * <SelectActionScreen navigation={navigation} route={route} />
 * 
 * @remarks
 * This component fetches actions from an API based on the provided service ID.
 * It allows users to search for actions, select an action, configure its options, and save the configuration.
 * If the user is not authenticated, they will be redirected to the login screen.
 * 
 * @function
 * @name SelectActionScreen
 * 
 * @typedef {object} Props
 * @property {object} navigation - The navigation object provided by React Navigation.
 * @property {object} route - The route object provided by React Navigation.
 * @property {object} route.params - The parameters passed to the route.
 * @property {string} route.params.serviceId - The ID of the service to fetch actions for.
 * 
 * @typedef {object} Action
 * @property {string} id - The ID of the action.
 * @property {string} name - The name of the action.
 * @property {object} option - The options available for the action.
 * 
 * @typedef {object} Service
 * @property {string} id - The ID of the service.
 * @property {string} name - The name of the service.
 * @property {object} service - The service details.
 * @property {string} service.color - The color associated with the service.
 * 
 * @typedef {object} SelectedActionOptions
 * @property {string} key - The key of the option.
 * @property {any} value - The value of the option.
 * 
 * @typedef {object} AppContext
 * @property {string} ipAddress - The IP address of the API server.
 * @property {string} token - The authentication token.
 */
const SelectActionScreen: React.FC<Props> = ({ navigation, route }) => {
  const [services, setServices] = useState<any[]>([]);
  const [filteredServices, setFilteredServices] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [selectedAction, setSelectedAction] = useState<any | null>(null);
  const [selectedActionOptions, setSelectedActionOptions] = useState<{
    [key: string]: any;
  }>({});
  const { ipAddress, token } = useContext(AppContext);
  const serviceId = route.params?.serviceId;

  useEffect(() => {
    /**
     * Fetches services from the API and updates the state with the retrieved data.
     * 
     * This function makes a GET request to the API endpoint using the provided
     * `ipAddress`, `serviceId`, and `token`. If the response is an array, it updates
     * the `services` and `filteredServices` state with the data. If the response is
     * not an array, it logs an error and sets the `services` and `filteredServices`
     * state to empty arrays. In case of an error, it navigates to the 'Login' screen
     * if the error code is 401, logs the error, and sets the `services` and 
     * `filteredServices` state to empty arrays. The `loading` state is set to false
     * after the request completes or fails.
     * 
     * @async
     * @function fetchServices
     * @returns {Promise<void>} A promise that resolves when the fetch operation is complete.
     */
    const fetchServices = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/action/info/${serviceId}`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const data = await response.json();
        if (Array.isArray(data)) {
          setServices(data);
          setFilteredServices(data);
        } else {
          console.error('Unexpected API response:', data);
          setServices([]);
          setFilteredServices([]);
        }
        setLoading(false);
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching services:', error);
        setServices([]);
        setFilteredServices([]);
        setLoading(false);
      }
    };

    fetchServices();
  }, [ipAddress]);

  /**
   * Handles the search functionality by filtering the services based on the input text.
   * 
   * @param {string} text - The search text input by the user.
   * 
   * - If the search text is empty, it resets the filtered services to the original list of services.
   * - If the search text is not empty, it filters the services whose names include the search text (case-insensitive).
   */
  const handleSearch = (text: string) => {
    setSearch(text);
    if (text === '') {
      setFilteredServices(services);
    } else {
      setFilteredServices(
        services.filter(service =>
          service.name.toLowerCase().includes(text.toLowerCase()),
        ),
      );
    }
  };

  /**
   * Handles the press event for an action.
   * 
   * @param {any} action - The action object that was pressed.
   * @returns {void}
   * 
   * This function sets the selected action and its options (if any).
   * If the action has an `option` property, it logs the options to the console,
   * parses them, and updates the state with the parsed options.
   * If the action does not have an `option` property, it sets the selected action options to an empty object.
   */
  const handleActionPress = (action: any) => {
    setSelectedAction(action);
    if (action.option) {
      console.log('Action Options:', action.option);
      const parsedOptions = action.option;
      setSelectedActionOptions(parsedOptions);
    } else {
      setSelectedActionOptions({});
    }
  };

  /**
   * Handles the change of an option in the selected action options.
   *
   * @param {string} key - The key of the option to change.
   * @param {any} value - The new value of the option.
   * @param {any} type - The type of the option, used to determine if the value should be converted to a number.
   */
  const handleOptionChange = (key: string, value: any, type: any) => {
    setSelectedActionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

  /**
   * Handles the saving of selected action options.
   * Logs the selected action and its configured options to the console.
   * Navigates to the 'WorkflowReactionScreen' with the selected action's ID and options.
   * Resets the selected action and its options to their initial states.
   *
   * @function
   * @name handleSaveOptions
   */
  const handleSaveOptions = () => {
    console.log('Selected Action:', selectedAction);
    console.log('Configured Options:', selectedActionOptions);
    navigation.navigate('WorkflowReactionScreen', {
      actionId: selectedAction.id,
      actionOptions: selectedActionOptions,
    });
    setSelectedAction(null);
    setSelectedActionOptions({});
  };

  if (loading) {
    return (
      <View style={styles.container}>
        <ActivityIndicator size="large" color="#000" />
      </View>
    );
  }

  /**
   * Formats a given text string by inserting spaces before each uppercase letter,
   * capitalizing the first letter of the string, and trimming any leading or trailing whitespace.
   *
   * @param text - The input string to be formatted.
   * @returns The formatted string with spaces before uppercase letters, 
   *          the first letter capitalized, and no leading or trailing whitespace.
   */
  const formatText = (text: string): string => {
    return text
      .replace(/([A-Z])/g, ' $1')
      .replace(/^./, str => str.toUpperCase())
      .trim();
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Choose action</Text>
      {selectedAction ? (
        <View style={styles.optionsContainer}>
          <Text style={styles.optionTitle}>
            Configure Options for {selectedAction.name}
          </Text>
          {Object.keys(selectedActionOptions).map(key => (
            <View key={key} style={styles.optionRow}>
              <Text style={styles.optionLabel}>{key}</Text>
              <TextInput
                accessibilityLabel={key}
                accessibilityHint={`Enter value for ${key}`}
                placeholder={`Ex: ${String(selectedActionOptions[key])}`}
                placeholderTextColor="#999"
                style={styles.optionInput}
                // value={String(selectedActionOptions[key])}
                onChangeText={text =>
                  handleOptionChange(
                    key,
                    text,
                    typeof selectedActionOptions[key],
                  )
                }
                keyboardType={`${typeof selectedActionOptions[key] === 'number' ? "numeric" : "default"}`}
              />
            </View>
          ))}
          <TouchableOpacity
            accessibilityLabel="Save Button"
            accessibilityHint="Save the configured options"
            style={styles.saveButton}
            onPress={handleSaveOptions}>
            <Text style={styles.saveButtonText}>Save</Text>
          </TouchableOpacity>
          <TouchableOpacity
            accessibilityLabel="Back Button"
            accessibilityHint="Go back to the previous screen"
            style={styles.backButton}
            onPress={() => setSelectedAction(null)}>
            <Text style={styles.backButtonText}>Back</Text>
          </TouchableOpacity>
        </View>
      ) : (
        <>
          <TextInput
            accessibilityLabel="Search Bar"
            accessibilityHint="Search for services"
            style={styles.searchBar}
            placeholder="Search services"
            value={search}
            onChangeText={handleSearch}
          />
          <ScrollView contentContainerStyle={styles.servicesContainer}>
            {filteredServices?.map(service => (
              <TouchableOpacity
                accessibilityLabel={service.name}
                accessibilityHint={`Select ${service.name}`}
                key={service.id}
                style={[
                  styles.serviceBox,
                  { backgroundColor: service.service.color },
                ]}
                onPress={() => handleActionPress(service)}>
                <Text style={styles.serviceText}>
                  {formatText(service.name)}
                </Text>
              </TouchableOpacity>
            ))}
          </ScrollView>
          <TouchableOpacity
            accessibilityLabel="Back Button"
            accessibilityHint="Go back to the previous screen"
            style={styles.backButton}
            onPress={() => navigation.goBack()}>
            <Text style={styles.backButtonText}>Back</Text>
          </TouchableOpacity>
        </>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    padding: 20,
  },
  title: {
    fontSize: 32,
    fontWeight: 'bold',
    marginVertical: 20,
  },
  searchBar: {
    width: '100%',
    backgroundColor: '#f9f9f9',
    color: '#000',
    borderRadius: 10,
    padding: 10,
    fontSize: 18,
    marginBottom: 20,
    borderColor: '#ccc',
    borderWidth: 1,
  },
  servicesContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  serviceBox: {
    width: 140,
    height: 140,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
    margin: 10,
  },
  serviceText: {
    color: '#fff',
    fontSize: 18,
    fontWeight: 'bold',
    textAlign: 'center',
  },
  optionsContainer: {
    width: '100%',
  },
  optionTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
    textAlign: 'center',
  },
  optionRow: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 15,
    width: '100%',
  },
  optionLabel: {
    fontSize: 16,
    flex: 1,
  },
  optionInput: {
    flex: 2,
    backgroundColor: '#f0f0f0',
    borderRadius: 5,
    padding: 10,
    color: '#000',
    borderColor: '#ccc',
    borderWidth: 1,
  },
  saveButton: {
    backgroundColor: '#1DC000',
    borderRadius: 5,
    padding: 10,
    marginTop: 20,
    alignItems: 'center',
  },
  saveButtonText: {
    color: '#fff',
    fontSize: 18,
  },
  backButton: {
    marginTop: 20,
    width: '90%',
    height: 50,
    borderRadius: 25,
    borderWidth: 2,
    borderColor: '#000',
    justifyContent: 'center',
    alignItems: 'center',
  },
  backButtonText: {
    fontSize: 18,
    fontWeight: 'bold',
  },
});

export default SelectActionScreen;
