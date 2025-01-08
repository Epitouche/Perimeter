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

  const handleOptionChange = (key: string, value: any, type: any) => {
    setSelectedActionOptions(prev => ({
      ...prev,
      [key]: type === 'number' ? Number(value) : value,
    }));
  };

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
                style={styles.optionInput}
                value={String(selectedActionOptions[key])}
                onChangeText={text =>
                  handleOptionChange(
                    key,
                    text,
                    typeof selectedActionOptions[key],
                  )
                }
                keyboardType="default" // Adjust as needed
              />
            </View>
          ))}
          <TouchableOpacity
            style={styles.saveButton}
            onPress={handleSaveOptions}>
            <Text style={styles.saveButtonText}>Save</Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.backButton}
            onPress={() => setSelectedAction(null)}>
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
          />
          <ScrollView contentContainerStyle={styles.servicesContainer}>
            {filteredServices?.map(service => (
              <TouchableOpacity
                key={service.id}
                style={styles.serviceBox}
                onPress={() => handleActionPress(service)}>
                <Text style={styles.serviceText}>
                  {formatText(service.name)}
                </Text>
              </TouchableOpacity>
            ))}
          </ScrollView>
          <TouchableOpacity
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
    borderColor: '#ccc',
    borderWidth: 1,
  },
  saveButton: {
    backgroundColor: '#4CAF50',
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
