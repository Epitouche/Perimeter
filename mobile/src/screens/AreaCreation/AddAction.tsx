import React, {useState, useEffect, useContext} from 'react';
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  ScrollView,
  ActivityIndicator,
} from 'react-native';
import {NativeStackScreenProps} from '@react-navigation/native-stack';
import {RootStackParamList} from '../../../App';
import {AppContext} from '../../context/AppContext';

type Props = NativeStackScreenProps<RootStackParamList, 'AddActionScreen'>;

const AddActionScreen: React.FC<Props> = ({navigation}) => {
  const [services, setServices] = useState<any[]>([]);
  const [filteredServices, setFilteredServices] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const {ipAddress, token} = useContext(AppContext);

  useEffect(() => {
    // Fetch services from API
    const fetchServices = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/service/info`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
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

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Add action</Text>
      <TextInput
        style={styles.searchBar}
        placeholder="Search services"
        value={search}
        onChangeText={handleSearch}
      />
      <ScrollView contentContainerStyle={styles.servicesContainer}>
        {filteredServices?.map(service => (
          <TouchableOpacity
            key={service.id}
            style={styles.serviceBox}
            onPress={() =>
              navigation.navigate('SelectActionScreen', {
                serviceId: service.id,
              })
            }>
            <Text style={styles.serviceText}>{service.name}</Text>
          </TouchableOpacity>
        ))}
      </ScrollView>
      <TouchableOpacity
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

export default AddActionScreen;