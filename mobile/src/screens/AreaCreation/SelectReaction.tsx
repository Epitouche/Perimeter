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

type Props = NativeStackScreenProps<RootStackParamList, 'SelectReactionScreen'>;

/**
 * SelectReactionScreen component allows users to select and configure a reaction for a specific service.
 *
 * @param {object} props - The props object.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @param {object} props.route - The route object provided by React Navigation.
 * @param {object} props.route.params - The parameters passed to the route.
 * @param {string} props.route.params.actionId - The ID of the action.
 * @param {object} props.route.params.actionOptions - The options for the action.
 * @param {string} props.route.params.serviceId - The ID of the service.
 *
 * @returns {JSX.Element} The rendered component.
 *
 * @component
 *
 * @example
 * return (
 *   <SelectReactionScreen navigation={navigation} route={route} />
 * )
 *
 * @remarks
 * This component fetches the available reactions for a given service and allows the user to search, select, and configure a reaction.
 * It handles the loading state, search functionality, and navigation to the next screen with the selected reaction and its options.
 *
 * @function
 * @name SelectReactionScreen
 */
const SelectReactionScreen: React.FC<Props> = ({ navigation, route }) => {
  const [services, setServices] = useState<any[]>([]);
  const [filteredServices, setFilteredServices] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [selectedReaction, setSelectedReaction] = useState<any | null>(null);
  const [selectedReactionOptions, setSelectedReactionOptions] = useState<{
    [key: string]: any;
  }>({});
  const { ipAddress, token } = useContext(AppContext);
  const { actionId, actionOptions, serviceId } = route.params;

  useEffect(() => {
    /**
     * Fetches services from the API and updates the state with the fetched data.
     *
     * This function sends a GET request to the API endpoint to retrieve information
     * about services. It handles the response by updating the `services` and
     * `filteredServices` state variables. If the response is not an array, it logs
     * an error and sets the state variables to empty arrays. It also handles errors
     * during the fetch operation, including navigating to the login screen if a 401
     * Unauthorized error occurs.
     *
     * @async
     * @function fetchServices
     * @returns {Promise<void>} A promise that resolves when the fetch operation is complete.
     * @throws Will log an error message if the fetch operation fails.
     */
    const fetchServices = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/reaction/info/${serviceId}`,
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
        if ((error as any).code === 401) {
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
   * Handles the search functionality by updating the search text and filtering the services.
   *
   * @param {string} text - The search text entered by the user.
   *
   * The function performs the following actions:
   * - Updates the search state with the provided text.
   * - If the search text is empty, it resets the filtered services to the original list of services.
   * - Otherwise, it filters the services based on whether their names include the search text (case-insensitive).
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
   * Handles the press action for a given reaction.
   *
   * @param {any} action - The action object containing the reaction details.
   * @returns {void}
   *
   * This function sets the selected reaction and its options if available.
   * If the action object contains an `option` property, it parses and sets the options.
   * Otherwise, it sets the selected reaction options to an empty object.
   */
  const handleActionPress = (action: any) => {
    setSelectedReaction(action);
    if (action.option) {
      const parsedOptions = action.option;
      setSelectedReactionOptions(parsedOptions);
    } else {
      setSelectedReactionOptions({});
    }
  };

  /**
   * Handles the change of an option in the reaction selection.
   *
   * @param {string} key - The key of the option being changed.
   * @param {any} value - The new value of the option.
   * @param {any} type - The type of the option, used to determine if the value should be parsed as a number.
   */
  const handleOptionChange = (key: string, value: any, type: any) => {
    setSelectedReactionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? parseFloat(value) : value,
    }));
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
   * @returns The formatted string with spaces before uppercase letters and the first letter capitalized.
   */
  const formatText = (text: string): string => {
    return text
      .replace(/([A-Z])/g, ' $1')
      .replace(/^./, str => str.toUpperCase())
      .trim();
  };

  return (
    <View style={styles.container}>
      <Text
        style={styles.title}
        accessibilityLabel="Choose reaction title"
        accessibilityHint="Title of the screen">
        Choose reaction
      </Text>
      {selectedReaction ? (
        <View style={styles.optionsContainer}>
          <Text
            style={styles.optionTitle}
            accessibilityLabel={`Configure Options for ${selectedReaction.name}`}
            accessibilityHint="Title for configuring options">
            Configure Options for {selectedReaction.name}
          </Text>
          {Object.keys(selectedReactionOptions).map(key => (
            <View key={key} style={styles.optionRow}>
              <Text
                style={styles.optionLabel}
                accessibilityLabel={`Option label ${key}`}
                accessibilityHint={`Label for the option ${key}`}>
                {key}
              </Text>
              <TextInput
                style={styles.optionInput}
                placeholder={`Ex: ${String(selectedReactionOptions[key])}`}
                placeholderTextColor="#999"
                // value={String(selectedReactionOptions[key])}
                onChangeText={text =>
                  handleOptionChange(
                    key,
                    text,
                    typeof selectedReactionOptions[key],
                  )
                }
                keyboardType={`${
                  typeof selectedReactionOptions[key] == 'number'
                    ? 'numeric'
                    : 'default'
                }`}
                accessibilityLabel={`Option input ${key}`}
                accessibilityHint={`Input field for the option ${key}`}
              />
            </View>
          ))}
          <TouchableOpacity
            style={styles.saveButton}
            onPress={() =>
              navigation.navigate('ValidateAreaScreen', {
                actionId,
                actionOptions,
                reactionId: selectedReaction.id,
                reactionOptions: selectedReactionOptions,
              })
            }
            accessibilityLabel="Save button"
            accessibilityHint="Press to save the selected reaction and options">
            <Text style={styles.saveButtonText}>Save</Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.backButton}
            onPress={() => setSelectedReaction(null)}
            accessibilityLabel="Back button"
            accessibilityHint="Press to go back to the previous screen">
            <Text style={styles.backButtonText}>Back</Text>
          </TouchableOpacity>
        </View>
      ) : (
        <>
          <TextInput
            style={styles.searchBar}
            placeholder="Search services"
            value={search}
            onChangeText={handleSearch}
            accessibilityLabel="Search bar"
            accessibilityHint="Input field to search for services"
          />
          <ScrollView contentContainerStyle={styles.servicesContainer}>
            {filteredServices?.map(service => (
              <TouchableOpacity
                key={service.id}
                style={[
                  styles.serviceBox,
                  { backgroundColor: service.service.color },
                ]}
                onPress={() => handleActionPress(service)}
                accessibilityLabel={`Service ${service.name}`}
                accessibilityHint={`Press to select the service ${service.name}`}>
                <Text style={styles.serviceText}>
                  {formatText(service.name)}
                </Text>
              </TouchableOpacity>
            ))}
          </ScrollView>
          <TouchableOpacity
            style={styles.backButton}
            onPress={() => navigation.goBack()}
            accessibilityLabel="Back button"
            accessibilityHint="Press to go back to the previous screen">
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
    backgroundColor: '#f0f0f0',
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

export default SelectReactionScreen;
