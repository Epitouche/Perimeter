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

type Props = NativeStackScreenProps<RootStackParamList, 'AddReactionScreen'>;

const AddReactionScreen: React.FC<Props> = ({ navigation, route }) => {
  const [connectedServices, setConnectedServices] = useState<string[]>([]);
  const [services, setServices] = useState<any[]>([]);
  const [filteredServices, setFilteredServices] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const { ipAddress, token } = useContext(AppContext);
  const { actionId, actionOptions } = route.params;

  useEffect(() => {
    const fetchServices = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/service/info/`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        const userResponse = await fetch(
          `http://${ipAddress}:8080/api/v1/user/info/all`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );

        const userData = await userResponse.json();
        const connected = userData.tokens.map(token => token.service.name);
        setConnectedServices(connected);
        console.log('Services:', response);
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

  if (loading) {
    return (
      <View style={styles.container}>
        <ActivityIndicator size="large" color="#000" />
      </View>
    );
  }

  const formatText = (text: string): string => {
    return text
      .replace(/([A-Z])/g, ' $1')
      .replace(/^./, str => str.toUpperCase())
      .trim();
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Choose reaction</Text>
      <TextInput
        accessibilityLabel='Search services'
        accessibilityHint='Enter text to search for services'
        style={styles.searchBar}
        placeholder="Search services"
        placeholderTextColor="#bbbbbb"
        value={search}
        onChangeText={handleSearch}
      />
      <ScrollView contentContainerStyle={styles.servicesContainer}>
        {filteredServices?.map(service => (
          <TouchableOpacity
            accessibilityLabel={`Select reaction for ${service.name}`}
            accessibilityHint={`Double tap to select the reaction for ${service.name}`}
            key={service.id}
            style={[
              styles.serviceBox,
              {
                backgroundColor:
                  connectedServices.includes(service.name) || !service.oauth
                    ? service.color
                    : '#d3d3d3',
              },
            ]}
            onPress={() =>
              navigation.navigate('SelectReactionScreen', {
                actionId: actionId,
                actionOptions: actionOptions,
                serviceId: service.id,
              })
            }
            disabled={
              !(connectedServices.includes(service.name) || !service.oauth)
            }>
            <Text style={styles.serviceText}>{formatText(service.name)}</Text>
          </TouchableOpacity>
        ))}
      </ScrollView>
      <TouchableOpacity
        accessibilityLabel='Back button'
        accessibilityHint='Double tap to go back to the previous screen'
        style={styles.backButton}
        onPress={() => navigation.goBack()}>
        <Text style={styles.backButtonText}>Back</Text>
      </TouchableOpacity>
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
    backgroundColor: '#add8e6',
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

export default AddReactionScreen;
