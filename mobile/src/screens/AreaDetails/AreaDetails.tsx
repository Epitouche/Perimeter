import React, { useState, useContext, useEffect } from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  FlatList,
  ScrollView,
} from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';
import { styles } from './StylesAreaDetails';
import AreaSections from './AreaSections';
import ActionsSections from './ActionsSections';
import ReactionsSections from './ReactionsSections';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaDetails'>;

/**
 * AreaDetailsScreen component displays the details of a specific area, including its results.
 * It fetches area results from an API and displays them in a table format.
 *
 * @param {object} props - The component props.
 * @param {object} props.navigation - The navigation object provided by React Navigation.
 * @param {object} props.route - The route object provided by React Navigation.
 * @param {object} props.route.params - The parameters passed to the route.
 * @param {object} props.route.params.area - The area object containing the area details.
 *
 * @returns {JSX.Element} The rendered component.
 *
 * @example
 * <AreaDetailsScreen navigation={navigation} route={route} />
 *
 * @component
 */
const AreaDetailsScreen = ({ navigation, route }: Props) => {
  const { area } = route.params;
  const { ipAddress, token } = useContext(AppContext);
  const [areaResults, setAreaResults] = useState([
    { created_at: '', result: '' },
  ]);

  useEffect(() => {
    /**
     * Fetches the area results from the server.
     *
     * This function sends a GET request to the server to retrieve the results
     * for a specific area. The request includes an authorization token in the
     * headers. If the request is successful, the area results are stored in the
     * state and logged to the console. If there is an error during the fetch
     * operation, it is caught and logged to the console.
     *
     * @async
     * @function fetchAreaResults
     * @returns {Promise<void>} A promise that resolves when the fetch operation is complete.
     * @throws Will log an error message to the console if the fetch operation fails.
     */
    const fetchAreaResults = async () => {
      console.log('Fetching area results');
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/area-result/${area.id}`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        if (response.ok) {
          const body = await response.json();
          setAreaResults(body);
        }
      } catch (error) {
        if ((error as any).response.status === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching area results:', error);
      }
    };

    fetchAreaResults();
  }, []);

  /**
   * Renders an item in a list with creation date and result.
   *
   * @param {Object} params - The parameters for the renderItem function.
   * @param {Object} params.item - The item to be rendered.
   * @param {string} params.item.created_at - The creation date and time of the result.
   * @param {string} params.item.result - The result of the area.
   * @returns {JSX.Element} A view containing the creation date and result.
   */
  const renderItem = ({
    item,
  }: {
    item: { created_at: string; result: string };
  }) => (
    <View style={styles.row}>
      <Text
        style={styles.cell}
        accessibilityLabel={`Created at: ${item.created_at.substring(0, 19)}`}
        accessibilityHint="Displays the creation date and time of the result">
        {item.created_at.substring(0, 19)}
      </Text>
      <Text
        style={styles.cell}
        accessibilityLabel={`Result: ${item.result}`}
        accessibilityHint="Displays the result of the area">
        {item.result}
      </Text>
    </View>
  );

  return (
    <ScrollView style={{ backgroundColor: 'white' }}>
      <View style={styles.container}>
        <Text
          style={styles.header}
          accessibilityLabel="Area Details Header"
          accessibilityHint="Header for the area details section">
          Area Details
        </Text>

        {/* Area Section */}
        {AreaSections({ navigation, route })}

        {/* Action Section */}
        {ActionsSections({ navigation, route })}

        {/* Reaction Section */}
        {ReactionsSections({ navigation, route })}

        {/* Area Results */}
        <Text
          style={[styles.header, { marginTop: 16 }]}
          accessibilityLabel="Area Results Header"
          accessibilityHint="Header for the area results section">
          Area Results
        </Text>
        <View
          style={[
            { borderColor: 'black', borderWidth: 1, borderRadius: 10, flex: 1 },
          ]}>
          <View style={[styles.row, { backgroundColor: 'white' }]}>
            <Text
              style={styles.cell}
              accessibilityLabel="Created At Column"
              accessibilityHint="Column header for creation date and time">
              created_at
            </Text>
            <Text
              style={styles.cell}
              accessibilityLabel="Result Column"
              accessibilityHint="Column header for results">
              result
            </Text>
          </View>
          <FlatList
            data={areaResults}
            renderItem={renderItem}
            keyExtractor={item => item.created_at}
            style={{ flex: 1 }}
          />
        </View>

        <TouchableOpacity
          onPress={() => navigation.navigate('AreaView')}
          accessibilityLabel="Back Button"
          accessibilityHint="Navigates back to the area view screen">
          <View style={{ alignItems: 'flex-end', justifyContent: 'flex-end' }}>
            <Text
              style={[
                styles.cancelButton,
                { color: '#E60000', width: '20%', margin: 10 },
              ]}>
              Back
            </Text>
          </View>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
};

export default AreaDetailsScreen;
