import React, { useContext, useEffect } from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from '../../Navigation/navigate';
import { AppContext } from '../../context/AppContext';

type Props = NativeStackScreenProps<
  RootStackParamList,
  'WorkflowReactionScreen'
>;

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
    const getService = async () => {
      try {
        const response = await fetch(
          `http://${ipAddress}:8080/api/v1/action/info/${actionId}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          },
        );

        let data = await response.json();
        setName(data[0].name);
        setService(data[0].service);
      } catch (error) {
        if (error.code === 401) {
          navigation.navigate('Login');
        }
        console.error('Error fetching service:', error);
      }
    };
    getService();
  }, [ipAddress, token]);

  return (
    <View style={styles.container}>
      <Text style={styles.title} accessibilityLabel="Add Area Title">Add Area</Text>
      <View style={[styles.actionBox, { backgroundColor: service?.color }]}>
      <Text style={styles.boxText} accessibilityLabel={`Service Name: ${name}`}>{name}</Text>
      </View>
      <View style={styles.line} accessibilityLabel="Separator Line" />
      <View style={styles.actionBox}>
      <Text style={styles.boxText} accessibilityLabel="Reaction Label">Reaction</Text>
      <TouchableOpacity
        style={styles.addButton}
        onPress={() =>
        navigation.navigate('AddReactionScreen', {
          actionId,
          actionOptions,
        })
        }
        accessibilityLabel="Add Reaction Button">
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
