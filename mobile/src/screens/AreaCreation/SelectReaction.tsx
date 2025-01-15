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

  const handleActionPress = (action: any) => {
    setSelectedReaction(action);
    if (action.option) {
      const parsedOptions = action.option;
      setSelectedReactionOptions(parsedOptions);
    } else {
      setSelectedReactionOptions({});
    }
  };

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

  const formatText = (text: string): string => {
    return text
      .replace(/([A-Z])/g, ' $1')
      .replace(/^./, str => str.toUpperCase())
      .trim();
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title} accessibilityLabel="Choose reaction title">Choose reaction</Text>
      {selectedReaction ? (
        <View style={styles.optionsContainer}>
          <Text style={styles.optionTitle} accessibilityLabel={`Configure Options for ${selectedReaction.name}`}>
            Configure Options for {selectedReaction.name}
          </Text>
          {Object.keys(selectedReactionOptions).map(key => (
            <View key={key} style={styles.optionRow}>
              <Text style={styles.optionLabel} accessibilityLabel={`Option label ${key}`}>{key}</Text>
              <TextInput
                style={styles.optionInput}
                value={String(selectedReactionOptions[key])}
                onChangeText={text =>
                  handleOptionChange(
                    key,
                    text,
                    typeof selectedReactionOptions[key],
                  )
                }
                keyboardType="default"
                accessibilityLabel={`Option input ${key}`}
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
            accessibilityLabel="Save button">
            <Text style={styles.saveButtonText}>Save</Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.backButton}
            onPress={() => setSelectedReaction(null)}
            accessibilityLabel="Back button">
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
          />
          <ScrollView contentContainerStyle={styles.servicesContainer}></ScrollView>
            {filteredServices?.map(service => (
              <TouchableOpacity
                key={service.id}
                style={styles.serviceBox}
                onPress={() => handleActionPress(service)}
                accessibilityLabel={`Service ${service.name}`}>
                <Text style={styles.serviceText}>
                  {formatText(service.name)}
                </Text>
              </TouchableOpacity>
            ))}
          </ScrollView>
          <TouchableOpacity
            style={styles.backButton}
            onPress={() => navigation.goBack()}
            accessibilityLabel="Back button">
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

export default SelectReactionScreen;
