import React, { useEffect, useContext, useState } from 'react';
import {
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  ActivityIndicator,
} from 'react-native';
import { SvgFromUri } from 'react-native-svg';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import BottomNavBar from './NavBar';
import { RootStackParamList } from '../Navigation/navigate';
import { AppContext } from '../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'AreaView'>;

/**
 * AreasScreen component fetches and displays a list of areas.
 *
 * @param {Props} props - The props object containing navigation.
 * @returns {JSX.Element} The rendered component.
 *
 * @component
 * @example
 * return (
 *   <AreasScreen navigation={navigation} />
 * )
 */
const AreasScreen = ({ navigation }: Props) => {
  const [areas, setAreas] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const { ipAddress, token } = useContext(AppContext);

  useEffect(() => {
    /**
     * Fetches areas from the API.
     *
     * This function sends a GET request to the API to retrieve area data.
     * If the request is successful, the data is stored in the state using `setAreas`.
     * If the request fails with a 401 status code, the user is navigated to the 'Login' screen.
     * Regardless of the outcome, the loading state is set to false after the request completes.
     *
     * @async
     * @function fetchAreas
     * @throws Will throw an error if the fetch request fails for reasons other than a 401 status code.
     */
    const fetchAreas = async () => {
      try {
        const response = await fetch(`http://${ipAddress}:8080/api/v1/area/`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        const data = await response.json();
        setAreas(data);
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching areas:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchAreas();
  }, []);

  if (loading) {
    return (
      <View style={styles.loadingContainer}>
        <ActivityIndicator size="large" color="#001DDA" />
        <Text>Loading AREAs...</Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Text style={styles.header}>My AREAs</Text>
      <View style={styles.areasContainer}>
        {areas.map((area, index) => (
          <TouchableOpacity
            key={index}
            style={[
              styles.areaBox,
              { backgroundColor: area.action.service.color },
            ]}
            onPress={() => navigation.navigate('AreaDetails', { area })}
            accessibilityHint={`Navigates to details of ${area.action.name} and ${area.reaction.name}`}>
            <Text
              style={
                styles.areaText
              }>{`${area.action.name} ~ ${area.reaction.name}`}</Text>
            <View style={styles.iconsContainer}>
              <SvgFromUri
                uri={area.action.service.icon}
                width={25}
                height={25}
              />
              <SvgFromUri
                uri={area.reaction.service.icon}
                width={25}
                height={25}
              />
            </View>
          </TouchableOpacity>
        ))}
      </View>
      <BottomNavBar navigation={navigation} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    padding: 16,
  },
  header: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 16,
  },
  areasContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-around',
  },
  areaBox: {
    width: 150,
    height: 150,
    borderRadius: 16,
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 16,
  },
  areaText: {
    color: 'white',
    fontSize: 16,
    fontWeight: 'bold',
    textAlign: 'center',
    marginBottom: 8,
  },
  iconsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    width: '80%',
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default AreasScreen;
